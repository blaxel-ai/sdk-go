package blaxel_test

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	blaxel "github.com/blaxel-ai/sdk-go"
	"github.com/blaxel-ai/sdk-go/option"
)

// Use a real HTTP server: generated mock tests do not check version negotiation.
func TestImagePaginationContract(t *testing.T) {
	requests := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		requests++
		if got := r.Header.Get("Blaxel-Version"); got != "2026-09-22" {
			t.Errorf("version = %q", got)
		}
		if r.URL.Path != "/images" {
			t.Errorf("path = %q", r.URL.Path)
		}
		q := r.URL.Query()
		if q.Get("limit") != "1" || q.Get("q") != "demo" || q.Get("sort") != "name:asc" {
			t.Errorf("query = %v", q)
		}
		w.Header().Set("Content-Type", "application/json")
		switch requests {
		case 1:
			if q.Get("cursor") != "" {
				t.Errorf("unexpected cursor: %v", q)
			}
			fmt.Fprint(w, `{"data":[{"metadata":{"name":"demo-a"},"spec":{"size":42,"tagCount":10000}}],"meta":{"hasMore":true,"nextCursor":"opaque+/cursor="}}`)
		case 2:
			if q.Get("cursor") != "opaque+/cursor=" {
				t.Errorf("cursor = %q", q.Get("cursor"))
			}
			fmt.Fprint(w, `{"data":[{"metadata":{"name":"demo-b"},"spec":{"size":12,"tagCount":1}}],"meta":{"hasMore":false}}`)
		default:
			t.Error("unexpected additional request")
			http.Error(w, "extra request", 400)
		}
	}))
	defer server.Close()
	client := blaxel.NewClient(option.WithBaseURL(server.URL), option.WithAPIKey("test"))
	iter := client.Images.ListAutoPaging(context.Background(), blaxel.ImageListParams{Limit: blaxel.Int(1), Q: blaxel.String("demo"), Sort: blaxel.String("name:asc")})
	var names []string
	for iter.Next() {
		item := iter.Current()
		names = append(names, item.Metadata.Name)
		if len(names) == 1 && (item.Spec.TagCount != 10000 || item.Spec.Size != 42) {
			t.Errorf("summary = %+v", item.Spec)
		}
	}
	if err := iter.Err(); err != nil {
		t.Fatal(err)
	}
	if fmt.Sprint(names) != "[demo-a demo-b]" || requests != 2 {
		t.Fatalf("names=%v requests=%d", names, requests)
	}
}

func TestImageSharedSummaryAndTagsContract(t *testing.T) {
	tagRequests := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if got := r.Header.Get("Blaxel-Version"); got != "2026-09-22" {
			t.Errorf("version=%q", got)
		}
		if r.URL.Query().Get("sourceWorkspace") != "owner" {
			t.Errorf("missing shared image owner: %s", r.URL)
		}
		w.Header().Set("Content-Type", "application/json")
		switch r.URL.Path {
		case "/images/sandbox/demo":
			fmt.Fprint(w, `{"metadata":{"name":"demo","sourceWorkspace":"owner","lastDeployedAt":"2026-09-22T12:00:00Z"},"spec":{"size":42,"tagCount":10000}}`)
		case "/images/sandbox/demo/tags":
			tagRequests++
			q := r.URL.Query()
			if q.Get("limit") != "1" || q.Get("q") != "v" || q.Get("sort") != "name:desc" {
				t.Errorf("query=%v", q)
			}
			if tagRequests == 1 {
				fmt.Fprint(w, `{"data":[{"name":"v2","size":30}],"meta":{"hasMore":true,"nextCursor":"next-tag"}}`)
			} else if tagRequests == 2 {
				if q.Get("cursor") != "next-tag" {
					t.Errorf("cursor=%q", q.Get("cursor"))
				}
				fmt.Fprint(w, `{"data":[{"name":"v1","size":12}],"meta":{"hasMore":false}}`)
			} else {
				t.Error("unexpected additional tag request")
				http.Error(w, "extra request", 400)
			}
		default:
			t.Errorf("path=%s", r.URL.Path)
			http.NotFound(w, r)
		}
	}))
	defer server.Close()
	client := blaxel.NewClient(option.WithBaseURL(server.URL), option.WithAPIKey("test"))
	summary, err := client.Images.Get(context.Background(), "demo", blaxel.ImageGetParams{ResourceType: "sandbox", SourceWorkspace: blaxel.String("owner")})
	if err != nil {
		t.Fatal(err)
	}
	if summary.Spec.TagCount != 10000 || summary.Metadata.LastDeployedAt != "2026-09-22T12:00:00Z" {
		t.Fatalf("summary=%+v", summary)
	}
	iter := client.Images.Tags.ListAutoPaging(context.Background(), "demo", blaxel.ImageTagListParams{ResourceType: "sandbox", SourceWorkspace: blaxel.String("owner"), Limit: blaxel.Int(1), Q: blaxel.String("v"), Sort: blaxel.ImageTagListParamsSortNameDesc})
	var names []string
	for iter.Next() {
		names = append(names, iter.Current().Name)
	}
	if err := iter.Err(); err != nil {
		t.Fatal(err)
	}
	if fmt.Sprint(names) != "[v2 v1]" || tagRequests != 2 {
		t.Fatalf("tags=%v requests=%d", names, tagRequests)
	}
}

func TestSandboxPaginationWithImageAPIVersion(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("Blaxel-Version") != "2026-09-22" || r.URL.Path != "/sandboxes" {
			t.Errorf("request=%s version=%q", r.URL, r.Header.Get("Blaxel-Version"))
		}
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprint(w, `{"data":[{"metadata":{"name":"sandbox"},"spec":{}}],"meta":{"hasMore":false}}`)
	}))
	defer server.Close()
	client := blaxel.NewClient(option.WithBaseURL(server.URL), option.WithAPIKey("test"))
	page, err := client.Sandboxes.List(context.Background(), blaxel.SandboxListParams{})
	if err != nil {
		t.Fatal(err)
	}
	if len(page.Data) != 1 || page.Data[0].Metadata.Name != "sandbox" {
		t.Fatalf("page=%+v", page)
	}
}
