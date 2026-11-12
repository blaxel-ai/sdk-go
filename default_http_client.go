// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blaxel

import (
	"net/http"
	"time"
)

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
		return &http.Client{Transport: t, CheckRedirect: noFollowRedirect}
	}
	return &http.Client{Transport: http.DefaultTransport, CheckRedirect: noFollowRedirect}
}

// noFollowRedirect stops the client from following 3xx responses. A redirect
// returned by a sandbox-api call can be attacker-controlled (a malicious
// sandbox returning a 302/307 to its own host): following it would re-send the
// request body (for example the command and environment variables of a process
// exec) to the attacker and make the SDK trust the attacker's forged response.
// The custom X-Blaxel-Authorization header is also not stripped by net/http on
// a redirect, so following one leaks the credential too. Returning
// [http.ErrUseLastResponse] hands the 3xx back to the caller instead.
func noFollowRedirect(*http.Request, []*http.Request) error {
	return http.ErrUseLastResponse
}
