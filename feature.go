// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blaxel

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"
	"time"

	"github.com/blaxel-ai/sdk-go/internal/apijson"
	"github.com/blaxel-ai/sdk-go/internal/requestconfig"
	"github.com/blaxel-ai/sdk-go/option"
	"github.com/blaxel-ai/sdk-go/packages/respjson"
)

// FeatureService contains methods and other services that help with interacting
// with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewFeatureService] method instead.
type FeatureService struct {
	Options []option.RequestOption
}

// NewFeatureService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewFeatureService(opts ...option.RequestOption) (r FeatureService) {
	r = FeatureService{}
	r.Options = opts
	return
}

// Evaluates a specific feature flag for the workspace with full details including
// variant and payload. Useful for testing and debugging feature flag targeting.
func (r *FeatureService) Get(ctx context.Context, featureKey string, opts ...option.RequestOption) (res *FeatureGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if featureKey == "" {
		err = errors.New("missing required featureKey parameter")
		return
	}
	path := fmt.Sprintf("features/%s", featureKey)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns only the feature flags that are currently enabled for the specified
// workspace. Disabled features are not included to prevent information leakage.
func (r *FeatureService) List(ctx context.Context, opts ...option.RequestOption) (res *FeatureListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "features"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type FeatureGetResponse struct {
	Enabled     bool      `json:"enabled"`
	EvaluatedAt time.Time `json:"evaluatedAt" format:"date-time"`
	FeatureKey  string    `json:"featureKey"`
	Payload     any       `json:"payload"`
	Variant     string    `json:"variant"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		EvaluatedAt respjson.Field
		FeatureKey  respjson.Field
		Payload     respjson.Field
		Variant     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FeatureGetResponse) RawJSON() string { return r.JSON.raw }
func (r *FeatureGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type FeatureListResponse struct {
	// Map of feature keys to enabled state (always true). Disabled features are
	// omitted.
	Features map[string]bool `json:"features"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Features    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r FeatureListResponse) RawJSON() string { return r.JSON.raw }
func (r *FeatureListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
