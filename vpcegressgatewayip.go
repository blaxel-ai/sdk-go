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

// VpcEgressgatewayIPService contains methods and other services that help with
// interacting with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVpcEgressgatewayIPService] method instead.
type VpcEgressgatewayIPService struct {
	Options []option.RequestOption
}

// NewVpcEgressgatewayIPService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewVpcEgressgatewayIPService(opts ...option.RequestOption) (r VpcEgressgatewayIPService) {
	r = VpcEgressgatewayIPService{}
	r.Options = opts
	return
}

// Get an egress IP by name
func (r *VpcEgressgatewayIPService) Get(ctx context.Context, ipName string, query VpcEgressgatewayIPGetParams, opts ...option.RequestOption) (res *VpcEgressgatewayIPGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.VpcName == "" {
		err = errors.New("missing required vpcName parameter")
		return nil, err
	}
	if query.GatewayName == "" {
		err = errors.New("missing required gatewayName parameter")
		return nil, err
	}
	if ipName == "" {
		err = errors.New("missing required ipName parameter")
		return nil, err
	}
	path := fmt.Sprintf("vpcs/%s/egressgateways/%s/ips/%s", query.VpcName, query.GatewayName, ipName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List egress IPs in a gateway
func (r *VpcEgressgatewayIPService) List(ctx context.Context, gatewayName string, query VpcEgressgatewayIPListParams, opts ...option.RequestOption) (res *[]VpcEgressgatewayIPListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.VpcName == "" {
		err = errors.New("missing required vpcName parameter")
		return nil, err
	}
	if gatewayName == "" {
		err = errors.New("missing required gatewayName parameter")
		return nil, err
	}
	path := fmt.Sprintf("vpcs/%s/egressgateways/%s/ips", query.VpcName, gatewayName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Delete an egress IP
func (r *VpcEgressgatewayIPService) Delete(ctx context.Context, ipName string, body VpcEgressgatewayIPDeleteParams, opts ...option.RequestOption) (res *VpcEgressgatewayIPDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.VpcName == "" {
		err = errors.New("missing required vpcName parameter")
		return nil, err
	}
	if body.GatewayName == "" {
		err = errors.New("missing required gatewayName parameter")
		return nil, err
	}
	if ipName == "" {
		err = errors.New("missing required ipName parameter")
		return nil, err
	}
	path := fmt.Sprintf("vpcs/%s/egressgateways/%s/ips/%s", body.VpcName, body.GatewayName, ipName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Allocate a new egress IP from the gateway
func (r *VpcEgressgatewayIPService) Allocate(ctx context.Context, gatewayName string, params VpcEgressgatewayIPAllocateParams, opts ...option.RequestOption) (res *VpcEgressgatewayIPAllocateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.VpcName == "" {
		err = errors.New("missing required vpcName parameter")
		return nil, err
	}
	if gatewayName == "" {
		err = errors.New("missing required gatewayName parameter")
		return nil, err
	}
	path := fmt.Sprintf("vpcs/%s/egressgateways/%s/ips", params.VpcName, gatewayName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// An individual IP address allocated from an egress gateway for dedicated outbound
// traffic
type VpcEgressgatewayIPGetResponse struct {
	// Metadata for an egress IP address including gateway association and name
	Metadata shared.EgressIPMetadata `json:"metadata" api:"required"`
	// Specification for an egress IP including IP family and the provisioned address
	Spec shared.EgressIPSpec `json:"spec" api:"required"`
	// Events happening on a resource deployed on Blaxel
	Events []CoreEvent `json:"events"`
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
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VpcEgressgatewayIPGetResponse) RawJSON() string { return r.JSON.raw }
func (r *VpcEgressgatewayIPGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An individual IP address allocated from an egress gateway for dedicated outbound
// traffic
type VpcEgressgatewayIPListResponse struct {
	// Metadata for an egress IP address including gateway association and name
	Metadata shared.EgressIPMetadata `json:"metadata" api:"required"`
	// Specification for an egress IP including IP family and the provisioned address
	Spec shared.EgressIPSpec `json:"spec" api:"required"`
	// Events happening on a resource deployed on Blaxel
	Events []CoreEvent `json:"events"`
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
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VpcEgressgatewayIPListResponse) RawJSON() string { return r.JSON.raw }
func (r *VpcEgressgatewayIPListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An individual IP address allocated from an egress gateway for dedicated outbound
// traffic
type VpcEgressgatewayIPDeleteResponse struct {
	// Metadata for an egress IP address including gateway association and name
	Metadata shared.EgressIPMetadata `json:"metadata" api:"required"`
	// Specification for an egress IP including IP family and the provisioned address
	Spec shared.EgressIPSpec `json:"spec" api:"required"`
	// Events happening on a resource deployed on Blaxel
	Events []CoreEvent `json:"events"`
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
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VpcEgressgatewayIPDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *VpcEgressgatewayIPDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An individual IP address allocated from an egress gateway for dedicated outbound
// traffic
type VpcEgressgatewayIPAllocateResponse struct {
	// Metadata for an egress IP address including gateway association and name
	Metadata shared.EgressIPMetadata `json:"metadata" api:"required"`
	// Specification for an egress IP including IP family and the provisioned address
	Spec shared.EgressIPSpec `json:"spec" api:"required"`
	// Events happening on a resource deployed on Blaxel
	Events []CoreEvent `json:"events"`
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
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VpcEgressgatewayIPAllocateResponse) RawJSON() string { return r.JSON.raw }
func (r *VpcEgressgatewayIPAllocateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VpcEgressgatewayIPGetParams struct {
	VpcName     string `path:"vpcName" api:"required" json:"-"`
	GatewayName string `path:"gatewayName" api:"required" json:"-"`
	paramObj
}

type VpcEgressgatewayIPListParams struct {
	VpcName string `path:"vpcName" api:"required" json:"-"`
	paramObj
}

type VpcEgressgatewayIPDeleteParams struct {
	VpcName     string `path:"vpcName" api:"required" json:"-"`
	GatewayName string `path:"gatewayName" api:"required" json:"-"`
	paramObj
}

type VpcEgressgatewayIPAllocateParams struct {
	VpcName string `path:"vpcName" api:"required" json:"-"`
	// Metadata for an egress IP address including gateway association and name
	Metadata shared.EgressIPMetadataParam `json:"metadata,omitzero" api:"required"`
	// Specification for an egress IP including IP family and the provisioned address
	Spec shared.EgressIPSpecParam `json:"spec,omitzero" api:"required"`
	paramObj
}

func (r VpcEgressgatewayIPAllocateParams) MarshalJSON() (data []byte, err error) {
	type shadow VpcEgressgatewayIPAllocateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VpcEgressgatewayIPAllocateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
