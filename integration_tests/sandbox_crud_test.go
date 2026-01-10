package integration_tests

import (
	"context"
	"strings"
	"sync"
	"testing"
	"time"

	blaxel "github.com/stainless-sdks/blaxel-go"
)

func TestSandboxCRUD(t *testing.T) {
	ctx := context.Background()
	client := newTestClient(t)
	var createdSandboxes []string
	var mu sync.Mutex

	trackSandbox := func(name string) {
		mu.Lock()
		createdSandboxes = append(createdSandboxes, name)
		mu.Unlock()
	}

	t.Cleanup(func() {
		for _, name := range createdSandboxes {
			_, _ = client.Sandboxes.Delete(ctx, name)
		}
	})

	t.Run("Create", func(t *testing.T) {
		t.Run("creates a sandbox with default settings", func(t *testing.T) {
			sandbox, err := client.Sandboxes.NewInstance(ctx, blaxel.SandboxNewParams{
				Sandbox: blaxel.SandboxParam{
					Metadata: blaxel.MetadataParam{
						Labels: defaultLabels,
					},
				},
			})
			if err != nil {
				t.Fatalf("failed to create sandbox: %v", err)
			}
			if sandbox.Metadata.Name == "" {
				t.Fatal("expected sandbox name to be set")
			}
			trackSandbox(sandbox.Metadata.Name)
			if !strings.HasPrefix(sandbox.Metadata.Name, "sandbox-") {
				t.Errorf("expected name to start with 'sandbox-', got %s", sandbox.Metadata.Name)
			}
		})

		t.Run("creates a sandbox with custom name", func(t *testing.T) {
			name := uniqueName("custom")
			sandbox, err := client.Sandboxes.NewInstance(ctx, blaxel.SandboxNewParams{
				Sandbox: blaxel.SandboxParam{
					Metadata: blaxel.MetadataParam{
						Name:   name,
						Labels: defaultLabels,
					},
				},
			})
			if err != nil {
				t.Fatalf("failed to create sandbox: %v", err)
			}
			trackSandbox(name)
			if sandbox.Metadata.Name != name {
				t.Errorf("expected name %s, got %s", name, sandbox.Metadata.Name)
			}
		})

		t.Run("creates a sandbox with specific image", func(t *testing.T) {
			name := uniqueName("image-test")
			sandbox, err := client.Sandboxes.NewInstance(ctx, blaxel.SandboxNewParams{
				Sandbox: blaxel.SandboxParam{
					Metadata: blaxel.MetadataParam{
						Name:   name,
						Labels: defaultLabels,
					},
					Spec: blaxel.SandboxSpecParam{
						Runtime: blaxel.SandboxRuntimeParam{
							Image: blaxel.String(defaultImage),
						},
					},
				},
			})
			if err != nil {
				t.Fatalf("failed to create sandbox: %v", err)
			}
			trackSandbox(name)
			if sandbox.Metadata.Name != name {
				t.Errorf("expected name %s, got %s", name, sandbox.Metadata.Name)
			}
		})

		t.Run("creates a sandbox with memory configuration", func(t *testing.T) {
			name := uniqueName("memory-test")
			_, err := client.Sandboxes.NewInstance(ctx, blaxel.SandboxNewParams{
				Sandbox: blaxel.SandboxParam{
					Metadata: blaxel.MetadataParam{
						Name:   name,
						Labels: defaultLabels,
					},
					Spec: blaxel.SandboxSpecParam{
						Runtime: blaxel.SandboxRuntimeParam{
							Image:  blaxel.String(defaultImage),
							Memory: blaxel.Int(8192),
						},
					},
				},
			})
			if err != nil {
				t.Fatalf("failed to create sandbox: %v", err)
			}
			trackSandbox(name)

			retrieved, err := client.Sandboxes.Get(ctx, name, blaxel.SandboxGetParams{})
			if err != nil {
				t.Fatalf("failed to get sandbox: %v", err)
			}
			if retrieved.Spec.Runtime.Memory != 8192 {
				t.Errorf("expected memory 8192, got %d", retrieved.Spec.Runtime.Memory)
			}
		})

		t.Run("creates a sandbox with labels", func(t *testing.T) {
			name := uniqueName("labels-test")
			labels := map[string]string{
				"env":        "test",
				"purpose":    "integration",
				"created-by": "go-integration-tests",
			}
			sandbox, err := client.Sandboxes.NewInstance(ctx, blaxel.SandboxNewParams{
				Sandbox: blaxel.SandboxParam{
					Metadata: blaxel.MetadataParam{
						Name:   name,
						Labels: labels,
					},
				},
			})
			if err != nil {
				t.Fatalf("failed to create sandbox: %v", err)
			}
			trackSandbox(name)
			if sandbox.Metadata.Labels["env"] != "test" {
				t.Errorf("expected label env=test, got %s", sandbox.Metadata.Labels["env"])
			}
			if sandbox.Metadata.Labels["purpose"] != "integration" {
				t.Errorf("expected label purpose=integration, got %s", sandbox.Metadata.Labels["purpose"])
			}
		})

		t.Run("creates a sandbox with ports", func(t *testing.T) {
			name := uniqueName("ports-test")
			_, err := client.Sandboxes.NewInstance(ctx, blaxel.SandboxNewParams{
				Sandbox: blaxel.SandboxParam{
					Metadata: blaxel.MetadataParam{
						Name:   name,
						Labels: defaultLabels,
					},
					Spec: blaxel.SandboxSpecParam{
						Runtime: blaxel.SandboxRuntimeParam{
							Image:  blaxel.String(defaultImage),
							Memory: blaxel.Int(2048),
							Ports: []blaxel.PortParam{
								{Name: blaxel.String("web"), Target: 3000},
								{Name: blaxel.String("api"), Target: 8080, Protocol: blaxel.PortProtocolTcp},
							},
						},
					},
				},
			})
			if err != nil {
				t.Fatalf("failed to create sandbox: %v", err)
			}
			trackSandbox(name)

			retrieved, err := client.Sandboxes.Get(ctx, name, blaxel.SandboxGetParams{})
			if err != nil {
				t.Fatalf("failed to get sandbox: %v", err)
			}
			if len(retrieved.Spec.Runtime.Ports) != 2 {
				t.Errorf("expected 2 ports, got %d", len(retrieved.Spec.Runtime.Ports))
			}
		})

		t.Run("creates a sandbox with environment variables", func(t *testing.T) {
			name := uniqueName("envs-test")
			_, err := client.Sandboxes.NewInstance(ctx, blaxel.SandboxNewParams{
				Sandbox: blaxel.SandboxParam{
					Metadata: blaxel.MetadataParam{
						Name:   name,
						Labels: defaultLabels,
					},
					Spec: blaxel.SandboxSpecParam{
						Runtime: blaxel.SandboxRuntimeParam{
							Image: blaxel.String(defaultImage),
							Envs: []blaxel.SandboxRuntimeEnvParam{
								{Name: blaxel.String("NODE_ENV"), Value: blaxel.String("test")},
								{Name: blaxel.String("DEBUG"), Value: blaxel.String("true")},
							},
						},
					},
				},
			})
			if err != nil {
				t.Fatalf("failed to create sandbox: %v", err)
			}
			trackSandbox(name)

			retrieved, err := client.Sandboxes.Get(ctx, name, blaxel.SandboxGetParams{})
			if err != nil {
				t.Fatalf("failed to get sandbox: %v", err)
			}
			if len(retrieved.Spec.Runtime.Envs) != 2 {
				t.Errorf("expected 2 envs, got %d", len(retrieved.Spec.Runtime.Envs))
			}
		})

		t.Run("creates a sandbox with region", func(t *testing.T) {
			name := uniqueName("region-test")
			_, err := client.Sandboxes.NewInstance(ctx, blaxel.SandboxNewParams{
				Sandbox: blaxel.SandboxParam{
					Metadata: blaxel.MetadataParam{
						Name:   name,
						Labels: defaultLabels,
					},
					Spec: blaxel.SandboxSpecParam{
						Region: blaxel.String(defaultRegion),
						Runtime: blaxel.SandboxRuntimeParam{
							Image: blaxel.String(defaultImage),
						},
					},
				},
			})
			if err != nil {
				t.Fatalf("failed to create sandbox: %v", err)
			}
			trackSandbox(name)

			retrieved, err := client.Sandboxes.Get(ctx, name, blaxel.SandboxGetParams{})
			if err != nil {
				t.Fatalf("failed to get sandbox: %v", err)
			}
			if retrieved.Spec.Region != defaultRegion {
				t.Errorf("expected region %s, got %s", defaultRegion, retrieved.Spec.Region)
			}
		})
	})

	t.Run("CreateIfNotExists", func(t *testing.T) {
		t.Run("creates a new sandbox if it does not exist", func(t *testing.T) {
			name := uniqueName("cine")
			sandbox, err := client.Sandboxes.CreateInstanceIfNotExists(ctx, blaxel.SandboxNewParams{
				Sandbox: blaxel.SandboxParam{
					Metadata: blaxel.MetadataParam{
						Name:   name,
						Labels: defaultLabels,
					},
				},
			})
			if err != nil {
				t.Fatalf("failed to create sandbox: %v", err)
			}
			trackSandbox(name)
			if sandbox.Metadata.Name != name {
				t.Errorf("expected name %s, got %s", name, sandbox.Metadata.Name)
			}
		})

		t.Run("returns existing sandbox if it already exists", func(t *testing.T) {
			name := uniqueName("cine-existing")

			// Create first
			first, err := client.Sandboxes.NewInstance(ctx, blaxel.SandboxNewParams{
				Sandbox: blaxel.SandboxParam{
					Metadata: blaxel.MetadataParam{
						Name:   name,
						Labels: defaultLabels,
					},
				},
			})
			if err != nil {
				t.Fatalf("failed to create sandbox: %v", err)
			}
			trackSandbox(name)

			// createIfNotExists should return the same sandbox
			second, err := client.Sandboxes.CreateInstanceIfNotExists(ctx, blaxel.SandboxNewParams{
				Sandbox: blaxel.SandboxParam{
					Metadata: blaxel.MetadataParam{
						Name:   name,
						Labels: defaultLabels,
					},
				},
			})
			if err != nil {
				t.Fatalf("failed to createIfNotExists: %v", err)
			}
			if second.Metadata.Name != first.Metadata.Name {
				t.Errorf("expected same sandbox, got different names: %s vs %s", first.Metadata.Name, second.Metadata.Name)
			}
		})

		t.Run("handles concurrent createIfNotExists calls", func(t *testing.T) {
			name := uniqueName("cine-race")
			concurrentCalls := 5
			results := make(chan *blaxel.SandboxInstance, concurrentCalls)
			errors := make(chan error, concurrentCalls)

			var wg sync.WaitGroup
			for i := 0; i < concurrentCalls; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					sandbox, err := client.Sandboxes.CreateInstanceIfNotExists(ctx, blaxel.SandboxNewParams{
						Sandbox: blaxel.SandboxParam{
							Metadata: blaxel.MetadataParam{
								Name:   name,
								Labels: defaultLabels,
							},
						},
					})
					if err != nil {
						errors <- err
					} else {
						results <- sandbox
					}
				}()
			}
			wg.Wait()
			close(results)
			close(errors)
			trackSandbox(name)

			var sandboxes []*blaxel.SandboxInstance
			for sb := range results {
				sandboxes = append(sandboxes, sb)
			}

			// All should have the same name
			names := make(map[string]bool)
			for _, sb := range sandboxes {
				names[sb.Metadata.Name] = true
			}
			if len(names) != 1 {
				t.Errorf("expected all sandboxes to have the same name, got %d unique names", len(names))
			}
			if len(sandboxes) < 3 {
				t.Errorf("expected at least 3 successes, got %d", len(sandboxes))
			}

			// Verify sandbox works
			retrieved, err := client.Sandboxes.GetInstance(ctx, name)
			if err != nil {
				t.Fatalf("failed to get sandbox: %v", err)
			}
			result, err := retrieved.Process.New(ctx, blaxel.ProcessRequestParam{
				Command:           "echo 'test'",
				WaitForCompletion: blaxel.Bool(true),
			})
			if err != nil {
				t.Fatalf("failed to exec: %v", err)
			}
			if !strings.Contains(result.Logs, "test") {
				t.Errorf("expected logs to contain 'test', got %s", result.Logs)
			}
		})
	})

	t.Run("Get", func(t *testing.T) {
		t.Run("retrieves an existing sandbox", func(t *testing.T) {
			name := uniqueName("get-test")
			_, err := client.Sandboxes.NewInstance(ctx, blaxel.SandboxNewParams{
				Sandbox: blaxel.SandboxParam{
					Metadata: blaxel.MetadataParam{
						Name:   name,
						Labels: defaultLabels,
					},
				},
			})
			if err != nil {
				t.Fatalf("failed to create sandbox: %v", err)
			}
			trackSandbox(name)

			retrieved, err := client.Sandboxes.GetInstance(ctx, name)
			if err != nil {
				t.Fatalf("failed to get sandbox: %v", err)
			}
			if retrieved.Metadata.Name != name {
				t.Errorf("expected name %s, got %s", name, retrieved.Metadata.Name)
			}
		})

		t.Run("returns error for non-existent sandbox", func(t *testing.T) {
			_, err := client.Sandboxes.GetInstance(ctx, "non-existent-sandbox-xyz")
			if err == nil {
				t.Error("expected error for non-existent sandbox")
			}
		})
	})

	t.Run("List", func(t *testing.T) {
		t.Run("lists all sandboxes", func(t *testing.T) {
			name := uniqueName("list-test")
			_, err := client.Sandboxes.NewInstance(ctx, blaxel.SandboxNewParams{
				Sandbox: blaxel.SandboxParam{
					Metadata: blaxel.MetadataParam{
						Name:   name,
						Labels: defaultLabels,
					},
				},
			})
			if err != nil {
				t.Fatalf("failed to create sandbox: %v", err)
			}
			trackSandbox(name)

			sandboxes, err := client.Sandboxes.List(ctx)
			if err != nil {
				t.Fatalf("failed to list sandboxes: %v", err)
			}
			if sandboxes == nil || len(*sandboxes) == 0 {
				t.Error("expected at least one sandbox")
			}

			found := false
			for _, sb := range *sandboxes {
				if sb.Metadata.Name == name {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("expected to find sandbox %s in list", name)
			}
		})
	})

	t.Run("Delete", func(t *testing.T) {
		t.Run("deletes an existing sandbox", func(t *testing.T) {
			name := uniqueName("delete-test")
			_, err := client.Sandboxes.NewInstance(ctx, blaxel.SandboxNewParams{
				Sandbox: blaxel.SandboxParam{
					Metadata: blaxel.MetadataParam{
						Name:   name,
						Labels: defaultLabels,
					},
				},
			})
			if err != nil {
				t.Fatalf("failed to create sandbox: %v", err)
			}

			_, err = client.Sandboxes.Delete(ctx, name)
			if err != nil {
				t.Fatalf("failed to delete sandbox: %v", err)
			}

			// Wait for deletion to complete
			deleted := waitForSandboxDeletion(ctx, client, name, 15)
			if !deleted {
				t.Error("sandbox was not fully deleted")
			}
		})

		t.Run("can delete using instance method", func(t *testing.T) {
			name := uniqueName("delete-instance")
			sandbox, err := client.Sandboxes.NewInstance(ctx, blaxel.SandboxNewParams{
				Sandbox: blaxel.SandboxParam{
					Metadata: blaxel.MetadataParam{
						Name:   name,
						Labels: defaultLabels,
					},
				},
			})
			if err != nil {
				t.Fatalf("failed to create sandbox: %v", err)
			}

			err = sandbox.Delete(ctx)
			if err != nil {
				t.Fatalf("failed to delete sandbox: %v", err)
			}

			// Wait for deletion to complete
			deleted := waitForSandboxDeletion(ctx, client, name, 15)
			if !deleted {
				t.Error("sandbox was not fully deleted")
			}
		})
	})

	t.Run("UpdateInstanceMetadata", func(t *testing.T) {
		t.Run("updates sandbox labels", func(t *testing.T) {
			name := uniqueName("update-meta")
			_, err := client.Sandboxes.NewInstance(ctx, blaxel.SandboxNewParams{
				Sandbox: blaxel.SandboxParam{
					Metadata: blaxel.MetadataParam{
						Name:   name,
						Labels: defaultLabels,
					},
				},
			})
			if err != nil {
				t.Fatalf("failed to create sandbox: %v", err)
			}
			trackSandbox(name)

			newLabels := map[string]string{
				"env":        "integration-test",
				"language":   "go",
				"created-by": "go-integration-tests",
				"updated":    "true",
			}
			updated, err := client.Sandboxes.UpdateInstanceMetadata(ctx, name, blaxel.SandboxUpdateMetadataParams{
				Labels: newLabels,
			})
			if err != nil {
				t.Fatalf("failed to update sandbox: %v", err)
			}
			if updated.Metadata.Labels["updated"] != "true" {
				t.Errorf("expected label updated=true, got %s", updated.Metadata.Labels["updated"])
			}
		})

		t.Run("updates sandbox displayName", func(t *testing.T) {
			name := uniqueName("update-display")
			_, err := client.Sandboxes.NewInstance(ctx, blaxel.SandboxNewParams{
				Sandbox: blaxel.SandboxParam{
					Metadata: blaxel.MetadataParam{
						Name:   name,
						Labels: defaultLabels,
					},
				},
			})
			if err != nil {
				t.Fatalf("failed to create sandbox: %v", err)
			}
			trackSandbox(name)

			updated, err := client.Sandboxes.UpdateInstanceMetadata(ctx, name, blaxel.SandboxUpdateMetadataParams{
				DisplayName: "My Test Sandbox",
			})
			if err != nil {
				t.Fatalf("failed to update sandbox: %v", err)
			}
			if updated.Metadata.DisplayName != "My Test Sandbox" {
				t.Errorf("expected displayName 'My Test Sandbox', got %s", updated.Metadata.DisplayName)
			}
		})
	})

	t.Run("Wait", func(t *testing.T) {
		t.Run("waits for sandbox to be ready", func(t *testing.T) {
			name := uniqueName("wait-test")
			sandbox, err := client.Sandboxes.NewInstance(ctx, blaxel.SandboxNewParams{
				Sandbox: blaxel.SandboxParam{
					Metadata: blaxel.MetadataParam{
						Name:   name,
						Labels: defaultLabels,
					},
				},
			})
			if err != nil {
				t.Fatalf("failed to create sandbox: %v", err)
			}
			trackSandbox(name)

			// Wait is implicit in NewInstance, but let's verify we can run commands
			result, err := sandbox.Process.New(ctx, blaxel.ProcessRequestParam{
				Command:           "echo 'ready'",
				WaitForCompletion: blaxel.Bool(true),
			})
			if err != nil {
				t.Fatalf("failed to exec: %v", err)
			}
			if !strings.Contains(result.Logs, "ready") {
				t.Errorf("expected logs to contain 'ready', got %s", result.Logs)
			}
		})
	})
}

func TestSandboxLifecycle(t *testing.T) {
	ctx := context.Background()
	client := newTestClient(t)
	var createdSandboxes []string
	var mu sync.Mutex

	trackSandbox := func(name string) {
		mu.Lock()
		createdSandboxes = append(createdSandboxes, name)
		mu.Unlock()
	}

	t.Cleanup(func() {
		for _, name := range createdSandboxes {
			_, _ = client.Sandboxes.Delete(ctx, name)
		}
	})

	t.Run("TTL", func(t *testing.T) {
		t.Run("creates sandbox with TTL string", func(t *testing.T) {
			name := uniqueName("ttl-string")
			sandbox, err := client.Sandboxes.NewInstance(ctx, blaxel.SandboxNewParams{
				Sandbox: blaxel.SandboxParam{
					Metadata: blaxel.MetadataParam{
						Name:   name,
						Labels: defaultLabels,
					},
					Spec: blaxel.SandboxSpecParam{
						Runtime: blaxel.SandboxRuntimeParam{
							Image: blaxel.String(defaultImage),
							Ttl:   blaxel.String("5m"),
						},
					},
				},
			})
			if err != nil {
				t.Fatalf("failed to create sandbox: %v", err)
			}
			trackSandbox(name)
			if sandbox.Metadata.Name != name {
				t.Errorf("expected name %s, got %s", name, sandbox.Metadata.Name)
			}

			time.Sleep(100 * time.Millisecond)

			// Verify sandbox is running
			status, err := client.Sandboxes.Get(ctx, name, blaxel.SandboxGetParams{})
			if err != nil {
				t.Fatalf("failed to get sandbox: %v", err)
			}
			if status.Status == "TERMINATED" {
				t.Error("sandbox should not be terminated yet")
			}
		})
	})

	t.Run("ExpirationPolicies", func(t *testing.T) {
		t.Run("creates sandbox with ttl-max-age policy", func(t *testing.T) {
			name := uniqueName("maxage-policy")
			sandbox, err := client.Sandboxes.NewInstance(ctx, blaxel.SandboxNewParams{
				Sandbox: blaxel.SandboxParam{
					Metadata: blaxel.MetadataParam{
						Name:   name,
						Labels: defaultLabels,
					},
					Spec: blaxel.SandboxSpecParam{
						Runtime: blaxel.SandboxRuntimeParam{
							Image: blaxel.String(defaultImage),
						},
						Lifecycle: blaxel.SandboxLifecycleParam{
							ExpirationPolicies: []blaxel.ExpirationPolicyParam{
								{Type: blaxel.ExpirationPolicyTypeTtlMaxAge, Value: blaxel.String("10m"), Action: blaxel.ExpirationPolicyActionDelete},
							},
						},
					},
				},
			})
			if err != nil {
				t.Fatalf("failed to create sandbox: %v", err)
			}
			trackSandbox(name)
			if sandbox.Metadata.Name != name {
				t.Errorf("expected name %s, got %s", name, sandbox.Metadata.Name)
			}
		})

		t.Run("creates sandbox with ttl-idle policy", func(t *testing.T) {
			name := uniqueName("idle-policy")
			sandbox, err := client.Sandboxes.NewInstance(ctx, blaxel.SandboxNewParams{
				Sandbox: blaxel.SandboxParam{
					Metadata: blaxel.MetadataParam{
						Name:   name,
						Labels: defaultLabels,
					},
					Spec: blaxel.SandboxSpecParam{
						Runtime: blaxel.SandboxRuntimeParam{
							Image: blaxel.String(defaultImage),
						},
						Lifecycle: blaxel.SandboxLifecycleParam{
							ExpirationPolicies: []blaxel.ExpirationPolicyParam{
								{Type: blaxel.ExpirationPolicyTypeTtlIdle, Value: blaxel.String("5m"), Action: blaxel.ExpirationPolicyActionDelete},
							},
						},
					},
				},
			})
			if err != nil {
				t.Fatalf("failed to create sandbox: %v", err)
			}
			trackSandbox(name)
			if sandbox.Metadata.Name != name {
				t.Errorf("expected name %s, got %s", name, sandbox.Metadata.Name)
			}
		})

		t.Run("creates sandbox with multiple policies", func(t *testing.T) {
			name := uniqueName("multi-policy")
			sandbox, err := client.Sandboxes.NewInstance(ctx, blaxel.SandboxNewParams{
				Sandbox: blaxel.SandboxParam{
					Metadata: blaxel.MetadataParam{
						Name:   name,
						Labels: defaultLabels,
					},
					Spec: blaxel.SandboxSpecParam{
						Runtime: blaxel.SandboxRuntimeParam{
							Image: blaxel.String(defaultImage),
						},
						Lifecycle: blaxel.SandboxLifecycleParam{
							ExpirationPolicies: []blaxel.ExpirationPolicyParam{
								{Type: blaxel.ExpirationPolicyTypeTtlIdle, Value: blaxel.String("5m"), Action: blaxel.ExpirationPolicyActionDelete},
								{Type: blaxel.ExpirationPolicyTypeTtlMaxAge, Value: blaxel.String("30m"), Action: blaxel.ExpirationPolicyActionDelete},
							},
						},
					},
				},
			})
			if err != nil {
				t.Fatalf("failed to create sandbox: %v", err)
			}
			trackSandbox(name)
			if sandbox.Metadata.Name != name {
				t.Errorf("expected name %s, got %s", name, sandbox.Metadata.Name)
			}
		})
	})

	// Note: SnapshotConfiguration tests are skipped as this feature is not yet available in the Go SDK
}
