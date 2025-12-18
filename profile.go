// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blaxel

import (
	"context"
	"net/http"
	"slices"

	"github.com/stainless-sdks/blaxel-go/internal/apijson"
	"github.com/stainless-sdks/blaxel-go/internal/requestconfig"
	"github.com/stainless-sdks/blaxel-go/option"
	"github.com/stainless-sdks/blaxel-go/packages/respjson"
)

// ProfileService contains methods and other services that help with interacting
// with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewProfileService] method instead.
type ProfileService struct {
	Options []option.RequestOption
}

// NewProfileService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewProfileService(opts ...option.RequestOption) (r ProfileService) {
	r = ProfileService{}
	r.Options = opts
	return
}

// Returns a list of all pending invitations in the workspace.
func (r *ProfileService) ListInvitations(ctx context.Context, opts ...option.RequestOption) (res *[]ProfileListInvitationsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "profile/invitations"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Pending invitation in workspace
type ProfileListInvitationsResponse struct {
	// User email
	Email string `json:"email"`
	// Invitation date
	InvitedAt string `json:"invitedAt"`
	// Invited by
	InvitedBy ProfileListInvitationsResponseInvitedBy `json:"invitedBy"`
	// ACL role
	Role string `json:"role"`
	// Workspace
	Workspace ProfileListInvitationsResponseWorkspace `json:"workspace"`
	// Workspace details
	WorkspaceDetails ProfileListInvitationsResponseWorkspaceDetails `json:"workspaceDetails"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email            respjson.Field
		InvitedAt        respjson.Field
		InvitedBy        respjson.Field
		Role             respjson.Field
		Workspace        respjson.Field
		WorkspaceDetails respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileListInvitationsResponse) RawJSON() string { return r.JSON.raw }
func (r *ProfileListInvitationsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Invited by
type ProfileListInvitationsResponseInvitedBy struct {
	// User email
	Email string `json:"email"`
	// User family name
	FamilyName string `json:"family_name"`
	// User given name
	GivenName string `json:"given_name"`
	// User sub
	Sub string `json:"sub"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email       respjson.Field
		FamilyName  respjson.Field
		GivenName   respjson.Field
		Sub         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileListInvitationsResponseInvitedBy) RawJSON() string { return r.JSON.raw }
func (r *ProfileListInvitationsResponseInvitedBy) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Workspace
type ProfileListInvitationsResponseWorkspace struct {
	// Workspace display name
	DisplayName string `json:"displayName"`
	// Workspace name
	Name string `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DisplayName respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileListInvitationsResponseWorkspace) RawJSON() string { return r.JSON.raw }
func (r *ProfileListInvitationsResponseWorkspace) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Workspace details
type ProfileListInvitationsResponseWorkspaceDetails struct {
	// List of user emails in the workspace
	Emails []map[string]any `json:"emails"`
	// Number of users in the workspace
	UserNumber float64 `json:"user_number"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Emails      respjson.Field
		UserNumber  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ProfileListInvitationsResponseWorkspaceDetails) RawJSON() string { return r.JSON.raw }
func (r *ProfileListInvitationsResponseWorkspaceDetails) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
