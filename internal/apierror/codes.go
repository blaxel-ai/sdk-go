package apierror

// Stable error codes emitted by the Blaxel gateway proxy via the
// X-Blaxel-Error-Code response header and the error.code JSON body field.
//
// Use these constants to match gateway errors programmatically:
//
//	var apierr *apierror.Error
//	if errors.As(err, &apierr) && apierr.BlaxelErrorCode == apierror.ErrWorkloadUnavailable {
//	    // handle cold-start / unavailable workload
//	}
const (
	ErrRouteNotFound          = "ROUTE_NOT_FOUND"
	ErrWorkloadNotFound       = "WORKLOAD_NOT_FOUND"
	ErrWorkspaceNotFound      = "WORKSPACE_NOT_FOUND"
	ErrWorkloadUnavailable    = "WORKLOAD_UNAVAILABLE"
	ErrAuthenticationRequired = "AUTHENTICATION_REQUIRED"
	ErrAuthenticationFailed   = "AUTHENTICATION_FAILED"
	ErrForbidden              = "FORBIDDEN"
	ErrBadRequest             = "BAD_REQUEST"
	ErrUsageLimitExceeded     = "USAGE_LIMIT_EXCEEDED"
	ErrPolicyViolation        = "POLICY_VIOLATION"
)
