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
	"github.com/stainless-sdks/blaxel-go/packages/param"
	"github.com/stainless-sdks/blaxel-go/packages/respjson"
)

// ServiceAccountService contains methods and other services that help with
// interacting with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewServiceAccountService] method instead.
type ServiceAccountService struct {
	Options []option.RequestOption
	APIKeys ServiceAccountAPIKeyService
}

// NewServiceAccountService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewServiceAccountService(opts ...option.RequestOption) (r ServiceAccountService) {
	r = ServiceAccountService{}
	r.Options = opts
	r.APIKeys = NewServiceAccountAPIKeyService(opts...)
	return
}

// Creates a service account in the workspace.
func (r *ServiceAccountService) New(ctx context.Context, body ServiceAccountNewParams, opts ...option.RequestOption) (res *ServiceAccountNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "service_accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Updates a service account.
func (r *ServiceAccountService) Update(ctx context.Context, clientID string, body ServiceAccountUpdateParams, opts ...option.RequestOption) (res *ServiceAccountUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if clientID == "" {
		err = errors.New("missing required clientId parameter")
		return
	}
	path := fmt.Sprintf("service_accounts/%s", clientID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Returns a list of all service accounts in the workspace.
func (r *ServiceAccountService) List(ctx context.Context, opts ...option.RequestOption) (res *[]ServiceAccountListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "service_accounts"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a service account.
func (r *ServiceAccountService) Delete(ctx context.Context, clientID string, opts ...option.RequestOption) (res *ServiceAccountDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if clientID == "" {
		err = errors.New("missing required clientId parameter")
		return
	}
	path := fmt.Sprintf("service_accounts/%s", clientID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type ServiceAccountNewResponse struct {
	// Service account client ID
	ClientID string `json:"client_id"`
	// Service account client secret (only returned on creation)
	ClientSecret string `json:"client_secret"`
	// Creation timestamp
	CreatedAt string `json:"created_at"`
	// Service account description
	Description string `json:"description"`
	// Service account name
	Name string `json:"name"`
	// Last update timestamp
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClientID     respjson.Field
		ClientSecret respjson.Field
		CreatedAt    respjson.Field
		Description  respjson.Field
		Name         respjson.Field
		UpdatedAt    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ServiceAccountNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ServiceAccountNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ServiceAccountUpdateResponse struct {
	// Service account client ID
	ClientID string `json:"client_id"`
	// Creation timestamp
	CreatedAt string `json:"created_at"`
	// Service account description
	Description string `json:"description"`
	// Service account name
	Name string `json:"name"`
	// Last update timestamp
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClientID    respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		Name        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ServiceAccountUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *ServiceAccountUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ServiceAccountListResponse struct {
	// Service account client ID
	ClientID string `json:"client_id"`
	// Creation timestamp
	CreatedAt string `json:"created_at"`
	// Service account description
	Description string `json:"description"`
	// Service account name
	Name string `json:"name"`
	// Last update timestamp
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClientID    respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		Name        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ServiceAccountListResponse) RawJSON() string { return r.JSON.raw }
func (r *ServiceAccountListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ServiceAccountDeleteResponse struct {
	// Service account client ID
	ClientID string `json:"client_id"`
	// Creation timestamp
	CreatedAt string `json:"created_at"`
	// Service account description
	Description string `json:"description"`
	// Service account name
	Name string `json:"name"`
	// Last update timestamp
	UpdatedAt string `json:"updated_at"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ClientID    respjson.Field
		CreatedAt   respjson.Field
		Description respjson.Field
		Name        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ServiceAccountDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *ServiceAccountDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ServiceAccountNewParams struct {
	// Service account name
	Name string `json:"name,required"`
	// Service account description
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r ServiceAccountNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ServiceAccountNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ServiceAccountNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ServiceAccountUpdateParams struct {
	// Service account description
	Description param.Opt[string] `json:"description,omitzero"`
	// Service account name
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r ServiceAccountUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow ServiceAccountUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ServiceAccountUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
