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

	"github.com/stainless-sdks/blaxel-go/internal/apijson"
	"github.com/stainless-sdks/blaxel-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/blaxel-go/internal/encoding/json"
	"github.com/stainless-sdks/blaxel-go/internal/requestconfig"
	"github.com/stainless-sdks/blaxel-go/option"
	"github.com/stainless-sdks/blaxel-go/packages/param"
	"github.com/stainless-sdks/blaxel-go/packages/respjson"
)

// JobService contains methods and other services that help with interacting with
// the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewJobService] method instead.
type JobService struct {
	Options    []option.RequestOption
	Executions JobExecutionService
}

// NewJobService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewJobService(opts ...option.RequestOption) (r JobService) {
	r = JobService{}
	r.Options = opts
	r.Executions = NewJobExecutionService(opts...)
	return
}

// Creates a job.
func (r *JobService) New(ctx context.Context, body JobNewParams, opts ...option.RequestOption) (res *Job, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "jobs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns a job by name.
func (r *JobService) Get(ctx context.Context, jobID string, query JobGetParams, opts ...option.RequestOption) (res *Job, err error) {
	opts = slices.Concat(r.Options, opts)
	if jobID == "" {
		err = errors.New("missing required jobId parameter")
		return
	}
	path := fmt.Sprintf("jobs/%s", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Update a job by name.
func (r *JobService) Update(ctx context.Context, jobID string, body JobUpdateParams, opts ...option.RequestOption) (res *Job, err error) {
	opts = slices.Concat(r.Options, opts)
	if jobID == "" {
		err = errors.New("missing required jobId parameter")
		return
	}
	path := fmt.Sprintf("jobs/%s", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Returns a list of all jobs in the workspace.
func (r *JobService) List(ctx context.Context, opts ...option.RequestOption) (res *[]Job, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "jobs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a job by name.
func (r *JobService) Delete(ctx context.Context, jobID string, opts ...option.RequestOption) (res *Job, err error) {
	opts = slices.Concat(r.Options, opts)
	if jobID == "" {
		err = errors.New("missing required jobId parameter")
		return
	}
	path := fmt.Sprintf("jobs/%s", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Returns revisions for a job by name.
func (r *JobService) ListRevisions(ctx context.Context, jobID string, opts ...option.RequestOption) (res *[]RevisionMetadata, err error) {
	opts = slices.Concat(r.Options, opts)
	if jobID == "" {
		err = errors.New("missing required jobId parameter")
		return
	}
	path := fmt.Sprintf("jobs/%s/revisions", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Request to create a job execution
type CreateJobExecutionRequestParam struct {
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

func (r CreateJobExecutionRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateJobExecutionRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateJobExecutionRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Job
type Job struct {
	// Metadata
	Metadata Metadata `json:"metadata,required"`
	// Job specification for API
	Spec JobSpec `json:"spec,required"`
	// Core events
	Events []CoreEvent `json:"events"`
	// Job status
	//
	// Any of "DELETING", "TERMINATED", "FAILED", "DEACTIVATED", "DEACTIVATING",
	// "UPLOADING", "BUILDING", "DEPLOYING", "DEPLOYED".
	Status Status `json:"status"`
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
func (r Job) RawJSON() string { return r.JSON.raw }
func (r *Job) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Job to a JobParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// JobParam.Overrides()
func (r Job) ToParam() JobParam {
	return param.Override[JobParam](json.RawMessage(r.RawJSON()))
}

// Job
//
// The properties Metadata, Spec are required.
type JobParam struct {
	// Metadata
	Metadata MetadataParam `json:"metadata,omitzero,required"`
	// Job specification for API
	Spec JobSpecParam `json:"spec,omitzero,required"`
	paramObj
}

func (r JobParam) MarshalJSON() (data []byte, err error) {
	type shadow JobParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JobParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	Tasks []JobExecutionSpecTask `json:"tasks"`
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

// Job execution task
type JobExecutionSpecTask struct {
	// Task conditions
	Conditions []JobExecutionSpecTaskCondition `json:"conditions"`
	// Job execution task metadata
	Metadata JobExecutionSpecTaskMetadata `json:"metadata"`
	// Job execution task specification
	Spec JobExecutionSpecTaskSpec `json:"spec"`
	// Job execution task status
	//
	// Any of "unspecified", "pending", "reconciling", "failed", "succeeded",
	// "running", "cancelled".
	Status      string         `json:"status"`
	ExtraFields map[string]any `json:",extras"`
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
func (r JobExecutionSpecTask) RawJSON() string { return r.JSON.raw }
func (r *JobExecutionSpecTask) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Job execution task condition
type JobExecutionSpecTaskCondition struct {
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
func (r JobExecutionSpecTaskCondition) RawJSON() string { return r.JSON.raw }
func (r *JobExecutionSpecTaskCondition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Job execution task metadata
type JobExecutionSpecTaskMetadata struct {
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
func (r JobExecutionSpecTaskMetadata) RawJSON() string { return r.JSON.raw }
func (r *JobExecutionSpecTaskMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Job execution task specification
type JobExecutionSpecTaskSpec struct {
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
func (r JobExecutionSpecTaskSpec) RawJSON() string { return r.JSON.raw }
func (r *JobExecutionSpecTaskSpec) UnmarshalJSON(data []byte) error {
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
	Status      string         `json:"status"`
	ExtraFields map[string]any `json:",extras"`
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

// Runtime configuration for Job
type JobRuntime struct {
	// The env variables to set in the job. Should be a list of Kubernetes EnvVar types
	Envs []map[string]any `json:"envs"`
	// The generation of the job
	//
	// Any of "mk2", "mk3".
	Generation JobRuntimeGeneration `json:"generation"`
	// The Docker image for the job
	Image string `json:"image"`
	// The maximum number of concurrent task for an execution
	MaxConcurrentTasks int64 `json:"maxConcurrentTasks"`
	// The maximum number of retries for the job
	MaxRetries int64 `json:"maxRetries"`
	// The memory for the job in MB
	Memory int64 `json:"memory"`
	// The exposed ports of the job
	Ports []Port `json:"ports"`
	// The timeout for the job in seconds
	Timeout int64 `json:"timeout"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Envs               respjson.Field
		Generation         respjson.Field
		Image              respjson.Field
		MaxConcurrentTasks respjson.Field
		MaxRetries         respjson.Field
		Memory             respjson.Field
		Ports              respjson.Field
		Timeout            respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobRuntime) RawJSON() string { return r.JSON.raw }
func (r *JobRuntime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this JobRuntime to a JobRuntimeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// JobRuntimeParam.Overrides()
func (r JobRuntime) ToParam() JobRuntimeParam {
	return param.Override[JobRuntimeParam](json.RawMessage(r.RawJSON()))
}

// The generation of the job
type JobRuntimeGeneration string

const (
	JobRuntimeGenerationMk2 JobRuntimeGeneration = "mk2"
	JobRuntimeGenerationMk3 JobRuntimeGeneration = "mk3"
)

// Runtime configuration for Job
type JobRuntimeParam struct {
	// The Docker image for the job
	Image param.Opt[string] `json:"image,omitzero"`
	// The maximum number of concurrent task for an execution
	MaxConcurrentTasks param.Opt[int64] `json:"maxConcurrentTasks,omitzero"`
	// The maximum number of retries for the job
	MaxRetries param.Opt[int64] `json:"maxRetries,omitzero"`
	// The memory for the job in MB
	Memory param.Opt[int64] `json:"memory,omitzero"`
	// The timeout for the job in seconds
	Timeout param.Opt[int64] `json:"timeout,omitzero"`
	// The env variables to set in the job. Should be a list of Kubernetes EnvVar types
	Envs []map[string]any `json:"envs,omitzero"`
	// The generation of the job
	//
	// Any of "mk2", "mk3".
	Generation JobRuntimeGeneration `json:"generation,omitzero"`
	// The exposed ports of the job
	Ports []PortParam `json:"ports,omitzero"`
	paramObj
}

func (r JobRuntimeParam) MarshalJSON() (data []byte, err error) {
	type shadow JobRuntimeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JobRuntimeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Job specification for API
type JobSpec struct {
	// Enable or disable the resource
	Enabled  bool     `json:"enabled"`
	Policies []string `json:"policies"`
	// Region where the job should be created (e.g. us-was-1, eu-lon-1)
	Region string `json:"region"`
	// Revision configuration
	Revision RevisionConfiguration `json:"revision"`
	// Runtime configuration for Job
	Runtime JobRuntime `json:"runtime"`
	// Triggers to use your agent
	Triggers []Trigger `json:"triggers"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Policies    respjson.Field
		Region      respjson.Field
		Revision    respjson.Field
		Runtime     respjson.Field
		Triggers    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobSpec) RawJSON() string { return r.JSON.raw }
func (r *JobSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this JobSpec to a JobSpecParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// JobSpecParam.Overrides()
func (r JobSpec) ToParam() JobSpecParam {
	return param.Override[JobSpecParam](json.RawMessage(r.RawJSON()))
}

// Job specification for API
type JobSpecParam struct {
	// Enable or disable the resource
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Region where the job should be created (e.g. us-was-1, eu-lon-1)
	Region   param.Opt[string] `json:"region,omitzero"`
	Policies []string          `json:"policies,omitzero"`
	// Revision configuration
	Revision RevisionConfigurationParam `json:"revision,omitzero"`
	// Runtime configuration for Job
	Runtime JobRuntimeParam `json:"runtime,omitzero"`
	// Triggers to use your agent
	Triggers []TriggerParam `json:"triggers,omitzero"`
	paramObj
}

func (r JobSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow JobSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JobSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Logical object representing a model
type Model struct {
	// Metadata
	Metadata Metadata `json:"metadata,required"`
	// Model specification for API
	Spec ModelSpec `json:"spec,required"`
	// Core events
	Events []CoreEvent `json:"events"`
	// Model status
	//
	// Any of "DELETING", "TERMINATED", "FAILED", "DEACTIVATED", "DEACTIVATING",
	// "UPLOADING", "BUILDING", "DEPLOYING", "DEPLOYED".
	Status Status `json:"status"`
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
func (r Model) RawJSON() string { return r.JSON.raw }
func (r *Model) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Model to a ModelParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ModelParam.Overrides()
func (r Model) ToParam() ModelParam {
	return param.Override[ModelParam](json.RawMessage(r.RawJSON()))
}

// Logical object representing a model
//
// The properties Metadata, Spec are required.
type ModelParam struct {
	// Metadata
	Metadata MetadataParam `json:"metadata,omitzero,required"`
	// Model specification for API
	Spec ModelSpecParam `json:"spec,omitzero,required"`
	paramObj
}

func (r ModelParam) MarshalJSON() (data []byte, err error) {
	type shadow ModelParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ModelParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Runtime configuration for Model
type ModelRuntime struct {
	// Endpoint Name of the model. In case of hf_private_endpoint, it is the endpoint
	// name. In case of hf_public_endpoint, it is not used.
	EndpointName string `json:"endpointName"`
	// The slug name of the origin model at HuggingFace.
	Model string `json:"model"`
	// The organization of the model
	Organization string `json:"organization"`
	// The type of origin for the deployment
	//
	// Any of "hf_private_endpoint", "hf_public_endpoint", "huggingface",
	// "public_model", "mcp", "openai", "anthropic", "gemini", "mistral", "deepseek",
	// "cohere", "cerebras", "xai", "vertexai", "azure-openai-service",
	// "azure-ai-inference", "azure-marketplace", "groq".
	Type ModelRuntimeType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndpointName respjson.Field
		Model        respjson.Field
		Organization respjson.Field
		Type         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelRuntime) RawJSON() string { return r.JSON.raw }
func (r *ModelRuntime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ModelRuntime to a ModelRuntimeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ModelRuntimeParam.Overrides()
func (r ModelRuntime) ToParam() ModelRuntimeParam {
	return param.Override[ModelRuntimeParam](json.RawMessage(r.RawJSON()))
}

// The type of origin for the deployment
type ModelRuntimeType string

const (
	ModelRuntimeTypeHfPrivateEndpoint  ModelRuntimeType = "hf_private_endpoint"
	ModelRuntimeTypeHfPublicEndpoint   ModelRuntimeType = "hf_public_endpoint"
	ModelRuntimeTypeHuggingface        ModelRuntimeType = "huggingface"
	ModelRuntimeTypePublicModel        ModelRuntimeType = "public_model"
	ModelRuntimeTypeMcp                ModelRuntimeType = "mcp"
	ModelRuntimeTypeOpenAI             ModelRuntimeType = "openai"
	ModelRuntimeTypeAnthropic          ModelRuntimeType = "anthropic"
	ModelRuntimeTypeGemini             ModelRuntimeType = "gemini"
	ModelRuntimeTypeMistral            ModelRuntimeType = "mistral"
	ModelRuntimeTypeDeepseek           ModelRuntimeType = "deepseek"
	ModelRuntimeTypeCohere             ModelRuntimeType = "cohere"
	ModelRuntimeTypeCerebras           ModelRuntimeType = "cerebras"
	ModelRuntimeTypeXai                ModelRuntimeType = "xai"
	ModelRuntimeTypeVertexai           ModelRuntimeType = "vertexai"
	ModelRuntimeTypeAzureOpenAIService ModelRuntimeType = "azure-openai-service"
	ModelRuntimeTypeAzureAIInference   ModelRuntimeType = "azure-ai-inference"
	ModelRuntimeTypeAzureMarketplace   ModelRuntimeType = "azure-marketplace"
	ModelRuntimeTypeGroq               ModelRuntimeType = "groq"
)

// Runtime configuration for Model
type ModelRuntimeParam struct {
	// Endpoint Name of the model. In case of hf_private_endpoint, it is the endpoint
	// name. In case of hf_public_endpoint, it is not used.
	EndpointName param.Opt[string] `json:"endpointName,omitzero"`
	// The slug name of the origin model at HuggingFace.
	Model param.Opt[string] `json:"model,omitzero"`
	// The organization of the model
	Organization param.Opt[string] `json:"organization,omitzero"`
	// The type of origin for the deployment
	//
	// Any of "hf_private_endpoint", "hf_public_endpoint", "huggingface",
	// "public_model", "mcp", "openai", "anthropic", "gemini", "mistral", "deepseek",
	// "cohere", "cerebras", "xai", "vertexai", "azure-openai-service",
	// "azure-ai-inference", "azure-marketplace", "groq".
	Type ModelRuntimeType `json:"type,omitzero"`
	paramObj
}

func (r ModelRuntimeParam) MarshalJSON() (data []byte, err error) {
	type shadow ModelRuntimeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ModelRuntimeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Model specification for API
type ModelSpec struct {
	// Enable or disable the resource
	Enabled bool `json:"enabled"`
	// Types of hardware available for deployments
	Flavors                []Flavor `json:"flavors"`
	IntegrationConnections []string `json:"integrationConnections"`
	Policies               []string `json:"policies"`
	// Runtime configuration for Model
	Runtime ModelRuntime `json:"runtime"`
	// Sandbox mode
	Sandbox bool `json:"sandbox"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled                respjson.Field
		Flavors                respjson.Field
		IntegrationConnections respjson.Field
		Policies               respjson.Field
		Runtime                respjson.Field
		Sandbox                respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ModelSpec) RawJSON() string { return r.JSON.raw }
func (r *ModelSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ModelSpec to a ModelSpecParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ModelSpecParam.Overrides()
func (r ModelSpec) ToParam() ModelSpecParam {
	return param.Override[ModelSpecParam](json.RawMessage(r.RawJSON()))
}

// Model specification for API
type ModelSpecParam struct {
	// Enable or disable the resource
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Sandbox mode
	Sandbox param.Opt[bool] `json:"sandbox,omitzero"`
	// Types of hardware available for deployments
	Flavors                []FlavorParam `json:"flavors,omitzero"`
	IntegrationConnections []string      `json:"integrationConnections,omitzero"`
	Policies               []string      `json:"policies,omitzero"`
	// Runtime configuration for Model
	Runtime ModelRuntimeParam `json:"runtime,omitzero"`
	paramObj
}

func (r ModelSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow ModelSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ModelSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobNewParams struct {
	// Job
	Job JobParam
	paramObj
}

func (r JobNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Job)
}
func (r *JobNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Job)
}

type JobGetParams struct {
	// Show secret values (admin only)
	ShowSecrets param.Opt[bool] `query:"show_secrets,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [JobGetParams]'s query parameters as `url.Values`.
func (r JobGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type JobUpdateParams struct {
	// Job
	Job JobParam
	paramObj
}

func (r JobUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Job)
}
func (r *JobUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Job)
}
