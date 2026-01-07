// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blaxel

import (
	"github.com/stainless-sdks/blaxel-go/option"
)

// SandboxFilesystemService contains methods and other services that help with
// interacting with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSandboxFilesystemService] method instead.
type SandboxFilesystemService struct {
	Options []option.RequestOption
}

// NewSandboxFilesystemService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewSandboxFilesystemService(opts ...option.RequestOption) (r SandboxFilesystemService) {
	r = SandboxFilesystemService{}
	r.Options = opts
	return
}
