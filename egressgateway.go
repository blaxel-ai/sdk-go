// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blaxel

import (
	"context"
	"net/http"
	"slices"

	"github.com/blaxel-ai/sdk-go/internal/apijson"
	"github.com/blaxel-ai/sdk-go/internal/requestconfig"
	"github.com/blaxel-ai/sdk-go/option"
	"github.com/blaxel-ai/sdk-go/packages/respjson"
	"github.com/blaxel-ai/sdk-go/shared"
)

// EgressgatewayService contains methods and other services that help with
// interacting with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEgressgatewayService] method instead.
type EgressgatewayService struct {
	Options []option.RequestOption
}

// NewEgressgatewayService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEgressgatewayService(opts ...option.RequestOption) (r EgressgatewayService) {
	r = EgressgatewayService{}
	r.Options = opts
	return
}

// List all egress gateways across all VPCs in the workspace
func (r *EgressgatewayService) List(ctx context.Context, opts ...option.RequestOption) (res *[]EgressgatewayListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "egressgateways"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// An egress gateway that manages outbound traffic routing within a VPC. Multiple
// egress IPs can be allocated from a single gateway.
type EgressgatewayListResponse struct {
	// Metadata for an egress gateway resource including workspace, VPC, and name
	Metadata shared.EgressGatewayMetadata `json:"metadata" api:"required"`
	// Specification for an egress gateway including region and capacity configuration
	Spec shared.EgressGatewaySpec `json:"spec" api:"required"`
	// Events happening on a resource deployed on Blaxel
	Events []CoreEvent `json:"events"`
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
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EgressgatewayListResponse) RawJSON() string { return r.JSON.raw }
func (r *EgressgatewayListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
