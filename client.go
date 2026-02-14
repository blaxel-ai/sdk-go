// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blaxel

import (
	"context"
	"encoding/base64"
	"net/http"
	"os"
	"slices"
	"strings"

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
}

// DefaultClientOptions read from the environment (BL_API_KEY, BL_CLIENT_CREDENTIALS,
// BL_WORKSPACE, BL_ENV). This should be used to initialize new clients.
func DefaultClientOptions() []option.RequestOption {
	// Set up save callback for automatic credential persistence on token refresh
	setupOAuth2RefreshSaveCallback()

	// Initialize environment from workspace config and apply env var overrides
	workspace := GetDefaultWorkspace()
	InitializeEnvironment(workspace)

	defaults := []option.RequestOption{option.WithBaseURL(GetBaseURL())}
	if o, ok := os.LookupEnv("BL_API_KEY"); ok {
		defaults = append(defaults, option.WithAPIKey(o))
	}
	if o, ok := os.LookupEnv("BL_CLIENT_CREDENTIALS"); ok {
		// BL_CLIENT_CREDENTIALS is base64(clientId:clientSecret)
		// Decode and split - handled by custom client options
		defaults = append(defaults, option.WithClientCredentials(o))
	}
	// Add workspace header if specified
	if workspace != "" {
		defaults = append(defaults, option.WithWorkspace(workspace))
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

// GetDefaultWorkspace returns the workspace from environment or config file
func GetDefaultWorkspace() string {
	// First check environment variable
	if workspace := os.Getenv("BL_WORKSPACE"); workspace != "" {
		return workspace
	}

	// Fall back to config file
	context, err := CurrentContext()
	if err == nil && context.Workspace != "" {
		return context.Workspace
	}

	return ""
}

// NewDefaultClient creates a client using environment variables and config file.
// It follows this precedence:
// 1. BL_API_KEY environment variable
// 2. BL_CLIENT_CREDENTIALS environment variable (base64 encoded clientId:clientSecret)
// 3. Credentials from ~/.blaxel/config.yaml for the workspace
//
// For workspace, it uses:
// 1. BL_WORKSPACE environment variable
// 2. Current context from ~/.blaxel/config.yaml
//
// When using access token authentication with a refresh token, this client
// automatically refreshes tokens when they are about to expire (at 20% remaining lifetime).
func NewDefaultClient(opts ...option.RequestOption) (Client, error) {
	// Determine workspace
	workspace := GetDefaultWorkspace()

	// Initialize environment from workspace config and apply env var overrides
	InitializeEnvironment(workspace)

	// Start with provided options
	clientOpts := make([]option.RequestOption, 0, len(opts)+4)

	// Add base URL from environment configuration
	clientOpts = append(clientOpts, option.WithBaseURL(GetBaseURL()))

	// Check for BL_API_KEY first
	if apiKey := os.Getenv("BL_API_KEY"); apiKey != "" {
		clientOpts = append(clientOpts, option.WithAPIKey(apiKey))
	} else if clientCreds := os.Getenv("BL_CLIENT_CREDENTIALS"); clientCreds != "" {
		// Decode base64 client credentials (format: base64(clientId:clientSecret))
		decoded, err := base64.StdEncoding.DecodeString(clientCreds)
		if err != nil {
			// Try without base64 decoding
			decoded = []byte(clientCreds)
		}

		parts := strings.SplitN(string(decoded), ":", 2)
		if len(parts) == 2 {
			clientOpts = append(clientOpts,
				option.WithClientID(parts[0]),
				option.WithClientSecret(parts[1]),
			)
		}
	} else {
		// Fall back to config file
		if workspace != "" {
			creds, err := LoadCredentials(workspace)
			if err == nil && creds.IsValid() {
				if creds.ExpiresIn > 0 {
					clientOpts = append(clientOpts, option.WithExpires(creds.ExpiresIn))
				}
				if creds.DeviceCode != "" {
					clientOpts = append(clientOpts, option.WithDeviceCode(creds.DeviceCode))
				}
				if creds.RefreshToken != "" {
					clientOpts = append(clientOpts, option.WithRefreshToken(creds.RefreshToken))
				}
				if creds.APIKey != "" {
					clientOpts = append(clientOpts, option.WithAPIKey(creds.APIKey))
				} else if creds.AccessToken != "" {
					clientOpts = append(clientOpts, option.WithAccessToken(creds.AccessToken))
				} else if creds.ClientCredentials != "" {
					// Decode base64 client credentials
					decoded, err := base64.StdEncoding.DecodeString(creds.ClientCredentials)
					if err != nil {
						decoded = []byte(creds.ClientCredentials)
					}

					parts := strings.SplitN(string(decoded), ":", 2)
					if len(parts) == 2 {
						clientOpts = append(clientOpts,
							option.WithClientID(parts[0]),
							option.WithClientSecret(parts[1]),
						)
					}
				}
			}
		}
	}

	// Add workspace header if specified
	if workspace != "" {
		clientOpts = append(clientOpts, option.WithWorkspace(workspace))
	}

	// Add user-provided options last (they override defaults)
	clientOpts = append(clientOpts, opts...)

	client := NewClient(clientOpts...)
	return client, nil
}

// setupOAuth2RefreshSaveCallback configures the OAuth2 refresh state to save
// credentials after a token refresh
func setupOAuth2RefreshSaveCallback() {
	requestconfig.OAuth2RefreshCache.SetSaveCallback(func(workspace string, accessToken string, refreshToken string, expiresIn int) error {
		// Load existing credentials to preserve other fields
		creds, _ := LoadCredentials(workspace)
		creds.AccessToken = accessToken
		if refreshToken != "" {
			creds.RefreshToken = refreshToken
		}
		if expiresIn > 0 {
			creds.ExpiresIn = expiresIn
		}
		return SaveCredentials(workspace, creds)
	})
}

// NewClientFromConfig creates a client using credentials from ~/.blaxel/config.yaml
// for the specified workspace. It supports automatic token refresh when using
// access tokens with refresh tokens.
func NewClientFromConfig(workspaceName string, opts ...option.RequestOption) (*Client, error) {
	// Initialize environment from workspace config and apply env var overrides
	InitializeEnvironment(workspaceName)

	creds, err := LoadCredentials(workspaceName)
	if err != nil {
		return nil, err
	}

	clientOpts := make([]option.RequestOption, 0, len(opts)+6)
	clientOpts = append(clientOpts, option.WithBaseURL(GetBaseURL()))

	if creds.APIKey != "" {
		clientOpts = append(clientOpts, option.WithAPIKey(creds.APIKey))
	} else if creds.AccessToken != "" {
		// Use access token with automatic refresh support
		clientOpts = append(clientOpts, option.WithAccessToken(creds.AccessToken))
		if creds.RefreshToken != "" {
			clientOpts = append(clientOpts, option.WithRefreshToken(creds.RefreshToken))
		}
		if creds.DeviceCode != "" {
			clientOpts = append(clientOpts, option.WithDeviceCode(creds.DeviceCode))
		}
		if creds.ExpiresIn > 0 {
			clientOpts = append(clientOpts, option.WithExpires(creds.ExpiresIn))
		}
	} else if creds.ClientCredentials != "" {
		// Decode base64 client credentials
		decoded, err := base64.StdEncoding.DecodeString(creds.ClientCredentials)
		if err != nil {
			decoded = []byte(creds.ClientCredentials)
		}

		parts := strings.SplitN(string(decoded), ":", 2)
		if len(parts) == 2 {
			clientOpts = append(clientOpts,
				option.WithClientID(parts[0]),
				option.WithClientSecret(parts[1]),
			)
		}
	}

	// Add workspace header
	clientOpts = append(clientOpts, option.WithWorkspace(workspaceName))

	// Add user-provided options
	clientOpts = append(clientOpts, opts...)

	client := NewClient(clientOpts...)
	return &client, nil
}

// NewClientFromCredentials creates a client using the provided Credentials struct.
// This encapsulates the pattern of checking APIKey/AccessToken/ClientCredentials.
func NewClientFromCredentials(credentials Credentials, opts ...option.RequestOption) Client {
	clientOpts := make([]option.RequestOption, 0, len(opts)+6)
	clientOpts = append(clientOpts, option.WithBaseURL(GetBaseURL()))

	if credentials.APIKey != "" {
		clientOpts = append(clientOpts, option.WithAPIKey(credentials.APIKey))
	} else if credentials.AccessToken != "" {
		clientOpts = append(clientOpts, option.WithAccessToken(credentials.AccessToken))
		if credentials.RefreshToken != "" {
			clientOpts = append(clientOpts, option.WithRefreshToken(credentials.RefreshToken))
		}
		if credentials.DeviceCode != "" {
			clientOpts = append(clientOpts, option.WithDeviceCode(credentials.DeviceCode))
		}
		if credentials.ExpiresIn > 0 {
			clientOpts = append(clientOpts, option.WithExpires(credentials.ExpiresIn))
		}
	} else if credentials.ClientCredentials != "" {
		// Decode base64 client credentials
		decoded, err := base64.StdEncoding.DecodeString(credentials.ClientCredentials)
		if err != nil {
			decoded = []byte(credentials.ClientCredentials)
		}

		parts := strings.SplitN(string(decoded), ":", 2)
		if len(parts) == 2 {
			clientOpts = append(clientOpts,
				option.WithClientID(parts[0]),
				option.WithClientSecret(parts[1]),
			)
		}
	}

	// Add user-provided options last (they override defaults)
	clientOpts = append(clientOpts, opts...)

	return NewClient(clientOpts...)
}
