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

func TestVolumeListInstancesReturnsPageAndNextPage(t *testing.T) {
	queries := []string{}
	client := blaxel.NewClient(
		option.WithAPIKey("My API Key"),
		option.WithHTTPClient(&http.Client{
			Transport: &closureTransport{
				fn: func(req *http.Request) (*http.Response, error) {
					queries = append(queries, req.URL.RawQuery)

					body := `{
						"data": [{
							"metadata": {"name": "first-volume", "displayName": "First Volume"},
							"spec": {"region": "us-pdx-1", "size": 1024},
							"state": {},
							"status": "READY"
						}],
						"meta": {"hasMore": true, "nextCursor": "next-page", "total": 2}
					}`
					if req.URL.Query().Get("cursor") == "next-page" {
						body = `{
							"data": [{
								"metadata": {"name": "second-volume", "displayName": "Second Volume"},
								"spec": {"region": "us-pdx-1", "size": 2048},
								"state": {"attachedTo": "sandbox:test"},
								"status": "READY"
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

	volumes, err := client.Volumes.ListInstances(context.Background(), blaxel.VolumeListParams{
		Limit: blaxel.Int(1),
		Q:     blaxel.String("volume"),
		Sort:  blaxel.VolumeListParamsSortNameAsc,
	})
	if err != nil {
		t.Fatalf("ListInstances returned error: %v", err)
	}
	if len(volumes.Data) != 1 {
		t.Fatalf("expected 1 volume, got %d", len(volumes.Data))
	}
	if volumes.Data[0].Name() != "first-volume" {
		t.Fatalf("unexpected first volume name: %q", volumes.Data[0].Name())
	}
	if !volumes.HasNextPage() {
		t.Fatal("expected first page to have a next page")
	}
	if volumes.NextCursor() != "next-page" {
		t.Fatalf("expected next cursor, got %q", volumes.NextCursor())
	}

	next, err := volumes.NextPage(context.Background())
	if err != nil {
		t.Fatalf("NextPage returned error: %v", err)
	}
	if next == nil {
		t.Fatal("expected next page")
	}
	if len(next.Data) != 1 {
		t.Fatalf("expected 1 volume on next page, got %d", len(next.Data))
	}
	if next.Data[0].Name() != "second-volume" {
		t.Fatalf("unexpected second volume name: %q", next.Data[0].Name())
	}
	if next.Data[0].Size() != 2048 {
		t.Fatalf("expected second volume size to be mapped, got %d", next.Data[0].Size())
	}
	if next.HasNextPage() {
		t.Fatal("expected second page to be terminal")
	}
	if len(queries) != 2 {
		t.Fatalf("expected 2 requests, got %d", len(queries))
	}
	if queries[0] != "limit=1&q=volume&sort=name%3Aasc" {
		t.Fatalf("expected first request with list params, got %q", queries[0])
	}
	if queries[1] != "cursor=next-page&limit=1&q=volume&sort=name%3Aasc" {
		t.Fatalf("expected second request cursor query with preserved params, got %q", queries[1])
	}
}
