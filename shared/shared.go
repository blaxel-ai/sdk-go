// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package shared

import (
	"encoding/json"

	"github.com/blaxel-ai/sdk-go/internal/apijson"
	"github.com/blaxel-ai/sdk-go/packages/param"
	"github.com/blaxel-ai/sdk-go/packages/respjson"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type EgressGatewayMetadata struct {
	// Unique identifier for the egress gateway within the VPC. Must be lowercase
	// alphanumeric with hyphens, max 49 characters. Immutable after creation.
	Name string `json:"name" api:"required"`
	// The date and time when the resource was created
	CreatedAt string `json:"createdAt"`
	// The user or service account who created the resource
	CreatedBy string `json:"createdBy"`
	// Human-readable name for display in the UI. Can contain spaces and special
	// characters, max 63 characters.
	DisplayName string `json:"displayName"`
	// The date and time when the resource was updated
	UpdatedAt string `json:"updatedAt"`
	// The user or service account who updated the resource
	UpdatedBy string `json:"updatedBy"`
	// Name of the VPC this egress gateway belongs to
	VpcName string `json:"vpcName"`
	// Name of the workspace this resource belongs to (read-only, set automatically)
	Workspace string `json:"workspace"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		CreatedAt   respjson.Field
		CreatedBy   respjson.Field
		DisplayName respjson.Field
		UpdatedAt   respjson.Field
		UpdatedBy   respjson.Field
		VpcName     respjson.Field
		Workspace   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EgressGatewayMetadata) RawJSON() string { return r.JSON.raw }
func (r *EgressGatewayMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EgressGatewayMetadata to a EgressGatewayMetadataParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EgressGatewayMetadataParam.Overrides()
func (r EgressGatewayMetadata) ToParam() EgressGatewayMetadataParam {
	return param.Override[EgressGatewayMetadataParam](json.RawMessage(r.RawJSON()))
}

// The property Name is required.
type EgressGatewayMetadataParam struct {
	// Unique identifier for the egress gateway within the VPC. Must be lowercase
	// alphanumeric with hyphens, max 49 characters. Immutable after creation.
	Name string `json:"name" api:"required"`
	// Human-readable name for display in the UI. Can contain spaces and special
	// characters, max 63 characters.
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	paramObj
}

func (r EgressGatewayMetadataParam) MarshalJSON() (data []byte, err error) {
	type shadow EgressGatewayMetadataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EgressGatewayMetadataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specification for an egress gateway including region and capacity configuration
type EgressGatewaySpec struct {
	// Region where the egress gateway is provisioned (e.g. us-pdx-1, eu-lon-1)
	Region string `json:"region" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Region      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EgressGatewaySpec) RawJSON() string { return r.JSON.raw }
func (r *EgressGatewaySpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EgressGatewaySpec to a EgressGatewaySpecParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EgressGatewaySpecParam.Overrides()
func (r EgressGatewaySpec) ToParam() EgressGatewaySpecParam {
	return param.Override[EgressGatewaySpecParam](json.RawMessage(r.RawJSON()))
}

// Specification for an egress gateway including region and capacity configuration
//
// The property Region is required.
type EgressGatewaySpecParam struct {
	// Region where the egress gateway is provisioned (e.g. us-pdx-1, eu-lon-1)
	Region string `json:"region" api:"required"`
	paramObj
}

func (r EgressGatewaySpecParam) MarshalJSON() (data []byte, err error) {
	type shadow EgressGatewaySpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EgressGatewaySpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type EgressIPMetadata struct {
	// Unique identifier for the egress IP within the gateway. Must be lowercase
	// alphanumeric with hyphens, max 49 characters. Immutable after creation.
	Name string `json:"name" api:"required"`
	// The date and time when the resource was created
	CreatedAt string `json:"createdAt"`
	// The user or service account who created the resource
	CreatedBy string `json:"createdBy"`
	// Human-readable name for display in the UI. Can contain spaces and special
	// characters, max 63 characters.
	DisplayName string `json:"displayName"`
	// Name of the egress gateway this IP is allocated from
	GatewayName string `json:"gatewayName"`
	// The date and time when the resource was updated
	UpdatedAt string `json:"updatedAt"`
	// The user or service account who updated the resource
	UpdatedBy string `json:"updatedBy"`
	// Name of the VPC this egress IP belongs to
	VpcName string `json:"vpcName"`
	// Name of the workspace this resource belongs to (read-only, set automatically)
	Workspace string `json:"workspace"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		CreatedAt   respjson.Field
		CreatedBy   respjson.Field
		DisplayName respjson.Field
		GatewayName respjson.Field
		UpdatedAt   respjson.Field
		UpdatedBy   respjson.Field
		VpcName     respjson.Field
		Workspace   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EgressIPMetadata) RawJSON() string { return r.JSON.raw }
func (r *EgressIPMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EgressIPMetadata to a EgressIPMetadataParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EgressIPMetadataParam.Overrides()
func (r EgressIPMetadata) ToParam() EgressIPMetadataParam {
	return param.Override[EgressIPMetadataParam](json.RawMessage(r.RawJSON()))
}

// The property Name is required.
type EgressIPMetadataParam struct {
	// Unique identifier for the egress IP within the gateway. Must be lowercase
	// alphanumeric with hyphens, max 49 characters. Immutable after creation.
	Name string `json:"name" api:"required"`
	// Human-readable name for display in the UI. Can contain spaces and special
	// characters, max 63 characters.
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	paramObj
}

func (r EgressIPMetadataParam) MarshalJSON() (data []byte, err error) {
	type shadow EgressIPMetadataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EgressIPMetadataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specification for an egress IP including IP family and the provisioned address
type EgressIPSpec struct {
	// IP address family, either IPv4 or IPv6
	//
	// Any of "IPv4", "IPv6".
	IPFamily EgressIPSpecIPFamily `json:"ipFamily" api:"required"`
	// Public IP address assigned to this egress IP (read-only, set after provisioning)
	PublicIP string `json:"publicIp"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IPFamily    respjson.Field
		PublicIP    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EgressIPSpec) RawJSON() string { return r.JSON.raw }
func (r *EgressIPSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this EgressIPSpec to a EgressIPSpecParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EgressIPSpecParam.Overrides()
func (r EgressIPSpec) ToParam() EgressIPSpecParam {
	return param.Override[EgressIPSpecParam](json.RawMessage(r.RawJSON()))
}

// IP address family, either IPv4 or IPv6
type EgressIPSpecIPFamily string

const (
	EgressIPSpecIPFamilyIPv4 EgressIPSpecIPFamily = "IPv4"
	EgressIPSpecIPFamilyIPv6 EgressIPSpecIPFamily = "IPv6"
)

// Specification for an egress IP including IP family and the provisioned address
//
// The property IPFamily is required.
type EgressIPSpecParam struct {
	// IP address family, either IPv4 or IPv6
	//
	// Any of "IPv4", "IPv6".
	IPFamily EgressIPSpecIPFamily `json:"ipFamily,omitzero" api:"required"`
	paramObj
}

func (r EgressIPSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow EgressIPSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EgressIPSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Environment variable with name and value
type Env struct {
	// Name of the environment variable
	Name string `json:"name"`
	// Whether the value is a secret
	Secret bool `json:"secret"`
	// Value of the environment variable
	Value string `json:"value"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Secret      respjson.Field
		Value       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Env) RawJSON() string { return r.JSON.raw }
func (r *Env) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Env to a EnvParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// EnvParam.Overrides()
func (r Env) ToParam() EnvParam {
	return param.Override[EnvParam](json.RawMessage(r.RawJSON()))
}

// Environment variable with name and value
type EnvParam struct {
	// Name of the environment variable
	Name param.Opt[string] `json:"name,omitzero"`
	// Whether the value is a secret
	Secret param.Opt[bool] `json:"secret,omitzero"`
	// Value of the environment variable
	Value param.Opt[string] `json:"value,omitzero"`
	paramObj
}

func (r EnvParam) MarshalJSON() (data []byte, err error) {
	type shadow EnvParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *EnvParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A type of hardware available for deployments
type Flavor struct {
	// Flavor name (e.g. t4)
	Name string `json:"name"`
	// Flavor type (e.g. cpu, gpu)
	//
	// Any of "cpu", "gpu".
	Type FlavorType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Flavor) RawJSON() string { return r.JSON.raw }
func (r *Flavor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Flavor to a FlavorParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// FlavorParam.Overrides()
func (r Flavor) ToParam() FlavorParam {
	return param.Override[FlavorParam](json.RawMessage(r.RawJSON()))
}

// Flavor type (e.g. cpu, gpu)
type FlavorType string

const (
	FlavorTypeCPU FlavorType = "cpu"
	FlavorTypeGPU FlavorType = "gpu"
)

// A type of hardware available for deployments
type FlavorParam struct {
	// Flavor name (e.g. t4)
	Name param.Opt[string] `json:"name,omitzero"`
	// Flavor type (e.g. cpu, gpu)
	//
	// Any of "cpu", "gpu".
	Type FlavorType `json:"type,omitzero"`
	paramObj
}

func (r FlavorParam) MarshalJSON() (data []byte, err error) {
	type shadow FlavorParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FlavorParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
