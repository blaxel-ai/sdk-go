package integration_tests

import (
	"context"
	"testing"

	blaxel "github.com/blaxel-ai/sdk-go"
)

// TEST_JOB_NAME is the name of an existing job in the workspace
// Note: These tests require a job named "mk3" to exist in your workspace.
const TEST_JOB_NAME = "mk3"

func TestJobs(t *testing.T) {
	ctx := context.Background()
	client := newTestClient(t)

	// Check if job exists
	_, err := client.Jobs.Get(ctx, TEST_JOB_NAME, blaxel.JobGetParams{})
	if err != nil {
		t.Skipf("Job %q does not exist. Skipping Jobs API Integration tests.", TEST_JOB_NAME)
	}

	t.Run("GetJob", func(t *testing.T) {
		t.Run("retrieves an existing job", func(t *testing.T) {
			job, err := client.Jobs.Get(ctx, TEST_JOB_NAME, blaxel.JobGetParams{})
			if err != nil {
				t.Fatalf("failed to get job: %v", err)
			}

			if job.Metadata.Name != TEST_JOB_NAME {
				t.Errorf("expected name %s, got %s", TEST_JOB_NAME, job.Metadata.Name)
			}
		})
	})

	t.Run("ListJobs", func(t *testing.T) {
		t.Run("lists all jobs", func(t *testing.T) {
			jobs, err := client.Jobs.List(ctx)
			if err != nil {
				t.Fatalf("failed to list jobs: %v", err)
			}

			if jobs == nil || len(*jobs) == 0 {
				t.Error("expected at least one job")
			}

			found := false
			for _, job := range *jobs {
				if job.Metadata.Name == TEST_JOB_NAME {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("expected to find job %s in list", TEST_JOB_NAME)
			}
		})
	})

	t.Run("JobExecutions", func(t *testing.T) {
		t.Run("creates and lists job executions", func(t *testing.T) {
			// Create an execution
			execution, err := client.Jobs.Executions.New(ctx, TEST_JOB_NAME, blaxel.JobExecutionNewParams{
				CreateJobExecutionRequest: blaxel.CreateJobExecutionRequestParam{
					Tasks: []any{
						map[string]interface{}{"name": "Richard"},
						map[string]interface{}{"name": "John"},
					},
				},
			})
			if err != nil {
				t.Fatalf("failed to create execution: %v", err)
			}

			if execution.ExecutionID == "" {
				t.Error("expected execution ID to be set")
			}

			// List executions
			executions, err := client.Jobs.Executions.List(ctx, TEST_JOB_NAME, blaxel.JobExecutionListParams{})
			if err != nil {
				t.Fatalf("failed to list executions: %v", err)
			}

			if executions == nil || len(*executions) == 0 {
				t.Error("expected at least one execution")
			}
		})

		t.Run("gets a job execution", func(t *testing.T) {
			// First create an execution
			execution, err := client.Jobs.Executions.New(ctx, TEST_JOB_NAME, blaxel.JobExecutionNewParams{
				CreateJobExecutionRequest: blaxel.CreateJobExecutionRequestParam{
					Tasks: []any{
						map[string]interface{}{"name": "Test"},
					},
				},
			})
			if err != nil {
				t.Fatalf("failed to create execution: %v", err)
			}

			// Get the execution
			retrieved, err := client.Jobs.Executions.Get(ctx, execution.ExecutionID, blaxel.JobExecutionGetParams{
				JobID: TEST_JOB_NAME,
			})
			if err != nil {
				t.Fatalf("failed to get execution: %v", err)
			}

			if retrieved.Metadata.ID != execution.ExecutionID {
				t.Errorf("expected ID %s, got %s", execution.ExecutionID, retrieved.Metadata.ID)
			}
		})
	})
}
