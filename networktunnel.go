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

// NetworkTunnelService contains methods and other services that help with
// interacting with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewNetworkTunnelService] method instead.
type NetworkTunnelService struct {
	Options []option.RequestOption
}

// NewNetworkTunnelService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewNetworkTunnelService(opts ...option.RequestOption) (r NetworkTunnelService) {
	r = NetworkTunnelService{}
	r.Options = opts
	return
}

// Stop the network tunnel and restore the original network configuration. WARNING:
// After disconnecting, the sandbox will lose all outbound internet connectivity
// (no egress). Inbound connections to the sandbox will still work. Use PUT
// /network/tunnel/config to re-establish the tunnel.
func (r *NetworkTunnelService) Disconnect(ctx context.Context, opts ...option.RequestOption) (res *NetworkTunnelDisconnectResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "network/tunnel"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Apply a new tunnel configuration on the fly. The existing tunnel is torn down
// and a new one is established. This endpoint is write-only; there is no
// corresponding GET to read the config back.
func (r *NetworkTunnelService) UpdateConfig(ctx context.Context, body NetworkTunnelUpdateConfigParams, opts ...option.RequestOption) (res *NetworkTunnelUpdateConfigResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "network/tunnel/config"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

type NetworkTunnelDisconnectResponse struct {
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
func (r NetworkTunnelDisconnectResponse) RawJSON() string { return r.JSON.raw }
func (r *NetworkTunnelDisconnectResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkTunnelUpdateConfigResponse struct {
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
func (r NetworkTunnelUpdateConfigResponse) RawJSON() string { return r.JSON.raw }
func (r *NetworkTunnelUpdateConfigResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type NetworkTunnelUpdateConfigParams struct {
	// Base64-encoded tunnel config JSON
	Config string `json:"config,required"`
	paramObj
}

func (r NetworkTunnelUpdateConfigParams) MarshalJSON() (data []byte, err error) {
	type shadow NetworkTunnelUpdateConfigParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *NetworkTunnelUpdateConfigParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
