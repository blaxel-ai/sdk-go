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
	shimjson "github.com/blaxel-ai/sdk-go/internal/encoding/json"
	"github.com/blaxel-ai/sdk-go/internal/requestconfig"
	"github.com/blaxel-ai/sdk-go/option"
	"github.com/blaxel-ai/sdk-go/packages/pagination"
	"github.com/blaxel-ai/sdk-go/packages/param"
)

// ModelService contains methods and other services that help with interacting with
// the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewModelService] method instead.
type ModelService struct {
	Options []option.RequestOption
}

// NewModelService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewModelService(opts ...option.RequestOption) (r ModelService) {
	r = ModelService{}
	r.Options = opts
	return
}

// Creates a new model gateway endpoint that proxies requests to an external LLM
// provider. Requires an integration connection with valid API credentials for the
// target provider.
func (r *ModelService) New(ctx context.Context, body ModelNewParams, opts ...option.RequestOption) (res *Model, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "models"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns detailed information about a model gateway endpoint including its
// provider configuration, integration connection, and usage status.
func (r *ModelService) Get(ctx context.Context, modelName string, opts ...option.RequestOption) (res *Model, err error) {
	opts = slices.Concat(r.Options, opts)
	if modelName == "" {
		err = errors.New("missing required modelName parameter")
		return nil, err
	}
	path := fmt.Sprintf("models/%s", modelName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates a model gateway endpoint's configuration. Changes to provider settings
// or integration connection take effect immediately.
func (r *ModelService) Update(ctx context.Context, modelName string, body ModelUpdateParams, opts ...option.RequestOption) (res *Model, err error) {
	opts = slices.Concat(r.Options, opts)
	if modelName == "" {
		err = errors.New("missing required modelName parameter")
		return nil, err
	}
	path := fmt.Sprintf("models/%s", modelName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return res, err
}

// Returns model gateway endpoints configured in the workspace. Each model
// represents a proxy to an external LLM provider (OpenAI, Anthropic, etc.) with
// unified access control. Starting with API version 2026-04-28 the response is
// wrapped in `{data, meta}` and supports cursor pagination via the `cursor` and
// `limit` query parameters; older versions keep returning a bare array with all
// models.
func (r *ModelService) List(ctx context.Context, query ModelListParams, opts ...option.RequestOption) (res *pagination.CursorPage[Model], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "models"
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

// Returns model gateway endpoints configured in the workspace. Each model
// represents a proxy to an external LLM provider (OpenAI, Anthropic, etc.) with
// unified access control. Starting with API version 2026-04-28 the response is
// wrapped in `{data, meta}` and supports cursor pagination via the `cursor` and
// `limit` query parameters; older versions keep returning a bare array with all
// models.
func (r *ModelService) ListAutoPaging(ctx context.Context, query ModelListParams, opts ...option.RequestOption) *pagination.CursorPageAutoPager[Model] {
	return pagination.NewCursorPageAutoPager(r.List(ctx, query, opts...))
}

// Permanently deletes a model gateway endpoint. Any agents or applications using
// this endpoint will need to be updated to use a different model.
func (r *ModelService) Delete(ctx context.Context, modelName string, opts ...option.RequestOption) (res *Model, err error) {
	opts = slices.Concat(r.Options, opts)
	if modelName == "" {
		err = errors.New("missing required modelName parameter")
		return nil, err
	}
	path := fmt.Sprintf("models/%s", modelName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Returns revisions for a model by name.
func (r *ModelService) ListRevisions(ctx context.Context, modelName string, opts ...option.RequestOption) (res *[]RevisionMetadata, err error) {
	opts = slices.Concat(r.Options, opts)
	if modelName == "" {
		err = errors.New("missing required modelName parameter")
		return nil, err
	}
	path := fmt.Sprintf("models/%s/revisions", modelName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type ModelNewParams struct {
	// Gateway endpoint to external LLM provider APIs (OpenAI, Anthropic, etc.) with
	// unified access control, credentials management, and usage tracking.
	Model ModelParam
	paramObj
}

func (r ModelNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Model)
}
func (r *ModelNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ModelUpdateParams struct {
	// Gateway endpoint to external LLM provider APIs (OpenAI, Anthropic, etc.) with
	// unified access control, credentials management, and usage tracking.
	Model ModelParam
	paramObj
}

func (r ModelUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Model)
}
func (r *ModelUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ModelListParams struct {
	// Opaque cursor returned by a previous response's meta.nextCursor. Only valid for
	// the same query (workspace + filters); the server rejects cursors bound to a
	// different query or older than 24h. Omit on the first page.
	Cursor param.Opt[string] `query:"cursor,omitzero" json:"-"`
	// Maximum number of items to return per page. Defaults to 50, clamped to 200.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Substring search across `metadata.name`, `metadata.displayName` and labels
	// (keys + values). Trimmed and lowercased server-side; queries shorter than 2
	// characters fall back to the unfiltered listing. Bound into the cursor
	// fingerprint so a cursor opened with one query cannot be reused with another.
	// Only honoured starting on Blaxel-Version 2026-04-28.
	Q param.Opt[string] `query:"q,omitzero" json:"-"`
	// Start from a known pagination boundary. `end` is only supported for
	// `createdAt:desc` listings and returns the oldest page directly without walking
	// every cursor from the first page.
	//
	// Any of "end".
	Anchor ModelListParamsAnchor `query:"anchor,omitzero" json:"-"`
	// Sort spec, formatted as `<key>:<direction>`. Allowed values are `createdAt:desc`
	// (default), `createdAt:asc`, `name:asc`, `name:desc`. The cursor fingerprint is
	// bound to the sort, so a cursor opened with one value cannot be reused with
	// another. Only honoured starting on Blaxel-Version 2026-04-28.
	//
	// Any of "createdAt:desc", "createdAt:asc", "name:asc", "name:desc".
	Sort ModelListParamsSort `query:"sort,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [ModelListParams]'s query parameters as `url.Values`.
func (r ModelListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Start from a known pagination boundary. `end` is only supported for
// `createdAt:desc` listings and returns the oldest page directly without walking
// every cursor from the first page.
type ModelListParamsAnchor string

const (
	ModelListParamsAnchorEnd ModelListParamsAnchor = "end"
)

// Sort spec, formatted as `<key>:<direction>`. Allowed values are `createdAt:desc`
// (default), `createdAt:asc`, `name:asc`, `name:desc`. The cursor fingerprint is
// bound to the sort, so a cursor opened with one value cannot be reused with
// another. Only honoured starting on Blaxel-Version 2026-04-28.
type ModelListParamsSort string

const (
	ModelListParamsSortCreatedAtDesc ModelListParamsSort = "createdAt:desc"
	ModelListParamsSortCreatedAtAsc  ModelListParamsSort = "createdAt:asc"
	ModelListParamsSortNameAsc       ModelListParamsSort = "name:asc"
	ModelListParamsSortNameDesc      ModelListParamsSort = "name:desc"
)
