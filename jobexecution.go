// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blaxel

import (
	"context"
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

// JobExecutionService contains methods and other services that help with
// interacting with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewJobExecutionService] method instead.
type JobExecutionService struct {
	Options []option.RequestOption
}

// NewJobExecutionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewJobExecutionService(opts ...option.RequestOption) (r JobExecutionService) {
	r = JobExecutionService{}
	r.Options = opts
	return
}

// Triggers a new execution of the batch job. Each execution runs multiple tasks in
// parallel according to the job's configured concurrency. Tasks can be
// parameterized via the request body.
func (r *JobExecutionService) New(ctx context.Context, jobID string, body JobExecutionNewParams, opts ...option.RequestOption) (res *JobExecutionNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if jobID == "" {
		err = errors.New("missing required jobId parameter")
		return nil, err
	}
	path := fmt.Sprintf("jobs/%s/executions", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns detailed information about a specific job execution including status,
// task statistics (success/failure/running counts), and timing information.
func (r *JobExecutionService) Get(ctx context.Context, executionID string, query JobExecutionGetParams, opts ...option.RequestOption) (res *JobExecution, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.JobID == "" {
		err = errors.New("missing required jobId parameter")
		return nil, err
	}
	if executionID == "" {
		err = errors.New("missing required executionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("jobs/%s/executions/%s", query.JobID, executionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns executions for a batch job. Starting with API version 2026-04-28 the
// response is wrapped in `{data, meta}` and supports cursor pagination via the
// `cursor` and `limit` query parameters; older versions keep the legacy
// offset/limit contract and return a bare array.
func (r *JobExecutionService) List(ctx context.Context, jobID string, query JobExecutionListParams, opts ...option.RequestOption) (res *pagination.CursorPage[JobExecution], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if jobID == "" {
		err = errors.New("missing required jobId parameter")
		return nil, err
	}
	path := fmt.Sprintf("jobs/%s/executions", jobID)
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

// Returns executions for a batch job. Starting with API version 2026-04-28 the
// response is wrapped in `{data, meta}` and supports cursor pagination via the
// `cursor` and `limit` query parameters; older versions keep the legacy
// offset/limit contract and return a bare array.
func (r *JobExecutionService) ListAutoPaging(ctx context.Context, jobID string, query JobExecutionListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[JobExecution] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, jobID, query, opts...))
}

// Cancels a running job execution. Tasks already in progress will complete, but no
// new tasks will be started. The execution status changes to 'cancelling' then
// 'cancelled'.
func (r *JobExecutionService) Delete(ctx context.Context, executionID string, body JobExecutionDeleteParams, opts ...option.RequestOption) (res *JobExecution, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.JobID == "" {
		err = errors.New("missing required jobId parameter")
		return nil, err
	}
	if executionID == "" {
		err = errors.New("missing required executionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("jobs/%s/executions/%s", body.JobID, executionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Returns one cursor-paginated page of an execution's tasks. Tasks are derived
// from event history each request; only the in-memory slicing is paginated, the
// events scan still fetches the whole event log behind the scenes. Available
// starting with API version 2026-04-28.
func (r *JobExecutionService) ListTasks(ctx context.Context, executionID string, params JobExecutionListTasksParams, opts ...option.RequestOption) (res *pagination.CursorPage[JobExecutionListTasksResponse], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.JobID == "" {
		err = errors.New("missing required jobId parameter")
		return nil, err
	}
	if executionID == "" {
		err = errors.New("missing required executionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("jobs/%s/executions/%s/tasks", params.JobID, executionID)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// Returns one cursor-paginated page of an execution's tasks. Tasks are derived
// from event history each request; only the in-memory slicing is paginated, the
// events scan still fetches the whole event log behind the scenes. Available
// starting with API version 2026-04-28.
func (r *JobExecutionService) ListTasksAutoPaging(ctx context.Context, executionID string, params JobExecutionListTasksParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[JobExecutionListTasksResponse] {
	return pagination.NewCursorPageAutoPager(r.ListTasks(ctx, executionID, params, opts...))
}

// Response returned when a job execution is successfully created. Contains
// identifiers and the tasks that will be executed.
type JobExecutionNewResponse struct {
	// Unique identifier for this request, used for idempotency and tracking.
	// Auto-generated if not provided in the request.
	ID string `json:"id"`
	// Unique identifier for the created execution. Use this ID to track execution
	// status, retrieve logs, or cancel the execution.
	ExecutionID string `json:"executionId"`
	// Name of the job that this execution belongs to
	JobID string `json:"jobId"`
	// Array of task configurations that will be executed in parallel according to the
	// job's concurrency settings. Each task can have custom parameters.
	Tasks []any `json:"tasks"`
	// Name of the workspace where the job execution was created
	WorkspaceID string `json:"workspaceId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ExecutionID respjson.Field
		JobID       respjson.Field
		Tasks       respjson.Field
		WorkspaceID respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobExecutionNewResponse) RawJSON() string { return r.JSON.raw }
func (r *JobExecutionNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Job execution task
type JobExecutionListTasksResponse struct {
	// Task conditions
	Conditions []JobExecutionListTasksResponseCondition `json:"conditions"`
	// Job execution task metadata
	Metadata JobExecutionListTasksResponseMetadata `json:"metadata"`
	// Job execution task specification
	Spec JobExecutionListTasksResponseSpec `json:"spec"`
	// Job execution task status
	//
	// Any of "unspecified", "pending", "reconciling", "failed", "succeeded",
	// "running", "cancelled".
	Status JobExecutionListTasksResponseStatus `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Conditions  respjson.Field
		Metadata    respjson.Field
		Spec        respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobExecutionListTasksResponse) RawJSON() string { return r.JSON.raw }
func (r *JobExecutionListTasksResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Job execution task condition
type JobExecutionListTasksResponseCondition struct {
	// Execution reason
	ExecutionReason string `json:"executionReason"`
	// Condition message
	Message string `json:"message"`
	// Condition reason
	Reason string `json:"reason"`
	// Condition severity
	Severity string `json:"severity"`
	// Condition state
	State string `json:"state"`
	// Condition type
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExecutionReason respjson.Field
		Message         respjson.Field
		Reason          respjson.Field
		Severity        respjson.Field
		State           respjson.Field
		Type            respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobExecutionListTasksResponseCondition) RawJSON() string { return r.JSON.raw }
func (r *JobExecutionListTasksResponseCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Job execution task metadata
type JobExecutionListTasksResponseMetadata struct {
	// Completion timestamp
	CompletedAt string `json:"completedAt"`
	// Creation timestamp
	CreatedAt string `json:"createdAt"`
	// Task name
	Name string `json:"name"`
	// Scheduled timestamp
	ScheduledAt string `json:"scheduledAt"`
	// Start timestamp
	StartedAt string `json:"startedAt"`
	// Last update timestamp
	UpdatedAt string `json:"updatedAt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletedAt respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		ScheduledAt respjson.Field
		StartedAt   respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobExecutionListTasksResponseMetadata) RawJSON() string { return r.JSON.raw }
func (r *JobExecutionListTasksResponseMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Job execution task specification
type JobExecutionListTasksResponseSpec struct {
	// Maximum number of retries
	MaxRetries int64 `json:"maxRetries"`
	// Task timeout duration
	Timeout string `json:"timeout"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaxRetries  respjson.Field
		Timeout     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobExecutionListTasksResponseSpec) RawJSON() string { return r.JSON.raw }
func (r *JobExecutionListTasksResponseSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Job execution task status
type JobExecutionListTasksResponseStatus string

const (
	JobExecutionListTasksResponseStatusUnspecified JobExecutionListTasksResponseStatus = "unspecified"
	JobExecutionListTasksResponseStatusPending     JobExecutionListTasksResponseStatus = "pending"
	JobExecutionListTasksResponseStatusReconciling JobExecutionListTasksResponseStatus = "reconciling"
	JobExecutionListTasksResponseStatusFailed      JobExecutionListTasksResponseStatus = "failed"
	JobExecutionListTasksResponseStatusSucceeded   JobExecutionListTasksResponseStatus = "succeeded"
	JobExecutionListTasksResponseStatusRunning     JobExecutionListTasksResponseStatus = "running"
	JobExecutionListTasksResponseStatusCancelled   JobExecutionListTasksResponseStatus = "cancelled"
)

type JobExecutionNewParams struct {
	// Request to create a job execution
	CreateJobExecutionRequest CreateJobExecutionRequestParam
	paramObj
}

func (r JobExecutionNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateJobExecutionRequest)
}
func (r *JobExecutionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobExecutionGetParams struct {
	JobID string `path:"jobId" api:"required" json:"-"`
	paramObj
}

type JobExecutionListParams struct {
	// Opaque cursor returned by a previous response's meta.nextCursor. Only valid for
	// the same query (workspace + filters); the server rejects cursors bound to a
	// different query or older than 24h. Omit on the first page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Number of items per page
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page offset (legacy, ignored when Blaxel-Version >= 2026-04-28)
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
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
	Sort JobExecutionListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [JobExecutionListParams]'s query parameters as `url.Values`.
func (r JobExecutionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort spec, formatted as `<key>:<direction>`. Allowed values are `createdAt:desc`
// (default), `createdAt:asc`, `name:asc`, `name:desc`. The cursor fingerprint is
// bound to the sort, so a cursor opened with one value cannot be reused with
// another. Only honoured starting on Blaxel-Version 2026-04-28.
type JobExecutionListParamsSort string

const (
	JobExecutionListParamsSortCreatedAtDesc JobExecutionListParamsSort = "createdAt:desc"
	JobExecutionListParamsSortCreatedAtAsc  JobExecutionListParamsSort = "createdAt:asc"
	JobExecutionListParamsSortNameAsc       JobExecutionListParamsSort = "name:asc"
	JobExecutionListParamsSortNameDesc      JobExecutionListParamsSort = "name:desc"
)

type JobExecutionDeleteParams struct {
	JobID string `path:"jobId" api:"required" json:"-"`
	paramObj
}

type JobExecutionListTasksParams struct {
	JobID string `path:"jobId" api:"required" json:"-"`
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
	// Sort spec, formatted as `<key>:<direction>`. Allowed values are `createdAt:desc`
	// (default), `createdAt:asc`, `name:asc`, `name:desc`. The cursor fingerprint is
	// bound to the sort, so a cursor opened with one value cannot be reused with
	// another. Only honoured starting on Blaxel-Version 2026-04-28.
	//
	// Any of "createdAt:desc", "createdAt:asc", "name:asc", "name:desc".
	Sort JobExecutionListTasksParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [JobExecutionListTasksParams]'s query parameters as
// `url.Values`.
func (r JobExecutionListTasksParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Sort spec, formatted as `<key>:<direction>`. Allowed values are `createdAt:desc`
// (default), `createdAt:asc`, `name:asc`, `name:desc`. The cursor fingerprint is
// bound to the sort, so a cursor opened with one value cannot be reused with
// another. Only honoured starting on Blaxel-Version 2026-04-28.
type JobExecutionListTasksParamsSort string

const (
	JobExecutionListTasksParamsSortCreatedAtDesc JobExecutionListTasksParamsSort = "createdAt:desc"
	JobExecutionListTasksParamsSortCreatedAtAsc  JobExecutionListTasksParamsSort = "createdAt:asc"
	JobExecutionListTasksParamsSortNameAsc       JobExecutionListTasksParamsSort = "name:asc"
	JobExecutionListTasksParamsSortNameDesc      JobExecutionListTasksParamsSort = "name:desc"
)
