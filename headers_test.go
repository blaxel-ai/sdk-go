package blaxel_test

import (
	"context"
	"net/http"
	"testing"

	blaxel "github.com/blaxel-ai/sdk-go"
	"github.com/blaxel-ai/sdk-go/option"
)

type headerTransport struct {
	fn func(req *http.Request) (*http.Response, error)
}

func (t *headerTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	return t.fn(req)
}

func TestBlaxelVersionHeader(t *testing.T) {
	var blaxelVersion string
	client := blaxel.NewClient(
		option.WithAPIKey("My API Key"),
		option.WithHTTPClient(&http.Client{
			Transport: &headerTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					blaxelVersion = req.Header.Get("Blaxel-Version")
					return &http.Response{
						StatusCode: http.StatusOK,
					}, nil
				},
			},
		}),
	)

	_, _ = client.Sandboxes.List(context.Background(), blaxel.SandboxListParams{})
	if blaxelVersion != "2026-04-28" {
		t.Errorf("Expected Blaxel-Version to be correct, but got: %#v", blaxelVersion)
	}
}
