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

// SandboxFService contains methods and other services that help with interacting
// with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSandboxFService] method instead.
type SandboxFService struct {
	Options []option.RequestOption
}

// NewSandboxFService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSandboxFService(opts ...option.RequestOption) (r SandboxFService) {
	r = SandboxFService{}
	r.Options = opts
	return
}

// Get content of a file or listing of a directory. Use Accept header to control
// response format for files.
func (r *SandboxFService) Get(ctx context.Context, path string, query SandboxFGetParams, opts ...option.RequestOption) (res *SandboxFGetResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	if path == "" {
		err = errors.New("missing required path parameter")
		return
	}
	path := fmt.Sprintf("filesystem/%s", path)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Create or update a file or directory
func (r *SandboxFService) Update(ctx context.Context, path string, body SandboxFUpdateParams, opts ...option.RequestOption) (res *SandboxFUpdateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if path == "" {
		err = errors.New("missing required path parameter")
		return
	}
	path := fmt.Sprintf("filesystem/%s", path)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Get content of a file or listing of a directory. Use Accept header to control
// response format for files.
func (r *SandboxFService) List(ctx context.Context, path string, query SandboxFListParams, opts ...option.RequestOption) (res *SandboxFListResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	if path == "" {
		err = errors.New("missing required path parameter")
		return
	}
	path := fmt.Sprintf("filesystem/%s", path)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Delete a file or directory
func (r *SandboxFService) Delete(ctx context.Context, path string, body SandboxFDeleteParams, opts ...option.RequestOption) (res *SandboxFDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if path == "" {
		err = errors.New("missing required path parameter")
		return
	}
	path := fmt.Sprintf("filesystem/%s", path)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, body, &res, opts...)
	return
}

type Directory struct {
	Files []File `json:"files,required"`
	Name  string `json:"name,required"`
	Path  string `json:"path,required"`
	// @name Subdirectories
	Subdirectories []Subdirectory `json:"subdirectories,required"`
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

type File struct {
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
func (r File) RawJSON() string { return r.JSON.raw }
func (r *File) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FileRequestParam struct {
	Content     param.Opt[string] `json:"content,omitzero"`
	IsDirectory param.Opt[bool]   `json:"isDirectory,omitzero"`
	Permissions param.Opt[string] `json:"permissions,omitzero"`
	paramObj
}

func (r FileRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow FileRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *FileRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Subdirectory struct {
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
func (r Subdirectory) RawJSON() string { return r.JSON.raw }
func (r *Subdirectory) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SandboxFGetResponseUnion contains all possible properties and values from
// [Directory], [SandboxFGetResponseFileWithContent], [io.Reader].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFile]
type SandboxFGetResponseUnion struct {
	// This field will be present if the value is a [io.Reader] instead of an object.
	OfFile io.Reader `json:",inline"`
	// This field is from variant [Directory].
	Files []File `json:"files"`
	Name  string `json:"name"`
	Path  string `json:"path"`
	// This field is from variant [Directory].
	Subdirectories []Subdirectory `json:"subdirectories"`
	// This field is from variant [SandboxFGetResponseFileWithContent].
	Content string `json:"content"`
	// This field is from variant [SandboxFGetResponseFileWithContent].
	Group string `json:"group"`
	// This field is from variant [SandboxFGetResponseFileWithContent].
	LastModified string `json:"lastModified"`
	// This field is from variant [SandboxFGetResponseFileWithContent].
	Owner string `json:"owner"`
	// This field is from variant [SandboxFGetResponseFileWithContent].
	Permissions string `json:"permissions"`
	// This field is from variant [SandboxFGetResponseFileWithContent].
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

func (u SandboxFGetResponseUnion) AsDirectory() (v Directory) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SandboxFGetResponseUnion) AsSandboxFGetResponseFileWithContent() (v SandboxFGetResponseFileWithContent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SandboxFGetResponseUnion) AsFile() (v io.Reader) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SandboxFGetResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *SandboxFGetResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SandboxFGetResponseFileWithContent struct {
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
func (r SandboxFGetResponseFileWithContent) RawJSON() string { return r.JSON.raw }
func (r *SandboxFGetResponseFileWithContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SandboxFUpdateResponse struct {
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
func (r SandboxFUpdateResponse) RawJSON() string { return r.JSON.raw }
func (r *SandboxFUpdateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// SandboxFListResponseUnion contains all possible properties and values from
// [Directory], [SandboxFListResponseFileWithContent], [io.Reader].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfFile]
type SandboxFListResponseUnion struct {
	// This field will be present if the value is a [io.Reader] instead of an object.
	OfFile io.Reader `json:",inline"`
	// This field is from variant [Directory].
	Files []File `json:"files"`
	Name  string `json:"name"`
	Path  string `json:"path"`
	// This field is from variant [Directory].
	Subdirectories []Subdirectory `json:"subdirectories"`
	// This field is from variant [SandboxFListResponseFileWithContent].
	Content string `json:"content"`
	// This field is from variant [SandboxFListResponseFileWithContent].
	Group string `json:"group"`
	// This field is from variant [SandboxFListResponseFileWithContent].
	LastModified string `json:"lastModified"`
	// This field is from variant [SandboxFListResponseFileWithContent].
	Owner string `json:"owner"`
	// This field is from variant [SandboxFListResponseFileWithContent].
	Permissions string `json:"permissions"`
	// This field is from variant [SandboxFListResponseFileWithContent].
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

func (u SandboxFListResponseUnion) AsDirectory() (v Directory) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SandboxFListResponseUnion) AsSandboxFListResponseFileWithContent() (v SandboxFListResponseFileWithContent) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u SandboxFListResponseUnion) AsFile() (v io.Reader) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u SandboxFListResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *SandboxFListResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SandboxFListResponseFileWithContent struct {
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
func (r SandboxFListResponseFileWithContent) RawJSON() string { return r.JSON.raw }
func (r *SandboxFListResponseFileWithContent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SandboxFDeleteResponse struct {
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
func (r SandboxFDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *SandboxFDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SandboxFGetParams struct {
	// Force download mode for files
	Download param.Opt[bool] `query:"download,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SandboxFGetParams]'s query parameters as `url.Values`.
func (r SandboxFGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SandboxFUpdateParams struct {
	FileRequest FileRequestParam
	paramObj
}

func (r SandboxFUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.FileRequest)
}
func (r *SandboxFUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.FileRequest)
}

type SandboxFListParams struct {
	// Force download mode for files
	Download param.Opt[bool] `query:"download,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SandboxFListParams]'s query parameters as `url.Values`.
func (r SandboxFListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type SandboxFDeleteParams struct {
	// Delete directory recursively
	Recursive param.Opt[bool] `query:"recursive,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SandboxFDeleteParams]'s query parameters as `url.Values`.
func (r SandboxFDeleteParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
