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

func TestAgentNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Agents.New(context.TODO(), blaxel.AgentNewParams{
		Agent: blaxel.AgentParam{
			Metadata: blaxel.MetadataParam{
				Name:        "my-resource",
				DisplayName: blaxel.String("My Resource"),
				Labels: map[string]string{
					"foo": "string",
				},
			},
			Spec: blaxel.AgentSpecParam{
				Enabled:  blaxel.Bool(true),
				Policies: []string{"string"},
				Repository: blaxel.RepositoryParam{
					Type: blaxel.String("github"),
					URL:  blaxel.String("https://github.com/my-org/my-agent"),
				},
				Revision: blaxel.RevisionConfigurationParam{
					Active:           blaxel.String("rev-abc123"),
					Canary:           blaxel.String("canary"),
					CanaryPercent:    blaxel.Int(10),
					StickySessionTtl: blaxel.Int(0),
					Traffic:          blaxel.Int(100),
				},
				Runtime: blaxel.AgentRuntimeParam{
					Envs: []blaxel.AgentRuntimeEnvParam{{
						Name:   blaxel.String("MY_ENV_VAR"),
						Secret: blaxel.Bool(true),
						Value:  blaxel.String("my-value"),
					}},
					Generation: blaxel.AgentRuntimeGenerationMk3,
					Image:      blaxel.String("image"),
					MaxScale:   blaxel.Int(10),
					Memory:     blaxel.Int(2048),
					MinScale:   blaxel.Int(0),
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

func TestAgentGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Agents.Get(
		context.TODO(),
		"agentName",
		blaxel.AgentGetParams{
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

func TestAgentUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Agents.Update(
		context.TODO(),
		"agentName",
		blaxel.AgentUpdateParams{
			Agent: blaxel.AgentParam{
				Metadata: blaxel.MetadataParam{
					Name:        "my-resource",
					DisplayName: blaxel.String("My Resource"),
					Labels: map[string]string{
						"foo": "string",
					},
				},
				Spec: blaxel.AgentSpecParam{
					Enabled:  blaxel.Bool(true),
					Policies: []string{"string"},
					Repository: blaxel.RepositoryParam{
						Type: blaxel.String("github"),
						URL:  blaxel.String("https://github.com/my-org/my-agent"),
					},
					Revision: blaxel.RevisionConfigurationParam{
						Active:           blaxel.String("rev-abc123"),
						Canary:           blaxel.String("canary"),
						CanaryPercent:    blaxel.Int(10),
						StickySessionTtl: blaxel.Int(0),
						Traffic:          blaxel.Int(100),
					},
					Runtime: blaxel.AgentRuntimeParam{
						Envs: []blaxel.AgentRuntimeEnvParam{{
							Name:   blaxel.String("MY_ENV_VAR"),
							Secret: blaxel.Bool(true),
							Value:  blaxel.String("my-value"),
						}},
						Generation: blaxel.AgentRuntimeGenerationMk3,
						Image:      blaxel.String("image"),
						MaxScale:   blaxel.Int(10),
						Memory:     blaxel.Int(2048),
						MinScale:   blaxel.Int(0),
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

func TestAgentList(t *testing.T) {
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
	_, err := client.Agents.List(context.TODO())
	if err != nil {
		var apierr *blaxel.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAgentDelete(t *testing.T) {
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
	_, err := client.Agents.Delete(context.TODO(), "agentName")
	if err != nil {
		var apierr *blaxel.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestAgentListRevisions(t *testing.T) {
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
	_, err := client.Agents.ListRevisions(context.TODO(), "agentName")
	if err != nil {
		var apierr *blaxel.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
