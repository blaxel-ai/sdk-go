// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blaxel

import (
	"context"
	"net/http"
	"slices"

	"github.com/blaxel-ai/sdk-go/internal/apijson"
	"github.com/blaxel-ai/sdk-go/internal/requestconfig"
	"github.com/blaxel-ai/sdk-go/option"
	"github.com/blaxel-ai/sdk-go/packages/param"
	"github.com/blaxel-ai/sdk-go/packages/respjson"
)

// UpgradeService contains methods and other services that help with interacting
// with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewUpgradeService] method instead.
type UpgradeService struct {
	Options []option.RequestOption
}

// NewUpgradeService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewUpgradeService(opts ...option.RequestOption) (r UpgradeService) {
	r = UpgradeService{}
	r.Options = opts
	return
}

// Triggers an upgrade of the sandbox-api process. Returns 200 immediately before
// upgrading. The upgrade will: download the specified binary from GitHub releases,
// validate it, and restart. All running processes will be preserved across the
// upgrade. Available versions: "develop" (default), "main", "latest", or specific
// tag like "v1.0.0" You can also specify a custom baseUrl for forks (defaults to
// https://github.com/blaxel-ai/sandbox/releases)
func (r *UpgradeService) Trigger(ctx context.Context, body UpgradeTriggerParams, opts ...option.RequestOption) (res *UpgradeTriggerResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "upgrade"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type UpgradeTriggerResponse struct {
	Message string `json:"message,required"`
	Path    string `json:"path"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message     respjson.Field
		Path        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r UpgradeTriggerResponse) RawJSON() string { return r.JSON.raw }
func (r *UpgradeTriggerResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type UpgradeTriggerParams struct {
	// Base URL for releases (useful for forks)
	BaseURL param.Opt[string] `json:"baseUrl,omitzero"`
	// Version to upgrade to: "develop", "main", "latest", or specific tag like
	// "v1.0.0"
	Version param.Opt[string] `json:"version,omitzero"`
	paramObj
}

func (r UpgradeTriggerParams) MarshalJSON() (data []byte, err error) {
	type shadow UpgradeTriggerParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *UpgradeTriggerParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
