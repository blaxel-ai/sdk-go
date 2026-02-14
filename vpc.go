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

// VpcService contains methods and other services that help with interacting with
// the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVpcService] method instead.
type VpcService struct {
	Options        []option.RequestOption
	Egressgateways VpcEgressgatewayService
}

// NewVpcService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewVpcService(opts ...option.RequestOption) (r VpcService) {
	r = VpcService{}
	r.Options = opts
	r.Egressgateways = NewVpcEgressgatewayService(opts...)
	return
}

// Create a VPC for the workspace
func (r *VpcService) New(ctx context.Context, body VpcNewParams, opts ...option.RequestOption) (res *VpcNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "vpcs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get a VPC by name
func (r *VpcService) Get(ctx context.Context, vpcName string, opts ...option.RequestOption) (res *VpcGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if vpcName == "" {
		err = errors.New("missing required vpcName parameter")
		return
	}
	path := fmt.Sprintf("vpcs/%s", vpcName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// List all VPCs in the workspace
func (r *VpcService) List(ctx context.Context, opts ...option.RequestOption) (res *[]VpcListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "vpcs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete a VPC
func (r *VpcService) Delete(ctx context.Context, vpcName string, opts ...option.RequestOption) (res *VpcDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if vpcName == "" {
		err = errors.New("missing required vpcName parameter")
		return
	}
	path := fmt.Sprintf("vpcs/%s", vpcName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Virtual Private Cloud scoped to a workspace for network isolation and dedicated
// egress
type VpcNewResponse struct {
	// Common metadata fields shared by all Blaxel resources including name, labels,
	// timestamps, and ownership information
	Metadata Metadata `json:"metadata,required"`
	// VPC specification (controlplane-only, no external resources)
	Spec map[string]any `json:"spec,required"`
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
func (r VpcNewResponse) RawJSON() string { return r.JSON.raw }
func (r *VpcNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Virtual Private Cloud scoped to a workspace for network isolation and dedicated
// egress
type VpcGetResponse struct {
	// Common metadata fields shared by all Blaxel resources including name, labels,
	// timestamps, and ownership information
	Metadata Metadata `json:"metadata,required"`
	// VPC specification (controlplane-only, no external resources)
	Spec map[string]any `json:"spec,required"`
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
func (r VpcGetResponse) RawJSON() string { return r.JSON.raw }
func (r *VpcGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Virtual Private Cloud scoped to a workspace for network isolation and dedicated
// egress
type VpcListResponse struct {
	// Common metadata fields shared by all Blaxel resources including name, labels,
	// timestamps, and ownership information
	Metadata Metadata `json:"metadata,required"`
	// VPC specification (controlplane-only, no external resources)
	Spec map[string]any `json:"spec,required"`
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
func (r VpcListResponse) RawJSON() string { return r.JSON.raw }
func (r *VpcListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Virtual Private Cloud scoped to a workspace for network isolation and dedicated
// egress
type VpcDeleteResponse struct {
	// Common metadata fields shared by all Blaxel resources including name, labels,
	// timestamps, and ownership information
	Metadata Metadata `json:"metadata,required"`
	// VPC specification (controlplane-only, no external resources)
	Spec map[string]any `json:"spec,required"`
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
func (r VpcDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *VpcDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VpcNewParams struct {
	// Common metadata fields shared by all Blaxel resources including name, labels,
	// timestamps, and ownership information
	Metadata MetadataParam `json:"metadata,omitzero,required"`
	// VPC specification (controlplane-only, no external resources)
	Spec map[string]any `json:"spec,omitzero,required"`
	paramObj
}

func (r VpcNewParams) MarshalJSON() (data []byte, err error) {
	type shadow VpcNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VpcNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
