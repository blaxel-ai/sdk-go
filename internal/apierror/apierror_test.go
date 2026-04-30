package apierror

import (
	"encoding/json"
	"testing"
)

func TestUnmarshalGatewayError(t *testing.T) {
	body := `{
		"error": {
			"code": "WORKLOAD_NOT_FOUND",
			"message": "The requested workload was not found",
			"origin": "platform",
			"retryable": false,
			"action": "Check the workload name and try again",
			"do_not": "Retry without checking the workload name",
			"docs_url": "https://docs.blaxel.ai/errors/workload-not-found",
			"status": 404
		}
	}`

	var apiErr Error
	if err := json.Unmarshal([]byte(body), &apiErr); err != nil {
		t.Fatalf("UnmarshalJSON failed: %v", err)
	}

	if apiErr.ErrorCode != ErrWorkloadNotFound {
		t.Errorf("ErrorCode = %q, want %q", apiErr.ErrorCode, ErrWorkloadNotFound)
	}
	if apiErr.Message != "The requested workload was not found" {
		t.Errorf("Message = %q, want %q", apiErr.Message, "The requested workload was not found")
	}
	if apiErr.Origin != "platform" {
		t.Errorf("Origin = %q, want %q", apiErr.Origin, "platform")
	}
	if apiErr.Retryable != false {
		t.Errorf("Retryable = %v, want false", apiErr.Retryable)
	}
	if apiErr.Action != "Check the workload name and try again" {
		t.Errorf("Action = %q, want %q", apiErr.Action, "Check the workload name and try again")
	}
	if apiErr.DoNot != "Retry without checking the workload name" {
		t.Errorf("DoNot = %q, want %q", apiErr.DoNot, "Retry without checking the workload name")
	}
	if apiErr.DocsURL != "https://docs.blaxel.ai/errors/workload-not-found" {
		t.Errorf("DocsURL = %q, want %q", apiErr.DocsURL, "https://docs.blaxel.ai/errors/workload-not-found")
	}
	if apiErr.Status != 404 {
		t.Errorf("Status = %d, want 404", apiErr.Status)
	}
	if !apiErr.IsGatewayError() {
		t.Error("IsGatewayError() = false, want true")
	}
	if apiErr.IsRetryable() {
		t.Error("IsRetryable() = true, want false")
	}
}

func TestUnmarshalGatewayErrorRetryable(t *testing.T) {
	body := `{
		"error": {
			"code": "WORKLOAD_UNAVAILABLE",
			"message": "The workload is temporarily unavailable",
			"origin": "platform",
			"retryable": true,
			"action": "Retry the request after a short delay",
			"do_not": "",
			"docs_url": "",
			"status": 503
		}
	}`

	var apiErr Error
	if err := json.Unmarshal([]byte(body), &apiErr); err != nil {
		t.Fatalf("UnmarshalJSON failed: %v", err)
	}

	if apiErr.ErrorCode != ErrWorkloadUnavailable {
		t.Errorf("ErrorCode = %q, want %q", apiErr.ErrorCode, ErrWorkloadUnavailable)
	}
	if !apiErr.IsGatewayError() {
		t.Error("IsGatewayError() = false, want true")
	}
	if !apiErr.IsRetryable() {
		t.Error("IsRetryable() = false, want true")
	}
	if apiErr.Status != 503 {
		t.Errorf("Status = %d, want 503", apiErr.Status)
	}
}

func TestUnmarshalNonGatewayError(t *testing.T) {
	body := `{"message": "something went wrong"}`

	var apiErr Error
	if err := json.Unmarshal([]byte(body), &apiErr); err != nil {
		t.Fatalf("UnmarshalJSON failed: %v", err)
	}

	if apiErr.IsGatewayError() {
		t.Error("IsGatewayError() = true, want false for non-gateway error")
	}
	if apiErr.IsRetryable() {
		t.Error("IsRetryable() = true, want false for non-gateway error")
	}
	if apiErr.ErrorCode != "" {
		t.Errorf("ErrorCode = %q, want empty", apiErr.ErrorCode)
	}
}

func TestUnmarshalEmptyBody(t *testing.T) {
	body := `{}`

	var apiErr Error
	if err := json.Unmarshal([]byte(body), &apiErr); err != nil {
		t.Fatalf("UnmarshalJSON failed: %v", err)
	}

	if apiErr.IsGatewayError() {
		t.Error("IsGatewayError() = true, want false for empty body")
	}
	if apiErr.ErrorCode != "" {
		t.Errorf("ErrorCode = %q, want empty", apiErr.ErrorCode)
	}
}

func TestErrorCodeConstants(t *testing.T) {
	codes := map[string]string{
		"ErrRouteNotFound":          ErrRouteNotFound,
		"ErrWorkloadNotFound":       ErrWorkloadNotFound,
		"ErrWorkspaceNotFound":      ErrWorkspaceNotFound,
		"ErrWorkloadUnavailable":    ErrWorkloadUnavailable,
		"ErrAuthenticationRequired": ErrAuthenticationRequired,
		"ErrAuthenticationFailed":   ErrAuthenticationFailed,
		"ErrForbidden":              ErrForbidden,
		"ErrBadRequest":             ErrBadRequest,
		"ErrUsageLimitExceeded":     ErrUsageLimitExceeded,
		"ErrPolicyViolation":        ErrPolicyViolation,
	}

	expected := map[string]string{
		"ErrRouteNotFound":          "ROUTE_NOT_FOUND",
		"ErrWorkloadNotFound":       "WORKLOAD_NOT_FOUND",
		"ErrWorkspaceNotFound":      "WORKSPACE_NOT_FOUND",
		"ErrWorkloadUnavailable":    "WORKLOAD_UNAVAILABLE",
		"ErrAuthenticationRequired": "AUTHENTICATION_REQUIRED",
		"ErrAuthenticationFailed":   "AUTHENTICATION_FAILED",
		"ErrForbidden":              "FORBIDDEN",
		"ErrBadRequest":             "BAD_REQUEST",
		"ErrUsageLimitExceeded":     "USAGE_LIMIT_EXCEEDED",
		"ErrPolicyViolation":        "POLICY_VIOLATION",
	}

	for name, got := range codes {
		want := expected[name]
		if got != want {
			t.Errorf("%s = %q, want %q", name, got, want)
		}
	}
}
