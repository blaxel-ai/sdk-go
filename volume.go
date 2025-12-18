// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blaxel

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/blaxel-go/internal/apijson"
	shimjson "github.com/stainless-sdks/blaxel-go/internal/encoding/json"
	"github.com/stainless-sdks/blaxel-go/internal/requestconfig"
	"github.com/stainless-sdks/blaxel-go/option"
	"github.com/stainless-sdks/blaxel-go/packages/param"
	"github.com/stainless-sdks/blaxel-go/packages/respjson"
)

// VolumeService contains methods and other services that help with interacting
// with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewVolumeService] method instead.
type VolumeService struct {
	Options []option.RequestOption
}

// NewVolumeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewVolumeService(opts ...option.RequestOption) (r VolumeService) {
	r = VolumeService{}
	r.Options = opts
	return
}

// Creates a volume.
func (r *VolumeService) New(ctx context.Context, body VolumeNewParams, opts ...option.RequestOption) (res *Volume, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "volumes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns a volume by name.
func (r *VolumeService) Get(ctx context.Context, volumeName string, opts ...option.RequestOption) (res *Volume, err error) {
	opts = slices.Concat(r.Options, opts)
	if volumeName == "" {
		err = errors.New("missing required volumeName parameter")
		return
	}
	path := fmt.Sprintf("volumes/%s", volumeName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns a list of all volumes in the workspace.
func (r *VolumeService) List(ctx context.Context, opts ...option.RequestOption) (res *[]Volume, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "volumes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a volume by name.
func (r *VolumeService) Delete(ctx context.Context, volumeName string, opts ...option.RequestOption) (res *Volume, err error) {
	opts = slices.Concat(r.Options, opts)
	if volumeName == "" {
		err = errors.New("missing required volumeName parameter")
		return
	}
	path := fmt.Sprintf("volumes/%s", volumeName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Volume resource for persistent storage
type Volume struct {
	// Metadata
	Metadata Metadata `json:"metadata,required"`
	// Volume specification - immutable configuration
	Spec VolumeSpec `json:"spec,required"`
	// Events are loaded from the events table, not stored in volume table
	Events []CoreEvent `json:"events"`
	// Volume state - mutable runtime state
	State VolumeState `json:"state"`
	// Volume status computed from events
	Status string `json:"status"`
	// Timestamp when the volume was marked for termination
	TerminatedAt string `json:"terminatedAt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metadata     respjson.Field
		Spec         respjson.Field
		Events       respjson.Field
		State        respjson.Field
		Status       respjson.Field
		TerminatedAt respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Volume) RawJSON() string { return r.JSON.raw }
func (r *Volume) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Volume to a VolumeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// VolumeParam.Overrides()
func (r Volume) ToParam() VolumeParam {
	return param.Override[VolumeParam](json.RawMessage(r.RawJSON()))
}

// Volume specification - immutable configuration
type VolumeSpec struct {
	// The internal infrastructure resource identifier for this volume
	InfrastructureID string `json:"infrastructureId"`
	// Region where the volume should be created (e.g. us-pdx-1, eu-lon-1)
	Region string `json:"region"`
	// Size of the volume in MB
	Size int64 `json:"size"`
	// Volume template with revision (e.g. "mytemplate:1" or "mytemplate:latest")
	Template string `json:"template"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		InfrastructureID respjson.Field
		Region           respjson.Field
		Size             respjson.Field
		Template         respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VolumeSpec) RawJSON() string { return r.JSON.raw }
func (r *VolumeSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Volume state - mutable runtime state
type VolumeState struct {
	// Resource this volume is attached to (e.g. "sandbox:my-sandbox",
	// "model:my-model")
	AttachedTo string `json:"attachedTo"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AttachedTo  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r VolumeState) RawJSON() string { return r.JSON.raw }
func (r *VolumeState) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Volume resource for persistent storage
//
// The properties Metadata, Spec are required.
type VolumeParam struct {
	// Metadata
	Metadata MetadataParam `json:"metadata,omitzero,required"`
	// Volume specification - immutable configuration
	Spec VolumeSpecParam `json:"spec,omitzero,required"`
	// Volume status computed from events
	Status param.Opt[string] `json:"status,omitzero"`
	// Timestamp when the volume was marked for termination
	TerminatedAt param.Opt[string] `json:"terminatedAt,omitzero"`
	// Events are loaded from the events table, not stored in volume table
	Events []CoreEventParam `json:"events,omitzero"`
	// Volume state - mutable runtime state
	State VolumeStateParam `json:"state,omitzero"`
	paramObj
}

func (r VolumeParam) MarshalJSON() (data []byte, err error) {
	type shadow VolumeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VolumeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Volume specification - immutable configuration
type VolumeSpecParam struct {
	// The internal infrastructure resource identifier for this volume
	InfrastructureID param.Opt[string] `json:"infrastructureId,omitzero"`
	// Region where the volume should be created (e.g. us-pdx-1, eu-lon-1)
	Region param.Opt[string] `json:"region,omitzero"`
	// Size of the volume in MB
	Size param.Opt[int64] `json:"size,omitzero"`
	// Volume template with revision (e.g. "mytemplate:1" or "mytemplate:latest")
	Template param.Opt[string] `json:"template,omitzero"`
	paramObj
}

func (r VolumeSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow VolumeSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VolumeSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Volume state - mutable runtime state
type VolumeStateParam struct {
	// Resource this volume is attached to (e.g. "sandbox:my-sandbox",
	// "model:my-model")
	AttachedTo param.Opt[string] `json:"attachedTo,omitzero"`
	paramObj
}

func (r VolumeStateParam) MarshalJSON() (data []byte, err error) {
	type shadow VolumeStateParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VolumeStateParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type VolumeNewParams struct {
	// Volume resource for persistent storage
	Volume VolumeParam
	paramObj
}

func (r VolumeNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Volume)
}
func (r *VolumeNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Volume)
}
