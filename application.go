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

// ApplicationService contains methods and other services that help with
// interacting with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewApplicationService] method instead.
type ApplicationService struct {
	Options []option.RequestOption
}

// NewApplicationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewApplicationService(opts ...option.RequestOption) (r ApplicationService) {
	r = ApplicationService{}
	r.Options = opts
	return
}

// Creates a new application deployment. Applications are long-running workloads
// that default to public access and mk3 generation.
func (r *ApplicationService) New(ctx context.Context, body ApplicationNewParams, opts ...option.RequestOption) (res *Application, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "applications"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns detailed information about an application including its current
// deployment status, configuration, events history, and endpoint URL.
func (r *ApplicationService) Get(ctx context.Context, applicationName string, opts ...option.RequestOption) (res *Application, err error) {
	opts = slices.Concat(r.Options, opts)
	if applicationName == "" {
		err = errors.New("missing required applicationName parameter")
		return nil, err
	}
	path := fmt.Sprintf("applications/%s", applicationName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates an application's configuration and triggers a new deployment. Changes to
// runtime settings, environment variables, or scaling parameters will be applied
// on the next deployment.
func (r *ApplicationService) Update(ctx context.Context, applicationName string, body ApplicationUpdateParams, opts ...option.RequestOption) (res *Application, err error) {
	opts = slices.Concat(r.Options, opts)
	if applicationName == "" {
		err = errors.New("missing required applicationName parameter")
		return nil, err
	}
	path := fmt.Sprintf("applications/%s", applicationName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Returns applications deployed in the workspace. Each application includes its
// deployment status, runtime configuration, and endpoint URL. Starting with API
// version 2026-04-28 the response is wrapped in `{data, meta}` and supports cursor
// pagination via the `cursor` and `limit` query parameters; older versions keep
// returning a bare array with all applications.
func (r *ApplicationService) List(ctx context.Context, query ApplicationListParams, opts ...option.RequestOption) (res *pagination.CursorPage[Application], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "applications"
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

// Returns applications deployed in the workspace. Each application includes its
// deployment status, runtime configuration, and endpoint URL. Starting with API
// version 2026-04-28 the response is wrapped in `{data, meta}` and supports cursor
// pagination via the `cursor` and `limit` query parameters; older versions keep
// returning a bare array with all applications.
func (r *ApplicationService) ListAutoPaging(ctx context.Context, query ApplicationListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[Application] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Permanently deletes an application and all its deployment history. The
// application endpoint will immediately stop responding. This action cannot be
// undone.
func (r *ApplicationService) Delete(ctx context.Context, applicationName string, opts ...option.RequestOption) (res *Application, err error) {
	opts = slices.Concat(r.Options, opts)
	if applicationName == "" {
		err = errors.New("missing required applicationName parameter")
		return nil, err
	}
	path := fmt.Sprintf("applications/%s", applicationName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// List all application revisions
func (r *ApplicationService) ListRevisions(ctx context.Context, applicationName string, opts ...option.RequestOption) (res *[]AppRevision, err error) {
	opts = slices.Concat(r.Options, opts)
	if applicationName == "" {
		err = errors.New("missing required applicationName parameter")
		return nil, err
	}
	path := fmt.Sprintf("applications/%s/revisions", applicationName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// A single application revision containing the deployed image and configuration
type AppRevision struct {
	// Container image for this revision (mandatory)
	Image string `json:"image" api:"required"`
	// Unique revision identifier
	ID string `json:"id"`
	// When this revision was created
	CreatedAt string `json:"createdAt"`
	// Who created this revision
	CreatedBy string `json:"createdBy"`
	// Environment variables for this revision
	Envs []shared.Env `json:"envs"`
	// Memory allocation in megabytes. Determines CPU allocation (CPU = memory / 2048).
	Memory int64 `json:"memory"`
	// Port the application listens on for this revision (default uses spec-level port
	// or 8080)
	Port int64 `json:"port"`
	// Snapshot ID this revision was forked from (optional, only set when created via
	// fork)
	Snapshot string `json:"snapshot"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Image       respjson.Field
		ID          respjson.Field
		CreatedAt   respjson.Field
		CreatedBy   respjson.Field
		Envs        respjson.Field
		Memory      respjson.Field
		Port        respjson.Field
		Snapshot    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppRevision) RawJSON() string { return r.JSON.raw }
func (r *AppRevision) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AppRevision to a AppRevisionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AppRevisionParam.Overrides()
func (r AppRevision) ToParam() AppRevisionParam {
	return param.Override[AppRevisionParam](json.RawMessage(r.RawJSON()))
}

// A single application revision containing the deployed image and configuration
//
// The property Image is required.
type AppRevisionParam struct {
	// Container image for this revision (mandatory)
	Image string `json:"image" api:"required"`
	// Unique revision identifier
	ID param.Opt[string] `json:"id,omitzero"`
	// When this revision was created
	CreatedAt param.Opt[string] `json:"createdAt,omitzero"`
	// Who created this revision
	CreatedBy param.Opt[string] `json:"createdBy,omitzero"`
	// Memory allocation in megabytes. Determines CPU allocation (CPU = memory / 2048).
	Memory param.Opt[int64] `json:"memory,omitzero"`
	// Port the application listens on for this revision (default uses spec-level port
	// or 8080)
	Port param.Opt[int64] `json:"port,omitzero"`
	// Snapshot ID this revision was forked from (optional, only set when created via
	// fork)
	Snapshot param.Opt[string] `json:"snapshot,omitzero"`
	// Environment variables for this revision
	Envs []shared.EnvParam `json:"envs,omitzero"`
	paramObj
}

func (r AppRevisionParam) MarshalJSON() (data []byte, err error) {
	type shadow AppRevisionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppRevisionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Routing configuration controlling which revision is active and canary traffic
// splitting
type AppRevisionConfiguration struct {
	// Active revision id
	Active string `json:"active"`
	// Canary revision id
	Canary string `json:"canary"`
	// Canary revision percent (0-100)
	CanaryPercent int64 `json:"canaryPercent"`
	// Sticky session TTL in seconds (0 = disabled)
	StickySessionTtl int64 `json:"stickySessionTtl"`
	// Traffic percentage for deployment
	Traffic int64 `json:"traffic"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Active           respjson.Field
		Canary           respjson.Field
		CanaryPercent    respjson.Field
		StickySessionTtl respjson.Field
		Traffic          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppRevisionConfiguration) RawJSON() string { return r.JSON.raw }
func (r *AppRevisionConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AppRevisionConfiguration to a
// AppRevisionConfigurationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AppRevisionConfigurationParam.Overrides()
func (r AppRevisionConfiguration) ToParam() AppRevisionConfigurationParam {
	return param.Override[AppRevisionConfigurationParam](json.RawMessage(r.RawJSON()))
}

// Routing configuration controlling which revision is active and canary traffic
// splitting
type AppRevisionConfigurationParam struct {
	// Active revision id
	Active param.Opt[string] `json:"active,omitzero"`
	// Canary revision id
	Canary param.Opt[string] `json:"canary,omitzero"`
	// Canary revision percent (0-100)
	CanaryPercent param.Opt[int64] `json:"canaryPercent,omitzero"`
	// Sticky session TTL in seconds (0 = disabled)
	StickySessionTtl param.Opt[int64] `json:"stickySessionTtl,omitzero"`
	// Traffic percentage for deployment
	Traffic param.Opt[int64] `json:"traffic,omitzero"`
	paramObj
}

func (r AppRevisionConfigurationParam) MarshalJSON() (data []byte, err error) {
	type shadow AppRevisionConfigurationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppRevisionConfigurationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A single URL entry for the application. If the domain is a wildcard custom
// domain (e.g. \*.sandbox.vybe.build), use subdomain to pick a specific subdomain.
// If the domain is a direct custom domain (e.g. app.vybe.build), subdomain is not
// needed.
type AppURL struct {
	// Custom domain (must be a verified custom domain in the workspace). Can be a
	// wildcard domain (e.g. sandbox.vybe.build registered as \*.sandbox.vybe.build) or
	// a direct domain (e.g. app.vybe.build).
	Domain string `json:"domain" api:"required"`
	// Subdomain to use with a wildcard custom domain (optional)
	Subdomain string `json:"subdomain"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Domain      respjson.Field
		Subdomain   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AppURL) RawJSON() string { return r.JSON.raw }
func (r *AppURL) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AppURL to a AppURLParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AppURLParam.Overrides()
func (r AppURL) ToParam() AppURLParam {
	return param.Override[AppURLParam](json.RawMessage(r.RawJSON()))
}

// A single URL entry for the application. If the domain is a wildcard custom
// domain (e.g. \*.sandbox.vybe.build), use subdomain to pick a specific subdomain.
// If the domain is a direct custom domain (e.g. app.vybe.build), subdomain is not
// needed.
//
// The property Domain is required.
type AppURLParam struct {
	// Custom domain (must be a verified custom domain in the workspace). Can be a
	// wildcard domain (e.g. sandbox.vybe.build registered as \*.sandbox.vybe.build) or
	// a direct domain (e.g. app.vybe.build).
	Domain string `json:"domain" api:"required"`
	// Subdomain to use with a wildcard custom domain (optional)
	Subdomain param.Opt[string] `json:"subdomain,omitzero"`
	paramObj
}

func (r AppURLParam) MarshalJSON() (data []byte, err error) {
	type shadow AppURLParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AppURLParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Long-running application deployment that runs your custom code as a publicly
// accessible endpoint. Applications are always public and use mk3 generation.
type Application struct {
	// Common metadata fields shared by all Blaxel resources including name, labels,
	// timestamps, and ownership information
	Metadata Metadata `json:"metadata" api:"required"`
	// Configuration for an application including revision management, URL routing, and
	// deployment region
	Spec ApplicationSpec `json:"spec" api:"required"`
	// Events happening on a resource deployed on Blaxel
	Events []CoreEvent `json:"events"`
	// Application status computed from events
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metadata    respjson.Field
		Spec        respjson.Field
		Events      respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Application) RawJSON() string { return r.JSON.raw }
func (r *Application) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Application to a ApplicationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ApplicationParam.Overrides()
func (r Application) ToParam() ApplicationParam {
	return param.Override[ApplicationParam](json.RawMessage(r.RawJSON()))
}

// Long-running application deployment that runs your custom code as a publicly
// accessible endpoint. Applications are always public and use mk3 generation.
//
// The properties Metadata, Spec are required.
type ApplicationParam struct {
	// Common metadata fields shared by all Blaxel resources including name, labels,
	// timestamps, and ownership information
	Metadata MetadataParam `json:"metadata,omitzero" api:"required"`
	// Configuration for an application including revision management, URL routing, and
	// deployment region
	Spec ApplicationSpecParam `json:"spec,omitzero" api:"required"`
	paramObj
}

func (r ApplicationParam) MarshalJSON() (data []byte, err error) {
	type shadow ApplicationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ApplicationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for an application including revision management, URL routing, and
// deployment region
type ApplicationSpec struct {
	// When false, the application is disabled and will not serve requests
	Enabled bool `json:"enabled"`
	// Port the application listens on (default 8080)
	Port int64 `json:"port"`
	// Region where the application is deployed (e.g. us-pdx-1, eu-lon-1)
	Region string `json:"region"`
	// Routing configuration controlling which revision is active and canary traffic
	// splitting
	Revision  AppRevisionConfiguration `json:"revision"`
	Revisions []AppRevision            `json:"revisions"`
	// URL configuration for the application. Each entry defines a custom URL through
	// which the application is accessible. The domain must be a verified custom domain
	// in the workspace.
	URLs []AppURL `json:"urls"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Port        respjson.Field
		Region      respjson.Field
		Revision    respjson.Field
		Revisions   respjson.Field
		URLs        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ApplicationSpec) RawJSON() string { return r.JSON.raw }
func (r *ApplicationSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ApplicationSpec to a ApplicationSpecParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ApplicationSpecParam.Overrides()
func (r ApplicationSpec) ToParam() ApplicationSpecParam {
	return param.Override[ApplicationSpecParam](json.RawMessage(r.RawJSON()))
}

// Configuration for an application including revision management, URL routing, and
// deployment region
type ApplicationSpecParam struct {
	// When false, the application is disabled and will not serve requests
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Port the application listens on (default 8080)
	Port param.Opt[int64] `json:"port,omitzero"`
	// Region where the application is deployed (e.g. us-pdx-1, eu-lon-1)
	Region param.Opt[string] `json:"region,omitzero"`
	// Routing configuration controlling which revision is active and canary traffic
	// splitting
	Revision  AppRevisionConfigurationParam `json:"revision,omitzero"`
	Revisions []AppRevisionParam            `json:"revisions,omitzero"`
	// URL configuration for the application. Each entry defines a custom URL through
	// which the application is accessible. The domain must be a verified custom domain
	// in the workspace.
	URLs []AppURLParam `json:"urls,omitzero"`
	paramObj
}

func (r ApplicationSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow ApplicationSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ApplicationSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ApplicationNewParams struct {
	// Long-running application deployment that runs your custom code as a publicly
	// accessible endpoint. Applications are always public and use mk3 generation.
	Application ApplicationParam
	paramObj
}

func (r ApplicationNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Application)
}
func (r *ApplicationNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ApplicationUpdateParams struct {
	// Long-running application deployment that runs your custom code as a publicly
	// accessible endpoint. Applications are always public and use mk3 generation.
	Application ApplicationParam
	paramObj
}

func (r ApplicationUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Application)
}
func (r *ApplicationUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ApplicationListParams struct {
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
	Anchor ApplicationListParamsAnchor `query:"anchor,omitzero" json:"-"`
	// Sort spec, formatted as `<key>:<direction>`. Allowed values are `createdAt:desc`
	// (default), `createdAt:asc`, `name:asc`, `name:desc`. The cursor fingerprint is
	// bound to the sort, so a cursor opened with one value cannot be reused with
	// another. Only honoured starting on Blaxel-Version 2026-04-28.
	//
	// Any of "createdAt:desc", "createdAt:asc", "name:asc", "name:desc".
	Sort ApplicationListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ApplicationListParams]'s query parameters as `url.Values`.
func (r ApplicationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Start from a known pagination boundary. `end` is only supported for `createdAt`
// listings (asc or desc) and returns the tail page directly without walking every
// cursor from the first page.
type ApplicationListParamsAnchor string

const (
	ApplicationListParamsAnchorEnd ApplicationListParamsAnchor = "end"
)

// Sort spec, formatted as `<key>:<direction>`. Allowed values are `createdAt:desc`
// (default), `createdAt:asc`, `name:asc`, `name:desc`. The cursor fingerprint is
// bound to the sort, so a cursor opened with one value cannot be reused with
// another. Only honoured starting on Blaxel-Version 2026-04-28.
type ApplicationListParamsSort string

const (
	ApplicationListParamsSortCreatedAtDesc ApplicationListParamsSort = "createdAt:desc"
	ApplicationListParamsSortCreatedAtAsc  ApplicationListParamsSort = "createdAt:asc"
	ApplicationListParamsSortNameAsc       ApplicationListParamsSort = "name:asc"
	ApplicationListParamsSortNameDesc      ApplicationListParamsSort = "name:desc"
)
