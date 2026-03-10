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
func (r EgressipListResponse) RawJSON() string { return r.JSON.raw }
func (r *EgressipListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
