// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blaxel

import (
	"context"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/blaxel-ai/sdk-go/internal/apijson"
	"github.com/blaxel-ai/sdk-go/internal/apiquery"
	"github.com/blaxel-ai/sdk-go/internal/requestconfig"
	"github.com/blaxel-ai/sdk-go/option"
	"github.com/blaxel-ai/sdk-go/packages/pagination"
	"github.com/blaxel-ai/sdk-go/packages/param"
	"github.com/blaxel-ai/sdk-go/packages/respjson"
)

// ScheduleService contains methods and other services that help with interacting
// with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewScheduleService] method instead.
type ScheduleService struct {
	Options []option.RequestOption
}

// NewScheduleService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewScheduleService(opts ...option.RequestOption) (r ScheduleService) {
	r = ScheduleService{}
	r.Options = opts
	return
}

// Returns schedule definitions across the workspace, newest first by default.
func (r *ScheduleService) List(ctx context.Context, query ScheduleListParams, opts ...option.RequestOption) (res *pagination.CursorPage[SandboxScheduleEntry], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "schedules"
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

// Returns schedule definitions across the workspace, newest first by default.
func (r *ScheduleService) ListAutoPaging(ctx context.Context, query ScheduleListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[SandboxScheduleEntry] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Returns active sandbox and scheduling metrics for a UTC minute window. since is
// inclusive and until is exclusive. The default window is the last 24 hours and
// the maximum is 7 days. Execution status describes submission acceptance, not
// process completion. Execution totals begin accumulating when metrics collection
// is enabled and can lag recent firings briefly.
func (r *ScheduleService) Metrics(ctx context.Context, query ScheduleMetricsParams, opts ...option.RequestOption) (res *ScheduleMetricsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "schedules/metrics"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Workspace sandbox scheduling metrics for a UTC minute window. since is inclusive
// and until is exclusive.
type ScheduleMetricsResponse struct {
	// Number of schedule execution submissions in the selected window.
	Executions int64 `json:"executions"`
	// Schedule execution counts grouped by submission acceptance status.
	ExecutionsByStatus ScheduleMetricsResponseExecutionsByStatus `json:"executionsByStatus"`
	// Number of active sandboxes in the workspace.
	Sandboxes int64 `json:"sandboxes"`
	// Number of schedules in the workspace.
	Schedules int64 `json:"schedules"`
	// Inclusive beginning of the metrics window, normalized to a UTC minute.
	Since string `json:"since"`
	// Exclusive end of the metrics window, normalized to a UTC minute.
	Until string `json:"until"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Executions         respjson.Field
		ExecutionsByStatus respjson.Field
		Sandboxes          respjson.Field
		Schedules          respjson.Field
		Since              respjson.Field
		Until              respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScheduleMetricsResponse) RawJSON() string { return r.JSON.raw }
func (r *ScheduleMetricsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Schedule execution counts grouped by submission acceptance status.
type ScheduleMetricsResponseExecutionsByStatus struct {
	// Number of process submissions that were not accepted.
	Failed int64 `json:"failed"`
	// Number of process submissions accepted by the sandbox.
	Succeeded int64 `json:"succeeded"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Failed      respjson.Field
		Succeeded   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ScheduleMetricsResponseExecutionsByStatus) RawJSON() string { return r.JSON.raw }
func (r *ScheduleMetricsResponseExecutionsByStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ScheduleListParams struct {
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
	// Sort by creation time. Defaults to newest first.
	//
	// Any of "createdAt:desc", "createdAt:asc".
	Sort ScheduleListParamsSort `query:"sort,omitzero" json:"-"`
	// Filter schedules by stored timing type. sleep resolves to at before persistence.
	//
	// Any of "cron", "at".
	Type ScheduleListParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ScheduleListParams]'s query parameters as `url.Values`.
func (r ScheduleListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort by creation time. Defaults to newest first.
type ScheduleListParamsSort string

const (
	ScheduleListParamsSortCreatedAtDesc ScheduleListParamsSort = "createdAt:desc"
	ScheduleListParamsSortCreatedAtAsc  ScheduleListParamsSort = "createdAt:asc"
)

// Filter schedules by stored timing type. sleep resolves to at before persistence.
type ScheduleListParamsType string

const (
	ScheduleListParamsTypeCron ScheduleListParamsType = "cron"
	ScheduleListParamsTypeAt   ScheduleListParamsType = "at"
)

type ScheduleMetricsParams struct {
	// Inclusive beginning of the metrics window, as RFC 3339. Normalized down to a UTC
	// minute, must be within the last 7 days, and defaults to 24 hours before until.
	Since param.Opt[time.Time] `query:"since,omitzero" format:"date-time" json:"-"`
	// Exclusive end of the metrics window, as RFC 3339. Normalized down to a UTC
	// minute and defaults to the current minute.
	Until param.Opt[time.Time] `query:"until,omitzero" format:"date-time" json:"-"`
	paramObj
}

// URLQuery serializes [ScheduleMetricsParams]'s query parameters as `url.Values`.
func (r ScheduleMetricsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
