package blaxel_test

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
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

// Run against a disposable sandbox API with LOCAL_PROCESS_API_URL=http://127.0.0.1:18083.
// The proxy drops a real status response after the first running observation.
func TestProcessWaitLiveNetworkRecovery(t *testing.T) {
	endpoint := os.Getenv("LOCAL_PROCESS_API_URL")
	if endpoint == "" {
		t.Skip("set LOCAL_PROCESS_API_URL to a disposable sandbox API")
	}
	target, err := url.Parse(endpoint)
	if err != nil {
		t.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	name := fmt.Sprintf("go-wait-recovery-%d", time.Now().UnixNano())
	startsFile, releaseFile := "/tmp/"+name+"-starts", "/tmp/"+name+"-release"
	client := blaxel.NewClient(option.WithAPIKey("test"))
	direct := client.Sandboxes.FromSession(blaxel.SessionWithToken{Name: "test", URL: endpoint, Token: "test"})
	t.Cleanup(func() {
		cleanupCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()
		_, _ = direct.Process.Kill(cleanupCtx, name)
		_, _ = direct.FS.RM(cleanupCtx, startsFile, false)
		_, _ = direct.FS.RM(cleanupCtx, releaseFile, false)
	})
	var posts, reads, dropped atomic.Int32
	releaseResult := make(chan error, 1)
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ModifyResponse = func(response *http.Response) error {
		if response.Request.Method == http.MethodPost {
			posts.Add(1)
		}
		if response.Request.Method == http.MethodGet && reads.Add(1) == 2 {
			var observed struct {
				Status string `json:"status"`
			}
			decodeErr := json.NewDecoder(response.Body).Decode(&observed)
			_ = response.Body.Close()
			if decodeErr != nil || observed.Status != "running" {
				t.Errorf("dropped observation must be running: status=%s error=%v", observed.Status, decodeErr)
			}
			dropped.Add(1)
			return io.ErrUnexpectedEOF
		}
		return nil
	}
	proxy.ErrorHandler = func(w http.ResponseWriter, r *http.Request, err error) {
		conn, _, hijackErr := w.(http.Hijacker).Hijack()
		releaseErr := hijackErr
		if hijackErr == nil {
			_ = conn.Close()
			// Release the real command only after dropping its running-state response.
			releaseCtx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
			_, releaseErr = direct.FS.Write(releaseCtx, releaseFile, "release")
			cancel()
		}
		select {
		case releaseResult <- releaseErr:
		default:
		}
	}
	server := httptest.NewServer(proxy)
	defer server.Close()
	// Disable HTTP transport's own connection reuse so it cannot transparently
	// retry the dropped idempotent GET before the SDK Wait loop sees it.
	transport := http.DefaultTransport.(*http.Transport).Clone()
	transport.DisableKeepAlives = true
	defer transport.CloseIdleConnections()
	sandbox := client.Sandboxes.FromSession(blaxel.SessionWithToken{Name: "test", URL: server.URL, Token: "test"}, option.WithHTTPClient(&http.Client{Transport: transport}), option.WithMaxRetries(0))
	command := fmt.Sprintf("printf 'started\\n' >> %s; while [ ! -f %s ]; do sleep 0.02; done; echo recovered-output; exit 7", startsFile, releaseFile)
	initial, err := sandbox.Process.New(ctx, blaxel.ProcessRequestParam{Name: blaxel.String(name), Command: command, WaitForCompletion: blaxel.Bool(false), KeepAlive: blaxel.Bool(false)})
	if err != nil {
		t.Fatal(err)
	}
	if initial.Status != "running" {
		t.Fatalf("expected initial running state: %+v", initial)
	}
	result, err := sandbox.Process.Wait(ctx, name, 5*time.Second, 20*time.Millisecond)
	if dropped.Load() > 0 {
		select {
		case releaseErr := <-releaseResult:
			if releaseErr != nil {
				t.Fatalf("release command: %v", releaseErr)
			}
		case <-ctx.Done():
			t.Fatalf("waiting for command release: %v", ctx.Err())
		}
	}
	if err != nil {
		t.Fatalf("Wait did not recover original command after the lost GET response: %v", err)
	}
	if result.ExitCode != 7 || result.Status != "failed" {
		t.Fatalf("terminal result=%+v", result)
	}
	logs, err := sandbox.Process.GetLogs(ctx, name)
	if err != nil {
		t.Fatal(err)
	}
	starts, err := direct.FS.Read(ctx, startsFile)
	if err != nil || starts != "started\n" || !strings.Contains(logs.Stdout, "recovered-output") || posts.Load() != 1 || dropped.Load() != 1 {
		t.Fatalf("starts=%q stdout=%q posts=%d dropped=%d error=%v", starts, logs.Stdout, posts.Load(), dropped.Load(), err)
	}
	t.Logf("one real process, one POST, one lost status response, recovered exit 7; status GET attempts=%d", reads.Load()-1)
}
