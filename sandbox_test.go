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

func TestSandboxNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Sandboxes.New(context.TODO(), blaxel.SandboxNewParams{
		Sandbox: blaxel.SandboxParam{
			Metadata: blaxel.MetadataParam{
				Name:        "my-resource",
				DisplayName: blaxel.String("My Resource"),
				Labels: map[string]string{
					"foo": "string",
				},
			},
			Spec: blaxel.SandboxSpecParam{
				Enabled: blaxel.Bool(true),
				Lifecycle: blaxel.SandboxLifecycleParam{
					ExpirationPolicies: []blaxel.ExpirationPolicyParam{{
						Action: blaxel.ExpirationPolicyActionDelete,
						Type:   blaxel.ExpirationPolicyTypeTtlIdle,
						Value:  blaxel.String("24h"),
					}},
				},
				Region: blaxel.String("us-pdx-1"),
				Runtime: blaxel.SandboxRuntimeParam{
					Envs: []blaxel.SandboxRuntimeEnvParam{{
						Name:   blaxel.String("MY_ENV_VAR"),
						Secret: blaxel.Bool(true),
						Value:  blaxel.String("my-value"),
					}},
					Expires: blaxel.String("2025-12-31T23:59:59Z"),
					Image:   blaxel.String("blaxel/base-image:latest"),
					Memory:  blaxel.Int(4096),
					Ports: []blaxel.PortParam{{
						Target:   8080,
						Name:     blaxel.String("http"),
						Protocol: blaxel.PortProtocolHTTP,
					}},
					Ttl: blaxel.String("24h"),
				},
				Volumes: []blaxel.VolumeAttachmentParam{{
					MountPath: blaxel.String("/mnt/data"),
					Name:      blaxel.String("my-volume"),
					ReadOnly:  blaxel.Bool(false),
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

func TestSandboxGetWithOptionalParams(t *testing.T) {
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
	_, err := client.Sandboxes.Get(
		context.TODO(),
		"sandboxName",
		blaxel.SandboxGetParams{
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

func TestSandboxUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Sandboxes.Update(
		context.TODO(),
		"sandboxName",
		blaxel.SandboxUpdateParams{
			Sandbox: blaxel.SandboxParam{
				Metadata: blaxel.MetadataParam{
					Name:        "my-resource",
					DisplayName: blaxel.String("My Resource"),
					Labels: map[string]string{
						"foo": "string",
					},
				},
				Spec: blaxel.SandboxSpecParam{
					Enabled: blaxel.Bool(true),
					Lifecycle: blaxel.SandboxLifecycleParam{
						ExpirationPolicies: []blaxel.ExpirationPolicyParam{{
							Action: blaxel.ExpirationPolicyActionDelete,
							Type:   blaxel.ExpirationPolicyTypeTtlIdle,
							Value:  blaxel.String("24h"),
						}},
					},
					Region: blaxel.String("us-pdx-1"),
					Runtime: blaxel.SandboxRuntimeParam{
						Envs: []blaxel.SandboxRuntimeEnvParam{{
							Name:   blaxel.String("MY_ENV_VAR"),
							Secret: blaxel.Bool(true),
							Value:  blaxel.String("my-value"),
						}},
						Expires: blaxel.String("2025-12-31T23:59:59Z"),
						Image:   blaxel.String("blaxel/base-image:latest"),
						Memory:  blaxel.Int(4096),
						Ports: []blaxel.PortParam{{
							Target:   8080,
							Name:     blaxel.String("http"),
							Protocol: blaxel.PortProtocolHTTP,
						}},
						Ttl: blaxel.String("24h"),
					},
					Volumes: []blaxel.VolumeAttachmentParam{{
						MountPath: blaxel.String("/mnt/data"),
						Name:      blaxel.String("my-volume"),
						ReadOnly:  blaxel.Bool(false),
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

func TestSandboxList(t *testing.T) {
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
	_, err := client.Sandboxes.List(context.TODO())
	if err != nil {
		var apierr *blaxel.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSandboxDelete(t *testing.T) {
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
	_, err := client.Sandboxes.Delete(context.TODO(), "sandboxName")
	if err != nil {
		var apierr *blaxel.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSandboxGetHub(t *testing.T) {
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
	_, err := client.Sandboxes.GetHub(context.TODO())
	if err != nil {
		var apierr *blaxel.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
