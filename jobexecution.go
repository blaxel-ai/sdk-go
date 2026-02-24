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
		return
	}
	path := fmt.Sprintf("jobs/%s/executions", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns detailed information about a specific job execution including status,
// task statistics (success/failure/running counts), and timing information.
func (r *JobExecutionService) Get(ctx context.Context, executionID string, query JobExecutionGetParams, opts ...option.RequestOption) (res *JobExecution, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.JobID == "" {
		err = errors.New("missing required jobId parameter")
		return
	}
	if executionID == "" {
		err = errors.New("missing required executionId parameter")
		return
	}
	path := fmt.Sprintf("jobs/%s/executions/%s", query.JobID, executionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns paginated list of executions for a batch job, sorted by creation time.
// Each execution contains status, task counts, and timing information.
func (r *JobExecutionService) List(ctx context.Context, jobID string, query JobExecutionListParams, opts ...option.RequestOption) (res *[]JobExecution, err error) {
	opts = slices.Concat(r.Options, opts)
	if jobID == "" {
		err = errors.New("missing required jobId parameter")
		return
	}
	path := fmt.Sprintf("jobs/%s/executions", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Cancels a running job execution. Tasks already in progress will complete, but no
// new tasks will be started. The execution status changes to 'cancelling' then
// 'cancelled'.
func (r *JobExecutionService) Delete(ctx context.Context, executionID string, body JobExecutionDeleteParams, opts ...option.RequestOption) (res *JobExecution, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.JobID == "" {
		err = errors.New("missing required jobId parameter")
		return
	}
	if executionID == "" {
		err = errors.New("missing required executionId parameter")
		return
	}
	path := fmt.Sprintf("jobs/%s/executions/%s", body.JobID, executionID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
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

type JobExecutionNewParams struct {
	// Request to create a job execution
	CreateJobExecutionRequest CreateJobExecutionRequestParam
	paramObj
}

func (r JobExecutionNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CreateJobExecutionRequest)
}
func (r *JobExecutionNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CreateJobExecutionRequest)
}

type JobExecutionGetParams struct {
	JobID string `path:"jobId" api:"required" json:"-"`
	paramObj
}

type JobExecutionListParams struct {
	// Number of items per page
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page offset
	Offset param.Opt[int64] `query:"offset,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [JobExecutionListParams]'s query parameters as `url.Values`.
func (r JobExecutionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type JobExecutionDeleteParams struct {
	JobID string `path:"jobId" api:"required" json:"-"`
	paramObj
}
