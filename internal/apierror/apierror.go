// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package apierror

import (
	"fmt"
	"net/http"
	"net/http/httputil"

	"github.com/blaxel-ai/sdk-go/internal/apijson"
	"github.com/blaxel-ai/sdk-go/packages/respjson"
)

// Error represents an error that originates from the API, i.e. when a request is
// made and the API returns a response with a HTTP status code. Other errors are
// not wrapped by this SDK.
type Error struct {
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	StatusCode int
	Request    *http.Request
	Response   *http.Response

	// BlaxelErrorCode is the stable error code from the X-Blaxel-Error-Code header,
	// set only on gateway-synthesized error responses (e.g. "WORKLOAD_UNAVAILABLE").
	BlaxelErrorCode string
	// BlaxelSource is the value of the X-Blaxel-Source header.
	// When equal to "platform", the error was synthesized by the Blaxel gateway proxy.
	BlaxelSource string
	// Retryable indicates whether retrying the same request may succeed.
	Retryable bool
	// Action is a directive telling the caller what to do next.
	Action string
	// DoNot is an optional anti-pattern warning discouraging common failure modes.
	DoNot string
	// DocsURL is an optional link to the relevant Blaxel documentation page.
	DocsURL string
}

// Returns the unmodified JSON received from the API
func (r Error) RawJSON() string { return r.JSON.raw }
func (r *Error) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func (r *Error) Error() string {
	// Attempt to re-populate the response body
	return fmt.Sprintf("%s %q: %d %s %s", r.Request.Method, r.Request.URL, r.Response.StatusCode, http.StatusText(r.Response.StatusCode), r.JSON.raw)
}

func (r *Error) DumpRequest(body bool) []byte {
	if r.Request.GetBody != nil {
		r.Request.Body, _ = r.Request.GetBody()
	}
	out, _ := httputil.DumpRequestOut(r.Request, body)
	return out
}

func (r *Error) DumpResponse(body bool) []byte {
	out, _ := httputil.DumpResponse(r.Response, body)
	return out
}

// IsGatewayError returns true when this error was synthesized by the Blaxel
// gateway proxy rather than the upstream workload.
func (r *Error) IsGatewayError() bool {
	return r.BlaxelSource == "platform"
}
