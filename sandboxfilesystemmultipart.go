// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blaxel

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"slices"

	"github.com/blaxel-ai/sdk-go/internal/apiform"
	"github.com/blaxel-ai/sdk-go/internal/apijson"
	"github.com/blaxel-ai/sdk-go/internal/apiquery"
	shimjson "github.com/blaxel-ai/sdk-go/internal/encoding/json"
	"github.com/blaxel-ai/sdk-go/internal/requestconfig"
	"github.com/blaxel-ai/sdk-go/option"
	"github.com/blaxel-ai/sdk-go/packages/param"
	"github.com/blaxel-ai/sdk-go/packages/respjson"
)

// SandboxFilesystemMultipartService contains methods and other services that help
// with interacting with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSandboxFilesystemMultipartService] method instead.
type SandboxFilesystemMultipartService struct {
	Options []option.RequestOption
}

// NewSandboxFilesystemMultipartService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewSandboxFilesystemMultipartService(opts ...option.RequestOption) (r SandboxFilesystemMultipartService) {
	r = SandboxFilesystemMultipartService{}
	r.Options = opts
	return
}

// List all active multipart uploads
func (r *SandboxFilesystemMultipartService) List(ctx context.Context, opts ...option.RequestOption) (res *ListUploadsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "filesystem-multipart"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Abort a multipart upload and clean up all parts
func (r *SandboxFilesystemMultipartService) Abort(ctx context.Context, uploadID string, opts ...option.RequestOption) (res *SandboxFilesystemMultipartAbortResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if uploadID == "" {
		err = errors.New("missing required uploadId parameter")
		return
	}
	path := fmt.Sprintf("filesystem-multipart/%s/abort", uploadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Complete a multipart upload by assembling all parts
func (r *SandboxFilesystemMultipartService) Complete(ctx context.Context, uploadID string, body SandboxFilesystemMultipartCompleteParams, opts ...option.RequestOption) (res *SandboxFilesystemMultipartCompleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if uploadID == "" {
		err = errors.New("missing required uploadId parameter")
		return
	}
	path := fmt.Sprintf("filesystem-multipart/%s/complete", uploadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Initiate a multipart upload session for a file
func (r *SandboxFilesystemMultipartService) Initiate(ctx context.Context, filePath string, body SandboxFilesystemMultipartInitiateParams, opts ...option.RequestOption) (res *InitiateResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if filePath == "" {
		err = errors.New("missing required filePath parameter")
		return
	}
	path := fmt.Sprintf("filesystem-multipart/initiate/%s", filePath)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// List all uploaded parts for a multipart upload
func (r *SandboxFilesystemMultipartService) ListParts(ctx context.Context, uploadID string, opts ...option.RequestOption) (res *ListPartsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if uploadID == "" {
		err = errors.New("missing required uploadId parameter")
		return
	}
	path := fmt.Sprintf("filesystem-multipart/%s/parts", uploadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Upload a single part of a multipart upload
func (r *SandboxFilesystemMultipartService) UploadPart(ctx context.Context, uploadID string, params SandboxFilesystemMultipartUploadPartParams, opts ...option.RequestOption) (res *UploadPartResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if uploadID == "" {
		err = errors.New("missing required uploadId parameter")
		return
	}
	path := fmt.Sprintf("filesystem-multipart/%s/part", uploadID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

type CompleteRequestParam struct {
	Parts []PartInfoParam `json:"parts,omitzero"`
	paramObj
}

func (r CompleteRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow CompleteRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CompleteRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InitiateRequestParam struct {
	Permissions param.Opt[string] `json:"permissions,omitzero"`
	paramObj
}

func (r InitiateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow InitiateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *InitiateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type InitiateResponse struct {
	Path     string `json:"path"`
	UploadID string `json:"uploadId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Path        respjson.Field
		UploadID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r InitiateResponse) RawJSON() string { return r.JSON.raw }
func (r *InitiateResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListPartsResponse struct {
	Parts    []ListPartsResponsePart `json:"parts"`
	UploadID string                  `json:"uploadId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Parts       respjson.Field
		UploadID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListPartsResponse) RawJSON() string { return r.JSON.raw }
func (r *ListPartsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListPartsResponsePart struct {
	Etag       string `json:"etag"`
	PartNumber int64  `json:"partNumber"`
	Size       int64  `json:"size"`
	UploadedAt string `json:"uploadedAt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Etag        respjson.Field
		PartNumber  respjson.Field
		Size        respjson.Field
		UploadedAt  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListPartsResponsePart) RawJSON() string { return r.JSON.raw }
func (r *ListPartsResponsePart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListUploadsResponse struct {
	Uploads []ListUploadsResponseUpload `json:"uploads"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Uploads     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListUploadsResponse) RawJSON() string { return r.JSON.raw }
func (r *ListUploadsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListUploadsResponseUpload struct {
	InitiatedAt string                                   `json:"initiatedAt"`
	Parts       map[string]ListUploadsResponseUploadPart `json:"parts"`
	Path        string                                   `json:"path"`
	Permissions int64                                    `json:"permissions"`
	UploadID    string                                   `json:"uploadId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InitiatedAt respjson.Field
		Parts       respjson.Field
		Path        respjson.Field
		Permissions respjson.Field
		UploadID    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListUploadsResponseUpload) RawJSON() string { return r.JSON.raw }
func (r *ListUploadsResponseUpload) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ListUploadsResponseUploadPart struct {
	Etag       string `json:"etag"`
	PartNumber int64  `json:"partNumber"`
	Size       int64  `json:"size"`
	UploadedAt string `json:"uploadedAt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Etag        respjson.Field
		PartNumber  respjson.Field
		Size        respjson.Field
		UploadedAt  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ListUploadsResponseUploadPart) RawJSON() string { return r.JSON.raw }
func (r *ListUploadsResponseUploadPart) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type PartInfoParam struct {
	Etag       param.Opt[string] `json:"etag,omitzero"`
	PartNumber param.Opt[int64]  `json:"partNumber,omitzero"`
	paramObj
}

func (r PartInfoParam) MarshalJSON() (data []byte, err error) {
	type shadow PartInfoParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *PartInfoParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UploadPartResponse struct {
	Etag       string `json:"etag"`
	PartNumber int64  `json:"partNumber"`
	Size       int64  `json:"size"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Etag        respjson.Field
		PartNumber  respjson.Field
		Size        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UploadPartResponse) RawJSON() string { return r.JSON.raw }
func (r *UploadPartResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SandboxFilesystemMultipartAbortResponse struct {
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
func (r SandboxFilesystemMultipartAbortResponse) RawJSON() string { return r.JSON.raw }
func (r *SandboxFilesystemMultipartAbortResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SandboxFilesystemMultipartCompleteResponse struct {
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
func (r SandboxFilesystemMultipartCompleteResponse) RawJSON() string { return r.JSON.raw }
func (r *SandboxFilesystemMultipartCompleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SandboxFilesystemMultipartCompleteParams struct {
	CompleteRequest CompleteRequestParam
	paramObj
}

func (r SandboxFilesystemMultipartCompleteParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CompleteRequest)
}
func (r *SandboxFilesystemMultipartCompleteParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CompleteRequest)
}

type SandboxFilesystemMultipartInitiateParams struct {
	InitiateRequest InitiateRequestParam
	paramObj
}

func (r SandboxFilesystemMultipartInitiateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.InitiateRequest)
}
func (r *SandboxFilesystemMultipartInitiateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.InitiateRequest)
}

type SandboxFilesystemMultipartUploadPartParams struct {
	// Part number (1-10000)
	PartNumber int64 `query:"partNumber" api:"required" json:"-"`
	// Part data
	File io.Reader `json:"file,omitzero" api:"required" format:"binary"`
	paramObj
}

func (r SandboxFilesystemMultipartUploadPartParams) MarshalMultipart() (data []byte, contentType string, err error) {
	buf := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(buf)
	err = apiform.MarshalRoot(r, writer)
	if err == nil {
		err = apiform.WriteExtras(writer, r.ExtraFields())
	}
	if err != nil {
		writer.Close()
		return nil, "", err
	}
	err = writer.Close()
	if err != nil {
		return nil, "", err
	}
	return buf.Bytes(), writer.FormDataContentType(), nil
}

// URLQuery serializes [SandboxFilesystemMultipartUploadPartParams]'s query
// parameters as `url.Values`.
func (r SandboxFilesystemMultipartUploadPartParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
