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

func TestJobExecutionNewWithOptionalParams(t *testing.T) {
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
	_, err := client.Jobs.Executions.New(
		context.TODO(),
		"jobId",
		blaxel.JobExecutionNewParams{
			CreateJobExecutionRequest: blaxel.CreateJobExecutionRequestParam{
				ID: blaxel.String("id"),
				Env: map[string]string{
					"0":  "{",
					"1":  "\"",
					"2":  "M",
					"3":  "Y",
					"4":  "_",
					"5":  "V",
					"6":  "A",
					"7":  "R",
					"8":  "\"",
					"9":  ":",
					"10": " ",
					"11": "\"",
					"12": "c",
					"13": "u",
					"14": "s",
					"15": "t",
					"16": "o",
					"17": "m",
					"18": "_",
					"19": "v",
					"20": "a",
					"21": "l",
					"22": "u",
					"23": "e",
					"24": "\"",
					"25": ",",
					"26": " ",
					"27": "\"",
					"28": "B",
					"29": "A",
					"30": "T",
					"31": "C",
					"32": "H",
					"33": "_",
					"34": "S",
					"35": "I",
					"36": "Z",
					"37": "E",
					"38": "\"",
					"39": ":",
					"40": " ",
					"41": "\"",
					"42": "1",
					"43": "0",
					"44": "0",
					"45": "\"",
					"46": "}",
				},
				ExecutionID: blaxel.String("executionId"),
				JobID:       blaxel.String("data-processing-job"),
				Memory:      blaxel.Int(2048),
				Tasks:       []any{map[string]any{}},
				WorkspaceID: blaxel.String("workspaceId"),
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

func TestJobExecutionGet(t *testing.T) {
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
	_, err := client.Jobs.Executions.Get(
		context.TODO(),
		"executionId",
		blaxel.JobExecutionGetParams{
			JobID: "jobId",
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

func TestJobExecutionListWithOptionalParams(t *testing.T) {
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
	_, err := client.Jobs.Executions.List(
		context.TODO(),
		"jobId",
		blaxel.JobExecutionListParams{
			Limit:  blaxel.Int(1),
			Offset: blaxel.Int(0),
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

func TestJobExecutionDelete(t *testing.T) {
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
	_, err := client.Jobs.Executions.Delete(
		context.TODO(),
		"executionId",
		blaxel.JobExecutionDeleteParams{
			JobID: "jobId",
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
