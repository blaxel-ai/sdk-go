// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blaxel

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/blaxel-go/internal/apijson"
	"github.com/stainless-sdks/blaxel-go/internal/requestconfig"
	"github.com/stainless-sdks/blaxel-go/option"
	"github.com/stainless-sdks/blaxel-go/packages/param"
	"github.com/stainless-sdks/blaxel-go/packages/respjson"
)

// UserService contains methods and other services that help with interacting with
// the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUserService] method instead.
type UserService struct {
	Options []option.RequestOption
}

// NewUserService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUserService(opts ...option.RequestOption) (r UserService) {
	r = UserService{}
	r.Options = opts
	return
}

// Returns a list of all users in the workspace.
func (r *UserService) List(ctx context.Context, opts ...option.RequestOption) (res *[]WorkspaceUser, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "users"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Invites a user to the workspace by email.
func (r *UserService) Invite(ctx context.Context, body UserInviteParams, opts ...option.RequestOption) (res *PendingInvitation, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "users"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Removes a user from the workspace (or revokes an invitation if the user has not
// accepted the invitation yet).
func (r *UserService) Remove(ctx context.Context, subOrEmail string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if subOrEmail == "" {
		err = errors.New("missing required subOrEmail parameter")
		return
	}
	path := fmt.Sprintf("users/%s", subOrEmail)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Updates the role of a user in the workspace.
func (r *UserService) UpdateRole(ctx context.Context, subOrEmail string, body UserUpdateRoleParams, opts ...option.RequestOption) (res *WorkspaceUser, err error) {
	opts = slices.Concat(r.Options, opts)
	if subOrEmail == "" {
		err = errors.New("missing required subOrEmail parameter")
		return
	}
	path := fmt.Sprintf("users/%s", subOrEmail)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type PendingInvitation struct {
	// The date and time when the resource was created
	CreatedAt string `json:"createdAt"`
	// The user or service account who created the resource
	CreatedBy string `json:"createdBy"`
	// User email
	Email string `json:"email"`
	// User sub
	InvitedBy string `json:"invitedBy"`
	// ACL role
	Role string `json:"role"`
	// The date and time when the resource was updated
	UpdatedAt string `json:"updatedAt"`
	// The user or service account who updated the resource
	UpdatedBy string `json:"updatedBy"`
	// Workspace name
	Workspace string `json:"workspace"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CreatedAt   respjson.Field
		CreatedBy   respjson.Field
		Email       respjson.Field
		InvitedBy   respjson.Field
		Role        respjson.Field
		UpdatedAt   respjson.Field
		UpdatedBy   respjson.Field
		Workspace   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r PendingInvitation) RawJSON() string { return r.JSON.raw }
func (r *PendingInvitation) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Workspace user
type WorkspaceUser struct {
	// Whether the user has accepted the workspace invitation
	Accepted bool `json:"accepted"`
	// Workspace user email
	Email string `json:"email"`
	// Whether the user's email has been verified
	EmailVerified bool `json:"email_verified"`
	// Workspace user family name
	FamilyName string `json:"family_name"`
	// Workspace user given name
	GivenName string `json:"given_name"`
	// Workspace user role
	Role string `json:"role"`
	// Workspace user identifier
	Sub string `json:"sub"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Accepted      respjson.Field
		Email         respjson.Field
		EmailVerified respjson.Field
		FamilyName    respjson.Field
		GivenName     respjson.Field
		Role          respjson.Field
		Sub           respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r WorkspaceUser) RawJSON() string { return r.JSON.raw }
func (r *WorkspaceUser) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserInviteParams struct {
	Email param.Opt[string] `json:"email,omitzero" format:"email"`
	paramObj
}

func (r UserInviteParams) MarshalJSON() (data []byte, err error) {
	type shadow UserInviteParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserInviteParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UserUpdateRoleParams struct {
	// The new role to assign to the user
	Role string `json:"role,required"`
	paramObj
}

func (r UserUpdateRoleParams) MarshalJSON() (data []byte, err error) {
	type shadow UserUpdateRoleParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UserUpdateRoleParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
