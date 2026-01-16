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

func TestJobNewWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := blaxel.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key", "Authorization"),
	)
	_, err := client.Jobs.New(context.TODO(), blaxel.JobNewParams{
		Job: blaxel.JobParam{
			Metadata: blaxel.MetadataParam{
				Name:        "my-resource",
				DisplayName: blaxel.String("My Resource"),
				Labels: map[string]string{
					"foo": "string",
				},
			},
			Spec: blaxel.JobSpecParam{
				Enabled:  blaxel.Bool(true),
				Policies: []string{"string"},
				Region:   blaxel.String("us-was-1"),
				Revision: blaxel.RevisionConfigurationParam{
					Active:           blaxel.String("rev-abc123"),
					Canary:           blaxel.String("canary"),
					CanaryPercent:    blaxel.Int(10),
					StickySessionTtl: blaxel.Int(0),
					Traffic:          blaxel.Int(100),
				},
				Runtime: blaxel.JobRuntimeParam{
					Envs: []blaxel.JobRuntimeEnvParam{{
						Name:   blaxel.String("MY_ENV_VAR"),
						Secret: blaxel.Bool(true),
						Value:  blaxel.String("my-value"),
					}},
					Generation:         blaxel.JobRuntimeGenerationMk3,
					Image:              blaxel.String("image"),
					MaxConcurrentTasks: blaxel.Int(10),
					MaxRetries:         blaxel.Int(3),
					Memory:             blaxel.Int(2048),
					Ports: []blaxel.PortParam{{
						Target:   8080,
						Name:     blaxel.String("http"),
						Protocol: blaxel.PortProtocolHTTP,
					}},
					Timeout: blaxel.Int(3600),
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

func TestJobGetWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := blaxel.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key", "Authorization"),
	)
	_, err := client.Jobs.Get(
		context.TODO(),
		"jobId",
		blaxel.JobGetParams{
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

func TestJobUpdateWithOptionalParams(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := blaxel.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key", "Authorization"),
	)
	_, err := client.Jobs.Update(
		context.TODO(),
		"jobId",
		blaxel.JobUpdateParams{
			Job: blaxel.JobParam{
				Metadata: blaxel.MetadataParam{
					Name:        "my-resource",
					DisplayName: blaxel.String("My Resource"),
					Labels: map[string]string{
						"foo": "string",
					},
				},
				Spec: blaxel.JobSpecParam{
					Enabled:  blaxel.Bool(true),
					Policies: []string{"string"},
					Region:   blaxel.String("us-was-1"),
					Revision: blaxel.RevisionConfigurationParam{
						Active:           blaxel.String("rev-abc123"),
						Canary:           blaxel.String("canary"),
						CanaryPercent:    blaxel.Int(10),
						StickySessionTtl: blaxel.Int(0),
						Traffic:          blaxel.Int(100),
					},
					Runtime: blaxel.JobRuntimeParam{
						Envs: []blaxel.JobRuntimeEnvParam{{
							Name:   blaxel.String("MY_ENV_VAR"),
							Secret: blaxel.Bool(true),
							Value:  blaxel.String("my-value"),
						}},
						Generation:         blaxel.JobRuntimeGenerationMk3,
						Image:              blaxel.String("image"),
						MaxConcurrentTasks: blaxel.Int(10),
						MaxRetries:         blaxel.Int(3),
						Memory:             blaxel.Int(2048),
						Ports: []blaxel.PortParam{{
							Target:   8080,
							Name:     blaxel.String("http"),
							Protocol: blaxel.PortProtocolHTTP,
						}},
						Timeout: blaxel.Int(3600),
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

func TestJobList(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := blaxel.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key", "Authorization"),
	)
	_, err := client.Jobs.List(context.TODO())
	if err != nil {
		var apierr *blaxel.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestJobDelete(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := blaxel.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key", "Authorization"),
	)
	_, err := client.Jobs.Delete(context.TODO(), "jobId")
	if err != nil {
		var apierr *blaxel.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestJobListRevisions(t *testing.T) {
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := blaxel.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key", "Authorization"),
	)
	_, err := client.Jobs.ListRevisions(context.TODO(), "jobId")
	if err != nil {
		var apierr *blaxel.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
