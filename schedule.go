// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blaxel

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/blaxel-ai/sdk-go/internal/apiquery"
	"github.com/blaxel-ai/sdk-go/internal/requestconfig"
	"github.com/blaxel-ai/sdk-go/option"
	"github.com/blaxel-ai/sdk-go/packages/pagination"
	"github.com/blaxel-ai/sdk-go/packages/param"
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
