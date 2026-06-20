// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blaxel_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/blaxel-ai/sdk-go"
	"github.com/blaxel-ai/sdk-go/internal/testutil"
	"github.com/blaxel-ai/sdk-go/option"
	"github.com/blaxel-ai/sdk-go/shared"
)

func TestPolicyNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := blaxel.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Policies.New(context.TODO(), blaxel.PolicyNewParams{
		Policy: blaxel.PolicyParam{
			Metadata: blaxel.MetadataParam{
				Name:        "my-resource",
				DisplayName: blaxel.String("My Resource"),
				ExternalID:  blaxel.String("my-session-123"),
				Labels: map[string]string{
					"foo": "string",
				},
			},
			Spec: blaxel.PolicySpecParam{
				Flavors: []shared.FlavorParam{{
					Name: blaxel.String("t4"),
					Type: shared.FlavorTypeCPU,
				}},
				Locations: []blaxel.PolicyLocationParam{{
					Name: blaxel.String("EU"),
					Type: blaxel.PolicyLocationTypeContinent,
				}},
				MaxTokens: blaxel.PolicyMaxTokensParam{
					Granularity:          blaxel.String("minute"),
					Input:                blaxel.Int(10000),
					Output:               blaxel.Int(5000),
					RatioInputOverOutput: blaxel.Int(2),
					Step:                 blaxel.Int(1),
					Total:                blaxel.Int(15000),
				},
				ResourceTypes: []string{"model"},
				Sandbox:       blaxel.Bool(false),
				Type:          blaxel.PolicySpecTypeLocation,
			},
			Usage: blaxel.PolicyUsageParam{
				Agents:    blaxel.Int(0),
				Functions: blaxel.Int(0),
				Jobs:      blaxel.Int(0),
				Models:    blaxel.Int(0),
				Sandboxes: blaxel.Int(0),
			},
		},
	})
	if err != nil {
		var apierr *blaxel.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPolicyGet(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := blaxel.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Policies.Get(context.TODO(), "policyName")
	if err != nil {
		var apierr *blaxel.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPolicyUpdateWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := blaxel.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Policies.Update(
		context.TODO(),
		"policyName",
		blaxel.PolicyUpdateParams{
			Policy: blaxel.PolicyParam{
				Metadata: blaxel.MetadataParam{
					Name:        "my-resource",
					DisplayName: blaxel.String("My Resource"),
					ExternalID:  blaxel.String("my-session-123"),
					Labels: map[string]string{
						"foo": "string",
					},
				},
				Spec: blaxel.PolicySpecParam{
					Flavors: []shared.FlavorParam{{
						Name: blaxel.String("t4"),
						Type: shared.FlavorTypeCPU,
					}},
					Locations: []blaxel.PolicyLocationParam{{
						Name: blaxel.String("EU"),
						Type: blaxel.PolicyLocationTypeContinent,
					}},
					MaxTokens: blaxel.PolicyMaxTokensParam{
						Granularity:          blaxel.String("minute"),
						Input:                blaxel.Int(10000),
						Output:               blaxel.Int(5000),
						RatioInputOverOutput: blaxel.Int(2),
						Step:                 blaxel.Int(1),
						Total:                blaxel.Int(15000),
					},
					ResourceTypes: []string{"model"},
					Sandbox:       blaxel.Bool(false),
					Type:          blaxel.PolicySpecTypeLocation,
				},
				Usage: blaxel.PolicyUsageParam{
					Agents:    blaxel.Int(0),
					Functions: blaxel.Int(0),
					Jobs:      blaxel.Int(0),
					Models:    blaxel.Int(0),
					Sandboxes: blaxel.Int(0),
				},
			},
		},
	)
	if err != nil {
		var apierr *blaxel.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPolicyListWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := blaxel.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Policies.List(context.TODO(), blaxel.PolicyListParams{
		Anchor: blaxel.PolicyListParamsAnchorEnd,
		Cursor: blaxel.String("cursor"),
		Limit:  blaxel.Int(1),
		Q:      blaxel.String("q"),
		Sort:   blaxel.PolicyListParamsSortCreatedAtDesc,
	})
	if err != nil {
		var apierr *blaxel.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPolicyDelete(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := blaxel.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Policies.Delete(context.TODO(), "policyName")
	if err != nil {
		var apierr *blaxel.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPolicyListUsages(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := blaxel.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Policies.ListUsages(context.TODO(), "policyName")
	if err != nil {
		var apierr *blaxel.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
