// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package blaxel_test

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/blaxel-ai/sdk-go"
	"github.com/blaxel-ai/sdk-go/internal/testutil"
	"github.com/blaxel-ai/sdk-go/option"
)

func TestScheduleExecutionListWithOptionalParams(t *testing.T) {
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
	_, err := client.ScheduleExecutions.List(context.TODO(), blaxel.ScheduleExecutionListParams{
		Cursor:   blaxel.String("cursor"),
		Limit:    blaxel.Int(1),
		Q:        blaxel.String("q"),
		Sandbox:  blaxel.String("sandbox"),
		Schedule: blaxel.String("schedule"),
		Since:    blaxel.Time(time.Now()),
		Sort:     blaxel.ScheduleExecutionListParamsSortCreatedAtDesc,
		Status:   blaxel.ScheduleExecutionListParamsStatusSucceeded,
		Until:    blaxel.Time(time.Now()),
	})
	if err != nil {
		var apierr *blaxel.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
