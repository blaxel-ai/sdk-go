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

func TestJobNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Jobs.New(context.TODO(), blaxel.JobNewParams{
		Job: blaxel.JobParam{
			Metadata: blaxel.MetadataParam{
				Name:        "name",
				DisplayName: blaxel.String("displayName"),
				Labels: map[string]string{
					"foo": "string",
				},
			},
			Spec: blaxel.JobSpecParam{
				Enabled:  blaxel.Bool(true),
				Policies: []string{"string"},
				Region:   blaxel.String("region"),
				Revision: blaxel.RevisionConfigurationParam{
					Active:           blaxel.String("active"),
					Canary:           blaxel.String("canary"),
					CanaryPercent:    blaxel.Int(0),
					StickySessionTtl: blaxel.Int(0),
					Traffic:          blaxel.Int(0),
				},
				Runtime: blaxel.JobRuntimeParam{
					Envs: []map[string]any{{
						"foo": "bar",
					}},
					Generation:         blaxel.JobRuntimeGenerationMk2,
					Image:              blaxel.String("image"),
					MaxConcurrentTasks: blaxel.Int(0),
					MaxRetries:         blaxel.Int(0),
					Memory:             blaxel.Int(0),
					Ports: []blaxel.PortParam{{
						Name:     blaxel.String("name"),
						Protocol: blaxel.PortProtocolHTTP,
						Target:   blaxel.Int(0),
					}},
					Timeout: blaxel.Int(0),
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

func TestJobGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Jobs.Update(
		context.TODO(),
		"jobId",
		blaxel.JobUpdateParams{
			Job: blaxel.JobParam{
				Metadata: blaxel.MetadataParam{
					Name:        "name",
					DisplayName: blaxel.String("displayName"),
					Labels: map[string]string{
						"foo": "string",
					},
				},
				Spec: blaxel.JobSpecParam{
					Enabled:  blaxel.Bool(true),
					Policies: []string{"string"},
					Region:   blaxel.String("region"),
					Revision: blaxel.RevisionConfigurationParam{
						Active:           blaxel.String("active"),
						Canary:           blaxel.String("canary"),
						CanaryPercent:    blaxel.Int(0),
						StickySessionTtl: blaxel.Int(0),
						Traffic:          blaxel.Int(0),
					},
					Runtime: blaxel.JobRuntimeParam{
						Envs: []map[string]any{{
							"foo": "bar",
						}},
						Generation:         blaxel.JobRuntimeGenerationMk2,
						Image:              blaxel.String("image"),
						MaxConcurrentTasks: blaxel.Int(0),
						MaxRetries:         blaxel.Int(0),
						Memory:             blaxel.Int(0),
						Ports: []blaxel.PortParam{{
							Name:     blaxel.String("name"),
							Protocol: blaxel.PortProtocolHTTP,
							Target:   blaxel.Int(0),
						}},
						Timeout: blaxel.Int(0),
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

func TestJobList(t *testing.T) {
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
	_, err := client.Jobs.ListRevisions(context.TODO(), "jobId")
	if err != nil {
		var apierr *blaxel.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
