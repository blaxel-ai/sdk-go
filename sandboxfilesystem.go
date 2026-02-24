// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blaxel

import (
	"context"
	"encoding/json"
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
	"github.com/blaxel-ai/sdk-go/packages/param"
	"github.com/blaxel-ai/sdk-go/packages/respjson"
)

// SandboxFilesystemService contains methods and other services that help with
// interacting with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSandboxFilesystemService] method instead.
type SandboxFilesystemService struct {
	Options   []option.RequestOption
	Watch     SandboxFilesystemWatchService
	Multipart SandboxFilesystemMultipartService
}

// NewSandboxFilesystemService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSandboxFilesystemService(opts ...option.RequestOption) (r SandboxFilesystemService) {
	r = SandboxFilesystemService{}
	r.Options = opts
	r.Watch = NewSandboxFilesystemWatchService(opts...)
	r.Multipart = NewSandboxFilesystemMultipartService(opts...)
	return
}

// Delete a file or directory
func (r *SandboxFilesystemService) Delete(ctx context.Context, filePath string, body SandboxFilesystemDeleteParams, opts ...option.RequestOption) (res *SandboxFilesystemDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if filePath == "" {
		err = errors.New("missing required filePath parameter")
		return
	}
	path := fmt.Sprintf("filesystem/%s", filePath)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Searches for text content inside files using ripgrep. Returns matching lines
// with context.
func (r *SandboxFilesystemService) ContentSearch(ctx context.Context, filePath string, query SandboxFilesystemContentSearchParams, opts ...option.RequestOption) (res *ContentSearchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if filePath == "" {
		err = errors.New("missing required filePath parameter")
		return
	}
	path := fmt.Sprintf("filesystem-content-search/%s", filePath)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a directory tree recursively
func (r *SandboxFilesystemService) DeleteTree(ctx context.Context, filePath string, body SandboxFilesystemDeleteTreeParams, opts ...option.RequestOption) (res *SandboxFilesystemDeleteTreeResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if filePath == "" {
		err = errors.New("missing required filePath parameter")
		return
	}
	path := fmt.Sprintf("filesystem/tree/%s", filePath)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

// Finds files and directories using the find command.
func (r *SandboxFilesystemService) Find(ctx context.Context, filePath string, query SandboxFilesystemFindParams, opts ...option.RequestOption) (res *FindResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if filePath == "" {
		err = errors.New("missing required filePath parameter")
		return
	}
	path := fmt.Sprintf("filesystem-find/%s", filePath)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get content of a file or listing of a directory. Use Accept header to control
// response format for files.
func (r *SandboxFilesystemService) Get(ctx context.Context, filePath string, query SandboxFilesystemGetParams, opts ...option.RequestOption) (res *SandboxFilesystemGetResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	if filePath == "" {
		err = errors.New("missing required filePath parameter")
		return
	}
	path := fmt.Sprintf("filesystem/%s", filePath)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Get a recursive directory tree structure starting from the specified path
func (r *SandboxFilesystemService) GetTree(ctx context.Context, filePath string, opts ...option.RequestOption) (res *SandboxFilesystemGetTreeResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	if filePath == "" {
		err = errors.New("missing required filePath parameter")
		return
	}
	path := fmt.Sprintf("filesystem/tree/%s", filePath)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Performs fuzzy search on filesystem paths using fuzzy matching algorithm.
// Optimized alternative to find and grep commands.
func (r *SandboxFilesystemService) Search(ctx context.Context, filePath string, query SandboxFilesystemSearchParams, opts ...option.RequestOption) (res *FuzzySearchResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if filePath == "" {
		err = errors.New("missing required filePath parameter")
		return
	}
	path := fmt.Sprintf("filesystem-search/%s", filePath)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Create or update a file or directory
func (r *SandboxFilesystemService) Write(ctx context.Context, filePath string, body SandboxFilesystemWriteParams, opts ...option.RequestOption) (res *SandboxFilesystemWriteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if filePath == "" {
		err = errors.New("missing required filePath parameter")
		return
	}
	path := fmt.Sprintf("filesystem/%s", filePath)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Create or update multiple files within a directory tree structure
func (r *SandboxFilesystemService) WriteTree(ctx context.Context, filePath string, body SandboxFilesystemWriteTreeParams, opts ...option.RequestOption) (res *SandboxFilesystemWriteTreeResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	if filePath == "" {
		err = errors.New("missing required filePath parameter")
		return
	}
	path := fmt.Sprintf("filesystem/tree/%s", filePath)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type ContentSearchMatch struct {
	Column  int64  `json:"column" api:"required"`
	Line    int64  `json:"line" api:"required"`
	Path    string `json:"path" api:"required"`
	Text    string `json:"text" api:"required"`
	Context string `json:"context"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Column      respjson.Field
		Line        respjson.Field
		Path        respjson.Field
		Text        respjson.Field
		Context     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContentSearchMatch) RawJSON() string { return r.JSON.raw }
func (r *ContentSearchMatch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ContentSearchResponse struct {
	Matches []ContentSearchMatch `json:"matches" api:"required"`
	Query   string               `json:"query" api:"required"`
	Total   int64                `json:"total" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Matches     respjson.Field
		Query       respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ContentSearchResponse) RawJSON() string { return r.JSON.raw }
func (r *ContentSearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Directory struct {
	Files []FilesystemRead `json:"files" api:"required"`
	Name  string           `json:"name" api:"required"`
	Path  string           `json:"path" api:"required"`
	// @name Subdirectories
	Subdirectories []Subdirectory `json:"subdirectories" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Files          respjson.Field
		Name           respjson.Field
		Path           respjson.Field
		Subdirectories respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Directory) RawJSON() string { return r.JSON.raw }
func (r *Directory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FilesystemRead struct {
	Group        string `json:"group" api:"required"`
	LastModified string `json:"lastModified" api:"required"`
	Name         string `json:"name" api:"required"`
	Owner        string `json:"owner" api:"required"`
	Path         string `json:"path" api:"required"`
	Permissions  string `json:"permissions" api:"required"`
	Size         int64  `json:"size" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Group        respjson.Field
		LastModified respjson.Field
		Name         respjson.Field
		Owner        respjson.Field
		Path         respjson.Field
		Permissions  respjson.Field
		Size         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FilesystemRead) RawJSON() string { return r.JSON.raw }
func (r *FilesystemRead) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FilesystemReadWithContent struct {
	Content      string `json:"content" api:"required"`
	Group        string `json:"group" api:"required"`
	LastModified string `json:"lastModified" api:"required"`
	Name         string `json:"name" api:"required"`
	Owner        string `json:"owner" api:"required"`
	Path         string `json:"path" api:"required"`
	Permissions  string `json:"permissions" api:"required"`
	Size         int64  `json:"size" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content      respjson.Field
		Group        respjson.Field
		LastModified respjson.Field
		Name         respjson.Field
		Owner        respjson.Field
		Path         respjson.Field
		Permissions  respjson.Field
		Size         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FilesystemReadWithContent) RawJSON() string { return r.JSON.raw }
func (r *FilesystemReadWithContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FilesystemWriteRequestParam struct {
	Content     param.Opt[string] `json:"content,omitzero"`
	IsDirectory param.Opt[bool]   `json:"isDirectory,omitzero"`
	Permissions param.Opt[string] `json:"permissions,omitzero"`
	paramObj
}

func (r FilesystemWriteRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow FilesystemWriteRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FilesystemWriteRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FindMatch struct {
	Path string `json:"path" api:"required"`
	// "file" or "directory"
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Path        respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FindMatch) RawJSON() string { return r.JSON.raw }
func (r *FindMatch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FindResponse struct {
	Matches []FindMatch `json:"matches" api:"required"`
	Total   int64       `json:"total" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Matches     respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FindResponse) RawJSON() string { return r.JSON.raw }
func (r *FindResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FuzzySearchMatch struct {
	Path  string `json:"path" api:"required"`
	Score int64  `json:"score" api:"required"`
	// "file" or "directory"
	Type string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Path        respjson.Field
		Score       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FuzzySearchMatch) RawJSON() string { return r.JSON.raw }
func (r *FuzzySearchMatch) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FuzzySearchResponse struct {
	Matches []FuzzySearchMatch `json:"matches" api:"required"`
	Total   int64              `json:"total" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Matches     respjson.Field
		Total       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FuzzySearchResponse) RawJSON() string { return r.JSON.raw }
func (r *FuzzySearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Subdirectory struct {
	Name string `json:"name" api:"required"`
	Path string `json:"path" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Path        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Subdirectory) RawJSON() string { return r.JSON.raw }
func (r *Subdirectory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TreeRequestParam struct {
	Files map[string]string `json:"files,omitzero"`
	paramObj
}

func (r TreeRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow TreeRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TreeRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SandboxFilesystemDeleteResponse struct {
	Message string `json:"message" api:"required"`
	Path    string `json:"path"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Path        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SandboxFilesystemDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *SandboxFilesystemDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SandboxFilesystemDeleteTreeResponse struct {
	Message string `json:"message" api:"required"`
	Path    string `json:"path"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Path        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SandboxFilesystemDeleteTreeResponse) RawJSON() string { return r.JSON.raw }
func (r *SandboxFilesystemDeleteTreeResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SandboxFilesystemGetResponseUnion contains all possible properties and values
// from [Directory], [FilesystemReadWithContent].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SandboxFilesystemGetResponseUnion struct {
	// This field is from variant [Directory].
	Files []FilesystemRead `json:"files"`
	Name  string           `json:"name"`
	Path  string           `json:"path"`
	// This field is from variant [Directory].
	Subdirectories []Subdirectory `json:"subdirectories"`
	// This field is from variant [FilesystemReadWithContent].
	Content string `json:"content"`
	// This field is from variant [FilesystemReadWithContent].
	Group string `json:"group"`
	// This field is from variant [FilesystemReadWithContent].
	LastModified string `json:"lastModified"`
	// This field is from variant [FilesystemReadWithContent].
	Owner string `json:"owner"`
	// This field is from variant [FilesystemReadWithContent].
	Permissions string `json:"permissions"`
	// This field is from variant [FilesystemReadWithContent].
	Size int64 `json:"size"`
	JSON struct {
		Files          respjson.Field
		Name           respjson.Field
		Path           respjson.Field
		Subdirectories respjson.Field
		Content        respjson.Field
		Group          respjson.Field
		LastModified   respjson.Field
		Owner          respjson.Field
		Permissions    respjson.Field
		Size           respjson.Field
		raw            string
	} `json:"-"`
}

func (u SandboxFilesystemGetResponseUnion) AsDirectory() (v Directory) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SandboxFilesystemGetResponseUnion) AsFilesystemReadWithContent() (v FilesystemReadWithContent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SandboxFilesystemGetResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *SandboxFilesystemGetResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SandboxFilesystemGetTreeResponseUnion contains all possible properties and
// values from [Directory], [FilesystemReadWithContent].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SandboxFilesystemGetTreeResponseUnion struct {
	// This field is from variant [Directory].
	Files []FilesystemRead `json:"files"`
	Name  string           `json:"name"`
	Path  string           `json:"path"`
	// This field is from variant [Directory].
	Subdirectories []Subdirectory `json:"subdirectories"`
	// This field is from variant [FilesystemReadWithContent].
	Content string `json:"content"`
	// This field is from variant [FilesystemReadWithContent].
	Group string `json:"group"`
	// This field is from variant [FilesystemReadWithContent].
	LastModified string `json:"lastModified"`
	// This field is from variant [FilesystemReadWithContent].
	Owner string `json:"owner"`
	// This field is from variant [FilesystemReadWithContent].
	Permissions string `json:"permissions"`
	// This field is from variant [FilesystemReadWithContent].
	Size int64 `json:"size"`
	JSON struct {
		Files          respjson.Field
		Name           respjson.Field
		Path           respjson.Field
		Subdirectories respjson.Field
		Content        respjson.Field
		Group          respjson.Field
		LastModified   respjson.Field
		Owner          respjson.Field
		Permissions    respjson.Field
		Size           respjson.Field
		raw            string
	} `json:"-"`
}

func (u SandboxFilesystemGetTreeResponseUnion) AsDirectory() (v Directory) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SandboxFilesystemGetTreeResponseUnion) AsFilesystemReadWithContent() (v FilesystemReadWithContent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SandboxFilesystemGetTreeResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *SandboxFilesystemGetTreeResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SandboxFilesystemWriteResponse struct {
	Message string `json:"message" api:"required"`
	Path    string `json:"path"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Path        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SandboxFilesystemWriteResponse) RawJSON() string { return r.JSON.raw }
func (r *SandboxFilesystemWriteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SandboxFilesystemWriteTreeResponseUnion contains all possible properties and
// values from [Directory], [FilesystemReadWithContent].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type SandboxFilesystemWriteTreeResponseUnion struct {
	// This field is from variant [Directory].
	Files []FilesystemRead `json:"files"`
	Name  string           `json:"name"`
	Path  string           `json:"path"`
	// This field is from variant [Directory].
	Subdirectories []Subdirectory `json:"subdirectories"`
	// This field is from variant [FilesystemReadWithContent].
	Content string `json:"content"`
	// This field is from variant [FilesystemReadWithContent].
	Group string `json:"group"`
	// This field is from variant [FilesystemReadWithContent].
	LastModified string `json:"lastModified"`
	// This field is from variant [FilesystemReadWithContent].
	Owner string `json:"owner"`
	// This field is from variant [FilesystemReadWithContent].
	Permissions string `json:"permissions"`
	// This field is from variant [FilesystemReadWithContent].
	Size int64 `json:"size"`
	JSON struct {
		Files          respjson.Field
		Name           respjson.Field
		Path           respjson.Field
		Subdirectories respjson.Field
		Content        respjson.Field
		Group          respjson.Field
		LastModified   respjson.Field
		Owner          respjson.Field
		Permissions    respjson.Field
		Size           respjson.Field
		raw            string
	} `json:"-"`
}

func (u SandboxFilesystemWriteTreeResponseUnion) AsDirectory() (v Directory) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SandboxFilesystemWriteTreeResponseUnion) AsFilesystemReadWithContent() (v FilesystemReadWithContent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SandboxFilesystemWriteTreeResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *SandboxFilesystemWriteTreeResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SandboxFilesystemDeleteParams struct {
	// Delete directory recursively
	Recursive param.Opt[bool] `query:"recursive,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SandboxFilesystemDeleteParams]'s query parameters as
// `url.Values`.
func (r SandboxFilesystemDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SandboxFilesystemContentSearchParams struct {
	// Text to search for
	Query string `query:"query" api:"required" json:"-"`
	// Case sensitive search (default: false)
	CaseSensitive param.Opt[bool] `query:"caseSensitive,omitzero" json:"-"`
	// Comma-separated directory names to skip (default:
	// node_modules,vendor,.git,dist,build,target,**pycache**,.venv,.next,coverage)
	ExcludeDirs param.Opt[string] `query:"excludeDirs,omitzero" json:"-"`
	// File pattern to include (e.g., \*.go)
	FilePattern param.Opt[string] `query:"filePattern,omitzero" json:"-"`
	// Maximum number of results to return (default: 100)
	MaxResults param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SandboxFilesystemContentSearchParams]'s query parameters as
// `url.Values`.
func (r SandboxFilesystemContentSearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SandboxFilesystemDeleteTreeParams struct {
	// Delete directory recursively
	Recursive param.Opt[bool] `query:"recursive,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SandboxFilesystemDeleteTreeParams]'s query parameters as
// `url.Values`.
func (r SandboxFilesystemDeleteTreeParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SandboxFilesystemFindParams struct {
	// Comma-separated directory names to skip (default:
	// node_modules,vendor,.git,dist,build,target,**pycache**,.venv,.next,coverage).
	// Use empty string to skip no directories.
	ExcludeDirs param.Opt[string] `query:"excludeDirs,omitzero" json:"-"`
	// Exclude hidden files and directories (default: true)
	ExcludeHidden param.Opt[bool] `query:"excludeHidden,omitzero" json:"-"`
	// Maximum number of results to return (default: 20). If set to 0, all results will
	// be returned.
	MaxResults param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	// Comma-separated file patterns to include (e.g., _.go,_.js)
	Patterns param.Opt[string] `query:"patterns,omitzero" json:"-"`
	// Type of search (file or directory)
	Type param.Opt[string] `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SandboxFilesystemFindParams]'s query parameters as
// `url.Values`.
func (r SandboxFilesystemFindParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SandboxFilesystemGetParams struct {
	// Force download mode for files
	Download param.Opt[bool] `query:"download,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SandboxFilesystemGetParams]'s query parameters as
// `url.Values`.
func (r SandboxFilesystemGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SandboxFilesystemSearchParams struct {
	// Comma-separated directory names to skip (default:
	// node_modules,vendor,.git,dist,build,target,**pycache**,.venv,.next,coverage).
	// Use empty string to skip no directories.
	ExcludeDirs param.Opt[string] `query:"excludeDirs,omitzero" json:"-"`
	// Exclude hidden files and directories (default: true)
	ExcludeHidden param.Opt[bool] `query:"excludeHidden,omitzero" json:"-"`
	// Maximum number of results to return (default: 20)
	MaxResults param.Opt[int64] `query:"maxResults,omitzero" json:"-"`
	// Comma-separated file patterns to include (e.g., _.go,_.js)
	Patterns param.Opt[string] `query:"patterns,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SandboxFilesystemSearchParams]'s query parameters as
// `url.Values`.
func (r SandboxFilesystemSearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SandboxFilesystemWriteParams struct {
	FilesystemWriteRequest FilesystemWriteRequestParam
	paramObj
}

func (r SandboxFilesystemWriteParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.FilesystemWriteRequest)
}
func (r *SandboxFilesystemWriteParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.FilesystemWriteRequest)
}

type SandboxFilesystemWriteTreeParams struct {
	TreeRequest TreeRequestParam
	paramObj
}

func (r SandboxFilesystemWriteTreeParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.TreeRequest)
}
func (r *SandboxFilesystemWriteTreeParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.TreeRequest)
}
