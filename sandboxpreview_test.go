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

func TestSandboxPreviewNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Sandboxes.Previews.New(
		context.TODO(),
		"sandboxName",
		blaxel.SandboxPreviewNewParams{
			Preview: blaxel.PreviewParam{
				Metadata: blaxel.PreviewMetadataParam{
					Name:         "name",
					DisplayName:  blaxel.String("displayName"),
					ResourceName: blaxel.String("resourceName"),
					ResourceType: blaxel.String("resourceType"),
					Workspace:    blaxel.String("workspace"),
				},
				Spec: blaxel.PreviewSpecParam{
					CustomDomain: blaxel.String("customDomain"),
					Expires:      blaxel.String("expires"),
					Port:         blaxel.Int(0),
					PrefixURL:    blaxel.String("prefixUrl"),
					Public:       blaxel.Bool(true),
					RequestHeaders: map[string]string{
						"foo": "string",
					},
					ResponseHeaders: map[string]string{
						"foo": "string",
					},
					Ttl: blaxel.String("ttl"),
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

func TestSandboxPreviewGet(t *testing.T) {
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
	_, err := client.Sandboxes.Previews.Get(
		context.TODO(),
		"previewName",
		blaxel.SandboxPreviewGetParams{
			SandboxName: "sandboxName",
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

func TestSandboxPreviewUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Sandboxes.Previews.Update(
		context.TODO(),
		"previewName",
		blaxel.SandboxPreviewUpdateParams{
			SandboxName: "sandboxName",
			Preview: blaxel.PreviewParam{
				Metadata: blaxel.PreviewMetadataParam{
					Name:         "name",
					DisplayName:  blaxel.String("displayName"),
					ResourceName: blaxel.String("resourceName"),
					ResourceType: blaxel.String("resourceType"),
					Workspace:    blaxel.String("workspace"),
				},
				Spec: blaxel.PreviewSpecParam{
					CustomDomain: blaxel.String("customDomain"),
					Expires:      blaxel.String("expires"),
					Port:         blaxel.Int(0),
					PrefixURL:    blaxel.String("prefixUrl"),
					Public:       blaxel.Bool(true),
					RequestHeaders: map[string]string{
						"foo": "string",
					},
					ResponseHeaders: map[string]string{
						"foo": "string",
					},
					Ttl: blaxel.String("ttl"),
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

func TestSandboxPreviewList(t *testing.T) {
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
	_, err := client.Sandboxes.Previews.List(context.TODO(), "sandboxName")
	if err != nil {
		var apierr *blaxel.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestSandboxPreviewDelete(t *testing.T) {
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
	_, err := client.Sandboxes.Previews.Delete(
		context.TODO(),
		"previewName",
		blaxel.SandboxPreviewDeleteParams{
			SandboxName: "sandboxName",
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
