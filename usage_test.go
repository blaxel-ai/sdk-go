// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blaxel_test

import (
	"context"
	"os"
	"testing"

	"github.com/blaxel-ai/sdk-go"
	"github.com/blaxel-ai/sdk-go/internal/testutil"
	"github.com/blaxel-ai/sdk-go/option"
)

func TestUsage(t *testing.T) {
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
	sandbox, err := client.Sandboxes.New(context.TODO(), blaxel.SandboxNewParams{
		Sandbox: blaxel.SandboxParam{
			Metadata: blaxel.MetadataParam{
				Name: "my-resource",
			},
			Spec: blaxel.SandboxSpecParam{},
		},
	})
	if err != nil {
		t.Fatalf("err should be nil: %s", err.Error())
	}
	t.Logf("%+v\n", sandbox.Metadata)
}
