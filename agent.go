// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blaxel

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/blaxel-go/internal/apijson"
	"github.com/stainless-sdks/blaxel-go/internal/apiquery"
	shimjson "github.com/stainless-sdks/blaxel-go/internal/encoding/json"
	"github.com/stainless-sdks/blaxel-go/internal/requestconfig"
	"github.com/stainless-sdks/blaxel-go/option"
	"github.com/stainless-sdks/blaxel-go/packages/param"
	"github.com/stainless-sdks/blaxel-go/packages/respjson"
)

// AgentService contains methods and other services that help with interacting with
// the blaxel API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentService] method instead.
type AgentService struct {
	Options []option.RequestOption
}

// NewAgentService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAgentService(opts ...option.RequestOption) (r AgentService) {
	r = AgentService{}
	r.Options = opts
	return
}

// Create agent by name
func (r *AgentService) New(ctx context.Context, body AgentNewParams, opts ...option.RequestOption) (res *Agent, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "agents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get agent by name
func (r *AgentService) Get(ctx context.Context, agentName string, query AgentGetParams, opts ...option.RequestOption) (res *Agent, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentName == "" {
		err = errors.New("missing required agentName parameter")
		return
	}
	path := fmt.Sprintf("agents/%s", agentName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Update agent by name
func (r *AgentService) Update(ctx context.Context, agentName string, body AgentUpdateParams, opts ...option.RequestOption) (res *Agent, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentName == "" {
		err = errors.New("missing required agentName parameter")
		return
	}
	path := fmt.Sprintf("agents/%s", agentName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// List all agents
func (r *AgentService) List(ctx context.Context, opts ...option.RequestOption) (res *[]Agent, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "agents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Delete agent by name
func (r *AgentService) Delete(ctx context.Context, agentName string, opts ...option.RequestOption) (res *Agent, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentName == "" {
		err = errors.New("missing required agentName parameter")
		return
	}
	path := fmt.Sprintf("agents/%s", agentName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return
}

// List all agent revisions
func (r *AgentService) ListRevisions(ctx context.Context, agentName string, opts ...option.RequestOption) (res *[]RevisionMetadata, err error) {
	opts = slices.Concat(r.Options, opts)
	if agentName == "" {
		err = errors.New("missing required agentName parameter")
		return
	}
	path := fmt.Sprintf("agents/%s/revisions", agentName)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Agent
type Agent struct {
	// Metadata
	Metadata Metadata `json:"metadata,required"`
	// Agent specification for API
	Spec AgentSpec `json:"spec,required"`
	// Core events
	Events []CoreEvent `json:"events"`
	// Agent status
	//
	// Any of "DELETING", "TERMINATED", "FAILED", "DEACTIVATED", "DEACTIVATING",
	// "UPLOADING", "BUILDING", "DEPLOYING", "DEPLOYED".
	Status Status `json:"status"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Metadata    respjson.Field
		Spec        respjson.Field
		Events      respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Agent) RawJSON() string { return r.JSON.raw }
func (r *Agent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Agent to a AgentParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AgentParam.Overrides()
func (r Agent) ToParam() AgentParam {
	return param.Override[AgentParam](json.RawMessage(r.RawJSON()))
}

// Agent
//
// The properties Metadata, Spec are required.
type AgentParam struct {
	// Metadata
	Metadata MetadataParam `json:"metadata,omitzero,required"`
	// Agent specification for API
	Spec AgentSpecParam `json:"spec,omitzero,required"`
	paramObj
}

func (r AgentParam) MarshalJSON() (data []byte, err error) {
	type shadow AgentParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Runtime configuration for Agent
type AgentRuntime struct {
	// The env variables to set in the agent. Should be a list of Kubernetes EnvVar
	// types
	Envs []map[string]any `json:"envs"`
	// The generation of the agent
	//
	// Any of "mk2", "mk3".
	Generation AgentRuntimeGeneration `json:"generation"`
	// The Docker image for the agent
	Image string `json:"image"`
	// The maximum number of replicas for the agent.
	MaxScale int64 `json:"maxScale"`
	// The memory for the agent in MB
	Memory int64 `json:"memory"`
	// The minimum number of replicas for the agent. Can be 0 or 1 (in which case the
	// agent is always running in at least one location).
	MinScale int64 `json:"minScale"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Envs        respjson.Field
		Generation  respjson.Field
		Image       respjson.Field
		MaxScale    respjson.Field
		Memory      respjson.Field
		MinScale    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRuntime) RawJSON() string { return r.JSON.raw }
func (r *AgentRuntime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AgentRuntime to a AgentRuntimeParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AgentRuntimeParam.Overrides()
func (r AgentRuntime) ToParam() AgentRuntimeParam {
	return param.Override[AgentRuntimeParam](json.RawMessage(r.RawJSON()))
}

// The generation of the agent
type AgentRuntimeGeneration string

const (
	AgentRuntimeGenerationMk2 AgentRuntimeGeneration = "mk2"
	AgentRuntimeGenerationMk3 AgentRuntimeGeneration = "mk3"
)

// Runtime configuration for Agent
type AgentRuntimeParam struct {
	// The Docker image for the agent
	Image param.Opt[string] `json:"image,omitzero"`
	// The maximum number of replicas for the agent.
	MaxScale param.Opt[int64] `json:"maxScale,omitzero"`
	// The memory for the agent in MB
	Memory param.Opt[int64] `json:"memory,omitzero"`
	// The minimum number of replicas for the agent. Can be 0 or 1 (in which case the
	// agent is always running in at least one location).
	MinScale param.Opt[int64] `json:"minScale,omitzero"`
	// The env variables to set in the agent. Should be a list of Kubernetes EnvVar
	// types
	Envs []map[string]any `json:"envs,omitzero"`
	// The generation of the agent
	//
	// Any of "mk2", "mk3".
	Generation AgentRuntimeGeneration `json:"generation,omitzero"`
	paramObj
}

func (r AgentRuntimeParam) MarshalJSON() (data []byte, err error) {
	type shadow AgentRuntimeParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentRuntimeParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Agent specification for API
type AgentSpec struct {
	// Enable or disable the resource
	Enabled  bool     `json:"enabled"`
	Policies []string `json:"policies"`
	// Repository
	Repository Repository `json:"repository"`
	// Revision configuration
	Revision RevisionConfiguration `json:"revision"`
	// Runtime configuration for Agent
	Runtime AgentRuntime `json:"runtime"`
	// Triggers to use your agent
	Triggers []Trigger `json:"triggers"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Enabled     respjson.Field
		Policies    respjson.Field
		Repository  respjson.Field
		Revision    respjson.Field
		Runtime     respjson.Field
		Triggers    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentSpec) RawJSON() string { return r.JSON.raw }
func (r *AgentSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AgentSpec to a AgentSpecParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AgentSpecParam.Overrides()
func (r AgentSpec) ToParam() AgentSpecParam {
	return param.Override[AgentSpecParam](json.RawMessage(r.RawJSON()))
}

// Agent specification for API
type AgentSpecParam struct {
	// Enable or disable the resource
	Enabled  param.Opt[bool] `json:"enabled,omitzero"`
	Policies []string        `json:"policies,omitzero"`
	// Repository
	Repository RepositoryParam `json:"repository,omitzero"`
	// Revision configuration
	Revision RevisionConfigurationParam `json:"revision,omitzero"`
	// Runtime configuration for Agent
	Runtime AgentRuntimeParam `json:"runtime,omitzero"`
	// Triggers to use your agent
	Triggers []TriggerParam `json:"triggers,omitzero"`
	paramObj
}

func (r AgentSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow AgentSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Core event
type CoreEvent struct {
	// Canary revisionID link to the event
	CanaryRevision string `json:"canaryRevision"`
	// Event message
	Message string `json:"message"`
	// RevisionID link to the event
	Revision string `json:"revision"`
	// Event status
	Status string `json:"status"`
	// Event time
	Time string `json:"time"`
	// Event type
	Type string `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CanaryRevision respjson.Field
		Message        respjson.Field
		Revision       respjson.Field
		Status         respjson.Field
		Time           respjson.Field
		Type           respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CoreEvent) RawJSON() string { return r.JSON.raw }
func (r *CoreEvent) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this CoreEvent to a CoreEventParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// CoreEventParam.Overrides()
func (r CoreEvent) ToParam() CoreEventParam {
	return param.Override[CoreEventParam](json.RawMessage(r.RawJSON()))
}

// Core event
type CoreEventParam struct {
	// Canary revisionID link to the event
	CanaryRevision param.Opt[string] `json:"canaryRevision,omitzero"`
	// Event message
	Message param.Opt[string] `json:"message,omitzero"`
	// RevisionID link to the event
	Revision param.Opt[string] `json:"revision,omitzero"`
	// Event status
	Status param.Opt[string] `json:"status,omitzero"`
	// Event time
	Time param.Opt[string] `json:"time,omitzero"`
	// Event type
	Type param.Opt[string] `json:"type,omitzero"`
	paramObj
}

func (r CoreEventParam) MarshalJSON() (data []byte, err error) {
	type shadow CoreEventParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *CoreEventParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type Metadata struct {
	// Model name
	Name string `json:"name,required"`
	// The date and time when the resource was created
	CreatedAt string `json:"createdAt"`
	// The user or service account who created the resource
	CreatedBy string `json:"createdBy"`
	// Model display name
	DisplayName string `json:"displayName"`
	// Labels
	Labels map[string]string `json:"labels"`
	// Plan
	Plan string `json:"plan"`
	// The date and time when the resource was updated
	UpdatedAt string `json:"updatedAt"`
	// The user or service account who updated the resource
	UpdatedBy string `json:"updatedBy"`
	// URL
	URL string `json:"url"`
	// Workspace name
	Workspace string `json:"workspace"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Name        respjson.Field
		CreatedAt   respjson.Field
		CreatedBy   respjson.Field
		DisplayName respjson.Field
		Labels      respjson.Field
		Plan        respjson.Field
		UpdatedAt   respjson.Field
		UpdatedBy   respjson.Field
		URL         respjson.Field
		Workspace   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Metadata) RawJSON() string { return r.JSON.raw }
func (r *Metadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Metadata to a MetadataParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// MetadataParam.Overrides()
func (r Metadata) ToParam() MetadataParam {
	return param.Override[MetadataParam](json.RawMessage(r.RawJSON()))
}

// The property Name is required.
type MetadataParam struct {
	// Model name
	Name string `json:"name,required"`
	// Model display name
	DisplayName param.Opt[string] `json:"displayName,omitzero"`
	// Labels
	Labels map[string]string `json:"labels,omitzero"`
	paramObj
}

func (r MetadataParam) MarshalJSON() (data []byte, err error) {
	type shadow MetadataParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *MetadataParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Repository
type Repository struct {
	// Repository type
	Type string `json:"type"`
	// Repository URL
	URL string `json:"url"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Type        respjson.Field
		URL         respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Repository) RawJSON() string { return r.JSON.raw }
func (r *Repository) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Repository to a RepositoryParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// RepositoryParam.Overrides()
func (r Repository) ToParam() RepositoryParam {
	return param.Override[RepositoryParam](json.RawMessage(r.RawJSON()))
}

// Repository
type RepositoryParam struct {
	// Repository type
	Type param.Opt[string] `json:"type,omitzero"`
	// Repository URL
	URL param.Opt[string] `json:"url,omitzero"`
	paramObj
}

func (r RepositoryParam) MarshalJSON() (data []byte, err error) {
	type shadow RepositoryParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RepositoryParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Revision configuration
type RevisionConfiguration struct {
	// Active revision id
	Active string `json:"active"`
	// Canary revision id
	Canary string `json:"canary"`
	// Canary revision percent
	CanaryPercent int64 `json:"canaryPercent"`
	// Sticky session TTL in seconds (0 = disabled)
	StickySessionTtl int64 `json:"stickySessionTtl"`
	// Traffic percentage
	Traffic int64 `json:"traffic"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Active           respjson.Field
		Canary           respjson.Field
		CanaryPercent    respjson.Field
		StickySessionTtl respjson.Field
		Traffic          respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RevisionConfiguration) RawJSON() string { return r.JSON.raw }
func (r *RevisionConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this RevisionConfiguration to a RevisionConfigurationParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// RevisionConfigurationParam.Overrides()
func (r RevisionConfiguration) ToParam() RevisionConfigurationParam {
	return param.Override[RevisionConfigurationParam](json.RawMessage(r.RawJSON()))
}

// Revision configuration
type RevisionConfigurationParam struct {
	// Active revision id
	Active param.Opt[string] `json:"active,omitzero"`
	// Canary revision id
	Canary param.Opt[string] `json:"canary,omitzero"`
	// Canary revision percent
	CanaryPercent param.Opt[int64] `json:"canaryPercent,omitzero"`
	// Sticky session TTL in seconds (0 = disabled)
	StickySessionTtl param.Opt[int64] `json:"stickySessionTtl,omitzero"`
	// Traffic percentage
	Traffic param.Opt[int64] `json:"traffic,omitzero"`
	paramObj
}

func (r RevisionConfigurationParam) MarshalJSON() (data []byte, err error) {
	type shadow RevisionConfigurationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *RevisionConfigurationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Revision metadata
type RevisionMetadata struct {
	// Revision ID
	ID string `json:"id"`
	// Is the revision active
	Active bool `json:"active"`
	// Is the revision canary
	Canary bool `json:"canary"`
	// Revision created at
	CreatedAt string `json:"createdAt"`
	// Revision created by
	CreatedBy string `json:"createdBy"`
	// Is the revision previous active
	PreviousActive bool `json:"previousActive"`
	// Status of the revision
	Status string `json:"status"`
	// Percent of traffic to the revision
	TrafficPercent int64 `json:"trafficPercent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		Active         respjson.Field
		Canary         respjson.Field
		CreatedAt      respjson.Field
		CreatedBy      respjson.Field
		PreviousActive respjson.Field
		Status         respjson.Field
		TrafficPercent respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RevisionMetadata) RawJSON() string { return r.JSON.raw }
func (r *RevisionMetadata) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of a resource
type Status string

const (
	StatusDeleting     Status = "DELETING"
	StatusTerminated   Status = "TERMINATED"
	StatusFailed       Status = "FAILED"
	StatusDeactivated  Status = "DEACTIVATED"
	StatusDeactivating Status = "DEACTIVATING"
	StatusUploading    Status = "UPLOADING"
	StatusBuilding     Status = "BUILDING"
	StatusDeploying    Status = "DEPLOYING"
	StatusDeployed     Status = "DEPLOYED"
)

// Trigger configuration
type Trigger struct {
	// The id of the trigger
	ID string `json:"id"`
	// The configuration of the trigger
	Configuration TriggerConfiguration `json:"configuration"`
	// Enable or disable the trigger (default: true)
	Enabled bool `json:"enabled"`
	// The type of trigger, can be http or http-async
	//
	// Any of "http", "http-async", "cron".
	Type TriggerType `json:"type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		Configuration respjson.Field
		Enabled       respjson.Field
		Type          respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Trigger) RawJSON() string { return r.JSON.raw }
func (r *Trigger) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this Trigger to a TriggerParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TriggerParam.Overrides()
func (r Trigger) ToParam() TriggerParam {
	return param.Override[TriggerParam](json.RawMessage(r.RawJSON()))
}

// The configuration of the trigger
type TriggerConfiguration struct {
	// The authentication type of the trigger
	AuthenticationType string `json:"authenticationType"`
	// The callback secret for async triggers (auto-generated, encrypted)
	CallbackSecret string `json:"callbackSecret"`
	// The callback URL for async triggers (optional)
	CallbackURL string `json:"callbackUrl"`
	// The path of the trigger
	Path string `json:"path"`
	// The retry of the trigger
	Retry int64 `json:"retry"`
	// The schedule of the trigger, cron expression \* \* \* \* \*
	Schedule string `json:"schedule"`
	// The tasks configuration of the cronjob
	Tasks []map[string]any `json:"tasks"`
	// The timeout in seconds for async triggers (max 900s, MK3 only)
	Timeout int64 `json:"timeout"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		AuthenticationType respjson.Field
		CallbackSecret     respjson.Field
		CallbackURL        respjson.Field
		Path               respjson.Field
		Retry              respjson.Field
		Schedule           respjson.Field
		Tasks              respjson.Field
		Timeout            respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TriggerConfiguration) RawJSON() string { return r.JSON.raw }
func (r *TriggerConfiguration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of trigger, can be http or http-async
type TriggerType string

const (
	TriggerTypeHTTP      TriggerType = "http"
	TriggerTypeHTTPAsync TriggerType = "http-async"
	TriggerTypeCron      TriggerType = "cron"
)

// Trigger configuration
type TriggerParam struct {
	// The id of the trigger
	ID param.Opt[string] `json:"id,omitzero"`
	// Enable or disable the trigger (default: true)
	Enabled param.Opt[bool] `json:"enabled,omitzero"`
	// The configuration of the trigger
	Configuration TriggerConfigurationParam `json:"configuration,omitzero"`
	// The type of trigger, can be http or http-async
	//
	// Any of "http", "http-async", "cron".
	Type TriggerType `json:"type,omitzero"`
	paramObj
}

func (r TriggerParam) MarshalJSON() (data []byte, err error) {
	type shadow TriggerParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TriggerParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration of the trigger
type TriggerConfigurationParam struct {
	// The authentication type of the trigger
	AuthenticationType param.Opt[string] `json:"authenticationType,omitzero"`
	// The callback secret for async triggers (auto-generated, encrypted)
	CallbackSecret param.Opt[string] `json:"callbackSecret,omitzero"`
	// The callback URL for async triggers (optional)
	CallbackURL param.Opt[string] `json:"callbackUrl,omitzero"`
	// The path of the trigger
	Path param.Opt[string] `json:"path,omitzero"`
	// The retry of the trigger
	Retry param.Opt[int64] `json:"retry,omitzero"`
	// The schedule of the trigger, cron expression \* \* \* \* \*
	Schedule param.Opt[string] `json:"schedule,omitzero"`
	// The timeout in seconds for async triggers (max 900s, MK3 only)
	Timeout param.Opt[int64] `json:"timeout,omitzero"`
	// The tasks configuration of the cronjob
	Tasks []map[string]any `json:"tasks,omitzero"`
	paramObj
}

func (r TriggerConfigurationParam) MarshalJSON() (data []byte, err error) {
	type shadow TriggerConfigurationParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TriggerConfigurationParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentNewParams struct {
	// Agent
	Agent AgentParam
	paramObj
}

func (r AgentNewParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Agent)
}
func (r *AgentNewParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Agent)
}

type AgentGetParams struct {
	// Show secret values (admin only)
	ShowSecrets param.Opt[bool] `query:"show_secrets,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AgentGetParams]'s query parameters as `url.Values`.
func (r AgentGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type AgentUpdateParams struct {
	// Agent
	Agent AgentParam
	paramObj
}

func (r AgentUpdateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.Agent)
}
func (r *AgentUpdateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Agent)
}
