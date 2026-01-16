// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blaxel

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/blaxel-ai/sdk-go/internal/requestconfig"
	"github.com/blaxel-ai/sdk-go/option"
)

// ImageTagService contains methods and other services that help with interacting
// with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewImageTagService] method instead.
type ImageTagService struct {
	Options []option.RequestOption
}

// NewImageTagService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewImageTagService(opts ...option.RequestOption) (r ImageTagService) {
	r = ImageTagService{}
	r.Options = opts
	return
}

// Deletes a specific tag from a container image. The underlying image layers are
// kept if other tags reference them. Will fail if the tag is currently in use.
func (r *ImageTagService) Delete(ctx context.Context, tagName string, body ImageTagDeleteParams, opts ...option.RequestOption) (res *Image, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.ResourceType == "" {
		err = errors.New("missing required resourceType parameter")
		return
	}
	if body.ImageName == "" {
		err = errors.New("missing required imageName parameter")
		return
	}
	if tagName == "" {
		err = errors.New("missing required tagName parameter")
		return
	}
	path := fmt.Sprintf("images/%s/%s/tags/%s", body.ResourceType, body.ImageName, tagName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type ImageTagDeleteParams struct {
	ResourceType string `path:"resourceType,required" json:"-"`
	ImageName    string `path:"imageName,required" json:"-"`
	paramObj
}
