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

func TestIntegrationConnectionNew(t *testing.T) {
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
	_, err := client.Integrations.Connections.New(context.TODO(), blaxel.IntegrationConnectionNewParams{
		IntegrationConnection: blaxel.IntegrationConnectionParam{
			Metadata: blaxel.MetadataParam{
				TimeFieldsParam:  blaxel.TimeFieldsParam{},
				OwnerFieldsParam: blaxel.OwnerFieldsParam{},
				Name:             "name",
				DisplayName:      blaxel.String("displayName"),
				Labels: map[string]string{
					"foo": "string",
				},
			},
			Spec: blaxel.IntegrationConnectionSpecParam{
				Config: map[string]string{
					"foo": "string",
				},
				Integration: blaxel.String("integration"),
				Sandbox:     blaxel.Bool(true),
				Secret: map[string]string{
					"foo": "string",
				},
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

func TestIntegrationConnectionGet(t *testing.T) {
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
	_, err := client.Integrations.Connections.Get(context.TODO(), "connectionName")
	if err != nil {
		var apierr *blaxel.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestIntegrationConnectionUpdate(t *testing.T) {
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
	_, err := client.Integrations.Connections.Update(
		context.TODO(),
		"connectionName",
		blaxel.IntegrationConnectionUpdateParams{
			IntegrationConnection: blaxel.IntegrationConnectionParam{
				Metadata: blaxel.MetadataParam{
					TimeFieldsParam:  blaxel.TimeFieldsParam{},
					OwnerFieldsParam: blaxel.OwnerFieldsParam{},
					Name:             "name",
					DisplayName:      blaxel.String("displayName"),
					Labels: map[string]string{
						"foo": "string",
					},
				},
				Spec: blaxel.IntegrationConnectionSpecParam{
					Config: map[string]string{
						"foo": "string",
					},
					Integration: blaxel.String("integration"),
					Sandbox:     blaxel.Bool(true),
					Secret: map[string]string{
						"foo": "string",
					},
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

func TestIntegrationConnectionList(t *testing.T) {
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
	_, err := client.Integrations.Connections.List(context.TODO())
	if err != nil {
		var apierr *blaxel.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestIntegrationConnectionDelete(t *testing.T) {
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
	_, err := client.Integrations.Connections.Delete(context.TODO(), "connectionName")
	if err != nil {
		var apierr *blaxel.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestIntegrationConnectionListEndpointConfigurations(t *testing.T) {
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
	err := client.Integrations.Connections.ListEndpointConfigurations(context.TODO(), "connectionName")
	if err != nil {
		var apierr *blaxel.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
