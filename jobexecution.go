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

	"github.com/stainless-sdks/blaxel-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/blaxel-go/internal/encoding/json"
	"github.com/stainless-sdks/blaxel-go/internal/requestconfig"
	"github.com/stainless-sdks/blaxel-go/option"
	"github.com/stainless-sdks/blaxel-go/packages/param"
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
