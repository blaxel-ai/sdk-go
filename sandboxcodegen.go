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

// Uses the configured LLM provider (Relace or Morph) to apply a code edit to the
// original content.
//
// To use this endpoint as an agent tool, follow these guidelines:
//
// Use this tool to make an edit to an existing file. This will be read by a less
// intelligent model, which will quickly apply the edit. You should make it clear
// what the edit is, while also minimizing the unchanged code you write.
//
// When writing the edit, you should specify each edit in sequence, with the
// special comment "// ... existing code ..." to represent unchanged code in
// between edited lines.
//
// Example format: // ... existing code ... FIRST_EDIT // ... existing code ...
// SECOND_EDIT // ... existing code ... THIRD_EDIT // ... existing code ...
//
// You should still bias towards repeating as few lines of the original file as
// possible to convey the change. But, each edit should contain minimally
// sufficient context of unchanged lines around the code you're editing to resolve
// ambiguity.
//
// DO NOT omit spans of pre-existing code (or comments) without using the "// ...
// existing code ..." comment to indicate its absence. If you omit the existing
// code comment, the model may inadvertently delete these lines.
//
// If you plan on deleting a section, you must provide context before and after to
// delete it. If the initial code is "Block 1\nBlock 2\nBlock 3", and you want to
// remove Block 2, you would output "// ... existing code ...\nBlock 1\nBlock 3\n//
// ... existing code ...".
//
// Make sure it is clear what the edit should be, and where it should be applied.
// Make edits to a file in a single edit_file call instead of multiple edit_file
// calls to the same file. The apply model can handle many distinct edits at once.
func (r *SandboxCodegenService) Fastapply(ctx context.Context, filePath string, body SandboxCodegenFastapplyParams, opts ...option.RequestOption) (res *ApplyEditResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if filePath == "" {
		err = errors.New("missing required filePath parameter")
		return
	}
	path := fmt.Sprintf("codegen/fastapply/%s", filePath)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, body, &res, opts...)
	return
}

// Uses Relace's code reranking model to find the most relevant files for a given
// query. This is useful as a first pass in agentic exploration to narrow down the
// search space.
//
// Based on: https://docs.relace.ai/docs/code-reranker/agent
//
// Query Construction: The query can be a short question or a more detailed
// conversation with the user request included. For a first pass, use the full
// conversation; for subsequent calls, use more targeted questions.
//
// Token Limit and Score Threshold: For 200k token context models like Claude 4
// Sonnet, recommended defaults are scoreThreshold=0.5 and tokenLimit=30000.
//
// The response will be a list of file paths and contents ordered from most
// relevant to least relevant.
func (r *SandboxCodegenService) Reranking(ctx context.Context, filePath string, query SandboxCodegenRerankingParams, opts ...option.RequestOption) (res *RerankingResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if filePath == "" {
		err = errors.New("missing required filePath parameter")
		return
	}
	path := fmt.Sprintf("codegen/reranking/%s", filePath)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// The property CodeEdit is required.
type ApplyEditRequestParam struct {
	CodeEdit string            `json:"codeEdit,required"`
	Model    param.Opt[string] `json:"model,omitzero"`
	paramObj
}

func (r ApplyEditRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow ApplyEditRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ApplyEditRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ApplyEditResponse struct {
	Message         string `json:"message"`
	OriginalContent string `json:"originalContent"`
	Path            string `json:"path"`
	Provider        string `json:"provider"`
	Success         bool   `json:"success"`
	UpdatedContent  string `json:"updatedContent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Message         respjson.Field
		OriginalContent respjson.Field
		Path            respjson.Field
		Provider        respjson.Field
		Success         respjson.Field
		UpdatedContent  respjson.Field
		ExtraFields     map[string]respjson.Field
		raw             string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ApplyEditResponse) RawJSON() string { return r.JSON.raw }
func (r *ApplyEditResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RerankingResponse struct {
	Files   []RerankingResponseFile `json:"files"`
	Message string                  `json:"message"`
	Success bool                    `json:"success"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Files       respjson.Field
		Message     respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RerankingResponse) RawJSON() string { return r.JSON.raw }
func (r *RerankingResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type RerankingResponseFile struct {
	Content string  `json:"content"`
	Path    string  `json:"path"`
	Score   float64 `json:"score"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		Path        respjson.Field
		Score       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r RerankingResponseFile) RawJSON() string { return r.JSON.raw }
func (r *RerankingResponseFile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SandboxCodegenFastapplyParams struct {
	ApplyEditRequest ApplyEditRequestParam
	paramObj
}

func (r SandboxCodegenFastapplyParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.ApplyEditRequest)
}
func (r *SandboxCodegenFastapplyParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.ApplyEditRequest)
}

type SandboxCodegenRerankingParams struct {
	// Natural language query to search for
	Query string `query:"query,required" json:"-"`
	// Regex pattern to filter files (e.g., .\*\\.ts$ for TypeScript files)
	FilePattern param.Opt[string] `query:"filePattern,omitzero" json:"-"`
	// Minimum relevance score (default: 0.5)
	ScoreThreshold param.Opt[float64] `query:"scoreThreshold,omitzero" json:"-"`
	// Maximum tokens to return (default: 30000)
	TokenLimit param.Opt[int64] `query:"tokenLimit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [SandboxCodegenRerankingParams]'s query parameters as
// `url.Values`.
func (r SandboxCodegenRerankingParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
