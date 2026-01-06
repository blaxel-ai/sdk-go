// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blaxel

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	shimjson "github.com/stainless-sdks/blaxel-go/internal/encoding/json"
	"github.com/stainless-sdks/blaxel-go/internal/requestconfig"
	"github.com/stainless-sdks/blaxel-go/option"
)

// ModelService contains methods and other services that help with interacting with
// the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewModelService] method instead.
type ModelService struct {
	Options []option.RequestOption
}

// NewModelService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewModelService(opts ...option.RequestOption) (r ModelService) {
	r = ModelService{}
	r.Options = opts
	return
}

// Creates a new model gateway endpoint that proxies requests to an external LLM
// provider. Requires an integration connection with valid API credentials for the
// target provider.
func (r *ModelService) New(ctx context.Context, body ModelNewParams, opts ...option.RequestOption) (res *Model, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "models"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns detailed information about a model gateway endpoint including its
// provider configuration, integration connection, and usage status.
func (r *ModelService) Get(ctx context.Context, modelName string, opts ...option.RequestOption) (res *Model, err error) {
	opts = slices.Concat(r.Options, opts)
	if modelName == "" {
		err = errors.New("missing required modelName parameter")
		return
	}
	path := fmt.Sprintf("models/%s", modelName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Updates a model gateway endpoint's configuration. Changes to provider settings
// or integration connection take effect immediately.
func (r *ModelService) Update(ctx context.Context, modelName string, body ModelUpdateParams, opts ...option.RequestOption) (res *Model, err error) {
	opts = slices.Concat(r.Options, opts)
	if modelName == "" {
		err = errors.New("missing required modelName parameter")
		return
	}
	path := fmt.Sprintf("models/%s", modelName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Returns all model gateway endpoints configured in the workspace. Each model
// represents a proxy to an external LLM provider (OpenAI, Anthropic, etc.) with
// unified access control.
func (r *ModelService) List(ctx context.Context, opts ...option.RequestOption) (res *[]Model, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "models"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Permanently deletes a model gateway endpoint. Any agents or applications using
// this endpoint will need to be updated to use a different model.
func (r *ModelService) Delete(ctx context.Context, modelName string, opts ...option.RequestOption) (res *Model, err error) {
	opts = slices.Concat(r.Options, opts)
	if modelName == "" {
		err = errors.New("missing required modelName parameter")
		return
	}
	path := fmt.Sprintf("models/%s", modelName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Returns revisions for a model by name.
func (r *ModelService) ListRevisions(ctx context.Context, modelName string, opts ...option.RequestOption) (res *[]RevisionMetadata, err error) {
	opts = slices.Concat(r.Options, opts)
	if modelName == "" {
		err = errors.New("missing required modelName parameter")
		return
	}
	path := fmt.Sprintf("models/%s/revisions", modelName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ModelNewParams struct {
	// Gateway endpoint to external LLM provider APIs (OpenAI, Anthropic, etc.) with
	// unified access control, credentials management, and usage tracking.
	Model ModelParam
	paramObj
}

func (r ModelNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Model)
}
func (r *ModelNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Model)
}

type ModelUpdateParams struct {
	// Gateway endpoint to external LLM provider APIs (OpenAI, Anthropic, etc.) with
	// unified access control, credentials management, and usage tracking.
	Model ModelParam
	paramObj
}

func (r ModelUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Model)
}
func (r *ModelUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Model)
}
