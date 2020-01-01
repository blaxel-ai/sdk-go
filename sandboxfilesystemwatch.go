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

// Streams change events for a directory until the client disconnects.
//
// The body is JSON lines (sent with `Content-Type: text/plain`): one event object
// per line, e.g. `{"op":"WRITE","name":"main.go","path":"/app/src","error":null}`.
// `op` is the fsnotify operation (CREATE, WRITE, REMOVE, RENAME or CHMOD, several
// can be joined with `|`), `name` the base name of the changed entry, `path` the
// directory containing it, and `error` is always null. A `[keepalive]` line (not
// JSON) is sent every 30 seconds. The response starts once the watch is
// registered, and the stream ends when the watched directory is removed or
// renamed, or when events were dropped because the kernel queue overflowed (reopen
// and resync). Only the directory's direct entries are watched. To also watch
// every subdirectory, including ones created later, end the path with `/**`. The
// stream reports subsequent changes without an initial snapshot. When a new
// directory is detected, synthetic CREATE events are emitted for entries already
// present inside it to cover changes before its watch was registered.
func (r *SandboxFilesystemWatchService) Start(ctx context.Context, filePath string, query SandboxFilesystemWatchStartParams, opts ...option.RequestOption) (res *string, err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "text/plain")}, opts...)
	if filePath == "" {
		err = errors.New("missing required filePath parameter")
		return nil, err
	}
	path := fmt.Sprintf("watch/filesystem/%s", filePath)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type SandboxFilesystemWatchStartParams struct {
	// Comma-separated substrings; events whose full path contains one are skipped, and
	// matching directories are not watched at all
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
