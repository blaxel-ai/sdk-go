// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blaxel

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"slices"

	"github.com/stainless-sdks/blaxel-go/internal/apijson"
	shimjson "github.com/stainless-sdks/blaxel-go/internal/encoding/json"
	"github.com/stainless-sdks/blaxel-go/internal/requestconfig"
	"github.com/stainless-sdks/blaxel-go/option"
	"github.com/stainless-sdks/blaxel-go/packages/param"
	"github.com/stainless-sdks/blaxel-go/packages/respjson"
)

// CustomdomainService contains methods and other services that help with
// interacting with the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewCustomdomainService] method instead.
type CustomdomainService struct {
	Options []option.RequestOption
}

// NewCustomdomainService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewCustomdomainService(opts ...option.RequestOption) (r CustomdomainService) {
	r = CustomdomainService{}
	r.Options = opts
	return
}

// Create custom domain
func (r *CustomdomainService) New(ctx context.Context, body CustomdomainNewParams, opts ...option.RequestOption) (res *CustomDomain, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "customdomains"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get custom domain
func (r *CustomdomainService) Get(ctx context.Context, domainName string, opts ...option.RequestOption) (res *CustomDomain, err error) {
	opts = slices.Concat(r.Options, opts)
	if domainName == "" {
		err = errors.New("missing required domainName parameter")
		return
	}
	path := fmt.Sprintf("customdomains/%s", domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update custom domain
func (r *CustomdomainService) Update(ctx context.Context, domainName string, body CustomdomainUpdateParams, opts ...option.RequestOption) (res *CustomDomain, err error) {
	opts = slices.Concat(r.Options, opts)
	if domainName == "" {
		err = errors.New("missing required domainName parameter")
		return
	}
	path := fmt.Sprintf("customdomains/%s", domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List all custom domains
func (r *CustomdomainService) List(ctx context.Context, opts ...option.RequestOption) (res *[]CustomDomain, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "customdomains"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete custom domain
func (r *CustomdomainService) Delete(ctx context.Context, domainName string, opts ...option.RequestOption) (res *CustomDomain, err error) {
	opts = slices.Concat(r.Options, opts)
	if domainName == "" {
		err = errors.New("missing required domainName parameter")
		return
	}
	path := fmt.Sprintf("customdomains/%s", domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// Verify custom domain
func (r *CustomdomainService) Verify(ctx context.Context, domainName string, opts ...option.RequestOption) (res *CustomDomain, err error) {
	opts = slices.Concat(r.Options, opts)
	if domainName == "" {
		err = errors.New("missing required domainName parameter")
		return
	}
	path := fmt.Sprintf("customdomains/%s/verify", domainName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return
}

// Custom domain for preview deployments The custom domain represents a base domain
// (e.g., example.com) that will be used to serve preview deployments. Each preview
// will be accessible at a subdomain: <preview-id>.preview.<base-domain> (e.g.,
// abc123.preview.example.com)
type CustomDomain struct {
	// Custom domain metadata
	Metadata CustomDomainMetadata `json:"metadata,required"`
	// Custom domain specification
	Spec CustomDomainSpec `json:"spec,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metadata    respjson.Field
		Spec        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomDomain) RawJSON() string { return r.JSON.raw }
func (r *CustomDomain) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CustomDomain to a CustomDomainParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CustomDomainParam.Overrides()
func (r CustomDomain) ToParam() CustomDomainParam {
	return param.Override[CustomDomainParam](json.RawMessage(r.RawJSON()))
}

// Custom domain metadata
type CustomDomainMetadata struct {
	// Display name for the custom domain
	DisplayName string `json:"displayName"`
	// Labels
	Labels map[string]string `json:"labels"`
	// Domain name (e.g., "example.com")
	Name string `json:"name"`
	// Workspace name
	Workspace string `json:"workspace"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DisplayName respjson.Field
		Labels      respjson.Field
		Name        respjson.Field
		Workspace   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	TimeFields
	OwnerFields
}

// Returns the unmodified JSON received from the API
func (r CustomDomainMetadata) RawJSON() string { return r.JSON.raw }
func (r *CustomDomainMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom domain specification
type CustomDomainSpec struct {
	// CNAME target for the domain
	CnameRecords string `json:"cnameRecords"`
	// Last verification attempt timestamp
	LastVerifiedAt string `json:"lastVerifiedAt"`
	// Region that the custom domain is associated with
	Region string `json:"region"`
	// Current status of the domain (pending, verified, failed)
	//
	// Any of "pending", "verified", "failed".
	Status string `json:"status"`
	// Map of TXT record names to values for domain verification
	TxtRecords map[string]string `json:"txtRecords"`
	// Error message if verification failed
	VerificationError string `json:"verificationError"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CnameRecords      respjson.Field
		LastVerifiedAt    respjson.Field
		Region            respjson.Field
		Status            respjson.Field
		TxtRecords        respjson.Field
		VerificationError respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CustomDomainSpec) RawJSON() string { return r.JSON.raw }
func (r *CustomDomainSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom domain for preview deployments The custom domain represents a base domain
// (e.g., example.com) that will be used to serve preview deployments. Each preview
// will be accessible at a subdomain: <preview-id>.preview.<base-domain> (e.g.,
// abc123.preview.example.com)
//
// The properties Metadata, Spec are required.
type CustomDomainParam struct {
	// Custom domain metadata
	Metadata CustomDomainMetadataParam `json:"metadata,omitzero,required"`
	// Custom domain specification
	Spec CustomDomainSpecParam `json:"spec,omitzero,required"`
	paramObj
}

func (r CustomDomainParam) MarshalJSON() (data []byte, err error) {
	type shadow CustomDomainParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomDomainParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Custom domain metadata
type CustomDomainMetadataParam struct {
	// Display name for the custom domain
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	// Labels
	Labels map[string]string `json:"labels,omitzero"`
	// Domain name (e.g., "example.com")
	Name param.Opt[string] `json:"name,omitzero"`
	// Workspace name
	Workspace param.Opt[string] `json:"workspace,omitzero"`
	TimeFieldsParam
	OwnerFieldsParam
	paramObj
}

func (r CustomDomainMetadataParam) MarshalJSON() (data []byte, err error) {
	type shadow CustomDomainMetadataParam
	return param.MarshalObject(r, (*shadow)(&r))
}

// Custom domain specification
type CustomDomainSpecParam struct {
	// CNAME target for the domain
	CnameRecords param.Opt[string] `json:"cnameRecords,omitzero"`
	// Last verification attempt timestamp
	LastVerifiedAt param.Opt[string] `json:"lastVerifiedAt,omitzero"`
	// Region that the custom domain is associated with
	Region param.Opt[string] `json:"region,omitzero"`
	// Error message if verification failed
	VerificationError param.Opt[string] `json:"verificationError,omitzero"`
	// Current status of the domain (pending, verified, failed)
	//
	// Any of "pending", "verified", "failed".
	Status string `json:"status,omitzero"`
	// Map of TXT record names to values for domain verification
	TxtRecords map[string]string `json:"txtRecords,omitzero"`
	paramObj
}

func (r CustomDomainSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow CustomDomainSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CustomDomainSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[CustomDomainSpecParam](
		"status", "pending", "verified", "failed",
	)
}

type CustomdomainNewParams struct {
	// Custom domain for preview deployments The custom domain represents a base domain
	// (e.g., example.com) that will be used to serve preview deployments. Each preview
	// will be accessible at a subdomain: <preview-id>.preview.<base-domain> (e.g.,
	// abc123.preview.example.com)
	CustomDomain CustomDomainParam
	paramObj
}

func (r CustomdomainNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CustomDomain)
}
func (r *CustomdomainNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CustomDomain)
}

type CustomdomainUpdateParams struct {
	// Custom domain for preview deployments The custom domain represents a base domain
	// (e.g., example.com) that will be used to serve preview deployments. Each preview
	// will be accessible at a subdomain: <preview-id>.preview.<base-domain> (e.g.,
	// abc123.preview.example.com)
	CustomDomain CustomDomainParam
	paramObj
}

func (r CustomdomainUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.CustomDomain)
}
func (r *CustomdomainUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.CustomDomain)
}
