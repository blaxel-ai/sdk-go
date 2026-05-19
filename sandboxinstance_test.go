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

func TestSandboxListInstancesReturnsPageAndNextPage(t *testing.T) {
	queries := []string{}
	client := blaxel.NewClient(
		option.WithAPIKey("My API Key"),
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					queries = append(queries, req.URL.RawQuery)

					body := `{
						"data": [{
							"metadata": {"name": "first-sandbox"},
							"spec": {},
							"status": "DEPLOYED"
						}],
						"meta": {"hasMore": true, "nextCursor": "next-page", "total": 2}
					}`
					if req.URL.Query().Get("cursor") == "next-page" {
						body = `{
							"data": [{
								"metadata": {"name": "second-sandbox"},
								"spec": {},
								"status": "DEPLOYED"
							}],
							"meta": {"hasMore": false, "nextCursor": "", "total": 2}
						}`
					}

					return &http.Response{
						StatusCode: http.StatusOK,
						Header:     http.Header{"Content-Type": []string{"application/json"}},
						Body:       io.NopCloser(strings.NewReader(body)),
					}, nil
				},
			},
		}),
	)

	sandboxes, err := client.Sandboxes.ListInstances(context.Background(), blaxel.SandboxListParams{
		Limit: blaxel.Int(1),
		Q:     blaxel.String("sandbox"),
		Sort:  blaxel.SandboxListParamsSortNameAsc,
	})
	if err != nil {
		t.Fatalf("ListInstances returned error: %v", err)
	}
	if len(sandboxes.Data) != 1 {
		t.Fatalf("expected 1 sandbox, got %d", len(sandboxes.Data))
	}
	if sandboxes.Data[0].Metadata.Name != "first-sandbox" {
		t.Fatalf("unexpected first sandbox name: %q", sandboxes.Data[0].Metadata.Name)
	}
	if !sandboxes.HasNextPage() {
		t.Fatal("expected first page to have a next page")
	}
	if sandboxes.NextCursor() != "next-page" {
		t.Fatalf("expected next cursor, got %q", sandboxes.NextCursor())
	}

	next, err := sandboxes.NextPage(context.Background())
	if err != nil {
		t.Fatalf("NextPage returned error: %v", err)
	}
	if next == nil {
		t.Fatal("expected next page")
	}
	if len(next.Data) != 1 {
		t.Fatalf("expected 1 sandbox on next page, got %d", len(next.Data))
	}
	if next.Data[0].Metadata.Name != "second-sandbox" {
		t.Fatalf("unexpected second sandbox name: %q", next.Data[0].Metadata.Name)
	}
	if next.HasNextPage() {
		t.Fatal("expected second page to be terminal")
	}
	if len(queries) != 2 {
		t.Fatalf("expected 2 requests, got %d", len(queries))
	}
	if queries[0] != "limit=1&q=sandbox&sort=name%3Aasc" {
		t.Fatalf("expected first request with list params, got %q", queries[0])
	}
	if queries[1] != "cursor=next-page&limit=1&q=sandbox&sort=name%3Aasc" {
		t.Fatalf("expected second request cursor query with preserved params, got %q", queries[1])
	}
}
