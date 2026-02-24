// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blaxel

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/blaxel-ai/sdk-go/internal/apijson"
	shimjson "github.com/blaxel-ai/sdk-go/internal/encoding/json"
	"github.com/blaxel-ai/sdk-go/internal/requestconfig"
	"github.com/blaxel-ai/sdk-go/option"
	"github.com/blaxel-ai/sdk-go/packages/param"
	"github.com/blaxel-ai/sdk-go/packages/respjson"
)

// SandboxProcessService contains methods and other services that help with
// interacting with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSandboxProcessService] method instead.
type SandboxProcessService struct {
	Options []option.RequestOption
}

// NewSandboxProcessService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSandboxProcessService(opts ...option.RequestOption) (r SandboxProcessService) {
	r = SandboxProcessService{}
	r.Options = opts
	return
}

// Execute a command and return process information. If Accept header is
// text/event-stream, streams logs in SSE format and returns the process response
// as a final event.
func (r *SandboxProcessService) New(ctx context.Context, body SandboxProcessNewParams, opts ...option.RequestOption) (res *ProcessResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "process"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get information about a process by its PID or name
func (r *SandboxProcessService) Get(ctx context.Context, identifier string, opts ...option.RequestOption) (res *ProcessResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("process/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get a list of all running and completed processes
func (r *SandboxProcessService) List(ctx context.Context, opts ...option.RequestOption) (res *[]ProcessResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "process"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Forcefully kill a running process
func (r *SandboxProcessService) Kill(ctx context.Context, identifier string, opts ...option.RequestOption) (res *SandboxProcessKillResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("process/%s/kill", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Get the stdout and stderr output of a process
func (r *SandboxProcessService) GetLogs(ctx context.Context, identifier string, opts ...option.RequestOption) (res *ProcessLogs, err error) {
	opts = slices.Concat(r.Options, opts)
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("process/%s/logs", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Gracefully stop a running process
func (r *SandboxProcessService) Stop(ctx context.Context, identifier string, opts ...option.RequestOption) (res *SandboxProcessStopResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if identifier == "" {
		err = errors.New("missing required identifier parameter")
		return
	}
	path := fmt.Sprintf("process/%s", identifier)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type ProcessLogs struct {
	Logs   string `json:"logs" api:"required"`
	Stderr string `json:"stderr" api:"required"`
	Stdout string `json:"stdout" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Logs        respjson.Field
		Stderr      respjson.Field
		Stdout      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProcessLogs) RawJSON() string { return r.JSON.raw }
func (r *ProcessLogs) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Command is required.
type ProcessRequestParam struct {
	Command           string            `json:"command" api:"required"`
	MaxRestarts       param.Opt[int64]  `json:"maxRestarts,omitzero"`
	Name              param.Opt[string] `json:"name,omitzero"`
	RestartOnFailure  param.Opt[bool]   `json:"restartOnFailure,omitzero"`
	Timeout           param.Opt[int64]  `json:"timeout,omitzero"`
	WaitForCompletion param.Opt[bool]   `json:"waitForCompletion,omitzero"`
	WorkingDir        param.Opt[string] `json:"workingDir,omitzero"`
	Env               map[string]string `json:"env,omitzero"`
	WaitForPorts      []int64           `json:"waitForPorts,omitzero"`
	paramObj
}

func (r ProcessRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ProcessRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ProcessRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProcessResponse struct {
	Command     string `json:"command" api:"required"`
	CompletedAt string `json:"completedAt" api:"required"`
	ExitCode    int64  `json:"exitCode" api:"required"`
	Logs        string `json:"logs" api:"required"`
	Name        string `json:"name" api:"required"`
	Pid         string `json:"pid" api:"required"`
	StartedAt   string `json:"startedAt" api:"required"`
	// Any of "failed", "killed", "stopped", "running", "completed".
	Status           ProcessResponseStatus `json:"status" api:"required"`
	Stderr           string                `json:"stderr" api:"required"`
	Stdout           string                `json:"stdout" api:"required"`
	WorkingDir       string                `json:"workingDir" api:"required"`
	MaxRestarts      int64                 `json:"maxRestarts"`
	RestartCount     int64                 `json:"restartCount"`
	RestartOnFailure bool                  `json:"restartOnFailure"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Command          respjson.Field
		CompletedAt      respjson.Field
		ExitCode         respjson.Field
		Logs             respjson.Field
		Name             respjson.Field
		Pid              respjson.Field
		StartedAt        respjson.Field
		Status           respjson.Field
		Stderr           respjson.Field
		Stdout           respjson.Field
		WorkingDir       respjson.Field
		MaxRestarts      respjson.Field
		RestartCount     respjson.Field
		RestartOnFailure respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProcessResponse) RawJSON() string { return r.JSON.raw }
func (r *ProcessResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ProcessResponseStatus string

const (
	ProcessResponseStatusFailed    ProcessResponseStatus = "failed"
	ProcessResponseStatusKilled    ProcessResponseStatus = "killed"
	ProcessResponseStatusStopped   ProcessResponseStatus = "stopped"
	ProcessResponseStatusRunning   ProcessResponseStatus = "running"
	ProcessResponseStatusCompleted ProcessResponseStatus = "completed"
)

type SandboxProcessKillResponse struct {
	Message string `json:"message" api:"required"`
	Path    string `json:"path"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Path        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SandboxProcessKillResponse) RawJSON() string { return r.JSON.raw }
func (r *SandboxProcessKillResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SandboxProcessStopResponse struct {
	Message string `json:"message" api:"required"`
	Path    string `json:"path"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Path        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SandboxProcessStopResponse) RawJSON() string { return r.JSON.raw }
func (r *SandboxProcessStopResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SandboxProcessNewParams struct {
	ProcessRequest ProcessRequestParam
	paramObj
}

func (r SandboxProcessNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ProcessRequest)
}
func (r *SandboxProcessNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ProcessRequest)
}
