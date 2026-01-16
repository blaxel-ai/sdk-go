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

// WorkspaceService contains methods and other services that help with interacting
// with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewWorkspaceService] method instead.
type WorkspaceService struct {
	Options []option.RequestOption
}

// NewWorkspaceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewWorkspaceService(opts ...option.RequestOption) (r WorkspaceService) {
	r = WorkspaceService{}
	r.Options = opts
	return
}

// Creates a new workspace tenant. The authenticated user becomes the workspace
// admin. Requires a linked billing account.
func (r *WorkspaceService) New(ctx context.Context, body WorkspaceNewParams, opts ...option.RequestOption) (res *Workspace, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "workspaces"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns detailed information about a workspace including its display name,
// account ID, status, and runtime configuration.
func (r *WorkspaceService) Get(ctx context.Context, workspaceName string, opts ...option.RequestOption) (res *Workspace, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceName == "" {
		err = errors.New("missing required workspaceName parameter")
		return
	}
	path := fmt.Sprintf("workspaces/%s", workspaceName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a workspace's settings such as display name and labels. The workspace
// name cannot be changed after creation.
func (r *WorkspaceService) Update(ctx context.Context, workspaceName string, body WorkspaceUpdateParams, opts ...option.RequestOption) (res *Workspace, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceName == "" {
		err = errors.New("missing required workspaceName parameter")
		return
	}
	path := fmt.Sprintf("workspaces/%s", workspaceName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Returns all workspaces the authenticated user has access to. Each workspace is a
// separate tenant with its own resources, team members, and billing.
func (r *WorkspaceService) List(ctx context.Context, opts ...option.RequestOption) (res *[]Workspace, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "workspaces"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Permanently deletes a workspace and ALL its resources (agents, functions,
// sandboxes, volumes, etc.). This action cannot be undone. Only workspace admins
// can delete a workspace.
func (r *WorkspaceService) Delete(ctx context.Context, workspaceName string, opts ...option.RequestOption) (res *Workspace, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceName == "" {
		err = errors.New("missing required workspaceName parameter")
		return
	}
	path := fmt.Sprintf("workspaces/%s", workspaceName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Accepts an invitation to a workspace.
func (r *WorkspaceService) AcceptInvitation(ctx context.Context, workspaceName string, opts ...option.RequestOption) (res *WorkspaceAcceptInvitationResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceName == "" {
		err = errors.New("missing required workspaceName parameter")
		return
	}
	path := fmt.Sprintf("workspaces/%s/join", workspaceName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Check if a workspace is available.
func (r *WorkspaceService) CheckAvailability(ctx context.Context, body WorkspaceCheckAvailabilityParams, opts ...option.RequestOption) (res *bool, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "workspaces/availability"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Declines an invitation to a workspace.
func (r *WorkspaceService) DeclineInvitation(ctx context.Context, workspaceName string, opts ...option.RequestOption) (res *PendingInvitation, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceName == "" {
		err = errors.New("missing required workspaceName parameter")
		return
	}
	path := fmt.Sprintf("workspaces/%s/decline", workspaceName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Leaves a workspace.
func (r *WorkspaceService) Leave(ctx context.Context, workspaceName string, opts ...option.RequestOption) (res *Workspace, err error) {
	opts = slices.Concat(r.Options, opts)
	if workspaceName == "" {
		err = errors.New("missing required workspaceName parameter")
		return
	}
	path := fmt.Sprintf("workspaces/%s/leave", workspaceName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

type Workspace struct {
	// Autogenerated unique workspace id
	ID string `json:"id"`
	// Workspace account id
	AccountID string `json:"accountId"`
	// The date and time when the resource was created
	CreatedAt string `json:"createdAt"`
	// The user or service account who created the resource
	CreatedBy string `json:"createdBy"`
	// Workspace display name
	DisplayName string `json:"displayName"`
	// Key-value pairs for organizing and filtering resources. Labels can be used to
	// categorize resources by environment, project, team, or any custom taxonomy.
	Labels map[string]string `json:"labels"`
	// Workspace name
	Name string `json:"name"`
	// Workspace write region
	Region string `json:"region"`
	// Runtime configuration for the workspace infrastructure
	Runtime WorkspaceRuntime `json:"runtime"`
	// Workspace status (created, account_binded, account_configured,
	// workspace_configured, ready, error)
	//
	// Any of "created", "account_binded", "account_configured",
	// "workspace_configured", "ready", "error".
	Status WorkspaceStatus `json:"status"`
	// Reason for current status (only set for error status)
	StatusReason string `json:"statusReason"`
	// The date and time when the resource was updated
	UpdatedAt string `json:"updatedAt"`
	// The user or service account who updated the resource
	UpdatedBy string `json:"updatedBy"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		AccountID    respjson.Field
		CreatedAt    respjson.Field
		CreatedBy    respjson.Field
		DisplayName  respjson.Field
		Labels       respjson.Field
		Name         respjson.Field
		Region       respjson.Field
		Runtime      respjson.Field
		Status       respjson.Field
		StatusReason respjson.Field
		UpdatedAt    respjson.Field
		UpdatedBy    respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Workspace) RawJSON() string { return r.JSON.raw }
func (r *Workspace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Workspace to a WorkspaceParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WorkspaceParam.Overrides()
func (r Workspace) ToParam() WorkspaceParam {
	return param.Override[WorkspaceParam](json.RawMessage(r.RawJSON()))
}

// Workspace status (created, account_binded, account_configured,
// workspace_configured, ready, error)
type WorkspaceStatus string

const (
	WorkspaceStatusCreated             WorkspaceStatus = "created"
	WorkspaceStatusAccountBinded       WorkspaceStatus = "account_binded"
	WorkspaceStatusAccountConfigured   WorkspaceStatus = "account_configured"
	WorkspaceStatusWorkspaceConfigured WorkspaceStatus = "workspace_configured"
	WorkspaceStatusReady               WorkspaceStatus = "ready"
	WorkspaceStatusError               WorkspaceStatus = "error"
)

type WorkspaceParam struct {
	// Workspace account id
	AccountID param.Opt[string] `json:"accountId,omitzero"`
	// Workspace display name
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	// Workspace name
	Name param.Opt[string] `json:"name,omitzero"`
	// Workspace write region
	Region param.Opt[string] `json:"region,omitzero"`
	// Key-value pairs for organizing and filtering resources. Labels can be used to
	// categorize resources by environment, project, team, or any custom taxonomy.
	Labels map[string]string `json:"labels,omitzero"`
	// Runtime configuration for the workspace infrastructure
	Runtime WorkspaceRuntimeParam `json:"runtime,omitzero"`
	paramObj
}

func (r WorkspaceParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkspaceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkspaceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Runtime configuration for the workspace infrastructure
type WorkspaceRuntime struct {
	// Infrastructure generation version for the workspace (affects available features
	// and deployment behavior)
	Generation string `json:"generation"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Generation  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkspaceRuntime) RawJSON() string { return r.JSON.raw }
func (r *WorkspaceRuntime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this WorkspaceRuntime to a WorkspaceRuntimeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// WorkspaceRuntimeParam.Overrides()
func (r WorkspaceRuntime) ToParam() WorkspaceRuntimeParam {
	return param.Override[WorkspaceRuntimeParam](json.RawMessage(r.RawJSON()))
}

// Runtime configuration for the workspace infrastructure
type WorkspaceRuntimeParam struct {
	// Infrastructure generation version for the workspace (affects available features
	// and deployment behavior)
	Generation param.Opt[string] `json:"generation,omitzero"`
	paramObj
}

func (r WorkspaceRuntimeParam) MarshalJSON() (data []byte, err error) {
	type shadow WorkspaceRuntimeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkspaceRuntimeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Pending invitation accept
type WorkspaceAcceptInvitationResponse struct {
	// User email
	Email string `json:"email"`
	// Tenant container that groups all Blaxel resources (agents, functions, models,
	// etc.) with shared team access control and billing.
	Workspace Workspace `json:"workspace"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email       respjson.Field
		Workspace   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkspaceAcceptInvitationResponse) RawJSON() string { return r.JSON.raw }
func (r *WorkspaceAcceptInvitationResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type WorkspaceNewParams struct {
	// Tenant container that groups all Blaxel resources (agents, functions, models,
	// etc.) with shared team access control and billing.
	Workspace WorkspaceParam
	paramObj
}

func (r WorkspaceNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Workspace)
}
func (r *WorkspaceNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Workspace)
}

type WorkspaceUpdateParams struct {
	// Tenant container that groups all Blaxel resources (agents, functions, models,
	// etc.) with shared team access control and billing.
	Workspace WorkspaceParam
	paramObj
}

func (r WorkspaceUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Workspace)
}
func (r *WorkspaceUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Workspace)
}

type WorkspaceCheckAvailabilityParams struct {
	Name string `json:"name,required"`
	paramObj
}

func (r WorkspaceCheckAvailabilityParams) MarshalJSON() (data []byte, err error) {
	type shadow WorkspaceCheckAvailabilityParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *WorkspaceCheckAvailabilityParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
