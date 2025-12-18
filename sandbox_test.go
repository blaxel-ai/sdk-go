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
				Name:        "name",
				DisplayName: blaxel.String("displayName"),
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
						Value:  blaxel.String("value"),
					}},
				},
				Region: blaxel.String("region"),
				Runtime: blaxel.SandboxRuntimeParam{
					Envs: []map[string]any{{
						"foo": "bar",
					}},
					Expires: blaxel.String("expires"),
					Image:   blaxel.String("image"),
					Memory:  blaxel.Int(0),
					Ports: []blaxel.PortParam{{
						Name:     blaxel.String("name"),
						Protocol: blaxel.PortProtocolHTTP,
						Target:   blaxel.Int(0),
					}},
					Ttl: blaxel.String("ttl"),
				},
				Volumes: []blaxel.VolumeAttachmentParam{{
					MountPath: blaxel.String("mountPath"),
					Name:      blaxel.String("name"),
					ReadOnly:  blaxel.Bool(true),
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
					Name:        "name",
					DisplayName: blaxel.String("displayName"),
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
							Value:  blaxel.String("value"),
						}},
					},
					Region: blaxel.String("region"),
					Runtime: blaxel.SandboxRuntimeParam{
						Envs: []map[string]any{{
							"foo": "bar",
						}},
						Expires: blaxel.String("expires"),
						Image:   blaxel.String("image"),
						Memory:  blaxel.Int(0),
						Ports: []blaxel.PortParam{{
							Name:     blaxel.String("name"),
							Protocol: blaxel.PortProtocolHTTP,
							Target:   blaxel.Int(0),
						}},
						Ttl: blaxel.String("ttl"),
					},
					Volumes: []blaxel.VolumeAttachmentParam{{
						MountPath: blaxel.String("mountPath"),
						Name:      blaxel.String("name"),
						ReadOnly:  blaxel.Bool(true),
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
