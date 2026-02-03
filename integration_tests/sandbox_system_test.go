package integration_tests

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"testing"
	"time"

	blaxel "github.com/blaxel-ai/sdk-go"
)

// waitForUpgradeComplete waits for sandbox upgrade to complete by polling the health endpoint
// using the sandbox's System.Health method which includes proper authentication
func waitForUpgradeComplete(ctx context.Context, sandbox *blaxel.SandboxInstance, maxWaitTime time.Duration) (*blaxel.HealthResponse, error) {
	fmt.Println("[TEST] Waiting for health upgrade count > 0...")
	startTime := time.Now()
	pollInterval := 500 * time.Millisecond
	var healthData *blaxel.HealthResponse

	for time.Since(startTime) < maxWaitTime {
		select {
		case <-ctx.Done():
			return nil, ctx.Err()
		default:
		}

		health, err := sandbox.System.Health(ctx)
		if err != nil {
			fmt.Printf("[TEST] Health check error: %v (elapsed: %dms)\n", err, time.Since(startTime).Milliseconds())
			time.Sleep(pollInterval)
			continue
		}

		healthData = health
		fmt.Printf("[TEST] Health check - upgradeCount: %d (elapsed: %dms)\n", healthData.UpgradeCount, time.Since(startTime).Milliseconds())

		if healthData.UpgradeCount > 0 {
			fmt.Printf("[TEST] Upgrade completed (took %dms)\n", time.Since(startTime).Milliseconds())
			return healthData, nil
		}

		if healthData.LastUpgrade.Status == "failed" {
			fmt.Printf("[TEST] Health check - last upgrade failed, health data: %+v\n", healthData)
			return nil, fmt.Errorf("upgrade failed: %+v", healthData)
		}

		time.Sleep(pollInterval)
	}

	return nil, fmt.Errorf("upgrade did not complete within %v. Last health data: %+v", maxWaitTime, healthData)
}

func TestSandboxSystem(t *testing.T) {
	ctx := context.Background()
	client := newTestClient(t)

	var createdSandboxes []string

	t.Cleanup(func() {
		for _, name := range createdSandboxes {
			_, _ = client.Sandboxes.Delete(ctx, name)
		}
	})

	t.Run("Upgrade", func(t *testing.T) {
		t.Run("upgrades sandbox and preview remains responsive", func(t *testing.T) {
			name := uniqueName("system-upgrade")
			t.Logf("[TEST] Starting test with sandbox name: %s", name)

			// Create sandbox with Next.js image
			t.Log("[TEST] Creating sandbox with blaxel/nextjs:latest image...")
			sandbox, err := client.Sandboxes.NewInstance(ctx, blaxel.SandboxNewParams{
				Sandbox: blaxel.SandboxParam{
					Metadata: blaxel.MetadataParam{
						Name:   name,
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
			createdSandboxes = append(createdSandboxes, name)
			t.Log("[TEST] Sandbox created")

			sandboxHost := sandbox.Metadata.URL
			t.Logf("[TEST] Sandbox host: %s", sandboxHost)
			if sandboxHost == "" {
				t.Fatal("sandbox host URL is empty")
			}

			// Do initial health check using the SDK
			t.Log("[TEST] Performing initial health check...")
			initialHealth, err := sandbox.System.Health(ctx)
			if err != nil {
				t.Logf("[TEST] Initial health check error (may be expected): %v", err)
			} else {
				t.Logf("[TEST] Initial health check - upgradeCount: %d", initialHealth.UpgradeCount)
			}

			httpClient := &http.Client{Timeout: 10 * time.Second}

			// Start the Next.js dev server
			t.Log("[TEST] Starting Next.js dev server...")
			_, err = sandbox.Process.New(ctx, blaxel.ProcessRequestParam{
				Name:         blaxel.Opt("nextjs-dev"),
				Command:      "npm run dev -- --port 3000",
				WorkingDir:   blaxel.String("/blaxel/app"),
				WaitForPorts: []int64{3000},
			})
			if err != nil {
				t.Fatalf("failed to start dev server: %v", err)
			}
			t.Log("[TEST] Next.js dev server started")

			// Create a public preview on port 3000
			t.Log("[TEST] Creating preview on port 3000...")
			preview, err := sandbox.Previews.New(ctx, blaxel.SandboxPreviewNewParams{
				Preview: blaxel.PreviewParam{
					Metadata: blaxel.PreviewMetadataParam{Name: "upgrade-test-preview"},
					Spec: blaxel.PreviewSpecParam{
						Port:   blaxel.Int(3000),
						Public: blaxel.Bool(true),
					},
				},
			})
			if err != nil {
				t.Fatalf("failed to create preview: %v", err)
			}

			previewURL := preview.Spec.URL
			if previewURL == "" {
				t.Fatal("preview URL is empty")
			}
			t.Logf("[TEST] Preview created with URL: %s", previewURL)

			// Wait for preview to be ready and verify it's responsive
			t.Log("[TEST] Waiting for preview to be ready...")
			previewReady := false
			maxWaitTime := 30 * time.Second
			startTime := time.Now()

			for time.Since(startTime) < maxWaitTime {
				resp, err := httpClient.Get(previewURL)
				if err != nil {
					t.Logf("[TEST] Preview check - error: %v (elapsed: %dms)", err, time.Since(startTime).Milliseconds())
				} else {
					t.Logf("[TEST] Preview check - status: %d (elapsed: %dms)", resp.StatusCode, time.Since(startTime).Milliseconds())
					resp.Body.Close()
					if resp.StatusCode == 200 {
						previewReady = true
						break
					}
				}
				time.Sleep(1 * time.Second)
			}

			t.Logf("[TEST] Preview ready: %v (took %dms)", previewReady, time.Since(startTime).Milliseconds())
			if !previewReady {
				t.Fatal("preview did not become ready")
			}

			// Verify preview is accessible before upgrade and capture content
			t.Log("[TEST] Verifying preview is accessible before upgrade...")
			preUpgradeResp, err := httpClient.Get(previewURL)
			if err != nil {
				t.Fatalf("failed to access preview before upgrade: %v", err)
			}
			preUpgradeContent, _ := io.ReadAll(preUpgradeResp.Body)
			preUpgradeResp.Body.Close()
			t.Logf("[TEST] Pre-upgrade preview status: %d", preUpgradeResp.StatusCode)
			t.Logf("[TEST] Pre-upgrade preview content length: %d bytes", len(preUpgradeContent))
			if preUpgradeResp.StatusCode != 200 {
				t.Fatalf("pre-upgrade preview returned status %d", preUpgradeResp.StatusCode)
			}

			// Upgrade the sandbox system
			t.Logf("[TEST] Calling sandbox.System.Upgrade()...")
			upgradeResult, err := sandbox.System.Upgrade(ctx, blaxel.SandboxSystemUpgradeParams{})
			if err != nil {
				t.Fatalf("failed to upgrade sandbox: %v", err)
			}
			t.Logf("[TEST] Upgrade call completed, result: %+v", upgradeResult)

			// Wait for health to show upgrade count > 0
			healthData, err := waitForUpgradeComplete(ctx, sandbox, maxWaitTime)
			if err != nil {
				t.Fatalf("upgrade did not complete: %v", err)
			}
			if healthData.UpgradeCount <= 0 {
				t.Fatalf("expected upgradeCount > 0, got %d", healthData.UpgradeCount)
			}

			// Wait a bit for everything to stabilize after upgrade
			t.Log("[TEST] Waiting 5s for stabilization...")
			time.Sleep(5 * time.Second)

			// Verify preview URL is still responsive after upgrade
			t.Log("[TEST] Verifying preview is still responsive after upgrade...")
			postUpgradeResp, err := httpClient.Get(previewURL)
			if err != nil {
				t.Fatalf("failed to access preview after upgrade: %v", err)
			}
			postUpgradeContent, _ := io.ReadAll(postUpgradeResp.Body)
			postUpgradeResp.Body.Close()
			t.Logf("[TEST] Post-upgrade preview status: %d", postUpgradeResp.StatusCode)
			t.Logf("[TEST] Post-upgrade preview content length: %d bytes", len(postUpgradeContent))
			if postUpgradeResp.StatusCode != 200 {
				t.Fatalf("post-upgrade preview returned status %d", postUpgradeResp.StatusCode)
			}

			// Verify the content size is similar before and after upgrade (allow delta of 200 bytes)
			sizeDelta := len(postUpgradeContent) - len(preUpgradeContent)
			if sizeDelta < 0 {
				sizeDelta = -sizeDelta
			}
			t.Logf("[TEST] Comparing content sizes - pre: %d, post: %d, delta: %d", len(preUpgradeContent), len(postUpgradeContent), sizeDelta)
			if sizeDelta > 200 {
				t.Fatalf("content size difference too large: %d bytes", sizeDelta)
			}

			t.Log("[TEST] Test completed successfully!")
		})

		t.Run("upgrades sandbox and running process persists", func(t *testing.T) {
			name := uniqueName("system-upgrade-process")
			t.Logf("[TEST] Starting process persistence test with sandbox name: %s", name)

			// Create sandbox
			t.Log("[TEST] Creating sandbox...")
			sandbox, err := client.Sandboxes.NewInstance(ctx, blaxel.SandboxNewParams{
				Sandbox: blaxel.SandboxParam{
					Metadata: blaxel.MetadataParam{
						Name:   name,
						Labels: defaultLabels,
					},
					Spec: blaxel.SandboxSpecParam{
						Region: blaxel.String(defaultRegion),
						Runtime: blaxel.SandboxRuntimeParam{
							Image:  blaxel.String("blaxel/base-image:latest"),
							Memory: blaxel.Int(1024),
						},
					},
				},
			})
			if err != nil {
				t.Fatalf("failed to create sandbox: %v", err)
			}
			createdSandboxes = append(createdSandboxes, name)
			t.Log("[TEST] Sandbox created")

			// Start a sleep process that will run for 6 seconds
			sleepDuration := 6
			t.Logf("[TEST] Starting sleep process for %d seconds...", sleepDuration)
			processStart := time.Now()
			sleepProcess, err := sandbox.Process.New(ctx, blaxel.ProcessRequestParam{
				Name:              blaxel.Opt("test-sleep"),
				Command:           fmt.Sprintf("sleep %d", sleepDuration),
				WaitForCompletion: blaxel.Bool(false),
			})
			if err != nil {
				t.Fatalf("failed to start sleep process: %v", err)
			}
			t.Logf("[TEST] Sleep process started with name: %s", sleepProcess.Name)
			if sleepProcess.Name != "test-sleep" {
				t.Fatalf("expected process name 'test-sleep', got %s", sleepProcess.Name)
			}

			// Wait a bit to ensure the process is running
			time.Sleep(2 * time.Second)

			// Verify the process is running before upgrade
			t.Log("[TEST] Checking process status before upgrade...")
			processBeforeUpgrade, err := sandbox.Process.Get(ctx, "test-sleep")
			if err != nil {
				t.Fatalf("failed to get process: %v", err)
			}
			t.Logf("[TEST] Process status before upgrade: %s", processBeforeUpgrade.Status)
			if processBeforeUpgrade.Status != "running" {
				t.Fatalf("expected process status 'running', got %s", processBeforeUpgrade.Status)
			}

			// Upgrade the sandbox system
			t.Logf("[TEST] Calling sandbox.System.Upgrade()...")
			upgradeResult, err := sandbox.System.Upgrade(ctx, blaxel.SandboxSystemUpgradeParams{})
			if err != nil {
				t.Fatalf("failed to upgrade sandbox: %v", err)
			}
			t.Logf("[TEST] Upgrade call completed, result: %+v", upgradeResult)

			// Wait for the upgrade to complete (check health)
			healthData, err := waitForUpgradeComplete(ctx, sandbox, 10*time.Second)
			if err != nil {
				t.Fatalf("upgrade did not complete: %v", err)
			}
			if healthData.UpgradeCount <= 0 {
				t.Fatalf("expected upgradeCount > 0, got %d", healthData.UpgradeCount)
			}

			// Check that the sleep process is still visible in the API after upgrade
			t.Log("[TEST] Checking process status after upgrade...")
			processAfterUpgrade, err := sandbox.Process.Get(ctx, "test-sleep")
			if err != nil {
				t.Fatalf("failed to get process after upgrade: %v", err)
			}
			t.Logf("[TEST] Process status after upgrade: %s", processAfterUpgrade.Status)
			// The process should still be running (or completed if we took too long)
			if processAfterUpgrade.Status != "running" && processAfterUpgrade.Status != "completed" {
				t.Fatalf("expected process status 'running' or 'completed', got %s", processAfterUpgrade.Status)
			}

			// Calculate remaining time for the sleep to complete
			elapsedSinceStart := time.Since(processStart)
			expectedTotalDuration := time.Duration(sleepDuration) * time.Second
			remainingTime := expectedTotalDuration - elapsedSinceStart
			t.Logf("[TEST] Elapsed since process start: %dms, remaining: %dms", elapsedSinceStart.Milliseconds(), remainingTime.Milliseconds())

			// If the process is still running, wait for it to complete
			if processAfterUpgrade.Status == "running" {
				// Wait for the process to complete with some buffer (5 seconds extra)
				waitTime := remainingTime + 5*time.Second
				if waitTime < 5*time.Second {
					waitTime = 5 * time.Second
				}
				t.Logf("[TEST] Waiting %dms for process to complete...", waitTime.Milliseconds())

				completedProcess, err := sandbox.Process.Wait(ctx, "test-sleep", waitTime, 1*time.Second)
				if err != nil {
					t.Fatalf("failed to wait for process: %v", err)
				}
				t.Logf("[TEST] Process completed with status: %s, exitCode: %d", completedProcess.Status, completedProcess.ExitCode)
				if completedProcess.Status != "completed" {
					t.Fatalf("expected process status 'completed', got %s", completedProcess.Status)
				}
				if completedProcess.ExitCode != 0 {
					t.Fatalf("expected exit code 0, got %d", completedProcess.ExitCode)
				}
			}

			// Verify the process completed in roughly the expected time (within 15 seconds tolerance)
			totalDuration := time.Since(processStart)
			t.Logf("[TEST] Total duration from process start to completion: %dms", totalDuration.Milliseconds())
			t.Logf("[TEST] Expected duration: ~%dms", expectedTotalDuration.Milliseconds())

			// The process should have completed close to the expected time
			// Allow 15 seconds tolerance for upgrade overhead
			tolerance := 15 * time.Second
			if totalDuration < expectedTotalDuration-2*time.Second {
				t.Fatalf("process completed too early: %v < %v", totalDuration, expectedTotalDuration-2*time.Second)
			}
			if totalDuration > expectedTotalDuration+tolerance {
				t.Fatalf("process took too long: %v > %v", totalDuration, expectedTotalDuration+tolerance)
			}

			t.Log("[TEST] Process persistence test completed successfully!")
		})
	})
}
