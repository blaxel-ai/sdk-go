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
	"github.com/blaxel-ai/sdk-go/packages/pagination"
	"github.com/blaxel-ai/sdk-go/packages/param"
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
	Share   ImageShareService
}

// NewImageService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewImageService(opts ...option.RequestOption) (r ImageService) {
	r = ImageService{}
	r.Options = opts
	r.Tags = NewImageTagService(opts...)
	r.Share = NewImageShareService(opts...)
	return
}

// Builds or imports a container image without creating a deployment. Provide a
// registry image reference to download and convert an existing image, or omit
// image to receive a presigned URL for uploading source code. Registry imports can
// specify memoryMb and volumeMb for the import worker. These settings do not
// change the resources of workloads using the image.
func (r *ImageService) New(ctx context.Context, body ImageNewParams, opts ...option.RequestOption) (res *ImageNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "images"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a bounded image summary starting with API version 2026-09-22. Older
// versions return the image with all tags.
func (r *ImageService) Get(ctx context.Context, imageName string, params ImageGetParams, opts ...option.RequestOption) (res *ImageSummary, err error) {
	opts = slices.Concat(r.Options, opts)
	if params.ResourceType == "" {
		err = errors.New("missing required resourceType parameter")
		return nil, err
	}
	if imageName == "" {
		err = errors.New("missing required imageName parameter")
		return nil, err
	}
	path := fmt.Sprintf("images/%s/%s", params.ResourceType, imageName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, params, &res, opts...)
	return res, err
}

// Returns paginated image summaries starting with API version 2026-09-22. Older
// versions return a bare array including all tags.
func (r *ImageService) List(ctx context.Context, query ImageListParams, opts ...option.RequestOption) (res *pagination.CursorPage[ImageSummary], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "images"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Returns paginated image summaries starting with API version 2026-09-22. Older
// versions return a bare array including all tags.
func (r *ImageService) ListAutoPaging(ctx context.Context, query ImageListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[ImageSummary] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Deletes a container image and all its tags from the workspace registry. Will
// fail if the image is currently in use by an active deployment.
func (r *ImageService) Delete(ctx context.Context, imageName string, body ImageDeleteParams, opts ...option.RequestOption) (res *Image, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.ResourceType == "" {
		err = errors.New("missing required resourceType parameter")
		return nil, err
	}
	if imageName == "" {
		err = errors.New("missing required imageName parameter")
		return nil, err
	}
	path := fmt.Sprintf("images/%s/%s", body.ResourceType, imageName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Cleans up unused container images in the workspace registry. Only removes images
// that are not currently referenced by any active agent, function, sandbox, or job
// deployment.
func (r *ImageService) Cleanup(ctx context.Context, opts ...option.RequestOption) (res *ImageCleanupResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "images"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
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
	// Events happening on a resource deployed on Blaxel
	Events []CoreEvent `json:"events"`
	// The date and time when the image was last deployed (most recent across all
	// tags).
	LastDeployedAt string `json:"lastDeployedAt"`
	// The name of the image (repository name).
	Name string `json:"name"`
	// The resource type of the image.
	ResourceType string `json:"resourceType"`
	// If this image is shared from another workspace, this field contains the name of
	// the source workspace. Empty for non-shared images.
	SourceWorkspace string `json:"sourceWorkspace"`
	// Deployment status of a resource deployed on Blaxel
	//
	// Any of "DELETING", "TERMINATED", "FAILED", "DEACTIVATED", "DEACTIVATING",
	// "UPLOADING", "BUILDING", "DEPLOYING", "DEPLOYED", "BUILT", "ARCHIVING",
	// "ARCHIVED", "UNARCHIVING".
	Status Status `json:"status"`
	// The date and time when the image was last updated.
	UpdatedAt string `json:"updatedAt"`
	// The workspace of the image.
	Workspace string `json:"workspace"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt       respjson.Field
		DisplayName     respjson.Field
		Events          respjson.Field
		LastDeployedAt  respjson.Field
		Name            respjson.Field
		ResourceType    respjson.Field
		SourceWorkspace respjson.Field
		Status          respjson.Field
		UpdatedAt       respjson.Field
		Workspace       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
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
	Tags []ImageTag `json:"tags"`
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

type ImageSummary struct {
	Metadata ImageSummaryMetadata `json:"metadata" api:"required"`
	Spec     ImageSummarySpec     `json:"spec" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metadata    respjson.Field
		Spec        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ImageSummary) RawJSON() string { return r.JSON.raw }
func (r *ImageSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ImageSummaryMetadata struct {
	// The date and time when the image was created.
	CreatedAt string `json:"createdAt"`
	// The display name of the image (registry/workspace/repository).
	DisplayName string `json:"displayName"`
	// Events happening on a resource deployed on Blaxel
	Events []CoreEvent `json:"events"`
	// The date and time when the image was last deployed (most recent across all
	// tags).
	LastDeployedAt string `json:"lastDeployedAt"`
	// The name of the image (repository name).
	Name string `json:"name"`
	// The resource type of the image.
	ResourceType string `json:"resourceType"`
	// If this image is shared from another workspace, this field contains the name of
	// the source workspace. Empty for non-shared images.
	SourceWorkspace string `json:"sourceWorkspace"`
	// Deployment status of a resource deployed on Blaxel
	//
	// Any of "DELETING", "TERMINATED", "FAILED", "DEACTIVATED", "DEACTIVATING",
	// "UPLOADING", "BUILDING", "DEPLOYING", "DEPLOYED", "BUILT", "ARCHIVING",
	// "ARCHIVED", "UNARCHIVING".
	Status Status `json:"status"`
	// The date and time when the image was last updated.
	UpdatedAt string `json:"updatedAt"`
	// The workspace of the image.
	Workspace string `json:"workspace"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt       respjson.Field
		DisplayName     respjson.Field
		Events          respjson.Field
		LastDeployedAt  respjson.Field
		Name            respjson.Field
		ResourceType    respjson.Field
		SourceWorkspace respjson.Field
		Status          respjson.Field
		UpdatedAt       respjson.Field
		Workspace       respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ImageSummaryMetadata) RawJSON() string { return r.JSON.raw }
func (r *ImageSummaryMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ImageSummarySpec struct {
	Size     int64 `json:"size"`
	TagCount int64 `json:"tagCount"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Size        respjson.Field
		TagCount    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ImageSummarySpec) RawJSON() string { return r.JSON.raw }
func (r *ImageSummarySpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ImageTag struct {
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
func (r ImageTag) RawJSON() string { return r.JSON.raw }
func (r *ImageTag) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ImageNewResponse struct {
	// The registered image reference (only present when image was provided in request)
	Image string `json:"image"`
	// Status message
	Message string `json:"message"`
	// Name of the image
	Name string `json:"name"`
	// Resource type
	ResourceType string `json:"resourceType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Image        respjson.Field
		Message      respjson.Field
		Name         respjson.Field
		ResourceType respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ImageNewResponse) RawJSON() string { return r.JSON.raw }
func (r *ImageNewResponse) UnmarshalJSON(data []byte) error {
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

type ImageNewParams struct {
	// Name of the image to build
	Name string `json:"name" api:"required"`
	// Resource type (agent, function, sandbox, job)
	ResourceType string `json:"resourceType" api:"required"`
	// Docker configuration JSON containing credentials for the source registry.
	DockerConfig param.Opt[string] `json:"dockerConfig,omitzero"`
	// Runtime generation (e.g., mk3). Defaults to mk3 if not specified.
	Generation param.Opt[string] `json:"generation,omitzero"`
	// A pre-built Docker image reference (e.g., docker.io/myorg/myimage:latest).
	// References with a registry hostname start an asynchronous import that downloads
	// and converts the image for the resource runtime.
	Image param.Opt[string] `json:"image,omitzero"`
	// Memory for the registry import worker in MiB. Only supported when image is a
	// registry reference. When omitted, the platform default is used.
	MemoryMB param.Opt[int64] `json:"memoryMb,omitzero"`
	// Temporary scratch disk for the registry import worker in MiB. Only supported
	// when image is a registry reference. When omitted, the platform default is used.
	// Set to 0 to use memory-backed scratch. Positive values are not supported for
	// HIPAA workspaces.
	VolumeMB param.Opt[int64] `json:"volumeMb,omitzero"`
	paramObj
}

func (r ImageNewParams) MarshalJSON() (data []byte, err error) {
	type shadow ImageNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ImageNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ImageGetParams struct {
	ResourceType string `path:"resourceType" api:"required" json:"-"`
	// Owner workspace for an account-shared image.
	SourceWorkspace param.Opt[string] `query:"sourceWorkspace,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ImageGetParams]'s query parameters as `url.Values`.
func (r ImageGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ImageListParams struct {
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit  param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	// Case-sensitive name prefix. Selects name order.
	Q    param.Opt[string] `query:"q,omitzero" json:"-"`
	Sort param.Opt[string] `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ImageListParams]'s query parameters as `url.Values`.
func (r ImageListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ImageDeleteParams struct {
	ResourceType string `path:"resourceType" api:"required" json:"-"`
	paramObj
}
