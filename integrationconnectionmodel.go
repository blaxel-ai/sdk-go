// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blaxel

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/blaxel-ai/sdk-go/internal/requestconfig"
	"github.com/blaxel-ai/sdk-go/option"
)

// IntegrationConnectionModelService contains methods and other services that help
// with interacting with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewIntegrationConnectionModelService] method instead.
type IntegrationConnectionModelService struct {
	Options []option.RequestOption
}

// NewIntegrationConnectionModelService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewIntegrationConnectionModelService(opts ...option.RequestOption) (r IntegrationConnectionModelService) {
	r = IntegrationConnectionModelService{}
	r.Options = opts
	return
}

// Returns a model for an integration connection by ID.
func (r *IntegrationConnectionModelService) Get(ctx context.Context, modelID string, query IntegrationConnectionModelGetParams, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if query.ConnectionName == "" {
		err = errors.New("missing required connectionName parameter")
		return
	}
	if modelID == "" {
		err = errors.New("missing required modelId parameter")
		return
	}
	path := fmt.Sprintf("integrations/connections/%s/models/%s", query.ConnectionName, modelID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

// Returns a list of all models for an integration connection.
func (r *IntegrationConnectionModelService) List(ctx context.Context, connectionName string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if connectionName == "" {
		err = errors.New("missing required connectionName parameter")
		return
	}
	path := fmt.Sprintf("integrations/connections/%s/models", connectionName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, nil, opts...)
	return
}

type IntegrationConnectionModelGetParams struct {
	ConnectionName string `path:"connectionName,required" json:"-"`
	paramObj
}
