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

// Creates a new MCP server function deployment. The function will expose tools via
// the Model Context Protocol that can be used by AI agents. Supports streamable
// HTTP transport.
func (r *FunctionService) New(ctx context.Context, body FunctionNewParams, opts ...option.RequestOption) (res *Function, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "functions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns detailed information about an MCP server function including its
// deployment status, available tools, transport configuration, and endpoint URL.
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

// Updates an MCP server function's configuration and triggers a new deployment.
// Changes to runtime settings, integrations, or transport protocol will be applied
// on the next deployment.
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

// Returns all MCP server functions deployed in the workspace. Each function
// includes its deployment status, transport protocol (websocket or http-stream),
// and endpoint URL.
func (r *FunctionService) List(ctx context.Context, opts ...option.RequestOption) (res *[]Function, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "functions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Permanently deletes an MCP server function and all its deployment history. Any
// agents using this function's tools will no longer be able to invoke them.
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

// MCP server deployment that exposes tools for AI agents via the Model Context
// Protocol (MCP). Deployed as a serverless auto-scaling endpoint using streamable
// HTTP transport.
type Function struct {
	// Common metadata fields shared by all Blaxel resources including name, labels,
	// timestamps, and ownership information
	Metadata Metadata `json:"metadata,required"`
	// Configuration for an MCP server function including runtime settings, transport
	// protocol, and connected integrations
	Spec FunctionSpec `json:"spec,required"`
	// Events happening on a resource deployed on Blaxel
	Events []CoreEvent `json:"events"`
	// Deployment status of a resource deployed on Blaxel
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

// MCP server deployment that exposes tools for AI agents via the Model Context
// Protocol (MCP). Deployed as a serverless auto-scaling endpoint using streamable
// HTTP transport.
//
// The properties Metadata, Spec are required.
type FunctionParam struct {
	// Common metadata fields shared by all Blaxel resources including name, labels,
	// timestamps, and ownership information
	Metadata MetadataParam `json:"metadata,omitzero,required"`
	// Configuration for an MCP server function including runtime settings, transport
	// protocol, and connected integrations
	Spec FunctionSpecParam `json:"spec,omitzero,required"`
	// Events happening on a resource deployed on Blaxel
	Events []CoreEventParam `json:"events,omitzero"`
	// Deployment status of a resource deployed on Blaxel
	//
	// Any of "DELETING", "TERMINATED", "FAILED", "DEACTIVATED", "DEACTIVATING",
	// "UPLOADING", "BUILDING", "DEPLOYING", "DEPLOYED".
	Status Status `json:"status,omitzero"`
	paramObj
}

func (r FunctionParam) MarshalJSON() (data []byte, err error) {
	type shadow FunctionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunctionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Runtime configuration defining how the MCP server function is deployed and
// scaled
type FunctionRuntime struct {
	// Environment variables injected into the function. Supports Kubernetes EnvVar
	// format with valueFrom references.
	Envs []FunctionRuntimeEnv `json:"envs"`
	// Infrastructure generation: mk2 (containers, 2-10s cold starts, 15+ global
	// regions) or mk3 (microVMs, sub-25ms cold starts)
	//
	// Any of "mk2", "mk3".
	Generation FunctionRuntimeGeneration `json:"generation"`
	// Container image built by Blaxel when deploying with 'bl deploy'. This field is
	// auto-populated during deployment.
	Image string `json:"image"`
	// Maximum number of concurrent function instances for auto-scaling
	MaxScale int64 `json:"maxScale"`
	// Memory allocation in megabytes. Also determines CPU allocation (CPU cores =
	// memory in MB / 2048, e.g., 4096MB = 2 CPUs).
	Memory int64 `json:"memory"`
	// Minimum instances to keep warm. Set to 1+ to eliminate cold starts, 0 for
	// scale-to-zero.
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

// Environment variable with name and value
type FunctionRuntimeEnv struct {
	// Name of the environment variable
	Name string `json:"name"`
	// Whether the value is a secret
	Secret bool `json:"secret"`
	// Value of the environment variable
	Value string `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Secret      respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FunctionRuntimeEnv) RawJSON() string { return r.JSON.raw }
func (r *FunctionRuntimeEnv) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Infrastructure generation: mk2 (containers, 2-10s cold starts, 15+ global
// regions) or mk3 (microVMs, sub-25ms cold starts)
type FunctionRuntimeGeneration string

const (
	FunctionRuntimeGenerationMk2 FunctionRuntimeGeneration = "mk2"
	FunctionRuntimeGenerationMk3 FunctionRuntimeGeneration = "mk3"
)

// Runtime configuration defining how the MCP server function is deployed and
// scaled
type FunctionRuntimeParam struct {
	// Container image built by Blaxel when deploying with 'bl deploy'. This field is
	// auto-populated during deployment.
	Image param.Opt[string] `json:"image,omitzero"`
	// Maximum number of concurrent function instances for auto-scaling
	MaxScale param.Opt[int64] `json:"maxScale,omitzero"`
	// Memory allocation in megabytes. Also determines CPU allocation (CPU cores =
	// memory in MB / 2048, e.g., 4096MB = 2 CPUs).
	Memory param.Opt[int64] `json:"memory,omitzero"`
	// Minimum instances to keep warm. Set to 1+ to eliminate cold starts, 0 for
	// scale-to-zero.
	MinScale param.Opt[int64] `json:"minScale,omitzero"`
	// Environment variables injected into the function. Supports Kubernetes EnvVar
	// format with valueFrom references.
	Envs []FunctionRuntimeEnvParam `json:"envs,omitzero"`
	// Infrastructure generation: mk2 (containers, 2-10s cold starts, 15+ global
	// regions) or mk3 (microVMs, sub-25ms cold starts)
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

// Environment variable with name and value
type FunctionRuntimeEnvParam struct {
	// Name of the environment variable
	Name param.Opt[string] `json:"name,omitzero"`
	// Whether the value is a secret
	Secret param.Opt[bool] `json:"secret,omitzero"`
	// Value of the environment variable
	Value param.Opt[string] `json:"value,omitzero"`
	paramObj
}

func (r FunctionRuntimeEnvParam) MarshalJSON() (data []byte, err error) {
	type shadow FunctionRuntimeEnvParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FunctionRuntimeEnvParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration for an MCP server function including runtime settings, transport
// protocol, and connected integrations
type FunctionSpec struct {
	// When false, the function is disabled and will not serve requests
	Enabled                bool     `json:"enabled"`
	IntegrationConnections []string `json:"integrationConnections"`
	Policies               []string `json:"policies"`
	// Revision configuration
	Revision RevisionConfiguration `json:"revision"`
	// Runtime configuration defining how the MCP server function is deployed and
	// scaled
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

// Configuration for an MCP server function including runtime settings, transport
// protocol, and connected integrations
type FunctionSpecParam struct {
	// When false, the function is disabled and will not serve requests
	Enabled                param.Opt[bool] `json:"enabled,omitzero"`
	IntegrationConnections []string        `json:"integrationConnections,omitzero"`
	Policies               []string        `json:"policies,omitzero"`
	// Revision configuration
	Revision RevisionConfigurationParam `json:"revision,omitzero"`
	// Runtime configuration defining how the MCP server function is deployed and
	// scaled
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
	// MCP server deployment that exposes tools for AI agents via the Model Context
	// Protocol (MCP). Deployed as a serverless auto-scaling endpoint using streamable
	// HTTP transport.
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
	// Show secret values (requires workspace admin role)
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
	// MCP server deployment that exposes tools for AI agents via the Model Context
	// Protocol (MCP). Deployed as a serverless auto-scaling endpoint using streamable
	// HTTP transport.
	Function FunctionParam
	paramObj
}

func (r FunctionUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Function)
}
func (r *FunctionUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Function)
}
