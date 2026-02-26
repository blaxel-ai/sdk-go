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
	"github.com/blaxel-ai/sdk-go/packages/respjson"
)

// ImageService contains methods and other services that help with interacting with
// the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewImageService] method instead.
type ImageService struct {
	Options []option.RequestOption
	Tags    ImageTagService
}

// NewImageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewImageService(opts ...option.RequestOption) (r ImageService) {
	r = ImageService{}
	r.Options = opts
	r.Tags = NewImageTagService(opts...)
	return
}

// Returns detailed information about a container image including all available
// tags, creation dates, and size information.
func (r *ImageService) Get(ctx context.Context, imageName string, query ImageGetParams, opts ...option.RequestOption) (res *Image, err error) {
	opts = slices.Concat(r.Options, opts)
	if query.ResourceType == "" {
		err = errors.New("missing required resourceType parameter")
		return
	}
	if imageName == "" {
		err = errors.New("missing required imageName parameter")
		return
	}
	path := fmt.Sprintf("images/%s/%s", query.ResourceType, imageName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns all container images stored in the workspace registry, grouped by
// repository with their available tags. Images are created during deployments of
// agents, functions, sandboxes, and jobs.
func (r *ImageService) List(ctx context.Context, opts ...option.RequestOption) (res *[]Image, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "images"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a container image and all its tags from the workspace registry. Will
// fail if the image is currently in use by an active deployment.
func (r *ImageService) Delete(ctx context.Context, imageName string, body ImageDeleteParams, opts ...option.RequestOption) (res *Image, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.ResourceType == "" {
		err = errors.New("missing required resourceType parameter")
		return
	}
	if imageName == "" {
		err = errors.New("missing required imageName parameter")
		return
	}
	path := fmt.Sprintf("images/%s/%s", body.ResourceType, imageName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Cleans up unused container images in the workspace registry. Only removes images
// that are not currently referenced by any active agent, function, sandbox, or job
// deployment.
func (r *ImageService) Cleanup(ctx context.Context, opts ...option.RequestOption) (res *ImageCleanupResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "images"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type Image struct {
	Metadata ImageMetadata `json:"metadata" api:"required"`
	Spec     ImageSpec     `json:"spec" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metadata    respjson.Field
		Spec        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Image) RawJSON() string { return r.JSON.raw }
func (r *Image) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ImageMetadata struct {
	// The date and time when the image was created.
	CreatedAt string `json:"createdAt"`
	// The display name of the image (registry/workspace/repository).
	DisplayName string `json:"displayName"`
	// The date and time when the image was last deployed (most recent across all
	// tags).
	LastDeployedAt string `json:"lastDeployedAt"`
	// The name of the image (repository name).
	Name string `json:"name"`
	// The resource type of the image.
	ResourceType string `json:"resourceType"`
	// The date and time when the image was last updated.
	UpdatedAt string `json:"updatedAt"`
	// The workspace of the image.
	Workspace string `json:"workspace"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt      respjson.Field
		DisplayName    respjson.Field
		LastDeployedAt respjson.Field
		Name           respjson.Field
		ResourceType   respjson.Field
		UpdatedAt      respjson.Field
		Workspace      respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ImageMetadata) RawJSON() string { return r.JSON.raw }
func (r *ImageMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ImageSpec struct {
	// The size of the image in bytes.
	Size int64 `json:"size"`
	// List of tags available for this image.
	Tags []ImageSpecTag `json:"tags"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Size        respjson.Field
		Tags        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ImageSpec) RawJSON() string { return r.JSON.raw }
func (r *ImageSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ImageSpecTag struct {
	// The date and time when the tag was created.
	CreatedAt string `json:"createdAt"`
	// The name of the tag.
	Name string `json:"name"`
	// The size of the image in bytes.
	Size int64 `json:"size"`
	// The date and time when the tag was last updated.
	UpdatedAt string `json:"updatedAt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		Name        respjson.Field
		Size        respjson.Field
		UpdatedAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ImageSpecTag) RawJSON() string { return r.JSON.raw }
func (r *ImageSpecTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ImageCleanupResponse struct {
	// Number of images deleted
	Deleted int64 `json:"deleted"`
	// Result message
	Message string `json:"message"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Deleted     respjson.Field
		Message     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ImageCleanupResponse) RawJSON() string { return r.JSON.raw }
func (r *ImageCleanupResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ImageGetParams struct {
	ResourceType string `path:"resourceType" api:"required" json:"-"`
	paramObj
}

type ImageDeleteParams struct {
	ResourceType string `path:"resourceType" api:"required" json:"-"`
	paramObj
}
