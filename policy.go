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

// PolicyService contains methods and other services that help with interacting
// with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPolicyService] method instead.
type PolicyService struct {
	Options []option.RequestOption
}

// NewPolicyService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewPolicyService(opts ...option.RequestOption) (r PolicyService) {
	r = PolicyService{}
	r.Options = opts
	return
}

// Creates a new governance policy to control where and how resources are deployed.
// Policies can restrict deployment to specific regions, countries, or continents
// for compliance.
func (r *PolicyService) New(ctx context.Context, body PolicyNewParams, opts ...option.RequestOption) (res *Policy, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "policies"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns detailed information about a governance policy including its type
// (location, flavor, or maxToken), restrictions, and which resource types it
// applies to.
func (r *PolicyService) Get(ctx context.Context, policyName string, opts ...option.RequestOption) (res *Policy, err error) {
	opts = slices.Concat(r.Options, opts)
	if policyName == "" {
		err = errors.New("missing required policyName parameter")
		return
	}
	path := fmt.Sprintf("policies/%s", policyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a governance policy's restrictions. Changes take effect on the next
// deployment of resources using this policy.
func (r *PolicyService) Update(ctx context.Context, policyName string, body PolicyUpdateParams, opts ...option.RequestOption) (res *Policy, err error) {
	opts = slices.Concat(r.Options, opts)
	if policyName == "" {
		err = errors.New("missing required policyName parameter")
		return
	}
	path := fmt.Sprintf("policies/%s", policyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Returns all governance policies in the workspace. Policies control deployment
// locations, hardware flavors, and token limits for agents, functions, and models.
func (r *PolicyService) List(ctx context.Context, opts ...option.RequestOption) (res *[]Policy, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "policies"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Permanently deletes a governance policy. Resources using this policy will need
// to be updated to use a different policy.
func (r *PolicyService) Delete(ctx context.Context, policyName string, opts ...option.RequestOption) (res *Policy, err error) {
	opts = slices.Concat(r.Options, opts)
	if policyName == "" {
		err = errors.New("missing required policyName parameter")
		return
	}
	path := fmt.Sprintf("policies/%s", policyName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Rule that controls how a deployment is made and served (e.g. location
// restrictions)
type Policy struct {
	// Common metadata fields shared by all Blaxel resources including name, labels,
	// timestamps, and ownership information
	Metadata Metadata `json:"metadata" api:"required"`
	// Policy specification
	Spec PolicySpec `json:"spec" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metadata    respjson.Field
		Spec        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Policy) RawJSON() string { return r.JSON.raw }
func (r *Policy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Policy to a PolicyParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PolicyParam.Overrides()
func (r Policy) ToParam() PolicyParam {
	return param.Override[PolicyParam](json.RawMessage(r.RawJSON()))
}

// Rule that controls how a deployment is made and served (e.g. location
// restrictions)
//
// The properties Metadata, Spec are required.
type PolicyParam struct {
	// Common metadata fields shared by all Blaxel resources including name, labels,
	// timestamps, and ownership information
	Metadata MetadataParam `json:"metadata,omitzero" api:"required"`
	// Policy specification
	Spec PolicySpecParam `json:"spec,omitzero" api:"required"`
	paramObj
}

func (r PolicyParam) MarshalJSON() (data []byte, err error) {
	type shadow PolicyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Policy location
type PolicyLocation struct {
	// Policy location name
	Name string `json:"name"`
	// Policy location type
	//
	// Any of "location", "country", "continent".
	Type PolicyLocationType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PolicyLocation) RawJSON() string { return r.JSON.raw }
func (r *PolicyLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PolicyLocation to a PolicyLocationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PolicyLocationParam.Overrides()
func (r PolicyLocation) ToParam() PolicyLocationParam {
	return param.Override[PolicyLocationParam](json.RawMessage(r.RawJSON()))
}

// Policy location type
type PolicyLocationType string

const (
	PolicyLocationTypeLocation  PolicyLocationType = "location"
	PolicyLocationTypeCountry   PolicyLocationType = "country"
	PolicyLocationTypeContinent PolicyLocationType = "continent"
)

// Policy location
type PolicyLocationParam struct {
	// Policy location name
	Name param.Opt[string] `json:"name,omitzero"`
	// Policy location type
	//
	// Any of "location", "country", "continent".
	Type PolicyLocationType `json:"type,omitzero"`
	paramObj
}

func (r PolicyLocationParam) MarshalJSON() (data []byte, err error) {
	type shadow PolicyLocationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyLocationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// PolicyMaxTokens is a local type that wraps a slice of PolicyMaxTokens
type PolicyMaxTokens struct {
	// Granularity
	Granularity string `json:"granularity"`
	// Input
	Input int64 `json:"input"`
	// Output
	Output int64 `json:"output"`
	// RatioInputOverOutput
	RatioInputOverOutput int64 `json:"ratioInputOverOutput"`
	// Step
	Step int64 `json:"step"`
	// Total
	Total int64 `json:"total"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Granularity          respjson.Field
		Input                respjson.Field
		Output               respjson.Field
		RatioInputOverOutput respjson.Field
		Step                 respjson.Field
		Total                respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PolicyMaxTokens) RawJSON() string { return r.JSON.raw }
func (r *PolicyMaxTokens) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PolicyMaxTokens to a PolicyMaxTokensParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PolicyMaxTokensParam.Overrides()
func (r PolicyMaxTokens) ToParam() PolicyMaxTokensParam {
	return param.Override[PolicyMaxTokensParam](json.RawMessage(r.RawJSON()))
}

// PolicyMaxTokens is a local type that wraps a slice of PolicyMaxTokens
type PolicyMaxTokensParam struct {
	// Granularity
	Granularity param.Opt[string] `json:"granularity,omitzero"`
	// Input
	Input param.Opt[int64] `json:"input,omitzero"`
	// Output
	Output param.Opt[int64] `json:"output,omitzero"`
	// RatioInputOverOutput
	RatioInputOverOutput param.Opt[int64] `json:"ratioInputOverOutput,omitzero"`
	// Step
	Step param.Opt[int64] `json:"step,omitzero"`
	// Total
	Total param.Opt[int64] `json:"total,omitzero"`
	paramObj
}

func (r PolicyMaxTokensParam) MarshalJSON() (data []byte, err error) {
	type shadow PolicyMaxTokensParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyMaxTokensParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Policy specification
type PolicySpec struct {
	// Types of hardware available for deployments
	Flavors []PolicySpecFlavor `json:"flavors"`
	// PolicyLocations is a local type that wraps a slice of Location
	Locations []PolicyLocation `json:"locations"`
	// PolicyMaxTokens is a local type that wraps a slice of PolicyMaxTokens
	MaxTokens PolicyMaxTokens `json:"maxTokens"`
	// PolicyResourceTypes is a local type that wraps a slice of PolicyResourceType
	//
	// Any of "model", "function", "agent", "sandbox".
	ResourceTypes []string `json:"resourceTypes"`
	// Sandbox mode
	Sandbox bool `json:"sandbox"`
	// Policy type, can be location or flavor
	//
	// Any of "location", "flavor", "maxToken".
	Type PolicySpecType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Flavors       respjson.Field
		Locations     respjson.Field
		MaxTokens     respjson.Field
		ResourceTypes respjson.Field
		Sandbox       respjson.Field
		Type          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PolicySpec) RawJSON() string { return r.JSON.raw }
func (r *PolicySpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this PolicySpec to a PolicySpecParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// PolicySpecParam.Overrides()
func (r PolicySpec) ToParam() PolicySpecParam {
	return param.Override[PolicySpecParam](json.RawMessage(r.RawJSON()))
}

// A type of hardware available for deployments
type PolicySpecFlavor struct {
	// Flavor name (e.g. t4)
	Name string `json:"name"`
	// Flavor type (e.g. cpu, gpu)
	//
	// Any of "cpu", "gpu".
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PolicySpecFlavor) RawJSON() string { return r.JSON.raw }
func (r *PolicySpecFlavor) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Policy type, can be location or flavor
type PolicySpecType string

const (
	PolicySpecTypeLocation PolicySpecType = "location"
	PolicySpecTypeFlavor   PolicySpecType = "flavor"
	PolicySpecTypeMaxToken PolicySpecType = "maxToken"
)

// Policy specification
type PolicySpecParam struct {
	// Sandbox mode
	Sandbox param.Opt[bool] `json:"sandbox,omitzero"`
	// Types of hardware available for deployments
	Flavors []PolicySpecFlavorParam `json:"flavors,omitzero"`
	// PolicyLocations is a local type that wraps a slice of Location
	Locations []PolicyLocationParam `json:"locations,omitzero"`
	// PolicyMaxTokens is a local type that wraps a slice of PolicyMaxTokens
	MaxTokens PolicyMaxTokensParam `json:"maxTokens,omitzero"`
	// PolicyResourceTypes is a local type that wraps a slice of PolicyResourceType
	//
	// Any of "model", "function", "agent", "sandbox".
	ResourceTypes []string `json:"resourceTypes,omitzero"`
	// Policy type, can be location or flavor
	//
	// Any of "location", "flavor", "maxToken".
	Type PolicySpecType `json:"type,omitzero"`
	paramObj
}

func (r PolicySpecParam) MarshalJSON() (data []byte, err error) {
	type shadow PolicySpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicySpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A type of hardware available for deployments
type PolicySpecFlavorParam struct {
	// Flavor name (e.g. t4)
	Name param.Opt[string] `json:"name,omitzero"`
	// Flavor type (e.g. cpu, gpu)
	//
	// Any of "cpu", "gpu".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r PolicySpecFlavorParam) MarshalJSON() (data []byte, err error) {
	type shadow PolicySpecFlavorParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicySpecFlavorParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicySpecFlavorParam](
		"type", "cpu", "gpu",
	)
}

type PolicyNewParams struct {
	// Rule that controls how a deployment is made and served (e.g. location
	// restrictions)
	Policy PolicyParam
	paramObj
}

func (r PolicyNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Policy)
}
func (r *PolicyNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Policy)
}

type PolicyUpdateParams struct {
	// Rule that controls how a deployment is made and served (e.g. location
	// restrictions)
	Policy PolicyParam
	paramObj
}

func (r PolicyUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Policy)
}
func (r *PolicyUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Policy)
}
