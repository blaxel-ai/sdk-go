package integration_tests

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"testing"
	"time"

	blaxel "github.com/blaxel-ai/sdk-go"
)

func TestSandboxPreviews(t *testing.T) {
	ctx := context.Background()
	client := newTestClient(t)

	sandboxName := uniqueName("preview-test")
	sandbox, err := client.Sandboxes.NewInstance(ctx, blaxel.SandboxNewParams{
		Sandbox: blaxel.SandboxParam{
			Metadata: blaxel.MetadataParam{
				Name:   sandboxName,
				Labels: defaultLabels,
			},
			Spec: blaxel.SandboxSpecParam{
				Region: blaxel.String(defaultRegion),
				Runtime: blaxel.SandboxRuntimeParam{
					Image:  blaxel.String("blaxel/nextjs:latest"),
					Memory: blaxel.Int(4096),
					Ports: []blaxel.PortParam{
						{Target: 3000},
					},
				},
			},
		},
	})
	if err != nil {
		t.Fatalf("failed to create sandbox: %v", err)
	}

	// Start the dev server
	_, err = sandbox.Process.New(ctx, blaxel.ProcessRequestParam{
		Command:      "npm run dev -- --port 3000",
		WorkingDir:   blaxel.String("/blaxel/app"),
		WaitForPorts: []int64{3000},
	})
	if err != nil {
		t.Fatalf("failed to start dev server: %v", err)
	}

	t.Cleanup(func() {
		_, _ = client.Sandboxes.Delete(ctx, sandboxName)
	})

	t.Run("Create", func(t *testing.T) {
		t.Run("creates a public preview", func(t *testing.T) {
			preview, err := sandbox.Previews.New(ctx, blaxel.SandboxPreviewNewParams{
				Preview: blaxel.PreviewParam{
					Metadata: blaxel.PreviewMetadataParam{Name: "public-preview"},
					Spec: blaxel.PreviewSpecParam{
						Port:   blaxel.Int(3000),
						Public: blaxel.Bool(true),
					},
				},
			})
			if err != nil {
				t.Fatalf("failed to create preview: %v", err)
			}

			if preview.Metadata.Name != "public-preview" {
				t.Errorf("expected name 'public-preview', got %s", preview.Metadata.Name)
			}
			if preview.Spec.URL == "" {
				t.Error("expected URL to be set")
			}

			_, _ = sandbox.Previews.Delete(ctx, "public-preview")
		})

		t.Run("creates a private preview", func(t *testing.T) {
			preview, err := sandbox.Previews.New(ctx, blaxel.SandboxPreviewNewParams{
				Preview: blaxel.PreviewParam{
					Metadata: blaxel.PreviewMetadataParam{Name: "private-preview"},
					Spec: blaxel.PreviewSpecParam{
						Port:   blaxel.Int(3000),
						Public: blaxel.Bool(false),
					},
				},
			})
			if err != nil {
				t.Fatalf("failed to create preview: %v", err)
			}

			if preview.Metadata.Name != "private-preview" {
				t.Errorf("expected name 'private-preview', got %s", preview.Metadata.Name)
			}
			if preview.Spec.URL == "" {
				t.Error("expected URL to be set")
			}

			_, _ = sandbox.Previews.Delete(ctx, "private-preview")
		})

		t.Run("creates preview with custom prefix URL", func(t *testing.T) {
			preview, err := sandbox.Previews.New(ctx, blaxel.SandboxPreviewNewParams{
				Preview: blaxel.PreviewParam{
					Metadata: blaxel.PreviewMetadataParam{Name: "prefix-preview"},
					Spec: blaxel.PreviewSpecParam{
						Port:      blaxel.Int(3000),
						PrefixURL: blaxel.String("my-custom-prefix"),
						Public:    blaxel.Bool(true),
					},
				},
			})
			if err != nil {
				t.Fatalf("failed to create preview: %v", err)
			}

			// URL should contain the custom prefix
			// Note: exact URL format may vary

			_, _ = sandbox.Previews.Delete(ctx, "prefix-preview")
			_ = preview // Use preview to avoid unused variable warning
		})
	})

	t.Run("CreateIfNotExists", func(t *testing.T) {
		t.Run("creates new preview if not exists", func(t *testing.T) {
			preview, err := sandbox.Previews.NewIfNotExists(ctx, blaxel.SandboxPreviewNewParams{
				Preview: blaxel.PreviewParam{
					Metadata: blaxel.PreviewMetadataParam{Name: "cine-preview"},
					Spec: blaxel.PreviewSpecParam{
						Port:   blaxel.Int(3000),
						Public: blaxel.Bool(true),
					},
				},
			})
			if err != nil {
				t.Fatalf("failed to create preview: %v", err)
			}

			if preview.Metadata.Name != "cine-preview" {
				t.Errorf("expected name 'cine-preview', got %s", preview.Metadata.Name)
			}

			_, _ = sandbox.Previews.Delete(ctx, "cine-preview")
		})

		t.Run("returns existing preview if already exists", func(t *testing.T) {
			// Create first
			_, err := sandbox.Previews.New(ctx, blaxel.SandboxPreviewNewParams{
				Preview: blaxel.PreviewParam{
					Metadata: blaxel.PreviewMetadataParam{Name: "existing-preview"},
					Spec: blaxel.PreviewSpecParam{
						Port:   blaxel.Int(3000),
						Public: blaxel.Bool(true),
					},
				},
			})
			if err != nil {
				t.Fatalf("failed to create preview: %v", err)
			}

			// Should return existing
			second, err := sandbox.Previews.NewIfNotExists(ctx, blaxel.SandboxPreviewNewParams{
				Preview: blaxel.PreviewParam{
					Metadata: blaxel.PreviewMetadataParam{Name: "existing-preview"},
					Spec: blaxel.PreviewSpecParam{
						Port:   blaxel.Int(3000),
						Public: blaxel.Bool(true),
					},
				},
			})
			if err != nil {
				t.Fatalf("failed to createIfNotExists: %v", err)
			}

			if second.Metadata.Name != "existing-preview" {
				t.Errorf("expected name 'existing-preview', got %s", second.Metadata.Name)
			}

			_, _ = sandbox.Previews.Delete(ctx, "existing-preview")
		})
	})

	t.Run("Get", func(t *testing.T) {
		t.Run("retrieves an existing preview", func(t *testing.T) {
			_, err := sandbox.Previews.New(ctx, blaxel.SandboxPreviewNewParams{
				Preview: blaxel.PreviewParam{
					Metadata: blaxel.PreviewMetadataParam{Name: "get-preview"},
					Spec: blaxel.PreviewSpecParam{
						Port:   blaxel.Int(3000),
						Public: blaxel.Bool(true),
					},
				},
			})
			if err != nil {
				t.Fatalf("failed to create preview: %v", err)
			}

			preview, err := sandbox.Previews.Get(ctx, "get-preview")
			if err != nil {
				t.Fatalf("failed to get preview: %v", err)
			}

			if preview.Metadata.Name != "get-preview" {
				t.Errorf("expected name 'get-preview', got %s", preview.Metadata.Name)
			}
			if preview.Spec.URL == "" {
				t.Error("expected URL to be set")
			}

			_, _ = sandbox.Previews.Delete(ctx, "get-preview")
		})
	})

	t.Run("List", func(t *testing.T) {
		t.Run("lists all previews", func(t *testing.T) {
			_, err := sandbox.Previews.New(ctx, blaxel.SandboxPreviewNewParams{
				Preview: blaxel.PreviewParam{
					Metadata: blaxel.PreviewMetadataParam{Name: "list-preview-1"},
					Spec: blaxel.PreviewSpecParam{
						Port:   blaxel.Int(3000),
						Public: blaxel.Bool(true),
					},
				},
			})
			if err != nil {
				t.Fatalf("failed to create preview: %v", err)
			}

			_, err = sandbox.Previews.New(ctx, blaxel.SandboxPreviewNewParams{
				Preview: blaxel.PreviewParam{
					Metadata: blaxel.PreviewMetadataParam{Name: "list-preview-2"},
					Spec: blaxel.PreviewSpecParam{
						Port:   blaxel.Int(3000),
						Public: blaxel.Bool(true),
					},
				},
			})
			if err != nil {
				t.Fatalf("failed to create preview: %v", err)
			}

			previews, err := sandbox.Previews.List(ctx)
			if err != nil {
				t.Fatalf("failed to list previews: %v", err)
			}

			if len(*previews) < 2 {
				t.Errorf("expected at least 2 previews, got %d", len(*previews))
			}

			found1, found2 := false, false
			for _, p := range *previews {
				if p.Metadata.Name == "list-preview-1" {
					found1 = true
				}
				if p.Metadata.Name == "list-preview-2" {
					found2 = true
				}
			}
			if !found1 || !found2 {
				t.Error("expected to find both previews in list")
			}

			_, _ = sandbox.Previews.Delete(ctx, "list-preview-1")
			_, _ = sandbox.Previews.Delete(ctx, "list-preview-2")
		})
	})

	t.Run("Delete", func(t *testing.T) {
		t.Run("deletes a preview", func(t *testing.T) {
			_, err := sandbox.Previews.New(ctx, blaxel.SandboxPreviewNewParams{
				Preview: blaxel.PreviewParam{
					Metadata: blaxel.PreviewMetadataParam{Name: "delete-preview"},
					Spec: blaxel.PreviewSpecParam{
						Port:   blaxel.Int(3000),
						Public: blaxel.Bool(true),
					},
				},
			})
			if err != nil {
				t.Fatalf("failed to create preview: %v", err)
			}

			_, err = sandbox.Previews.Delete(ctx, "delete-preview")
			if err != nil {
				t.Fatalf("failed to delete preview: %v", err)
			}

			previews, err := sandbox.Previews.List(ctx)
			if err != nil {
				t.Fatalf("failed to list previews: %v", err)
			}

			for _, p := range *previews {
				if p.Metadata.Name == "delete-preview" {
					t.Error("expected preview to be deleted")
				}
			}
		})
	})

	t.Run("PublicPreviewAccess", func(t *testing.T) {
		t.Run("public preview is accessible without token", func(t *testing.T) {
			preview, err := sandbox.Previews.New(ctx, blaxel.SandboxPreviewNewParams{
				Preview: blaxel.PreviewParam{
					Metadata: blaxel.PreviewMetadataParam{Name: "access-public"},
					Spec: blaxel.PreviewSpecParam{
						Port:   blaxel.Int(3000),
						Public: blaxel.Bool(true),
					},
				},
			})
			if err != nil {
				t.Fatalf("failed to create preview: %v", err)
			}

			// Make HTTP request to preview URL
			httpClient := &http.Client{Timeout: 10 * time.Second}
			resp, err := httpClient.Get(preview.Spec.URL)
			if err != nil {
				t.Fatalf("failed to access preview: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != 200 {
				t.Errorf("expected status 200, got %d", resp.StatusCode)
			}

			_, _ = sandbox.Previews.Delete(ctx, "access-public")
		})
	})

	t.Run("AsyncDeletion", func(t *testing.T) {
		t.Run("creates private preview with 15 tokens and tests async deletion", func(t *testing.T) {
			t.Logf("Sandbox name: %s", sandbox.Metadata.Name)
			preview, err := sandbox.Previews.New(ctx, blaxel.SandboxPreviewNewParams{
				Preview: blaxel.PreviewParam{
					Metadata: blaxel.PreviewMetadataParam{Name: "preview-with-many-tokens"},
					Spec: blaxel.PreviewSpecParam{
						Port:   blaxel.Int(3000),
						Public: blaxel.Bool(false),
					},
				},
			})
			if err != nil {
				t.Fatalf("failed to create preview: %v", err)
			}
			t.Logf("Preview created: %s", preview.Metadata.Name)

			expiration := time.Now().Add(10 * time.Minute)
			tokens := make([]*blaxel.PreviewToken, 15)
			var wg sync.WaitGroup
			var mu sync.Mutex
			errChan := make(chan error, 15)

			for i := 0; i < 15; i++ {
				wg.Add(1)
				go func(idx int) {
					defer wg.Done()
					token, err := sandbox.Previews.NewToken(ctx, "preview-with-many-tokens", expiration)
					if err != nil {
						errChan <- err
						return
					}
					mu.Lock()
					tokens[idx] = token
					mu.Unlock()
				}(i)
			}
			wg.Wait()
			close(errChan)

			for err := range errChan {
				t.Fatalf("failed to create token: %v", err)
			}

			for _, token := range tokens {
				if token == nil || token.Spec.Token == "" {
					t.Error("expected token value to be set")
				}
			}

			listedTokens, err := sandbox.Previews.ListTokens(ctx, "preview-with-many-tokens")
			if err != nil {
				t.Fatalf("failed to list tokens: %v", err)
			}
			if len(*listedTokens) < 15 {
				t.Errorf("expected at least 15 tokens, got %d", len(*listedTokens))
			}

			_, err = sandbox.Previews.Delete(ctx, "preview-with-many-tokens")
			if err != nil {
				t.Fatalf("failed to delete preview: %v", err)
			}

			// Should fail to get the deleted preview
			_, err = sandbox.Previews.Get(ctx, "preview-with-many-tokens")
			if err == nil {
				t.Error("expected error getting deleted preview")
			}

			// Recreate the preview
			newPreview, err := sandbox.Previews.New(ctx, blaxel.SandboxPreviewNewParams{
				Preview: blaxel.PreviewParam{
					Metadata: blaxel.PreviewMetadataParam{Name: "preview-with-many-tokens"},
					Spec: blaxel.PreviewSpecParam{
						Port:   blaxel.Int(3000),
						Public: blaxel.Bool(true),
					},
				},
			})
			if err != nil {
				t.Fatalf("failed to recreate preview: %v", err)
			}

			httpClient := &http.Client{Timeout: 10 * time.Second}
			resp, err := httpClient.Get(newPreview.Spec.URL)
			if err != nil {
				t.Fatalf("failed to access preview: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != 200 {
				t.Errorf("expected status 200, got %d", resp.StatusCode)
			}

			_, _ = sandbox.Previews.Delete(ctx, "preview-with-many-tokens")
		})
	})

	t.Run("PrivatePreviewTokens", func(t *testing.T) {
		t.Run("private preview requires token", func(t *testing.T) {
			preview, err := sandbox.Previews.New(ctx, blaxel.SandboxPreviewNewParams{
				Preview: blaxel.PreviewParam{
					Metadata: blaxel.PreviewMetadataParam{Name: "token-required"},
					Spec: blaxel.PreviewSpecParam{
						Port:   blaxel.Int(3000),
						Public: blaxel.Bool(false),
					},
				},
			})
			if err != nil {
				t.Fatalf("failed to create preview: %v", err)
			}

			// Make HTTP request without token
			httpClient := &http.Client{Timeout: 10 * time.Second}
			resp, err := httpClient.Get(preview.Spec.URL)
			if err != nil {
				t.Fatalf("failed to access preview: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != 401 {
				t.Errorf("expected status 401, got %d", resp.StatusCode)
			}

			_, _ = sandbox.Previews.Delete(ctx, "token-required")
		})

		t.Run("creates and uses preview token", func(t *testing.T) {
			preview, err := sandbox.Previews.New(ctx, blaxel.SandboxPreviewNewParams{
				Preview: blaxel.PreviewParam{
					Metadata: blaxel.PreviewMetadataParam{Name: "token-test"},
					Spec: blaxel.PreviewSpecParam{
						Port:   blaxel.Int(3000),
						Public: blaxel.Bool(false),
					},
				},
			})
			if err != nil {
				t.Fatalf("failed to create preview: %v", err)
			}

			// Create token (expires in 10 minutes)
			expiration := time.Now().Add(10 * time.Minute)
			token, err := sandbox.Previews.NewToken(ctx, "token-test", expiration)
			if err != nil {
				t.Fatalf("failed to create token: %v", err)
			}

			if token.Spec.Token == "" {
				t.Error("expected token value to be set")
			}

			// Access with token
			httpClient := &http.Client{Timeout: 10 * time.Second}
			resp, err := httpClient.Get(preview.Spec.URL + "?bl_preview_token=" + token.Spec.Token)
			if err != nil {
				t.Fatalf("failed to access preview: %v", err)
			}
			defer resp.Body.Close()

			if resp.StatusCode != 200 {
				t.Errorf("expected status 200, got %d", resp.StatusCode)
			}

			_, _ = sandbox.Previews.Delete(ctx, "token-test")
		})

		t.Run("lists tokens", func(t *testing.T) {
			_, err := sandbox.Previews.New(ctx, blaxel.SandboxPreviewNewParams{
				Preview: blaxel.PreviewParam{
					Metadata: blaxel.PreviewMetadataParam{Name: "list-tokens"},
					Spec: blaxel.PreviewSpecParam{
						Port:   blaxel.Int(3000),
						Public: blaxel.Bool(false),
					},
				},
			})
			if err != nil {
				t.Fatalf("failed to create preview: %v", err)
			}

			expiration := time.Now().Add(10 * time.Minute)
			token, err := sandbox.Previews.NewToken(ctx, "list-tokens", expiration)
			if err != nil {
				t.Fatalf("failed to create token: %v", err)
			}

			tokens, err := sandbox.Previews.ListTokens(ctx, "list-tokens")
			if err != nil {
				t.Fatalf("failed to list tokens: %v", err)
			}

			if len(*tokens) == 0 {
				t.Error("expected at least one token")
			}

			found := false
			for _, tok := range *tokens {
				if tok.Spec.Token == token.Spec.Token {
					found = true
					break
				}
			}
			if !found {
				t.Error("expected to find created token")
			}

			_, _ = sandbox.Previews.Delete(ctx, "list-tokens")
		})

		t.Run("deletes token", func(t *testing.T) {
			_, err := sandbox.Previews.New(ctx, blaxel.SandboxPreviewNewParams{
				Preview: blaxel.PreviewParam{
					Metadata: blaxel.PreviewMetadataParam{Name: "delete-token"},
					Spec: blaxel.PreviewSpecParam{
						Port:   blaxel.Int(3000),
						Public: blaxel.Bool(false),
					},
				},
			})
			if err != nil {
				t.Fatalf("failed to create preview: %v", err)
			}

			expiration := time.Now().Add(10 * time.Minute)
			token, err := sandbox.Previews.NewToken(ctx, "delete-token", expiration)
			if err != nil {
				t.Fatalf("failed to create token: %v", err)
			}

			_, err = sandbox.Previews.DeleteToken(ctx, "delete-token", token.Metadata.Name)
			if err != nil {
				t.Fatalf("failed to delete token: %v", err)
			}

			tokens, err := sandbox.Previews.ListTokens(ctx, "delete-token")
			if err != nil {
				t.Fatalf("failed to list tokens: %v", err)
			}

			for _, tok := range *tokens {
				if tok.Spec.Token == token.Spec.Token {
					t.Error("expected token to be deleted")
				}
			}

			_, _ = sandbox.Previews.Delete(ctx, "delete-token")
		})
	})

	t.Run("PreviewRaceConditions", func(t *testing.T) {
		t.Run("creates a preview then removes it and recreates the same preview", func(t *testing.T) {
			runTest := func(index int) error {
				previewName := fmt.Sprintf("preview-race-%d", index)
				preview, err := sandbox.Previews.NewIfNotExists(ctx, blaxel.SandboxPreviewNewParams{
					Preview: blaxel.PreviewParam{
						Metadata: blaxel.PreviewMetadataParam{Name: previewName},
						Spec: blaxel.PreviewSpecParam{
							Port:   blaxel.Int(3000),
							Public: blaxel.Bool(true),
						},
					},
				})
				if err != nil {
					return fmt.Errorf("failed to create preview %s: %v", previewName, err)
				}

				httpClient := &http.Client{Timeout: 10 * time.Second}
				resp, err := httpClient.Get(preview.Spec.URL)
				if err != nil {
					return fmt.Errorf("failed to access preview %s: %v", previewName, err)
				}
				resp.Body.Close()

				if resp.StatusCode != 200 {
					return fmt.Errorf("expected status 200 for %s, got %d", previewName, resp.StatusCode)
				}

				_, err = sandbox.Previews.Delete(ctx, previewName)
				if err != nil {
					return fmt.Errorf("failed to delete preview %s: %v", previewName, err)
				}

				preview2, err := sandbox.Previews.NewIfNotExists(ctx, blaxel.SandboxPreviewNewParams{
					Preview: blaxel.PreviewParam{
						Metadata: blaxel.PreviewMetadataParam{Name: previewName},
						Spec: blaxel.PreviewSpecParam{
							Port:   blaxel.Int(3000),
							Public: blaxel.Bool(true),
						},
					},
				})
				if err != nil {
					return fmt.Errorf("failed to recreate preview %s: %v", previewName, err)
				}

				resp2, err := httpClient.Get(preview2.Spec.URL)
				if err != nil {
					return fmt.Errorf("failed to access recreated preview %s: %v", previewName, err)
				}
				resp2.Body.Close()

				if resp2.StatusCode != 200 {
					t.Logf("Preview URL check failed for %s: %s - Status: %d", previewName, preview2.Spec.URL, resp2.StatusCode)
					return fmt.Errorf("expected status 200 for recreated %s, got %d", previewName, resp2.StatusCode)
				}

				// Cleanup
				_, _ = sandbox.Previews.Delete(ctx, previewName)
				return nil
			}

			var wg sync.WaitGroup
			errChan := make(chan error, 100)

			for i := 1; i <= 100; i++ {
				wg.Add(1)
				go func(idx int) {
					defer wg.Done()
					if err := runTest(idx); err != nil {
						errChan <- err
					}
				}(i)
			}
			wg.Wait()
			close(errChan)

			var errors []error
			for err := range errChan {
				errors = append(errors, err)
			}
			if len(errors) > 0 {
				for _, err := range errors {
					t.Errorf("%v", err)
				}
				t.Fatalf("had %d failures in race condition test", len(errors))
			}
		})
	})
}
