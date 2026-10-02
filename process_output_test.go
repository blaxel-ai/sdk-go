package blaxel_test

import (
	"context"
	"io"
	"net/http"
	"reflect"
	"sync"
	"testing"

	blaxel "github.com/blaxel-ai/sdk-go"
)

type outputRecorder struct {
	mu                   sync.Mutex
	logs, stdout, stderr []string
}

func (o *outputRecorder) options() blaxel.ProcessStreamOptions {
	record := func(target *[]string) func(string) {
		return func(chunk string) {
			o.mu.Lock()
			defer o.mu.Unlock()
			*target = append(*target, chunk)
		}
	}
	return blaxel.ProcessStreamOptions{OnLog: record(&o.logs), OnStdout: record(&o.stdout), OnStderr: record(&o.stderr)}
}

func TestExecStreamingNegotiatesAndParsesNDJSON(t *testing.T) {
	var accept string
	sandbox := processServer(t, func(w http.ResponseWriter, r *http.Request) {
		accept = r.Header.Get("Accept")
		w.Header().Set("Content-Type", "application/x-ndjson")
		io.WriteString(w, `{"type":"stdout","data":"hi\n"}`+"\n\n"+`{"type":"stderr","data":"warn\n"}`+"\n")
		w.(http.Flusher).Flush()
		io.WriteString(w, `{"type":"result","data":"{\"name\":\"p\",\"status\":\"completed\",\"exitCode\":0}"}`+"\n")
	})
	var out outputRecorder
	result, err := sandbox.Process.ExecWithStreaming(context.Background(), blaxel.ProcessRequestParam{Command: "echo"}, out.options())
	if err != nil || result.Status != "completed" {
		t.Fatalf("result=%+v err=%v", result, err)
	}
	if accept != "application/x-ndjson, text/event-stream" {
		t.Fatalf("Accept=%q", accept)
	}
	if !reflect.DeepEqual(out.logs, []string{"hi\n", "warn\n"}) || !reflect.DeepEqual(out.stdout, []string{"hi\n"}) || !reflect.DeepEqual(out.stderr, []string{"warn\n"}) {
		t.Fatalf("logs=%q stdout=%q stderr=%q", out.logs, out.stdout, out.stderr)
	}
}
