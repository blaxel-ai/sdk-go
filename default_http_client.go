// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blaxel

import (
	"errors"
	"net/http"
	"time"
)

// stripBlaxelAuthOnRedirect drops Blaxel's custom auth headers when a redirect
// crosses to a different host. Go's net/http only strips well-known headers
// (e.g. Authorization) on cross-host redirects, not custom ones, so without
// this a 3xx to an attacker-controlled host could leak the bearer/preview
// token. It also preserves the stdlib default of stopping after 10 redirects.
func stripBlaxelAuthOnRedirect(req *http.Request, via []*http.Request) error {
	if len(via) >= 10 {
		return errors.New("stopped after 10 redirects")
	}
	if len(via) > 0 && req.URL.Host != via[len(via)-1].URL.Host {
		req.Header.Del("X-Blaxel-Authorization")
		req.Header.Del("X-Blaxel-Preview-Token")
	}
	return nil
}

// defaultResponseHeaderTimeout bounds the time between a fully written request
// and the server's response headers. It does not apply to the response body,
// so long-running streams are unaffected. Without this, a server that accepts
// the connection but never responds would hang the request indefinitely.
const defaultResponseHeaderTimeout = 10 * time.Minute

// defaultHTTPClient returns an [*http.Client] used when the caller does not
// supply one via [option.WithHTTPClient]. When [http.DefaultTransport] is the
// stdlib [*http.Transport], it is cloned and a [http.Transport.ResponseHeaderTimeout]
// is set so stuck connections fail fast instead of compounding across retries.
// If [http.DefaultTransport] has been wrapped (for example by otelhttp for
// distributed tracing), the wrapping is preserved and the header timeout is
// skipped.
func defaultHTTPClient() *http.Client {
	if t, ok := http.DefaultTransport.(*http.Transport); ok {
		t = t.Clone()
		t.ResponseHeaderTimeout = defaultResponseHeaderTimeout
		return &http.Client{Transport: t, CheckRedirect: stripBlaxelAuthOnRedirect}
	}
	return &http.Client{Transport: http.DefaultTransport, CheckRedirect: stripBlaxelAuthOnRedirect}
}
