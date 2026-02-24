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
		return
	}
	if query.GatewayName == "" {
		err = errors.New("missing required gatewayName parameter")
		return
	}
	if ipName == "" {
		err = errors.New("missing required ipName parameter")
		return
	}
	path := fmt.Sprintf("vpcs/%s/egressgateways/%s/ips/%s", query.VpcName, query.GatewayName, ipName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List egress IPs in a gateway
func (r *VpcEgressgatewayIPService) List(ctx context.Context, gatewayName string, query VpcEgressgatewayIPListParams, opts ...option.RequestOption) (res *[]VpcEgressgatewayIPListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.VpcName == "" {
		err = errors.New("missing required vpcName parameter")
		return
	}
	if gatewayName == "" {
		err = errors.New("missing required gatewayName parameter")
		return
	}
	path := fmt.Sprintf("vpcs/%s/egressgateways/%s/ips", query.VpcName, gatewayName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete an egress IP
func (r *VpcEgressgatewayIPService) Delete(ctx context.Context, ipName string, body VpcEgressgatewayIPDeleteParams, opts ...option.RequestOption) (res *VpcEgressgatewayIPDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.VpcName == "" {
		err = errors.New("missing required vpcName parameter")
		return
	}
	if body.GatewayName == "" {
		err = errors.New("missing required gatewayName parameter")
		return
	}
	if ipName == "" {
		err = errors.New("missing required ipName parameter")
		return
	}
	path := fmt.Sprintf("vpcs/%s/egressgateways/%s/ips/%s", body.VpcName, body.GatewayName, ipName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Allocate a new egress IP from the gateway
func (r *VpcEgressgatewayIPService) Allocate(ctx context.Context, gatewayName string, params VpcEgressgatewayIPAllocateParams, opts ...option.RequestOption) (res *VpcEgressgatewayIPAllocateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.VpcName == "" {
		err = errors.New("missing required vpcName parameter")
		return
	}
	if gatewayName == "" {
		err = errors.New("missing required gatewayName parameter")
		return
	}
	path := fmt.Sprintf("vpcs/%s/egressgateways/%s/ips", params.VpcName, gatewayName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// An individual IP address allocated from an egress gateway for dedicated outbound
// traffic
type VpcEgressgatewayIPGetResponse struct {
	// Metadata for an egress IP address including gateway association and name
	Metadata VpcEgressgatewayIPGetResponseMetadata `json:"metadata" api:"required"`
	// Specification for an egress IP including IP family and the provisioned address
	Spec VpcEgressgatewayIPGetResponseSpec `json:"spec" api:"required"`
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

// Metadata for an egress IP address including gateway association and name
type VpcEgressgatewayIPGetResponseMetadata struct {
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
func (r VpcEgressgatewayIPGetResponseMetadata) RawJSON() string { return r.JSON.raw }
func (r *VpcEgressgatewayIPGetResponseMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specification for an egress IP including IP family and the provisioned address
type VpcEgressgatewayIPGetResponseSpec struct {
	// IP address family, either IPv4 or IPv6
	//
	// Any of "IPv4", "IPv6".
	IPFamily string `json:"ipFamily" api:"required"`
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
func (r VpcEgressgatewayIPGetResponseSpec) RawJSON() string { return r.JSON.raw }
func (r *VpcEgressgatewayIPGetResponseSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An individual IP address allocated from an egress gateway for dedicated outbound
// traffic
type VpcEgressgatewayIPListResponse struct {
	// Metadata for an egress IP address including gateway association and name
	Metadata VpcEgressgatewayIPListResponseMetadata `json:"metadata" api:"required"`
	// Specification for an egress IP including IP family and the provisioned address
	Spec VpcEgressgatewayIPListResponseSpec `json:"spec" api:"required"`
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

// Metadata for an egress IP address including gateway association and name
type VpcEgressgatewayIPListResponseMetadata struct {
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
func (r VpcEgressgatewayIPListResponseMetadata) RawJSON() string { return r.JSON.raw }
func (r *VpcEgressgatewayIPListResponseMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specification for an egress IP including IP family and the provisioned address
type VpcEgressgatewayIPListResponseSpec struct {
	// IP address family, either IPv4 or IPv6
	//
	// Any of "IPv4", "IPv6".
	IPFamily string `json:"ipFamily" api:"required"`
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
func (r VpcEgressgatewayIPListResponseSpec) RawJSON() string { return r.JSON.raw }
func (r *VpcEgressgatewayIPListResponseSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An individual IP address allocated from an egress gateway for dedicated outbound
// traffic
type VpcEgressgatewayIPDeleteResponse struct {
	// Metadata for an egress IP address including gateway association and name
	Metadata VpcEgressgatewayIPDeleteResponseMetadata `json:"metadata" api:"required"`
	// Specification for an egress IP including IP family and the provisioned address
	Spec VpcEgressgatewayIPDeleteResponseSpec `json:"spec" api:"required"`
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

// Metadata for an egress IP address including gateway association and name
type VpcEgressgatewayIPDeleteResponseMetadata struct {
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
func (r VpcEgressgatewayIPDeleteResponseMetadata) RawJSON() string { return r.JSON.raw }
func (r *VpcEgressgatewayIPDeleteResponseMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specification for an egress IP including IP family and the provisioned address
type VpcEgressgatewayIPDeleteResponseSpec struct {
	// IP address family, either IPv4 or IPv6
	//
	// Any of "IPv4", "IPv6".
	IPFamily string `json:"ipFamily" api:"required"`
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
func (r VpcEgressgatewayIPDeleteResponseSpec) RawJSON() string { return r.JSON.raw }
func (r *VpcEgressgatewayIPDeleteResponseSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An individual IP address allocated from an egress gateway for dedicated outbound
// traffic
type VpcEgressgatewayIPAllocateResponse struct {
	// Metadata for an egress IP address including gateway association and name
	Metadata VpcEgressgatewayIPAllocateResponseMetadata `json:"metadata" api:"required"`
	// Specification for an egress IP including IP family and the provisioned address
	Spec VpcEgressgatewayIPAllocateResponseSpec `json:"spec" api:"required"`
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

// Metadata for an egress IP address including gateway association and name
type VpcEgressgatewayIPAllocateResponseMetadata struct {
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
func (r VpcEgressgatewayIPAllocateResponseMetadata) RawJSON() string { return r.JSON.raw }
func (r *VpcEgressgatewayIPAllocateResponseMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specification for an egress IP including IP family and the provisioned address
type VpcEgressgatewayIPAllocateResponseSpec struct {
	// IP address family, either IPv4 or IPv6
	//
	// Any of "IPv4", "IPv6".
	IPFamily string `json:"ipFamily" api:"required"`
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
func (r VpcEgressgatewayIPAllocateResponseSpec) RawJSON() string { return r.JSON.raw }
func (r *VpcEgressgatewayIPAllocateResponseSpec) UnmarshalJSON(data []byte) error {
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
	Metadata VpcEgressgatewayIPAllocateParamsMetadata `json:"metadata,omitzero" api:"required"`
	// Specification for an egress IP including IP family and the provisioned address
	Spec VpcEgressgatewayIPAllocateParamsSpec `json:"spec,omitzero" api:"required"`
	paramObj
}

func (r VpcEgressgatewayIPAllocateParams) MarshalJSON() (data []byte, err error) {
	type shadow VpcEgressgatewayIPAllocateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VpcEgressgatewayIPAllocateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata for an egress IP address including gateway association and name
//
// The property Name is required.
type VpcEgressgatewayIPAllocateParamsMetadata struct {
	// Unique identifier for the egress IP within the gateway. Must be lowercase
	// alphanumeric with hyphens, max 49 characters. Immutable after creation.
	Name string `json:"name" api:"required"`
	// Human-readable name for display in the UI. Can contain spaces and special
	// characters, max 63 characters.
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	paramObj
}

func (r VpcEgressgatewayIPAllocateParamsMetadata) MarshalJSON() (data []byte, err error) {
	type shadow VpcEgressgatewayIPAllocateParamsMetadata
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VpcEgressgatewayIPAllocateParamsMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specification for an egress IP including IP family and the provisioned address
//
// The property IPFamily is required.
type VpcEgressgatewayIPAllocateParamsSpec struct {
	// IP address family, either IPv4 or IPv6
	//
	// Any of "IPv4", "IPv6".
	IPFamily string `json:"ipFamily,omitzero" api:"required"`
	paramObj
}

func (r VpcEgressgatewayIPAllocateParamsSpec) MarshalJSON() (data []byte, err error) {
	type shadow VpcEgressgatewayIPAllocateParamsSpec
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VpcEgressgatewayIPAllocateParamsSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[VpcEgressgatewayIPAllocateParamsSpec](
		"ipFamily", "IPv4", "IPv6",
	)
}
