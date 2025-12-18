// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blaxel

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/blaxel-go/internal/apijson"
	"github.com/stainless-sdks/blaxel-go/internal/apiquery"
	"github.com/stainless-sdks/blaxel-go/internal/requestconfig"
	"github.com/stainless-sdks/blaxel-go/option"
	"github.com/stainless-sdks/blaxel-go/packages/param"
	"github.com/stainless-sdks/blaxel-go/packages/respjson"
)

// PublicIPService contains methods and other services that help with interacting
// with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPublicIPService] method instead.
type PublicIPService struct {
	Options []option.RequestOption
}

// NewPublicIPService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewPublicIPService(opts ...option.RequestOption) (r PublicIPService) {
	r = PublicIPService{}
	r.Options = opts
	return
}

// Returns a list of all public ips used in Blaxel..
func (r *PublicIPService) List(ctx context.Context, query PublicIPListParams, opts ...option.RequestOption) (res *PublicIPListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "publicIps"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type PublicIPListResponse map[string]PublicIPListResponseItem

type PublicIPListResponseItem struct {
	// Description of the region/location
	Description string `json:"description"`
	// List of public ipv4 addresses
	Ipv4Cidrs []string `json:"ipv4Cidrs"`
	// List of public ipv6 addresses
	Ipv6Cidrs []string `json:"ipv6Cidrs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Ipv4Cidrs   respjson.Field
		Ipv6Cidrs   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PublicIPListResponseItem) RawJSON() string { return r.JSON.raw }
func (r *PublicIPListResponseItem) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PublicIPListParams struct {
	// Filter by region name (only returns mk3 region data)
	Region param.Opt[string] `query:"region,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [PublicIPListParams]'s query parameters as `url.Values`.
func (r PublicIPListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
