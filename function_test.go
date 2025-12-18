// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blaxel_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/stainless-sdks/blaxel-go"
	"github.com/stainless-sdks/blaxel-go/internal/testutil"
	"github.com/stainless-sdks/blaxel-go/option"
)

func TestFunctionNewWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := blaxel.NewClient(
		option.WithBaseURL(baseURL),
		option.WithClientID("My Client ID"),
		option.WithClientSecret("My Client Secret"),
	)
	_, err := client.Functions.New(context.TODO(), blaxel.FunctionNewParams{
		Function: blaxel.FunctionParam{
			Metadata: blaxel.MetadataParam{
				Name:        "name",
				DisplayName: blaxel.String("displayName"),
				Labels: map[string]string{
					"foo": "string",
				},
			},
			Spec: blaxel.FunctionSpecParam{
				Enabled:                blaxel.Bool(true),
				IntegrationConnections: []string{"string"},
				Policies:               []string{"string"},
				Revision: blaxel.RevisionConfigurationParam{
					Active:           blaxel.String("active"),
					Canary:           blaxel.String("canary"),
					CanaryPercent:    blaxel.Int(0),
					StickySessionTtl: blaxel.Int(0),
					Traffic:          blaxel.Int(0),
				},
				Runtime: blaxel.FunctionRuntimeParam{
					Envs: []map[string]any{{
						"foo": "bar",
					}},
					Generation: blaxel.FunctionRuntimeGenerationMk2,
					Image:      blaxel.String("image"),
					MaxScale:   blaxel.Int(0),
					Memory:     blaxel.Int(0),
					MinScale:   blaxel.Int(0),
				},
				Transport: blaxel.FunctionSpecTransportWebsocket,
				Triggers: []blaxel.TriggerParam{{
					ID: blaxel.String("id"),
					Configuration: blaxel.TriggerConfigurationParam{
						AuthenticationType: blaxel.String("authenticationType"),
						CallbackSecret:     blaxel.String("callbackSecret"),
						CallbackURL:        blaxel.String("callbackUrl"),
						Path:               blaxel.String("path"),
						Retry:              blaxel.Int(0),
						Schedule:           blaxel.String("schedule"),
						Tasks: []map[string]any{{
							"foo": "bar",
						}},
						Timeout: blaxel.Int(0),
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
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := blaxel.NewClient(
		option.WithBaseURL(baseURL),
		option.WithClientID("My Client ID"),
		option.WithClientSecret("My Client Secret"),
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
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := blaxel.NewClient(
		option.WithBaseURL(baseURL),
		option.WithClientID("My Client ID"),
		option.WithClientSecret("My Client Secret"),
	)
	_, err := client.Functions.Update(
		context.TODO(),
		"functionName",
		blaxel.FunctionUpdateParams{
			Function: blaxel.FunctionParam{
				Metadata: blaxel.MetadataParam{
					Name:        "name",
					DisplayName: blaxel.String("displayName"),
					Labels: map[string]string{
						"foo": "string",
					},
				},
				Spec: blaxel.FunctionSpecParam{
					Enabled:                blaxel.Bool(true),
					IntegrationConnections: []string{"string"},
					Policies:               []string{"string"},
					Revision: blaxel.RevisionConfigurationParam{
						Active:           blaxel.String("active"),
						Canary:           blaxel.String("canary"),
						CanaryPercent:    blaxel.Int(0),
						StickySessionTtl: blaxel.Int(0),
						Traffic:          blaxel.Int(0),
					},
					Runtime: blaxel.FunctionRuntimeParam{
						Envs: []map[string]any{{
							"foo": "bar",
						}},
						Generation: blaxel.FunctionRuntimeGenerationMk2,
						Image:      blaxel.String("image"),
						MaxScale:   blaxel.Int(0),
						Memory:     blaxel.Int(0),
						MinScale:   blaxel.Int(0),
					},
					Transport: blaxel.FunctionSpecTransportWebsocket,
					Triggers: []blaxel.TriggerParam{{
						ID: blaxel.String("id"),
						Configuration: blaxel.TriggerConfigurationParam{
							AuthenticationType: blaxel.String("authenticationType"),
							CallbackSecret:     blaxel.String("callbackSecret"),
							CallbackURL:        blaxel.String("callbackUrl"),
							Path:               blaxel.String("path"),
							Retry:              blaxel.Int(0),
							Schedule:           blaxel.String("schedule"),
							Tasks: []map[string]any{{
								"foo": "bar",
							}},
							Timeout: blaxel.Int(0),
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
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := blaxel.NewClient(
		option.WithBaseURL(baseURL),
		option.WithClientID("My Client ID"),
		option.WithClientSecret("My Client Secret"),
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
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := blaxel.NewClient(
		option.WithBaseURL(baseURL),
		option.WithClientID("My Client ID"),
		option.WithClientSecret("My Client Secret"),
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
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := blaxel.NewClient(
		option.WithBaseURL(baseURL),
		option.WithClientID("My Client ID"),
		option.WithClientSecret("My Client Secret"),
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
