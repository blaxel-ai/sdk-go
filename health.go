// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blaxel

import (
	"context"
	"net/http"
	"slices"

	"github.com/blaxel-ai/sdk-go/internal/apijson"
	"github.com/blaxel-ai/sdk-go/internal/requestconfig"
	"github.com/blaxel-ai/sdk-go/option"
	"github.com/blaxel-ai/sdk-go/packages/respjson"
)

// HealthService contains methods and other services that help with interacting
// with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewHealthService] method instead.
type HealthService struct {
	Options []option.RequestOption
}

// NewHealthService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewHealthService(opts ...option.RequestOption) (r HealthService) {
	r = HealthService{}
	r.Options = opts
	return
}

// Returns health status and system information including upgrade count and binary
// details Also includes last upgrade attempt status with detailed error
// information if available
func (r *HealthService) Check(ctx context.Context, opts ...option.RequestOption) (res *HealthCheckResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "health"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type HealthCheckResponse struct {
	Arch          string                         `json:"arch" api:"required"`
	BuildTime     string                         `json:"buildTime" api:"required"`
	GitCommit     string                         `json:"gitCommit" api:"required"`
	GoVersion     string                         `json:"goVersion" api:"required"`
	LastUpgrade   HealthCheckResponseLastUpgrade `json:"lastUpgrade" api:"required"`
	Os            string                         `json:"os" api:"required"`
	StartedAt     string                         `json:"startedAt" api:"required"`
	Status        string                         `json:"status" api:"required"`
	UpgradeCount  int64                          `json:"upgradeCount" api:"required"`
	Uptime        string                         `json:"uptime" api:"required"`
	UptimeSeconds float64                        `json:"uptimeSeconds" api:"required"`
	Version       string                         `json:"version" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Arch          respjson.Field
		BuildTime     respjson.Field
		GitCommit     respjson.Field
		GoVersion     respjson.Field
		LastUpgrade   respjson.Field
		Os            respjson.Field
		StartedAt     respjson.Field
		Status        respjson.Field
		UpgradeCount  respjson.Field
		Uptime        respjson.Field
		UptimeSeconds respjson.Field
		Version       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HealthCheckResponse) RawJSON() string { return r.JSON.raw }
func (r *HealthCheckResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type HealthCheckResponseLastUpgrade struct {
	// Current state (idle, running, completed, failed)
	//
	// Any of "idle", "running", "completed", "failed".
	Status string `json:"status" api:"required"`
	// Current/last step (none, starting, download, validate, replace, completed,
	// skipped)
	Step string `json:"step" api:"required"`
	// Version being upgraded to
	Version string `json:"version" api:"required"`
	// Path to downloaded binary
	BinaryPath string `json:"binaryPath"`
	// Bytes downloaded
	BytesDownloaded int64 `json:"bytesDownloaded"`
	// URL used for download
	DownloadURL string `json:"downloadUrl"`
	// Error message if failed
	Error string `json:"error"`
	// When the upgrade was attempted
	LastAttempt string `json:"lastAttempt"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status          respjson.Field
		Step            respjson.Field
		Version         respjson.Field
		BinaryPath      respjson.Field
		BytesDownloaded respjson.Field
		DownloadURL     respjson.Field
		Error           respjson.Field
		LastAttempt     respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r HealthCheckResponseLastUpgrade) RawJSON() string { return r.JSON.raw }
func (r *HealthCheckResponseLastUpgrade) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
