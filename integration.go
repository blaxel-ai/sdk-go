// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blaxel

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/blaxel-go/internal/apijson"
	"github.com/stainless-sdks/blaxel-go/internal/requestconfig"
	"github.com/stainless-sdks/blaxel-go/option"
	"github.com/stainless-sdks/blaxel-go/packages/respjson"
)

// IntegrationService contains methods and other services that help with
// interacting with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIntegrationService] method instead.
type IntegrationService struct {
	Options     []option.RequestOption
	Connections IntegrationConnectionService
}

// NewIntegrationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewIntegrationService(opts ...option.RequestOption) (r IntegrationService) {
	r = IntegrationService{}
	r.Options = opts
	r.Connections = NewIntegrationConnectionService(opts...)
	return
}

// Returns integration information by name.
func (r *IntegrationService) Get(ctx context.Context, integrationName string, opts ...option.RequestOption) (res *IntegrationGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if integrationName == "" {
		err = errors.New("missing required integrationName parameter")
		return
	}
	path := fmt.Sprintf("integrations/%s", integrationName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Integration
type IntegrationGetResponse struct {
	// Integration additional infos
	AdditionalInfos map[string]string `json:"additionalInfos"`
	// Integration endpoints
	Endpoints map[string]IntegrationGetResponseEndpoint `json:"endpoints"`
	// Integration headers
	Headers map[string]string `json:"headers"`
	// Integration name
	Name string `json:"name"`
	// Integration organizations
	Organizations []IntegrationGetResponseOrganization `json:"organizations"`
	// Integration query params
	Params map[string]string `json:"params"`
	// Integration repositories
	Repositories []IntegrationGetResponseRepository `json:"repositories"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AdditionalInfos respjson.Field
		Endpoints       respjson.Field
		Headers         respjson.Field
		Name            respjson.Field
		Organizations   respjson.Field
		Params          respjson.Field
		Repositories    respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntegrationGetResponse) RawJSON() string { return r.JSON.raw }
func (r *IntegrationGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Integration endpoint
type IntegrationGetResponseEndpoint struct {
	// Integration endpoint token
	Token IntegrationGetResponseEndpointToken `json:"token"`
	// Integration endpoint body
	Body string `json:"body"`
	// Integration endpoint ignore models
	IgnoreModels []map[string]any `json:"ignoreModels"`
	// Integration endpoint method
	Method string `json:"method"`
	// Integration endpoint models
	Models []map[string]any `json:"models"`
	// Integration endpoint stream key
	StreamKey string `json:"streamKey"`
	// Integration endpoint token
	StreamToken IntegrationGetResponseEndpointStreamToken `json:"streamToken"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token        respjson.Field
		Body         respjson.Field
		IgnoreModels respjson.Field
		Method       respjson.Field
		Models       respjson.Field
		StreamKey    respjson.Field
		StreamToken  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntegrationGetResponseEndpoint) RawJSON() string { return r.JSON.raw }
func (r *IntegrationGetResponseEndpoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Integration endpoint token
type IntegrationGetResponseEndpointToken struct {
	// Integration endpoint token received
	Received string `json:"received"`
	// Integration endpoint token sent
	Sent string `json:"sent"`
	// Integration endpoint token total
	Total string `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Received    respjson.Field
		Sent        respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntegrationGetResponseEndpointToken) RawJSON() string { return r.JSON.raw }
func (r *IntegrationGetResponseEndpointToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Integration endpoint token
type IntegrationGetResponseEndpointStreamToken struct {
	// Integration endpoint token received
	Received string `json:"received"`
	// Integration endpoint token sent
	Sent string `json:"sent"`
	// Integration endpoint token total
	Total string `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Received    respjson.Field
		Sent        respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntegrationGetResponseEndpointStreamToken) RawJSON() string { return r.JSON.raw }
func (r *IntegrationGetResponseEndpointStreamToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Integration organization
type IntegrationGetResponseOrganization struct {
	// Provider organization ID
	ID string `json:"id"`
	// Provider organization avatar URL
	AvatarURL string `json:"avatar_url"`
	// Provider organization display name
	DisplayName string `json:"displayName"`
	// Provider organization name
	Name        string         `json:"name"`
	ExtraFields map[string]any `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		AvatarURL   respjson.Field
		DisplayName respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntegrationGetResponseOrganization) RawJSON() string { return r.JSON.raw }
func (r *IntegrationGetResponseOrganization) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Integration repository
type IntegrationGetResponseRepository struct {
	// Repository ID
	ID string `json:"id"`
	// Whether the repository has Blaxel imports
	IsBl bool `json:"isBl"`
	// Repository name
	Name string `json:"name"`
	// Repository owner
	Organization string `json:"organization"`
	// Repository URL
	URL         string         `json:"url"`
	ExtraFields map[string]any `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		IsBl         respjson.Field
		Name         respjson.Field
		Organization respjson.Field
		URL          respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntegrationGetResponseRepository) RawJSON() string { return r.JSON.raw }
func (r *IntegrationGetResponseRepository) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
