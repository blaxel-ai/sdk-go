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

// EgressipService contains methods and other services that help with interacting
// with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewEgressipService] method instead.
type EgressipService struct {
	Options []option.RequestOption
}

// NewEgressipService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewEgressipService(opts ...option.RequestOption) (r EgressipService) {
	r = EgressipService{}
	r.Options = opts
	return
}

// List all egress IPs across all VPCs and gateways in the workspace
func (r *EgressipService) List(ctx context.Context, opts ...option.RequestOption) (res *[]EgressipListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "egressips"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// An individual IP address allocated from an egress gateway for dedicated outbound
// traffic
type EgressipListResponse struct {
	// Metadata for an egress IP address including gateway association and name
	Metadata EgressipListResponseMetadata `json:"metadata,required"`
	// Specification for an egress IP including IP family and the provisioned address
	Spec EgressipListResponseSpec `json:"spec,required"`
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
func (r EgressipListResponse) RawJSON() string { return r.JSON.raw }
func (r *EgressipListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Metadata for an egress IP address including gateway association and name
type EgressipListResponseMetadata struct {
	// Unique identifier for the egress IP within the gateway. Must be lowercase
	// alphanumeric with hyphens, max 49 characters. Immutable after creation.
	Name string `json:"name,required"`
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
func (r EgressipListResponseMetadata) RawJSON() string { return r.JSON.raw }
func (r *EgressipListResponseMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Specification for an egress IP including IP family and the provisioned address
type EgressipListResponseSpec struct {
	// IP address family, either IPv4 or IPv6
	//
	// Any of "IPv4", "IPv6".
	IPFamily string `json:"ipFamily,required"`
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
func (r EgressipListResponseSpec) RawJSON() string { return r.JSON.raw }
func (r *EgressipListResponseSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
