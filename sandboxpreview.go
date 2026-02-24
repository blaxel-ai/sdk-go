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

// SandboxPreviewService contains methods and other services that help with
// interacting with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSandboxPreviewService] method instead.
type SandboxPreviewService struct {
	Options []option.RequestOption
	Tokens  SandboxPreviewTokenService
}

// NewSandboxPreviewService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSandboxPreviewService(opts ...option.RequestOption) (r SandboxPreviewService) {
	r = SandboxPreviewService{}
	r.Options = opts
	r.Tokens = NewSandboxPreviewTokenService(opts...)
	return
}

// Create a preview
func (r *SandboxPreviewService) New(ctx context.Context, sandboxName string, body SandboxPreviewNewParams, opts ...option.RequestOption) (res *Preview, err error) {
	opts = slices.Concat(r.Options, opts)
	if sandboxName == "" {
		err = errors.New("missing required sandboxName parameter")
		return
	}
	path := fmt.Sprintf("sandboxes/%s/previews", sandboxName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns a Sandbox Preview by name.
func (r *SandboxPreviewService) Get(ctx context.Context, previewName string, query SandboxPreviewGetParams, opts ...option.RequestOption) (res *Preview, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.SandboxName == "" {
		err = errors.New("missing required sandboxName parameter")
		return
	}
	if previewName == "" {
		err = errors.New("missing required previewName parameter")
		return
	}
	path := fmt.Sprintf("sandboxes/%s/previews/%s", query.SandboxName, previewName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a Sandbox Preview by name.
func (r *SandboxPreviewService) Update(ctx context.Context, previewName string, params SandboxPreviewUpdateParams, opts ...option.RequestOption) (res *Preview, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.SandboxName == "" {
		err = errors.New("missing required sandboxName parameter")
		return
	}
	if previewName == "" {
		err = errors.New("missing required previewName parameter")
		return
	}
	path := fmt.Sprintf("sandboxes/%s/previews/%s", params.SandboxName, previewName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Returns a list of Sandbox Previews in the workspace.
func (r *SandboxPreviewService) List(ctx context.Context, sandboxName string, opts ...option.RequestOption) (res *[]Preview, err error) {
	opts = slices.Concat(r.Options, opts)
	if sandboxName == "" {
		err = errors.New("missing required sandboxName parameter")
		return
	}
	path := fmt.Sprintf("sandboxes/%s/previews", sandboxName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a Sandbox Preview by name.
func (r *SandboxPreviewService) Delete(ctx context.Context, previewName string, body SandboxPreviewDeleteParams, opts ...option.RequestOption) (res *Preview, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.SandboxName == "" {
		err = errors.New("missing required sandboxName parameter")
		return
	}
	if previewName == "" {
		err = errors.New("missing required previewName parameter")
		return
	}
	path := fmt.Sprintf("sandboxes/%s/previews/%s", body.SandboxName, previewName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Preview of a Resource
type Preview struct {
	// PreviewMetadata
	Metadata PreviewMetadata `json:"metadata" api:"required"`
	// Preview of a Resource
	Spec PreviewSpec `json:"spec" api:"required"`
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
func (r Preview) RawJSON() string { return r.JSON.raw }
func (r *Preview) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Preview to a PreviewParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PreviewParam.Overrides()
func (r Preview) ToParam() PreviewParam {
	return param.Override[PreviewParam](json.RawMessage(r.RawJSON()))
}

// Preview of a Resource
//
// The properties Metadata, Spec are required.
type PreviewParam struct {
	// PreviewMetadata
	Metadata PreviewMetadataParam `json:"metadata,omitzero" api:"required"`
	// Preview of a Resource
	Spec PreviewSpecParam `json:"spec,omitzero" api:"required"`
	paramObj
}

func (r PreviewParam) MarshalJSON() (data []byte, err error) {
	type shadow PreviewParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PreviewParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PreviewMetadata struct {
	// Preview name
	Name string `json:"name" api:"required"`
	// The date and time when the resource was created
	CreatedAt string `json:"createdAt"`
	// The user or service account who created the resource
	CreatedBy string `json:"createdBy"`
	// Model display name
	DisplayName string `json:"displayName"`
	// Resource name
	ResourceName string `json:"resourceName"`
	// Resource type
	ResourceType string `json:"resourceType"`
	// The date and time when the resource was updated
	UpdatedAt string `json:"updatedAt"`
	// The user or service account who updated the resource
	UpdatedBy string `json:"updatedBy"`
	// Workspace name
	Workspace string `json:"workspace"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name         respjson.Field
		CreatedAt    respjson.Field
		CreatedBy    respjson.Field
		DisplayName  respjson.Field
		ResourceName respjson.Field
		ResourceType respjson.Field
		UpdatedAt    respjson.Field
		UpdatedBy    respjson.Field
		Workspace    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PreviewMetadata) RawJSON() string { return r.JSON.raw }
func (r *PreviewMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PreviewMetadata to a PreviewMetadataParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PreviewMetadataParam.Overrides()
func (r PreviewMetadata) ToParam() PreviewMetadataParam {
	return param.Override[PreviewMetadataParam](json.RawMessage(r.RawJSON()))
}

// The property Name is required.
type PreviewMetadataParam struct {
	// Preview name
	Name string `json:"name" api:"required"`
	// Model display name
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	// Resource name
	ResourceName param.Opt[string] `json:"resourceName,omitzero"`
	// Resource type
	ResourceType param.Opt[string] `json:"resourceType,omitzero"`
	// Workspace name
	Workspace param.Opt[string] `json:"workspace,omitzero"`
	paramObj
}

func (r PreviewMetadataParam) MarshalJSON() (data []byte, err error) {
	type shadow PreviewMetadataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PreviewMetadataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Preview of a Resource
type PreviewSpec struct {
	// Custom domain bound to this preview
	CustomDomain string `json:"customDomain"`
	// The expiration date for the preview in ISO 8601 format - 2024-12-31T23:59:59Z
	Expires string `json:"expires"`
	// Port of the preview
	Port int64 `json:"port"`
	// Prefix URL
	PrefixURL string `json:"prefixUrl"`
	// Whether the preview is public
	Public bool `json:"public"`
	// Region where the preview is deployed, this is readonly
	Region string `json:"region"`
	// Those headers will be set in all requests to your preview. This is especially
	// useful to set the Authorization header.
	RequestHeaders map[string]string `json:"requestHeaders"`
	// Those headers will be set in all responses of your preview. This is especially
	// useful to set the CORS headers.
	ResponseHeaders map[string]string `json:"responseHeaders"`
	// Time to live for the preview (e.g., "1h", "24h", "7d"). After this duration, the
	// preview will be automatically deleted.
	Ttl string `json:"ttl"`
	// URL of the preview
	URL string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CustomDomain    respjson.Field
		Expires         respjson.Field
		Port            respjson.Field
		PrefixURL       respjson.Field
		Public          respjson.Field
		Region          respjson.Field
		RequestHeaders  respjson.Field
		ResponseHeaders respjson.Field
		Ttl             respjson.Field
		URL             respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PreviewSpec) RawJSON() string { return r.JSON.raw }
func (r *PreviewSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PreviewSpec to a PreviewSpecParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PreviewSpecParam.Overrides()
func (r PreviewSpec) ToParam() PreviewSpecParam {
	return param.Override[PreviewSpecParam](json.RawMessage(r.RawJSON()))
}

// Preview of a Resource
type PreviewSpecParam struct {
	// Custom domain bound to this preview
	CustomDomain param.Opt[string] `json:"customDomain,omitzero"`
	// The expiration date for the preview in ISO 8601 format - 2024-12-31T23:59:59Z
	Expires param.Opt[string] `json:"expires,omitzero"`
	// Port of the preview
	Port param.Opt[int64] `json:"port,omitzero"`
	// Prefix URL
	PrefixURL param.Opt[string] `json:"prefixUrl,omitzero"`
	// Whether the preview is public
	Public param.Opt[bool] `json:"public,omitzero"`
	// Time to live for the preview (e.g., "1h", "24h", "7d"). After this duration, the
	// preview will be automatically deleted.
	Ttl param.Opt[string] `json:"ttl,omitzero"`
	// Those headers will be set in all requests to your preview. This is especially
	// useful to set the Authorization header.
	RequestHeaders map[string]string `json:"requestHeaders,omitzero"`
	// Those headers will be set in all responses of your preview. This is especially
	// useful to set the CORS headers.
	ResponseHeaders map[string]string `json:"responseHeaders,omitzero"`
	paramObj
}

func (r PreviewSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow PreviewSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PreviewSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SandboxPreviewNewParams struct {
	// Preview of a Resource
	Preview PreviewParam
	paramObj
}

func (r SandboxPreviewNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Preview)
}
func (r *SandboxPreviewNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Preview)
}

type SandboxPreviewGetParams struct {
	SandboxName string `path:"sandboxName" api:"required" json:"-"`
	paramObj
}

type SandboxPreviewUpdateParams struct {
	SandboxName string `path:"sandboxName" api:"required" json:"-"`
	// Preview of a Resource
	Preview PreviewParam
	paramObj
}

func (r SandboxPreviewUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Preview)
}
func (r *SandboxPreviewUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Preview)
}

type SandboxPreviewDeleteParams struct {
	SandboxName string `path:"sandboxName" api:"required" json:"-"`
	paramObj
}
