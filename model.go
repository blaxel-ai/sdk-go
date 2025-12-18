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

// Creates a model.
func (r *ModelService) New(ctx context.Context, body ModelNewParams, opts ...option.RequestOption) (res *Model, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "models"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Returns a model by name.
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

// Update a model by name.
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

// Returns a list of all models in the workspace.
func (r *ModelService) List(ctx context.Context, opts ...option.RequestOption) (res *[]Model, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "models"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Deletes a model by name.
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
	// Logical object representing a model
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
	// Logical object representing a model
	Model ModelParam
	paramObj
}

func (r ModelUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Model)
}
func (r *ModelUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Model)
}
