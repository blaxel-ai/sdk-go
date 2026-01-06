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

	"github.com/stainless-sdks/blaxel-go/internal/apijson"
	"github.com/stainless-sdks/blaxel-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/blaxel-go/internal/encoding/json"
	"github.com/stainless-sdks/blaxel-go/internal/requestconfig"
	"github.com/stainless-sdks/blaxel-go/option"
	"github.com/stainless-sdks/blaxel-go/packages/param"
	"github.com/stainless-sdks/blaxel-go/packages/respjson"
)

// VolumeTemplateService contains methods and other services that help with
// interacting with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVolumeTemplateService] method instead.
type VolumeTemplateService struct {
	Options []option.RequestOption
}

// NewVolumeTemplateService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewVolumeTemplateService(opts ...option.RequestOption) (r VolumeTemplateService) {
	r = VolumeTemplateService{}
	r.Options = opts
	return
}

// Creates a new volume template for initializing volumes with pre-configured
// filesystem contents. Optionally returns a presigned URL for uploading the
// template archive.
func (r *VolumeTemplateService) New(ctx context.Context, params VolumeTemplateNewParams, opts ...option.RequestOption) (res *VolumeTemplate, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "volume_templates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

// Returns a volume template by name.
func (r *VolumeTemplateService) Get(ctx context.Context, volumeTemplateName string, opts ...option.RequestOption) (res *VolumeTemplate, err error) {
	opts = slices.Concat(r.Options, opts)
	if volumeTemplateName == "" {
		err = errors.New("missing required volumeTemplateName parameter")
		return
	}
	path := fmt.Sprintf("volume_templates/%s", volumeTemplateName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns all volume templates in the workspace. Volume templates contain
// pre-configured filesystem snapshots that can be used to initialize new volumes.
func (r *VolumeTemplateService) List(ctx context.Context, opts ...option.RequestOption) (res *[]VolumeTemplate, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "volume_templates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a volume template by name.
func (r *VolumeTemplateService) Delete(ctx context.Context, volumeTemplateName string, opts ...option.RequestOption) (res *VolumeTemplate, err error) {
	opts = slices.Concat(r.Options, opts)
	if volumeTemplateName == "" {
		err = errors.New("missing required volumeTemplateName parameter")
		return
	}
	path := fmt.Sprintf("volume_templates/%s", volumeTemplateName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Deletes a specific version of a volume template.
func (r *VolumeTemplateService) DeleteVersion(ctx context.Context, versionName string, body VolumeTemplateDeleteVersionParams, opts ...option.RequestOption) (res *VolumeTemplateDeleteVersionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.VolumeTemplateName == "" {
		err = errors.New("missing required volumeTemplateName parameter")
		return
	}
	if versionName == "" {
		err = errors.New("missing required versionName parameter")
		return
	}
	path := fmt.Sprintf("volume_templates/%s/versions/%s", body.VolumeTemplateName, versionName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Creates or updates a volume template.
func (r *VolumeTemplateService) Upsert(ctx context.Context, volumeTemplateName string, params VolumeTemplateUpsertParams, opts ...option.RequestOption) (res *VolumeTemplate, err error) {
	opts = slices.Concat(r.Options, opts)
	if volumeTemplateName == "" {
		err = errors.New("missing required volumeTemplateName parameter")
		return
	}
	path := fmt.Sprintf("volume_templates/%s", volumeTemplateName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return
}

// Volume template for creating pre-configured volumes
type VolumeTemplate struct {
	// Common metadata fields shared by all Blaxel resources including name, labels,
	// timestamps, and ownership information
	Metadata Metadata `json:"metadata,required"`
	// Volume template specification
	Spec VolumeTemplateSpec `json:"spec,required"`
	// Volume template state
	State VolumeTemplateState `json:"state"`
	// List of versions for this template
	Versions []VolumeTemplateVersion `json:"versions"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metadata    respjson.Field
		Spec        respjson.Field
		State       respjson.Field
		Versions    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VolumeTemplate) RawJSON() string { return r.JSON.raw }
func (r *VolumeTemplate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this VolumeTemplate to a VolumeTemplateParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// VolumeTemplateParam.Overrides()
func (r VolumeTemplate) ToParam() VolumeTemplateParam {
	return param.Override[VolumeTemplateParam](json.RawMessage(r.RawJSON()))
}

// Volume template for creating pre-configured volumes
//
// The properties Metadata, Spec are required.
type VolumeTemplateParam struct {
	// Common metadata fields shared by all Blaxel resources including name, labels,
	// timestamps, and ownership information
	Metadata MetadataParam `json:"metadata,omitzero,required"`
	// Volume template specification
	Spec VolumeTemplateSpecParam `json:"spec,omitzero,required"`
	// Volume template state
	State VolumeTemplateStateParam `json:"state,omitzero"`
	// List of versions for this template
	Versions []VolumeTemplateVersionParam `json:"versions,omitzero"`
	paramObj
}

func (r VolumeTemplateParam) MarshalJSON() (data []byte, err error) {
	type shadow VolumeTemplateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VolumeTemplateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Volume template specification
type VolumeTemplateSpec struct {
	// Default size of the volume in MB
	DefaultSize int64 `json:"defaultSize"`
	// Description of the volume template
	Description string `json:"description"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DefaultSize respjson.Field
		Description respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VolumeTemplateSpec) RawJSON() string { return r.JSON.raw }
func (r *VolumeTemplateSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this VolumeTemplateSpec to a VolumeTemplateSpecParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// VolumeTemplateSpecParam.Overrides()
func (r VolumeTemplateSpec) ToParam() VolumeTemplateSpecParam {
	return param.Override[VolumeTemplateSpecParam](json.RawMessage(r.RawJSON()))
}

// Volume template specification
type VolumeTemplateSpecParam struct {
	// Default size of the volume in MB
	DefaultSize param.Opt[int64] `json:"defaultSize,omitzero"`
	// Description of the volume template
	Description param.Opt[string] `json:"description,omitzero"`
	paramObj
}

func (r VolumeTemplateSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow VolumeTemplateSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VolumeTemplateSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Volume template state
type VolumeTemplateState struct {
	// Timestamp of last version upload
	LastVersionUploadedAt string `json:"lastVersionUploadedAt"`
	// Current/latest S3 version ID
	LatestVersion string `json:"latestVersion"`
	// Status of the volume template (created, ready, error)
	//
	// Any of "created", "ready", "error".
	Status VolumeTemplateStateStatus `json:"status"`
	// Total number of versions for this template
	VersionCount int64 `json:"versionCount"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		LastVersionUploadedAt respjson.Field
		LatestVersion         respjson.Field
		Status                respjson.Field
		VersionCount          respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VolumeTemplateState) RawJSON() string { return r.JSON.raw }
func (r *VolumeTemplateState) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this VolumeTemplateState to a VolumeTemplateStateParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// VolumeTemplateStateParam.Overrides()
func (r VolumeTemplateState) ToParam() VolumeTemplateStateParam {
	return param.Override[VolumeTemplateStateParam](json.RawMessage(r.RawJSON()))
}

// Status of the volume template (created, ready, error)
type VolumeTemplateStateStatus string

const (
	VolumeTemplateStateStatusCreated VolumeTemplateStateStatus = "created"
	VolumeTemplateStateStatusReady   VolumeTemplateStateStatus = "ready"
	VolumeTemplateStateStatusError   VolumeTemplateStateStatus = "error"
)

// Volume template state
type VolumeTemplateStateParam struct {
	// Timestamp of last version upload
	LastVersionUploadedAt param.Opt[string] `json:"lastVersionUploadedAt,omitzero"`
	// Current/latest S3 version ID
	LatestVersion param.Opt[string] `json:"latestVersion,omitzero"`
	// Total number of versions for this template
	VersionCount param.Opt[int64] `json:"versionCount,omitzero"`
	// Status of the volume template (created, ready, error)
	//
	// Any of "created", "ready", "error".
	Status VolumeTemplateStateStatus `json:"status,omitzero"`
	paramObj
}

func (r VolumeTemplateStateParam) MarshalJSON() (data []byte, err error) {
	type shadow VolumeTemplateStateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VolumeTemplateStateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Volume template version tracking individual versions of template content
type VolumeTemplateVersion struct {
	// S3 bucket name where this version is stored
	Bucket string `json:"bucket"`
	// Size of the template content in bytes
	ContentSize int64 `json:"contentSize"`
	// Name of the template version
	Name string `json:"name"`
	// AWS region where this version is stored
	Region string `json:"region"`
	// Status of the version (CREATED, READY, FAILED)
	//
	// Any of "CREATED", "READY", "FAILED".
	Status VolumeTemplateVersionStatus `json:"status"`
	// Template name this version belongs to
	TemplateName string `json:"templateName"`
	// S3 version ID for this template version
	VersionID string `json:"versionId"`
	// Workspace name
	Workspace string `json:"workspace"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Bucket       respjson.Field
		ContentSize  respjson.Field
		Name         respjson.Field
		Region       respjson.Field
		Status       respjson.Field
		TemplateName respjson.Field
		VersionID    respjson.Field
		Workspace    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VolumeTemplateVersion) RawJSON() string { return r.JSON.raw }
func (r *VolumeTemplateVersion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this VolumeTemplateVersion to a VolumeTemplateVersionParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// VolumeTemplateVersionParam.Overrides()
func (r VolumeTemplateVersion) ToParam() VolumeTemplateVersionParam {
	return param.Override[VolumeTemplateVersionParam](json.RawMessage(r.RawJSON()))
}

// Status of the version (CREATED, READY, FAILED)
type VolumeTemplateVersionStatus string

const (
	VolumeTemplateVersionStatusCreated VolumeTemplateVersionStatus = "CREATED"
	VolumeTemplateVersionStatusReady   VolumeTemplateVersionStatus = "READY"
	VolumeTemplateVersionStatusFailed  VolumeTemplateVersionStatus = "FAILED"
)

// Volume template version tracking individual versions of template content
type VolumeTemplateVersionParam struct {
	// S3 bucket name where this version is stored
	Bucket param.Opt[string] `json:"bucket,omitzero"`
	// Size of the template content in bytes
	ContentSize param.Opt[int64] `json:"contentSize,omitzero"`
	// Name of the template version
	Name param.Opt[string] `json:"name,omitzero"`
	// AWS region where this version is stored
	Region param.Opt[string] `json:"region,omitzero"`
	// Template name this version belongs to
	TemplateName param.Opt[string] `json:"templateName,omitzero"`
	// S3 version ID for this template version
	VersionID param.Opt[string] `json:"versionId,omitzero"`
	// Workspace name
	Workspace param.Opt[string] `json:"workspace,omitzero"`
	// Status of the version (CREATED, READY, FAILED)
	//
	// Any of "CREATED", "READY", "FAILED".
	Status VolumeTemplateVersionStatus `json:"status,omitzero"`
	paramObj
}

func (r VolumeTemplateVersionParam) MarshalJSON() (data []byte, err error) {
	type shadow VolumeTemplateVersionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VolumeTemplateVersionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VolumeTemplateDeleteVersionResponse struct {
	Message string `json:"message"`
	// Volume template for creating pre-configured volumes
	Template VolumeTemplate `json:"template"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Template    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VolumeTemplateDeleteVersionResponse) RawJSON() string { return r.JSON.raw }
func (r *VolumeTemplateDeleteVersionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VolumeTemplateNewParams struct {
	// Volume template for creating pre-configured volumes
	VolumeTemplate VolumeTemplateParam
	// If true, returns a presigned URL for uploading the template content
	Upload param.Opt[bool] `query:"upload,omitzero" json:"-"`
	// Version name for the template version being created
	Version param.Opt[string] `query:"version,omitzero" json:"-"`
	paramObj
}

func (r VolumeTemplateNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.VolumeTemplate)
}
func (r *VolumeTemplateNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.VolumeTemplate)
}

// URLQuery serializes [VolumeTemplateNewParams]'s query parameters as
// `url.Values`.
func (r VolumeTemplateNewParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type VolumeTemplateDeleteVersionParams struct {
	VolumeTemplateName string `path:"volumeTemplateName,required" json:"-"`
	paramObj
}

type VolumeTemplateUpsertParams struct {
	// Volume template for creating pre-configured volumes
	VolumeTemplate VolumeTemplateParam
	// If true, returns a presigned URL for uploading the template content
	Upload param.Opt[bool] `query:"upload,omitzero" json:"-"`
	// Version name for the template version being created
	Version param.Opt[string] `query:"version,omitzero" json:"-"`
	paramObj
}

func (r VolumeTemplateUpsertParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.VolumeTemplate)
}
func (r *VolumeTemplateUpsertParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.VolumeTemplate)
}

// URLQuery serializes [VolumeTemplateUpsertParams]'s query parameters as
// `url.Values`.
func (r VolumeTemplateUpsertParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
