// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blaxel

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/blaxel-ai/sdk-go/internal/apiquery"
	"github.com/blaxel-ai/sdk-go/internal/requestconfig"
	"github.com/blaxel-ai/sdk-go/option"
	"github.com/blaxel-ai/sdk-go/packages/pagination"
	"github.com/blaxel-ai/sdk-go/packages/param"
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

// Returns a bounded page of image tags. Search by prefix or exact name.
func (r *ImageTagService) List(ctx context.Context, imageName string, params ImageTagListParams, opts ...option.RequestOption) (res *pagination.CursorPage[ImageTag], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	if params.ResourceType == "" {
		err = errors.New("missing required resourceType parameter")
		return nil, err
	}
	if imageName == "" {
		err = errors.New("missing required imageName parameter")
		return nil, err
	}
	path := fmt.Sprintf("images/%s/%s/tags", params.ResourceType, imageName)
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, params, &res, opts...)
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

// Returns a bounded page of image tags. Search by prefix or exact name.
func (r *ImageTagService) ListAutoPaging(ctx context.Context, imageName string, params ImageTagListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[ImageTag] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, imageName, params, opts...))
}

// Deletes a specific tag from a container image. The underlying image layers are
// kept if other tags reference them. Will fail if the tag is currently in use.
func (r *ImageTagService) Delete(ctx context.Context, tagName string, body ImageTagDeleteParams, opts ...option.RequestOption) (res *Image, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.ResourceType == "" {
		err = errors.New("missing required resourceType parameter")
		return nil, err
	}
	if body.ImageName == "" {
		err = errors.New("missing required imageName parameter")
		return nil, err
	}
	if tagName == "" {
		err = errors.New("missing required tagName parameter")
		return nil, err
	}
	path := fmt.Sprintf("images/%s/%s/tags/%s", body.ResourceType, body.ImageName, tagName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type ImageTagListParams struct {
	ResourceType string            `path:"resourceType" api:"required" json:"-"`
	Cursor       param.Opt[string] `query:"cursor,omitzero" json:"-"`
	Limit        param.Opt[int64]  `query:"limit,omitzero" json:"-"`
	// Exact tag name, mutually exclusive with q.
	Name param.Opt[string] `query:"name,omitzero" json:"-"`
	// Case-sensitive name prefix. Selects name order.
	Q    param.Opt[string] `query:"q,omitzero" json:"-"`
	Sort param.Opt[string] `query:"sort,omitzero" json:"-"`
	// Owner workspace for an account-shared image.
	SourceWorkspace param.Opt[string] `query:"sourceWorkspace,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ImageTagListParams]'s query parameters as `url.Values`.
func (r ImageTagListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type ImageTagDeleteParams struct {
	ResourceType string `path:"resourceType" api:"required" json:"-"`
	ImageName    string `path:"imageName" api:"required" json:"-"`
	paramObj
}
