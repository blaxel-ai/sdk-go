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
	"github.com/stainless-sdks/blaxel-go/internal/requestconfig"
	"github.com/stainless-sdks/blaxel-go/option"
	"github.com/stainless-sdks/blaxel-go/packages/param"
	"github.com/stainless-sdks/blaxel-go/packages/respjson"
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

// Creates an API key for a service account.
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

// Returns a list of all API keys for a service account.
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

// Deletes an API key for a service account.
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

// Long-lived API key for accessing Blaxel
type APIKey struct {
	// Api key id, to retrieve it from the API
	ID string `json:"id"`
	// Api key
	APIKey string `json:"apiKey"`
	// Duration until expiration (in seconds)
	ExpiresIn string `json:"expires_in"`
	// Name for the API key
	Name string `json:"name"`
	// User subject identifier
	Sub string `json:"sub"`
	// Subject type
	SubType string `json:"sub_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		APIKey      respjson.Field
		ExpiresIn   respjson.Field
		Name        respjson.Field
		Sub         respjson.Field
		SubType     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	TimeFields
	OwnerFields
}

// Returns the unmodified JSON received from the API
func (r APIKey) RawJSON() string { return r.JSON.raw }
func (r *APIKey) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Owner fields for Persistance
type OwnerFields struct {
	// The user or service account who created the resource
	CreatedBy string `json:"createdBy"`
	// The user or service account who updated the resource
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedBy   respjson.Field
		UpdatedBy   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OwnerFields) RawJSON() string { return r.JSON.raw }
func (r *OwnerFields) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this OwnerFields to a OwnerFieldsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// OwnerFieldsParam.Overrides()
func (r OwnerFields) ToParam() OwnerFieldsParam {
	return param.Override[OwnerFieldsParam](json.RawMessage(r.RawJSON()))
}

// Owner fields for Persistance
type OwnerFieldsParam struct {
	paramObj
}

func (r OwnerFieldsParam) MarshalJSON() (data []byte, err error) {
	type shadow OwnerFieldsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *OwnerFieldsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Time fields for Persistance
type TimeFields struct {
	// The date and time when the resource was created
	CreatedAt string `json:"createdAt"`
	// The date and time when the resource was updated
	UpdatedAt string `json:"updatedAt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TimeFields) RawJSON() string { return r.JSON.raw }
func (r *TimeFields) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TimeFields to a TimeFieldsParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TimeFieldsParam.Overrides()
func (r TimeFields) ToParam() TimeFieldsParam {
	return param.Override[TimeFieldsParam](json.RawMessage(r.RawJSON()))
}

// Time fields for Persistance
type TimeFieldsParam struct {
	paramObj
}

func (r TimeFieldsParam) MarshalJSON() (data []byte, err error) {
	type shadow TimeFieldsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TimeFieldsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ServiceAccountAPIKeyNewParams struct {
	// Expiration period for the API key
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
