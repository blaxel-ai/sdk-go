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

func TestApplicationNewWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Applications.New(context.TODO(), blaxel.ApplicationNewParams{
		Application: blaxel.ApplicationParam{
			Metadata: blaxel.MetadataParam{
				Name:        "my-resource",
				DisplayName: blaxel.String("My Resource"),
				ExternalID:  blaxel.String("my-session-123"),
				Labels: map[string]string{
					"foo": "string",
				},
			},
			Spec: blaxel.ApplicationSpecParam{
				Enabled: blaxel.Bool(true),
				Envs: []shared.EnvParam{{
					Name:   blaxel.String("MY_ENV_VAR"),
					Secret: blaxel.Bool(true),
					Value:  blaxel.String("my-value"),
				}},
				Extensions: map[string]blaxel.ApplicationSpecExtensionParam{
					"foo": {
						Sandbox: blaxel.String("my-sandbox"),
					},
				},
				Image:  blaxel.String("my-registry/my-app:latest"),
				Memory: blaxel.Int(2048),
				Port:   blaxel.Int(8080),
				Proxy:  blaxel.Bool(false),
				Region: blaxel.String("us-pdx-1"),
				Revision: blaxel.AppRevisionConfigurationParam{
					Active:           blaxel.String("active"),
					Canary:           blaxel.String("canary"),
					CanaryPercent:    blaxel.Int(10),
					StickySessionTtl: blaxel.Int(0),
					Traffic:          blaxel.Int(100),
				},
				Revisions: []blaxel.AppRevisionParam{{
					Image:     "image",
					ID:        blaxel.String("id"),
					CreatedAt: blaxel.String("createdAt"),
					CreatedBy: blaxel.String("createdBy"),
					Envs: []shared.EnvParam{{
						Name:   blaxel.String("MY_ENV_VAR"),
						Secret: blaxel.Bool(true),
						Value:  blaxel.String("my-value"),
					}},
					Memory:   blaxel.Int(2048),
					Port:     blaxel.Int(8080),
					Snapshot: blaxel.String("snapshot"),
				}},
				URLs: []blaxel.AppURLParam{{
					Domain:    "app.example.com",
					Subdomain: blaxel.String("www"),
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

func TestApplicationGet(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Applications.Get(context.TODO(), "applicationName")
	if err != nil {
		var apierr *blaxel.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestApplicationUpdateWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Applications.Update(
		context.TODO(),
		"applicationName",
		blaxel.ApplicationUpdateParams{
			Application: blaxel.ApplicationParam{
				Metadata: blaxel.MetadataParam{
					Name:        "my-resource",
					DisplayName: blaxel.String("My Resource"),
					ExternalID:  blaxel.String("my-session-123"),
					Labels: map[string]string{
						"foo": "string",
					},
				},
				Spec: blaxel.ApplicationSpecParam{
					Enabled: blaxel.Bool(true),
					Envs: []shared.EnvParam{{
						Name:   blaxel.String("MY_ENV_VAR"),
						Secret: blaxel.Bool(true),
						Value:  blaxel.String("my-value"),
					}},
					Extensions: map[string]blaxel.ApplicationSpecExtensionParam{
						"foo": {
							Sandbox: blaxel.String("my-sandbox"),
						},
					},
					Image:  blaxel.String("my-registry/my-app:latest"),
					Memory: blaxel.Int(2048),
					Port:   blaxel.Int(8080),
					Proxy:  blaxel.Bool(false),
					Region: blaxel.String("us-pdx-1"),
					Revision: blaxel.AppRevisionConfigurationParam{
						Active:           blaxel.String("active"),
						Canary:           blaxel.String("canary"),
						CanaryPercent:    blaxel.Int(10),
						StickySessionTtl: blaxel.Int(0),
						Traffic:          blaxel.Int(100),
					},
					Revisions: []blaxel.AppRevisionParam{{
						Image:     "image",
						ID:        blaxel.String("id"),
						CreatedAt: blaxel.String("createdAt"),
						CreatedBy: blaxel.String("createdBy"),
						Envs: []shared.EnvParam{{
							Name:   blaxel.String("MY_ENV_VAR"),
							Secret: blaxel.Bool(true),
							Value:  blaxel.String("my-value"),
						}},
						Memory:   blaxel.Int(2048),
						Port:     blaxel.Int(8080),
						Snapshot: blaxel.String("snapshot"),
					}},
					URLs: []blaxel.AppURLParam{{
						Domain:    "app.example.com",
						Subdomain: blaxel.String("www"),
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

func TestApplicationListWithOptionalParams(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Applications.List(context.TODO(), blaxel.ApplicationListParams{
		Anchor: blaxel.ApplicationListParamsAnchorEnd,
		Cursor: blaxel.String("cursor"),
		Limit:  blaxel.Int(1),
		Q:      blaxel.String("q"),
		Sort:   blaxel.ApplicationListParamsSortCreatedAtDesc,
	})
	if err != nil {
		var apierr *blaxel.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestApplicationDelete(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Applications.Delete(context.TODO(), "applicationName")
	if err != nil {
		var apierr *blaxel.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestApplicationListRevisions(t *testing.T) {
	t.Skip("Mock server tests are disabled")
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
	_, err := client.Applications.ListRevisions(context.TODO(), "applicationName")
	if err != nil {
		var apierr *blaxel.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
