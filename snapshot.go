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
	"github.com/blaxel-ai/sdk-go/shared"
)

// SnapshotService contains methods and other services that help with interacting
// with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSnapshotService] method instead.
type SnapshotService struct {
	Options []option.RequestOption
}

// NewSnapshotService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSnapshotService(opts ...option.RequestOption) (r SnapshotService) {
	r = SnapshotService{}
	r.Options = opts
	return
}

// Captures a snapshot from a source object. The snapshot belongs to the workspace
// rather than to its source, so it survives the deletion of the object it was
// captured from, and carries what creating a sandbox or an application from it
// needs.
func (r *SnapshotService) New(ctx context.Context, body SnapshotNewParams, opts ...option.RequestOption) (res *SandboxSnapshot, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "snapshots"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a snapshot of the workspace by ID.
func (r *SnapshotService) Get(ctx context.Context, snapshotName string, opts ...option.RequestOption) (res *SandboxSnapshot, err error) {
	opts = slices.Concat(r.Options, opts)
	if snapshotName == "" {
		err = errors.New("missing required snapshotName parameter")
		return nil, err
	}
	path := fmt.Sprintf("snapshots/%s", snapshotName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns the snapshots of the workspace, newest first by default, including the
// ones whose source object has been deleted. Starting with API version 2026-04-28
// the response is wrapped in `{data, meta}` and supports cursor pagination via the
// `cursor` and `limit` query parameters; older versions keep returning a bare
// array.
func (r *SnapshotService) List(ctx context.Context, query SnapshotListParams, opts ...option.RequestOption) (res *pagination.CursorPage[SandboxSnapshot], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "snapshots"
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

// Returns the snapshots of the workspace, newest first by default, including the
// ones whose source object has been deleted. Starting with API version 2026-04-28
// the response is wrapped in `{data, meta}` and supports cursor pagination via the
// `cursor` and `limit` query parameters; older versions keep returning a bare
// array.
func (r *SnapshotService) ListAutoPaging(ctx context.Context, query SnapshotListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[SandboxSnapshot] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Deletes a snapshot. There is a single snapshot object, so this removes it for
// the whole workspace, whether or not the object it was captured from still
// exists.
func (r *SnapshotService) Delete(ctx context.Context, snapshotName string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if snapshotName == "" {
		err = errors.New("missing required snapshotName parameter")
		return err
	}
	path := fmt.Sprintf("snapshots/%s", snapshotName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Creates a new sandbox or application from a snapshot. The snapshot is enough on
// its own, so this works after the object it was captured from has been deleted.
func (r *SnapshotService) Fork(ctx context.Context, snapshotName string, params SnapshotForkParams, opts ...option.RequestOption) (res *SnapshotForkResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	if snapshotName == "" {
		err = errors.New("missing required snapshotName parameter")
		return nil, err
	}
	path := fmt.Sprintf("snapshots/%s/fork", snapshotName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Request body for forking a sandbox into an application. Creates a new
// application or adds a canary revision to an existing one.
//
// The properties TargetName, TargetType are required.
type SandboxForkRequestParam struct {
	// Name of the target application to create or update
	TargetName string `json:"targetName" api:"required"`
	// Target resource type to fork into
	TargetType string `json:"targetType" api:"required"`
	// Custom domain for the application
	CustomDomain param.Opt[string] `json:"customDomain,omitzero"`
	// Port to expose from the sandbox
	Port param.Opt[int64] `json:"port,omitzero"`
	// URL prefix for the application
	Prefix param.Opt[string] `json:"prefix,omitzero"`
	// Snapshot ID to fork from. When set, the fork is created from that existing
	// snapshot (and an application revision references it). When omitted, a fork to a
	// sandbox copies the source sandbox's live state directly and no snapshot is
	// persisted; a fork to an application still takes a snapshot, since its revision
	// references one.
	SnapshotID param.Opt[string] `json:"snapshotId,omitzero"`
	// Traffic percentage for canary deployment (0-100). When set on an existing
	// target, creates a new revision with this traffic percentage.
	Traffic param.Opt[int64] `json:"traffic,omitzero"`
	// Environment variables the fork runs with, on top of the ones the source has. A
	// variable the source already carries takes this value in the fork, one it does
	// not is added, and every other variable of the source is kept.
	Envs []shared.EnvParam `json:"envs,omitzero"`
	paramObj
}

func (r SandboxForkRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow SandboxForkRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SandboxForkRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Response returned after forking a sandbox. Contains either the new sandbox or
// application depending on the fork type.
type SandboxForkResponse struct {
	// Name of the created or updated resource
	Name string `json:"name"`
	// The snapshot ID the fork was created from. Set only when the fork went through a
	// snapshot, meaning an explicit snapshotId was supplied or the fork target is an
	// application. Empty when the fork copied the source sandbox's live state
	// directly.
	SnapshotID string `json:"snapshotId"`
	// Type of resource that was created (sandbox or application)
	//
	// Any of "sandbox", "application".
	Type SandboxForkResponseType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		SnapshotID  respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SandboxForkResponse) RawJSON() string { return r.JSON.raw }
func (r *SandboxForkResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Type of resource that was created (sandbox or application)
type SandboxForkResponseType string

const (
	SandboxForkResponseTypeSandbox     SandboxForkResponseType = "sandbox"
	SandboxForkResponseTypeApplication SandboxForkResponseType = "application"
)

// A point-in-time snapshot of a sandbox. It is a workspace-level object: it
// outlives the sandbox it was captured from, and can be restored onto a sandbox or
// forked into a new sandbox or application on its own.
type SandboxSnapshot struct {
	// Identifier of the snapshot, unique in the workspace. Workspace-level routes
	// address the snapshot by it.
	ID string `json:"id" api:"required"`
	// When the snapshot was created
	CreatedAt string `json:"createdAt" api:"required"`
	// Display name of the snapshot, unique among the snapshots of the sandbox it was
	// captured from. Defaults to the identifier.
	Name string `json:"name" api:"required"`
	// Status of the snapshot (pending, ready, failed)
	Status string `json:"status" api:"required"`
	// Workspace owning the snapshot
	Workspace string `json:"workspace" api:"required"`
	// Who created the snapshot
	CreatedBy string `json:"createdBy"`
	// Name of the source sandbox. Kept for compatibility, read source.name instead.
	SandboxName string `json:"sandboxName"`
	// The object a snapshot was captured from.
	Source SandboxSnapshotSource `json:"source"`
	// The configuration a snapshot carries, so a sandbox or an application can be
	// created from it once its source object is gone.
	Spec SandboxSnapshotSpec `json:"spec"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		Status      respjson.Field
		Workspace   respjson.Field
		CreatedBy   respjson.Field
		SandboxName respjson.Field
		Source      respjson.Field
		Spec        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SandboxSnapshot) RawJSON() string { return r.JSON.raw }
func (r *SandboxSnapshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request body for creating a snapshot. The source object is required at the root
// endpoint and implied by the path on the nested one.
type SandboxSnapshotRequestParam struct {
	// Display name of the snapshot, unique among the snapshots of the sandbox it is
	// captured from; the same name may be reused on another sandbox. Defaults to the
	// snapshot identifier when absent.
	Name param.Opt[string] `json:"name,omitzero"`
	// The object a snapshot was captured from.
	Source SandboxSnapshotSourceParam `json:"source,omitzero"`
	paramObj
}

func (r SandboxSnapshotRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow SandboxSnapshotRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SandboxSnapshotRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The object a snapshot was captured from.
type SandboxSnapshotSource struct {
	// Name of the object the snapshot was captured from
	Name string `json:"name" api:"required"`
	// Whether the source object has since been deleted. The snapshot stays usable, the
	// link is only kept for context.
	Deleted bool `json:"deleted"`
	// Kind of the object the snapshot was captured from. Defaults to sandbox, the only
	// kind that can be captured today.
	//
	// Any of "sandbox".
	Kind SandboxSnapshotSourceKind `json:"kind"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Deleted     respjson.Field
		Kind        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SandboxSnapshotSource) RawJSON() string { return r.JSON.raw }
func (r *SandboxSnapshotSource) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SandboxSnapshotSource to a SandboxSnapshotSourceParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SandboxSnapshotSourceParam.Overrides()
func (r SandboxSnapshotSource) ToParam() SandboxSnapshotSourceParam {
	return param.Override[SandboxSnapshotSourceParam](json.RawMessage(r.RawJSON()))
}

// Kind of the object the snapshot was captured from. Defaults to sandbox, the only
// kind that can be captured today.
type SandboxSnapshotSourceKind string

const (
	SandboxSnapshotSourceKindSandbox SandboxSnapshotSourceKind = "sandbox"
)

// The object a snapshot was captured from.
//
// The property Name is required.
type SandboxSnapshotSourceParam struct {
	// Name of the object the snapshot was captured from
	Name string `json:"name" api:"required"`
	// Kind of the object the snapshot was captured from. Defaults to sandbox, the only
	// kind that can be captured today.
	//
	// Any of "sandbox".
	Kind SandboxSnapshotSourceKind `json:"kind,omitzero"`
	paramObj
}

func (r SandboxSnapshotSourceParam) MarshalJSON() (data []byte, err error) {
	type shadow SandboxSnapshotSourceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SandboxSnapshotSourceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration a snapshot carries, so a sandbox or an application can be
// created from it once its source object is gone.
type SandboxSnapshotSpec struct {
	// Infrastructure generation the snapshot was captured on. A snapshot only restores
	// on the generation it came from.
	//
	// Any of "mk2", "mk3".
	Generation SandboxSnapshotSpecGeneration `json:"generation"`
	// Image the source object ran
	Image string `json:"image"`
	// Memory in MB the source object ran with
	Memory int64 `json:"memory"`
	// Set of ports for a resource
	Ports []Port `json:"ports"`
	// Region holding the snapshot. Restores and forks land in it.
	Region  string             `json:"region"`
	Volumes []VolumeAttachment `json:"volumes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Generation  respjson.Field
		Image       respjson.Field
		Memory      respjson.Field
		Ports       respjson.Field
		Region      respjson.Field
		Volumes     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SandboxSnapshotSpec) RawJSON() string { return r.JSON.raw }
func (r *SandboxSnapshotSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Infrastructure generation the snapshot was captured on. A snapshot only restores
// on the generation it came from.
type SandboxSnapshotSpecGeneration string

const (
	SandboxSnapshotSpecGenerationMk2 SandboxSnapshotSpecGeneration = "mk2"
	SandboxSnapshotSpecGenerationMk3 SandboxSnapshotSpecGeneration = "mk3"
)

// SnapshotForkResponseUnion contains all possible properties and values from
// [SandboxForkResponse], [[]SandboxForkResponse].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfSandboxForkResponseArray]
type SnapshotForkResponseUnion struct {
	// This field will be present if the value is a [[]SandboxForkResponse] instead of
	// an object.
	OfSandboxForkResponseArray []SandboxForkResponse `json:",inline"`
	// This field is from variant [SandboxForkResponse].
	Name string `json:"name"`
	// This field is from variant [SandboxForkResponse].
	SnapshotID string `json:"snapshotId"`
	// This field is from variant [SandboxForkResponse].
	Type SandboxForkResponseType `json:"type"`
	JSON struct {
		OfSandboxForkResponseArray respjson.Field
		Name                       respjson.Field
		SnapshotID                 respjson.Field
		Type                       respjson.Field
		raw                        string
	} `json:"-"`
}

func (u SnapshotForkResponseUnion) AsSandboxForkResponse() (v SandboxForkResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SnapshotForkResponseUnion) AsSandboxForkResponseArray() (v []SandboxForkResponse) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SnapshotForkResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *SnapshotForkResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SnapshotNewParams struct {
	// Request body for creating a snapshot. The source object is required at the root
	// endpoint and implied by the path on the nested one.
	SandboxSnapshotRequest SandboxSnapshotRequestParam
	paramObj
}

func (r SnapshotNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SandboxSnapshotRequest)
}
func (r *SnapshotNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SnapshotListParams struct {
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
	Anchor SnapshotListParamsAnchor `query:"anchor,omitzero" json:"-"`
	// Sort spec, formatted as `<key>:<direction>`. Allowed values are `createdAt:desc`
	// (default), `createdAt:asc`, `name:asc`, `name:desc`. The cursor fingerprint is
	// bound to the sort, so a cursor opened with one value cannot be reused with
	// another. Only honoured starting on Blaxel-Version 2026-04-28.
	//
	// Any of "createdAt:desc", "createdAt:asc", "name:asc", "name:desc".
	Sort SnapshotListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SnapshotListParams]'s query parameters as `url.Values`.
func (r SnapshotListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Start from a known pagination boundary. `end` is only supported for `createdAt`
// listings (asc or desc) and returns the tail page directly without walking every
// cursor from the first page.
type SnapshotListParamsAnchor string

const (
	SnapshotListParamsAnchorEnd SnapshotListParamsAnchor = "end"
)

// Sort spec, formatted as `<key>:<direction>`. Allowed values are `createdAt:desc`
// (default), `createdAt:asc`, `name:asc`, `name:desc`. The cursor fingerprint is
// bound to the sort, so a cursor opened with one value cannot be reused with
// another. Only honoured starting on Blaxel-Version 2026-04-28.
type SnapshotListParamsSort string

const (
	SnapshotListParamsSortCreatedAtDesc SnapshotListParamsSort = "createdAt:desc"
	SnapshotListParamsSortCreatedAtAsc  SnapshotListParamsSort = "createdAt:asc"
	SnapshotListParamsSortNameAsc       SnapshotListParamsSort = "name:asc"
	SnapshotListParamsSortNameDesc      SnapshotListParamsSort = "name:desc"
)

type SnapshotForkParams struct {
	// Request body for forking a sandbox into an application. Creates a new
	// application or adds a canary revision to an existing one.
	SandboxForkRequest SandboxForkRequestParam
	// Bulk fork. When set, `count` sandboxes are created from the snapshot with
	// server-generated names and the response is an array of fork results (even for
	// `count=1`). The quota is validated for the whole batch before anything is
	// created and the request never returns a partial batch. Only supported for the
	// sandbox target type; cannot be combined with `targetName`.
	Count param.Opt[int64] `query:"count,omitzero" json:"-"`
	paramObj
}

func (r SnapshotForkParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SandboxForkRequest)
}
func (r *SnapshotForkParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [SnapshotForkParams]'s query parameters as `url.Values`.
func (r SnapshotForkParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
