package integration_tests

import (
	"context"
	"os"
	"strings"
	"testing"

	blaxel "github.com/stainless-sdks/blaxel-go"
)

func TestSandboxCodegen(t *testing.T) {
	// These tests require RELACE_API_KEY or MORPH_API_KEY environment variables
	hasRelaceKey := os.Getenv("RELACE_API_KEY") != ""
	hasMorphKey := os.Getenv("MORPH_API_KEY") != ""

	ctx := context.Background()
	client := newTestClient(t)

	if !hasRelaceKey && !hasMorphKey {
		t.Log("Codegen tests skipped - set RELACE_API_KEY or MORPH_API_KEY to run")
		return
	}

	t.Run("FastapplyWithRelace", func(t *testing.T) {
		if !hasRelaceKey {
			t.Skip("RELACE_API_KEY not set")
		}

		sandboxName := uniqueName("codegen-relace")
		sandbox, err := client.Sandboxes.NewInstance(ctx, blaxel.SandboxNewParams{
			Sandbox: blaxel.SandboxParam{
				Metadata: blaxel.MetadataParam{
					Name:   sandboxName,
					Labels: defaultLabels,
				},
				Spec: blaxel.SandboxSpecParam{
					Runtime: blaxel.SandboxRuntimeParam{
						Image: blaxel.String(defaultImage),
						Envs: []blaxel.SandboxRuntimeEnvParam{
							{Name: blaxel.String("RELACE_API_KEY"), Value: blaxel.String(os.Getenv("RELACE_API_KEY"))},
						},
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

		t.Run("applies code edit to create new file", func(t *testing.T) {
			_, err := sandbox.Codegen.Fastapply(ctx, "/tmp/test.txt", blaxel.ApplyEditRequestParam{
				CodeEdit: "// ... existing code ...\nconsole.log('Hello, world!');",
			})
			if err != nil {
				t.Fatalf("failed to fastapply: %v", err)
			}

			content, err := sandbox.FS.Read(ctx, "/tmp/test.txt")
			if err != nil {
				t.Fatalf("failed to read file: %v", err)
			}
			if !strings.Contains(content, "Hello, world!") {
				t.Errorf("expected content to contain 'Hello, world!', got %s", content)
			}
		})

		t.Run("preserves existing content when applying edits", func(t *testing.T) {
			// First edit
			_, err := sandbox.Codegen.Fastapply(ctx, "/tmp/preserve-test.txt", blaxel.ApplyEditRequestParam{
				CodeEdit: "// ... existing code ...\nconsole.log('First line');",
			})
			if err != nil {
				t.Fatalf("failed to fastapply: %v", err)
			}

			// Second edit - should preserve first line
			_, err = sandbox.Codegen.Fastapply(ctx, "/tmp/preserve-test.txt", blaxel.ApplyEditRequestParam{
				CodeEdit: "// ... keep existing code\nconsole.log('Second line');",
			})
			if err != nil {
				t.Fatalf("failed to fastapply: %v", err)
			}

			content, err := sandbox.FS.Read(ctx, "/tmp/preserve-test.txt")
			if err != nil {
				t.Fatalf("failed to read file: %v", err)
			}
			if !strings.Contains(content, "Second line") {
				t.Errorf("expected content to contain 'Second line', got %s", content)
			}
		})

		t.Run("performs reranking search", func(t *testing.T) {
			// Create test file
			_, err := sandbox.FS.Write(ctx, "/tmp/search-test.txt", "The meaning of life is 42")
			if err != nil {
				t.Fatalf("failed to write file: %v", err)
			}

			result, err := sandbox.Codegen.Reranking(ctx, "/tmp", blaxel.SandboxCodegenRerankingParams{
				Query:          "What is the meaning of life?",
				ScoreThreshold: blaxel.Float(0.01),
				TokenLimit:     blaxel.Int(1000),
				FilePattern:    blaxel.String(`.*\.txt$`),
			})
			if err != nil {
				t.Fatalf("failed to reranking: %v", err)
			}

			if result == nil {
				t.Fatal("expected result to be returned")
			}
			if len(result.Files) == 0 {
				t.Error("expected at least one file in result")
			}

			found := false
			for _, f := range result.Files {
				if strings.Contains(f.Path, "search-test.txt") {
					found = true
					break
				}
			}
			if !found {
				t.Error("expected to find search-test.txt in results")
			}
		})
	})

	t.Run("FastapplyWithMorph", func(t *testing.T) {
		if !hasMorphKey {
			t.Skip("MORPH_API_KEY not set")
		}

		sandboxName := uniqueName("codegen-morph")
		sandbox, err := client.Sandboxes.NewInstance(ctx, blaxel.SandboxNewParams{
			Sandbox: blaxel.SandboxParam{
				Metadata: blaxel.MetadataParam{
					Name:   sandboxName,
					Labels: defaultLabels,
				},
				Spec: blaxel.SandboxSpecParam{
					Runtime: blaxel.SandboxRuntimeParam{
						Image: blaxel.String(defaultImage),
						Envs: []blaxel.SandboxRuntimeEnvParam{
							{Name: blaxel.String("MORPH_API_KEY"), Value: blaxel.String(os.Getenv("MORPH_API_KEY"))},
						},
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

		t.Run("applies code edit with Morph backend", func(t *testing.T) {
			_, err := sandbox.Codegen.Fastapply(ctx, "/tmp/morph-test.txt", blaxel.ApplyEditRequestParam{
				CodeEdit: "// ... existing code ...\nconsole.log('Hello from Morph!');",
			})
			if err != nil {
				t.Fatalf("failed to fastapply: %v", err)
			}

			content, err := sandbox.FS.Read(ctx, "/tmp/morph-test.txt")
			if err != nil {
				t.Fatalf("failed to read file: %v", err)
			}
			if !strings.Contains(content, "Hello from Morph!") {
				t.Errorf("expected content to contain 'Hello from Morph!', got %s", content)
			}
		})

		t.Run("performs reranking with Morph", func(t *testing.T) {
			_, err := sandbox.FS.Write(ctx, "/tmp/morph-search.txt", "The answer is always 42")
			if err != nil {
				t.Fatalf("failed to write file: %v", err)
			}

			result, err := sandbox.Codegen.Reranking(ctx, "/tmp", blaxel.SandboxCodegenRerankingParams{
				Query:          "What is the answer?",
				ScoreThreshold: blaxel.Float(0.01),
				TokenLimit:     blaxel.Int(1000000),
				FilePattern:    blaxel.String(`.*\.txt$`),
			})
			if err != nil {
				t.Fatalf("failed to reranking: %v", err)
			}

			if result == nil {
				t.Fatal("expected result to be returned")
			}
			if len(result.Files) == 0 {
				t.Error("expected at least one file in result")
			}
		})
	})
}
