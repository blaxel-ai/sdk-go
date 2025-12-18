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

// Create a connection for an integration.
func (r *IntegrationConnectionService) New(ctx context.Context, body IntegrationConnectionNewParams, opts ...option.RequestOption) (res *IntegrationConnection, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "integrations/connections"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns an integration connection by integration name and connection name.
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

// Update an integration connection by integration name and connection name.
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

// Returns a list of all connections integrations in the workspace.
func (r *IntegrationConnectionService) List(ctx context.Context, opts ...option.RequestOption) (res *[]IntegrationConnection, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "integrations/connections"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes an integration connection by integration name and connection name.
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

// Integration Connection
type IntegrationConnection struct {
	// Metadata
	Metadata Metadata `json:"metadata,required"`
	// Integration connection specification
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

// Integration connection specification
type IntegrationConnectionSpec struct {
	// Additional configuration for the integration
	Config map[string]string `json:"config"`
	// Integration type
	Integration string `json:"integration"`
	// Sandbox mode
	Sandbox bool `json:"sandbox"`
	// Integration secret
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

// Integration Connection
//
// The properties Metadata, Spec are required.
type IntegrationConnectionParam struct {
	// Metadata
	Metadata MetadataParam `json:"metadata,omitzero,required"`
	// Integration connection specification
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

// Integration connection specification
type IntegrationConnectionSpecParam struct {
	// Integration type
	Integration param.Opt[string] `json:"integration,omitzero"`
	// Sandbox mode
	Sandbox param.Opt[bool] `json:"sandbox,omitzero"`
	// Additional configuration for the integration
	Config map[string]string `json:"config,omitzero"`
	// Integration secret
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
	// Integration Connection
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
	// Integration Connection
	IntegrationConnection IntegrationConnectionParam
	paramObj
}

func (r IntegrationConnectionUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.IntegrationConnection)
}
func (r *IntegrationConnectionUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.IntegrationConnection)
}
