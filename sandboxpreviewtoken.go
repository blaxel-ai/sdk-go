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

// SandboxPreviewTokenService contains methods and other services that help with
// interacting with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSandboxPreviewTokenService] method instead.
type SandboxPreviewTokenService struct {
	Options []option.RequestOption
}

// NewSandboxPreviewTokenService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSandboxPreviewTokenService(opts ...option.RequestOption) (r SandboxPreviewTokenService) {
	r = SandboxPreviewTokenService{}
	r.Options = opts
	return
}

// Creates a token for a Sandbox Preview.
func (r *SandboxPreviewTokenService) New(ctx context.Context, previewName string, params SandboxPreviewTokenNewParams, opts ...option.RequestOption) (res *PreviewToken, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.SandboxName == "" {
		err = errors.New("missing required sandboxName parameter")
		return
	}
	if previewName == "" {
		err = errors.New("missing required previewName parameter")
		return
	}
	path := fmt.Sprintf("sandboxes/%s/previews/%s/tokens", params.SandboxName, previewName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Gets tokens for a Sandbox Preview.
func (r *SandboxPreviewTokenService) Get(ctx context.Context, previewName string, query SandboxPreviewTokenGetParams, opts ...option.RequestOption) (res *[]PreviewToken, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.SandboxName == "" {
		err = errors.New("missing required sandboxName parameter")
		return
	}
	if previewName == "" {
		err = errors.New("missing required previewName parameter")
		return
	}
	path := fmt.Sprintf("sandboxes/%s/previews/%s/tokens", query.SandboxName, previewName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a token for a Sandbox Preview by name.
func (r *SandboxPreviewTokenService) Delete(ctx context.Context, tokenName string, body SandboxPreviewTokenDeleteParams, opts ...option.RequestOption) (res *SandboxPreviewTokenDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.SandboxName == "" {
		err = errors.New("missing required sandboxName parameter")
		return
	}
	if body.PreviewName == "" {
		err = errors.New("missing required previewName parameter")
		return
	}
	if tokenName == "" {
		err = errors.New("missing required tokenName parameter")
		return
	}
	path := fmt.Sprintf("sandboxes/%s/previews/%s/tokens/%s", body.SandboxName, body.PreviewName, tokenName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Token for a Preview
type PreviewToken struct {
	// PreviewTokenMetadata
	Metadata PreviewTokenMetadata `json:"metadata,required"`
	// Spec for a Preview Token
	Spec PreviewTokenSpec `json:"spec,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metadata    respjson.Field
		Spec        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PreviewToken) RawJSON() string { return r.JSON.raw }
func (r *PreviewToken) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PreviewToken to a PreviewTokenParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PreviewTokenParam.Overrides()
func (r PreviewToken) ToParam() PreviewTokenParam {
	return param.Override[PreviewTokenParam](json.RawMessage(r.RawJSON()))
}

// PreviewTokenMetadata
type PreviewTokenMetadata struct {
	// Token name
	Name string `json:"name,required"`
	// Preview name
	PreviewName string `json:"previewName"`
	// Resource name
	ResourceName string `json:"resourceName"`
	// Resource type
	ResourceType string `json:"resourceType"`
	// Workspace name
	Workspace string `json:"workspace"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name         respjson.Field
		PreviewName  respjson.Field
		ResourceName respjson.Field
		ResourceType respjson.Field
		Workspace    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PreviewTokenMetadata) RawJSON() string { return r.JSON.raw }
func (r *PreviewTokenMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Spec for a Preview Token
type PreviewTokenSpec struct {
	// Token
	Token string `json:"token"`
	// Whether the token is expired
	Expired bool `json:"expired"`
	// Expiration time of the token
	ExpiresAt string `json:"expiresAt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Token       respjson.Field
		Expired     respjson.Field
		ExpiresAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PreviewTokenSpec) RawJSON() string { return r.JSON.raw }
func (r *PreviewTokenSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Token for a Preview
//
// The properties Metadata, Spec are required.
type PreviewTokenParam struct {
	// PreviewTokenMetadata
	Metadata PreviewTokenMetadataParam `json:"metadata,omitzero,required"`
	// Spec for a Preview Token
	Spec PreviewTokenSpecParam `json:"spec,omitzero,required"`
	paramObj
}

func (r PreviewTokenParam) MarshalJSON() (data []byte, err error) {
	type shadow PreviewTokenParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PreviewTokenParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PreviewTokenMetadata
//
// The property Name is required.
type PreviewTokenMetadataParam struct {
	// Token name
	Name string `json:"name,required"`
	// Preview name
	PreviewName param.Opt[string] `json:"previewName,omitzero"`
	// Resource name
	ResourceName param.Opt[string] `json:"resourceName,omitzero"`
	// Resource type
	ResourceType param.Opt[string] `json:"resourceType,omitzero"`
	// Workspace name
	Workspace param.Opt[string] `json:"workspace,omitzero"`
	paramObj
}

func (r PreviewTokenMetadataParam) MarshalJSON() (data []byte, err error) {
	type shadow PreviewTokenMetadataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PreviewTokenMetadataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Spec for a Preview Token
type PreviewTokenSpecParam struct {
	// Expiration time of the token
	ExpiresAt param.Opt[string] `json:"expiresAt,omitzero"`
	paramObj
}

func (r PreviewTokenSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow PreviewTokenSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PreviewTokenSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SandboxPreviewTokenDeleteResponse struct {
	// Success message
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SandboxPreviewTokenDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *SandboxPreviewTokenDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SandboxPreviewTokenNewParams struct {
	SandboxName string `path:"sandboxName,required" json:"-"`
	// Token for a Preview
	PreviewToken PreviewTokenParam
	paramObj
}

func (r SandboxPreviewTokenNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.PreviewToken)
}
func (r *SandboxPreviewTokenNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.PreviewToken)
}

type SandboxPreviewTokenGetParams struct {
	SandboxName string `path:"sandboxName,required" json:"-"`
	paramObj
}

type SandboxPreviewTokenDeleteParams struct {
	SandboxName string `path:"sandboxName,required" json:"-"`
	PreviewName string `path:"previewName,required" json:"-"`
	paramObj
}
