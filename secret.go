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

// SecretService contains methods and other services that help with interacting
// with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSecretService] method instead.
type SecretService struct {
	Options []option.RequestOption
}

// NewSecretService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSecretService(opts ...option.RequestOption) (r SecretService) {
	r = SecretService{}
	r.Options = opts
	return
}

// Creates or updates a secret by name. Each call stores a new immutable version;
// the latest version is the one resolved by the proxy.
func (r *SecretService) New(ctx context.Context, body SecretNewParams, opts ...option.RequestOption) (res *Secret, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "secrets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Lists the secrets of the workspace. Values are write-only and never returned.
func (r *SecretService) List(ctx context.Context, opts ...option.RequestOption) (res *[]Secret, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "secrets"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Deletes a secret and all its versions from the workspace.
func (r *SecretService) Delete(ctx context.Context, secretName string, opts ...option.RequestOption) (res *Secret, err error) {
	opts = slices.Concat(r.Options, opts)
	if secretName == "" {
		err = errors.New("missing required secretName parameter")
		return nil, err
	}
	path := fmt.Sprintf("secrets/%s", secretName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Workspace secret, referenced from proxy routing as {{SECRET:name}}. Values are
// write-only and can never be read back. Every upsert stores a new immutable
// version and the latest one is resolved at runtime, only for requests to its
// destinations.
type Secret struct {
	// Destinations the secret resolves for, same syntax as routing destinations: a
	// host, optionally followed by a path glob. The secret is never substituted toward
	// any other destination.
	Destinations []string `json:"destinations" api:"required"`
	// Case-sensitive secret name
	Name string `json:"name" api:"required"`
	// The date and time when the secret was first created
	CreatedAt string `json:"createdAt"`
	// Delivery of the latest change to every region. Absent once every region applied
	// it; PENDING while a region has not yet; DEGRADED when a region has been behind
	// for more than 15 minutes; DELETING for a deleted secret not yet removed from
	// every region.
	//
	// Any of "PENDING", "DEGRADED", "DELETING".
	Status SecretStatus `json:"status"`
	// The date and time of the latest version
	UpdatedAt string `json:"updatedAt"`
	// Secret value. Write-only, never returned by the API
	Value string `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Destinations respjson.Field
		Name         respjson.Field
		CreatedAt    respjson.Field
		Status       respjson.Field
		UpdatedAt    respjson.Field
		Value        respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Secret) RawJSON() string { return r.JSON.raw }
func (r *Secret) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Secret to a SecretParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SecretParam.Overrides()
func (r Secret) ToParam() SecretParam {
	return param.Override[SecretParam](json.RawMessage(r.RawJSON()))
}

// Delivery of the latest change to every region. Absent once every region applied
// it; PENDING while a region has not yet; DEGRADED when a region has been behind
// for more than 15 minutes; DELETING for a deleted secret not yet removed from
// every region.
type SecretStatus string

const (
	SecretStatusPending  SecretStatus = "PENDING"
	SecretStatusDegraded SecretStatus = "DEGRADED"
	SecretStatusDeleting SecretStatus = "DELETING"
)

// Workspace secret, referenced from proxy routing as {{SECRET:name}}. Values are
// write-only and can never be read back. Every upsert stores a new immutable
// version and the latest one is resolved at runtime, only for requests to its
// destinations.
//
// The properties Destinations, Name are required.
type SecretParam struct {
	// Destinations the secret resolves for, same syntax as routing destinations: a
	// host, optionally followed by a path glob. The secret is never substituted toward
	// any other destination.
	Destinations []string `json:"destinations,omitzero" api:"required"`
	// Case-sensitive secret name
	Name string `json:"name" api:"required"`
	// Secret value. Write-only, never returned by the API
	Value param.Opt[string] `json:"value,omitzero"`
	paramObj
}

func (r SecretParam) MarshalJSON() (data []byte, err error) {
	type shadow SecretParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SecretParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SecretNewParams struct {
	// Workspace secret, referenced from proxy routing as {{SECRET:name}}. Values are
	// write-only and can never be read back. Every upsert stores a new immutable
	// version and the latest one is resolved at runtime, only for requests to its
	// destinations.
	Secret SecretParam
	paramObj
}

func (r SecretNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Secret)
}
func (r *SecretNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
