package mcp

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	officialMcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// echoParams is the input for the fixture "echo" tool below.
type echoParams struct {
	Message string `json:"message"`
}

func echoHandler(ctx context.Context, req *officialMcp.CallToolRequest, args echoParams) (*officialMcp.CallToolResult, any, error) {
	return &officialMcp.CallToolResult{
		Content: []officialMcp.Content{
			&officialMcp.TextContent{Text: "echo: " + args.Message},
		},
	}, nil, nil
}

// newFixtureServer starts a loopback HTTP server hosting a real
// github.com/modelcontextprotocol/go-sdk MCP server with a single "echo"
// tool. seenAuth, if non-nil, is set to the Authorization header of the
// last request the server received, so tests can confirm MCPClient forwards
// custom headers end to end.
func newFixtureServer(t *testing.T, seenAuth *string) *httptest.Server {
	t.Helper()

	server := officialMcp.NewServer(&officialMcp.Implementation{Name: "fixture-server", Version: "v0.0.1"}, nil)
	officialMcp.AddTool(server, &officialMcp.Tool{Name: "echo", Description: "echoes the given message"}, echoHandler)

	handler := officialMcp.NewStreamableHTTPHandler(func(r *http.Request) *officialMcp.Server {
		return server
	}, nil)

	mux := http.NewServeMux()
	mux.HandleFunc("/mcp", func(w http.ResponseWriter, r *http.Request) {
		if seenAuth != nil {
			*seenAuth = r.Header.Get("Authorization")
		}
		handler.ServeHTTP(w, r)
	})

	srv := httptest.NewServer(mux)
	t.Cleanup(srv.Close)
	return srv
}

// TestMCPClientHTTPStreamRoundTrip exercises the exact construction path
// production code uses: NewMCPClient, which auto-negotiates a transport and
// calls into officialMcp.NewClient(...).Connect(...), against a real
// in-process MCP server built from the same dependency. It is a genuine
// round trip (list + call a tool over the wire on a loopback listener), so
// it fails if github.com/modelcontextprotocol/go-sdk is removed, stubbed,
// or if a bump changes the client/transport wire format in a way that is
// incompatible with the server side of the same version -- exactly the risk
// a dependency bump introduces. It does not attempt to reproduce any of the
// four CVEs fixed between v1.2.0 and v1.4.1 (case-insensitive JSON field
// matching, DNS-rebinding protection, cross-origin tool execution, and
// null-Unicode JSON handling): those are server-hardening and parser
// behaviors best proven by the advisory's own fix plus the lockfile check
// (see VALIDATION in the report), not by a functional client test. This
// test's job is narrower and complementary: prove the bump did not break
// the client behaviour this repo ships.
func TestMCPClientHTTPStreamRoundTrip(t *testing.T) {
	var seenAuth string
	srv := newFixtureServer(t, &seenAuth)

	client, err := NewMCPClient(srv.URL, map[string]string{"Authorization": "Bearer test-token"})
	if err != nil {
		t.Fatalf("NewMCPClient: %v", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	tools, err := client.ListTools(ctx)
	if err != nil {
		t.Fatalf("ListTools: %v", err)
	}
	if len(tools.Tools) != 1 || tools.Tools[0].Name != "echo" {
		t.Fatalf("unexpected tools: %+v", tools.Tools)
	}

	result, err := client.CallTool(ctx, "echo", map[string]any{"message": "hello"})
	if err != nil {
		t.Fatalf("CallTool: %v", err)
	}
	if len(result.Content) != 1 {
		t.Fatalf("expected 1 content item, got %d", len(result.Content))
	}
	text, ok := result.Content[0].(*officialMcp.TextContent)
	if !ok {
		t.Fatalf("expected TextContent, got %T", result.Content[0])
	}
	if text.Text != "echo: hello" {
		t.Fatalf("unexpected echo result: %q", text.Text)
	}

	if seenAuth != "Bearer test-token" {
		t.Fatalf("expected Authorization header to be forwarded to the server, got %q", seenAuth)
	}
}
