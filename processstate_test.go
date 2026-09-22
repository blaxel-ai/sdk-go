package blaxel_test

import (
	"context"
	"encoding/json"
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

func processTestSandbox(t *testing.T, handler http.HandlerFunc) *blaxel.SandboxInstance {
	t.Helper()
	server := httptest.NewServer(handler)
	t.Cleanup(server.Close)
	client := blaxel.NewClient(option.WithAPIKey("test"))
	return client.Sandboxes.FromSession(blaxel.SessionWithToken{Name: "test", URL: server.URL, Token: "test"}, option.WithHTTPClient(server.Client()))
}
func writeProcess(w http.ResponseWriter, status string) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintf(w, `{"name":"original","pid":"123","status":%q,"exitCode":7,"logs":"preserved output"}`, status)
}
func requireProcessError(t *testing.T, err error) *blaxel.ProcessError {
	t.Helper()
	var p *blaxel.ProcessError
	if !errors.As(err, &p) {
		t.Fatalf("expected ProcessError, got %T: %v", err, err)
	}
	return p
}

func TestProcessWaitTerminalStates(t *testing.T) {
	for _, status := range []string{"completed", "failed", "killed", "stopped"} {
		t.Run(status, func(t *testing.T) {
			var calls atomic.Int32
			sandbox := processTestSandbox(t, func(w http.ResponseWriter, r *http.Request) { calls.Add(1); writeProcess(w, status) })
			result, err := sandbox.Process.Wait(context.Background(), "original", time.Second, time.Millisecond)
			if err != nil || string(result.Status) != status || result.ExitCode != 7 || result.Logs != "preserved output" || calls.Load() != 1 {
				t.Fatalf("result=%+v err=%v calls=%d", result, err, calls.Load())
			}
		})
	}
}
func TestProcessWaitRecoversTransientStatusFailure(t *testing.T) {
	for _, status := range []int{408, 429, 500, 502, 503, 504} {
		t.Run(fmt.Sprint(status), func(t *testing.T) {
			var calls atomic.Int32
			sandbox := processTestSandbox(t, func(w http.ResponseWriter, r *http.Request) {
				switch calls.Add(1) {
				case 1:
					writeProcess(w, "running")
				case 2:
					http.Error(w, "temporary", status)
				default:
					writeProcess(w, "completed")
				}
			})
			result, err := sandbox.Process.Wait(context.Background(), "original", time.Second, time.Millisecond)
			if err != nil || result.Status != "completed" || calls.Load() != 3 {
				t.Fatalf("result=%+v err=%v calls=%d", result, err, calls.Load())
			}
		})
	}
}
func TestProcessWaitPermanentFailurePreservesLastObservation(t *testing.T) {
	for _, status := range []int{400, 401, 403, 404, 409, 501} {
		t.Run(fmt.Sprint(status), func(t *testing.T) {
			var calls atomic.Int32
			sandbox := processTestSandbox(t, func(w http.ResponseWriter, r *http.Request) {
				if calls.Add(1) == 1 {
					writeProcess(w, "running")
				} else {
					http.Error(w, "unavailable", status)
				}
			})
			result, err := sandbox.Process.Wait(context.Background(), "original", time.Second, time.Millisecond)
			p := requireProcessError(t, err)
			var apiErr *blaxel.Error
			if result != nil || p.Identifier != "original" || p.LastKnownProcess == nil || p.LastKnownProcess.Status != "running" || !errors.As(err, &apiErr) || apiErr.StatusCode != status || calls.Load() != 2 {
				t.Fatalf("processError=%+v calls=%d", p, calls.Load())
			}
		})
	}
}
func TestProcessWaitRejectsUnknownStatus(t *testing.T) {
	sandbox := processTestSandbox(t, func(w http.ResponseWriter, r *http.Request) { writeProcess(w, "unknown") })
	result, err := sandbox.Process.Wait(context.Background(), "original", time.Second, time.Millisecond)
	p := requireProcessError(t, err)
	if result != nil || p.LastKnownProcess.Status != "unknown" {
		t.Fatalf("result=%+v error=%+v", result, p)
	}
}
func TestProcessWaitDeadlineBoundsHTTPRequest(t *testing.T) {
	sandbox := processTestSandbox(t, func(w http.ResponseWriter, r *http.Request) { <-r.Context().Done() })
	start := time.Now()
	_, err := sandbox.Process.Wait(context.Background(), "original", 30*time.Millisecond, time.Second)
	if !errors.Is(err, context.DeadlineExceeded) || time.Since(start) > 500*time.Millisecond {
		t.Fatalf("error=%v elapsed=%s", err, time.Since(start))
	}
	requireProcessError(t, err)
}
func TestProcessWaitCancellationDoesNotStopCommand(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	var deletes atomic.Int32
	sandbox := processTestSandbox(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "DELETE" {
			deletes.Add(1)
		}
		writeProcess(w, "running")
	})
	// Cancel during the polling delay, after the running observation is received.
	time.AfterFunc(25*time.Millisecond, cancel)
	_, err := sandbox.Process.Wait(ctx, "original", time.Second, time.Second)
	p := requireProcessError(t, err)
	if !errors.Is(err, context.Canceled) || p.LastKnownProcess == nil || deletes.Load() != 0 {
		t.Fatalf("error=%+v deletes=%d", p, deletes.Load())
	}
}
func TestProcessWaitTransientFailureDeadlineKeepsCause(t *testing.T) {
	sandbox := processTestSandbox(t, func(w http.ResponseWriter, r *http.Request) { http.Error(w, "temporary", 503) })
	_, err := sandbox.Process.Wait(context.Background(), "original", 20*time.Millisecond, time.Millisecond)
	var apiErr *blaxel.Error
	if !errors.Is(err, context.DeadlineExceeded) || !errors.As(err, &apiErr) || apiErr.StatusCode != 503 {
		t.Fatalf("error=%v", err)
	}
}
func TestProcessCreationIsNotReplayed(t *testing.T) {
	for _, stream := range []bool{false, true} {
		t.Run(fmt.Sprint(stream), func(t *testing.T) {
			var calls atomic.Int32
			var submitted atomic.Value
			sandbox := processTestSandbox(t, func(w http.ResponseWriter, r *http.Request) {
				calls.Add(1)
				var body struct {
					Name string `json:"name"`
				}
				_ = json.NewDecoder(r.Body).Decode(&body)
				submitted.Store(body.Name)
				conn, _, err := w.(http.Hijacker).Hijack()
				if err == nil {
					_ = conn.Close()
				}
			})
			var err error
			if stream {
				_, err = sandbox.Process.ExecWithStreaming(context.Background(), blaxel.ProcessRequestParam{Command: "side effect"}, blaxel.ProcessStreamOptions{})
			} else {
				_, err = sandbox.Process.New(context.Background(), blaxel.ProcessRequestParam{Command: "side effect"}, option.WithMaxRetries(5))
			}
			p := requireProcessError(t, err)
			submittedName, _ := submitted.Load().(string)
			if calls.Load() != 1 || submittedName == "" || p.Identifier != submittedName {
				t.Fatalf("error=%+v calls=%d submitted=%s", p, calls.Load(), submittedName)
			}
		})
	}
}
func TestProcessCreationPreservesProvidedNameAndCanRecover(t *testing.T) {
	var posts atomic.Int32
	sandbox := processTestSandbox(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			posts.Add(1)
			http.Error(w, "response lost", 504)
			return
		}
		if r.URL.Path != "/process/original" {
			t.Errorf("wrong recovery URL: %s", r.URL.Path)
		}
		writeProcess(w, "completed")
	})
	_, err := sandbox.Process.New(context.Background(), blaxel.ProcessRequestParam{Command: "side effect", Name: blaxel.String("original")})
	p := requireProcessError(t, err)
	result, err := sandbox.Process.Wait(context.Background(), p.Identifier, time.Second, time.Millisecond)
	if err != nil || result.ExitCode != 7 || posts.Load() != 1 {
		t.Fatalf("result=%+v err=%v posts=%d", result, err, posts.Load())
	}
}
func TestProcessStreamMissingResultPreservesIdentity(t *testing.T) {
	sandbox := processTestSandbox(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/x-ndjson")
		io.WriteString(w, "{\"type\":\"stdout\",\"data\":\"partial\"}\n")
	})
	var logs string
	_, err := sandbox.Process.ExecWithStreaming(context.Background(), blaxel.ProcessRequestParam{Command: "side effect", Name: blaxel.String("original")}, blaxel.ProcessStreamOptions{OnStdout: func(s string) { logs += s }})
	p := requireProcessError(t, err)
	if p.Identifier != "original" || logs != "partial" {
		t.Fatalf("error=%+v logs=%s", p, logs)
	}
}
func TestProcessStopAndWaitConfirmsAPIState(t *testing.T) {
	for _, force := range []bool{false, true} {
		t.Run(fmt.Sprint(force), func(t *testing.T) {
			var deletes, reads atomic.Int32
			sandbox := processTestSandbox(t, func(w http.ResponseWriter, r *http.Request) {
				if r.Method == http.MethodDelete {
					deletes.Add(1)
					if strings.HasSuffix(r.URL.Path, "/kill") != force {
						t.Errorf("unexpected path %s", r.URL.Path)
					}
					w.Header().Set("Content-Type", "application/json")
					io.WriteString(w, `{"message":"requested"}`)
					return
				}
				if reads.Add(1) == 1 {
					writeProcess(w, "running")
				} else {
					writeProcess(w, "killed")
				}
			})
			var result *blaxel.ProcessResponse
			var err error
			if force {
				result, err = sandbox.Process.KillAndWait(context.Background(), "original", time.Second, time.Millisecond)
			} else {
				result, err = sandbox.Process.StopAndWait(context.Background(), "original", time.Second, time.Millisecond)
			}
			if err != nil || result.Status != "killed" || deletes.Load() != 1 || reads.Load() != 2 {
				t.Fatalf("result=%+v err=%v deletes=%d reads=%d", result, err, deletes.Load(), reads.Load())
			}
		})
	}
}

func TestProcessStopAndWaitRecoversLostAcknowledgment(t *testing.T) {
	var deletes atomic.Int32
	sandbox := processTestSandbox(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			deletes.Add(1)
			http.Error(w, "lost response", 502)
			return
		}
		writeProcess(w, "killed")
	})
	result, err := sandbox.Process.KillAndWait(context.Background(), "original", time.Second, time.Millisecond)
	if err != nil || result.Status != "killed" || deletes.Load() != 1 {
		t.Fatalf("result=%+v err=%v deletes=%d", result, err, deletes.Load())
	}
}
func TestProcessStopAndWaitDoesNotClaimStopOnTimeout(t *testing.T) {
	sandbox := processTestSandbox(t, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodDelete {
			w.Header().Set("Content-Type", "application/json")
			io.WriteString(w, `{"message":"requested"}`)
			return
		}
		writeProcess(w, "running")
	})
	result, err := sandbox.Process.KillAndWait(context.Background(), "original", 20*time.Millisecond, time.Millisecond)
	p := requireProcessError(t, err)
	if result != nil || !errors.Is(err, context.DeadlineExceeded) || p.LastKnownProcess.Status != "running" {
		t.Fatalf("result=%+v err=%+v", result, p)
	}
}

func TestProcessStreamLogsExposesHTTPFailureWithoutCallback(t *testing.T) {
	sandbox := processTestSandbox(t, func(w http.ResponseWriter, r *http.Request) { http.Error(w, "missing", 404) })
	stream := sandbox.Process.StreamLogs(context.Background(), "original", blaxel.ProcessStreamOptions{})
	stream.Wait()
	p := requireProcessError(t, stream.Err())
	if p.Identifier != "original" || !strings.Contains(p.Error(), "HTTP 404") {
		t.Fatal(p)
	}
}
func TestProcessStreamLogsExposesReadFailure(t *testing.T) {
	sandbox := processTestSandbox(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Length", "1000")
		io.WriteString(w, "stdout:partial\n")
	})
	stream := sandbox.Process.StreamLogs(context.Background(), "original", blaxel.ProcessStreamOptions{})
	stream.Wait()
	if !errors.Is(stream.Err(), io.ErrUnexpectedEOF) {
		t.Fatalf("error=%v", stream.Err())
	}
}
func TestProcessStreamLogsCloseIsNotFailure(t *testing.T) {
	started := make(chan struct{})
	sandbox := processTestSandbox(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.(http.Flusher).Flush()
		close(started)
		<-r.Context().Done()
	})
	stream := sandbox.Process.StreamLogs(context.Background(), "original", blaxel.ProcessStreamOptions{})
	<-started
	stream.Close()
	stream.Wait()
	if stream.Err() != nil {
		t.Fatal(stream.Err())
	}
}
func TestProcessStreamLogsContextCancellationIsObservable(t *testing.T) {
	started := make(chan struct{})
	sandbox := processTestSandbox(t, func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.(http.Flusher).Flush()
		close(started)
		<-r.Context().Done()
	})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	stream := sandbox.Process.StreamLogs(ctx, "original", blaxel.ProcessStreamOptions{})
	<-started
	cancel()
	stream.Wait()
	if !errors.Is(stream.Err(), context.Canceled) {
		t.Fatalf("error=%v", stream.Err())
	}
}

func TestProcessSetupFailuresAreNotRecoverable(t *testing.T) {
	for _, streaming := range []bool{false, true} {
		t.Run(fmt.Sprint(streaming), func(t *testing.T) {
			var calls atomic.Int32
			client := blaxel.NewClient()
			sandbox := client.Sandboxes.FromSession(blaxel.SessionWithToken{Name: "test", URL: "https://sandbox.test", Token: "test"}, option.WithBaseURL("://invalid"), option.WithHTTPClient(&http.Client{Transport: &closureTransport{fn: func(r *http.Request) (*http.Response, error) { calls.Add(1); return nil, io.EOF }}}))
			var err error
			if streaming {
				_, err = sandbox.Process.ExecWithStreaming(context.Background(), blaxel.ProcessRequestParam{Command: "side effect"}, blaxel.ProcessStreamOptions{})
			} else {
				_, err = sandbox.Process.New(context.Background(), blaxel.ProcessRequestParam{Command: "side effect"})
			}
			var processErr *blaxel.ProcessError
			if err == nil || errors.As(err, &processErr) || calls.Load() != 0 {
				t.Fatalf("error=%v processError=%+v requests=%d", err, processErr, calls.Load())
			}
		})
	}
}
func TestProcessAlreadyCanceledCreationIsNotRecoverable(t *testing.T) {
	for _, streaming := range []bool{false, true} {
		t.Run(fmt.Sprint(streaming), func(t *testing.T) {
			var calls atomic.Int32
			sandbox := processTestSandbox(t, func(w http.ResponseWriter, r *http.Request) { calls.Add(1); writeProcess(w, "completed") })
			ctx, cancel := context.WithCancel(context.Background())
			cancel()
			var err error
			if streaming {
				_, err = sandbox.Process.ExecWithStreaming(ctx, blaxel.ProcessRequestParam{Command: "side effect"}, blaxel.ProcessStreamOptions{})
			} else {
				_, err = sandbox.Process.New(ctx, blaxel.ProcessRequestParam{Command: "side effect"})
			}
			var processErr *blaxel.ProcessError
			if !errors.Is(err, context.Canceled) || errors.As(err, &processErr) || calls.Load() != 0 {
				t.Fatalf("error=%v processError=%+v requests=%d", err, processErr, calls.Load())
			}
		})
	}
}
func TestWatchStreamControlExposesHTTPFailure(t *testing.T) {
	sandbox := processTestSandbox(t, func(w http.ResponseWriter, r *http.Request) { http.Error(w, "missing", 500) })
	stream := sandbox.FS.Watch(context.Background(), "tmp", func(blaxel.WatchEvent) {}, nil)
	stream.Wait()
	if stream.Err() == nil || !strings.Contains(stream.Err().Error(), "HTTP 500") {
		t.Fatalf("error=%v", stream.Err())
	}
}
func TestWatchStreamControlExposesReadFailure(t *testing.T) {
	sandbox := processTestSandbox(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Length", "1000")
		io.WriteString(w, "\n")
	})
	stream := sandbox.FS.Watch(context.Background(), "tmp", func(blaxel.WatchEvent) {}, nil)
	stream.Wait()
	if !errors.Is(stream.Err(), io.ErrUnexpectedEOF) {
		t.Fatalf("error=%v", stream.Err())
	}
}
func TestWatchStreamControlCloseAndCancel(t *testing.T) {
	for _, explicitClose := range []bool{false, true} {
		t.Run(fmt.Sprint(explicitClose), func(t *testing.T) {
			started := make(chan struct{})
			sandbox := processTestSandbox(t, func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(200)
				w.(http.Flusher).Flush()
				close(started)
				<-r.Context().Done()
			})
			ctx, cancel := context.WithCancel(context.Background())
			defer cancel()
			stream := sandbox.FS.Watch(ctx, "tmp", func(blaxel.WatchEvent) {}, nil)
			<-started
			if explicitClose {
				stream.Close()
			} else {
				cancel()
			}
			stream.Wait()
			if explicitClose && stream.Err() != nil {
				t.Fatal(stream.Err())
			}
			if !explicitClose && !errors.Is(stream.Err(), context.Canceled) {
				t.Fatalf("error=%v", stream.Err())
			}
		})
	}
}
