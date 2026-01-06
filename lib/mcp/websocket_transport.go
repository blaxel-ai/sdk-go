package mcp

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
	"github.com/modelcontextprotocol/go-sdk/jsonrpc"
	officialMcp "github.com/modelcontextprotocol/go-sdk/mcp"
)

// WebSocketTransport implements the MCP Transport interface for WebSocket connections
type WebSocketTransport struct {
	url     string
	headers map[string]string
}

// NewWebSocketTransport creates a new WebSocket transport
func NewWebSocketTransport(url string, headers map[string]string) *WebSocketTransport {
	return &WebSocketTransport{
		url:     url,
		headers: headers,
	}
}

// Connect implements the Transport interface
func (t *WebSocketTransport) Connect(ctx context.Context) (officialMcp.Connection, error) {
	// Build HTTP headers for WebSocket connection
	httpHeaders := http.Header{}
	for key, value := range t.headers {
		httpHeaders.Set(key, value)
	}

	// Dial WebSocket
	dialer := websocket.Dialer{}
	conn, _, err := dialer.DialContext(ctx, t.url, httpHeaders)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to WebSocket: %w", err)
	}

	return &WebSocketConnection{
		conn: conn,
	}, nil
}

// WebSocketConnection implements the MCP Connection interface
type WebSocketConnection struct {
	conn   *websocket.Conn
	mu     sync.Mutex
	closed bool
}

// Read reads the next message from the WebSocket connection
func (c *WebSocketConnection) Read(ctx context.Context) (jsonrpc.Message, error) {
	for {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
			_, message, err := c.conn.ReadMessage()
			if err != nil {
				if websocket.IsCloseError(err, websocket.CloseNormalClosure, websocket.CloseGoingAway) {
					return nil, io.EOF
				}
				return nil, err
			}

			// Parse the JSON-RPC message
			var raw json.RawMessage
			if err := json.Unmarshal(message, &raw); err != nil {
				continue // Skip invalid messages
			}

			msg, err := jsonrpc.DecodeMessage(raw)
			if err != nil {
				continue // Skip invalid messages
			}

			return msg, nil
		}
	}
}

// Write writes a message to the WebSocket connection
func (c *WebSocketConnection) Write(ctx context.Context, msg jsonrpc.Message) error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.closed {
		return fmt.Errorf("connection closed")
	}

	data, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	return c.conn.WriteMessage(websocket.TextMessage, data)
}

// Close closes the WebSocket connection
func (c *WebSocketConnection) Close() error {
	c.mu.Lock()
	defer c.mu.Unlock()

	if c.closed {
		return nil
	}
	c.closed = true

	return c.conn.Close()
}

// SessionID returns the session ID (not applicable for WebSocket)
func (c *WebSocketConnection) SessionID() string {
	return ""
}
