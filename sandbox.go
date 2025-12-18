// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blaxel

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/blaxel-go/internal/apijson"
	"github.com/stainless-sdks/blaxel-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/blaxel-go/internal/encoding/json"
	"github.com/stainless-sdks/blaxel-go/internal/requestconfig"
	"github.com/stainless-sdks/blaxel-go/option"
	"github.com/stainless-sdks/blaxel-go/packages/param"
	"github.com/stainless-sdks/blaxel-go/packages/respjson"
)

// SandboxService contains methods and other services that help with interacting
// with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSandboxService] method instead.
type SandboxService struct {
	Options  []option.RequestOption
	Previews SandboxPreviewService
}

// NewSandboxService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSandboxService(opts ...option.RequestOption) (r SandboxService) {
	r = SandboxService{}
	r.Options = opts
	r.Previews = NewSandboxPreviewService(opts...)
	return
}

// Creates a Sandbox.
func (r *SandboxService) New(ctx context.Context, body SandboxNewParams, opts ...option.RequestOption) (res *Sandbox, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "sandboxes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns a Sandbox by name.
func (r *SandboxService) Get(ctx context.Context, sandboxName string, query SandboxGetParams, opts ...option.RequestOption) (res *Sandbox, err error) {
	opts = slices.Concat(r.Options, opts)
	if sandboxName == "" {
		err = errors.New("missing required sandboxName parameter")
		return
	}
	path := fmt.Sprintf("sandboxes/%s", sandboxName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Update a Sandbox by name.
func (r *SandboxService) Update(ctx context.Context, sandboxName string, body SandboxUpdateParams, opts ...option.RequestOption) (res *Sandbox, err error) {
	opts = slices.Concat(r.Options, opts)
	if sandboxName == "" {
		err = errors.New("missing required sandboxName parameter")
		return
	}
	path := fmt.Sprintf("sandboxes/%s", sandboxName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Returns a list of all Sandboxes in the workspace.
func (r *SandboxService) List(ctx context.Context, opts ...option.RequestOption) (res *[]Sandbox, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "sandboxes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a Sandbox by name.
func (r *SandboxService) Delete(ctx context.Context, sandboxName string, opts ...option.RequestOption) (res *Sandbox, err error) {
	opts = slices.Concat(r.Options, opts)
	if sandboxName == "" {
		err = errors.New("missing required sandboxName parameter")
		return
	}
	path := fmt.Sprintf("sandboxes/%s", sandboxName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// List sandbox hub definitions
func (r *SandboxService) GetHub(ctx context.Context, opts ...option.RequestOption) (res *[]SandboxGetHubResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "sandbox/hub"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Expiration policy for sandbox lifecycle management
type ExpirationPolicy struct {
	// Action to take when policy is triggered
	//
	// Any of "delete".
	Action ExpirationPolicyAction `json:"action"`
	// Type of expiration policy
	//
	// Any of "ttl-idle", "ttl-max-age", "date".
	Type ExpirationPolicyType `json:"type"`
	// Duration value (e.g., '1h', '24h', '7d')
	Value string `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Action      respjson.Field
		Type        respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ExpirationPolicy) RawJSON() string { return r.JSON.raw }
func (r *ExpirationPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this ExpirationPolicy to a ExpirationPolicyParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// ExpirationPolicyParam.Overrides()
func (r ExpirationPolicy) ToParam() ExpirationPolicyParam {
	return param.Override[ExpirationPolicyParam](json.RawMessage(r.RawJSON()))
}

// Action to take when policy is triggered
type ExpirationPolicyAction string

const (
	ExpirationPolicyActionDelete ExpirationPolicyAction = "delete"
)

// Type of expiration policy
type ExpirationPolicyType string

const (
	ExpirationPolicyTypeTtlIdle   ExpirationPolicyType = "ttl-idle"
	ExpirationPolicyTypeTtlMaxAge ExpirationPolicyType = "ttl-max-age"
	ExpirationPolicyTypeDate      ExpirationPolicyType = "date"
)

// Expiration policy for sandbox lifecycle management
type ExpirationPolicyParam struct {
	// Duration value (e.g., '1h', '24h', '7d')
	Value param.Opt[string] `json:"value,omitzero"`
	// Action to take when policy is triggered
	//
	// Any of "delete".
	Action ExpirationPolicyAction `json:"action,omitzero"`
	// Type of expiration policy
	//
	// Any of "ttl-idle", "ttl-max-age", "date".
	Type ExpirationPolicyType `json:"type,omitzero"`
	paramObj
}

func (r ExpirationPolicyParam) MarshalJSON() (data []byte, err error) {
	type shadow ExpirationPolicyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ExpirationPolicyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A port for a resource
type Port struct {
	// The name of the port
	Name string `json:"name"`
	// The protocol of the port
	//
	// Any of "HTTP", "TCP", "UDP".
	Protocol PortProtocol `json:"protocol"`
	// The target port of the port
	Target int64 `json:"target"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Protocol    respjson.Field
		Target      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Port) RawJSON() string { return r.JSON.raw }
func (r *Port) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Port to a PortParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PortParam.Overrides()
func (r Port) ToParam() PortParam {
	return param.Override[PortParam](json.RawMessage(r.RawJSON()))
}

// The protocol of the port
type PortProtocol string

const (
	PortProtocolHTTP PortProtocol = "HTTP"
	PortProtocolTcp  PortProtocol = "TCP"
	PortProtocolUdp  PortProtocol = "UDP"
)

// A port for a resource
type PortParam struct {
	// The name of the port
	Name param.Opt[string] `json:"name,omitzero"`
	// The target port of the port
	Target param.Opt[int64] `json:"target,omitzero"`
	// The protocol of the port
	//
	// Any of "HTTP", "TCP", "UDP".
	Protocol PortProtocol `json:"protocol,omitzero"`
	paramObj
}

func (r PortParam) MarshalJSON() (data []byte, err error) {
	type shadow PortParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PortParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Micro VM for running agentic tasks
type Sandbox struct {
	// Metadata
	Metadata Metadata `json:"metadata,required"`
	// Sandbox specification for API
	Spec SandboxSpec `json:"spec,required"`
	// Core events
	Events []CoreEvent `json:"events"`
	// Last time the sandbox was used (read-only, managed by the system)
	LastUsedAt string `json:"lastUsedAt"`
	// Sandbox status
	//
	// Any of "DELETING", "TERMINATED", "FAILED", "DEACTIVATED", "DEACTIVATING",
	// "UPLOADING", "BUILDING", "DEPLOYING", "DEPLOYED".
	Status Status `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metadata    respjson.Field
		Spec        respjson.Field
		Events      respjson.Field
		LastUsedAt  respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Sandbox) RawJSON() string { return r.JSON.raw }
func (r *Sandbox) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Sandbox to a SandboxParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SandboxParam.Overrides()
func (r Sandbox) ToParam() SandboxParam {
	return param.Override[SandboxParam](json.RawMessage(r.RawJSON()))
}

// Micro VM for running agentic tasks
//
// The properties Metadata, Spec are required.
type SandboxParam struct {
	// Metadata
	Metadata MetadataParam `json:"metadata,omitzero,required"`
	// Sandbox specification for API
	Spec SandboxSpecParam `json:"spec,omitzero,required"`
	paramObj
}

func (r SandboxParam) MarshalJSON() (data []byte, err error) {
	type shadow SandboxParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SandboxParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lifecycle configuration for sandbox management
type SandboxLifecycle struct {
	// List of expiration policies
	ExpirationPolicies []ExpirationPolicy `json:"expirationPolicies"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpirationPolicies respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SandboxLifecycle) RawJSON() string { return r.JSON.raw }
func (r *SandboxLifecycle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SandboxLifecycle to a SandboxLifecycleParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SandboxLifecycleParam.Overrides()
func (r SandboxLifecycle) ToParam() SandboxLifecycleParam {
	return param.Override[SandboxLifecycleParam](json.RawMessage(r.RawJSON()))
}

// Lifecycle configuration for sandbox management
type SandboxLifecycleParam struct {
	// List of expiration policies
	ExpirationPolicies []ExpirationPolicyParam `json:"expirationPolicies,omitzero"`
	paramObj
}

func (r SandboxLifecycleParam) MarshalJSON() (data []byte, err error) {
	type shadow SandboxLifecycleParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SandboxLifecycleParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Runtime configuration for Sandbox
type SandboxRuntime struct {
	// The env variables to set in the sandbox. Should be a list of Kubernetes EnvVar
	// types
	Envs []map[string]any `json:"envs"`
	// The expiration date for the sandbox in ISO 8601 format - 2024-12-31T23:59:59Z
	Expires string `json:"expires"`
	// The Docker image for the sandbox
	Image string `json:"image"`
	// The memory for the sandbox in MB
	Memory int64 `json:"memory"`
	// The exposed ports of the sandbox
	Ports []Port `json:"ports"`
	// The TTL for the sandbox in seconds - 30m, 24h, 7d
	Ttl string `json:"ttl"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Envs        respjson.Field
		Expires     respjson.Field
		Image       respjson.Field
		Memory      respjson.Field
		Ports       respjson.Field
		Ttl         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SandboxRuntime) RawJSON() string { return r.JSON.raw }
func (r *SandboxRuntime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SandboxRuntime to a SandboxRuntimeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SandboxRuntimeParam.Overrides()
func (r SandboxRuntime) ToParam() SandboxRuntimeParam {
	return param.Override[SandboxRuntimeParam](json.RawMessage(r.RawJSON()))
}

// Runtime configuration for Sandbox
type SandboxRuntimeParam struct {
	// The expiration date for the sandbox in ISO 8601 format - 2024-12-31T23:59:59Z
	Expires param.Opt[string] `json:"expires,omitzero"`
	// The Docker image for the sandbox
	Image param.Opt[string] `json:"image,omitzero"`
	// The memory for the sandbox in MB
	Memory param.Opt[int64] `json:"memory,omitzero"`
	// The TTL for the sandbox in seconds - 30m, 24h, 7d
	Ttl param.Opt[string] `json:"ttl,omitzero"`
	// The env variables to set in the sandbox. Should be a list of Kubernetes EnvVar
	// types
	Envs []map[string]any `json:"envs,omitzero"`
	// The exposed ports of the sandbox
	Ports []PortParam `json:"ports,omitzero"`
	paramObj
}

func (r SandboxRuntimeParam) MarshalJSON() (data []byte, err error) {
	type shadow SandboxRuntimeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SandboxRuntimeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Sandbox specification for API
type SandboxSpec struct {
	// Enable or disable the resource
	Enabled bool `json:"enabled"`
	// Lifecycle configuration for the sandbox
	Lifecycle SandboxLifecycle `json:"lifecycle"`
	// Region where the sandbox should be created (e.g. us-pdx-1, eu-lon-1)
	Region string `json:"region"`
	// Runtime configuration for Sandbox
	Runtime SandboxRuntime `json:"runtime"`
	// Volumes to attach to the sandbox
	Volumes []VolumeAttachment `json:"volumes"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Lifecycle   respjson.Field
		Region      respjson.Field
		Runtime     respjson.Field
		Volumes     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SandboxSpec) RawJSON() string { return r.JSON.raw }
func (r *SandboxSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SandboxSpec to a SandboxSpecParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SandboxSpecParam.Overrides()
func (r SandboxSpec) ToParam() SandboxSpecParam {
	return param.Override[SandboxSpecParam](json.RawMessage(r.RawJSON()))
}

// Sandbox specification for API
type SandboxSpecParam struct {
	// Enable or disable the resource
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Region where the sandbox should be created (e.g. us-pdx-1, eu-lon-1)
	Region param.Opt[string] `json:"region,omitzero"`
	// Lifecycle configuration for the sandbox
	Lifecycle SandboxLifecycleParam `json:"lifecycle,omitzero"`
	// Runtime configuration for Sandbox
	Runtime SandboxRuntimeParam `json:"runtime,omitzero"`
	// Volumes to attach to the sandbox
	Volumes []VolumeAttachmentParam `json:"volumes,omitzero"`
	paramObj
}

func (r SandboxSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow SandboxSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SandboxSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Volume attachment configuration for sandbox
type VolumeAttachment struct {
	// Mount path in the container
	MountPath string `json:"mountPath"`
	// Name of the volume to attach
	Name string `json:"name"`
	// Whether the volume is mounted as read-only
	ReadOnly bool `json:"readOnly"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MountPath   respjson.Field
		Name        respjson.Field
		ReadOnly    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VolumeAttachment) RawJSON() string { return r.JSON.raw }
func (r *VolumeAttachment) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this VolumeAttachment to a VolumeAttachmentParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// VolumeAttachmentParam.Overrides()
func (r VolumeAttachment) ToParam() VolumeAttachmentParam {
	return param.Override[VolumeAttachmentParam](json.RawMessage(r.RawJSON()))
}

// Volume attachment configuration for sandbox
type VolumeAttachmentParam struct {
	// Mount path in the container
	MountPath param.Opt[string] `json:"mountPath,omitzero"`
	// Name of the volume to attach
	Name param.Opt[string] `json:"name,omitzero"`
	// Whether the volume is mounted as read-only
	ReadOnly param.Opt[bool] `json:"readOnly,omitzero"`
	paramObj
}

func (r VolumeAttachmentParam) MarshalJSON() (data []byte, err error) {
	type shadow VolumeAttachmentParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VolumeAttachmentParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Sandbox definition for admin store operations
type SandboxGetHubResponse struct {
	// Categories of the defintion
	Categories []map[string]any `json:"categories"`
	// If the definition is coming soon
	ComingSoon bool `json:"coming_soon"`
	// Description of the defintion
	Description string `json:"description"`
	// Display name of the definition
	DisplayName string `json:"displayName"`
	// If the definition is enterprise
	Enterprise bool `json:"enterprise"`
	// If the definition is hidden
	Hidden bool `json:"hidden"`
	// Icon of the definition
	Icon string `json:"icon"`
	// Image of the Sandbox definition
	Image string `json:"image"`
	// Long description of the defintion
	LongDescription string `json:"longDescription"`
	// Memory of the Sandbox definition in MB
	Memory int64 `json:"memory"`
	// Name of the artifact
	Name string `json:"name"`
	// Set of ports for a resource
	Ports []Port `json:"ports"`
	// Tags of the definition
	Tags string `json:"tags"`
	// URL of the definition
	URL string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Categories      respjson.Field
		ComingSoon      respjson.Field
		Description     respjson.Field
		DisplayName     respjson.Field
		Enterprise      respjson.Field
		Hidden          respjson.Field
		Icon            respjson.Field
		Image           respjson.Field
		LongDescription respjson.Field
		Memory          respjson.Field
		Name            respjson.Field
		Ports           respjson.Field
		Tags            respjson.Field
		URL             respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SandboxGetHubResponse) RawJSON() string { return r.JSON.raw }
func (r *SandboxGetHubResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SandboxNewParams struct {
	// Micro VM for running agentic tasks
	Sandbox SandboxParam
	paramObj
}

func (r SandboxNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Sandbox)
}
func (r *SandboxNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Sandbox)
}

type SandboxGetParams struct {
	// Show secret values (admin only)
	ShowSecrets param.Opt[bool] `query:"show_secrets,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SandboxGetParams]'s query parameters as `url.Values`.
func (r SandboxGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SandboxUpdateParams struct {
	// Micro VM for running agentic tasks
	Sandbox SandboxParam
	paramObj
}

func (r SandboxUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Sandbox)
}
func (r *SandboxUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Sandbox)
}
