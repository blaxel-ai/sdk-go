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

func TestAgentNewWithOptionalParams(t *testing.T) {
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
	)
	_, err := client.Agents.New(context.TODO(), blaxel.AgentNewParams{
		Agent: blaxel.AgentParam{
			Metadata: blaxel.MetadataParam{
				TimeFieldsParam:  blaxel.TimeFieldsParam{},
				OwnerFieldsParam: blaxel.OwnerFieldsParam{},
				Name:             "name",
				DisplayName:      blaxel.String("displayName"),
				Labels: map[string]string{
					"foo": "string",
				},
			},
			Spec: blaxel.AgentSpecParam{
				Enabled:  blaxel.Bool(true),
				Policies: []string{"string"},
				Repository: blaxel.AgentSpecRepositoryParam{
					Type: blaxel.String("type"),
					URL:  blaxel.String("url"),
				},
				Revision: blaxel.RevisionConfigurationParam{
					Active:           blaxel.String("active"),
					Canary:           blaxel.String("canary"),
					CanaryPercent:    blaxel.Int(0),
					StickySessionTtl: blaxel.Int(0),
					Traffic:          blaxel.Int(0),
				},
				Runtime: blaxel.AgentSpecRuntimeParam{
					Envs: []map[string]any{{
						"foo": "bar",
					}},
					Generation: "mk2",
					Image:      blaxel.String("image"),
					MaxScale:   blaxel.Int(0),
					Memory:     blaxel.Int(0),
					MinScale:   blaxel.Int(0),
				},
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

func TestAgentGetWithOptionalParams(t *testing.T) {
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
	)
	_, err := client.Agents.Update(
		context.TODO(),
		"agentName",
		blaxel.AgentUpdateParams{
			Agent: blaxel.AgentParam{
				Metadata: blaxel.MetadataParam{
					TimeFieldsParam:  blaxel.TimeFieldsParam{},
					OwnerFieldsParam: blaxel.OwnerFieldsParam{},
					Name:             "name",
					DisplayName:      blaxel.String("displayName"),
					Labels: map[string]string{
						"foo": "string",
					},
				},
				Spec: blaxel.AgentSpecParam{
					Enabled:  blaxel.Bool(true),
					Policies: []string{"string"},
					Repository: blaxel.AgentSpecRepositoryParam{
						Type: blaxel.String("type"),
						URL:  blaxel.String("url"),
					},
					Revision: blaxel.RevisionConfigurationParam{
						Active:           blaxel.String("active"),
						Canary:           blaxel.String("canary"),
						CanaryPercent:    blaxel.Int(0),
						StickySessionTtl: blaxel.Int(0),
						Traffic:          blaxel.Int(0),
					},
					Runtime: blaxel.AgentSpecRuntimeParam{
						Envs: []map[string]any{{
							"foo": "bar",
						}},
						Generation: "mk2",
						Image:      blaxel.String("image"),
						MaxScale:   blaxel.Int(0),
						Memory:     blaxel.Int(0),
						MinScale:   blaxel.Int(0),
					},
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

func TestAgentList(t *testing.T) {
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
