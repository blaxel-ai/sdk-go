// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blaxel

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/blaxel-ai/sdk-go/internal/apijson"
	shimjson "github.com/blaxel-ai/sdk-go/internal/encoding/json"
	"github.com/blaxel-ai/sdk-go/internal/requestconfig"
	"github.com/blaxel-ai/sdk-go/option"
	"github.com/blaxel-ai/sdk-go/packages/param"
	"github.com/blaxel-ai/sdk-go/packages/respjson"
)

// DriveService contains methods and other services that help with interacting with
// the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDriveService] method instead.
type DriveService struct {
	Options []option.RequestOption
}

// NewDriveService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewDriveService(opts ...option.RequestOption) (r DriveService) {
	r = DriveService{}
	r.Options = opts
	return
}

// Creates a new drive in the workspace. Drives are backed by SeaweedFS buckets and
// can be mounted at runtime to sandboxes.
func (r *DriveService) New(ctx context.Context, body DriveNewParams, opts ...option.RequestOption) (res *Drive, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "drives"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieves details of a specific drive including its status and events.
func (r *DriveService) Get(ctx context.Context, driveName string, opts ...option.RequestOption) (res *Drive, err error) {
	opts = slices.Concat(r.Options, opts)
	if driveName == "" {
		err = errors.New("missing required driveName parameter")
		return
	}
	path := fmt.Sprintf("drives/%s", driveName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates an existing drive. Metadata fields like displayName and labels can be
// changed. Size can be set if not already configured.
func (r *DriveService) Update(ctx context.Context, driveName string, body DriveUpdateParams, opts ...option.RequestOption) (res *Drive, err error) {
	opts = slices.Concat(r.Options, opts)
	if driveName == "" {
		err = errors.New("missing required driveName parameter")
		return
	}
	path := fmt.Sprintf("drives/%s", driveName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Returns all drives in the workspace. Drives provide persistent storage that can
// be attached to agents, functions, and sandboxes.
func (r *DriveService) List(ctx context.Context, opts ...option.RequestOption) (res *[]Drive, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "drives"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a drive immediately. The drive record is removed from the database
// synchronously.
func (r *DriveService) Delete(ctx context.Context, driveName string, opts ...option.RequestOption) (res *DriveDeleteResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if driveName == "" {
		err = errors.New("missing required driveName parameter")
		return
	}
	path := fmt.Sprintf("drives/%s", driveName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Drive providing persistent storage that can be attached to agents, functions,
// and sandboxes. Drives are backed by SeaweedFS buckets and can be mounted at
// runtime via the sbx API.
type Drive struct {
	// Common metadata fields shared by all Blaxel resources including name, labels,
	// timestamps, and ownership information
	Metadata Metadata `json:"metadata" api:"required"`
	// Immutable drive configuration set at creation time
	Spec DriveSpec `json:"spec" api:"required"`
	// Events happening on a resource deployed on Blaxel
	Events []CoreEvent `json:"events"`
	// Current runtime state of the drive
	State DriveState `json:"state"`
	// Drive status computed from events
	Status string `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metadata    respjson.Field
		Spec        respjson.Field
		Events      respjson.Field
		State       respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Drive) RawJSON() string { return r.JSON.raw }
func (r *Drive) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Drive to a DriveParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// DriveParam.Overrides()
func (r Drive) ToParam() DriveParam {
	return param.Override[DriveParam](json.RawMessage(r.RawJSON()))
}

// Drive providing persistent storage that can be attached to agents, functions,
// and sandboxes. Drives are backed by SeaweedFS buckets and can be mounted at
// runtime via the sbx API.
//
// The properties Metadata, Spec are required.
type DriveParam struct {
	// Common metadata fields shared by all Blaxel resources including name, labels,
	// timestamps, and ownership information
	Metadata MetadataParam `json:"metadata,omitzero" api:"required"`
	// Immutable drive configuration set at creation time
	Spec DriveSpecParam `json:"spec,omitzero" api:"required"`
	paramObj
}

func (r DriveParam) MarshalJSON() (data []byte, err error) {
	type shadow DriveParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DriveParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Immutable drive configuration set at creation time
type DriveSpec struct {
	// The internal infrastructure resource identifier for this drive (bucket name)
	InfrastructureID string `json:"infrastructureId"`
	// Deployment region for the drive (e.g., us-pdx-1, eu-lon-1). Must match the
	// region of resources it attaches to.
	Region string `json:"region"`
	// Optional size limit for the drive in GB. If not specified, drive has no size
	// limit.
	Size int64 `json:"size"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InfrastructureID respjson.Field
		Region           respjson.Field
		Size             respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DriveSpec) RawJSON() string { return r.JSON.raw }
func (r *DriveSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this DriveSpec to a DriveSpecParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// DriveSpecParam.Overrides()
func (r DriveSpec) ToParam() DriveSpecParam {
	return param.Override[DriveSpecParam](json.RawMessage(r.RawJSON()))
}

// Immutable drive configuration set at creation time
type DriveSpecParam struct {
	// Deployment region for the drive (e.g., us-pdx-1, eu-lon-1). Must match the
	// region of resources it attaches to.
	Region param.Opt[string] `json:"region,omitzero"`
	// Optional size limit for the drive in GB. If not specified, drive has no size
	// limit.
	Size param.Opt[int64] `json:"size,omitzero"`
	paramObj
}

func (r DriveSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow DriveSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DriveSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DriveState map[string]any

type DriveDeleteResponse struct {
	Message string `json:"message"`
	Name    string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DriveDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *DriveDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DriveNewParams struct {
	// Drive providing persistent storage that can be attached to agents, functions,
	// and sandboxes. Drives are backed by SeaweedFS buckets and can be mounted at
	// runtime via the sbx API.
	Drive DriveParam
	paramObj
}

func (r DriveNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Drive)
}
func (r *DriveNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Drive)
}

type DriveUpdateParams struct {
	// Drive providing persistent storage that can be attached to agents, functions,
	// and sandboxes. Drives are backed by SeaweedFS buckets and can be mounted at
	// runtime via the sbx API.
	Drive DriveParam
	paramObj
}

func (r DriveUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Drive)
}
func (r *DriveUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Drive)
}
