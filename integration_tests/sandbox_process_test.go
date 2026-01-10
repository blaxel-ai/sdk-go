package integration_tests

import (
	"context"
	"strings"
	"sync"
	"testing"
	"time"

	blaxel "github.com/stainless-sdks/blaxel-go"
)

func TestSandboxProcess(t *testing.T) {
	ctx := context.Background()
	client := newTestClient(t)

	sandboxName := uniqueName("process-test")
	sandbox, err := client.Sandboxes.NewInstance(ctx, blaxel.SandboxNewParams{
		Sandbox: blaxel.SandboxParam{
			Metadata: blaxel.MetadataParam{
				Name:   sandboxName,
				Labels: defaultLabels,
			},
			Spec: blaxel.SandboxSpecParam{
				Runtime: blaxel.SandboxRuntimeParam{
					Image:  blaxel.String(defaultImage),
					Memory: blaxel.Int(2048),
				},
			},
		},
	})
	if err != nil {
		t.Fatalf("failed to create sandbox: %v", err)
	}

	t.Cleanup(func() {
		_, _ = client.Sandboxes.Delete(ctx, sandboxName)
	})

	t.Run("Exec", func(t *testing.T) {
		t.Run("executes a simple command", func(t *testing.T) {
			result, err := sandbox.Process.New(ctx, blaxel.ProcessRequestParam{
				Command:           "echo 'Hello World'",
				WaitForCompletion: blaxel.Bool(true),
			})
			if err != nil {
				t.Fatalf("failed to exec: %v", err)
			}
			if result.Status != "completed" {
				t.Errorf("expected status completed, got %s", result.Status)
			}
			if !strings.Contains(result.Logs, "Hello World") {
				t.Errorf("expected logs to contain 'Hello World', got %s", result.Logs)
			}
		})

		t.Run("executes command with custom name", func(t *testing.T) {
			result, err := sandbox.Process.New(ctx, blaxel.ProcessRequestParam{
				Name:              blaxel.String("custom-named-process"),
				Command:           "echo 'named'",
				WaitForCompletion: blaxel.Bool(true),
			})
			if err != nil {
				t.Fatalf("failed to exec: %v", err)
			}
			if result.Name != "custom-named-process" {
				t.Errorf("expected name 'custom-named-process', got %s", result.Name)
			}
		})

		t.Run("generates name when not provided", func(t *testing.T) {
			result, err := sandbox.Process.New(ctx, blaxel.ProcessRequestParam{
				Command:           "echo 'auto'",
				WaitForCompletion: blaxel.Bool(true),
			})
			if err != nil {
				t.Fatalf("failed to exec: %v", err)
			}
			if result.Name == "" {
				t.Error("expected name to be generated")
			}
			if !strings.HasPrefix(result.Name, "proc-") {
				t.Errorf("expected name to start with 'proc-', got %s", result.Name)
			}
		})

		t.Run("executes command with working directory", func(t *testing.T) {
			// Create directory first
			_, err := sandbox.Process.New(ctx, blaxel.ProcessRequestParam{
				Command:           "mkdir -p /tmp/workdir",
				WaitForCompletion: blaxel.Bool(true),
			})
			if err != nil {
				t.Fatalf("failed to create dir: %v", err)
			}

			result, err := sandbox.Process.New(ctx, blaxel.ProcessRequestParam{
				Command:           "pwd",
				WorkingDir:        blaxel.String("/tmp/workdir"),
				WaitForCompletion: blaxel.Bool(true),
			})
			if err != nil {
				t.Fatalf("failed to exec: %v", err)
			}
			if !strings.Contains(result.Logs, "/tmp/workdir") {
				t.Errorf("expected logs to contain '/tmp/workdir', got %s", result.Logs)
			}
		})

		t.Run("captures stdout", func(t *testing.T) {
			result, err := sandbox.Process.New(ctx, blaxel.ProcessRequestParam{
				Command:           "echo 'stdout output'",
				WaitForCompletion: blaxel.Bool(true),
			})
			if err != nil {
				t.Fatalf("failed to exec: %v", err)
			}
			if !strings.Contains(result.Logs, "stdout output") {
				t.Errorf("expected logs to contain 'stdout output', got %s", result.Logs)
			}
		})

		t.Run("captures stderr", func(t *testing.T) {
			result, err := sandbox.Process.New(ctx, blaxel.ProcessRequestParam{
				Command:           "echo 'stderr output' >&2",
				WaitForCompletion: blaxel.Bool(true),
			})
			if err != nil {
				t.Fatalf("failed to exec: %v", err)
			}
			if !strings.Contains(result.Logs, "stderr output") {
				t.Errorf("expected logs to contain 'stderr output', got %s", result.Logs)
			}
		})

		t.Run("returns exit code", func(t *testing.T) {
			successResult, err := sandbox.Process.New(ctx, blaxel.ProcessRequestParam{
				Command:           "exit 0",
				WaitForCompletion: blaxel.Bool(true),
			})
			if err != nil {
				t.Fatalf("failed to exec: %v", err)
			}
			if successResult.ExitCode != 0 {
				t.Errorf("expected exit code 0, got %d", successResult.ExitCode)
			}

			failResult, err := sandbox.Process.New(ctx, blaxel.ProcessRequestParam{
				Command:           "exit 42",
				WaitForCompletion: blaxel.Bool(true),
			})
			if err != nil {
				t.Fatalf("failed to exec: %v", err)
			}
			if failResult.ExitCode != 42 {
				t.Errorf("expected exit code 42, got %d", failResult.ExitCode)
			}
		})
	})

	t.Run("ExecWithoutWaiting", func(t *testing.T) {
		t.Run("returns immediately when waitForCompletion is false", func(t *testing.T) {
			startTime := time.Now()

			result, err := sandbox.Process.New(ctx, blaxel.ProcessRequestParam{
				Name:              blaxel.String("no-wait-test"),
				Command:           "sleep 5",
				WaitForCompletion: blaxel.Bool(false),
			})
			if err != nil {
				t.Fatalf("failed to exec: %v", err)
			}

			elapsed := time.Since(startTime)
			if elapsed > 4*time.Second {
				t.Errorf("expected to return quickly, took %v", elapsed)
			}
			if result.Name != "no-wait-test" {
				t.Errorf("expected name 'no-wait-test', got %s", result.Name)
			}

			// Cleanup
			_, _ = sandbox.Process.Kill(ctx, "no-wait-test")
		})
	})

	t.Run("Get", func(t *testing.T) {
		t.Run("retrieves process information", func(t *testing.T) {
			_, err := sandbox.Process.New(ctx, blaxel.ProcessRequestParam{
				Name:              blaxel.String("get-test"),
				Command:           "echo 'test'",
				WaitForCompletion: blaxel.Bool(true),
			})
			if err != nil {
				t.Fatalf("failed to exec: %v", err)
			}

			process, err := sandbox.Process.Get(ctx, "get-test")
			if err != nil {
				t.Fatalf("failed to get process: %v", err)
			}
			if process.Name != "get-test" {
				t.Errorf("expected name 'get-test', got %s", process.Name)
			}
			if process.Status != "completed" {
				t.Errorf("expected status 'completed', got %s", process.Status)
			}
		})

		t.Run("shows running status for long process", func(t *testing.T) {
			_, err := sandbox.Process.New(ctx, blaxel.ProcessRequestParam{
				Name:              blaxel.String("long-running"),
				Command:           "sleep 30",
				WaitForCompletion: blaxel.Bool(false),
			})
			if err != nil {
				t.Fatalf("failed to exec: %v", err)
			}

			process, err := sandbox.Process.Get(ctx, "long-running")
			if err != nil {
				t.Fatalf("failed to get process: %v", err)
			}
			if process.Status != "running" {
				t.Errorf("expected status 'running', got %s", process.Status)
			}

			// Cleanup
			_, _ = sandbox.Process.Kill(ctx, "long-running")
		})
	})

	t.Run("Logs", func(t *testing.T) {
		t.Run("retrieves all logs", func(t *testing.T) {
			_, err := sandbox.Process.New(ctx, blaxel.ProcessRequestParam{
				Name:              blaxel.String("logs-test"),
				Command:           "echo 'stdout' && echo 'stderr' >&2",
				WaitForCompletion: blaxel.Bool(true),
			})
			if err != nil {
				t.Fatalf("failed to exec: %v", err)
			}

			logs, err := sandbox.Process.GetLogs(ctx, "logs-test")
			if err != nil {
				t.Fatalf("failed to get logs: %v", err)
			}
			if !strings.Contains(logs.Stdout, "stdout") {
				t.Errorf("expected stdout to contain 'stdout', got %s", logs.Stdout)
			}
			if !strings.Contains(logs.Stderr, "stderr") {
				t.Errorf("expected stderr to contain 'stderr', got %s", logs.Stderr)
			}
		})

		t.Run("retrieves stdout only", func(t *testing.T) {
			_, err := sandbox.Process.New(ctx, blaxel.ProcessRequestParam{
				Name:              blaxel.String("stdout-only"),
				Command:           "echo 'out' && echo 'err' >&2",
				WaitForCompletion: blaxel.Bool(true),
			})
			if err != nil {
				t.Fatalf("failed to exec: %v", err)
			}

			logs, err := sandbox.Process.GetLogs(ctx, "stdout-only")
			if err != nil {
				t.Fatalf("failed to get logs: %v", err)
			}
			if !strings.Contains(logs.Stdout, "out") {
				t.Errorf("expected stdout to contain 'out', got %s", logs.Stdout)
			}
		})

		t.Run("retrieves stderr only", func(t *testing.T) {
			_, err := sandbox.Process.New(ctx, blaxel.ProcessRequestParam{
				Name:              blaxel.String("stderr-only"),
				Command:           "echo 'out' && echo 'err' >&2",
				WaitForCompletion: blaxel.Bool(true),
			})
			if err != nil {
				t.Fatalf("failed to exec: %v", err)
			}

			logs, err := sandbox.Process.GetLogs(ctx, "stderr-only")
			if err != nil {
				t.Fatalf("failed to get logs: %v", err)
			}
			if !strings.Contains(logs.Stderr, "err") {
				t.Errorf("expected stderr to contain 'err', got %s", logs.Stderr)
			}
		})
	})

	t.Run("StreamLogs", func(t *testing.T) {
		t.Run("streams logs in real-time", func(t *testing.T) {
			var logs []string
			var mu sync.Mutex

			_, err := sandbox.Process.New(ctx, blaxel.ProcessRequestParam{
				Name:              blaxel.String("stream-test"),
				Command:           "for i in 1 2 3; do echo \"msg $i\"; sleep 1; done",
				WaitForCompletion: blaxel.Bool(false),
			})
			if err != nil {
				t.Fatalf("failed to exec: %v", err)
			}

			stream := sandbox.Process.StreamLogs(ctx, "stream-test", blaxel.ProcessStreamOptions{
				OnLog: func(log string) {
					mu.Lock()
					logs = append(logs, log)
					mu.Unlock()
				},
			})

			_, err = sandbox.Process.Wait(ctx, "stream-test", 30*time.Second, time.Second)
			if err != nil {
				t.Fatalf("failed to wait: %v", err)
			}
			time.Sleep(1 * time.Second)
			stream.Close()

			mu.Lock()
			logCount := len(logs)
			mu.Unlock()

			if logCount == 0 {
				t.Error("expected to receive logs")
			}
		})

		t.Run("can close stream early", func(t *testing.T) {
			var logs []string
			var mu sync.Mutex

			_, err := sandbox.Process.New(ctx, blaxel.ProcessRequestParam{
				Name:              blaxel.String("close-early"),
				Command:           "for i in $(seq 1 10); do echo $i; sleep 1; done",
				WaitForCompletion: blaxel.Bool(false),
			})
			if err != nil {
				t.Fatalf("failed to exec: %v", err)
			}

			stream := sandbox.Process.StreamLogs(ctx, "close-early", blaxel.ProcessStreamOptions{
				OnLog: func(log string) {
					mu.Lock()
					logs = append(logs, log)
					mu.Unlock()
				},
			})

			time.Sleep(2 * time.Second)
			stream.Close()

			mu.Lock()
			logsAtClose := len(logs)
			mu.Unlock()

			time.Sleep(3 * time.Second)

			mu.Lock()
			finalLogs := len(logs)
			mu.Unlock()

			// No new logs should arrive after close
			if finalLogs != logsAtClose {
				t.Errorf("expected no new logs after close, got %d more", finalLogs-logsAtClose)
			}

			// Cleanup
			_, _ = sandbox.Process.Kill(ctx, "close-early")
		})
	})

	t.Run("Wait", func(t *testing.T) {
		t.Run("waits for process completion", func(t *testing.T) {
			_, err := sandbox.Process.New(ctx, blaxel.ProcessRequestParam{
				Name:              blaxel.String("wait-test"),
				Command:           "sleep 2 && echo 'done'",
				WaitForCompletion: blaxel.Bool(false),
			})
			if err != nil {
				t.Fatalf("failed to exec: %v", err)
			}

			_, err = sandbox.Process.Wait(ctx, "wait-test", 30*time.Second, time.Second)
			if err != nil {
				t.Fatalf("failed to wait: %v", err)
			}

			process, err := sandbox.Process.Get(ctx, "wait-test")
			if err != nil {
				t.Fatalf("failed to get process: %v", err)
			}
			if process.Status != "completed" {
				t.Errorf("expected status 'completed', got %s", process.Status)
			}
		})

		t.Run("respects maxWait timeout", func(t *testing.T) {
			_, err := sandbox.Process.New(ctx, blaxel.ProcessRequestParam{
				Name:              blaxel.String("timeout-test"),
				Command:           "sleep 60",
				WaitForCompletion: blaxel.Bool(false),
			})
			if err != nil {
				t.Fatalf("failed to exec: %v", err)
			}

			_, err = sandbox.Process.Wait(ctx, "timeout-test", 2*time.Second, 500*time.Millisecond)
			if err == nil {
				t.Error("expected timeout error")
			}
			if !strings.Contains(err.Error(), "did not finish in time") {
				t.Errorf("expected timeout error message, got %v", err)
			}

			// Cleanup
			_, _ = sandbox.Process.Kill(ctx, "timeout-test")
		})
	})

	t.Run("Kill", func(t *testing.T) {
		t.Run("kills a running process", func(t *testing.T) {
			_, err := sandbox.Process.New(ctx, blaxel.ProcessRequestParam{
				Name:              blaxel.String("kill-test"),
				Command:           "sleep 60",
				WaitForCompletion: blaxel.Bool(false),
			})
			if err != nil {
				t.Fatalf("failed to exec: %v", err)
			}

			process, err := sandbox.Process.Get(ctx, "kill-test")
			if err != nil {
				t.Fatalf("failed to get process: %v", err)
			}
			if process.Status != "running" {
				t.Errorf("expected status 'running', got %s", process.Status)
			}

			_, err = sandbox.Process.Kill(ctx, "kill-test")
			if err != nil {
				t.Fatalf("failed to kill process: %v", err)
			}

			time.Sleep(1 * time.Second)

			process, err = sandbox.Process.Get(ctx, "kill-test")
			if err != nil {
				t.Fatalf("failed to get process: %v", err)
			}
			validStatuses := map[blaxel.ProcessResponseStatus]bool{
				blaxel.ProcessResponseStatusKilled:    true,
				blaxel.ProcessResponseStatusFailed:    true,
				blaxel.ProcessResponseStatusCompleted: true,
			}
			if !validStatuses[process.Status] {
				t.Errorf("expected status to be killed/failed/completed, got %s", process.Status)
			}
		})
	})

	t.Run("RestartOnFailure", func(t *testing.T) {
		t.Run("restarts process on failure", func(t *testing.T) {
			result, err := sandbox.Process.New(ctx, blaxel.ProcessRequestParam{
				Name:              blaxel.String("restart-test"),
				Command:           "exit 1",
				RestartOnFailure:  blaxel.Bool(true),
				MaxRestarts:       blaxel.Int(3),
				WaitForCompletion: blaxel.Bool(true),
			})
			if err != nil {
				t.Fatalf("failed to exec: %v", err)
			}
			if result.RestartCount == 0 {
				t.Error("expected restart count > 0")
			}
			if result.RestartCount > 3 {
				t.Errorf("expected restart count <= 3, got %d", result.RestartCount)
			}
		})
	})
}
