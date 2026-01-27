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
)

func TestFunctionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Functions.New(context.TODO(), blaxel.FunctionNewParams{
		Function: blaxel.FunctionParam{
			Metadata: blaxel.MetadataParam{
				Name:        "my-resource",
				DisplayName: blaxel.String("My Resource"),
				Labels: map[string]string{
					"foo": "string",
				},
			},
			Spec: blaxel.FunctionSpecParam{
				Enabled:                blaxel.Bool(true),
				IntegrationConnections: []string{"string"},
				Policies:               []string{"string"},
				Revision: blaxel.RevisionConfigurationParam{
					Active:           blaxel.String("rev-abc123"),
					Canary:           blaxel.String("canary"),
					CanaryPercent:    blaxel.Int(10),
					StickySessionTtl: blaxel.Int(0),
					Traffic:          blaxel.Int(100),
				},
				Runtime: blaxel.FunctionRuntimeParam{
					Envs: []blaxel.FunctionRuntimeEnvParam{{
						Name:   blaxel.String("MY_ENV_VAR"),
						Secret: blaxel.Bool(true),
						Value:  blaxel.String("my-value"),
					}},
					Generation: blaxel.FunctionRuntimeGenerationMk3,
					Image:      blaxel.String("image"),
					MaxScale:   blaxel.Int(10),
					Memory:     blaxel.Int(2048),
					MinScale:   blaxel.Int(0),
					Transport:  blaxel.FunctionRuntimeTransportHTTPStream,
				},
				Triggers: []blaxel.TriggerParam{{
					ID: blaxel.String("trigger-1"),
					Configuration: blaxel.TriggerConfigurationParam{
						AuthenticationType: blaxel.String("blaxel"),
						CallbackURL:        blaxel.String("https://myapp.com/webhook"),
						Path:               blaxel.String("/invoke"),
						Retry:              blaxel.Int(3),
						Schedule:           blaxel.String("0 * * * *"),
						Tasks:              []any{map[string]any{}},
						Timeout:            blaxel.Int(300),
					},
					Enabled: blaxel.Bool(true),
					Type:    blaxel.TriggerTypeHTTP,
				}},
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

func TestFunctionGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Functions.Get(
		context.TODO(),
		"functionName",
		blaxel.FunctionGetParams{
			ShowSecrets: blaxel.Bool(true),
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

func TestFunctionUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Functions.Update(
		context.TODO(),
		"functionName",
		blaxel.FunctionUpdateParams{
			Function: blaxel.FunctionParam{
				Metadata: blaxel.MetadataParam{
					Name:        "my-resource",
					DisplayName: blaxel.String("My Resource"),
					Labels: map[string]string{
						"foo": "string",
					},
				},
				Spec: blaxel.FunctionSpecParam{
					Enabled:                blaxel.Bool(true),
					IntegrationConnections: []string{"string"},
					Policies:               []string{"string"},
					Revision: blaxel.RevisionConfigurationParam{
						Active:           blaxel.String("rev-abc123"),
						Canary:           blaxel.String("canary"),
						CanaryPercent:    blaxel.Int(10),
						StickySessionTtl: blaxel.Int(0),
						Traffic:          blaxel.Int(100),
					},
					Runtime: blaxel.FunctionRuntimeParam{
						Envs: []blaxel.FunctionRuntimeEnvParam{{
							Name:   blaxel.String("MY_ENV_VAR"),
							Secret: blaxel.Bool(true),
							Value:  blaxel.String("my-value"),
						}},
						Generation: blaxel.FunctionRuntimeGenerationMk3,
						Image:      blaxel.String("image"),
						MaxScale:   blaxel.Int(10),
						Memory:     blaxel.Int(2048),
						MinScale:   blaxel.Int(0),
						Transport:  blaxel.FunctionRuntimeTransportHTTPStream,
					},
					Triggers: []blaxel.TriggerParam{{
						ID: blaxel.String("trigger-1"),
						Configuration: blaxel.TriggerConfigurationParam{
							AuthenticationType: blaxel.String("blaxel"),
							CallbackURL:        blaxel.String("https://myapp.com/webhook"),
							Path:               blaxel.String("/invoke"),
							Retry:              blaxel.Int(3),
							Schedule:           blaxel.String("0 * * * *"),
							Tasks:              []any{map[string]any{}},
							Timeout:            blaxel.Int(300),
						},
						Enabled: blaxel.Bool(true),
						Type:    blaxel.TriggerTypeHTTP,
					}},
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

func TestFunctionList(t *testing.T) {
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
	_, err := client.Functions.List(context.TODO())
	if err != nil {
		var apierr *blaxel.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFunctionDelete(t *testing.T) {
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
	_, err := client.Functions.Delete(context.TODO(), "functionName")
	if err != nil {
		var apierr *blaxel.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestFunctionListRevisions(t *testing.T) {
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
	_, err := client.Functions.ListRevisions(context.TODO(), "functionName")
	if err != nil {
		var apierr *blaxel.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
