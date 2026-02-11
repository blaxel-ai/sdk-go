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
		return
	}
	path := fmt.Sprintf("vpcs/%s/egressgateways", vpcName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get an egress gateway by name
func (r *VpcEgressgatewayService) Get(ctx context.Context, gatewayName string, query VpcEgressgatewayGetParams, opts ...option.RequestOption) (res *VpcEgressgatewayGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.VpcName == "" {
		err = errors.New("missing required vpcName parameter")
		return
	}
	if gatewayName == "" {
		err = errors.New("missing required gatewayName parameter")
		return
	}
	path := fmt.Sprintf("vpcs/%s/egressgateways/%s", query.VpcName, gatewayName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List egress gateways in a VPC
func (r *VpcEgressgatewayService) List(ctx context.Context, vpcName string, opts ...option.RequestOption) (res *[]VpcEgressgatewayListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if vpcName == "" {
		err = errors.New("missing required vpcName parameter")
		return
	}
	path := fmt.Sprintf("vpcs/%s/egressgateways", vpcName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete an egress gateway
func (r *VpcEgressgatewayService) Delete(ctx context.Context, gatewayName string, body VpcEgressgatewayDeleteParams, opts ...option.RequestOption) (res *VpcEgressgatewayDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.VpcName == "" {
		err = errors.New("missing required vpcName parameter")
		return
	}
	if gatewayName == "" {
		err = errors.New("missing required gatewayName parameter")
		return
	}
	path := fmt.Sprintf("vpcs/%s/egressgateways/%s", body.VpcName, gatewayName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// An egress gateway that manages outbound traffic routing within a VPC. Multiple
// egress IPs can be allocated from a single gateway.
type VpcEgressgatewayNewResponse struct {
	// Metadata for an egress gateway resource including workspace, VPC, and name
	Metadata VpcEgressgatewayNewResponseMetadata `json:"metadata,required"`
	// Specification for an egress gateway including region and capacity configuration
	Spec VpcEgressgatewayNewResponseSpec `json:"spec,required"`
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
func (r VpcEgressgatewayNewResponse) RawJSON() string { return r.JSON.raw }
func (r *VpcEgressgatewayNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata for an egress gateway resource including workspace, VPC, and name
type VpcEgressgatewayNewResponseMetadata struct {
	// Unique identifier for the egress gateway within the VPC. Must be lowercase
	// alphanumeric with hyphens, max 49 characters. Immutable after creation.
	Name string `json:"name,required"`
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
func (r VpcEgressgatewayNewResponseMetadata) RawJSON() string { return r.JSON.raw }
func (r *VpcEgressgatewayNewResponseMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specification for an egress gateway including region and capacity configuration
type VpcEgressgatewayNewResponseSpec struct {
	// Region where the egress gateway is provisioned (e.g. us-pdx-1, eu-lon-1)
	Region string `json:"region,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Region      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VpcEgressgatewayNewResponseSpec) RawJSON() string { return r.JSON.raw }
func (r *VpcEgressgatewayNewResponseSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An egress gateway that manages outbound traffic routing within a VPC. Multiple
// egress IPs can be allocated from a single gateway.
type VpcEgressgatewayGetResponse struct {
	// Metadata for an egress gateway resource including workspace, VPC, and name
	Metadata VpcEgressgatewayGetResponseMetadata `json:"metadata,required"`
	// Specification for an egress gateway including region and capacity configuration
	Spec VpcEgressgatewayGetResponseSpec `json:"spec,required"`
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
func (r VpcEgressgatewayGetResponse) RawJSON() string { return r.JSON.raw }
func (r *VpcEgressgatewayGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata for an egress gateway resource including workspace, VPC, and name
type VpcEgressgatewayGetResponseMetadata struct {
	// Unique identifier for the egress gateway within the VPC. Must be lowercase
	// alphanumeric with hyphens, max 49 characters. Immutable after creation.
	Name string `json:"name,required"`
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
func (r VpcEgressgatewayGetResponseMetadata) RawJSON() string { return r.JSON.raw }
func (r *VpcEgressgatewayGetResponseMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specification for an egress gateway including region and capacity configuration
type VpcEgressgatewayGetResponseSpec struct {
	// Region where the egress gateway is provisioned (e.g. us-pdx-1, eu-lon-1)
	Region string `json:"region,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Region      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VpcEgressgatewayGetResponseSpec) RawJSON() string { return r.JSON.raw }
func (r *VpcEgressgatewayGetResponseSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An egress gateway that manages outbound traffic routing within a VPC. Multiple
// egress IPs can be allocated from a single gateway.
type VpcEgressgatewayListResponse struct {
	// Metadata for an egress gateway resource including workspace, VPC, and name
	Metadata VpcEgressgatewayListResponseMetadata `json:"metadata,required"`
	// Specification for an egress gateway including region and capacity configuration
	Spec VpcEgressgatewayListResponseSpec `json:"spec,required"`
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
func (r VpcEgressgatewayListResponse) RawJSON() string { return r.JSON.raw }
func (r *VpcEgressgatewayListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata for an egress gateway resource including workspace, VPC, and name
type VpcEgressgatewayListResponseMetadata struct {
	// Unique identifier for the egress gateway within the VPC. Must be lowercase
	// alphanumeric with hyphens, max 49 characters. Immutable after creation.
	Name string `json:"name,required"`
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
func (r VpcEgressgatewayListResponseMetadata) RawJSON() string { return r.JSON.raw }
func (r *VpcEgressgatewayListResponseMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specification for an egress gateway including region and capacity configuration
type VpcEgressgatewayListResponseSpec struct {
	// Region where the egress gateway is provisioned (e.g. us-pdx-1, eu-lon-1)
	Region string `json:"region,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Region      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VpcEgressgatewayListResponseSpec) RawJSON() string { return r.JSON.raw }
func (r *VpcEgressgatewayListResponseSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// An egress gateway that manages outbound traffic routing within a VPC. Multiple
// egress IPs can be allocated from a single gateway.
type VpcEgressgatewayDeleteResponse struct {
	// Metadata for an egress gateway resource including workspace, VPC, and name
	Metadata VpcEgressgatewayDeleteResponseMetadata `json:"metadata,required"`
	// Specification for an egress gateway including region and capacity configuration
	Spec VpcEgressgatewayDeleteResponseSpec `json:"spec,required"`
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
func (r VpcEgressgatewayDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *VpcEgressgatewayDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata for an egress gateway resource including workspace, VPC, and name
type VpcEgressgatewayDeleteResponseMetadata struct {
	// Unique identifier for the egress gateway within the VPC. Must be lowercase
	// alphanumeric with hyphens, max 49 characters. Immutable after creation.
	Name string `json:"name,required"`
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
func (r VpcEgressgatewayDeleteResponseMetadata) RawJSON() string { return r.JSON.raw }
func (r *VpcEgressgatewayDeleteResponseMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specification for an egress gateway including region and capacity configuration
type VpcEgressgatewayDeleteResponseSpec struct {
	// Region where the egress gateway is provisioned (e.g. us-pdx-1, eu-lon-1)
	Region string `json:"region,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Region      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VpcEgressgatewayDeleteResponseSpec) RawJSON() string { return r.JSON.raw }
func (r *VpcEgressgatewayDeleteResponseSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VpcEgressgatewayNewParams struct {
	// Metadata for an egress gateway resource including workspace, VPC, and name
	Metadata VpcEgressgatewayNewParamsMetadata `json:"metadata,omitzero,required"`
	// Specification for an egress gateway including region and capacity configuration
	Spec VpcEgressgatewayNewParamsSpec `json:"spec,omitzero,required"`
	paramObj
}

func (r VpcEgressgatewayNewParams) MarshalJSON() (data []byte, err error) {
	type shadow VpcEgressgatewayNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VpcEgressgatewayNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata for an egress gateway resource including workspace, VPC, and name
//
// The property Name is required.
type VpcEgressgatewayNewParamsMetadata struct {
	// Unique identifier for the egress gateway within the VPC. Must be lowercase
	// alphanumeric with hyphens, max 49 characters. Immutable after creation.
	Name string `json:"name,required"`
	// Human-readable name for display in the UI. Can contain spaces and special
	// characters, max 63 characters.
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	paramObj
}

func (r VpcEgressgatewayNewParamsMetadata) MarshalJSON() (data []byte, err error) {
	type shadow VpcEgressgatewayNewParamsMetadata
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VpcEgressgatewayNewParamsMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specification for an egress gateway including region and capacity configuration
//
// The property Region is required.
type VpcEgressgatewayNewParamsSpec struct {
	// Region where the egress gateway is provisioned (e.g. us-pdx-1, eu-lon-1)
	Region string `json:"region,required"`
	paramObj
}

func (r VpcEgressgatewayNewParamsSpec) MarshalJSON() (data []byte, err error) {
	type shadow VpcEgressgatewayNewParamsSpec
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VpcEgressgatewayNewParamsSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VpcEgressgatewayGetParams struct {
	VpcName string `path:"vpcName,required" json:"-"`
	paramObj
}

type VpcEgressgatewayDeleteParams struct {
	VpcName string `path:"vpcName,required" json:"-"`
	paramObj
}
