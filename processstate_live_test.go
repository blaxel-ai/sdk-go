package blaxel_test

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	blaxel "github.com/blaxel-ai/sdk-go"
	"github.com/blaxel-ai/sdk-go/option"
)

// TestProcessStateLive requires an explicitly supplied disposable sandbox API.
// Example: BLAXEL_PROCESS_TEST_URL=http://127.0.0.1:18083 go test . -run TestProcessStateLive -v
// It runs commands directly in that sandbox; do not point it at production.
func TestProcessStateLive(t *testing.T) {
	endpoint := os.Getenv("BLAXEL_PROCESS_TEST_URL")
	if endpoint == "" {
		t.Skip("set BLAXEL_PROCESS_TEST_URL to a disposable sandbox API")
	}
	target, err := url.Parse(endpoint)
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	client := blaxel.NewClient(option.WithAPIKey("test"))
	sandbox := client.Sandboxes.FromSession(blaxel.SessionWithToken{Name: "process-test", URL: endpoint, Token: "test"})
	name := fmt.Sprintf("go-state-%d", time.Now().UnixNano())
	var creates, transient atomic.Int32
	proxy := httputil.NewSingleHostReverseProxy(target)
	// Forward creation to the real server, then hide its response. This reproduces
	// a command that started while the caller received only a connection failure.
	proxy.ModifyResponse = func(resp *http.Response) error {
		if resp.Request.Method == http.MethodPost {
			creates.Add(1)
			_ = resp.Body.Close()
			return fmt.Errorf("injected lost creation response")
		}
		return nil
	}
	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		http.Error(w, "injected response loss", http.StatusBadGateway)
	}
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet && transient.Add(1) == 1 {
			http.Error(w, "injected status failure", 503)
			return
		}
		proxy.ServeHTTP(w, r)
	}))
	defer server.Close()
	unreliable := client.Sandboxes.FromSession(blaxel.SessionWithToken{Name: "process-test", URL: server.URL, Token: "test"}, option.WithHTTPClient(server.Client()))
	_, err = unreliable.Process.New(ctx, blaxel.ProcessRequestParam{Name: blaxel.String(name), Command: "sleep 0.2; echo recovered-go-output; exit 7", WaitForCompletion: blaxel.Bool(false), KeepAlive: blaxel.Bool(false)})
	p := requireProcessError(t, err)
	if p.Identifier != name || creates.Load() != 1 {
		t.Fatalf("identity=%s creations=%d", p.Identifier, creates.Load())
	}
	result, err := unreliable.Process.Wait(ctx, p.Identifier, 10*time.Second, 10*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	if result.ExitCode != 7 || !strings.Contains(result.Logs, "recovered-go-output") {
		t.Fatalf("result=%+v", result)
	}
	logs, err := sandbox.Process.GetLogs(ctx, name)
	if err != nil || !strings.Contains(logs.Stdout, "recovered-go-output") {
		t.Fatalf("logs=%+v err=%v", logs, err)
	}
	if creates.Load() != 1 {
		t.Fatalf("command replayed: %d", creates.Load())
	}

	name = fmt.Sprintf("go-cancel-%d", time.Now().UnixNano())
	_, err = sandbox.Process.New(ctx, blaxel.ProcessRequestParam{Name: blaxel.String(name), Command: "sleep 30", WaitForCompletion: blaxel.Bool(false), KeepAlive: blaxel.Bool(false)})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		cleanupCtx, stop := context.WithTimeout(context.Background(), 3*time.Second)
		defer stop()
		_, _ = sandbox.Process.Kill(cleanupCtx, name)
	})
	waitCtx, stopWait := context.WithTimeout(ctx, 30*time.Millisecond)
	_, err = sandbox.Process.Wait(waitCtx, name, time.Second, time.Millisecond)
	stopWait()
	if !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("expected deadline: %v", err)
	}
	running, err := sandbox.Process.Get(ctx, name)
	if err != nil || running.Status != "running" {
		t.Fatalf("wait stopped command: %+v %v", running, err)
	}
	result, err = sandbox.Process.KillAndWait(ctx, name, 5*time.Second, 10*time.Millisecond)
	if err != nil || result.Status != "killed" {
		t.Fatalf("kill confirmation: %+v %v", result, err)
	}
	// Exercise the normal streaming path and client-generated recovery identity.
	streamed, err := sandbox.Process.ExecWithStreaming(ctx, blaxel.ProcessRequestParam{Command: "echo go-stream-live", WaitForCompletion: blaxel.Bool(true), KeepAlive: blaxel.Bool(false)}, blaxel.ProcessStreamOptions{})
	if err != nil || streamed.Name == "" || streamed.ExitCode != 0 || !strings.Contains(streamed.Logs, "go-stream-live") {
		t.Fatalf("stream result=%+v err=%v", streamed, err)
	}
	graceful, err := sandbox.Process.New(ctx, blaxel.ProcessRequestParam{Command: "sleep 30", WaitForCompletion: blaxel.Bool(false), KeepAlive: blaxel.Bool(false)})
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		cleanupCtx, stop := context.WithTimeout(context.Background(), 3*time.Second)
		defer stop()
		_, _ = sandbox.Process.Kill(cleanupCtx, graceful.Name)
	})
	stopped, err := sandbox.Process.StopAndWait(ctx, graceful.Name, 5*time.Second, 10*time.Millisecond)
	if err != nil || (stopped.Status != "stopped" && stopped.Status != "failed") {
		t.Fatalf("graceful stop=%+v err=%v", stopped, err)
	}
	t.Log("real command response lost, recovered original exit=7/logs after transient GET; no POST replay; wait timeout leaves command running; KillAndWait confirms API killed")
}
