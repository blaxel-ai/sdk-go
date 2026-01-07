// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blaxel

import (
	"github.com/stainless-sdks/blaxel-go/option"
)

// SandboxCodegenService contains methods and other services that help with
// interacting with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSandboxCodegenService] method instead.
type SandboxCodegenService struct {
	Options []option.RequestOption
}

// NewSandboxCodegenService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSandboxCodegenService(opts ...option.RequestOption) (r SandboxCodegenService) {
	r = SandboxCodegenService{}
	r.Options = opts
	return
}
