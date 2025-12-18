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

// McpService contains methods and other services that help with interacting with
// the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewMcpService] method instead.
type McpService struct {
	Options []option.RequestOption
}

// NewMcpService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewMcpService(opts ...option.RequestOption) (r McpService) {
	r = McpService{}
	r.Options = opts
	return
}

// List MCP hub definitions
func (r *McpService) ListHubs(ctx context.Context, opts ...option.RequestOption) (res *[]McpListHubsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "mcp/hub"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Definition of an MCP from the MCP Hub
type McpListHubsResponse struct {
	// Categories of the artifact
	Categories []map[string]any `json:"categories"`
	// If the artifact is coming soon
	ComingSoon bool `json:"coming_soon"`
	// Description of the artifact
	Description string `json:"description"`
	// Display name of the artifact
	DisplayName string `json:"displayName"`
	// If the artifact is enterprise
	Enterprise bool `json:"enterprise"`
	// Entrypoint of the artifact
	Entrypoint McpListHubsResponseEntrypoint `json:"entrypoint"`
	// Form of the artifact
	Form McpListHubsResponseForm `json:"form"`
	// If the artifact is hidden
	Hidden bool `json:"hidden"`
	// Hidden secrets of the artifact
	HiddenSecrets []string `json:"hiddenSecrets"`
	// Icon of the artifact
	Icon string `json:"icon"`
	// Image of the artifact
	Image string `json:"image"`
	// Integration of the artifact
	Integration string `json:"integration"`
	// Long description of the artifact
	LongDescription string `json:"longDescription"`
	// Name of the artifact
	Name string `json:"name"`
	// Transport compatibility for the MCP, can be "websocket" or "http-stream"
	Transport string `json:"transport"`
	// URL of the artifact
	URL string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Categories      respjson.Field
		ComingSoon      respjson.Field
		Description     respjson.Field
		DisplayName     respjson.Field
		Enterprise      respjson.Field
		Entrypoint      respjson.Field
		Form            respjson.Field
		Hidden          respjson.Field
		HiddenSecrets   respjson.Field
		Icon            respjson.Field
		Image           respjson.Field
		Integration     respjson.Field
		LongDescription respjson.Field
		Name            respjson.Field
		Transport       respjson.Field
		URL             respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
	TimeFields
}

// Returns the unmodified JSON received from the API
func (r McpListHubsResponse) RawJSON() string { return r.JSON.raw }
func (r *McpListHubsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Entrypoint of the artifact
type McpListHubsResponseEntrypoint struct {
	// Args of the entrypoint
	Args []map[string]any `json:"args"`
	// Command of the entrypoint
	Command string `json:"command"`
	// Env of the entrypoint
	Env map[string]string `json:"env"`
	// Super Gateway args of the entrypoint
	SuperGatewayArgs []map[string]any `json:"superGatewayArgs"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Args             respjson.Field
		Command          respjson.Field
		Env              respjson.Field
		SuperGatewayArgs respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r McpListHubsResponseEntrypoint) RawJSON() string { return r.JSON.raw }
func (r *McpListHubsResponseEntrypoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Form of the artifact
type McpListHubsResponseForm struct {
	// Config of the artifact
	Config map[string]any `json:"config"`
	// OAuth of the artifact
	OAuth McpListHubsResponseFormOAuth `json:"oauth"`
	// Secrets of the artifact
	Secrets map[string]any `json:"secrets"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Config      respjson.Field
		OAuth       respjson.Field
		Secrets     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r McpListHubsResponseForm) RawJSON() string { return r.JSON.raw }
func (r *McpListHubsResponseForm) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OAuth of the artifact
type McpListHubsResponseFormOAuth struct {
	// Scope of the OAuth
	Scope []map[string]any `json:"scope"`
	// Type of the OAuth
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Scope       respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r McpListHubsResponseFormOAuth) RawJSON() string { return r.JSON.raw }
func (r *McpListHubsResponseFormOAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
