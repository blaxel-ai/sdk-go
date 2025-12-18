// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blaxel

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/blaxel-go/internal/apijson"
	"github.com/stainless-sdks/blaxel-go/internal/apiquery"
	"github.com/stainless-sdks/blaxel-go/internal/requestconfig"
	"github.com/stainless-sdks/blaxel-go/option"
	"github.com/stainless-sdks/blaxel-go/packages/param"
	"github.com/stainless-sdks/blaxel-go/packages/respjson"
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

// Creates a new execution for a job by name.
func (r *JobExecutionService) New(ctx context.Context, jobID string, body JobExecutionNewParams, opts ...option.RequestOption) (res *JobExecution, err error) {
	opts = slices.Concat(r.Options, opts)
	if jobID == "" {
		err = errors.New("missing required jobId parameter")
		return
	}
	path := fmt.Sprintf("jobs/%s/executions", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns an execution for a job by name.
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

// Returns a list of all executions for a job by name.
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

// Stop an execution for a job by name.
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

// Job execution
type JobExecution struct {
	// Job execution metadata
	Metadata JobExecutionMetadata `json:"metadata,required"`
	// Job execution specification
	Spec JobExecutionSpec `json:"spec,required"`
	// Job execution statistics
	Stats JobExecutionStats `json:"stats"`
	// Execution status
	//
	// Any of "queued", "pending", "running", "cancelling", "cancelled", "failed",
	// "succeeded", "timeout".
	Status JobExecutionStatus `json:"status"`
	// List of execution tasks
	Tasks []JobExecutionTask `json:"tasks"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metadata    respjson.Field
		Spec        respjson.Field
		Stats       respjson.Field
		Status      respjson.Field
		Tasks       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobExecution) RawJSON() string { return r.JSON.raw }
func (r *JobExecution) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Job execution metadata
type JobExecutionMetadata struct {
	// Execution ID
	ID string `json:"id"`
	// Cluster ID
	Cluster string `json:"cluster"`
	// Completion timestamp
	CompletedAt string `json:"completedAt"`
	// Creation timestamp
	CreatedAt string `json:"createdAt"`
	// Deletion timestamp
	DeletedAt string `json:"deletedAt"`
	// Expiration timestamp
	ExpiredAt string `json:"expiredAt"`
	// Job name
	Job string `json:"job"`
	// Start timestamp
	StartedAt string `json:"startedAt"`
	// Last update timestamp
	UpdatedAt string `json:"updatedAt"`
	// Workspace ID
	Workspace string `json:"workspace"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Cluster     respjson.Field
		CompletedAt respjson.Field
		CreatedAt   respjson.Field
		DeletedAt   respjson.Field
		ExpiredAt   respjson.Field
		Job         respjson.Field
		StartedAt   respjson.Field
		UpdatedAt   respjson.Field
		Workspace   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobExecutionMetadata) RawJSON() string { return r.JSON.raw }
func (r *JobExecutionMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Job execution specification
type JobExecutionSpec struct {
	// Number of parallel tasks
	Parallelism int64 `json:"parallelism"`
	// List of execution tasks
	Tasks []JobExecutionTask `json:"tasks"`
	// Job timeout in seconds (captured at execution creation time)
	Timeout int64 `json:"timeout"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Parallelism respjson.Field
		Tasks       respjson.Field
		Timeout     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobExecutionSpec) RawJSON() string { return r.JSON.raw }
func (r *JobExecutionSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Job execution statistics
type JobExecutionStats struct {
	// Number of cancelled tasks
	Cancelled int64 `json:"cancelled"`
	// Number of failed tasks
	Failure int64 `json:"failure"`
	// Number of retried tasks
	Retried int64 `json:"retried"`
	// Number of running tasks
	Running int64 `json:"running"`
	// Number of successful tasks
	Success int64 `json:"success"`
	// Total number of tasks
	Total int64 `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Cancelled   respjson.Field
		Failure     respjson.Field
		Retried     respjson.Field
		Running     respjson.Field
		Success     respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobExecutionStats) RawJSON() string { return r.JSON.raw }
func (r *JobExecutionStats) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Execution status
type JobExecutionStatus string

const (
	JobExecutionStatusQueued     JobExecutionStatus = "queued"
	JobExecutionStatusPending    JobExecutionStatus = "pending"
	JobExecutionStatusRunning    JobExecutionStatus = "running"
	JobExecutionStatusCancelling JobExecutionStatus = "cancelling"
	JobExecutionStatusCancelled  JobExecutionStatus = "cancelled"
	JobExecutionStatusFailed     JobExecutionStatus = "failed"
	JobExecutionStatusSucceeded  JobExecutionStatus = "succeeded"
	JobExecutionStatusTimeout    JobExecutionStatus = "timeout"
)

// Job execution task
type JobExecutionTask struct {
	// Task conditions
	Conditions []JobExecutionTaskCondition `json:"conditions"`
	// Job execution task metadata
	Metadata JobExecutionTaskMetadata `json:"metadata"`
	// Job execution task specification
	Spec JobExecutionTaskSpec `json:"spec"`
	// Job execution task status
	//
	// Any of "unspecified", "pending", "reconciling", "failed", "succeeded",
	// "running", "cancelled".
	Status JobExecutionTaskStatus `json:"status"`
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
func (r JobExecutionTask) RawJSON() string { return r.JSON.raw }
func (r *JobExecutionTask) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Job execution task condition
type JobExecutionTaskCondition struct {
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
	Type        string         `json:"type"`
	ExtraFields map[string]any `json:",extras"`
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
func (r JobExecutionTaskCondition) RawJSON() string { return r.JSON.raw }
func (r *JobExecutionTaskCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Job execution task metadata
type JobExecutionTaskMetadata struct {
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
func (r JobExecutionTaskMetadata) RawJSON() string { return r.JSON.raw }
func (r *JobExecutionTaskMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Job execution task specification
type JobExecutionTaskSpec struct {
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
func (r JobExecutionTaskSpec) RawJSON() string { return r.JSON.raw }
func (r *JobExecutionTaskSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Job execution task status
type JobExecutionTaskStatus string

const (
	JobExecutionTaskStatusUnspecified JobExecutionTaskStatus = "unspecified"
	JobExecutionTaskStatusPending     JobExecutionTaskStatus = "pending"
	JobExecutionTaskStatusReconciling JobExecutionTaskStatus = "reconciling"
	JobExecutionTaskStatusFailed      JobExecutionTaskStatus = "failed"
	JobExecutionTaskStatusSucceeded   JobExecutionTaskStatus = "succeeded"
	JobExecutionTaskStatusRunning     JobExecutionTaskStatus = "running"
	JobExecutionTaskStatusCancelled   JobExecutionTaskStatus = "cancelled"
)

type JobExecutionNewParams struct {
	// Unique message ID
	ID param.Opt[string] `json:"id,omitzero"`
	// Execution ID (optional, will be generated if not provided)
	ExecutionID param.Opt[string] `json:"executionId,omitzero"`
	// Job ID
	JobID param.Opt[string] `json:"jobId,omitzero"`
	// Workspace ID
	WorkspaceID param.Opt[string] `json:"workspaceId,omitzero"`
	// Array of task parameters for parallel execution
	Tasks []map[string]any `json:"tasks,omitzero"`
	paramObj
}

func (r JobExecutionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow JobExecutionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JobExecutionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobExecutionGetParams struct {
	JobID string `path:"jobId,required" json:"-"`
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
	JobID string `path:"jobId,required" json:"-"`
	paramObj
}
