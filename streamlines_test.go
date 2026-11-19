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
	"testing"

	blaxel "github.com/blaxel-ai/sdk-go"
	"github.com/blaxel-ai/sdk-go/option"
)

// ndjsonEvent is one line of a streamed exec, as the sandbox API writes it.
func ndjsonEvent(t *testing.T, eventType, data string) string {
	t.Helper()
	line, err := json.Marshal(map[string]string{"type": eventType, "data": data})
	if err != nil {
		t.Fatal(err)
	}
	return string(line) + "\n"
}

func TestProcessStreamReturnsResultLinesOver64KB(t *testing.T) {
	output := strings.Repeat("a", 150_000)
	result, err := json.Marshal(map[string]any{
		"name": "proc", "pid": "123", "status": "completed", "exitCode": 0,
		"stdout": output, "logs": output,
	})
	if err != nil {
		t.Fatal(err)
	}
	resultLine := ndjsonEvent(t, "result", string(result))
	if len(resultLine) < 200_000 {
		t.Fatalf("result line is %d bytes, want over 200 KB", len(resultLine))
	}
	sandbox := processServer(t, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/x-ndjson")
		io.WriteString(w, ndjsonEvent(t, "stdout", output))
		io.WriteString(w, ndjsonEvent(t, "keepalive", ""))
		io.WriteString(w, resultLine)
	})

	var streamed strings.Builder
	var streamErrors []error
	res, err := sandbox.Process.ExecWithStreaming(context.Background(), blaxel.ProcessRequestParam{Command: "big"}, blaxel.ProcessStreamOptions{
		OnStdout: func(s string) { streamed.WriteString(s) },
		OnError:  func(err error) { streamErrors = append(streamErrors, err) },
	})
	if err != nil {
		t.Fatalf("ExecWithStreaming: %v", err)
	}
	if res.Stdout != output || res.Logs != output || res.ExitCode != 0 {
		t.Fatalf("result has stdout %d bytes, logs %d bytes, exit %d", len(res.Stdout), len(res.Logs), res.ExitCode)
	}
	if streamed.String() != output {
		t.Fatalf("OnStdout got %d bytes, want %d", streamed.Len(), len(output))
	}
	if len(streamErrors) != 0 {
		t.Fatalf("OnError called: %v", streamErrors)
	}
}

func TestProcessStreamReturnsTheErrorEvent(t *testing.T) {
	cases := []struct {
		name      string
		event     string
		wantError string
		check     func(t *testing.T, e *blaxel.ProcessStreamError)
	}{
		{
			name:      "message only",
			event:     `{"type":"error","data":"could not start process: chdir /does/not/exist: no such file or directory"}`,
			wantError: "could not start process: chdir /does/not/exist: no such file or directory",
			check: func(t *testing.T, e *blaxel.ProcessStreamError) {
				if e.Code != "" || e.Status != 0 || e.DispatchState != "" {
					t.Fatalf("unexpected fields: %+v", e)
				}
			},
		},
		{
			name: "code and platform fields",
			event: `{"type":"error","data":"working directory /etc/hosts is not a directory",` +
				`"code":"PROCESS_WORKING_DIR_UNUSABLE","error_code":"PROCESS_WORKING_DIR_UNUSABLE",` +
				`"message":"working directory /etc/hosts is not a directory","action":"Pass a directory.",` +
				`"do_not":"Do not resend unchanged.","docs_url":"https://docs.blaxel.ai/troubleshooting/error-codes#process-working-dir-unusable",` +
				`"retryable":false,"safe_to_retry_request":false,"dispatch_state":"not_dispatched","status":422,` +
				`"origin":"platform","timestamp":"2026-11-18T12:00:00Z","request_id":"req-1"}`,
			wantError: "working directory /etc/hosts is not a directory (code PROCESS_WORKING_DIR_UNUSABLE)",
			check: func(t *testing.T, e *blaxel.ProcessStreamError) {
				if e.Message != "working directory /etc/hosts is not a directory" || e.Code != "PROCESS_WORKING_DIR_UNUSABLE" ||
					e.Action != "Pass a directory." || e.DoNot != "Do not resend unchanged." ||
					!strings.HasSuffix(e.DocsURL, "#process-working-dir-unusable") || e.Retryable || e.SafeToRetryRequest ||
					e.DispatchState != "not_dispatched" || e.Status != 422 || e.RequestID != "req-1" {
					t.Fatalf("unexpected fields: %+v", e)
				}
			},
		},
		{
			name:      "message field and an unexpected status type",
			event:     `{"type":"error","message":"stream failed","error_code":"PROCESS_STREAM_FAILED","status":{"state":"frozen"},"retryable":true}`,
			wantError: "stream failed (code PROCESS_STREAM_FAILED)",
			check: func(t *testing.T, e *blaxel.ProcessStreamError) {
				if e.Status != 0 || !e.Retryable {
					t.Fatalf("unexpected fields: %+v", e)
				}
			},
		},
		{
			name:      "no message",
			event:     `{"type":"error"}`,
			wantError: "the sandbox API ended the stream with an error and no message",
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			sandbox := processServer(t, func(w http.ResponseWriter, r *http.Request) {
				w.Header().Set("Content-Type", "application/x-ndjson")
				io.WriteString(w, ndjsonEvent(t, "stdout", "before\n"))
				io.WriteString(w, tc.event+"\n")
			})
			var stdout []string
			var onError []error
			res, err := sandbox.Process.ExecWithStreaming(context.Background(), blaxel.ProcessRequestParam{Command: "pwd"}, blaxel.ProcessStreamOptions{
				OnStdout: func(s string) { stdout = append(stdout, s) },
				OnError:  func(err error) { onError = append(onError, err) },
			})
			if res != nil {
				t.Fatalf("result=%+v, want nil", res)
			}
			var streamErr *blaxel.ProcessStreamError
			if !errors.As(err, &streamErr) {
				t.Fatalf("error=%v (%T), want *ProcessStreamError", err, err)
			}
			if err.Error() != tc.wantError {
				t.Fatalf("error=%q, want %q", err.Error(), tc.wantError)
			}
			if streamErr.RawJSON() != tc.event {
				t.Fatalf("RawJSON=%q", streamErr.RawJSON())
			}
			if len(onError) != 1 || onError[0] != err {
				t.Fatalf("OnError got %v, want the returned error once", onError)
			}
			if len(stdout) != 1 || stdout[0] != "before\n" {
				t.Fatalf("stdout=%q", stdout)
			}
			if tc.check != nil {
				tc.check(t, streamErr)
			}
		})
	}
}

func TestProcessAndWatchStreamReadLinesOver64KB(t *testing.T) {
	long := strings.Repeat("b", 200_000)
	t.Run("logs", func(t *testing.T) {
		sandbox := processServer(t, func(w http.ResponseWriter, r *http.Request) {
			io.WriteString(w, "[keepalive]\nstdout:"+long+"\nstderr:short\n")
		})
		var stdout, stderr []string
		stream := sandbox.Process.StreamLogs(context.Background(), "proc", blaxel.ProcessStreamOptions{
			OnStdout: func(s string) { stdout = append(stdout, s) },
			OnStderr: func(s string) { stderr = append(stderr, s) },
		})
		if err := stream.Err(); err != nil {
			t.Fatalf("Err: %v", err)
		}
		if len(stdout) != 1 || stdout[0] != long || len(stderr) != 1 || stderr[0] != "short" {
			t.Fatalf("stdout %d lines, stderr %q", len(stdout), stderr)
		}
	})
	t.Run("watch", func(t *testing.T) {
		sandbox := processServer(t, func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprintf(w, "{\"op\":\"WRITE\",\"path\":\"/app\",\"name\":%q}\n{\"op\":\"CREATE\",\"path\":\"/app\",\"name\":\"b\"}\n", long)
		})
		var events []blaxel.WatchEvent
		stream := sandbox.FS.Watch(context.Background(), "/app", func(e blaxel.WatchEvent) { events = append(events, e) }, nil)
		if err := stream.Err(); err != nil {
			t.Fatalf("Err: %v", err)
		}
		if len(events) != 2 || events[0].Name != long || events[1].Name != "b" {
			t.Fatalf("got %d events", len(events))
		}
	})
}

func TestProcessStreamRunCodeReadsLinesOver64KB(t *testing.T) {
	image := strings.Repeat("i", 200_000)
	var server *httptest.Server
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case r.Method == http.MethodGet && strings.HasSuffix(r.URL.Path, "/sandboxes/jupyter"):
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintf(w, `{"metadata":{"name":"jupyter","url":%q},"spec":{},"status":"DEPLOYED"}`, server.URL)
		case r.Method == http.MethodPost && r.URL.Path == "/port/8888/execute":
			fmt.Fprintf(w, "{\"type\":\"stdout\",\"text\":\"hi\"}\n{\"type\":\"result\",\"png\":%q}\n", image)
		default:
			http.NotFound(w, r)
		}
	}))
	t.Cleanup(server.Close)
	client := blaxel.NewClient(option.WithAPIKey("test"), option.WithBaseURL(server.URL), option.WithHTTPClient(server.Client()), option.WithMaxRetries(0))
	interpreter, err := client.Sandboxes.GetCodeInterpreter(context.Background(), "jupyter")
	if err != nil {
		t.Fatalf("GetCodeInterpreter: %v", err)
	}
	execution, err := interpreter.RunCode(context.Background(), "plot()", nil)
	if err != nil {
		t.Fatalf("RunCode: %v", err)
	}
	if len(execution.Results) != 1 || execution.Results[0]["png"] != image {
		t.Fatalf("got %d results", len(execution.Results))
	}
	if len(execution.Logs.Stdout) != 1 || execution.Logs.Stdout[0] != "hi" {
		t.Fatalf("stdout=%q", execution.Logs.Stdout)
	}
}
