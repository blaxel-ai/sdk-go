package blaxel

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"slices"
	"strings"
	"sync"
	"testing"
	"time"

	"github.com/blaxel-ai/sdk-go/option"
)

// The tests' sandbox name contains "409", so an error that carries the name
// cannot pass for a 409 conflict.

// writeSandbox answers with the sandbox sbx-409 in this status, with extra
// JSON fields appended to the body.
func writeSandbox(code int, status string, extra string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(code)
		fmt.Fprintf(w, `{"metadata":{"name":"sbx-409"},"spec":{},"status":%q%s}`, status, extra)
	}
}

// createWaitService is a SandboxService on a test server that answers the
// create with create and successive reads of sbx-409 with gets, repeating the
// last one. calls returns the requests the server saw, as "METHOD /path".
func createWaitService(t *testing.T, create http.HandlerFunc, gets ...http.HandlerFunc) (service *SandboxService, calls func() []string) {
	t.Helper()
	var mu sync.Mutex
	var seen []string
	reads := 0
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		mu.Lock()
		seen = append(seen, r.Method+" "+r.URL.Path)
		handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) { http.NotFound(w, r) })
		switch {
		case r.Method == http.MethodPost && r.URL.Path == "/sandboxes":
			handler = create
		case r.Method == http.MethodGet && r.URL.Path == "/sandboxes/sbx-409" && len(gets) > 0:
			handler = gets[min(reads, len(gets)-1)]
			reads++
		}
		mu.Unlock()
		handler(w, r)
	}))
	t.Cleanup(server.Close)
	client := NewClient(option.WithAPIKey("test"), option.WithBaseURL(server.URL), option.WithHTTPClient(server.Client()))
	return &client.Sandboxes, func() []string {
		mu.Lock()
		defer mu.Unlock()
		return slices.Clone(seen)
	}
}

var createWaitParams = SandboxNewParams{Sandbox: SandboxParam{
	Metadata: MetadataParam{Name: "sbx-409"},
	Spec:     SandboxSpecParam{Region: String("us-pdx-1")},
}}

func TestNewInstanceAfter202WaitsUntilDeployed(t *testing.T) {
	t.Parallel()
	service, calls := createWaitService(t,
		writeSandbox(http.StatusAccepted, "DEPLOYING", ""),
		writeSandbox(http.StatusOK, "DEPLOYING", ""),
		writeSandbox(http.StatusOK, "DEPLOYED", ""),
	)
	instance, err := service.NewInstance(context.Background(), createWaitParams)
	if err != nil {
		t.Fatalf("NewInstance: %v", err)
	}
	if instance.Status != StatusDeployed || instance.Metadata.Name != "sbx-409" {
		t.Fatalf("got %s %s, want sbx-409 DEPLOYED", instance.Metadata.Name, instance.Status)
	}
	want := []string{"POST /sandboxes", "GET /sandboxes/sbx-409", "GET /sandboxes/sbx-409"}
	if got := calls(); !slices.Equal(got, want) {
		t.Fatalf("requests %v, want %v", got, want)
	}
}

func TestNewInstanceAfter202ReturnsFailureReason(t *testing.T) {
	t.Parallel()
	service, calls := createWaitService(t,
		writeSandbox(http.StatusAccepted, "DEPLOYING", ""),
		writeSandbox(http.StatusOK, "FAILED", `,"events":[{"status":"DEPLOYING","message":"starting"},{"status":"FAILED","message":"image could not be pulled"}]`),
	)
	instance, err := service.NewInstance(context.Background(), createWaitParams)
	if instance != nil || err == nil {
		t.Fatalf("got instance %v and error %v, want an error", instance, err)
	}
	if want := "sandbox sbx-409 is FAILED instead of DEPLOYED: image could not be pulled"; err.Error() != want {
		t.Fatalf("error %q, want %q", err, want)
	}
	if got := len(calls()); got != 2 {
		t.Fatalf("%d requests, want 2: %v", got, calls())
	}
}

func TestNewInstanceAfter202UsesFatalErrorWithoutFailedEvent(t *testing.T) {
	t.Parallel()
	service, _ := createWaitService(t,
		writeSandbox(http.StatusAccepted, "DEPLOYING", ""),
		writeSandbox(http.StatusOK, "FAILED", `,"errors":[{"fatal":true,"message":"out of capacity"},{"fatal":false,"message":"restarted"}]`),
	)
	_, err := service.NewInstance(context.Background(), createWaitParams)
	if err == nil || !strings.HasSuffix(err.Error(), ": out of capacity") {
		t.Fatalf("error %v, want the fatal error's message", err)
	}
}

func TestNewInstanceAfter202RetriesTransientReads(t *testing.T) {
	t.Parallel()
	service, calls := createWaitService(t,
		writeSandbox(http.StatusAccepted, "DEPLOYING", ""),
		func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, "unavailable", http.StatusServiceUnavailable)
		},
		writeSandbox(http.StatusOK, "DEPLOYED", ""),
	)
	instance, err := service.NewInstance(context.Background(), createWaitParams)
	if err != nil || instance.Status != StatusDeployed {
		t.Fatalf("got %v, %v; want a DEPLOYED sandbox", instance, err)
	}
	// One read per poll: the wait does not stack the client's own retries on
	// top of its polling.
	if got := len(calls()); got != 3 {
		t.Fatalf("%d requests, want 3: %v", got, calls())
	}
}

func TestNewInstanceAfter202StopsOnPermanentReadError(t *testing.T) {
	t.Parallel()
	service, _ := createWaitService(t,
		writeSandbox(http.StatusAccepted, "DEPLOYING", ""),
		func(w http.ResponseWriter, r *http.Request) { http.Error(w, "forbidden", http.StatusForbidden) },
	)
	_, err := service.NewInstance(context.Background(), createWaitParams)
	var apiError *Error
	if !errors.As(err, &apiError) || apiError.StatusCode != http.StatusForbidden {
		t.Fatalf("error %v, want the 403", err)
	}
}

func TestNewInstanceWith200DoesNotWait(t *testing.T) {
	t.Parallel()
	for _, status := range []Status{StatusDeploying, StatusBuilding, StatusDeployed} {
		t.Run(string(status), func(t *testing.T) {
			t.Parallel()
			service, calls := createWaitService(t, writeSandbox(http.StatusOK, string(status), ""))
			instance, err := service.NewInstance(context.Background(), createWaitParams)
			if err != nil || instance.Status != status {
				t.Fatalf("got %v, %v; want a %s sandbox", instance, err, status)
			}
			if want := []string{"POST /sandboxes"}; !slices.Equal(calls(), want) {
				t.Fatalf("requests %v, want %v", calls(), want)
			}
		})
	}
}

func TestNewInstanceIfNotExistsAfter202ReturnsTheFailure(t *testing.T) {
	t.Parallel()
	service, calls := createWaitService(t,
		writeSandbox(http.StatusAccepted, "DEPLOYING", ""),
		writeSandbox(http.StatusOK, "FAILED", `,"events":[{"status":"FAILED","message":"image could not be pulled"}]`),
	)
	instance, err := service.NewInstanceIfNotExists(context.Background(), createWaitParams)
	if instance != nil || err == nil || !strings.HasSuffix(err.Error(), ": image could not be pulled") {
		t.Fatalf("got instance %v and error %v, want the failure", instance, err)
	}
	if got := len(calls()); got != 2 {
		t.Fatalf("%d requests, want 2: %v", got, calls())
	}
}

func TestNewInstanceIfNotExistsReturnsTheExistingSandboxOn409(t *testing.T) {
	t.Parallel()
	service, calls := createWaitService(t,
		func(w http.ResponseWriter, r *http.Request) {
			http.Error(w, `{"error":"sandbox already exists"}`, http.StatusConflict)
		},
		writeSandbox(http.StatusOK, "DEPLOYED", ""),
	)
	instance, err := service.NewInstanceIfNotExists(context.Background(), createWaitParams, option.WithMaxRetries(0))
	if err != nil || instance.Status != StatusDeployed {
		t.Fatalf("got %v, %v; want the existing sandbox", instance, err)
	}
	if want := []string{"POST /sandboxes", "GET /sandboxes/sbx-409"}; !slices.Equal(calls(), want) {
		t.Fatalf("requests %v, want %v", calls(), want)
	}
}

func TestWaitForCreationTimesOut(t *testing.T) {
	t.Parallel()
	service, _ := createWaitService(t, nil, writeSandbox(http.StatusOK, "BUILDING", ""))
	_, err := service.waitForCreation(context.Background(), "sbx-409", 700*time.Millisecond, nil)
	if err == nil || !strings.HasPrefix(err.Error(), "sandbox sbx-409 is still BUILDING 700ms after the API accepted its creation") || !errors.Is(err, context.DeadlineExceeded) {
		t.Fatalf("error %v, want sbx-409 still BUILDING", err)
	}
}

func TestWaitForCreationStopsWhenTheCallerCancels(t *testing.T) {
	t.Parallel()
	service, _ := createWaitService(t, nil, writeSandbox(http.StatusOK, "DEPLOYING", ""))
	ctx, cancel := context.WithCancel(context.Background())
	time.AfterFunc(100*time.Millisecond, cancel)
	_, err := service.waitForCreation(ctx, "sbx-409", sandboxCreateWaitMax, nil)
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("error %v, want context.Canceled", err)
	}
}
