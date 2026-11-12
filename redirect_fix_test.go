package blaxel

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

// A malicious sandbox-api can return a 3xx to an attacker host. The default
// client must not follow it: following re-sends the X-Blaxel-Authorization
// credential and the request body (e.g. a process exec's env) to the attacker.
func TestDefaultClientDoesNotFollowRedirects(t *testing.T) {
	for _, code := range []int{302, 307} {
		t.Run(http.StatusText(code), func(t *testing.T) {
			var gotAuth, gotBody string
			hits := 0
			atk := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				hits++
				gotAuth = r.Header.Get("X-Blaxel-Authorization")
				b, _ := io.ReadAll(r.Body)
				gotBody = string(b)
				w.Write([]byte("{}"))
			}))
			defer atk.Close()
			vic := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				http.Redirect(w, r, atk.URL+"/steal", code)
			}))
			defer vic.Close()

			req, _ := http.NewRequest("POST", vic.URL, strings.NewReader(`{"command":"env","env":{"DB_PASSWORD":"SECRET"}}`))
			req.Header.Set("X-Blaxel-Authorization", "Bearer SECRET-KEY")
			resp, err := defaultHTTPClient().Do(req)
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}
			defer resp.Body.Close()
			if resp.StatusCode != code {
				t.Fatalf("expected %d returned to caller, got %d", code, resp.StatusCode)
			}
			if hits != 0 {
				t.Fatalf("attacker reached %d times (redirect followed); auth=%q body=%q", hits, gotAuth, gotBody)
			}
		})
	}
}
