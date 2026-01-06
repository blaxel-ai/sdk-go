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

// Creates a new sandbox VM for secure AI code execution. Sandboxes automatically
// scale to zero when idle and resume instantly, preserving memory state including
// running processes and filesystem.
func (r *SandboxService) New(ctx context.Context, body SandboxNewParams, opts ...option.RequestOption) (res *Sandbox, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "sandboxes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns detailed information about a sandbox including its current state
// (active/standby), configuration, attached volumes, lifecycle policies, and API
// endpoint URL.
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

// Updates a sandbox's configuration. Note that certain changes (like image or
// memory) may reset the sandbox state. Use lifecycle policies to control automatic
// cleanup.
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

// Returns all sandboxes in the workspace. Each sandbox includes its current state
// (active/standby), configuration, and endpoint URL. Sandboxes resume from standby
// in under 25ms.
func (r *SandboxService) List(ctx context.Context, opts ...option.RequestOption) (res *[]Sandbox, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "sandboxes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Permanently deletes a sandbox and all its data. If no volumes are attached, this
// guarantees zero data retention (ZDR). This action cannot be undone.
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

// Returns all pre-built sandbox templates available in the Blaxel Hub. These
// include popular development environments with pre-installed tools and
// frameworks.
func (r *SandboxService) GetHub(ctx context.Context, opts ...option.RequestOption) (res *[]SandboxGetHubResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "sandbox/hub"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Expiration policy for automatic sandbox cleanup based on time conditions
type ExpirationPolicy struct {
	// Action to take when the expiration condition is met
	//
	// Any of "delete".
	Action ExpirationPolicyAction `json:"action"`
	// Type of expiration policy: ttl-idle (delete after inactivity), ttl-max-age
	// (delete after total lifetime), or date (delete at specific time)
	//
	// Any of "ttl-idle", "ttl-max-age", "date".
	Type ExpirationPolicyType `json:"type"`
	// Duration value for TTL policies (e.g., '30m', '24h', '7d') or ISO 8601 date for
	// date policies
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

// Action to take when the expiration condition is met
type ExpirationPolicyAction string

const (
	ExpirationPolicyActionDelete ExpirationPolicyAction = "delete"
)

// Type of expiration policy: ttl-idle (delete after inactivity), ttl-max-age
// (delete after total lifetime), or date (delete at specific time)
type ExpirationPolicyType string

const (
	ExpirationPolicyTypeTtlIdle   ExpirationPolicyType = "ttl-idle"
	ExpirationPolicyTypeTtlMaxAge ExpirationPolicyType = "ttl-max-age"
	ExpirationPolicyTypeDate      ExpirationPolicyType = "date"
)

// Expiration policy for automatic sandbox cleanup based on time conditions
type ExpirationPolicyParam struct {
	// Duration value for TTL policies (e.g., '30m', '24h', '7d') or ISO 8601 date for
	// date policies
	Value param.Opt[string] `json:"value,omitzero"`
	// Action to take when the expiration condition is met
	//
	// Any of "delete".
	Action ExpirationPolicyAction `json:"action,omitzero"`
	// Type of expiration policy: ttl-idle (delete after inactivity), ttl-max-age
	// (delete after total lifetime), or date (delete at specific time)
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

// Lightweight virtual machine for secure AI code execution. Sandboxes resume from
// standby in under 25ms and automatically scale to zero after inactivity,
// preserving memory state including running processes and filesystem.
type Sandbox struct {
	// Common metadata fields shared by all Blaxel resources including name, labels,
	// timestamps, and ownership information
	Metadata Metadata `json:"metadata,required"`
	// Configuration for a sandbox including its image, memory, ports, region, and
	// lifecycle policies
	Spec SandboxSpec `json:"spec,required"`
	// Events happening on a resource deployed on Blaxel
	Events []CoreEvent `json:"events"`
	// Last time the sandbox was used (read-only, managed by the system)
	LastUsedAt string `json:"lastUsedAt"`
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

// Lightweight virtual machine for secure AI code execution. Sandboxes resume from
// standby in under 25ms and automatically scale to zero after inactivity,
// preserving memory state including running processes and filesystem.
//
// The properties Metadata, Spec are required.
type SandboxParam struct {
	// Common metadata fields shared by all Blaxel resources including name, labels,
	// timestamps, and ownership information
	Metadata MetadataParam `json:"metadata,omitzero,required"`
	// Configuration for a sandbox including its image, memory, ports, region, and
	// lifecycle policies
	Spec SandboxSpecParam `json:"spec,omitzero,required"`
	// Events happening on a resource deployed on Blaxel
	Events []CoreEventParam `json:"events,omitzero"`
	// Deployment status of a resource deployed on Blaxel
	//
	// Any of "DELETING", "TERMINATED", "FAILED", "DEACTIVATED", "DEACTIVATING",
	// "UPLOADING", "BUILDING", "DEPLOYING", "DEPLOYED".
	Status Status `json:"status,omitzero"`
	paramObj
}

func (r SandboxParam) MarshalJSON() (data []byte, err error) {
	type shadow SandboxParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SandboxParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Lifecycle configuration controlling automatic sandbox deletion based on idle
// time, max age, or specific dates
type SandboxLifecycle struct {
	// List of expiration policies. Multiple policies can be combined; whichever
	// condition is met first triggers the action.
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

// Lifecycle configuration controlling automatic sandbox deletion based on idle
// time, max age, or specific dates
type SandboxLifecycleParam struct {
	// List of expiration policies. Multiple policies can be combined; whichever
	// condition is met first triggers the action.
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

// Runtime configuration defining how the sandbox VM is provisioned and its
// resource limits
type SandboxRuntime struct {
	// Environment variables injected into the sandbox. Supports Kubernetes EnvVar
	// format with valueFrom references.
	Envs []map[string]any `json:"envs"`
	// Absolute expiration timestamp in ISO 8601 format when the sandbox will be
	// deleted
	Expires string `json:"expires"`
	// Sandbox image to use. Can be a public Blaxel image (e.g.,
	// blaxel/base-image:latest) or a custom template image built with 'bl deploy'.
	Image string `json:"image"`
	// Memory allocation in megabytes. Also determines CPU allocation (CPU cores =
	// memory in MB / 2048, e.g., 4096MB = 2 CPUs).
	Memory int64 `json:"memory"`
	// Set of ports for a resource
	Ports []Port `json:"ports"`
	// Time-to-live duration after which the sandbox is automatically deleted (e.g.,
	// '30m', '24h', '7d')
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

// Runtime configuration defining how the sandbox VM is provisioned and its
// resource limits
type SandboxRuntimeParam struct {
	// Absolute expiration timestamp in ISO 8601 format when the sandbox will be
	// deleted
	Expires param.Opt[string] `json:"expires,omitzero"`
	// Sandbox image to use. Can be a public Blaxel image (e.g.,
	// blaxel/base-image:latest) or a custom template image built with 'bl deploy'.
	Image param.Opt[string] `json:"image,omitzero"`
	// Memory allocation in megabytes. Also determines CPU allocation (CPU cores =
	// memory in MB / 2048, e.g., 4096MB = 2 CPUs).
	Memory param.Opt[int64] `json:"memory,omitzero"`
	// Time-to-live duration after which the sandbox is automatically deleted (e.g.,
	// '30m', '24h', '7d')
	Ttl param.Opt[string] `json:"ttl,omitzero"`
	// Environment variables injected into the sandbox. Supports Kubernetes EnvVar
	// format with valueFrom references.
	Envs []map[string]any `json:"envs,omitzero"`
	// Set of ports for a resource
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

// Configuration for a sandbox including its image, memory, ports, region, and
// lifecycle policies
type SandboxSpec struct {
	// When false, the sandbox is disabled and will not accept connections
	Enabled bool `json:"enabled"`
	// Lifecycle configuration controlling automatic sandbox deletion based on idle
	// time, max age, or specific dates
	Lifecycle SandboxLifecycle `json:"lifecycle"`
	// Region where the sandbox should be created (e.g. us-pdx-1, eu-lon-1). If not
	// specified, defaults to the region closest to the user.
	Region string `json:"region"`
	// Runtime configuration defining how the sandbox VM is provisioned and its
	// resource limits
	Runtime SandboxRuntime     `json:"runtime"`
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

// Configuration for a sandbox including its image, memory, ports, region, and
// lifecycle policies
type SandboxSpecParam struct {
	// When false, the sandbox is disabled and will not accept connections
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// Region where the sandbox should be created (e.g. us-pdx-1, eu-lon-1). If not
	// specified, defaults to the region closest to the user.
	Region param.Opt[string] `json:"region,omitzero"`
	// Lifecycle configuration controlling automatic sandbox deletion based on idle
	// time, max age, or specific dates
	Lifecycle SandboxLifecycleParam `json:"lifecycle,omitzero"`
	// Runtime configuration defining how the sandbox VM is provisioned and its
	// resource limits
	Runtime SandboxRuntimeParam     `json:"runtime,omitzero"`
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

// Configuration for attaching a volume to a sandbox at a specific filesystem path
type VolumeAttachment struct {
	// Absolute filesystem path where the volume will be mounted inside the sandbox
	MountPath string `json:"mountPath"`
	// Name of the volume resource to attach (must exist in the same workspace and
	// region)
	Name string `json:"name"`
	// If true, the volume is mounted read-only and cannot be modified by the sandbox
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

// Configuration for attaching a volume to a sandbox at a specific filesystem path
type VolumeAttachmentParam struct {
	// Absolute filesystem path where the volume will be mounted inside the sandbox
	MountPath param.Opt[string] `json:"mountPath,omitzero"`
	// Name of the volume resource to attach (must exist in the same workspace and
	// region)
	Name param.Opt[string] `json:"name,omitzero"`
	// If true, the volume is mounted read-only and cannot be modified by the sandbox
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

// Pre-configured sandbox template available in the Sandbox Hub for quick
// deployment with predefined tools and configurations
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
	// Lightweight virtual machine for secure AI code execution. Sandboxes resume from
	// standby in under 25ms and automatically scale to zero after inactivity,
	// preserving memory state including running processes and filesystem.
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
	// Show secret values (requires workspace admin role)
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
	// Lightweight virtual machine for secure AI code execution. Sandboxes resume from
	// standby in under 25ms and automatically scale to zero after inactivity,
	// preserving memory state including running processes and filesystem.
	Sandbox SandboxParam
	paramObj
}

func (r SandboxUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Sandbox)
}
func (r *SandboxUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Sandbox)
}
