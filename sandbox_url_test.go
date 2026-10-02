package blaxel_test

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	blaxel "github.com/blaxel-ai/sdk-go"
	"github.com/blaxel-ai/sdk-go/option"
)

func TestSandboxURLWithOrWithoutTrailingSlash(t *testing.T) {
	for _, suffix := range []string{"", "/", "//"} {
		t.Run("suffix="+suffix, func(t *testing.T) {
			var paths []string
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				paths = append(paths, r.URL.Path)
				w.Header().Set("Content-Type", "application/json")
				w.Write([]byte(`{"name":"p","status":"completed","path":"/tmp/a"}`))
			}))
			defer server.Close()
			client := blaxel.NewClient(option.WithAPIKey("test"))
			sandbox := client.Sandboxes.FromSession(blaxel.SessionWithToken{Name: "test", URL: server.URL + suffix, Token: "test"}, option.WithHTTPClient(server.Client()), option.WithMaxRetries(0))
			if _, err := sandbox.FS.Write(context.Background(), "tmp/a", "x"); err != nil {
				t.Fatal(err)
			}
			if _, err := sandbox.Process.Get(context.Background(), "p"); err != nil {
				t.Fatal(err)
			}
			if want := []string{"/filesystem/tmp/a", "/process/p"}; len(paths) != 2 || paths[0] != want[0] || paths[1] != want[1] {
				t.Fatalf("paths=%q want %q", paths, want)
			}
		})
	}
}
