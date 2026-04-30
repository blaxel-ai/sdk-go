// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package apierror

import (
	"encoding/json"
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

	// Gateway error body fields (populated when the response contains {"error": {...}})
	ErrorCode        string `json:"code"`
	Message          string `json:"message"`
	Origin           string `json:"origin"`
	Retryable        bool   `json:"retryable"`
	Action           string `json:"action"`
	DoNot            string `json:"do_not"`
	DocsURL          string `json:"docs_url"`
	Status           int    `json:"status"`
	hasErrorEnvelope bool
}

// IsGatewayError returns true if the error originated from the Blaxel platform gateway.
// Detection is based on the presence of the {"error": {...}} envelope in the response body.
func (r *Error) IsGatewayError() bool {
	return r.hasErrorEnvelope
}

// IsRetryable returns true if the gateway indicates the request can be retried.
func (r *Error) IsRetryable() bool {
	return r.Retryable
}

// Returns the unmodified JSON received from the API
func (r Error) RawJSON() string { return r.JSON.raw }
func (r *Error) UnmarshalJSON(data []byte) error {
	if err := apijson.UnmarshalRoot(data, r); err != nil {
		return err
	}

	// The gateway error body has the shape {"error": {"code": "...", ...}}.
	// Try to unwrap the nested "error" object and populate the typed fields.
	var envelope struct {
		Error json.RawMessage `json:"error"`
	}
	if err := json.Unmarshal(data, &envelope); err == nil && len(envelope.Error) > 0 {
		var inner struct {
			Code      string `json:"code"`
			Message   string `json:"message"`
			Origin    string `json:"origin"`
			Retryable bool   `json:"retryable"`
			Action    string `json:"action"`
			DoNot     string `json:"do_not"`
			DocsURL   string `json:"docs_url"`
			Status    int    `json:"status"`
		}
		if err := json.Unmarshal(envelope.Error, &inner); err == nil {
			r.hasErrorEnvelope = true
			r.ErrorCode = inner.Code
			r.Message = inner.Message
			if inner.Origin != "" {
				r.Origin = inner.Origin
			} else {
				r.Origin = "platform"
			}
			r.Retryable = inner.Retryable
			r.Action = inner.Action
			r.DoNot = inner.DoNot
			r.DocsURL = inner.DocsURL
			r.Status = inner.Status
		}
	}

	return nil
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
