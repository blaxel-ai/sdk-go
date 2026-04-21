// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blaxel

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/blaxel-ai/sdk-go/internal/apijson"
	"github.com/blaxel-ai/sdk-go/internal/apiquery"
	"github.com/blaxel-ai/sdk-go/internal/requestconfig"
	"github.com/blaxel-ai/sdk-go/option"
	"github.com/blaxel-ai/sdk-go/packages/param"
	"github.com/blaxel-ai/sdk-go/packages/respjson"
)

// ImageSharePendingService contains methods and other services that help with
// interacting with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewImageSharePendingService] method instead.
type ImageSharePendingService struct {
	Options []option.RequestOption
}

// NewImageSharePendingService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewImageSharePendingService(opts ...option.RequestOption) (r ImageSharePendingService) {
	r = ImageSharePendingService{}
	r.Options = opts
	return
}

// Lists pending cross-account image shares targeting the caller's workspace
// (incoming) or originating from it (outgoing). Expired shares are cleaned up
// opportunistically.
func (r *ImageSharePendingService) List(ctx context.Context, query ImageSharePendingListParams, opts ...option.RequestOption) (res *[]ImageSharePendingListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "pending-image-shares"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Accepts a pending cross-account image share and copies the image metadata to the
// target workspace. Caller must be an admin of the target workspace.
func (r *ImageSharePendingService) Accept(ctx context.Context, pendingShareID string, body ImageSharePendingAcceptParams, opts ...option.RequestOption) (res *Image, err error) {
	opts = slices.Concat(r.Options, opts)
	if pendingShareID == "" {
		err = errors.New("missing required pendingShareId parameter")
		return nil, err
	}
	path := fmt.Sprintf("pending-image-shares/%s/accept", pendingShareID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Declines a pending cross-account image share. Caller must be an admin of the
// target workspace.
func (r *ImageSharePendingService) Decline(ctx context.Context, pendingShareID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if pendingShareID == "" {
		err = errors.New("missing required pendingShareId parameter")
		return err
	}
	path := fmt.Sprintf("pending-image-shares/%s/decline", pendingShareID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, nil, opts...)
	return err
}

// Rendered pending image share with source/target workspace metadata
type ImageSharePendingListResponse struct {
	// Unique identifier for the pending image share
	ID string `json:"id"`
	// Creation date
	CreatedAt string `json:"createdAt"`
	// Expiration date
	ExpiresAt string `json:"expiresAt"`
	// Whether the target workspace already has an image with the same name (potential
	// conflict)
	HasConflict bool `json:"hasConflict"`
	// Image name (repository)
	ImageName string `json:"imageName"`
	// Resource type (agent, function, sandbox, job)
	ResourceType string `json:"resourceType"`
	// User sub who initiated the share
	SharedBy string `json:"sharedBy"`
	// Email of the user who initiated the share
	SharedByEmail string `json:"sharedByEmail"`
	// Source account ID
	SourceAccountID string `json:"sourceAccountId"`
	// Source workspace name
	SourceWorkspace string `json:"sourceWorkspace"`
	// Source workspace display name
	SourceWorkspaceDisplayName string `json:"sourceWorkspaceDisplayName"`
	// Target account ID
	TargetAccountID string `json:"targetAccountId"`
	// Target workspace name
	TargetWorkspace string `json:"targetWorkspace"`
	// Target workspace display name
	TargetWorkspaceDisplayName string `json:"targetWorkspaceDisplayName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                         respjson.Field
		CreatedAt                  respjson.Field
		ExpiresAt                  respjson.Field
		HasConflict                respjson.Field
		ImageName                  respjson.Field
		ResourceType               respjson.Field
		SharedBy                   respjson.Field
		SharedByEmail              respjson.Field
		SourceAccountID            respjson.Field
		SourceWorkspace            respjson.Field
		SourceWorkspaceDisplayName respjson.Field
		TargetAccountID            respjson.Field
		TargetWorkspace            respjson.Field
		TargetWorkspaceDisplayName respjson.Field
		ExtraFields                map[string]respjson.Field
		raw                        string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ImageSharePendingListResponse) RawJSON() string { return r.JSON.raw }
func (r *ImageSharePendingListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ImageSharePendingListParams struct {
	// Direction of pending shares: "incoming" (default) or "outgoing"
	//
	// Any of "incoming", "outgoing".
	Direction ImageSharePendingListParamsDirection `query:"direction,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ImageSharePendingListParams]'s query parameters as
// `url.Values`.
func (r ImageSharePendingListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Direction of pending shares: "incoming" (default) or "outgoing"
type ImageSharePendingListParamsDirection string

const (
	ImageSharePendingListParamsDirectionIncoming ImageSharePendingListParamsDirection = "incoming"
	ImageSharePendingListParamsDirectionOutgoing ImageSharePendingListParamsDirection = "outgoing"
)

type ImageSharePendingAcceptParams struct {
	// When true, overwrite conflicting tags in the target workspace instead of
	// returning 409.
	Force param.Opt[bool] `query:"force,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ImageSharePendingAcceptParams]'s query parameters as
// `url.Values`.
func (r ImageSharePendingAcceptParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
