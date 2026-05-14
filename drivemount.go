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
	"github.com/blaxel-ai/sdk-go/packages/param"
	"github.com/blaxel-ai/sdk-go/packages/respjson"
)

// DriveMountService contains methods and other services that help with interacting
// with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewDriveMountService] method instead.
type DriveMountService struct {
	Options []option.RequestOption
}

// NewDriveMountService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewDriveMountService(opts ...option.RequestOption) (r DriveMountService) {
	r = DriveMountService{}
	r.Options = opts
	return
}

// Returns a list of all currently mounted drives managed by blfs
func (r *DriveMountService) List(ctx context.Context, opts ...option.RequestOption) (res *DriveMountListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "drives/mount"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Mounts an agent drive using the blfs binary to a local path, optionally mounting
// a subpath within the drive
func (r *DriveMountService) Attach(ctx context.Context, body DriveMountAttachParams, opts ...option.RequestOption) (res *DriveMountAttachResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "drives/mount"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Unmounts a previously mounted drive from the specified local path
func (r *DriveMountService) Detach(ctx context.Context, mountPath string, opts ...option.RequestOption) (res *DriveMountDetachResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if mountPath == "" {
		err = errors.New("missing required mountPath parameter")
		return nil, err
	}
	path := fmt.Sprintf("drives/mount/%s", mountPath)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type DriveMountListResponse struct {
	Mounts []DriveMountListResponseMount `json:"mounts"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Mounts      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DriveMountListResponse) RawJSON() string { return r.JSON.raw }
func (r *DriveMountListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DriveMountListResponseMount struct {
	DriveName string `json:"driveName"`
	DrivePath string `json:"drivePath"`
	MountPath string `json:"mountPath"`
	ReadOnly  bool   `json:"readOnly"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DriveName   respjson.Field
		DrivePath   respjson.Field
		MountPath   respjson.Field
		ReadOnly    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DriveMountListResponseMount) RawJSON() string { return r.JSON.raw }
func (r *DriveMountListResponseMount) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DriveMountAttachResponse struct {
	DriveName string `json:"driveName"`
	DrivePath string `json:"drivePath"`
	Message   string `json:"message"`
	MountPath string `json:"mountPath"`
	ReadOnly  bool   `json:"readOnly"`
	Success   bool   `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DriveName   respjson.Field
		DrivePath   respjson.Field
		Message     respjson.Field
		MountPath   respjson.Field
		ReadOnly    respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DriveMountAttachResponse) RawJSON() string { return r.JSON.raw }
func (r *DriveMountAttachResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DriveMountDetachResponse struct {
	Message   string `json:"message"`
	MountPath string `json:"mountPath"`
	Success   bool   `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		MountPath   respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r DriveMountDetachResponse) RawJSON() string { return r.JSON.raw }
func (r *DriveMountDetachResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type DriveMountAttachParams struct {
	DriveName string `json:"driveName" api:"required"`
	MountPath string `json:"mountPath" api:"required"`
	// Optional, defaults to "/"
	DrivePath param.Opt[string] `json:"drivePath,omitzero"`
	// Optional, defaults to false
	ReadOnly param.Opt[bool] `json:"readOnly,omitzero"`
	paramObj
}

func (r DriveMountAttachParams) MarshalJSON() (data []byte, err error) {
	type shadow DriveMountAttachParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *DriveMountAttachParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
