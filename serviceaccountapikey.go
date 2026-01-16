// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blaxel

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/blaxel-ai/sdk-go/internal/apijson"
	"github.com/blaxel-ai/sdk-go/internal/requestconfig"
	"github.com/blaxel-ai/sdk-go/option"
	"github.com/blaxel-ai/sdk-go/packages/param"
	"github.com/blaxel-ai/sdk-go/packages/respjson"
)

// ServiceAccountAPIKeyService contains methods and other services that help with
// interacting with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewServiceAccountAPIKeyService] method instead.
type ServiceAccountAPIKeyService struct {
	Options []option.RequestOption
}

// NewServiceAccountAPIKeyService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewServiceAccountAPIKeyService(opts ...option.RequestOption) (r ServiceAccountAPIKeyService) {
	r = ServiceAccountAPIKeyService{}
	r.Options = opts
	return
}

// Creates a new long-lived API key for a service account. The full key value is
// only returned once at creation. API keys can have optional expiration dates.
func (r *ServiceAccountAPIKeyService) New(ctx context.Context, clientID string, body ServiceAccountAPIKeyNewParams, opts ...option.RequestOption) (res *APIKey, err error) {
	opts = slices.Concat(r.Options, opts)
	if clientID == "" {
		err = errors.New("missing required clientId parameter")
		return
	}
	path := fmt.Sprintf("service_accounts/%s/api_keys", clientID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns all long-lived API keys created for a service account. API keys provide
// an alternative to OAuth for simpler authentication scenarios.
func (r *ServiceAccountAPIKeyService) List(ctx context.Context, clientID string, opts ...option.RequestOption) (res *[]APIKey, err error) {
	opts = slices.Concat(r.Options, opts)
	if clientID == "" {
		err = errors.New("missing required clientId parameter")
		return
	}
	path := fmt.Sprintf("service_accounts/%s/api_keys", clientID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Revokes an API key for a service account. The key becomes invalid immediately
// and any requests using it will fail authentication.
func (r *ServiceAccountAPIKeyService) Delete(ctx context.Context, apiKeyID string, body ServiceAccountAPIKeyDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.ClientID == "" {
		err = errors.New("missing required clientId parameter")
		return
	}
	if apiKeyID == "" {
		err = errors.New("missing required apiKeyId parameter")
		return
	}
	path := fmt.Sprintf("service_accounts/%s/api_keys/%s", body.ClientID, apiKeyID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

type APIKey struct {
	// Api key id, to retrieve it from the API
	ID string `json:"id"`
	// Api key
	APIKey string `json:"apiKey"`
	// The date and time when the resource was created
	CreatedAt string `json:"createdAt"`
	// The user or service account who created the resource
	CreatedBy string `json:"createdBy"`
	// Duration until expiration. Supports formats like '30d' (30 days), '24h' (24
	// hours), '1w' (1 week). If not set, the API key never expires.
	ExpiresIn string `json:"expires_in"`
	// Name for the API key
	Name string `json:"name"`
	// User subject identifier
	Sub string `json:"sub"`
	// Subject type (user or service_account)
	SubType string `json:"sub_type"`
	// The date and time when the resource was updated
	UpdatedAt string `json:"updatedAt"`
	// The user or service account who updated the resource
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		APIKey      respjson.Field
		CreatedAt   respjson.Field
		CreatedBy   respjson.Field
		ExpiresIn   respjson.Field
		Name        respjson.Field
		Sub         respjson.Field
		SubType     respjson.Field
		UpdatedAt   respjson.Field
		UpdatedBy   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r APIKey) RawJSON() string { return r.JSON.raw }
func (r *APIKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ServiceAccountAPIKeyNewParams struct {
	// Expiration period for the API key. Supports formats like '30d' (30 days), '24h'
	// (24 hours), '1w' (1 week). If not set, the API key never expires.
	ExpiresIn param.Opt[string] `json:"expires_in,omitzero"`
	// Name for the API key
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r ServiceAccountAPIKeyNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ServiceAccountAPIKeyNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ServiceAccountAPIKeyNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ServiceAccountAPIKeyDeleteParams struct {
	ClientID string `path:"clientId,required" json:"-"`
	paramObj
}
