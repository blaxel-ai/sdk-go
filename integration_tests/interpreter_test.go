package integration_tests

import (
	"context"
	"strings"
	"testing"
	"time"

	blaxel "github.com/blaxel-ai/sdk-go"
)

func TestCodeInterpreter(t *testing.T) {
	ctx := context.Background()
	client := newTestClient(t)

	interpreterName := uniqueName("interp")
	interpreter, err := client.Sandboxes.NewCodeInterpreter(ctx, &blaxel.CodeInterpreterCreateConfig{
		Name:   interpreterName,
		Labels: defaultLabels,
	})
	if err != nil {
		t.Fatalf("failed to create code interpreter: %v", err)
	}

	t.Cleanup(func() {
		_, _ = client.Sandboxes.Delete(ctx, interpreterName)
	})

	t.Run("Create", func(t *testing.T) {
		t.Run("creates a code interpreter", func(t *testing.T) {
			if interpreter.Metadata.Name == "" {
				t.Error("expected interpreter name to be set")
			}
		})
	})

	t.Run("CreateCodeContext", func(t *testing.T) {
		t.Run("creates a Python code context", func(t *testing.T) {
			codeCtx, err := interpreter.CreateCodeContext(ctx, &blaxel.CreateContextOptions{
				Language: "python",
			})
			if err != nil {
				t.Fatalf("failed to create code context: %v", err)
			}

			if codeCtx == nil {
				t.Error("expected context to be returned")
			}
			if codeCtx.ID == "" {
				t.Error("expected context ID to be set")
			}
		})
	})

	t.Run("RunCode", func(t *testing.T) {
		t.Run("executes simple Python code", func(t *testing.T) {
			var stdoutLines []string

			_, err := interpreter.RunCode(ctx, "print('Hello from interpreter')", &blaxel.RunCodeOptions{
				Language: "python",
				OnStdout: func(msg blaxel.OutputMessage) {
					stdoutLines = append(stdoutLines, msg.Text)
				},
				Timeout: 30 * time.Second,
			})
			if err != nil {
				t.Fatalf("failed to run code: %v", err)
			}

			stdout := strings.Join(stdoutLines, "")
			if !strings.Contains(stdout, "Hello from interpreter") {
				t.Errorf("expected stdout to contain 'Hello from interpreter', got %s", stdout)
			}
		})

		t.Run("captures stderr output", func(t *testing.T) {
			var stderrLines []string

			_, err := interpreter.RunCode(ctx, "import sys; sys.stderr.write('error message')", &blaxel.RunCodeOptions{
				Language: "python",
				OnStderr: func(msg blaxel.OutputMessage) {
					stderrLines = append(stderrLines, msg.Text)
				},
				Timeout: 30 * time.Second,
			})
			if err != nil {
				t.Fatalf("failed to run code: %v", err)
			}

			stderr := strings.Join(stderrLines, "")
			if !strings.Contains(stderr, "error message") {
				t.Errorf("expected stderr to contain 'error message', got %s", stderr)
			}
		})

		t.Run("returns execution results", func(t *testing.T) {
			var results []blaxel.ExecutionResult

			_, err := interpreter.RunCode(ctx, "2 + 2", &blaxel.RunCodeOptions{
				Language: "python",
				OnResult: func(res blaxel.ExecutionResult) {
					results = append(results, res)
				},
				Timeout: 30 * time.Second,
			})
			if err != nil {
				t.Fatalf("failed to run code: %v", err)
			}

			if len(results) == 0 {
				t.Error("expected at least one result")
			}
		})

		t.Run("captures execution errors", func(t *testing.T) {
			var errors []blaxel.ExecutionError

			_, err := interpreter.RunCode(ctx, "raise ValueError('test error')", &blaxel.RunCodeOptions{
				Language: "python",
				OnError: func(err blaxel.ExecutionError) {
					errors = append(errors, err)
				},
				Timeout: 30 * time.Second,
			})
			if err != nil {
				t.Fatalf("failed to run code: %v", err)
			}

			if len(errors) == 0 {
				t.Error("expected at least one error")
			}
			if errors[0].Name != "ValueError" {
				t.Errorf("expected error name 'ValueError', got %s", errors[0].Name)
			}
		})

		t.Run("persists state across runs", func(t *testing.T) {
			// Define a function in first run
			_, err := interpreter.RunCode(ctx, "def add(a, b):\n    return a + b", &blaxel.RunCodeOptions{
				Language: "python",
				Timeout:  30 * time.Second,
			})
			if err != nil {
				t.Fatalf("failed to run code: %v", err)
			}

			// Call the function in second run
			var stdoutLines []string
			_, err = interpreter.RunCode(ctx, "print(add(2, 3))", &blaxel.RunCodeOptions{
				Language: "python",
				OnStdout: func(msg blaxel.OutputMessage) {
					stdoutLines = append(stdoutLines, msg.Text)
				},
				Timeout: 30 * time.Second,
			})
			if err != nil {
				t.Fatalf("failed to run code: %v", err)
			}

			stdout := strings.Join(stdoutLines, "")
			if !strings.Contains(stdout, "5") {
				t.Errorf("expected stdout to contain '5', got %s", stdout)
			}
		})

		t.Run("handles variables across runs", func(t *testing.T) {
			// Set variable
			_, err := interpreter.RunCode(ctx, "x = 42", &blaxel.RunCodeOptions{
				Language: "python",
				Timeout:  30 * time.Second,
			})
			if err != nil {
				t.Fatalf("failed to run code: %v", err)
			}

			// Read variable
			var stdoutLines []string
			_, err = interpreter.RunCode(ctx, "print(x)", &blaxel.RunCodeOptions{
				Language: "python",
				OnStdout: func(msg blaxel.OutputMessage) {
					stdoutLines = append(stdoutLines, msg.Text)
				},
				Timeout: 30 * time.Second,
			})
			if err != nil {
				t.Fatalf("failed to run code: %v", err)
			}

			stdout := strings.Join(stdoutLines, "")
			if !strings.Contains(stdout, "42") {
				t.Errorf("expected stdout to contain '42', got %s", stdout)
			}
		})
	})

	t.Run("StaticMethods", func(t *testing.T) {
		t.Run("gets an existing interpreter", func(t *testing.T) {
			retrieved, err := client.Sandboxes.GetCodeInterpreter(ctx, interpreterName)
			if err != nil {
				t.Fatalf("failed to get interpreter: %v", err)
			}

			if retrieved.Metadata.Name != interpreterName {
				t.Errorf("expected name %s, got %s", interpreterName, retrieved.Metadata.Name)
			}
		})
	})
}
