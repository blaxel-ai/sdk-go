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

// Returns all pre-built MCP server definitions available in the Blaxel Hub. These
// can be deployed directly to your workspace with pre-configured tools and
// integrations.
func (r *McpService) GetHub(ctx context.Context, opts ...option.RequestOption) (res *[]McpGetHubResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "mcp/hub"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type McpGetHubResponse struct {
	// Categories of the artifact
	Categories []map[string]any `json:"categories"`
	// If the artifact is coming soon
	ComingSoon bool `json:"coming_soon"`
	// The date and time when the resource was created
	CreatedAt string `json:"createdAt"`
	// Description of the artifact
	Description string `json:"description"`
	// Display name of the artifact
	DisplayName string `json:"displayName"`
	// If the artifact is enterprise
	Enterprise bool `json:"enterprise"`
	// Entrypoint of the artifact
	Entrypoint McpGetHubResponseEntrypoint `json:"entrypoint"`
	// Form of the artifact
	Form McpGetHubResponseForm `json:"form"`
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
	// The date and time when the resource was updated
	UpdatedAt string `json:"updatedAt"`
	// URL of the artifact
	URL string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Categories      respjson.Field
		ComingSoon      respjson.Field
		CreatedAt       respjson.Field
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
		UpdatedAt       respjson.Field
		URL             respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r McpGetHubResponse) RawJSON() string { return r.JSON.raw }
func (r *McpGetHubResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Entrypoint of the artifact
type McpGetHubResponseEntrypoint struct {
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
func (r McpGetHubResponseEntrypoint) RawJSON() string { return r.JSON.raw }
func (r *McpGetHubResponseEntrypoint) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Form of the artifact
type McpGetHubResponseForm struct {
	// Config of the artifact
	Config map[string]any `json:"config"`
	// OAuth of the artifact
	OAuth McpGetHubResponseFormOAuth `json:"oauth"`
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
func (r McpGetHubResponseForm) RawJSON() string { return r.JSON.raw }
func (r *McpGetHubResponseForm) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// OAuth of the artifact
type McpGetHubResponseFormOAuth struct {
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
func (r McpGetHubResponseFormOAuth) RawJSON() string { return r.JSON.raw }
func (r *McpGetHubResponseFormOAuth) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
