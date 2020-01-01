// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blaxel

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/blaxel-ai/sdk-go/internal/apijson"
	"github.com/blaxel-ai/sdk-go/internal/apiquery"
	"github.com/blaxel-ai/sdk-go/internal/requestconfig"
	"github.com/blaxel-ai/sdk-go/option"
	"github.com/blaxel-ai/sdk-go/packages/pagination"
	"github.com/blaxel-ai/sdk-go/packages/param"
	"github.com/blaxel-ai/sdk-go/packages/respjson"
)

// DriveService contains methods and other services that help with interacting with
// the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDriveService] method instead.
type DriveService struct {
	Options []option.RequestOption
	Mount   DriveMountService
}

// NewDriveService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDriveService(opts ...option.RequestOption) (r DriveService) {
	r = DriveService{}
	r.Options = opts
	r.Mount = NewDriveMountService(opts...)
	return
}

// Creates a new drive in the workspace. Drives can be buckets and can be mounted
// at runtime to sandboxes.
func (r *DriveService) New(ctx context.Context, body DriveNewParams, opts ...option.RequestOption) (res *DriveNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "drives"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Retrieves details of a specific drive including its status and events.
func (r *DriveService) Get(ctx context.Context, driveName string, opts ...option.RequestOption) (res *DriveGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if driveName == "" {
		err = errors.New("missing required driveName parameter")
		return nil, err
	}
	path := fmt.Sprintf("drives/%s", driveName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates an existing drive. Metadata fields like displayName and labels can be
// changed. Size can be set if not already configured.
func (r *DriveService) Update(ctx context.Context, driveName string, body DriveUpdateParams, opts ...option.RequestOption) (res *DriveUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if driveName == "" {
		err = errors.New("missing required driveName parameter")
		return nil, err
	}
	path := fmt.Sprintf("drives/%s", driveName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Returns all drives in the workspace. Drives provide persistent storage that can
// be attached to agents, functions, and sandboxes. Starting with API version
// 2026-04-28, the response wraps items in `{data, meta}` and supports cursor
// pagination via the `cursor` and `limit` query parameters; older versions keep
// returning a bare array with all drives.
func (r *DriveService) List(ctx context.Context, query DriveListParams, opts ...option.RequestOption) (res *pagination.CursorPage[DriveListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "drives"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Returns all drives in the workspace. Drives provide persistent storage that can
// be attached to agents, functions, and sandboxes. Starting with API version
// 2026-04-28, the response wraps items in `{data, meta}` and supports cursor
// pagination via the `cursor` and `limit` query parameters; older versions keep
// returning a bare array with all drives.
func (r *DriveService) ListAutoPaging(ctx context.Context, query DriveListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[DriveListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Deletes a drive immediately. The drive record is removed from the database
// synchronously.
func (r *DriveService) Delete(ctx context.Context, driveName string, opts ...option.RequestOption) (res *DriveDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if driveName == "" {
		err = errors.New("missing required driveName parameter")
		return nil, err
	}
	path := fmt.Sprintf("drives/%s", driveName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Issues a short-lived JWT access token scoped to a specific drive. The token can
// be used as Bearer authentication for direct S3 operations against the drive's
// bucket.
func (r *DriveService) NewAccessToken(ctx context.Context, driveName string, opts ...option.RequestOption) (res *DriveNewAccessTokenResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if driveName == "" {
		err = errors.New("missing required driveName parameter")
		return nil, err
	}
	path := fmt.Sprintf("drives/%s/access-token", driveName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Returns the JSON Web Key Set containing the Ed25519 public key used to verify
// drive access tokens. Other S3-compatible storage can use this endpoint to
// validate Bearer tokens.
func (r *DriveService) GetJwks(ctx context.Context, opts ...option.RequestOption) (res *DriveGetJwksResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "drives/jwks.json"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Immutable drive configuration set at creation time
type DriveSpec struct {
	// The internal infrastructure resource identifier for this drive (bucket name)
	InfrastructureID string `json:"infrastructureId"`
	// Permissions controlling which workloads can access this drive. Empty means all
	// workloads in the workspace can access the drive. Maximum 3 permissions.
	Permissions []DriveSpecPermission `json:"permissions"`
	// Deployment region for the drive (e.g., us-pdx-1, eu-lon-1). Must match the
	// region of resources it attaches to.
	Region string `json:"region"`
	// Optional size limit for the drive in GB. If not specified, drive has no size
	// limit.
	Size int64 `json:"size"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InfrastructureID respjson.Field
		Permissions      respjson.Field
		Region           respjson.Field
		Size             respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DriveSpec) RawJSON() string { return r.JSON.raw }
func (r *DriveSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this DriveSpec to a DriveSpecParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// DriveSpecParam.Overrides()
func (r DriveSpec) ToParam() DriveSpecParam {
	return param.Override[DriveSpecParam](json.RawMessage(r.RawJSON()))
}

// Permission that controls which workloads can access a drive. A workload must
// have ALL specified labels (AND logic). Permissions are evaluated with OR logic —
// the first matching permission grants access.
type DriveSpecPermission struct {
	// Labels that the workload must have. All labels must match (AND logic). Empty
	// labels match all workloads.
	Labels any `json:"labels"`
	// Access mode granted by this permission
	//
	// Any of "read", "read-write".
	Mode string `json:"mode"`
	// Subfolder path to restrict access to. Defaults to / (full drive).
	Path string `json:"path"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Labels      respjson.Field
		Mode        respjson.Field
		Path        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DriveSpecPermission) RawJSON() string { return r.JSON.raw }
func (r *DriveSpecPermission) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Immutable drive configuration set at creation time
type DriveSpecParam struct {
	// Deployment region for the drive (e.g., us-pdx-1, eu-lon-1). Must match the
	// region of resources it attaches to.
	Region param.Opt[string] `json:"region,omitzero"`
	// Optional size limit for the drive in GB. If not specified, drive has no size
	// limit.
	Size param.Opt[int64] `json:"size,omitzero"`
	// Permissions controlling which workloads can access this drive. Empty means all
	// workloads in the workspace can access the drive. Maximum 3 permissions.
	Permissions []DriveSpecPermissionParam `json:"permissions,omitzero"`
	paramObj
}

func (r DriveSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow DriveSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DriveSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Permission that controls which workloads can access a drive. A workload must
// have ALL specified labels (AND logic). Permissions are evaluated with OR logic —
// the first matching permission grants access.
type DriveSpecPermissionParam struct {
	// Subfolder path to restrict access to. Defaults to / (full drive).
	Path param.Opt[string] `json:"path,omitzero"`
	// Labels that the workload must have. All labels must match (AND logic). Empty
	// labels match all workloads.
	Labels any `json:"labels,omitzero"`
	// Access mode granted by this permission
	//
	// Any of "read", "read-write".
	Mode string `json:"mode,omitzero"`
	paramObj
}

func (r DriveSpecPermissionParam) MarshalJSON() (data []byte, err error) {
	type shadow DriveSpecPermissionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DriveSpecPermissionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[DriveSpecPermissionParam](
		"mode", "read", "read-write",
	)
}

// Drive providing persistent storage that can be attached to agents, functions,
// and sandboxes. Drives can be mounted at runtime via the sbx API.
type DriveNewResponse struct {
	// Common metadata fields shared by all Blaxel resources including name, labels,
	// timestamps, and ownership information
	Metadata Metadata `json:"metadata" api:"required"`
	// Immutable drive configuration set at creation time
	Spec DriveSpec `json:"spec" api:"required"`
	// Events happening on a resource deployed on Blaxel
	Events []CoreEvent `json:"events"`
	// Current runtime state of the drive
	State DriveNewResponseState `json:"state"`
	// Drive status computed from events
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metadata    respjson.Field
		Spec        respjson.Field
		Events      respjson.Field
		State       respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DriveNewResponse) RawJSON() string { return r.JSON.raw }
func (r *DriveNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current runtime state of the drive
type DriveNewResponseState struct {
	// S3-compatible endpoint URL for accessing drive contents
	S3URL string `json:"s3Url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		S3URL       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DriveNewResponseState) RawJSON() string { return r.JSON.raw }
func (r *DriveNewResponseState) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Drive providing persistent storage that can be attached to agents, functions,
// and sandboxes. Drives can be mounted at runtime via the sbx API.
type DriveGetResponse struct {
	// Common metadata fields shared by all Blaxel resources including name, labels,
	// timestamps, and ownership information
	Metadata Metadata `json:"metadata" api:"required"`
	// Immutable drive configuration set at creation time
	Spec DriveSpec `json:"spec" api:"required"`
	// Events happening on a resource deployed on Blaxel
	Events []CoreEvent `json:"events"`
	// Current runtime state of the drive
	State DriveGetResponseState `json:"state"`
	// Drive status computed from events
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metadata    respjson.Field
		Spec        respjson.Field
		Events      respjson.Field
		State       respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DriveGetResponse) RawJSON() string { return r.JSON.raw }
func (r *DriveGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current runtime state of the drive
type DriveGetResponseState struct {
	// S3-compatible endpoint URL for accessing drive contents
	S3URL string `json:"s3Url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		S3URL       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DriveGetResponseState) RawJSON() string { return r.JSON.raw }
func (r *DriveGetResponseState) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Drive providing persistent storage that can be attached to agents, functions,
// and sandboxes. Drives can be mounted at runtime via the sbx API.
type DriveUpdateResponse struct {
	// Common metadata fields shared by all Blaxel resources including name, labels,
	// timestamps, and ownership information
	Metadata Metadata `json:"metadata" api:"required"`
	// Immutable drive configuration set at creation time
	Spec DriveSpec `json:"spec" api:"required"`
	// Events happening on a resource deployed on Blaxel
	Events []CoreEvent `json:"events"`
	// Current runtime state of the drive
	State DriveUpdateResponseState `json:"state"`
	// Drive status computed from events
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metadata    respjson.Field
		Spec        respjson.Field
		Events      respjson.Field
		State       respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DriveUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *DriveUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current runtime state of the drive
type DriveUpdateResponseState struct {
	// S3-compatible endpoint URL for accessing drive contents
	S3URL string `json:"s3Url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		S3URL       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DriveUpdateResponseState) RawJSON() string { return r.JSON.raw }
func (r *DriveUpdateResponseState) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Drive providing persistent storage that can be attached to agents, functions,
// and sandboxes. Drives can be mounted at runtime via the sbx API.
type DriveListResponse struct {
	// Common metadata fields shared by all Blaxel resources including name, labels,
	// timestamps, and ownership information
	Metadata Metadata `json:"metadata" api:"required"`
	// Immutable drive configuration set at creation time
	Spec DriveSpec `json:"spec" api:"required"`
	// Events happening on a resource deployed on Blaxel
	Events []CoreEvent `json:"events"`
	// Current runtime state of the drive
	State DriveListResponseState `json:"state"`
	// Drive status computed from events
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metadata    respjson.Field
		Spec        respjson.Field
		Events      respjson.Field
		State       respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DriveListResponse) RawJSON() string { return r.JSON.raw }
func (r *DriveListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current runtime state of the drive
type DriveListResponseState struct {
	// S3-compatible endpoint URL for accessing drive contents
	S3URL string `json:"s3Url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		S3URL       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DriveListResponseState) RawJSON() string { return r.JSON.raw }
func (r *DriveListResponseState) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DriveDeleteResponse struct {
	Message string `json:"message"`
	Name    string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DriveDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *DriveDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DriveNewAccessTokenResponse struct {
	AccessToken string  `json:"access_token"`
	ExpiresIn   float64 `json:"expires_in"`
	TokenType   string  `json:"token_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccessToken respjson.Field
		ExpiresIn   respjson.Field
		TokenType   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DriveNewAccessTokenResponse) RawJSON() string { return r.JSON.raw }
func (r *DriveNewAccessTokenResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DriveGetJwksResponse struct {
	Keys []any `json:"keys"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Keys        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DriveGetJwksResponse) RawJSON() string { return r.JSON.raw }
func (r *DriveGetJwksResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DriveNewParams struct {
	// Common metadata fields shared by all Blaxel resources including name, labels,
	// timestamps, and ownership information
	Metadata MetadataParam `json:"metadata,omitzero" api:"required"`
	// Immutable drive configuration set at creation time
	Spec DriveSpecParam `json:"spec,omitzero" api:"required"`
	paramObj
}

func (r DriveNewParams) MarshalJSON() (data []byte, err error) {
	type shadow DriveNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DriveNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DriveUpdateParams struct {
	// Common metadata fields shared by all Blaxel resources including name, labels,
	// timestamps, and ownership information
	Metadata MetadataParam `json:"metadata,omitzero" api:"required"`
	// Immutable drive configuration set at creation time
	Spec DriveSpecParam `json:"spec,omitzero" api:"required"`
	paramObj
}

func (r DriveUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow DriveUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DriveUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DriveListParams struct {
	// Opaque cursor returned by a previous response's meta.nextCursor. Only valid for
	// the same query (workspace + filters); the server rejects cursors bound to a
	// different query or older than 24h. Omit on the first page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of items to return per page. Defaults to 50, clamped to 200.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Substring search across `metadata.name`, `metadata.displayName` and labels
	// (keys + values). Trimmed and lowercased server-side; queries shorter than 2
	// characters fall back to the unfiltered listing. Bound into the cursor
	// fingerprint so a cursor opened with one query cannot be reused with another.
	// Only honoured starting on Blaxel-Version 2026-04-28.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Start from a known pagination boundary. `end` is only supported for `createdAt`
	// listings (asc or desc) and returns the tail page directly without walking every
	// cursor from the first page.
	//
	// Any of "end".
	Anchor DriveListParamsAnchor `query:"anchor,omitzero" json:"-"`
	// Sort spec, formatted as `<key>:<direction>`. Allowed values are `createdAt:desc`
	// (default), `createdAt:asc`, `name:asc`, `name:desc`. The cursor fingerprint is
	// bound to the sort, so a cursor opened with one value cannot be reused with
	// another. Only honoured starting on Blaxel-Version 2026-04-28.
	//
	// Any of "createdAt:desc", "createdAt:asc", "name:asc", "name:desc".
	Sort DriveListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [DriveListParams]'s query parameters as `url.Values`.
func (r DriveListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Start from a known pagination boundary. `end` is only supported for `createdAt`
// listings (asc or desc) and returns the tail page directly without walking every
// cursor from the first page.
type DriveListParamsAnchor string

const (
	DriveListParamsAnchorEnd DriveListParamsAnchor = "end"
)

// Sort spec, formatted as `<key>:<direction>`. Allowed values are `createdAt:desc`
// (default), `createdAt:asc`, `name:asc`, `name:desc`. The cursor fingerprint is
// bound to the sort, so a cursor opened with one value cannot be reused with
// another. Only honoured starting on Blaxel-Version 2026-04-28.
type DriveListParamsSort string

const (
	DriveListParamsSortCreatedAtDesc DriveListParamsSort = "createdAt:desc"
	DriveListParamsSortCreatedAtAsc  DriveListParamsSort = "createdAt:asc"
	DriveListParamsSortNameAsc       DriveListParamsSort = "name:asc"
	DriveListParamsSortNameDesc      DriveListParamsSort = "name:desc"
)
