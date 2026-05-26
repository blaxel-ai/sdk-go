package mcp

import (
	"context"
	"fmt"
	"net/http"
	"strings"
	"time"

	officialMcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// TransportType defines the type of transport to use
type TransportType string

const (
	TransportTypeAuto       TransportType = "auto"
	TransportTypeWebSocket  TransportType = "websocket"
	TransportTypeHTTPStream TransportType = "http-stream"
)

// MCPClient wraps the official MCP client with additional functionality
type MCPClient struct {
	session     *officialMcp.ClientSession
	transport   interface{}
	serverURL   string
	authHeaders map[string]string
}

// headerTransport adds headers to HTTP requests
type headerTransport struct {
	base    http.RoundTripper
	headers map[string]string
}

func (t *headerTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	// Clone the request to avoid modifying the original
	req = req.Clone(req.Context())

	// Add headers
	for key, value := range t.headers {
		req.Header.Set(key, value)
	}

	// Always include Accept header for SSE/streaming
	if req.Header.Get("Accept") == "" {
		req.Header.Set("Accept", "application/json, text/event-stream")
	}

	return t.base.RoundTrip(req)
}

// NewMCPClient creates a new MCP client with auto-detected transport
func NewMCPClient(serverURL string, authHeaders map[string]string) (*MCPClient, error) {
	return NewMCPClientWithTransport(serverURL, authHeaders, TransportTypeAuto)
}

// NewMCPClientWithTransport creates a new MCP client with specified transport type
func NewMCPClientWithTransport(serverURL string, authHeaders map[string]string, transportType TransportType) (*MCPClient, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Determine transport type
	if transportType == TransportTypeAuto {
		// Try HTTP Stream first (more reliable), fall back to WebSocket
		client, err := createHTTPStreamClient(ctx, serverURL, authHeaders)
		if err == nil {
			return client, nil
		}
		// Fall back to WebSocket
		return createWebSocketClient(ctx, serverURL, authHeaders)
	}

	switch transportType {
	case TransportTypeWebSocket:
		return createWebSocketClient(ctx, serverURL, authHeaders)
	case TransportTypeHTTPStream:
		return createHTTPStreamClient(ctx, serverURL, authHeaders)
	default:
		return nil, fmt.Errorf("unknown transport type: %s", transportType)
	}
}

func createWebSocketClient(ctx context.Context, serverURL string, authHeaders map[string]string) (*MCPClient, error) {
	// Convert HTTP URL to WebSocket URL
	wsURL := serverURL
	if strings.HasPrefix(wsURL, "http://") {
		wsURL = "ws://" + strings.TrimPrefix(wsURL, "http://")
	} else if strings.HasPrefix(wsURL, "https://") {
		wsURL = "wss://" + strings.TrimPrefix(wsURL, "https://")
	}

	// Ensure URL ends with /mcp for WebSocket
	if !strings.HasSuffix(wsURL, "/mcp") {
		wsURL = wsURL + "/mcp"
	}

	// Create WebSocket transport with headers
	transport := NewWebSocketTransport(wsURL, authHeaders)

	// Create client implementation
	impl := &officialMcp.Implementation{
		Name:    "blaxel-cli",
		Version: "1.0.0",
	}

	client := officialMcp.NewClient(impl, nil)

	// Connect
	session, err := client.Connect(ctx, transport, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to connect via WebSocket: %w", err)
	}

	return &MCPClient{
		session:     session,
		transport:   transport,
		serverURL:   serverURL,
		authHeaders: authHeaders,
	}, nil
}

func createHTTPStreamClient(ctx context.Context, serverURL string, authHeaders map[string]string) (*MCPClient, error) {
	// Create HTTP client with headers
	httpClient := &http.Client{
		Transport: &headerTransport{
			base:    http.DefaultTransport,
			headers: authHeaders,
		},
	}

	// Ensure URL ends with /mcp for HTTP stream
	if !strings.HasSuffix(serverURL, "/mcp") {
		serverURL = serverURL + "/mcp"
	}

	// Create transport
	transport := &officialMcp.StreamableClientTransport{
		Endpoint:   serverURL,
		HTTPClient: httpClient,
		MaxRetries: 3,
	}

	// Create client implementation
	impl := &officialMcp.Implementation{
		Name:    "blaxel-cli",
		Version: "1.0.0",
	}

	client := officialMcp.NewClient(impl, nil)

	// Connect
	session, err := client.Connect(ctx, transport, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to connect via HTTP Stream: %w", err)
	}

	return &MCPClient{
		session:     session,
		transport:   transport,
		serverURL:   serverURL,
		authHeaders: authHeaders,
	}, nil
}

// ListTools lists available tools from the MCP server
func (c *MCPClient) ListTools(ctx context.Context) (*officialMcp.ListToolsResult, error) {
	return c.session.ListTools(ctx, &officialMcp.ListToolsParams{})
}

// CallTool calls a tool on the MCP server
func (c *MCPClient) CallTool(ctx context.Context, name string, args map[string]interface{}) (*officialMcp.CallToolResult, error) {
	return c.session.CallTool(ctx, &officialMcp.CallToolParams{
		Name:      name,
		Arguments: args,
	})
}

// Close closes the MCP client connection
func (c *MCPClient) Close() error {
	if c.session != nil {
		return c.session.Close()
	}
	return nil
}
