package blaxel_test

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"sync"
	"testing"
	"time"

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

func TestSandboxInstanceFSWriteBinaryMultipartLimitsConcurrentPartUploads(t *testing.T) {
	transport := newSandboxMultipartUploadTransport()
	sandbox := multipartUploadSandbox(transport)
	data := bytes.Repeat([]byte("a"), 3*sandboxMultipartUploadChunkSize+1)

	errCh := make(chan error, 1)
	var resp *blaxel.SandboxFilesystemWriteResponse
	go func() {
		var err error
		resp, err = sandbox.FS.WriteBinary(context.Background(), "large.bin", data, "0644")
		errCh <- err
	}()

	select {
	case <-transport.reachedLimit:
	case <-time.After(2 * time.Second):
		close(transport.releaseParts)
		t.Fatal("expected first three multipart part uploads to start")
	}

	select {
	case <-transport.overLimit:
	case <-time.After(250 * time.Millisecond):
	}

	close(transport.releaseParts)

	select {
	case err := <-errCh:
		if err != nil {
			t.Fatalf("WriteBinary returned error: %v", err)
		}
	case <-time.After(2 * time.Second):
		t.Fatal("timed out waiting for WriteBinary")
	}

	transport.mu.Lock()
	defer transport.mu.Unlock()

	if transport.maxActiveParts > sandboxMultipartUploadLimit {
		t.Fatalf("expected at most %d concurrent part uploads, got %d", sandboxMultipartUploadLimit, transport.maxActiveParts)
	}
	if transport.initiateCount != 1 {
		t.Fatalf("expected 1 initiate request, got %d", transport.initiateCount)
	}
	if transport.completeCount != 1 {
		t.Fatalf("expected 1 complete request, got %d", transport.completeCount)
	}
	if transport.abortCount != 0 {
		t.Fatalf("expected no abort request, got %d", transport.abortCount)
	}
	if len(transport.partNumbers) != 4 {
		t.Fatalf("expected 4 part uploads, got %d", len(transport.partNumbers))
	}
	if len(transport.completeParts) != 4 {
		t.Fatalf("expected complete request with 4 parts, got %d", len(transport.completeParts))
	}
	if resp == nil || resp.Path != "large.bin" || resp.Message != "uploaded" {
		t.Fatalf("unexpected WriteBinary response: %#v", resp)
	}
}

func TestSandboxInstanceFSWriteBinaryMultipartAbortsWhenPartUploadFails(t *testing.T) {
	transport := newSandboxMultipartUploadTransport()
	transport.failPartNumber = "2"
	close(transport.releaseParts)

	sandbox := multipartUploadSandbox(transport)
	data := bytes.Repeat([]byte("a"), sandboxMultipartUploadChunkSize+1)

	_, err := sandbox.FS.WriteBinary(context.Background(), "large.bin", data, "0644")
	if err == nil {
		t.Fatal("expected WriteBinary to return a part upload error")
	}

	transport.mu.Lock()
	defer transport.mu.Unlock()

	if transport.initiateCount != 1 {
		t.Fatalf("expected 1 initiate request, got %d", transport.initiateCount)
	}
	if transport.abortCount != 1 {
		t.Fatalf("expected 1 abort request, got %d", transport.abortCount)
	}
	if transport.completeCount != 0 {
		t.Fatalf("expected no complete request, got %d", transport.completeCount)
	}
}

const (
	sandboxMultipartUploadChunkSize = 5 * 1024 * 1024
	sandboxMultipartUploadLimit     = 3
)

type sandboxMultipartUploadTransport struct {
	mu sync.Mutex

	releaseParts chan struct{}
	reachedLimit chan struct{}
	overLimit    chan struct{}
	reachedOnce  sync.Once
	overOnce     sync.Once

	activeParts    int
	maxActiveParts int
	initiateCount  int
	completeCount  int
	abortCount     int
	partNumbers    []string
	completeParts  []sandboxMultipartCompletePart
	failPartNumber string
}

type sandboxMultipartCompletePart struct {
	Etag       string `json:"etag"`
	PartNumber int64  `json:"partNumber"`
}

func newSandboxMultipartUploadTransport() *sandboxMultipartUploadTransport {
	return &sandboxMultipartUploadTransport{
		releaseParts: make(chan struct{}),
		reachedLimit: make(chan struct{}),
		overLimit:    make(chan struct{}),
	}
}

func multipartUploadSandbox(transport *sandboxMultipartUploadTransport) *blaxel.SandboxInstance {
	client := blaxel.NewClient()
	return client.Sandboxes.FromSession(
		blaxel.SessionWithToken{
			Name:  "sandbox-session",
			URL:   "https://sandbox.test",
			Token: "preview-token",
		},
		option.WithHTTPClient(&http.Client{Transport: transport}),
		option.WithMaxRetries(0),
	)
}

func (t *sandboxMultipartUploadTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	switch {
	case req.Method == http.MethodPost && req.URL.Path == "/filesystem-multipart/initiate/large.bin":
		t.mu.Lock()
		t.initiateCount++
		t.mu.Unlock()
		return sandboxMultipartJSONResponse(http.StatusOK, `{"uploadId":"upload-1","path":"large.bin"}`), nil

	case req.Method == http.MethodPut && req.URL.Path == "/filesystem-multipart/upload-1/part":
		return t.handlePartUpload(req)

	case req.Method == http.MethodPost && req.URL.Path == "/filesystem-multipart/upload-1/complete":
		return t.handleComplete(req)

	case req.Method == http.MethodDelete && req.URL.Path == "/filesystem-multipart/upload-1/abort":
		t.mu.Lock()
		t.abortCount++
		t.mu.Unlock()
		return sandboxMultipartJSONResponse(http.StatusOK, `{"message":"aborted"}`), nil

	default:
		return sandboxMultipartJSONResponse(
			http.StatusNotFound,
			fmt.Sprintf(`{"error":{"message":"unexpected request: %s %s","status":404}}`, req.Method, req.URL.String()),
		), nil
	}
}

func (t *sandboxMultipartUploadTransport) handlePartUpload(req *http.Request) (*http.Response, error) {
	partNumber := req.URL.Query().Get("partNumber")
	t.mu.Lock()
	t.partNumbers = append(t.partNumbers, partNumber)
	t.activeParts++
	if t.activeParts > t.maxActiveParts {
		t.maxActiveParts = t.activeParts
	}
	if t.activeParts == sandboxMultipartUploadLimit {
		t.reachedOnce.Do(func() { close(t.reachedLimit) })
	}
	if t.activeParts > sandboxMultipartUploadLimit {
		t.overOnce.Do(func() { close(t.overLimit) })
	}
	t.mu.Unlock()

	<-t.releaseParts
	_, _ = io.Copy(io.Discard, req.Body)
	_ = req.Body.Close()

	t.mu.Lock()
	t.activeParts--
	shouldFail := partNumber == t.failPartNumber
	t.mu.Unlock()

	if shouldFail {
		return sandboxMultipartJSONResponse(http.StatusBadRequest, `{"error":{"message":"part failed","status":400}}`), nil
	}
	return sandboxMultipartJSONResponse(http.StatusOK, fmt.Sprintf(`{"etag":"etag-%s","partNumber":%s}`, partNumber, partNumber)), nil
}

func (t *sandboxMultipartUploadTransport) handleComplete(req *http.Request) (*http.Response, error) {
	defer req.Body.Close()
	var body struct {
		Parts []sandboxMultipartCompletePart `json:"parts"`
	}
	if err := json.NewDecoder(req.Body).Decode(&body); err != nil {
		return nil, err
	}

	t.mu.Lock()
	t.completeCount++
	t.completeParts = body.Parts
	t.mu.Unlock()

	return sandboxMultipartJSONResponse(http.StatusOK, `{"message":"uploaded","path":"large.bin"}`), nil
}

func sandboxMultipartJSONResponse(status int, body string) *http.Response {
	return &http.Response{
		StatusCode: status,
		Header:     http.Header{"Content-Type": []string{"application/json"}},
		Body:       io.NopCloser(strings.NewReader(body)),
	}
}
