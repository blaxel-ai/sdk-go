// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blaxel

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/blaxel-go/internal/apijson"
	"github.com/stainless-sdks/blaxel-go/internal/requestconfig"
	"github.com/stainless-sdks/blaxel-go/option"
	"github.com/stainless-sdks/blaxel-go/packages/respjson"
)

// ConfigurationService contains methods and other services that help with
// interacting with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewConfigurationService] method instead.
type ConfigurationService struct {
	Options []option.RequestOption
}

// NewConfigurationService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewConfigurationService(opts ...option.RequestOption) (r ConfigurationService) {
	r = ConfigurationService{}
	r.Options = opts
	return
}

// List all configurations
func (r *ConfigurationService) List(ctx context.Context, opts ...option.RequestOption) (res *ConfigurationListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "configuration"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Configuration
type ConfigurationListResponse struct {
	// Continents
	Continents []ConfigurationListResponseContinent `json:"continents"`
	// Countries
	Countries []ConfigurationListResponseCountry `json:"countries"`
	// Private locations managed with blaxel operator
	PrivateLocations []ConfigurationListResponsePrivateLocation `json:"privateLocations"`
	// Regions
	Regions []ConfigurationListResponseRegion `json:"regions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Continents       respjson.Field
		Countries        respjson.Field
		PrivateLocations respjson.Field
		Regions          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigurationListResponse) RawJSON() string { return r.JSON.raw }
func (r *ConfigurationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Continent
type ConfigurationListResponseContinent struct {
	// Continent display name
	DisplayName string `json:"displayName"`
	// Continent code
	Name        string         `json:"name"`
	ExtraFields map[string]any `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DisplayName respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigurationListResponseContinent) RawJSON() string { return r.JSON.raw }
func (r *ConfigurationListResponseContinent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Configuration
type ConfigurationListResponseCountry struct {
	// Country display name
	DisplayName string `json:"displayName"`
	// Country code
	Name        string         `json:"name"`
	ExtraFields map[string]any `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DisplayName respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigurationListResponseCountry) RawJSON() string { return r.JSON.raw }
func (r *ConfigurationListResponseCountry) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Private location available for policies
type ConfigurationListResponsePrivateLocation struct {
	// Location name
	Name        string         `json:"name"`
	ExtraFields map[string]any `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigurationListResponsePrivateLocation) RawJSON() string { return r.JSON.raw }
func (r *ConfigurationListResponsePrivateLocation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Region
type ConfigurationListResponseRegion struct {
	// Region display name
	Allowed string `json:"allowed"`
	// Region display name
	Continent string `json:"continent"`
	// Region display name
	Country string `json:"country"`
	// Region display name
	InfoGeneration string `json:"infoGeneration"`
	// Region display name
	Location string `json:"location"`
	// Region name
	Name        string         `json:"name"`
	ExtraFields map[string]any `json:",extras"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Allowed        respjson.Field
		Continent      respjson.Field
		Country        respjson.Field
		InfoGeneration respjson.Field
		Location       respjson.Field
		Name           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConfigurationListResponseRegion) RawJSON() string { return r.JSON.raw }
func (r *ConfigurationListResponseRegion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
