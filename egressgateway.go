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
	return
}

// An egress gateway that manages outbound traffic routing within a VPC. Multiple
// egress IPs can be allocated from a single gateway.
type EgressgatewayListResponse struct {
	// Metadata for an egress gateway resource including workspace, VPC, and name
	Metadata EgressgatewayListResponseMetadata `json:"metadata,required"`
	// Specification for an egress gateway including region and capacity configuration
	Spec EgressgatewayListResponseSpec `json:"spec,required"`
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
func (r EgressgatewayListResponse) RawJSON() string { return r.JSON.raw }
func (r *EgressgatewayListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata for an egress gateway resource including workspace, VPC, and name
type EgressgatewayListResponseMetadata struct {
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
func (r EgressgatewayListResponseMetadata) RawJSON() string { return r.JSON.raw }
func (r *EgressgatewayListResponseMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specification for an egress gateway including region and capacity configuration
type EgressgatewayListResponseSpec struct {
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
func (r EgressgatewayListResponseSpec) RawJSON() string { return r.JSON.raw }
func (r *EgressgatewayListResponseSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
