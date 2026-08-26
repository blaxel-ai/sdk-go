// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blaxel

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/blaxel-ai/sdk-go/internal/apiquery"
	"github.com/blaxel-ai/sdk-go/internal/requestconfig"
	"github.com/blaxel-ai/sdk-go/option"
	"github.com/blaxel-ai/sdk-go/packages/pagination"
	"github.com/blaxel-ai/sdk-go/packages/param"
)

// ScheduleExecutionService contains methods and other services that help with
// interacting with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScheduleExecutionService] method instead.
type ScheduleExecutionService struct {
	Options []option.RequestOption
}

// NewScheduleExecutionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewScheduleExecutionService(opts ...option.RequestOption) (r ScheduleExecutionService) {
	r = ScheduleExecutionService{}
	r.Options = opts
	return
}

// Returns schedule execution submissions across the workspace, newest first by
// default. Status describes submission acceptance, not process completion. since
// and until are inclusive RFC 3339 bounds on creation time.
func (r *ScheduleExecutionService) List(ctx context.Context, query ScheduleExecutionListParams, opts ...option.RequestOption) (res *pagination.CursorPage[SandboxScheduleExecution], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "schedule-executions"
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

// Returns schedule execution submissions across the workspace, newest first by
// default. Status describes submission acceptance, not process completion. since
// and until are inclusive RFC 3339 bounds on creation time.
func (r *ScheduleExecutionService) ListAutoPaging(ctx context.Context, query ScheduleExecutionListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[SandboxScheduleExecution] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

type ScheduleExecutionListParams struct {
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
	// Filter by sandbox name.
	Sandbox param.Opt[string] `query:"sandbox,omitzero" json:"-"`
	// Filter by schedule id.
	Schedule param.Opt[string] `query:"schedule,omitzero" json:"-"`
	// Inclusive beginning of the createdAt window, as RFC 3339.
	Since param.Opt[time.Time] `query:"since,omitzero" format:"date-time" json:"-"`
	// Inclusive end of the createdAt window, as RFC 3339.
	Until param.Opt[time.Time] `query:"until,omitzero" format:"date-time" json:"-"`
	// Sort by creation time. Defaults to newest first.
	//
	// Any of "createdAt:desc", "createdAt:asc".
	Sort ScheduleExecutionListParamsSort `query:"sort,omitzero" json:"-"`
	// Filter by process submission acceptance status. Historical rows without status
	// do not match this filter.
	//
	// Any of "succeeded", "failed".
	Status ScheduleExecutionListParamsStatus `query:"status,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ScheduleExecutionListParams]'s query parameters as
// `url.Values`.
func (r ScheduleExecutionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort by creation time. Defaults to newest first.
type ScheduleExecutionListParamsSort string

const (
	ScheduleExecutionListParamsSortCreatedAtDesc ScheduleExecutionListParamsSort = "createdAt:desc"
	ScheduleExecutionListParamsSortCreatedAtAsc  ScheduleExecutionListParamsSort = "createdAt:asc"
)

// Filter by process submission acceptance status. Historical rows without status
// do not match this filter.
type ScheduleExecutionListParamsStatus string

const (
	ScheduleExecutionListParamsStatusSucceeded ScheduleExecutionListParamsStatus = "succeeded"
	ScheduleExecutionListParamsStatusFailed    ScheduleExecutionListParamsStatus = "failed"
)
