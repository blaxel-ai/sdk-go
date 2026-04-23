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

// ImageShareService contains methods and other services that help with interacting
// with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewImageShareService] method instead.
type ImageShareService struct {
	Options []option.RequestOption
}

// NewImageShareService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewImageShareService(opts ...option.RequestOption) (r ImageShareService) {
	r = ImageShareService{}
	r.Options = opts
	return
}

// Shares a container image with another workspace by copying the metadata record.
// The underlying storage (S3) data is not duplicated. For same-account targets the
// share is applied immediately. For cross-account targets, a pending image share
// is created and must be explicitly accepted by an admin of the target workspace;
// the correct target account ID must be supplied as an anti-spam measure.
func (r *ImageShareService) New(ctx context.Context, imageName string, params ImageShareNewParams, opts ...option.RequestOption) (res *Image, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ResourceType == "" {
		err = errors.New("missing required resourceType parameter")
		return nil, err
	}
	if imageName == "" {
		err = errors.New("missing required imageName parameter")
		return nil, err
	}
	path := fmt.Sprintf("images/%s/%s/share", params.ResourceType, imageName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Returns the list of workspaces that a container image is currently shared with.
func (r *ImageShareService) List(ctx context.Context, imageName string, query ImageShareListParams, opts ...option.RequestOption) (res *[]ImageShareListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.ResourceType == "" {
		err = errors.New("missing required resourceType parameter")
		return nil, err
	}
	if imageName == "" {
		err = errors.New("missing required imageName parameter")
		return nil, err
	}
	path := fmt.Sprintf("images/%s/%s/share", query.ResourceType, imageName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Revokes sharing of a container image with a target workspace. Removes the
// metadata copy from the target workspace. The source image is not affected.
func (r *ImageShareService) Delete(ctx context.Context, targetWorkspace string, body ImageShareDeleteParams, opts ...option.RequestOption) (res *Image, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.ResourceType == "" {
		err = errors.New("missing required resourceType parameter")
		return nil, err
	}
	if body.ImageName == "" {
		err = errors.New("missing required imageName parameter")
		return nil, err
	}
	if targetWorkspace == "" {
		err = errors.New("missing required targetWorkspace parameter")
		return nil, err
	}
	path := fmt.Sprintf("images/%s/%s/share/%s", body.ResourceType, body.ImageName, targetWorkspace)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type ImageShareListResponse struct {
	// ID of the account that owns the target workspace.
	AccountID string `json:"accountId" api:"required"`
	// "active" when the share is applied in the target workspace, "pending" when it is
	// awaiting accept on a cross-account share.
	Status string `json:"status" api:"required"`
	// The workspace the image is shared with.
	Workspace string `json:"workspace" api:"required"`
	// Email of the account owner for the target workspace (when available).
	AccountOwnerEmail string `json:"accountOwnerEmail"`
	// ID of the pending share record when status is "pending".
	PendingShareID string `json:"pendingShareId"`
	// Display name of the target workspace.
	WorkspaceDisplayName string `json:"workspaceDisplayName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AccountID            respjson.Field
		Status               respjson.Field
		Workspace            respjson.Field
		AccountOwnerEmail    respjson.Field
		PendingShareID       respjson.Field
		WorkspaceDisplayName respjson.Field
		ExtraFields          map[string]respjson.Field
		raw                  string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ImageShareListResponse) RawJSON() string { return r.JSON.raw }
func (r *ImageShareListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ImageShareNewParams struct {
	ResourceType string `path:"resourceType" api:"required" json:"-"`
	// Name of the workspace to share the image with
	TargetWorkspace string `json:"targetWorkspace" api:"required"`
	// Account ID of the target workspace. Required when the target workspace belongs
	// to a different account than the source workspace (anti-spam).
	TargetAccountID param.Opt[string] `json:"targetAccountId,omitzero"`
	paramObj
}

func (r ImageShareNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ImageShareNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ImageShareNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ImageShareListParams struct {
	ResourceType string `path:"resourceType" api:"required" json:"-"`
	paramObj
}

type ImageShareDeleteParams struct {
	ResourceType string `path:"resourceType" api:"required" json:"-"`
	ImageName    string `path:"imageName" api:"required" json:"-"`
	paramObj
}
