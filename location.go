// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blaxel

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"

	"github.com/stainless-sdks/blaxel-go/internal/apijson"
	"github.com/stainless-sdks/blaxel-go/internal/requestconfig"
	"github.com/stainless-sdks/blaxel-go/option"
	"github.com/stainless-sdks/blaxel-go/packages/param"
	"github.com/stainless-sdks/blaxel-go/packages/respjson"
)

// LocationService contains methods and other services that help with interacting
// with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewLocationService] method instead.
type LocationService struct {
	Options []option.RequestOption
}

// NewLocationService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewLocationService(opts ...option.RequestOption) (r LocationService) {
	r = LocationService{}
	r.Options = opts
	return
}

// Returns all deployment regions with their current availability status and
// supported hardware flavors. Use this to discover where resources can be
// deployed.
func (r *LocationService) List(ctx context.Context, opts ...option.RequestOption) (res *[]LocationListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "locations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
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

// Location availability for policies
type LocationListResponse struct {
	// Continent of the location
	Continent string `json:"continent"`
	// Country of the location
	Country string `json:"country"`
	// Hardware flavors available in the location
	Flavors []Flavor `json:"flavors"`
	// Name of the location
	Location string `json:"location"`
	// Region of the location
	Region string `json:"region"`
	// Status of the location
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Continent   respjson.Field
		Country     respjson.Field
		Flavors     respjson.Field
		Location    respjson.Field
		Region      respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r LocationListResponse) RawJSON() string { return r.JSON.raw }
func (r *LocationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
