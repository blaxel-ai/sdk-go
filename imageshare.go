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
// The underlying storage (S3) data is not duplicated. The target workspace must
// belong to the same account.
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
func (r *ImageShareService) List(ctx context.Context, imageName string, query ImageShareListParams, opts ...option.RequestOption) (res *[]string, err error) {
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

type ImageShareNewParams struct {
	ResourceType string `path:"resourceType" api:"required" json:"-"`
	// Name of the workspace to share the image with
	TargetWorkspace string `json:"targetWorkspace" api:"required"`
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
