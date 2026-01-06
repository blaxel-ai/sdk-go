// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blaxel

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/blaxel-go/internal/apijson"
	"github.com/stainless-sdks/blaxel-go/internal/requestconfig"
	"github.com/stainless-sdks/blaxel-go/option"
	"github.com/stainless-sdks/blaxel-go/packages/respjson"
)

// TemplateService contains methods and other services that help with interacting
// with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTemplateService] method instead.
type TemplateService struct {
	Options []option.RequestOption
}

// NewTemplateService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTemplateService(opts ...option.RequestOption) (r TemplateService) {
	r = TemplateService{}
	r.Options = opts
	return
}

// Returns detailed information about a deployment template including its
// configuration, source code reference, and available parameters.
func (r *TemplateService) Get(ctx context.Context, templateName string, opts ...option.RequestOption) (res *Template, err error) {
	opts = slices.Concat(r.Options, opts)
	if templateName == "" {
		err = errors.New("missing required templateName parameter")
		return
	}
	path := fmt.Sprintf("templates/%s", templateName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Returns all deployment templates available for creating agents, functions, and
// other resources with pre-configured settings and code.
func (r *TemplateService) List(ctx context.Context, opts ...option.RequestOption) (res *[]Template, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "templates"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Blaxel template
type Template struct {
	// Default branch of the template
	DefaultBranch string `json:"defaultBranch"`
	// Description of the template
	Description string `json:"description"`
	// Number of downloads/clones of the repository
	DownloadCount int64 `json:"downloadCount"`
	// Number of forks the repository has
	ForksCount int64 `json:"forksCount"`
	// URL to the template's icon
	Icon string `json:"icon"`
	// URL to the template's icon in dark mode
	IconDark string `json:"iconDark"`
	// Name of the template
	Name string `json:"name"`
	// SHA of the variable
	Sha string `json:"sha"`
	// Number of stars the repository has
	StarCount int64 `json:"starCount"`
	// Topic of the template
	Topics []string `json:"topics"`
	// URL of the template
	URL string `json:"url"`
	// Variables of the template
	Variables []TemplateVariable `json:"variables"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DefaultBranch respjson.Field
		Description   respjson.Field
		DownloadCount respjson.Field
		ForksCount    respjson.Field
		Icon          respjson.Field
		IconDark      respjson.Field
		Name          respjson.Field
		Sha           respjson.Field
		StarCount     respjson.Field
		Topics        respjson.Field
		URL           respjson.Field
		Variables     respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Template) RawJSON() string { return r.JSON.raw }
func (r *Template) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Blaxel template variable
type TemplateVariable struct {
	// Description of the variable
	Description string `json:"description"`
	// Integration of the variable
	Integration string `json:"integration"`
	// Name of the variable
	Name string `json:"name"`
	// Path of the variable
	Path string `json:"path"`
	// Whether the variable is secret
	Secret bool `json:"secret"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		Integration respjson.Field
		Name        respjson.Field
		Path        respjson.Field
		Secret      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TemplateVariable) RawJSON() string { return r.JSON.raw }
func (r *TemplateVariable) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
