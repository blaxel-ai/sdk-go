// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blaxel

import (
	"context"
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

// SandboxSnapshotService contains methods and other services that help with
// interacting with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSandboxSnapshotService] method instead.
type SandboxSnapshotService struct {
	Options []option.RequestOption
}

// NewSandboxSnapshotService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSandboxSnapshotService(opts ...option.RequestOption) (r SandboxSnapshotService) {
	r = SandboxSnapshotService{}
	r.Options = opts
	return
}

// Creates a point-in-time snapshot of a sandbox. Snapshots capture the sandbox
// state and can be used for forking into new sandboxes or applications. This is a
// WIP endpoint — the full implementation depends on the execution plane.
func (r *SandboxSnapshotService) New(ctx context.Context, sandboxName string, body SandboxSnapshotNewParams, opts ...option.RequestOption) (res *SandboxSnapshot, err error) {
	opts = slices.Concat(r.Options, opts)
	if sandboxName == "" {
		err = errors.New("missing required sandboxName parameter")
		return nil, err
	}
	path := fmt.Sprintf("sandboxes/%s/snapshots", sandboxName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a list of snapshots for the specified sandbox.
func (r *SandboxSnapshotService) List(ctx context.Context, sandboxName string, opts ...option.RequestOption) (res *[]SandboxSnapshot, err error) {
	opts = slices.Concat(r.Options, opts)
	if sandboxName == "" {
		err = errors.New("missing required sandboxName parameter")
		return nil, err
	}
	path := fmt.Sprintf("sandboxes/%s/snapshots", sandboxName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Deletes a snapshot of a sandbox by its ID.
func (r *SandboxSnapshotService) Delete(ctx context.Context, snapshotID string, body SandboxSnapshotDeleteParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if body.SandboxName == "" {
		err = errors.New("missing required sandboxName parameter")
		return err
	}
	if snapshotID == "" {
		err = errors.New("missing required snapshotId parameter")
		return err
	}
	path := fmt.Sprintf("sandboxes/%s/snapshots/%s", body.SandboxName, snapshotID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return err
}

// Restores a sandbox to one of its own snapshots. The running sandbox is torn down
// and rebuilt from the snapshot under the same name and URLs, so everything it
// held since the snapshot was taken is lost unless it was itself snapshotted.
func (r *SandboxSnapshotService) Restore(ctx context.Context, snapshotID string, body SandboxSnapshotRestoreParams, opts ...option.RequestOption) (res *SandboxRestoreResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if body.SandboxName == "" {
		err = errors.New("missing required sandboxName parameter")
		return nil, err
	}
	if snapshotID == "" {
		err = errors.New("missing required snapshotId parameter")
		return nil, err
	}
	path := fmt.Sprintf("sandboxes/%s/snapshots/%s/restore", body.SandboxName, snapshotID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Result of restoring a sandbox to one of its snapshots. The sandbox keeps its
// name and URLs; everything it held since the snapshot was taken is gone.
type SandboxRestoreResponse struct {
	// Name of the restored sandbox
	Name string `json:"name" api:"required"`
	// Snapshot the sandbox was restored from
	SnapshotID string `json:"snapshotId" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		SnapshotID  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SandboxRestoreResponse) RawJSON() string { return r.JSON.raw }
func (r *SandboxRestoreResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A point-in-time snapshot of a sandbox that can be used for forking into a new
// sandbox or application.
type SandboxSnapshot struct {
	// Unique snapshot identifier
	ID string `json:"id" api:"required"`
	// When the snapshot was created
	CreatedAt string `json:"createdAt" api:"required"`
	// Name of the source sandbox
	SandboxName string `json:"sandboxName" api:"required"`
	// Status of the snapshot (pending, ready, failed)
	Status string `json:"status" api:"required"`
	// Workspace of the source sandbox
	Workspace string `json:"workspace" api:"required"`
	// Who created the snapshot
	CreatedBy string `json:"createdBy"`
	// Optional human-readable name for the snapshot
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		SandboxName respjson.Field
		Status      respjson.Field
		Workspace   respjson.Field
		CreatedBy   respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SandboxSnapshot) RawJSON() string { return r.JSON.raw }
func (r *SandboxSnapshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Request body for creating a snapshot of a sandbox. Captures the current sandbox
// state.
type SandboxSnapshotRequestParam struct {
	// Optional human-readable name for the snapshot
	Name param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r SandboxSnapshotRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow SandboxSnapshotRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SandboxSnapshotRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SandboxSnapshotNewParams struct {
	// Request body for creating a snapshot of a sandbox. Captures the current sandbox
	// state.
	SandboxSnapshotRequest SandboxSnapshotRequestParam
	paramObj
}

func (r SandboxSnapshotNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SandboxSnapshotRequest)
}
func (r *SandboxSnapshotNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SandboxSnapshotDeleteParams struct {
	SandboxName string `path:"sandboxName" api:"required" json:"-"`
	paramObj
}

type SandboxSnapshotRestoreParams struct {
	SandboxName string `path:"sandboxName" api:"required" json:"-"`
	paramObj
}
