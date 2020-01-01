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

func TestSandboxScheduleNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Sandboxes.Schedules.New(
		context.TODO(),
		"sandboxName",
		blaxel.SandboxScheduleNewParams{
			SandboxScheduleEntry: blaxel.SandboxScheduleEntryParam{
				ID: blaxel.String("schedule-0"),
				Input: blaxel.SandboxScheduleInputParam{
					Command: blaxel.String("python train.py --epochs 10"),
					Env: map[string]string{
						"foo": "string",
					},
					KeepAlive:  blaxel.Bool(true),
					Name:       blaxel.String("training-job"),
					Timeout:    blaxel.Int(3600),
					WorkingDir: blaxel.String("/app"),
				},
				MaxExecutions: blaxel.Int(100),
				Type:          blaxel.SandboxScheduleEntryTypeCron,
				Value:         blaxel.String("0 8 * * 1-5"),
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

func TestSandboxScheduleGet(t *testing.T) {
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
	_, err := client.Sandboxes.Schedules.Get(
		context.TODO(),
		"scheduleId",
		blaxel.SandboxScheduleGetParams{
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

func TestSandboxScheduleUpdateWithOptionalParams(t *testing.T) {
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
	_, err := client.Sandboxes.Schedules.Update(
		context.TODO(),
		"scheduleId",
		blaxel.SandboxScheduleUpdateParams{
			SandboxName: "sandboxName",
			SandboxScheduleEntry: blaxel.SandboxScheduleEntryParam{
				ID: blaxel.String("schedule-0"),
				Input: blaxel.SandboxScheduleInputParam{
					Command: blaxel.String("python train.py --epochs 10"),
					Env: map[string]string{
						"foo": "string",
					},
					KeepAlive:  blaxel.Bool(true),
					Name:       blaxel.String("training-job"),
					Timeout:    blaxel.Int(3600),
					WorkingDir: blaxel.String("/app"),
				},
				MaxExecutions: blaxel.Int(100),
				Type:          blaxel.SandboxScheduleEntryTypeCron,
				Value:         blaxel.String("0 8 * * 1-5"),
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

func TestSandboxScheduleListWithOptionalParams(t *testing.T) {
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
	_, err := client.Sandboxes.Schedules.List(
		context.TODO(),
		"sandboxName",
		blaxel.SandboxScheduleListParams{
			Cursor: blaxel.String("cursor"),
			Limit:  blaxel.Int(1),
			Q:      blaxel.String("q"),
			Sort:   blaxel.SandboxScheduleListParamsSortCreatedAtDesc,
			Type:   blaxel.SandboxScheduleListParamsTypeCron,
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

func TestSandboxScheduleDelete(t *testing.T) {
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
	_, err := client.Sandboxes.Schedules.Delete(
		context.TODO(),
		"scheduleId",
		blaxel.SandboxScheduleDeleteParams{
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

func TestSandboxScheduleListExecutionsWithOptionalParams(t *testing.T) {
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
	_, err := client.Sandboxes.Schedules.ListExecutions(
		context.TODO(),
		"sandboxName",
		blaxel.SandboxScheduleListExecutionsParams{
			Cursor: blaxel.String("cursor"),
			Limit:  blaxel.Int(1),
			Q:      blaxel.String("q"),
			Sort:   blaxel.SandboxScheduleListExecutionsParamsSortCreatedAtDesc,
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
