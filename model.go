// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blaxel

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/blaxel-go/internal/apijson"
	shimjson "github.com/stainless-sdks/blaxel-go/internal/encoding/json"
	"github.com/stainless-sdks/blaxel-go/internal/requestconfig"
	"github.com/stainless-sdks/blaxel-go/option"
	"github.com/stainless-sdks/blaxel-go/packages/param"
	"github.com/stainless-sdks/blaxel-go/packages/respjson"
)

// ModelService contains methods and other services that help with interacting with
// the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewModelService] method instead.
type ModelService struct {
	Options []option.RequestOption
}

// NewModelService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewModelService(opts ...option.RequestOption) (r ModelService) {
	r = ModelService{}
	r.Options = opts
	return
}

// Creates a model.
func (r *ModelService) New(ctx context.Context, body ModelNewParams, opts ...option.RequestOption) (res *Model, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "models"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns a model by name.
func (r *ModelService) Get(ctx context.Context, modelName string, opts ...option.RequestOption) (res *Model, err error) {
	opts = slices.Concat(r.Options, opts)
	if modelName == "" {
		err = errors.New("missing required modelName parameter")
		return
	}
	path := fmt.Sprintf("models/%s", modelName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update a model by name.
func (r *ModelService) Update(ctx context.Context, modelName string, body ModelUpdateParams, opts ...option.RequestOption) (res *Model, err error) {
	opts = slices.Concat(r.Options, opts)
	if modelName == "" {
		err = errors.New("missing required modelName parameter")
		return
	}
	path := fmt.Sprintf("models/%s", modelName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Returns a list of all models in the workspace.
func (r *ModelService) List(ctx context.Context, opts ...option.RequestOption) (res *[]Model, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "models"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a model by name.
func (r *ModelService) Delete(ctx context.Context, modelName string, opts ...option.RequestOption) (res *Model, err error) {
	opts = slices.Concat(r.Options, opts)
	if modelName == "" {
		err = errors.New("missing required modelName parameter")
		return
	}
	path := fmt.Sprintf("models/%s", modelName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Returns revisions for a model by name.
func (r *ModelService) ListRevisions(ctx context.Context, modelName string, opts ...option.RequestOption) (res *[]RevisionMetadata, err error) {
	opts = slices.Concat(r.Options, opts)
	if modelName == "" {
		err = errors.New("missing required modelName parameter")
		return
	}
	path := fmt.Sprintf("models/%s/revisions", modelName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
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

// Model specification for API
type ModelSpec struct {
	// Enable or disable the resource
	Enabled bool `json:"enabled"`
	// Types of hardware available for deployments
	Flavors                []Flavor `json:"flavors"`
	IntegrationConnections []string `json:"integrationConnections"`
	Policies               []string `json:"policies"`
	// Runtime configuration for Model
	Runtime ModelSpecRuntime `json:"runtime"`
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

// Runtime configuration for Model
type ModelSpecRuntime struct {
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
	Type string `json:"type"`
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
func (r ModelSpecRuntime) RawJSON() string { return r.JSON.raw }
func (r *ModelSpecRuntime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
	Runtime ModelSpecRuntimeParam `json:"runtime,omitzero"`
	paramObj
}

func (r ModelSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow ModelSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ModelSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Runtime configuration for Model
type ModelSpecRuntimeParam struct {
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
	Type string `json:"type,omitzero"`
	paramObj
}

func (r ModelSpecRuntimeParam) MarshalJSON() (data []byte, err error) {
	type shadow ModelSpecRuntimeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ModelSpecRuntimeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[ModelSpecRuntimeParam](
		"type", "hf_private_endpoint", "hf_public_endpoint", "huggingface", "public_model", "mcp", "openai", "anthropic", "gemini", "mistral", "deepseek", "cohere", "cerebras", "xai", "vertexai", "azure-openai-service", "azure-ai-inference", "azure-marketplace", "groq",
	)
}

type ModelNewParams struct {
	// Logical object representing a model
	Model ModelParam
	paramObj
}

func (r ModelNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Model)
}
func (r *ModelNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Model)
}

type ModelUpdateParams struct {
	// Logical object representing a model
	Model ModelParam
	paramObj
}

func (r ModelUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Model)
}
func (r *ModelUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Model)
}
