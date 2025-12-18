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

func TestPolicyNew(t *testing.T) {
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
	_, err := client.Policies.New(context.TODO(), blaxel.PolicyNewParams{
		Policy: blaxel.PolicyParam{
			Metadata: blaxel.MetadataParam{
				TimeFieldsParam:  blaxel.TimeFieldsParam{},
				OwnerFieldsParam: blaxel.OwnerFieldsParam{},
				Name:             "name",
				DisplayName:      blaxel.String("displayName"),
				Labels: map[string]string{
					"foo": "string",
				},
			},
			Spec: blaxel.PolicySpecParam{
				Flavors: []blaxel.FlavorParam{{
					Name: blaxel.String("name"),
					Type: blaxel.FlavorTypeCPU,
				}},
				Locations: []blaxel.PolicySpecLocationParam{{
					Name: blaxel.String("name"),
					Type: "location",
				}},
				MaxTokens: blaxel.PolicySpecMaxTokensParam{
					Granularity:          blaxel.String("granularity"),
					Input:                blaxel.Int(0),
					Output:               blaxel.Int(0),
					RatioInputOverOutput: blaxel.Int(0),
					Step:                 blaxel.Int(0),
					Total:                blaxel.Int(0),
				},
				ResourceTypes: []string{"model"},
				Sandbox:       blaxel.Bool(true),
				Type:          "location",
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

func TestPolicyGet(t *testing.T) {
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
	_, err := client.Policies.Get(context.TODO(), "policyName")
	if err != nil {
		var apierr *blaxel.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPolicyUpdate(t *testing.T) {
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
	_, err := client.Policies.Update(
		context.TODO(),
		"policyName",
		blaxel.PolicyUpdateParams{
			Policy: blaxel.PolicyParam{
				Metadata: blaxel.MetadataParam{
					TimeFieldsParam:  blaxel.TimeFieldsParam{},
					OwnerFieldsParam: blaxel.OwnerFieldsParam{},
					Name:             "name",
					DisplayName:      blaxel.String("displayName"),
					Labels: map[string]string{
						"foo": "string",
					},
				},
				Spec: blaxel.PolicySpecParam{
					Flavors: []blaxel.FlavorParam{{
						Name: blaxel.String("name"),
						Type: blaxel.FlavorTypeCPU,
					}},
					Locations: []blaxel.PolicySpecLocationParam{{
						Name: blaxel.String("name"),
						Type: "location",
					}},
					MaxTokens: blaxel.PolicySpecMaxTokensParam{
						Granularity:          blaxel.String("granularity"),
						Input:                blaxel.Int(0),
						Output:               blaxel.Int(0),
						RatioInputOverOutput: blaxel.Int(0),
						Step:                 blaxel.Int(0),
						Total:                blaxel.Int(0),
					},
					ResourceTypes: []string{"model"},
					Sandbox:       blaxel.Bool(true),
					Type:          "location",
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

func TestPolicyList(t *testing.T) {
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
	_, err := client.Policies.List(context.TODO())
	if err != nil {
		var apierr *blaxel.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestPolicyDelete(t *testing.T) {
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
	_, err := client.Policies.Delete(context.TODO(), "policyName")
	if err != nil {
		var apierr *blaxel.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
