package blaxel_test

import (
	"context"
	"io"
	"net/http"
	"strings"
	"testing"

	blaxel "github.com/blaxel-ai/sdk-go"
	"github.com/blaxel-ai/sdk-go/option"
)

func TestRunWithMetadataUsesJobMetadataURL(t *testing.T) {
	const (
		workspace       = "test-workspace"
		jobName         = "daily-import"
		metadataURL     = "https://job-runtime.example/daily-import"
		runtimeURL      = metadataURL + "/"
		controlPlaneURL = "https://api.example/v0/jobs/" + jobName
	)

	var requests []string
	client := blaxel.NewClient(
		option.WithAPIKey("My API Key"),
		option.WithBaseURL("https://api.example/v0"),
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					requests = append(requests, req.Method+" "+req.URL.String())

					switch req.URL.String() {
					case controlPlaneURL:
						if req.Method != http.MethodGet {
							t.Fatalf("metadata lookup used %s, want GET", req.Method)
						}
						return jsonResponse(http.StatusOK, `{"metadata":{"name":"`+jobName+`","url":"`+metadataURL+`"},"spec":{},"events":[],"status":"DEPLOYED"}`), nil
					case runtimeURL:
						if req.Method != http.MethodPost {
							t.Fatalf("runtime invocation used %s, want POST", req.Method)
						}
						return jsonResponse(http.StatusAccepted, `{"executionId":"exec-123"}`), nil
					default:
						t.Fatalf("unexpected request %s %s", req.Method, req.URL.String())
						return nil, nil
					}
				},
			},
		}),
	)

	resp, err := client.RunWithMetadata(context.Background(), workspace, "job", jobName, http.MethodPost, "", map[string]any{"tasks": []any{}})
	if err != nil {
		t.Fatalf("RunWithMetadata returned error: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusAccepted {
		t.Fatalf("status = %d, want %d", resp.StatusCode, http.StatusAccepted)
	}

	want := []string{
		http.MethodGet + " " + controlPlaneURL,
		http.MethodPost + " " + runtimeURL,
	}
	if strings.Join(requests, "\n") != strings.Join(want, "\n") {
		t.Fatalf("requests = %q, want %q", requests, want)
	}
}

func jsonResponse(status int, body string) *http.Response {
	return &http.Response{
		StatusCode: status,
		Header: http.Header{
			"Content-Type": []string{"application/json"},
		},
		Body: io.NopCloser(strings.NewReader(body)),
	}
}
