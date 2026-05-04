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

	"github.com/blaxel-ai/sdk-go/internal/apijson"
	"github.com/blaxel-ai/sdk-go/internal/apiquery"
	shimjson "github.com/blaxel-ai/sdk-go/internal/encoding/json"
	"github.com/blaxel-ai/sdk-go/internal/requestconfig"
	"github.com/blaxel-ai/sdk-go/option"
	"github.com/blaxel-ai/sdk-go/packages/param"
	"github.com/blaxel-ai/sdk-go/packages/respjson"
	"github.com/blaxel-ai/sdk-go/shared"
)

// SandboxService contains methods and other services that help with interacting
// with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSandboxService] method instead.
type SandboxService struct {
	Options    []option.RequestOption
	Process    SandboxProcessService
	Filesystem SandboxFilesystemService
	Codegen    SandboxCodegenService
	Previews   SandboxPreviewService
}

// NewSandboxService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSandboxService(opts ...option.RequestOption) (r SandboxService) {
	r = SandboxService{}
	r.Options = opts
	r.Process = NewSandboxProcessService(opts...)
	r.Filesystem = NewSandboxFilesystemService(opts...)
	r.Codegen = NewSandboxCodegenService(opts...)
	r.Previews = NewSandboxPreviewService(opts...)
	return
}

// Creates a new sandbox VM for secure AI code execution. Sandboxes automatically
// scale to zero when idle and resume instantly, preserving memory state including
// running processes and filesystem.
func (r *SandboxService) New(ctx context.Context, params SandboxNewParams, opts ...option.RequestOption) (res *Sandbox, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "sandboxes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns detailed information about a sandbox including its configuration,
// attached volumes, lifecycle policies, and API endpoint URL.
func (r *SandboxService) Get(ctx context.Context, sandboxName string, query SandboxGetParams, opts ...option.RequestOption) (res *Sandbox, err error) {
	opts = slices.Concat(r.Options, opts)
	if sandboxName == "" {
		err = errors.New("missing required sandboxName parameter")
		return nil, err
	}
	path := fmt.Sprintf("sandboxes/%s", sandboxName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Updates a sandbox's configuration. Note that certain changes (like image or
// memory) may reset the sandbox state. Use lifecycle policies to control automatic
// cleanup.
func (r *SandboxService) Update(ctx context.Context, sandboxName string, body SandboxUpdateParams, opts ...option.RequestOption) (res *Sandbox, err error) {
	opts = slices.Concat(r.Options, opts)
	if sandboxName == "" {
		err = errors.New("missing required sandboxName parameter")
		return nil, err
	}
	path := fmt.Sprintf("sandboxes/%s", sandboxName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Returns all sandboxes in the workspace. Each sandbox includes its configuration,
// status, and endpoint URL. Terminated sandboxes are hidden by default; pass
// `showTerminated=true` to include them.
func (r *SandboxService) List(ctx context.Context, query SandboxListParams, opts ...option.RequestOption) (res *[]Sandbox, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "sandboxes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Permanently deletes a sandbox and all its data. If no volumes are attached, this
// guarantees zero data retention (ZDR). This action cannot be undone.
func (r *SandboxService) Delete(ctx context.Context, sandboxName string, opts ...option.RequestOption) (res *Sandbox, err error) {
	opts = slices.Concat(r.Options, opts)
	if sandboxName == "" {
		err = errors.New("missing required sandboxName parameter")
		return nil, err
	}
	path := fmt.Sprintf("sandboxes/%s", sandboxName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Returns all pre-built sandbox templates available in the Blaxel Hub. These
// include popular development environments with pre-installed tools and
// frameworks.
func (r *SandboxService) GetHub(ctx context.Context, opts ...option.RequestOption) (res *[]SandboxGetHubResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "sandbox/hub"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
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
	// The target port of the port
	Target int64 `json:"target" api:"required"`
	// The name of the port
	Name string `json:"name"`
	// The protocol of the port
	//
	// Any of "HTTP", "TCP", "UDP", "TLS".
	Protocol PortProtocol `json:"protocol"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Target      respjson.Field
		Name        respjson.Field
		Protocol    respjson.Field
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
	PortProtocolTls  PortProtocol = "TLS"
)

// A port for a resource
//
// The property Target is required.
type PortParam struct {
	// The target port of the port
	Target int64 `json:"target" api:"required"`
	// The name of the port
	Name param.Opt[string] `json:"name,omitzero"`
	// The protocol of the port
	//
	// Any of "HTTP", "TCP", "UDP", "TLS".
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
	Metadata Metadata `json:"metadata" api:"required"`
	// Configuration for a sandbox including its image, memory, ports, region, and
	// lifecycle policies
	Spec SandboxSpec `json:"spec" api:"required"`
	// Events happening on a resource deployed on Blaxel
	Events []CoreEvent `json:"events"`
	// Time in seconds until the sandbox is automatically deleted based on TTL and
	// lifecycle policies. Only present for sandboxes with lifecycle configured.
	ExpiresIn int64 `json:"expiresIn"`
	// Last time the sandbox was used (read-only, managed by the system)
	LastUsedAt string `json:"lastUsedAt"`
	// Current state of the sandbox (read-only, managed by the system)
	//
	// Any of "RUNNING", "STANDBY".
	State SandboxState `json:"state"`
	// Deployment status of a resource deployed on Blaxel
	//
	// Any of "DELETING", "TERMINATED", "FAILED", "DEACTIVATED", "DEACTIVATING",
	// "UPLOADING", "BUILDING", "DEPLOYING", "DEPLOYED", "BUILT".
	Status Status `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metadata    respjson.Field
		Spec        respjson.Field
		Events      respjson.Field
		ExpiresIn   respjson.Field
		LastUsedAt  respjson.Field
		State       respjson.Field
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

// Current state of the sandbox (read-only, managed by the system)
type SandboxState string

const (
	SandboxStateRunning SandboxState = "RUNNING"
	SandboxStateStandby SandboxState = "STANDBY"
)

// Lightweight virtual machine for secure AI code execution. Sandboxes resume from
// standby in under 25ms and automatically scale to zero after inactivity,
// preserving memory state including running processes and filesystem.
//
// The properties Metadata, Spec are required.
type SandboxParam struct {
	// Common metadata fields shared by all Blaxel resources including name, labels,
	// timestamps, and ownership information
	Metadata MetadataParam `json:"metadata,omitzero" api:"required"`
	// Configuration for a sandbox including its image, memory, ports, region, and
	// lifecycle policies
	Spec SandboxSpecParam `json:"spec,omitzero" api:"required"`
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
	// Duration to keep the sandbox record after termination for log access (e.g.,
	// '1h', '24h', '7d'). Defaults to 5m. Subject to maximum quota limits.
	TerminatedRetention string `json:"terminatedRetention"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ExpirationPolicies  respjson.Field
		TerminatedRetention respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
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
	// Duration to keep the sandbox record after termination for log access (e.g.,
	// '1h', '24h', '7d'). Defaults to 5m. Subject to maximum quota limits.
	TerminatedRetention param.Opt[string] `json:"terminatedRetention,omitzero"`
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

// Network configuration for a sandbox including domain filtering, egress IP
// binding, and proxy settings
type SandboxNetwork struct {
	// List of allowed external domains (allowlist). When set, only these domains are
	// reachable. Supports wildcards (e.g. \*.s3.amazonaws.com).
	AllowedDomains []string `json:"allowedDomains"`
	// Egress configuration for routing sandbox outbound traffic through a dedicated IP
	// gateway
	Egress SandboxNetworkEgress `json:"egress"`
	// List of forbidden external domains (denylist). When set, all domains except
	// these are reachable. Supports wildcards (e.g. \*.malware.com). If both
	// allowedDomains and forbiddenDomains are set, allowedDomains takes precedence.
	ForbiddenDomains []string `json:"forbiddenDomains"`
	// Proxy configuration for routing sandbox HTTP traffic through the platform proxy
	// with MITM inspection and per-destination header/body injection
	Proxy SandboxNetworkProxy `json:"proxy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AllowedDomains   respjson.Field
		Egress           respjson.Field
		ForbiddenDomains respjson.Field
		Proxy            respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SandboxNetwork) RawJSON() string { return r.JSON.raw }
func (r *SandboxNetwork) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this SandboxNetwork to a SandboxNetworkParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// SandboxNetworkParam.Overrides()
func (r SandboxNetwork) ToParam() SandboxNetworkParam {
	return param.Override[SandboxNetworkParam](json.RawMessage(r.RawJSON()))
}

// Egress configuration for routing sandbox outbound traffic through a dedicated IP
// gateway
type SandboxNetworkEgress struct {
	// Name of the egress gateway to route traffic through. The gateway must exist in
	// the default VPC.
	Gateway string `json:"gateway"`
	// Egress mode. Use 'dedicated' for a dedicated egress IP.
	Mode string `json:"mode"`
	// Per-destination egress policies (not yet supported)
	Policies []SandboxNetworkEgressPolicy `json:"policies"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Gateway     respjson.Field
		Mode        respjson.Field
		Policies    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SandboxNetworkEgress) RawJSON() string { return r.JSON.raw }
func (r *SandboxNetworkEgress) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Egress policy routing specific destinations through dedicated or shared gateways
// (not yet supported)
type SandboxNetworkEgressPolicy struct {
	// Destination domains or IPs this policy applies to
	Destinations []string `json:"destinations"`
	// Egress mode for these destinations (dedicated or shared)
	Mode string `json:"mode"`
	// Name of this egress policy
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Destinations respjson.Field
		Mode         respjson.Field
		Name         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SandboxNetworkEgressPolicy) RawJSON() string { return r.JSON.raw }
func (r *SandboxNetworkEgressPolicy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Proxy configuration for routing sandbox HTTP traffic through the platform proxy
// with MITM inspection and per-destination header/body injection
type SandboxNetworkProxy struct {
	// Domains that bypass the proxy entirely via the NO_PROXY directive. Traffic to
	// these destinations goes direct, not through the CONNECT tunnel. Supports
	// wildcards. Note that localhost, private ranges (10.0.0.0/8, 172.16.0.0/12,
	// 192.168.0.0/16), 169.254.169.254, .local and .internal are always bypassed by
	// default.
	Bypass []string `json:"bypass"`
	// Per-destination routing rules with header/body injection and secrets. Use
	// destinations ["*"] for global rules that apply to all destinations.
	Routing []SandboxNetworkProxyRouting `json:"routing"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bypass      respjson.Field
		Routing     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SandboxNetworkProxy) RawJSON() string { return r.JSON.raw }
func (r *SandboxNetworkProxy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Routing rule that injects headers and body fields into requests matching the
// given destinations. Use destinations ["*"] for a global rule that applies to all
// proxied traffic.
type SandboxNetworkProxyRouting struct {
	// Body fields to inject into matching requests. Values may contain {{SECRET:name}}
	// references resolved from this rule's secrets.
	Body map[string]string `json:"body"`
	// Destination domains this rule applies to. Use ["*"] for a global rule that
	// matches all destinations.
	Destinations []string `json:"destinations"`
	// Headers to inject into matching requests. Values may contain {{SECRET:name}}
	// references resolved from this rule's secrets.
	Headers map[string]string `json:"headers"`
	// Named secret values for this routing rule, referenced in headers/body via
	// {{SECRET:name}}. Stored encrypted at rest. Write-only: never returned in API
	// responses.
	Secrets map[string]string `json:"secrets"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Body         respjson.Field
		Destinations respjson.Field
		Headers      respjson.Field
		Secrets      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SandboxNetworkProxyRouting) RawJSON() string { return r.JSON.raw }
func (r *SandboxNetworkProxyRouting) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Network configuration for a sandbox including domain filtering, egress IP
// binding, and proxy settings
type SandboxNetworkParam struct {
	// List of allowed external domains (allowlist). When set, only these domains are
	// reachable. Supports wildcards (e.g. \*.s3.amazonaws.com).
	AllowedDomains []string `json:"allowedDomains,omitzero"`
	// Egress configuration for routing sandbox outbound traffic through a dedicated IP
	// gateway
	Egress SandboxNetworkEgressParam `json:"egress,omitzero"`
	// List of forbidden external domains (denylist). When set, all domains except
	// these are reachable. Supports wildcards (e.g. \*.malware.com). If both
	// allowedDomains and forbiddenDomains are set, allowedDomains takes precedence.
	ForbiddenDomains []string `json:"forbiddenDomains,omitzero"`
	// Proxy configuration for routing sandbox HTTP traffic through the platform proxy
	// with MITM inspection and per-destination header/body injection
	Proxy SandboxNetworkProxyParam `json:"proxy,omitzero"`
	paramObj
}

func (r SandboxNetworkParam) MarshalJSON() (data []byte, err error) {
	type shadow SandboxNetworkParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SandboxNetworkParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Egress configuration for routing sandbox outbound traffic through a dedicated IP
// gateway
type SandboxNetworkEgressParam struct {
	// Name of the egress gateway to route traffic through. The gateway must exist in
	// the default VPC.
	Gateway param.Opt[string] `json:"gateway,omitzero"`
	// Egress mode. Use 'dedicated' for a dedicated egress IP.
	Mode param.Opt[string] `json:"mode,omitzero"`
	// Per-destination egress policies (not yet supported)
	Policies []SandboxNetworkEgressPolicyParam `json:"policies,omitzero"`
	paramObj
}

func (r SandboxNetworkEgressParam) MarshalJSON() (data []byte, err error) {
	type shadow SandboxNetworkEgressParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SandboxNetworkEgressParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Egress policy routing specific destinations through dedicated or shared gateways
// (not yet supported)
type SandboxNetworkEgressPolicyParam struct {
	// Egress mode for these destinations (dedicated or shared)
	Mode param.Opt[string] `json:"mode,omitzero"`
	// Name of this egress policy
	Name param.Opt[string] `json:"name,omitzero"`
	// Destination domains or IPs this policy applies to
	Destinations []string `json:"destinations,omitzero"`
	paramObj
}

func (r SandboxNetworkEgressPolicyParam) MarshalJSON() (data []byte, err error) {
	type shadow SandboxNetworkEgressPolicyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SandboxNetworkEgressPolicyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Proxy configuration for routing sandbox HTTP traffic through the platform proxy
// with MITM inspection and per-destination header/body injection
type SandboxNetworkProxyParam struct {
	// Domains that bypass the proxy entirely via the NO_PROXY directive. Traffic to
	// these destinations goes direct, not through the CONNECT tunnel. Supports
	// wildcards. Note that localhost, private ranges (10.0.0.0/8, 172.16.0.0/12,
	// 192.168.0.0/16), 169.254.169.254, .local and .internal are always bypassed by
	// default.
	Bypass []string `json:"bypass,omitzero"`
	// Per-destination routing rules with header/body injection and secrets. Use
	// destinations ["*"] for global rules that apply to all destinations.
	Routing []SandboxNetworkProxyRoutingParam `json:"routing,omitzero"`
	paramObj
}

func (r SandboxNetworkProxyParam) MarshalJSON() (data []byte, err error) {
	type shadow SandboxNetworkProxyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SandboxNetworkProxyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Routing rule that injects headers and body fields into requests matching the
// given destinations. Use destinations ["*"] for a global rule that applies to all
// proxied traffic.
type SandboxNetworkProxyRoutingParam struct {
	// Body fields to inject into matching requests. Values may contain {{SECRET:name}}
	// references resolved from this rule's secrets.
	Body map[string]string `json:"body,omitzero"`
	// Destination domains this rule applies to. Use ["*"] for a global rule that
	// matches all destinations.
	Destinations []string `json:"destinations,omitzero"`
	// Headers to inject into matching requests. Values may contain {{SECRET:name}}
	// references resolved from this rule's secrets.
	Headers map[string]string `json:"headers,omitzero"`
	// Named secret values for this routing rule, referenced in headers/body via
	// {{SECRET:name}}. Stored encrypted at rest. Write-only: never returned in API
	// responses.
	Secrets map[string]string `json:"secrets,omitzero"`
	paramObj
}

func (r SandboxNetworkProxyRoutingParam) MarshalJSON() (data []byte, err error) {
	type shadow SandboxNetworkProxyRoutingParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SandboxNetworkProxyRoutingParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Runtime configuration defining how the sandbox VM is provisioned and its
// resource limits
type SandboxRuntime struct {
	// Environment variables injected into the sandbox. Supports Kubernetes EnvVar
	// format with valueFrom references.
	Envs []shared.Env `json:"envs"`
	// Absolute expiration timestamp in ISO 8601 format when the sandbox will be
	// deleted
	Expires string `json:"expires"`
	// Extra arguments for sandbox kernel selection. Supported keys: 'iptables',
	// 'nvme'. Values: 'enabled' or 'disabled'. Determines which kernel variant the
	// sandbox runs on. Immutable after creation.
	ExtraArgs map[string]string `json:"extraArgs"`
	// Sandbox image to use. Can be a public Blaxel image (e.g.,
	// blaxel/base-image:latest) or a custom template image built with 'bl deploy'.
	Image string `json:"image"`
	// Memory allocation in megabytes. Also determines CPU allocation (CPU cores =
	// memory in MB / 2048, e.g., 4096MB = 2 CPUs).
	Memory int64 `json:"memory"`
	// Set of ports for a resource
	Ports []Port `json:"ports"`
	// Duration in seconds the pod needs to terminate gracefully. Defaults to 0 for
	// immediate termination.
	TerminationGracePeriodSeconds int64 `json:"terminationGracePeriodSeconds"`
	// Time-to-live duration after which the sandbox is automatically deleted (e.g.,
	// '30m', '24h', '7d')
	Ttl string `json:"ttl"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Envs                          respjson.Field
		Expires                       respjson.Field
		ExtraArgs                     respjson.Field
		Image                         respjson.Field
		Memory                        respjson.Field
		Ports                         respjson.Field
		TerminationGracePeriodSeconds respjson.Field
		Ttl                           respjson.Field
		ExtraFields                   map[string]respjson.Field
		raw                           string
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
	// Duration in seconds the pod needs to terminate gracefully. Defaults to 0 for
	// immediate termination.
	TerminationGracePeriodSeconds param.Opt[int64] `json:"terminationGracePeriodSeconds,omitzero"`
	// Time-to-live duration after which the sandbox is automatically deleted (e.g.,
	// '30m', '24h', '7d')
	Ttl param.Opt[string] `json:"ttl,omitzero"`
	// Environment variables injected into the sandbox. Supports Kubernetes EnvVar
	// format with valueFrom references.
	Envs []shared.EnvParam `json:"envs,omitzero"`
	// Extra arguments for sandbox kernel selection. Supported keys: 'iptables',
	// 'nvme'. Values: 'enabled' or 'disabled'. Determines which kernel variant the
	// sandbox runs on. Immutable after creation.
	ExtraArgs map[string]string `json:"extraArgs,omitzero"`
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
	// Network configuration for a sandbox including domain filtering, egress IP
	// binding, and proxy settings
	Network SandboxNetwork `json:"network"`
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
		Network     respjson.Field
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
	// Network configuration for a sandbox including domain filtering, egress IP
	// binding, and proxy settings
	Network SandboxNetworkParam `json:"network,omitzero"`
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

// Configuration for attaching a persistent volume to a sandbox at a specific
// filesystem path
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

// Configuration for attaching a persistent volume to a sandbox at a specific
// filesystem path
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
	// If true, return existing sandbox instead of 409 error when sandbox exists and is
	// not in FAILED/TERMINATED/TERMINATING state
	CreateIfNotExist param.Opt[bool] `query:"createIfNotExist,omitzero" json:"-"`
	paramObj
}

func (r SandboxNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Sandbox)
}
func (r *SandboxNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// URLQuery serializes [SandboxNewParams]'s query parameters as `url.Values`.
func (r SandboxNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
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
	return apijson.UnmarshalRoot(data, r)
}

type SandboxListParams struct {
	// If true, include terminated sandboxes in the response. Defaults to false.
	ShowTerminated param.Opt[bool] `query:"showTerminated,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SandboxListParams]'s query parameters as `url.Values`.
func (r SandboxListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
