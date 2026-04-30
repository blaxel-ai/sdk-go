package apierror

// Gateway error codes returned by the Blaxel platform.
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
