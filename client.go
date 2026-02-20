// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blaxel

import (
	"context"
	"net/http"
	"os"
	"slices"

	"github.com/blaxel-ai/sdk-go/internal/requestconfig"
	"github.com/blaxel-ai/sdk-go/option"
)

// Client creates a struct with services and top level methods that help with
// interacting with the blaxel API. You should not instantiate this client
// directly, and instead use the [NewClient] method instead.
type Client struct {
	Options         []option.RequestOption
	Agents          AgentService
	Functions       FunctionService
	Integrations    IntegrationService
	Images          ImageService
	Jobs            JobService
	Models          ModelService
	Policies        PolicyService
	PublicIPs       PublicIPService
	VolumeTemplates VolumeTemplateService
	Volumes         VolumeService
	Templates       TemplateService
	Workspaces      WorkspaceService
	Sandboxes       SandboxService
	Health          HealthService
	Upgrade         UpgradeService
	Features        FeatureService
	Vpcs            VpcService
	Network         NetworkService
	Egressgateways  EgressgatewayService
	Egressips       EgressipService
}

// DefaultClientOptions read from the environment (BL_API_KEY, BL_CLIENT_ID,
// BL_CLIENT_SECRET, BLAXEL_BASE_URL). This should be used to initialize new
// clients.
func DefaultClientOptions() []option.RequestOption {
	defaults := []option.RequestOption{option.WithEnvironmentProduction()}
	if o, ok := os.LookupEnv("BLAXEL_BASE_URL"); ok {
		defaults = append(defaults, option.WithBaseURL(o))
	}
	if o, ok := os.LookupEnv("BL_API_KEY"); ok {
		defaults = append(defaults, option.WithAPIKey(o))
	}
	if o, ok := os.LookupEnv("BL_CLIENT_ID"); ok {
		defaults = append(defaults, option.WithClientID(o))
	}
	if o, ok := os.LookupEnv("BL_CLIENT_SECRET"); ok {
		defaults = append(defaults, option.WithClientSecret(o))
	}
	return defaults
}

// NewClient generates a new client with the default option read from the
// environment (BL_API_KEY, BL_CLIENT_ID, BL_CLIENT_SECRET, BLAXEL_BASE_URL). The
// option passed in as arguments are applied after these default arguments, and all
// option will be passed down to the services and requests that this client makes.
func NewClient(opts ...option.RequestOption) (r Client) {
	opts = append(DefaultClientOptions(), opts...)

	r = Client{Options: opts}

	r.Agents = NewAgentService(opts...)
	r.Functions = NewFunctionService(opts...)
	r.Integrations = NewIntegrationService(opts...)
	r.Images = NewImageService(opts...)
	r.Jobs = NewJobService(opts...)
	r.Models = NewModelService(opts...)
	r.Policies = NewPolicyService(opts...)
	r.PublicIPs = NewPublicIPService(opts...)
	r.VolumeTemplates = NewVolumeTemplateService(opts...)
	r.Volumes = NewVolumeService(opts...)
	r.Templates = NewTemplateService(opts...)
	r.Workspaces = NewWorkspaceService(opts...)
	r.Sandboxes = NewSandboxService(opts...)
	r.Health = NewHealthService(opts...)
	r.Upgrade = NewUpgradeService(opts...)
	r.Features = NewFeatureService(opts...)
	r.Vpcs = NewVpcService(opts...)
	r.Network = NewNetworkService(opts...)
	r.Egressgateways = NewEgressgatewayService(opts...)
	r.Egressips = NewEgressipService(opts...)

	return
}

// Execute makes a request with the given context, method, URL, request params,
// response, and request options. This is useful for hitting undocumented endpoints
// while retaining the base URL, auth, retries, and other options from the client.
//
// If a byte slice or an [io.Reader] is supplied to params, it will be used as-is
// for the request body.
//
// The params is by default serialized into the body using [encoding/json]. If your
// type implements a MarshalJSON function, it will be used instead to serialize the
// request. If a URLQuery method is implemented, the returned [url.Values] will be
// used as query strings to the url.
//
// If your params struct uses [param.Field], you must provide either [MarshalJSON],
// [URLQuery], and/or [MarshalForm] functions. It is undefined behavior to use a
// struct uses [param.Field] without specifying how it is serialized.
//
// Any "…Params" object defined in this library can be used as the request
// argument. Note that 'path' arguments will not be forwarded into the url.
//
// The response body will be deserialized into the res variable, depending on its
// type:
//
//   - A pointer to a [*http.Response] is populated by the raw response.
//   - A pointer to a byte array will be populated with the contents of the request
//     body.
//   - A pointer to any other type uses this library's default JSON decoding, which
//     respects UnmarshalJSON if it is defined on the type.
//   - A nil value will not read the response body.
//
// For even greater flexibility, see [option.WithResponseInto] and
// [option.WithResponseBodyInto].
func (r *Client) Execute(ctx context.Context, method string, path string, params any, res any, opts ...option.RequestOption) error {
	opts = slices.Concat(r.Options, opts)
	return requestconfig.ExecuteNewRequest(ctx, method, path, params, res, opts...)
}

// Get makes a GET request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Get(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodGet, path, params, res, opts...)
}

// Post makes a POST request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Post(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPost, path, params, res, opts...)
}

// Put makes a PUT request with the given URL, params, and optionally deserializes
// to a response. See [Execute] documentation on the params and response.
func (r *Client) Put(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPut, path, params, res, opts...)
}

// Patch makes a PATCH request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Patch(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodPatch, path, params, res, opts...)
}

// Delete makes a DELETE request with the given URL, params, and optionally
// deserializes to a response. See [Execute] documentation on the params and
// response.
func (r *Client) Delete(ctx context.Context, path string, params any, res any, opts ...option.RequestOption) error {
	return r.Execute(ctx, http.MethodDelete, path, params, res, opts...)
}
