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
	shimjson "github.com/blaxel-ai/sdk-go/internal/encoding/json"
	"github.com/blaxel-ai/sdk-go/internal/requestconfig"
	"github.com/blaxel-ai/sdk-go/option"
	"github.com/blaxel-ai/sdk-go/packages/pagination"
	"github.com/blaxel-ai/sdk-go/packages/param"
	"github.com/blaxel-ai/sdk-go/packages/respjson"
)

// VolumeService contains methods and other services that help with interacting
// with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVolumeService] method instead.
type VolumeService struct {
	Options []option.RequestOption
}

// NewVolumeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewVolumeService(opts ...option.RequestOption) (r VolumeService) {
	r = VolumeService{}
	r.Options = opts
	return
}

// Creates a new persistent storage volume that can be attached to sandboxes.
// Volumes must be created in a specific region and can only attach to sandboxes in
// the same region.
func (r *VolumeService) New(ctx context.Context, body VolumeNewParams, opts ...option.RequestOption) (res *Volume, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "volumes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns detailed information about a volume including its size, region,
// attachment status, and any events history.
func (r *VolumeService) Get(ctx context.Context, volumeName string, opts ...option.RequestOption) (res *Volume, err error) {
	opts = slices.Concat(r.Options, opts)
	if volumeName == "" {
		err = errors.New("missing required volumeName parameter")
		return nil, err
	}
	path := fmt.Sprintf("volumes/%s", volumeName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates a volume.
func (r *VolumeService) Update(ctx context.Context, volumeName string, body VolumeUpdateParams, opts ...option.RequestOption) (res *Volume, err error) {
	opts = slices.Concat(r.Options, opts)
	if volumeName == "" {
		err = errors.New("missing required volumeName parameter")
		return nil, err
	}
	path := fmt.Sprintf("volumes/%s", volumeName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Returns persistent storage volumes in the workspace. Volumes can be attached to
// sandboxes for durable file storage that persists across sessions and sandbox
// deletions. Starting with API version 2026-04-28 the response is wrapped in
// `{data, meta}` and supports cursor pagination via the `cursor` and `limit` query
// parameters; older versions keep returning a bare array of volumes.
func (r *VolumeService) List(ctx context.Context, query VolumeListParams, opts ...option.RequestOption) (res *pagination.CursorPage[VolumeListResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "volumes"
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

// Returns persistent storage volumes in the workspace. Volumes can be attached to
// sandboxes for durable file storage that persists across sessions and sandbox
// deletions. Starting with API version 2026-04-28 the response is wrapped in
// `{data, meta}` and supports cursor pagination via the `cursor` and `limit` query
// parameters; older versions keep returning a bare array of volumes.
func (r *VolumeService) ListAutoPaging(ctx context.Context, query VolumeListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[VolumeListResponse] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Permanently deletes a volume and all its data. The volume must not be attached
// to any sandbox. This action cannot be undone.
func (r *VolumeService) Delete(ctx context.Context, volumeName string, opts ...option.RequestOption) (res *Volume, err error) {
	opts = slices.Concat(r.Options, opts)
	if volumeName == "" {
		err = errors.New("missing required volumeName parameter")
		return nil, err
	}
	path := fmt.Sprintf("volumes/%s", volumeName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Returns an active volume matching the given external ID. If no active volume is
// found, returns 404.
func (r *VolumeService) GetByExternalID(ctx context.Context, externalID string, opts ...option.RequestOption) (res *Volume, err error) {
	opts = slices.Concat(r.Options, opts)
	if externalID == "" {
		err = errors.New("missing required externalId parameter")
		return nil, err
	}
	path := fmt.Sprintf("volumes/by-external-id/%s", externalID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Persistent storage volume that can be attached to sandboxes for durable file
// storage across sessions. Volumes survive sandbox deletion and can be reattached
// to new sandboxes.
type Volume struct {
	// Common metadata fields shared by all Blaxel resources including name, labels,
	// timestamps, and ownership information
	Metadata Metadata `json:"metadata" api:"required"`
	// Immutable volume configuration set at creation time (size and region cannot be
	// changed after creation)
	Spec VolumeSpec `json:"spec" api:"required"`
	// Events happening on a resource deployed on Blaxel
	Events []CoreEvent `json:"events"`
	// Current runtime state of the volume including attachment status
	State VolumeState `json:"state"`
	// Volume status computed from events
	Status string `json:"status"`
	// Timestamp when the volume was marked for termination
	TerminatedAt string `json:"terminatedAt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metadata     respjson.Field
		Spec         respjson.Field
		Events       respjson.Field
		State        respjson.Field
		Status       respjson.Field
		TerminatedAt respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Volume) RawJSON() string { return r.JSON.raw }
func (r *Volume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Volume to a VolumeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// VolumeParam.Overrides()
func (r Volume) ToParam() VolumeParam {
	return param.Override[VolumeParam](json.RawMessage(r.RawJSON()))
}

// Current runtime state of the volume including attachment status
type VolumeState struct {
	// Resource currently using this volume in format "type:name" (e.g.,
	// "sandbox:my-sandbox"). Empty if not attached.
	AttachedTo string `json:"attachedTo"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AttachedTo  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VolumeState) RawJSON() string { return r.JSON.raw }
func (r *VolumeState) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Persistent storage volume that can be attached to sandboxes for durable file
// storage across sessions. Volumes survive sandbox deletion and can be reattached
// to new sandboxes.
//
// The properties Metadata, Spec are required.
type VolumeParam struct {
	// Common metadata fields shared by all Blaxel resources including name, labels,
	// timestamps, and ownership information
	Metadata MetadataParam `json:"metadata,omitzero" api:"required"`
	// Immutable volume configuration set at creation time (size and region cannot be
	// changed after creation)
	Spec VolumeSpecParam `json:"spec,omitzero" api:"required"`
	paramObj
}

func (r VolumeParam) MarshalJSON() (data []byte, err error) {
	type shadow VolumeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VolumeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current runtime state of the volume including attachment status
type VolumeStateParam struct {
	// Resource currently using this volume in format "type:name" (e.g.,
	// "sandbox:my-sandbox"). Empty if not attached.
	AttachedTo param.Opt[string] `json:"attachedTo,omitzero"`
	paramObj
}

func (r VolumeStateParam) MarshalJSON() (data []byte, err error) {
	type shadow VolumeStateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VolumeStateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Immutable volume configuration set at creation time (size and region cannot be
// changed after creation)
type VolumeSpec struct {
	// The internal infrastructure resource identifier for this volume
	InfrastructureID string `json:"infrastructureId"`
	// Deployment region for the volume (e.g., us-pdx-1, eu-lon-1). Must match the
	// region of sandboxes it attaches to.
	Region string `json:"region"`
	// Storage capacity in megabytes. Can be increased after creation but not
	// decreased.
	Size int64 `json:"size"`
	// Volume template to initialize from, with optional revision (e.g., "mytemplate:1"
	// or "mytemplate:latest")
	Template string `json:"template"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InfrastructureID respjson.Field
		Region           respjson.Field
		Size             respjson.Field
		Template         respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VolumeSpec) RawJSON() string { return r.JSON.raw }
func (r *VolumeSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this VolumeSpec to a VolumeSpecParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// VolumeSpecParam.Overrides()
func (r VolumeSpec) ToParam() VolumeSpecParam {
	return param.Override[VolumeSpecParam](json.RawMessage(r.RawJSON()))
}

// Immutable volume configuration set at creation time (size and region cannot be
// changed after creation)
type VolumeSpecParam struct {
	// Deployment region for the volume (e.g., us-pdx-1, eu-lon-1). Must match the
	// region of sandboxes it attaches to.
	Region param.Opt[string] `json:"region,omitzero"`
	// Storage capacity in megabytes. Can be increased after creation but not
	// decreased.
	Size param.Opt[int64] `json:"size,omitzero"`
	// Volume template to initialize from, with optional revision (e.g., "mytemplate:1"
	// or "mytemplate:latest")
	Template param.Opt[string] `json:"template,omitzero"`
	paramObj
}

func (r VolumeSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow VolumeSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VolumeSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// LiteVolume is the listing-shape projection of a Volume. Drops events to keep
// page payloads small.
type VolumeListResponse struct {
	// Compact metadata for a Volume, returned in listing responses.
	Metadata VolumeListResponseMetadata `json:"metadata"`
	// Compact spec for a Volume, returned in listing responses.
	Spec VolumeListResponseSpec `json:"spec"`
	// Current runtime state of the volume including attachment status
	State VolumeListResponseState `json:"state"`
	// Computed status of the volume.
	Status string `json:"status"`
	// Termination timestamp for soft-deleted volumes.
	TerminatedAt string `json:"terminatedAt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metadata     respjson.Field
		Spec         respjson.Field
		State        respjson.Field
		Status       respjson.Field
		TerminatedAt respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VolumeListResponse) RawJSON() string { return r.JSON.raw }
func (r *VolumeListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Compact metadata for a Volume, returned in listing responses.
type VolumeListResponseMetadata struct {
	CreatedAt   string `json:"createdAt"`
	DisplayName string `json:"displayName"`
	// Caller-owned identifier for external lookups.
	ExternalID string `json:"externalId"`
	Name       string `json:"name"`
	UpdatedAt  string `json:"updatedAt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		DisplayName respjson.Field
		ExternalID  respjson.Field
		Name        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VolumeListResponseMetadata) RawJSON() string { return r.JSON.raw }
func (r *VolumeListResponseMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Compact spec for a Volume, returned in listing responses.
type VolumeListResponseSpec struct {
	// Region the volume is provisioned in.
	Region string `json:"region"`
	// Volume size in gigabytes.
	Size int64 `json:"size"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Region      respjson.Field
		Size        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VolumeListResponseSpec) RawJSON() string { return r.JSON.raw }
func (r *VolumeListResponseSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current runtime state of the volume including attachment status
type VolumeListResponseState struct {
	// Resource currently using this volume in format "type:name" (e.g.,
	// "sandbox:my-sandbox"). Empty if not attached.
	AttachedTo string `json:"attachedTo"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AttachedTo  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VolumeListResponseState) RawJSON() string { return r.JSON.raw }
func (r *VolumeListResponseState) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VolumeNewParams struct {
	// Persistent storage volume that can be attached to sandboxes for durable file
	// storage across sessions. Volumes survive sandbox deletion and can be reattached
	// to new sandboxes.
	Volume VolumeParam
	paramObj
}

func (r VolumeNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Volume)
}
func (r *VolumeNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VolumeUpdateParams struct {
	// Persistent storage volume that can be attached to sandboxes for durable file
	// storage across sessions. Volumes survive sandbox deletion and can be reattached
	// to new sandboxes.
	Volume VolumeParam
	paramObj
}

func (r VolumeUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Volume)
}
func (r *VolumeUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VolumeListParams struct {
	// Opaque cursor returned by a previous response's meta.nextCursor. Only valid for
	// the same query (workspace + filters); the server rejects cursors bound to a
	// different query or older than 24h. Omit on the first page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Filter volumes by external ID. When set, only volumes matching this caller-owned
	// identifier are returned.
	ExternalID param.Opt[string] `query:"externalId,omitzero" json:"-"`
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
	Anchor VolumeListParamsAnchor `query:"anchor,omitzero" json:"-"`
	// Sort spec, formatted as `<key>:<direction>`. Allowed values are `createdAt:desc`
	// (default), `createdAt:asc`, `name:asc`, `name:desc`. The cursor fingerprint is
	// bound to the sort, so a cursor opened with one value cannot be reused with
	// another. Only honoured starting on Blaxel-Version 2026-04-28.
	//
	// Any of "createdAt:desc", "createdAt:asc", "name:asc", "name:desc".
	Sort VolumeListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [VolumeListParams]'s query parameters as `url.Values`.
func (r VolumeListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Start from a known pagination boundary. `end` is only supported for `createdAt`
// listings (asc or desc) and returns the tail page directly without walking every
// cursor from the first page.
type VolumeListParamsAnchor string

const (
	VolumeListParamsAnchorEnd VolumeListParamsAnchor = "end"
)

// Sort spec, formatted as `<key>:<direction>`. Allowed values are `createdAt:desc`
// (default), `createdAt:asc`, `name:asc`, `name:desc`. The cursor fingerprint is
// bound to the sort, so a cursor opened with one value cannot be reused with
// another. Only honoured starting on Blaxel-Version 2026-04-28.
type VolumeListParamsSort string

const (
	VolumeListParamsSortCreatedAtDesc VolumeListParamsSort = "createdAt:desc"
	VolumeListParamsSortCreatedAtAsc  VolumeListParamsSort = "createdAt:asc"
	VolumeListParamsSortNameAsc       VolumeListParamsSort = "name:asc"
	VolumeListParamsSortNameDesc      VolumeListParamsSort = "name:desc"
)
