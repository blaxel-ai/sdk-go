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

// Creates a new persistent storage volume that can be attached to sandboxes.
// Volumes must be created in a specific region and can only attach to sandboxes in
// the same region.
func (r *VolumeService) New(ctx context.Context, body VolumeNewParams, opts ...option.RequestOption) (res *Volume, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "volumes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns detailed information about a volume including its size, region,
// attachment status, and any events history.
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

// Returns all persistent storage volumes in the workspace. Volumes can be attached
// to sandboxes for durable file storage that persists across sessions and sandbox
// deletions.
func (r *VolumeService) List(ctx context.Context, opts ...option.RequestOption) (res *[]Volume, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "volumes"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Permanently deletes a volume and all its data. The volume must not be attached
// to any sandbox. This action cannot be undone.
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

// Persistent storage volume that can be attached to sandboxes for durable file
// storage across sessions. Volumes survive sandbox deletion and can be reattached
// to new sandboxes.
type Volume struct {
	// Common metadata fields shared by all Blaxel resources including name, labels,
	// timestamps, and ownership information
	Metadata Metadata `json:"metadata,required"`
	// Immutable volume configuration set at creation time (size and region cannot be
	// changed after creation)
	Spec VolumeSpec `json:"spec,required"`
	// Events happening on a resource deployed on Blaxel
	Events []CoreEvent `json:"events"`
	// Current runtime state of the volume including attachment status
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

// Immutable volume configuration set at creation time (size and region cannot be
// changed after creation)
type VolumeSpec struct {
	// The internal infrastructure resource identifier for this volume
	InfrastructureID string `json:"infrastructureId"`
	// Deployment region for the volume (e.g., us-pdx-1, eu-lon-1). Must match the
	// region of sandboxes it attaches to.
	Region string `json:"region"`
	// Storage capacity in megabytes. Can be increased after creation but not
	// decreased.
	Size int64 `json:"size"`
	// Volume template to initialize from, with optional revision (e.g., "mytemplate:1"
	// or "mytemplate:latest")
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

// Current runtime state of the volume including attachment status
type VolumeState struct {
	// Resource currently using this volume in format "type:name" (e.g.,
	// "sandbox:my-sandbox"). Empty if not attached.
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

// Persistent storage volume that can be attached to sandboxes for durable file
// storage across sessions. Volumes survive sandbox deletion and can be reattached
// to new sandboxes.
//
// The properties Metadata, Spec are required.
type VolumeParam struct {
	// Common metadata fields shared by all Blaxel resources including name, labels,
	// timestamps, and ownership information
	Metadata MetadataParam `json:"metadata,omitzero,required"`
	// Immutable volume configuration set at creation time (size and region cannot be
	// changed after creation)
	Spec VolumeSpecParam `json:"spec,omitzero,required"`
	// Volume status computed from events
	Status param.Opt[string] `json:"status,omitzero"`
	// Timestamp when the volume was marked for termination
	TerminatedAt param.Opt[string] `json:"terminatedAt,omitzero"`
	paramObj
}

func (r VolumeParam) MarshalJSON() (data []byte, err error) {
	type shadow VolumeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *VolumeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Immutable volume configuration set at creation time (size and region cannot be
// changed after creation)
type VolumeSpecParam struct {
	// Deployment region for the volume (e.g., us-pdx-1, eu-lon-1). Must match the
	// region of sandboxes it attaches to.
	Region param.Opt[string] `json:"region,omitzero"`
	// Storage capacity in megabytes. Can be increased after creation but not
	// decreased.
	Size param.Opt[int64] `json:"size,omitzero"`
	// Volume template to initialize from, with optional revision (e.g., "mytemplate:1"
	// or "mytemplate:latest")
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

// Current runtime state of the volume including attachment status
type VolumeStateParam struct {
	// Resource currently using this volume in format "type:name" (e.g.,
	// "sandbox:my-sandbox"). Empty if not attached.
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
	// Persistent storage volume that can be attached to sandboxes for durable file
	// storage across sessions. Volumes survive sandbox deletion and can be reattached
	// to new sandboxes.
	Volume VolumeParam
	paramObj
}

func (r VolumeNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Volume)
}
func (r *VolumeNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Volume)
}
