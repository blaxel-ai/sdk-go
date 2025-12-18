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
	shimjson "github.com/stainless-sdks/blaxel-go/internal/encoding/json"
	"github.com/stainless-sdks/blaxel-go/internal/requestconfig"
	"github.com/stainless-sdks/blaxel-go/option"
	"github.com/stainless-sdks/blaxel-go/packages/param"
	"github.com/stainless-sdks/blaxel-go/packages/respjson"
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

// Creates a policy.
func (r *PolicyService) New(ctx context.Context, body PolicyNewParams, opts ...option.RequestOption) (res *Policy, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "policies"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns a policy by name.
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

// Updates a policy.
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

// Returns a list of all policies in the workspace.
func (r *PolicyService) List(ctx context.Context, opts ...option.RequestOption) (res *[]Policy, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "policies"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a policy by name.
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
	// Metadata
	Metadata Metadata `json:"metadata,required"`
	// Policy specification
	Spec PolicySpec `json:"spec,required"`
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

// Policy specification
type PolicySpec struct {
	// Flavors allowed by the policy. If not set, all flavors are allowed.
	Flavors []Flavor `json:"flavors"`
	// Locations allowed by the policy. If not set, all locations are allowed.
	Locations []PolicySpecLocation `json:"locations"`
	// Max token allowed by the policy. If not set, no max token is allowed.
	MaxTokens PolicySpecMaxTokens `json:"maxTokens"`
	// ResourceTypes where the policy is applied. If not set, the policy is applied to
	// all resource types.
	//
	// Any of "model", "function", "agent", "sandbox".
	ResourceTypes []string `json:"resourceTypes"`
	// Sandbox mode
	Sandbox bool `json:"sandbox"`
	// Policy type, can be location or flavor
	//
	// Any of "location", "flavor", "maxToken".
	Type string `json:"type"`
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

// Policy location
type PolicySpecLocation struct {
	// Policy location name
	Name string `json:"name"`
	// Policy location type
	//
	// Any of "location", "country", "continent".
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
func (r PolicySpecLocation) RawJSON() string { return r.JSON.raw }
func (r *PolicySpecLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Max token allowed by the policy. If not set, no max token is allowed.
type PolicySpecMaxTokens struct {
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
func (r PolicySpecMaxTokens) RawJSON() string { return r.JSON.raw }
func (r *PolicySpecMaxTokens) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Rule that controls how a deployment is made and served (e.g. location
// restrictions)
//
// The properties Metadata, Spec are required.
type PolicyParam struct {
	// Metadata
	Metadata MetadataParam `json:"metadata,omitzero,required"`
	// Policy specification
	Spec PolicySpecParam `json:"spec,omitzero,required"`
	paramObj
}

func (r PolicyParam) MarshalJSON() (data []byte, err error) {
	type shadow PolicyParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicyParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Policy specification
type PolicySpecParam struct {
	// Sandbox mode
	Sandbox param.Opt[bool] `json:"sandbox,omitzero"`
	// Flavors allowed by the policy. If not set, all flavors are allowed.
	Flavors []FlavorParam `json:"flavors,omitzero"`
	// Locations allowed by the policy. If not set, all locations are allowed.
	Locations []PolicySpecLocationParam `json:"locations,omitzero"`
	// Max token allowed by the policy. If not set, no max token is allowed.
	MaxTokens PolicySpecMaxTokensParam `json:"maxTokens,omitzero"`
	// ResourceTypes where the policy is applied. If not set, the policy is applied to
	// all resource types.
	//
	// Any of "model", "function", "agent", "sandbox".
	ResourceTypes []string `json:"resourceTypes,omitzero"`
	// Policy type, can be location or flavor
	//
	// Any of "location", "flavor", "maxToken".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r PolicySpecParam) MarshalJSON() (data []byte, err error) {
	type shadow PolicySpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicySpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicySpecParam](
		"type", "location", "flavor", "maxToken",
	)
}

// Policy location
type PolicySpecLocationParam struct {
	// Policy location name
	Name param.Opt[string] `json:"name,omitzero"`
	// Policy location type
	//
	// Any of "location", "country", "continent".
	Type string `json:"type,omitzero"`
	paramObj
}

func (r PolicySpecLocationParam) MarshalJSON() (data []byte, err error) {
	type shadow PolicySpecLocationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicySpecLocationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[PolicySpecLocationParam](
		"type", "location", "country", "continent",
	)
}

// Max token allowed by the policy. If not set, no max token is allowed.
type PolicySpecMaxTokensParam struct {
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

func (r PolicySpecMaxTokensParam) MarshalJSON() (data []byte, err error) {
	type shadow PolicySpecMaxTokensParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PolicySpecMaxTokensParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
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
