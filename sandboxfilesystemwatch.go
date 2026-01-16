// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blaxel

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/blaxel-ai/sdk-go/internal/apiquery"
	"github.com/blaxel-ai/sdk-go/internal/requestconfig"
	"github.com/blaxel-ai/sdk-go/option"
	"github.com/blaxel-ai/sdk-go/packages/param"
)

// SandboxFilesystemWatchService contains methods and other services that help with
// interacting with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSandboxFilesystemWatchService] method instead.
type SandboxFilesystemWatchService struct {
	Options []option.RequestOption
}

// NewSandboxFilesystemWatchService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSandboxFilesystemWatchService(opts ...option.RequestOption) (r SandboxFilesystemWatchService) {
	r = SandboxFilesystemWatchService{}
	r.Options = opts
	return
}

// Streams the path of modified files (one per line) in the given directory. Closes
// when the client disconnects.
func (r *SandboxFilesystemWatchService) Start(ctx context.Context, filePath string, query SandboxFilesystemWatchStartParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	if filePath == "" {
		err = errors.New("missing required filePath parameter")
		return
	}
	path := fmt.Sprintf("watch/filesystem/%s", filePath)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

type SandboxFilesystemWatchStartParams struct {
	// Ignore patterns (comma-separated)
	Ignore param.Opt[string] `query:"ignore,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SandboxFilesystemWatchStartParams]'s query parameters as
// `url.Values`.
func (r SandboxFilesystemWatchStartParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
