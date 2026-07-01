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

// SandboxScheduleService contains methods and other services that help with
// interacting with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSandboxScheduleService] method instead.
type SandboxScheduleService struct {
	Options []option.RequestOption
}

// NewSandboxScheduleService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSandboxScheduleService(opts ...option.RequestOption) (r SandboxScheduleService) {
	r = SandboxScheduleService{}
	r.Options = opts
	return
}

// Adds a schedule to a Sandbox.
func (r *SandboxScheduleService) New(ctx context.Context, sandboxName string, body SandboxScheduleNewParams, opts ...option.RequestOption) (res *SandboxScheduleEntry, err error) {
	opts = slices.Concat(r.Options, opts)
	if sandboxName == "" {
		err = errors.New("missing required sandboxName parameter")
		return nil, err
	}
	path := fmt.Sprintf("sandboxes/%s/schedules", sandboxName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a Sandbox Schedule by id.
func (r *SandboxScheduleService) Get(ctx context.Context, scheduleID string, query SandboxScheduleGetParams, opts ...option.RequestOption) (res *SandboxScheduleEntry, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.SandboxName == "" {
		err = errors.New("missing required sandboxName parameter")
		return nil, err
	}
	if scheduleID == "" {
		err = errors.New("missing required scheduleId parameter")
		return nil, err
	}
	path := fmt.Sprintf("sandboxes/%s/schedules/%s", query.SandboxName, scheduleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates a Sandbox Schedule by id.
func (r *SandboxScheduleService) Update(ctx context.Context, scheduleID string, params SandboxScheduleUpdateParams, opts ...option.RequestOption) (res *SandboxScheduleEntry, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.SandboxName == "" {
		err = errors.New("missing required sandboxName parameter")
		return nil, err
	}
	if scheduleID == "" {
		err = errors.New("missing required scheduleId parameter")
		return nil, err
	}
	path := fmt.Sprintf("sandboxes/%s/schedules/%s", params.SandboxName, scheduleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Returns the schedules configured on a Sandbox. Starting with API version
// 2026-04-28 the response is wrapped in `{data, meta}` and supports cursor
// pagination via the `cursor` and `limit` query parameters; older versions return
// a bare array.
func (r *SandboxScheduleService) List(ctx context.Context, sandboxName string, query SandboxScheduleListParams, opts ...option.RequestOption) (res *pagination.CursorPage[SandboxScheduleEntry], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if sandboxName == "" {
		err = errors.New("missing required sandboxName parameter")
		return nil, err
	}
	path := fmt.Sprintf("sandboxes/%s/schedules", sandboxName)
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

// Returns the schedules configured on a Sandbox. Starting with API version
// 2026-04-28 the response is wrapped in `{data, meta}` and supports cursor
// pagination via the `cursor` and `limit` query parameters; older versions return
// a bare array.
func (r *SandboxScheduleService) ListAutoPaging(ctx context.Context, sandboxName string, query SandboxScheduleListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[SandboxScheduleEntry] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, sandboxName, query, opts...))
}

// Deletes a Sandbox Schedule by id.
func (r *SandboxScheduleService) Delete(ctx context.Context, scheduleID string, body SandboxScheduleDeleteParams, opts ...option.RequestOption) (res *SandboxScheduleEntry, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.SandboxName == "" {
		err = errors.New("missing required sandboxName parameter")
		return nil, err
	}
	if scheduleID == "" {
		err = errors.New("missing required scheduleId parameter")
		return nil, err
	}
	path := fmt.Sprintf("sandboxes/%s/schedules/%s", body.SandboxName, scheduleID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Returns the execution history of a Sandbox's schedules (across all schedules of
// the sandbox), newest first. Cursor-paginated via the `cursor` and `limit` query
// parameters. Each item records the HTTP status of submitting the scheduled
// command and the process name for log lookup.
func (r *SandboxScheduleService) ListExecutions(ctx context.Context, sandboxName string, query SandboxScheduleListExecutionsParams, opts ...option.RequestOption) (res *pagination.CursorPage[SandboxScheduleExecution], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if sandboxName == "" {
		err = errors.New("missing required sandboxName parameter")
		return nil, err
	}
	path := fmt.Sprintf("sandboxes/%s/schedule-executions", sandboxName)
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

// Returns the execution history of a Sandbox's schedules (across all schedules of
// the sandbox), newest first. Cursor-paginated via the `cursor` and `limit` query
// parameters. Each item records the HTTP status of submitting the scheduled
// command and the process name for log lookup.
func (r *SandboxScheduleService) ListExecutionsAutoPaging(ctx context.Context, sandboxName string, query SandboxScheduleListExecutionsParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[SandboxScheduleExecution] {
	return pagination.NewCursorPageAutoPager(r.ListExecutions(ctx, sandboxName, query, opts...))
}

// A scheduled task that executes a process inside the sandbox at specified times.
// Stored in the dedicated schedules table (no longer embedded in the sandbox
// spec).
type SandboxScheduleEntry struct {
	// Unique identifier for this schedule within its sandbox. Auto-generated if not
	// provided.
	ID string `json:"id"`
	// Creation timestamp (read-only).
	CreatedAt string `json:"createdAt"`
	// Process execution configuration for a scheduled sandbox task
	Input SandboxScheduleInput `json:"input"`
	// Maximum number of execution records kept for this schedule. Once reached,
	// recording a new execution deletes the oldest. Defaults to 100.
	MaxExecutions int64 `json:"maxExecutions"`
	// Type of schedule timing. 'cron' for recurring (5-field expression), 'at' for a
	// specific RFC 3339 datetime, 'sleep' for a duration from now (resolved to 'at' on
	// creation).
	//
	// Any of "cron", "at", "sleep".
	Type SandboxScheduleEntryType `json:"type"`
	// Timing value. For 'cron': a 5-field cron expression (e.g. '0 8 \* \* 1-5'). For
	// 'at': an RFC 3339 datetime (e.g. '2026-07-01T09:00:00Z'). For 'sleep': a
	// duration (e.g. '2h', '30m', '7d').
	Value string `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CreatedAt     respjson.Field
		Input         respjson.Field
		MaxExecutions respjson.Field
		Type          respjson.Field
		Value         respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SandboxScheduleEntry) RawJSON() string { return r.JSON.raw }
func (r *SandboxScheduleEntry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SandboxScheduleEntry to a SandboxScheduleEntryParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SandboxScheduleEntryParam.Overrides()
func (r SandboxScheduleEntry) ToParam() SandboxScheduleEntryParam {
	return param.Override[SandboxScheduleEntryParam](json.RawMessage(r.RawJSON()))
}

// Type of schedule timing. 'cron' for recurring (5-field expression), 'at' for a
// specific RFC 3339 datetime, 'sleep' for a duration from now (resolved to 'at' on
// creation).
type SandboxScheduleEntryType string

const (
	SandboxScheduleEntryTypeCron  SandboxScheduleEntryType = "cron"
	SandboxScheduleEntryTypeAt    SandboxScheduleEntryType = "at"
	SandboxScheduleEntryTypeSleep SandboxScheduleEntryType = "sleep"
)

// A scheduled task that executes a process inside the sandbox at specified times.
// Stored in the dedicated schedules table (no longer embedded in the sandbox
// spec).
type SandboxScheduleEntryParam struct {
	// Unique identifier for this schedule within its sandbox. Auto-generated if not
	// provided.
	ID param.Opt[string] `json:"id,omitzero"`
	// Maximum number of execution records kept for this schedule. Once reached,
	// recording a new execution deletes the oldest. Defaults to 100.
	MaxExecutions param.Opt[int64] `json:"maxExecutions,omitzero"`
	// Timing value. For 'cron': a 5-field cron expression (e.g. '0 8 \* \* 1-5'). For
	// 'at': an RFC 3339 datetime (e.g. '2026-07-01T09:00:00Z'). For 'sleep': a
	// duration (e.g. '2h', '30m', '7d').
	Value param.Opt[string] `json:"value,omitzero"`
	// Process execution configuration for a scheduled sandbox task
	Input SandboxScheduleInputParam `json:"input,omitzero"`
	// Type of schedule timing. 'cron' for recurring (5-field expression), 'at' for a
	// specific RFC 3339 datetime, 'sleep' for a duration from now (resolved to 'at' on
	// creation).
	//
	// Any of "cron", "at", "sleep".
	Type SandboxScheduleEntryType `json:"type,omitzero"`
	paramObj
}

func (r SandboxScheduleEntryParam) MarshalJSON() (data []byte, err error) {
	type shadow SandboxScheduleEntryParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SandboxScheduleEntryParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// One recorded execution of a sandbox schedule. statusCode is the HTTP status from
// submitting the command to the sandbox (the scheduler does not wait for the
// command to finish). Stored in the dedicated scheduleexecutions table.
type SandboxScheduleExecution struct {
	// Unique id of this execution within the schedule.
	ID string `json:"id"`
	// Creation timestamp (read-only).
	CreatedAt string `json:"createdAt"`
	// RFC 3339 time at which the command was submitted.
	ExecutedAt string `json:"executedAt"`
	// Name of the process started in the sandbox for this execution, used to look up
	// its logs.
	ProcessName string `json:"processName"`
	// Id of the schedule this execution belongs to.
	ScheduleID string `json:"scheduleId"`
	// HTTP status code returned when the scheduled command was submitted to the
	// sandbox (0 if the sandbox could not be reached). 2xx/3xx means the command was
	// accepted.
	StatusCode int64 `json:"statusCode"`
	// Process timeout in seconds for this execution. The UI uses it to scope the log
	// view to [executedAt, executedAt+timeout]. 0 when the schedule set no timeout.
	Timeout int64 `json:"timeout"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		ExecutedAt  respjson.Field
		ProcessName respjson.Field
		ScheduleID  respjson.Field
		StatusCode  respjson.Field
		Timeout     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SandboxScheduleExecution) RawJSON() string { return r.JSON.raw }
func (r *SandboxScheduleExecution) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Process execution configuration for a scheduled sandbox task
type SandboxScheduleInput struct {
	// Shell command to execute inside the sandbox
	Command string `json:"command"`
	// Environment variables to set for the process. May contain secrets, so values are
	// encrypted at rest and masked in API responses unless an admin requests
	// show_secrets=true.
	Env map[string]string `json:"env"`
	// Keep the sandbox alive (disable scale-to-zero) while the process runs. Defaults
	// to true.
	KeepAlive bool `json:"keepAlive"`
	// Optional name for the process (used to retrieve status/logs)
	Name string `json:"name"`
	// Timeout in seconds for the process. Defaults to 600 (10 minutes). Set to 0 for
	// no timeout.
	Timeout int64 `json:"timeout"`
	// Working directory for the command
	WorkingDir string `json:"workingDir"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Command     respjson.Field
		Env         respjson.Field
		KeepAlive   respjson.Field
		Name        respjson.Field
		Timeout     respjson.Field
		WorkingDir  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SandboxScheduleInput) RawJSON() string { return r.JSON.raw }
func (r *SandboxScheduleInput) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SandboxScheduleInput to a SandboxScheduleInputParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SandboxScheduleInputParam.Overrides()
func (r SandboxScheduleInput) ToParam() SandboxScheduleInputParam {
	return param.Override[SandboxScheduleInputParam](json.RawMessage(r.RawJSON()))
}

// Process execution configuration for a scheduled sandbox task
type SandboxScheduleInputParam struct {
	// Shell command to execute inside the sandbox
	Command param.Opt[string] `json:"command,omitzero"`
	// Keep the sandbox alive (disable scale-to-zero) while the process runs. Defaults
	// to true.
	KeepAlive param.Opt[bool] `json:"keepAlive,omitzero"`
	// Optional name for the process (used to retrieve status/logs)
	Name param.Opt[string] `json:"name,omitzero"`
	// Timeout in seconds for the process. Defaults to 600 (10 minutes). Set to 0 for
	// no timeout.
	Timeout param.Opt[int64] `json:"timeout,omitzero"`
	// Working directory for the command
	WorkingDir param.Opt[string] `json:"workingDir,omitzero"`
	// Environment variables to set for the process. May contain secrets, so values are
	// encrypted at rest and masked in API responses unless an admin requests
	// show_secrets=true.
	Env map[string]string `json:"env,omitzero"`
	paramObj
}

func (r SandboxScheduleInputParam) MarshalJSON() (data []byte, err error) {
	type shadow SandboxScheduleInputParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SandboxScheduleInputParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SandboxScheduleNewParams struct {
	// A scheduled task that executes a process inside the sandbox at specified times.
	// Stored in the dedicated schedules table (no longer embedded in the sandbox
	// spec).
	SandboxScheduleEntry SandboxScheduleEntryParam
	paramObj
}

func (r SandboxScheduleNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SandboxScheduleEntry)
}
func (r *SandboxScheduleNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SandboxScheduleGetParams struct {
	SandboxName string `path:"sandboxName" api:"required" json:"-"`
	paramObj
}

type SandboxScheduleUpdateParams struct {
	SandboxName string `path:"sandboxName" api:"required" json:"-"`
	// A scheduled task that executes a process inside the sandbox at specified times.
	// Stored in the dedicated schedules table (no longer embedded in the sandbox
	// spec).
	SandboxScheduleEntry SandboxScheduleEntryParam
	paramObj
}

func (r SandboxScheduleUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SandboxScheduleEntry)
}
func (r *SandboxScheduleUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SandboxScheduleListParams struct {
	// Opaque cursor returned by a previous response's meta.nextCursor. Only valid for
	// the same query (workspace + filters); the server rejects cursors bound to a
	// different query or older than 24h. Omit on the first page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Number of items per page
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Substring search across `metadata.name`, `metadata.displayName` and labels
	// (keys + values). Trimmed and lowercased server-side; queries shorter than 2
	// characters fall back to the unfiltered listing. Bound into the cursor
	// fingerprint so a cursor opened with one query cannot be reused with another.
	// Only honoured starting on Blaxel-Version 2026-04-28.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Sort spec, formatted as `<key>:<direction>`. Allowed values are `createdAt:desc`
	// (default), `createdAt:asc`, `name:asc`, `name:desc`. The cursor fingerprint is
	// bound to the sort, so a cursor opened with one value cannot be reused with
	// another. Only honoured starting on Blaxel-Version 2026-04-28.
	//
	// Any of "createdAt:desc", "createdAt:asc", "name:asc", "name:desc".
	Sort SandboxScheduleListParamsSort `query:"sort,omitzero" json:"-"`
	// Filter schedules by timing type. Only cron and at are stored (sleep resolves to
	// at on creation); any other value is ignored.
	//
	// Any of "cron", "at".
	Type SandboxScheduleListParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SandboxScheduleListParams]'s query parameters as
// `url.Values`.
func (r SandboxScheduleListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort spec, formatted as `<key>:<direction>`. Allowed values are `createdAt:desc`
// (default), `createdAt:asc`, `name:asc`, `name:desc`. The cursor fingerprint is
// bound to the sort, so a cursor opened with one value cannot be reused with
// another. Only honoured starting on Blaxel-Version 2026-04-28.
type SandboxScheduleListParamsSort string

const (
	SandboxScheduleListParamsSortCreatedAtDesc SandboxScheduleListParamsSort = "createdAt:desc"
	SandboxScheduleListParamsSortCreatedAtAsc  SandboxScheduleListParamsSort = "createdAt:asc"
	SandboxScheduleListParamsSortNameAsc       SandboxScheduleListParamsSort = "name:asc"
	SandboxScheduleListParamsSortNameDesc      SandboxScheduleListParamsSort = "name:desc"
)

// Filter schedules by timing type. Only cron and at are stored (sleep resolves to
// at on creation); any other value is ignored.
type SandboxScheduleListParamsType string

const (
	SandboxScheduleListParamsTypeCron SandboxScheduleListParamsType = "cron"
	SandboxScheduleListParamsTypeAt   SandboxScheduleListParamsType = "at"
)

type SandboxScheduleDeleteParams struct {
	SandboxName string `path:"sandboxName" api:"required" json:"-"`
	paramObj
}

type SandboxScheduleListExecutionsParams struct {
	// Opaque cursor returned by a previous response's meta.nextCursor. Only valid for
	// the same query (workspace + filters); the server rejects cursors bound to a
	// different query or older than 24h. Omit on the first page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Number of items per page
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Substring search across `metadata.name`, `metadata.displayName` and labels
	// (keys + values). Trimmed and lowercased server-side; queries shorter than 2
	// characters fall back to the unfiltered listing. Bound into the cursor
	// fingerprint so a cursor opened with one query cannot be reused with another.
	// Only honoured starting on Blaxel-Version 2026-04-28.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Sort spec, formatted as `<key>:<direction>`. Allowed values are `createdAt:desc`
	// (default), `createdAt:asc`, `name:asc`, `name:desc`. The cursor fingerprint is
	// bound to the sort, so a cursor opened with one value cannot be reused with
	// another. Only honoured starting on Blaxel-Version 2026-04-28.
	//
	// Any of "createdAt:desc", "createdAt:asc", "name:asc", "name:desc".
	Sort SandboxScheduleListExecutionsParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SandboxScheduleListExecutionsParams]'s query parameters as
// `url.Values`.
func (r SandboxScheduleListExecutionsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort spec, formatted as `<key>:<direction>`. Allowed values are `createdAt:desc`
// (default), `createdAt:asc`, `name:asc`, `name:desc`. The cursor fingerprint is
// bound to the sort, so a cursor opened with one value cannot be reused with
// another. Only honoured starting on Blaxel-Version 2026-04-28.
type SandboxScheduleListExecutionsParamsSort string

const (
	SandboxScheduleListExecutionsParamsSortCreatedAtDesc SandboxScheduleListExecutionsParamsSort = "createdAt:desc"
	SandboxScheduleListExecutionsParamsSortCreatedAtAsc  SandboxScheduleListExecutionsParamsSort = "createdAt:asc"
	SandboxScheduleListExecutionsParamsSortNameAsc       SandboxScheduleListExecutionsParamsSort = "name:asc"
	SandboxScheduleListExecutionsParamsSortNameDesc      SandboxScheduleListExecutionsParamsSort = "name:desc"
)
