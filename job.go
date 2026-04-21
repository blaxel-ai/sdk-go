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
	"github.com/blaxel-ai/sdk-go/shared"
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

// Creates a new batch job definition for parallel AI task processing. Jobs can be
// triggered via API or scheduled, and support configurable parallelism, timeouts,
// and retry logic.
func (r *JobService) New(ctx context.Context, body JobNewParams, opts ...option.RequestOption) (res *Job, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "jobs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns detailed information about a batch job including its runtime
// configuration, execution history, and deployment status.
func (r *JobService) Get(ctx context.Context, jobID string, query JobGetParams, opts ...option.RequestOption) (res *Job, err error) {
	opts = slices.Concat(r.Options, opts)
	if jobID == "" {
		err = errors.New("missing required jobId parameter")
		return nil, err
	}
	path := fmt.Sprintf("jobs/%s", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Updates a batch job's configuration. Changes affect new executions; running
// executions continue with their original configuration.
func (r *JobService) Update(ctx context.Context, jobID string, body JobUpdateParams, opts ...option.RequestOption) (res *Job, err error) {
	opts = slices.Concat(r.Options, opts)
	if jobID == "" {
		err = errors.New("missing required jobId parameter")
		return nil, err
	}
	path := fmt.Sprintf("jobs/%s", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Returns all batch job definitions in the workspace. Each job can be triggered to
// run multiple parallel tasks with configurable concurrency and retry settings.
func (r *JobService) List(ctx context.Context, opts ...option.RequestOption) (res *[]Job, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "jobs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Permanently deletes a batch job definition and cancels any running executions.
// Historical execution data will be retained for a limited time.
func (r *JobService) Delete(ctx context.Context, jobID string, opts ...option.RequestOption) (res *Job, err error) {
	opts = slices.Concat(r.Options, opts)
	if jobID == "" {
		err = errors.New("missing required jobId parameter")
		return nil, err
	}
	path := fmt.Sprintf("jobs/%s", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Returns revisions for a job by name.
func (r *JobService) ListRevisions(ctx context.Context, jobID string, opts ...option.RequestOption) (res *[]RevisionMetadata, err error) {
	opts = slices.Concat(r.Options, opts)
	if jobID == "" {
		err = errors.New("missing required jobId parameter")
		return nil, err
	}
	path := fmt.Sprintf("jobs/%s/revisions", jobID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Request to create a job execution
type CreateJobExecutionRequestParam struct {
	// Unique message ID
	ID param.Opt[string] `json:"id,omitzero"`
	// Execution ID (optional, will be generated if not provided)
	ExecutionID param.Opt[string] `json:"executionId,omitzero"`
	// Job ID
	JobID param.Opt[string] `json:"jobId,omitzero"`
	// Memory override in megabytes (optional, must be lower than or equal to job's
	// configured memory)
	Memory param.Opt[int64] `json:"memory,omitzero"`
	// Workspace ID
	WorkspaceID param.Opt[string] `json:"workspaceId,omitzero"`
	// Environment variable overrides (optional, will merge with job's environment
	// variables)
	Env map[string]string `json:"env,omitzero"`
	// Array of task parameters for parallel execution
	Tasks []any `json:"tasks,omitzero"`
	paramObj
}

func (r CreateJobExecutionRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CreateJobExecutionRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CreateJobExecutionRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Batch processing job definition for running parallel AI tasks. Jobs can execute
// multiple tasks concurrently with configurable parallelism, retries, and
// timeouts.
type Job struct {
	// Common metadata fields shared by all Blaxel resources including name, labels,
	// timestamps, and ownership information
	Metadata Metadata `json:"metadata" api:"required"`
	// Configuration for a batch job including execution parameters, parallelism
	// settings, and deployment region
	Spec JobSpec `json:"spec" api:"required"`
	// Events happening on a resource deployed on Blaxel
	Events []CoreEvent `json:"events"`
	// Deployment status of a resource deployed on Blaxel
	//
	// Any of "DELETING", "TERMINATED", "FAILED", "DEACTIVATED", "DEACTIVATING",
	// "UPLOADING", "BUILDING", "DEPLOYING", "DEPLOYED", "BUILT".
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

// Batch processing job definition for running parallel AI tasks. Jobs can execute
// multiple tasks concurrently with configurable parallelism, retries, and
// timeouts.
//
// The properties Metadata, Spec are required.
type JobParam struct {
	// Common metadata fields shared by all Blaxel resources including name, labels,
	// timestamps, and ownership information
	Metadata MetadataParam `json:"metadata,omitzero" api:"required"`
	// Configuration for a batch job including execution parameters, parallelism
	// settings, and deployment region
	Spec JobSpecParam `json:"spec,omitzero" api:"required"`
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
	Metadata JobExecutionMetadata `json:"metadata" api:"required"`
	// Job execution specification
	Spec JobExecutionSpec `json:"spec" api:"required"`
	// Job execution statistics
	Stats JobExecutionStats `json:"stats"`
	// Job execution status
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
	// Environment variable overrides (if provided for this execution, values are
	// masked with \*\*\*)
	EnvOverride any `json:"envOverride"`
	// Memory override in megabytes (if provided for this execution)
	MemoryOverride int64 `json:"memoryOverride"`
	// Number of parallel tasks
	Parallelism int64 `json:"parallelism"`
	// List of execution tasks
	Tasks []JobExecutionSpecTask `json:"tasks"`
	// Job timeout in seconds (captured at execution creation time)
	Timeout int64 `json:"timeout"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EnvOverride    respjson.Field
		MemoryOverride respjson.Field
		Parallelism    respjson.Field
		Tasks          respjson.Field
		Timeout        respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
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
	Status string `json:"status"`
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

// Job execution status
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
	Status string `json:"status"`
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

// Runtime configuration defining how batch job tasks are executed with parallelism
// and retry settings
type JobRuntime struct {
	// Percentage of VM RAM allocated for disk storage (tmpfs overlay). Valid range
	// 10-95, default 50. Only applies to mk3.1 (microVM) generation.
	DiskPercent int64 `json:"diskPercent"`
	// Environment variables injected into job tasks. Supports Kubernetes EnvVar format
	// with valueFrom references.
	Envs []shared.Env `json:"envs"`
	// Infrastructure generation: mk2 (containers, 2-10s cold starts) or mk3 (microVMs,
	// sub-25ms cold starts)
	//
	// Any of "mk2", "mk3".
	Generation JobRuntimeGeneration `json:"generation"`
	// Container image built by Blaxel when deploying with 'bl deploy'. This field is
	// auto-populated during deployment.
	Image string `json:"image"`
	// Number of automatic retry attempts for failed tasks before marking as failed
	MaxRetries int64 `json:"maxRetries"`
	// Memory allocation in megabytes. Also determines CPU allocation (CPU cores =
	// memory in MB / 2048, e.g., 4096MB = 2 CPUs).
	Memory int64 `json:"memory"`
	// Set of ports for a resource
	Ports []Port `json:"ports"`
	// Maximum execution time in seconds before a task is terminated
	Timeout int64 `json:"timeout"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DiskPercent respjson.Field
		Envs        respjson.Field
		Generation  respjson.Field
		Image       respjson.Field
		MaxRetries  respjson.Field
		Memory      respjson.Field
		Ports       respjson.Field
		Timeout     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
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

// Infrastructure generation: mk2 (containers, 2-10s cold starts) or mk3 (microVMs,
// sub-25ms cold starts)
type JobRuntimeGeneration string

const (
	JobRuntimeGenerationMk2 JobRuntimeGeneration = "mk2"
	JobRuntimeGenerationMk3 JobRuntimeGeneration = "mk3"
)

// Runtime configuration defining how batch job tasks are executed with parallelism
// and retry settings
type JobRuntimeParam struct {
	// Percentage of VM RAM allocated for disk storage (tmpfs overlay). Valid range
	// 10-95, default 50. Only applies to mk3.1 (microVM) generation.
	DiskPercent param.Opt[int64] `json:"diskPercent,omitzero"`
	// Container image built by Blaxel when deploying with 'bl deploy'. This field is
	// auto-populated during deployment.
	Image param.Opt[string] `json:"image,omitzero"`
	// Number of automatic retry attempts for failed tasks before marking as failed
	MaxRetries param.Opt[int64] `json:"maxRetries,omitzero"`
	// Memory allocation in megabytes. Also determines CPU allocation (CPU cores =
	// memory in MB / 2048, e.g., 4096MB = 2 CPUs).
	Memory param.Opt[int64] `json:"memory,omitzero"`
	// Maximum execution time in seconds before a task is terminated
	Timeout param.Opt[int64] `json:"timeout,omitzero"`
	// Environment variables injected into job tasks. Supports Kubernetes EnvVar format
	// with valueFrom references.
	Envs []shared.EnvParam `json:"envs,omitzero"`
	// Infrastructure generation: mk2 (containers, 2-10s cold starts) or mk3 (microVMs,
	// sub-25ms cold starts)
	//
	// Any of "mk2", "mk3".
	Generation JobRuntimeGeneration `json:"generation,omitzero"`
	// Set of ports for a resource
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

// Configuration for a batch job including execution parameters, parallelism
// settings, and deployment region
type JobSpec struct {
	// When false, the job is disabled and new executions cannot be triggered
	Enabled bool `json:"enabled"`
	// Configuration for running GitHub Actions workflow jobs on Blaxel infrastructure.
	// When repositories are configured, the job acts as a self-hosted GitHub Actions
	// runner. Workflow jobs use runs-on with the Blaxel job name to target a specific
	// runner.
	GitHubRunner JobSpecGitHubRunner `json:"githubRunner"`
	Policies     []string            `json:"policies"`
	// Region where the job should be created (e.g. us-was-1, eu-lon-1)
	Region string `json:"region"`
	// Revision configuration
	Revision RevisionConfiguration `json:"revision"`
	// Runtime configuration defining how batch job tasks are executed with parallelism
	// and retry settings
	Runtime JobRuntime `json:"runtime"`
	// Triggers to use your agent
	Triggers []Trigger       `json:"triggers"`
	Volumes  []JobSpecVolume `json:"volumes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled      respjson.Field
		GitHubRunner respjson.Field
		Policies     respjson.Field
		Region       respjson.Field
		Revision     respjson.Field
		Runtime      respjson.Field
		Triggers     respjson.Field
		Volumes      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
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

// Configuration for running GitHub Actions workflow jobs on Blaxel infrastructure.
// When repositories are configured, the job acts as a self-hosted GitHub Actions
// runner. Workflow jobs use runs-on with the Blaxel job name to target a specific
// runner.
type JobSpecGitHubRunner struct {
	// Repositories in owner/repo format that this runner is associated with. The
	// runner will pick up workflow jobs from any of these repositories. If non-empty,
	// the runner is considered enabled.
	Repositories []string `json:"repositories"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Repositories respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobSpecGitHubRunner) RawJSON() string { return r.JSON.raw }
func (r *JobSpecGitHubRunner) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ephemeral volume for a job. Temporary disk-backed storage that is created when
// the job starts and destroyed when it completes.
type JobSpecVolume struct {
	// Absolute filesystem path where the volume will be mounted inside the container
	MountPath string `json:"mountPath" api:"required"`
	// Identifier for the volume, used to reference it internally
	Name string `json:"name" api:"required"`
	// Storage capacity in megabytes
	SizeMB int64 `json:"sizeMb" api:"required"`
	// Type of volume. Currently only "ephemeral" is supported.
	//
	// Any of "ephemeral".
	Type string `json:"type" api:"required"`
	// If true, the volume is mounted read-only
	ReadOnly bool `json:"readOnly"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MountPath   respjson.Field
		Name        respjson.Field
		SizeMB      respjson.Field
		Type        respjson.Field
		ReadOnly    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r JobSpecVolume) RawJSON() string { return r.JSON.raw }
func (r *JobSpecVolume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for a batch job including execution parameters, parallelism
// settings, and deployment region
type JobSpecParam struct {
	// When false, the job is disabled and new executions cannot be triggered
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Region where the job should be created (e.g. us-was-1, eu-lon-1)
	Region param.Opt[string] `json:"region,omitzero"`
	// Configuration for running GitHub Actions workflow jobs on Blaxel infrastructure.
	// When repositories are configured, the job acts as a self-hosted GitHub Actions
	// runner. Workflow jobs use runs-on with the Blaxel job name to target a specific
	// runner.
	GitHubRunner JobSpecGitHubRunnerParam `json:"githubRunner,omitzero"`
	Policies     []string                 `json:"policies,omitzero"`
	// Revision configuration
	Revision RevisionConfigurationParam `json:"revision,omitzero"`
	// Runtime configuration defining how batch job tasks are executed with parallelism
	// and retry settings
	Runtime JobRuntimeParam `json:"runtime,omitzero"`
	// Triggers to use your agent
	Triggers []TriggerParam       `json:"triggers,omitzero"`
	Volumes  []JobSpecVolumeParam `json:"volumes,omitzero"`
	paramObj
}

func (r JobSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow JobSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JobSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for running GitHub Actions workflow jobs on Blaxel infrastructure.
// When repositories are configured, the job acts as a self-hosted GitHub Actions
// runner. Workflow jobs use runs-on with the Blaxel job name to target a specific
// runner.
type JobSpecGitHubRunnerParam struct {
	// Repositories in owner/repo format that this runner is associated with. The
	// runner will pick up workflow jobs from any of these repositories. If non-empty,
	// the runner is considered enabled.
	Repositories []string `json:"repositories,omitzero"`
	paramObj
}

func (r JobSpecGitHubRunnerParam) MarshalJSON() (data []byte, err error) {
	type shadow JobSpecGitHubRunnerParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JobSpecGitHubRunnerParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Ephemeral volume for a job. Temporary disk-backed storage that is created when
// the job starts and destroyed when it completes.
//
// The properties MountPath, Name, SizeMB, Type are required.
type JobSpecVolumeParam struct {
	// Absolute filesystem path where the volume will be mounted inside the container
	MountPath string `json:"mountPath" api:"required"`
	// Identifier for the volume, used to reference it internally
	Name string `json:"name" api:"required"`
	// Storage capacity in megabytes
	SizeMB int64 `json:"sizeMb" api:"required"`
	// Type of volume. Currently only "ephemeral" is supported.
	//
	// Any of "ephemeral".
	Type string `json:"type,omitzero" api:"required"`
	// If true, the volume is mounted read-only
	ReadOnly param.Opt[bool] `json:"readOnly,omitzero"`
	paramObj
}

func (r JobSpecVolumeParam) MarshalJSON() (data []byte, err error) {
	type shadow JobSpecVolumeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *JobSpecVolumeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[JobSpecVolumeParam](
		"type", "ephemeral",
	)
}

// Gateway endpoint to external LLM provider APIs (OpenAI, Anthropic, etc.) with
// unified access control, credentials management, and usage tracking.
type Model struct {
	// Common metadata fields shared by all Blaxel resources including name, labels,
	// timestamps, and ownership information
	Metadata Metadata `json:"metadata" api:"required"`
	// Configuration for a model gateway endpoint including provider type, credentials,
	// and access policies
	Spec ModelSpec `json:"spec" api:"required"`
	// Events happening on a resource deployed on Blaxel
	Events []CoreEvent `json:"events"`
	// Deployment status of a resource deployed on Blaxel
	//
	// Any of "DELETING", "TERMINATED", "FAILED", "DEACTIVATED", "DEACTIVATING",
	// "UPLOADING", "BUILDING", "DEPLOYING", "DEPLOYED", "BUILT".
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

// Gateway endpoint to external LLM provider APIs (OpenAI, Anthropic, etc.) with
// unified access control, credentials management, and usage tracking.
//
// The properties Metadata, Spec are required.
type ModelParam struct {
	// Common metadata fields shared by all Blaxel resources including name, labels,
	// timestamps, and ownership information
	Metadata MetadataParam `json:"metadata,omitzero" api:"required"`
	// Configuration for a model gateway endpoint including provider type, credentials,
	// and access policies
	Spec ModelSpecParam `json:"spec,omitzero" api:"required"`
	paramObj
}

func (r ModelParam) MarshalJSON() (data []byte, err error) {
	type shadow ModelParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ModelParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration identifying which external LLM provider and model this gateway
// endpoint proxies to
type ModelRuntime struct {
	// Provider-specific endpoint name (e.g., HuggingFace Inference Endpoints name)
	EndpointName string `json:"endpointName"`
	// Infrastructure generation. Empty (default) uses the classic deployment path. mk3
	// deploys through the model-gateway on microVM clusters.
	//
	// Any of "mk3", "mk2".
	Generation ModelRuntimeGeneration `json:"generation"`
	// Model identifier at the provider (e.g., gpt-4.1, claude-sonnet-4-20250514,
	// mistral-large-latest)
	Model string `json:"model"`
	// Organization or account identifier at the provider (required for some providers
	// like OpenAI)
	Organization string `json:"organization"`
	// LLM provider type determining the API protocol and authentication method
	//
	// Any of "hf_private_endpoint", "hf_public_endpoint", "huggingface",
	// "public_model", "mcp", "openai", "anthropic", "gemini", "mistral", "deepseek",
	// "cohere", "cerebras", "xai", "vertexai", "azure-openai-service",
	// "azure-ai-inference", "azure-marketplace", "groq".
	Type ModelRuntimeType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndpointName respjson.Field
		Generation   respjson.Field
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

// Infrastructure generation. Empty (default) uses the classic deployment path. mk3
// deploys through the model-gateway on microVM clusters.
type ModelRuntimeGeneration string

const (
	ModelRuntimeGenerationMk3 ModelRuntimeGeneration = "mk3"
	ModelRuntimeGenerationMk2 ModelRuntimeGeneration = "mk2"
)

// LLM provider type determining the API protocol and authentication method
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

// Configuration identifying which external LLM provider and model this gateway
// endpoint proxies to
type ModelRuntimeParam struct {
	// Provider-specific endpoint name (e.g., HuggingFace Inference Endpoints name)
	EndpointName param.Opt[string] `json:"endpointName,omitzero"`
	// Model identifier at the provider (e.g., gpt-4.1, claude-sonnet-4-20250514,
	// mistral-large-latest)
	Model param.Opt[string] `json:"model,omitzero"`
	// Organization or account identifier at the provider (required for some providers
	// like OpenAI)
	Organization param.Opt[string] `json:"organization,omitzero"`
	// Infrastructure generation. Empty (default) uses the classic deployment path. mk3
	// deploys through the model-gateway on microVM clusters.
	//
	// Any of "mk3", "mk2".
	Generation ModelRuntimeGeneration `json:"generation,omitzero"`
	// LLM provider type determining the API protocol and authentication method
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

// Configuration for a model gateway endpoint including provider type, credentials,
// and access policies
type ModelSpec struct {
	// When false, the model endpoint is disabled and will not accept inference
	// requests
	Enabled bool `json:"enabled"`
	// Types of hardware available for deployments
	Flavors                []shared.Flavor `json:"flavors"`
	IntegrationConnections []string        `json:"integrationConnections"`
	Policies               []string        `json:"policies"`
	// Configuration identifying which external LLM provider and model this gateway
	// endpoint proxies to
	Runtime ModelRuntime `json:"runtime"`
	// When true, uses sandbox/test credentials from the integration connection
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

// Configuration for a model gateway endpoint including provider type, credentials,
// and access policies
type ModelSpecParam struct {
	// When false, the model endpoint is disabled and will not accept inference
	// requests
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// When true, uses sandbox/test credentials from the integration connection
	Sandbox param.Opt[bool] `json:"sandbox,omitzero"`
	// Types of hardware available for deployments
	Flavors                []shared.FlavorParam `json:"flavors,omitzero"`
	IntegrationConnections []string             `json:"integrationConnections,omitzero"`
	Policies               []string             `json:"policies,omitzero"`
	// Configuration identifying which external LLM provider and model this gateway
	// endpoint proxies to
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
	// Batch processing job definition for running parallel AI tasks. Jobs can execute
	// multiple tasks concurrently with configurable parallelism, retries, and
	// timeouts.
	Job JobParam
	paramObj
}

func (r JobNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Job)
}
func (r *JobNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type JobGetParams struct {
	// Show secret values (requires workspace admin role)
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
	// Batch processing job definition for running parallel AI tasks. Jobs can execute
	// multiple tasks concurrently with configurable parallelism, retries, and
	// timeouts.
	Job JobParam
	paramObj
}

func (r JobUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Job)
}
func (r *JobUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
