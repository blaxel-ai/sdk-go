// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blaxel

import (
	"github.com/blaxel-ai/sdk-go/internal/apierror"
	"github.com/blaxel-ai/sdk-go/packages/param"
	"github.com/blaxel-ai/sdk-go/shared"
)

// aliased to make [param.APIUnion] private when embedding
type paramUnion = param.APIUnion

// aliased to make [param.APIObject] private when embedding
type paramObj = param.APIObject

type Error = apierror.Error

// Gateway error codes returned by the Blaxel platform.
const (
	ErrRouteNotFound          = apierror.ErrRouteNotFound
	ErrWorkloadNotFound       = apierror.ErrWorkloadNotFound
	ErrWorkspaceNotFound      = apierror.ErrWorkspaceNotFound
	ErrWorkloadUnavailable    = apierror.ErrWorkloadUnavailable
	ErrAuthenticationRequired = apierror.ErrAuthenticationRequired
	ErrAuthenticationFailed   = apierror.ErrAuthenticationFailed
	ErrForbidden              = apierror.ErrForbidden
	ErrBadRequest             = apierror.ErrBadRequest
	ErrUsageLimitExceeded     = apierror.ErrUsageLimitExceeded
	ErrPolicyViolation        = apierror.ErrPolicyViolation
)

// This is an alias to an internal type.
type EgressGatewayMetadata = shared.EgressGatewayMetadata

// This is an alias to an internal type.
type EgressGatewayMetadataParam = shared.EgressGatewayMetadataParam

// Specification for an egress gateway including region and capacity configuration
//
// This is an alias to an internal type.
type EgressGatewaySpec = shared.EgressGatewaySpec

// Specification for an egress gateway including region and capacity configuration
//
// This is an alias to an internal type.
type EgressGatewaySpecParam = shared.EgressGatewaySpecParam

// This is an alias to an internal type.
type EgressIPMetadata = shared.EgressIPMetadata

// This is an alias to an internal type.
type EgressIPMetadataParam = shared.EgressIPMetadataParam

// Specification for an egress IP including IP family and the provisioned address
//
// This is an alias to an internal type.
type EgressIPSpec = shared.EgressIPSpec

// IP address family, either IPv4 or IPv6
//
// This is an alias to an internal type.
type EgressIPSpecIPFamily = shared.EgressIPSpecIPFamily

// Equals "IPv4"
const EgressIPSpecIPFamilyIPv4 = shared.EgressIPSpecIPFamilyIPv4

// Equals "IPv6"
const EgressIPSpecIPFamilyIPv6 = shared.EgressIPSpecIPFamilyIPv6

// Specification for an egress IP including IP family and the provisioned address
//
// This is an alias to an internal type.
type EgressIPSpecParam = shared.EgressIPSpecParam

// Environment variable with name and value
//
// This is an alias to an internal type.
type Env = shared.Env

// Environment variable with name and value
//
// This is an alias to an internal type.
type EnvParam = shared.EnvParam

// A type of hardware available for deployments
//
// This is an alias to an internal type.
type Flavor = shared.Flavor

// Flavor type (e.g. cpu, gpu)
//
// This is an alias to an internal type.
type FlavorType = shared.FlavorType

// Equals "cpu"
const FlavorTypeCPU = shared.FlavorTypeCPU

// Equals "gpu"
const FlavorTypeGPU = shared.FlavorTypeGPU

// A type of hardware available for deployments
//
// This is an alias to an internal type.
type FlavorParam = shared.FlavorParam
