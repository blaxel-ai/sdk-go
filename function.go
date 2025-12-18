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

// FunctionService contains methods and other services that help with interacting
// with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFunctionService] method instead.
type FunctionService struct {
	Options []option.RequestOption
}

// NewFunctionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewFunctionService(opts ...option.RequestOption) (r FunctionService) {
	r = FunctionService{}
	r.Options = opts
	return
}

// Create function
func (r *FunctionService) New(ctx context.Context, body FunctionNewParams, opts ...option.RequestOption) (res *Function, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "functions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get function by name
func (r *FunctionService) Get(ctx context.Context, functionName string, query FunctionGetParams, opts ...option.RequestOption) (res *Function, err error) {
	opts = slices.Concat(r.Options, opts)
	if functionName == "" {
		err = errors.New("missing required functionName parameter")
		return
	}
	path := fmt.Sprintf("functions/%s", functionName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Update function by name
func (r *FunctionService) Update(ctx context.Context, functionName string, body FunctionUpdateParams, opts ...option.RequestOption) (res *Function, err error) {
	opts = slices.Concat(r.Options, opts)
	if functionName == "" {
		err = errors.New("missing required functionName parameter")
		return
	}
	path := fmt.Sprintf("functions/%s", functionName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List all functions
func (r *FunctionService) List(ctx context.Context, opts ...option.RequestOption) (res *[]Function, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "functions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete function by name
func (r *FunctionService) Delete(ctx context.Context, functionName string, opts ...option.RequestOption) (res *Function, err error) {
	opts = slices.Concat(r.Options, opts)
	if functionName == "" {
		err = errors.New("missing required functionName parameter")
		return
	}
	path := fmt.Sprintf("functions/%s", functionName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Returns revisions for a function by name.
func (r *FunctionService) ListRevisions(ctx context.Context, functionName string, opts ...option.RequestOption) (res *[]RevisionMetadata, err error) {
	opts = slices.Concat(r.Options, opts)
	if functionName == "" {
		err = errors.New("missing required functionName parameter")
		return
	}
	path := fmt.Sprintf("functions/%s/revisions", functionName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Function
type Function struct {
	// Metadata
	Metadata Metadata `json:"metadata,required"`
	// Function specification for API
	Spec FunctionSpec `json:"spec,required"`
	// Core events
	Events []CoreEvent `json:"events"`
	// Function status
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
func (r Function) RawJSON() string { return r.JSON.raw }
func (r *Function) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Function to a FunctionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// FunctionParam.Overrides()
func (r Function) ToParam() FunctionParam {
	return param.Override[FunctionParam](json.RawMessage(r.RawJSON()))
}

// Function
//
// The properties Metadata, Spec are required.
type FunctionParam struct {
	// Metadata
	Metadata MetadataParam `json:"metadata,omitzero,required"`
	// Function specification for API
	Spec FunctionSpecParam `json:"spec,omitzero,required"`
	paramObj
}

func (r FunctionParam) MarshalJSON() (data []byte, err error) {
	type shadow FunctionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunctionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Runtime configuration for Function
type FunctionRuntime struct {
	// The env variables to set in the function. Should be a list of Kubernetes EnvVar
	// types
	Envs []map[string]any `json:"envs"`
	// The generation of the function
	//
	// Any of "mk2", "mk3".
	Generation FunctionRuntimeGeneration `json:"generation"`
	// The Docker image for the function
	Image string `json:"image"`
	// The maximum number of replicas for the function.
	MaxScale int64 `json:"maxScale"`
	// The memory for the function in MB
	Memory int64 `json:"memory"`
	// The minimum number of replicas for the function. Can be 0 or 1 (in which case
	// the function is always running in at least one location).
	MinScale int64 `json:"minScale"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Envs        respjson.Field
		Generation  respjson.Field
		Image       respjson.Field
		MaxScale    respjson.Field
		Memory      respjson.Field
		MinScale    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionRuntime) RawJSON() string { return r.JSON.raw }
func (r *FunctionRuntime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this FunctionRuntime to a FunctionRuntimeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// FunctionRuntimeParam.Overrides()
func (r FunctionRuntime) ToParam() FunctionRuntimeParam {
	return param.Override[FunctionRuntimeParam](json.RawMessage(r.RawJSON()))
}

// The generation of the function
type FunctionRuntimeGeneration string

const (
	FunctionRuntimeGenerationMk2 FunctionRuntimeGeneration = "mk2"
	FunctionRuntimeGenerationMk3 FunctionRuntimeGeneration = "mk3"
)

// Runtime configuration for Function
type FunctionRuntimeParam struct {
	// The Docker image for the function
	Image param.Opt[string] `json:"image,omitzero"`
	// The maximum number of replicas for the function.
	MaxScale param.Opt[int64] `json:"maxScale,omitzero"`
	// The memory for the function in MB
	Memory param.Opt[int64] `json:"memory,omitzero"`
	// The minimum number of replicas for the function. Can be 0 or 1 (in which case
	// the function is always running in at least one location).
	MinScale param.Opt[int64] `json:"minScale,omitzero"`
	// The env variables to set in the function. Should be a list of Kubernetes EnvVar
	// types
	Envs []map[string]any `json:"envs,omitzero"`
	// The generation of the function
	//
	// Any of "mk2", "mk3".
	Generation FunctionRuntimeGeneration `json:"generation,omitzero"`
	paramObj
}

func (r FunctionRuntimeParam) MarshalJSON() (data []byte, err error) {
	type shadow FunctionRuntimeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunctionRuntimeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Function specification for API
type FunctionSpec struct {
	// Enable or disable the resource
	Enabled                bool     `json:"enabled"`
	IntegrationConnections []string `json:"integrationConnections"`
	Policies               []string `json:"policies"`
	// Revision configuration
	Revision RevisionConfiguration `json:"revision"`
	// Runtime configuration for Function
	Runtime FunctionRuntime `json:"runtime"`
	// Transport compatibility for the MCP, can be "websocket" or "http-stream"
	//
	// Any of "websocket", "http-stream".
	Transport FunctionSpecTransport `json:"transport"`
	// Triggers to use your agent
	Triggers []Trigger `json:"triggers"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled                respjson.Field
		IntegrationConnections respjson.Field
		Policies               respjson.Field
		Revision               respjson.Field
		Runtime                respjson.Field
		Transport              respjson.Field
		Triggers               respjson.Field
		ExtraFields            map[string]respjson.Field
		raw                    string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionSpec) RawJSON() string { return r.JSON.raw }
func (r *FunctionSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this FunctionSpec to a FunctionSpecParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// FunctionSpecParam.Overrides()
func (r FunctionSpec) ToParam() FunctionSpecParam {
	return param.Override[FunctionSpecParam](json.RawMessage(r.RawJSON()))
}

// Transport compatibility for the MCP, can be "websocket" or "http-stream"
type FunctionSpecTransport string

const (
	FunctionSpecTransportWebsocket  FunctionSpecTransport = "websocket"
	FunctionSpecTransportHTTPStream FunctionSpecTransport = "http-stream"
)

// Function specification for API
type FunctionSpecParam struct {
	// Enable or disable the resource
	Enabled                param.Opt[bool] `json:"enabled,omitzero"`
	IntegrationConnections []string        `json:"integrationConnections,omitzero"`
	Policies               []string        `json:"policies,omitzero"`
	// Revision configuration
	Revision RevisionConfigurationParam `json:"revision,omitzero"`
	// Runtime configuration for Function
	Runtime FunctionRuntimeParam `json:"runtime,omitzero"`
	// Transport compatibility for the MCP, can be "websocket" or "http-stream"
	//
	// Any of "websocket", "http-stream".
	Transport FunctionSpecTransport `json:"transport,omitzero"`
	// Triggers to use your agent
	Triggers []TriggerParam `json:"triggers,omitzero"`
	paramObj
}

func (r FunctionSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow FunctionSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunctionSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FunctionNewParams struct {
	// Function
	Function FunctionParam
	paramObj
}

func (r FunctionNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Function)
}
func (r *FunctionNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Function)
}

type FunctionGetParams struct {
	// Show secret values (admin only)
	ShowSecrets param.Opt[bool] `query:"show_secrets,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [FunctionGetParams]'s query parameters as `url.Values`.
func (r FunctionGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type FunctionUpdateParams struct {
	// Function
	Function FunctionParam
	paramObj
}

func (r FunctionUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Function)
}
func (r *FunctionUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Function)
}
