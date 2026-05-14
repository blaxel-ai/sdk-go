// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blaxel

import (
	"github.com/blaxel-ai/sdk-go/option"
)

// PendingImageSharePendingService contains methods and other services that help
// with interacting with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewPendingImageSharePendingService] method instead.
type PendingImageSharePendingService struct {
	Options []option.RequestOption
}

// NewPendingImageSharePendingService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewPendingImageSharePendingService(opts ...option.RequestOption) (r PendingImageSharePendingService) {
	r = PendingImageSharePendingService{}
	r.Options = opts
	return
}
