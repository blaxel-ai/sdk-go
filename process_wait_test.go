package blaxel_test

import (
	"context"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	blaxel "github.com/blaxel-ai/sdk-go"
	"github.com/blaxel-ai/sdk-go/option"
)

func processServer(t *testing.T, handler http.HandlerFunc) *blaxel.SandboxInstance {
	t.Helper()
	server := httptest.NewServer(handler)
	t.Cleanup(server.Close)
	client := blaxel.NewClient(option.WithAPIKey("test"))
	return client.Sandboxes.FromSession(blaxel.SessionWithToken{Name: "test", URL: server.URL, Token: "test"}, option.WithHTTPClient(server.Client()), option.WithMaxRetries(0))
}
func processJSON(w http.ResponseWriter, status string) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"name":"original","pid":"123","status":%q,"exitCode":7,"logs":"output"}`, status)
}
func TestProcessWaitTerminalStates(t *testing.T) {
	for _, status := range []string{"completed", "failed", "killed", "stopped"} {
		t.Run(status, func(t *testing.T) {
			sandbox := processServer(t, func(w http.ResponseWriter, r *http.Request) { processJSON(w, status) })
			result, err := sandbox.Process.Wait(context.Background(), "original", time.Second, time.Millisecond)
			if err != nil || string(result.Status) != status || result.ExitCode != 7 {
				t.Fatalf("result=%+v err=%v", result, err)
			}
		})
	}
}
func TestProcessWaitRecoversStatusFailures(t *testing.T) {
	for _, status := range []int{408, 429, 500, 502, 503, 504} {
		t.Run(fmt.Sprint(status), func(t *testing.T) {
			var calls atomic.Int32
			sandbox := processServer(t, func(w http.ResponseWriter, r *http.Request) {
				switch calls.Add(1) {
				case 1:
					processJSON(w, "running")
				case 2:
					http.Error(w, "temporary", status)
				default:
					processJSON(w, "completed")
				}
			})
			result, err := sandbox.Process.Wait(context.Background(), "original", time.Second, time.Millisecond)
			if err != nil || result.Status != "completed" || calls.Load() != 3 {
				t.Fatalf("result=%+v err=%v calls=%d", result, err, calls.Load())
			}
		})
	}
}
func TestProcessWaitPreservesPermanentError(t *testing.T) {
	for _, status := range []int{400, 401, 403, 404, 409, 501} {
		t.Run(fmt.Sprint(status), func(t *testing.T) {
			var calls atomic.Int32
			sandbox := processServer(t, func(w http.ResponseWriter, r *http.Request) { calls.Add(1); http.Error(w, "permanent", status) })
			result, err := sandbox.Process.Wait(context.Background(), "original", time.Second, time.Millisecond)
			apiErr, ok := err.(*blaxel.Error)
			if result != nil || !ok || apiErr.StatusCode != status || calls.Load() != 1 {
				t.Fatalf("result=%+v err=%v calls=%d", result, err, calls.Load())
			}
		})
	}
}
func TestProcessWaitUnknownStatus(t *testing.T) {
	sandbox := processServer(t, func(w http.ResponseWriter, r *http.Request) { processJSON(w, "unknown") })
	result, err := sandbox.Process.Wait(context.Background(), "original", time.Second, time.Millisecond)
	if result != nil || err == nil || !strings.Contains(err.Error(), "unknown process status") {
		t.Fatalf("result=%+v err=%v", result, err)
	}
}
func TestProcessWaitDeadlineBoundsRequest(t *testing.T) {
	sandbox := processServer(t, func(w http.ResponseWriter, r *http.Request) { <-r.Context().Done() })
	start := time.Now()
	_, err := sandbox.Process.Wait(context.Background(), "original", 30*time.Millisecond, time.Second)
	if !errors.Is(err, context.DeadlineExceeded) || time.Since(start) > 500*time.Millisecond {
		t.Fatalf("err=%v elapsed=%s", err, time.Since(start))
	}
}
func TestProcessWaitDeadlineKeepsLastError(t *testing.T) {
	sandbox := processServer(t, func(w http.ResponseWriter, r *http.Request) { http.Error(w, "unavailable", 503) })
	_, err := sandbox.Process.Wait(context.Background(), "original", 30*time.Millisecond, time.Millisecond)
	var apiErr *blaxel.Error
	if !errors.Is(err, context.DeadlineExceeded) || !errors.As(err, &apiErr) || apiErr.StatusCode != 503 {
		t.Fatal(err)
	}
}
func TestProcessWaitCancellationDoesNotKill(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	sandbox := processServer(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			t.Errorf("unexpected method: %s", r.Method)
		}
		processJSON(w, "running")
		cancel()
	})
	_, err := sandbox.Process.Wait(ctx, "original", time.Second, time.Second)
	if err != context.Canceled {
		t.Fatalf("error=%v", err)
	}
}
func TestProcessCreationDoesNotReplay(t *testing.T) {
	var calls atomic.Int32
	sandbox := processServer(t, func(w http.ResponseWriter, r *http.Request) {
		calls.Add(1)
		conn, _, err := w.(http.Hijacker).Hijack()
		if err == nil {
			_ = conn.Close()
		}
	})
	_, err := sandbox.Process.New(context.Background(), blaxel.ProcessRequestParam{Name: blaxel.String("original"), Command: "side effect"}, option.WithMaxRetries(5))
	if err == nil || calls.Load() != 1 {
		t.Fatalf("error=%v calls=%d", err, calls.Load())
	}
}
func TestProcessStreamMissingResult(t *testing.T) {
	sandbox := processServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/x-ndjson")
		io.WriteString(w, "{\"type\":\"stdout\",\"data\":\"partial\"}\n")
	})
	result, err := sandbox.Process.ExecWithStreaming(context.Background(), blaxel.ProcessRequestParam{Command: "echo output"}, blaxel.ProcessStreamOptions{})
	if result != nil || err == nil || !strings.Contains(err.Error(), "no result received") {
		t.Fatalf("result=%+v error=%v", result, err)
	}
}
func TestProcessAndWatchStreamErrors(t *testing.T) {
	for _, watch := range []bool{false, true} {
		for _, readError := range []bool{false, true} {
			t.Run(fmt.Sprintf("watch=%t/read=%t", watch, readError), func(t *testing.T) {
				sandbox := processServer(t, func(w http.ResponseWriter, r *http.Request) {
					if readError {
						w.Header().Set("Content-Length", "1000")
						io.WriteString(w, "\n")
					} else {
						http.Error(w, "missing", 404)
					}
				})
				var stream *blaxel.StreamControl
				if watch {
					stream = sandbox.FS.Watch(context.Background(), "tmp", func(blaxel.WatchEvent) {}, nil)
				} else {
					stream = sandbox.Process.StreamLogs(context.Background(), "original", blaxel.ProcessStreamOptions{})
				}
				err := stream.Err()
				if err == nil || (readError && !errors.Is(err, io.ErrUnexpectedEOF)) {
					t.Fatalf("error=%v", err)
				}
			})
		}
	}
}
func TestProcessAndWatchStreamCloseAndCancellation(t *testing.T) {
	for _, watch := range []bool{false, true} {
		for _, explicitClose := range []bool{false, true} {
			t.Run(fmt.Sprintf("watch=%t/close=%t", watch, explicitClose), func(t *testing.T) {
				started := make(chan struct{})
				sandbox := processServer(t, func(w http.ResponseWriter, r *http.Request) {
					w.WriteHeader(200)
					w.(http.Flusher).Flush()
					close(started)
					<-r.Context().Done()
				})
				ctx, cancel := context.WithCancel(context.Background())
				defer cancel()
				var stream *blaxel.StreamControl
				if watch {
					stream = sandbox.FS.Watch(ctx, "tmp", func(blaxel.WatchEvent) {}, nil)
				} else {
					stream = sandbox.Process.StreamLogs(ctx, "original", blaxel.ProcessStreamOptions{})
				}
				<-started
				if explicitClose {
					stream.Close()
				} else {
					cancel()
				}
				err := stream.Err()
				if (explicitClose && err != nil) || (!explicitClose && !errors.Is(err, context.Canceled)) {
					t.Fatalf("error=%v", err)
				}
			})
		}
	}
}

func TestProcessWaitPermanentTransportError(t *testing.T) {
	sentinel := errors.New("permanent transport setup error")
	var calls atomic.Int32
	client := blaxel.NewClient()
	sandbox := client.Sandboxes.FromSession(blaxel.SessionWithToken{Name: "test", URL: "https://sandbox.test", Token: "test"}, option.WithHTTPClient(&http.Client{Transport: &closureTransport{fn: func(r *http.Request) (*http.Response, error) { calls.Add(1); return nil, sentinel }}}))
	_, err := sandbox.Process.Wait(context.Background(), "original", time.Second, time.Millisecond)
	if !errors.Is(err, sentinel) || errors.Is(err, context.DeadlineExceeded) || calls.Load() != 1 {
		t.Fatalf("error=%v calls=%d", err, calls.Load())
	}
}
