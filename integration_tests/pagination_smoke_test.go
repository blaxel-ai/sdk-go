package integration_tests

import (
	"context"
	"testing"

	blaxel "github.com/blaxel-ai/sdk-go"
)

// TestDrivePaginationSmoke exercises native cursor pagination end-to-end against
// the real control plane: a small page size must span multiple pages, expose a
// cursor, fetch the next page without duplicates, and auto-paging must walk the
// whole set.
func TestDrivePaginationSmoke(t *testing.T) {
	client := newTestClient(t)
	ctx := context.Background()

	// Baseline: walk everything via the auto-pager.
	all := []string{}
	seen := map[string]bool{}
	autopager := client.Drives.ListAutoPaging(ctx, blaxel.DriveListParams{})
	for autopager.Next() {
		d := autopager.Current()
		all = append(all, d.Metadata.Name)
	}
	if err := autopager.Err(); err != nil {
		t.Fatalf("auto-paging failed: %v", err)
	}
	t.Logf("drives total via auto-paging = %d", len(all))
	if len(all) < 2 {
		t.Skipf("need >=2 drives to exercise multi-page pagination, found %d", len(all))
	}

	// First page with limit 1 must report more pages and expose a cursor.
	page1, err := client.Drives.List(ctx, blaxel.DriveListParams{Limit: blaxel.Int(1)})
	if err != nil {
		t.Fatalf("list page1: %v", err)
	}
	if len(page1.Data) != 1 {
		t.Fatalf("expected 1 item on page1, got %d", len(page1.Data))
	}
	if !page1.Meta.HasMore || page1.Meta.NextCursor == "" {
		t.Fatalf("expected HasMore=true and a NextCursor, got HasMore=%v cursor=%q", page1.Meta.HasMore, page1.Meta.NextCursor)
	}

	// Next page must be a different item (no repeated cursor / no duplicate).
	page2, err := page1.GetNextPage()
	if err != nil {
		t.Fatalf("get next page: %v", err)
	}
	if page2 == nil || len(page2.Data) < 1 {
		t.Fatalf("expected data on page2")
	}
	if page1.Data[0].Metadata.Name == page2.Data[0].Metadata.Name {
		t.Fatalf("page1 and page2 returned the same first item: %s", page1.Data[0].Metadata.Name)
	}

	// Auto-paging with a small page size walks every page without duplicates.
	walked := 0
	smallPager := client.Drives.ListAutoPaging(ctx, blaxel.DriveListParams{Limit: blaxel.Int(1)})
	for smallPager.Next() {
		name := smallPager.Current().Metadata.Name
		if seen[name] {
			t.Fatalf("duplicate item across pages: %s", name)
		}
		seen[name] = true
		walked++
	}
	if err := smallPager.Err(); err != nil {
		t.Fatalf("small-page auto-paging failed: %v", err)
	}
	// The auto-pager must advance past the first page (page size is 1, and we
	// have >=2 drives). We don't require walked == an earlier unbounded count:
	// the workspace is shared, so concurrent create/delete from other tests makes
	// an exact total flaky. No-duplicates (checked above) is the key invariant.
	if walked < 2 {
		t.Fatalf("auto-pager did not advance past the first page: walked=%d", walked)
	}
	t.Logf("PASS drive pagination: walked=%d unique=%d", walked, len(seen))
}

// TestSandboxListInstancesPaginationSmoke checks the custom ListInstances wrapper
// now built on top of native cursor pagination still paginates correctly.
func TestSandboxListInstancesPaginationSmoke(t *testing.T) {
	client := newTestClient(t)
	ctx := context.Background()

	first, err := client.Sandboxes.ListInstances(ctx, blaxel.SandboxListParams{Limit: blaxel.Int(1)})
	if err != nil {
		t.Fatalf("list instances: %v", err)
	}
	total := len(first.Data)
	if !first.HasNextPage() {
		t.Logf("only %d sandbox(es); single page", total)
		return
	}
	if first.Meta.NextCursor == "" {
		t.Fatalf("expected a NextCursor when HasNextPage is true")
	}
	next, err := first.NextPage(ctx)
	if err != nil {
		t.Fatalf("next page: %v", err)
	}
	if next == nil || len(next.Data) < 1 {
		t.Fatalf("expected sandboxes on the next page")
	}
	if first.Data[0].Metadata.Name == next.Data[0].Metadata.Name {
		t.Fatalf("ListInstances returned duplicate across pages")
	}
	t.Logf("PASS sandbox ListInstances pagination: page1=%d next=%d", total, len(next.Data))
}
