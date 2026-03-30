// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blaxel

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/blaxel-ai/sdk-go/internal/apijson"
	"github.com/blaxel-ai/sdk-go/internal/requestconfig"
	"github.com/blaxel-ai/sdk-go/option"
	"github.com/blaxel-ai/sdk-go/packages/param"
	"github.com/blaxel-ai/sdk-go/packages/respjson"
	"github.com/blaxel-ai/sdk-go/shared"
)

// VpcEgressgatewayService contains methods and other services that help with
// interacting with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVpcEgressgatewayService] method instead.
type VpcEgressgatewayService struct {
	Options []option.RequestOption
	IPs     VpcEgressgatewayIPService
}

// NewVpcEgressgatewayService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewVpcEgressgatewayService(opts ...option.RequestOption) (r VpcEgressgatewayService) {
	r = VpcEgressgatewayService{}
	r.Options = opts
	r.IPs = NewVpcEgressgatewayIPService(opts...)
	return
}

// Create an egress gateway in a VPC
func (r *VpcEgressgatewayService) New(ctx context.Context, vpcName string, body VpcEgressgatewayNewParams, opts ...option.RequestOption) (res *VpcEgressgatewayNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if vpcName == "" {
		err = errors.New("missing required vpcName parameter")
		return nil, err
	}
	path := fmt.Sprintf("vpcs/%s/egressgateways", vpcName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get an egress gateway by name
func (r *VpcEgressgatewayService) Get(ctx context.Context, gatewayName string, query VpcEgressgatewayGetParams, opts ...option.RequestOption) (res *VpcEgressgatewayGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.VpcName == "" {
		err = errors.New("missing required vpcName parameter")
		return nil, err
	}
	if gatewayName == "" {
		err = errors.New("missing required gatewayName parameter")
		return nil, err
	}
	path := fmt.Sprintf("vpcs/%s/egressgateways/%s", query.VpcName, gatewayName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List egress gateways in a VPC
func (r *VpcEgressgatewayService) List(ctx context.Context, vpcName string, opts ...option.RequestOption) (res *[]VpcEgressgatewayListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if vpcName == "" {
		err = errors.New("missing required vpcName parameter")
		return nil, err
	}
	path := fmt.Sprintf("vpcs/%s/egressgateways", vpcName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Delete an egress gateway
func (r *VpcEgressgatewayService) Delete(ctx context.Context, gatewayName string, body VpcEgressgatewayDeleteParams, opts ...option.RequestOption) (res *VpcEgressgatewayDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.VpcName == "" {
		err = errors.New("missing required vpcName parameter")
		return nil, err
	}
	if gatewayName == "" {
		err = errors.New("missing required gatewayName parameter")
		return nil, err
	}
	path := fmt.Sprintf("vpcs/%s/egressgateways/%s", body.VpcName, gatewayName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// An egress gateway that manages outbound traffic routing within a VPC. Multiple
// egress IPs can be allocated from a single gateway.
type VpcEgressgatewayNewResponse struct {
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
func (r VpcEgressgatewayNewResponse) RawJSON() string { return r.JSON.raw }
func (r *VpcEgressgatewayNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An egress gateway that manages outbound traffic routing within a VPC. Multiple
// egress IPs can be allocated from a single gateway.
type VpcEgressgatewayGetResponse struct {
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
func (r VpcEgressgatewayGetResponse) RawJSON() string { return r.JSON.raw }
func (r *VpcEgressgatewayGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An egress gateway that manages outbound traffic routing within a VPC. Multiple
// egress IPs can be allocated from a single gateway.
type VpcEgressgatewayListResponse struct {
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
func (r VpcEgressgatewayListResponse) RawJSON() string { return r.JSON.raw }
func (r *VpcEgressgatewayListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An egress gateway that manages outbound traffic routing within a VPC. Multiple
// egress IPs can be allocated from a single gateway.
type VpcEgressgatewayDeleteResponse struct {
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
func (r VpcEgressgatewayDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *VpcEgressgatewayDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VpcEgressgatewayNewParams struct {
	// Metadata for an egress gateway resource including workspace, VPC, and name
	Metadata shared.EgressGatewayMetadataParam `json:"metadata,omitzero" api:"required"`
	// Specification for an egress gateway including region and capacity configuration
	Spec shared.EgressGatewaySpecParam `json:"spec,omitzero" api:"required"`
	paramObj
}

func (r VpcEgressgatewayNewParams) MarshalJSON() (data []byte, err error) {
	type shadow VpcEgressgatewayNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VpcEgressgatewayNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VpcEgressgatewayGetParams struct {
	VpcName string `path:"vpcName" api:"required" json:"-"`
	paramObj
}

type VpcEgressgatewayDeleteParams struct {
	VpcName string `path:"vpcName" api:"required" json:"-"`
	paramObj
}
