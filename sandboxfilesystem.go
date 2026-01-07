// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blaxel

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/blaxel-go/internal/apijson"
	"github.com/stainless-sdks/blaxel-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/blaxel-go/internal/encoding/json"
	"github.com/stainless-sdks/blaxel-go/internal/requestconfig"
	"github.com/stainless-sdks/blaxel-go/option"
	"github.com/stainless-sdks/blaxel-go/packages/param"
	"github.com/stainless-sdks/blaxel-go/packages/respjson"
)

// SandboxFilesystemService contains methods and other services that help with
// interacting with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSandboxFilesystemService] method instead.
type SandboxFilesystemService struct {
	Options []option.RequestOption
}

// NewSandboxFilesystemService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSandboxFilesystemService(opts ...option.RequestOption) (r SandboxFilesystemService) {
	r = SandboxFilesystemService{}
	r.Options = opts
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

// Create or update a file or directory
func (r *SandboxFilesystemService) Update(ctx context.Context, filePath string, body SandboxFilesystemUpdateParams, opts ...option.RequestOption) (res *SandboxFilesystemUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if filePath == "" {
		err = errors.New("missing required filePath parameter")
		return
	}
	path := fmt.Sprintf("filesystem/%s", filePath)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Get content of a file or listing of a directory. Use Accept header to control
// response format for files.
func (r *SandboxFilesystemService) List(ctx context.Context, filePath string, query SandboxFilesystemListParams, opts ...option.RequestOption) (res *SandboxFilesystemListResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	if filePath == "" {
		err = errors.New("missing required filePath parameter")
		return
	}
	path := fmt.Sprintf("filesystem/%s", filePath)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
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

type FilesystemDirectory struct {
	Files []FilesystemRead `json:"files,required"`
	Name  string           `json:"name,required"`
	Path  string           `json:"path,required"`
	// @name Subdirectories
	Subdirectories []FilesystemSubdirectory `json:"subdirectories,required"`
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
func (r FilesystemDirectory) RawJSON() string { return r.JSON.raw }
func (r *FilesystemDirectory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FilesystemRead struct {
	Group        string `json:"group,required"`
	LastModified string `json:"lastModified,required"`
	Name         string `json:"name,required"`
	Owner        string `json:"owner,required"`
	Path         string `json:"path,required"`
	Permissions  string `json:"permissions,required"`
	Size         int64  `json:"size,required"`
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
	Content      string `json:"content,required"`
	Group        string `json:"group,required"`
	LastModified string `json:"lastModified,required"`
	Name         string `json:"name,required"`
	Owner        string `json:"owner,required"`
	Path         string `json:"path,required"`
	Permissions  string `json:"permissions,required"`
	Size         int64  `json:"size,required"`
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

type FilesystemSubdirectory struct {
	Name string `json:"name,required"`
	Path string `json:"path,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		Path        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FilesystemSubdirectory) RawJSON() string { return r.JSON.raw }
func (r *FilesystemSubdirectory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FilesystemWriteParam struct {
	Content     param.Opt[string] `json:"content,omitzero"`
	IsDirectory param.Opt[bool]   `json:"isDirectory,omitzero"`
	Permissions param.Opt[string] `json:"permissions,omitzero"`
	paramObj
}

func (r FilesystemWriteParam) MarshalJSON() (data []byte, err error) {
	type shadow FilesystemWriteParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FilesystemWriteParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SandboxFilesystemGetResponseUnion contains all possible properties and values
// from [FilesystemDirectory], [FilesystemReadWithContent], [io.Reader].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFile]
type SandboxFilesystemGetResponseUnion struct {
	// This field will be present if the value is a [io.Reader] instead of an object.
	OfFile io.Reader `json:",inline"`
	// This field is from variant [FilesystemDirectory].
	Files []FilesystemRead `json:"files"`
	Name  string           `json:"name"`
	Path  string           `json:"path"`
	// This field is from variant [FilesystemDirectory].
	Subdirectories []FilesystemSubdirectory `json:"subdirectories"`
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
		OfFile         respjson.Field
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

func (u SandboxFilesystemGetResponseUnion) AsFilesystemDirectory() (v FilesystemDirectory) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SandboxFilesystemGetResponseUnion) AsFilesystemReadWithContent() (v FilesystemReadWithContent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SandboxFilesystemGetResponseUnion) AsFile() (v io.Reader) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SandboxFilesystemGetResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *SandboxFilesystemGetResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SandboxFilesystemUpdateResponse struct {
	Message string `json:"message,required"`
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
func (r SandboxFilesystemUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *SandboxFilesystemUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SandboxFilesystemListResponseUnion contains all possible properties and values
// from [FilesystemDirectory], [FilesystemReadWithContent], [io.Reader].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFile]
type SandboxFilesystemListResponseUnion struct {
	// This field will be present if the value is a [io.Reader] instead of an object.
	OfFile io.Reader `json:",inline"`
	// This field is from variant [FilesystemDirectory].
	Files []FilesystemRead `json:"files"`
	Name  string           `json:"name"`
	Path  string           `json:"path"`
	// This field is from variant [FilesystemDirectory].
	Subdirectories []FilesystemSubdirectory `json:"subdirectories"`
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
		OfFile         respjson.Field
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

func (u SandboxFilesystemListResponseUnion) AsFilesystemDirectory() (v FilesystemDirectory) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SandboxFilesystemListResponseUnion) AsFilesystemReadWithContent() (v FilesystemReadWithContent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SandboxFilesystemListResponseUnion) AsFile() (v io.Reader) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SandboxFilesystemListResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *SandboxFilesystemListResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SandboxFilesystemDeleteResponse struct {
	Message string `json:"message,required"`
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

type SandboxFilesystemUpdateParams struct {
	FilesystemWrite FilesystemWriteParam
	paramObj
}

func (r SandboxFilesystemUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.FilesystemWrite)
}
func (r *SandboxFilesystemUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.FilesystemWrite)
}

type SandboxFilesystemListParams struct {
	// Force download mode for files
	Download param.Opt[bool] `query:"download,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SandboxFilesystemListParams]'s query parameters as
// `url.Values`.
func (r SandboxFilesystemListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
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
