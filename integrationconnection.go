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

// IntegrationConnectionService contains methods and other services that help with
// interacting with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIntegrationConnectionService] method instead.
type IntegrationConnectionService struct {
	Options []option.RequestOption
	Models  IntegrationConnectionModelService
}

// NewIntegrationConnectionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewIntegrationConnectionService(opts ...option.RequestOption) (r IntegrationConnectionService) {
	r = IntegrationConnectionService{}
	r.Options = opts
	r.Models = NewIntegrationConnectionModelService(opts...)
	return
}

// Creates a new integration connection with credentials for an external service.
// The connection can then be used by models, functions, and other resources to
// authenticate with the service.
func (r *IntegrationConnectionService) New(ctx context.Context, body IntegrationConnectionNewParams, opts ...option.RequestOption) (res *IntegrationConnection, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "integrations/connections"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns detailed information about an integration connection including its
// provider type, configuration (secrets are masked), and usage status.
func (r *IntegrationConnectionService) Get(ctx context.Context, connectionName string, opts ...option.RequestOption) (res *IntegrationConnection, err error) {
	opts = slices.Concat(r.Options, opts)
	if connectionName == "" {
		err = errors.New("missing required connectionName parameter")
		return
	}
	path := fmt.Sprintf("integrations/connections/%s", connectionName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an integration connection's configuration or credentials. Changes take
// effect immediately for all resources using this connection.
func (r *IntegrationConnectionService) Update(ctx context.Context, connectionName string, body IntegrationConnectionUpdateParams, opts ...option.RequestOption) (res *IntegrationConnection, err error) {
	opts = slices.Concat(r.Options, opts)
	if connectionName == "" {
		err = errors.New("missing required connectionName parameter")
		return
	}
	path := fmt.Sprintf("integrations/connections/%s", connectionName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Returns all configured integration connections in the workspace. Each connection
// stores credentials and settings for an external service (LLM provider, API,
// database).
func (r *IntegrationConnectionService) List(ctx context.Context, opts ...option.RequestOption) (res *[]IntegrationConnection, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "integrations/connections"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Permanently deletes an integration connection. Any resources using this
// connection will lose access to the external service.
func (r *IntegrationConnectionService) Delete(ctx context.Context, connectionName string, opts ...option.RequestOption) (res *IntegrationConnection, err error) {
	opts = slices.Concat(r.Options, opts)
	if connectionName == "" {
		err = errors.New("missing required connectionName parameter")
		return
	}
	path := fmt.Sprintf("integrations/connections/%s", connectionName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Returns a list of all endpoint configurations for a model.
func (r *IntegrationConnectionService) ListEndpointConfigurations(ctx context.Context, connectionName string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if connectionName == "" {
		err = errors.New("missing required connectionName parameter")
		return
	}
	path := fmt.Sprintf("integrations/connections/%s/endpointConfigurations", connectionName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

// Configured connection to an external service (LLM provider, API, SaaS, database)
// storing credentials and settings for use by workspace resources.
type IntegrationConnection struct {
	// Common metadata fields shared by all Blaxel resources including name, labels,
	// timestamps, and ownership information
	Metadata Metadata `json:"metadata,required"`
	// Specification defining the integration type, configuration parameters, and
	// encrypted credentials
	Spec IntegrationConnectionSpec `json:"spec,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metadata    respjson.Field
		Spec        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntegrationConnection) RawJSON() string { return r.JSON.raw }
func (r *IntegrationConnection) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this IntegrationConnection to a IntegrationConnectionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// IntegrationConnectionParam.Overrides()
func (r IntegrationConnection) ToParam() IntegrationConnectionParam {
	return param.Override[IntegrationConnectionParam](json.RawMessage(r.RawJSON()))
}

// Configured connection to an external service (LLM provider, API, SaaS, database)
// storing credentials and settings for use by workspace resources.
//
// The properties Metadata, Spec are required.
type IntegrationConnectionParam struct {
	// Common metadata fields shared by all Blaxel resources including name, labels,
	// timestamps, and ownership information
	Metadata MetadataParam `json:"metadata,omitzero,required"`
	// Specification defining the integration type, configuration parameters, and
	// encrypted credentials
	Spec IntegrationConnectionSpecParam `json:"spec,omitzero,required"`
	paramObj
}

func (r IntegrationConnectionParam) MarshalJSON() (data []byte, err error) {
	type shadow IntegrationConnectionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntegrationConnectionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specification defining the integration type, configuration parameters, and
// encrypted credentials
type IntegrationConnectionSpec struct {
	// Non-sensitive configuration parameters for the integration (e.g., organization
	// ID, region)
	Config map[string]string `json:"config"`
	// Integration provider type (e.g., openai, anthropic, github, slack)
	Integration string `json:"integration"`
	// Whether this connection uses sandbox/test credentials instead of production
	Sandbox bool `json:"sandbox"`
	// Encrypted credentials and API keys for authenticating with the external service
	Secret map[string]string `json:"secret"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Config      respjson.Field
		Integration respjson.Field
		Sandbox     respjson.Field
		Secret      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntegrationConnectionSpec) RawJSON() string { return r.JSON.raw }
func (r *IntegrationConnectionSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this IntegrationConnectionSpec to a
// IntegrationConnectionSpecParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// IntegrationConnectionSpecParam.Overrides()
func (r IntegrationConnectionSpec) ToParam() IntegrationConnectionSpecParam {
	return param.Override[IntegrationConnectionSpecParam](json.RawMessage(r.RawJSON()))
}

// Specification defining the integration type, configuration parameters, and
// encrypted credentials
type IntegrationConnectionSpecParam struct {
	// Integration provider type (e.g., openai, anthropic, github, slack)
	Integration param.Opt[string] `json:"integration,omitzero"`
	// Whether this connection uses sandbox/test credentials instead of production
	Sandbox param.Opt[bool] `json:"sandbox,omitzero"`
	// Non-sensitive configuration parameters for the integration (e.g., organization
	// ID, region)
	Config map[string]string `json:"config,omitzero"`
	// Encrypted credentials and API keys for authenticating with the external service
	Secret map[string]string `json:"secret,omitzero"`
	paramObj
}

func (r IntegrationConnectionSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow IntegrationConnectionSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *IntegrationConnectionSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IntegrationConnectionNewParams struct {
	// Configured connection to an external service (LLM provider, API, SaaS, database)
	// storing credentials and settings for use by workspace resources.
	IntegrationConnection IntegrationConnectionParam
	paramObj
}

func (r IntegrationConnectionNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.IntegrationConnection)
}
func (r *IntegrationConnectionNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.IntegrationConnection)
}

type IntegrationConnectionUpdateParams struct {
	// Configured connection to an external service (LLM provider, API, SaaS, database)
	// storing credentials and settings for use by workspace resources.
	IntegrationConnection IntegrationConnectionParam
	paramObj
}

func (r IntegrationConnectionUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.IntegrationConnection)
}
func (r *IntegrationConnectionUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.IntegrationConnection)
}
