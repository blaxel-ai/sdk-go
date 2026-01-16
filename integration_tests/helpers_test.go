package integration_tests

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"os"
	"testing"
	"time"

	blaxel "github.com/blaxel-ai/sdk-go"
)

// Test configuration
var (
	defaultRegion = "us-pdx-1"
	defaultImage  = "blaxel/base-image:latest"
	defaultLabels = map[string]string{
		"env":        "integration-test",
		"language":   "go",
		"created-by": "go-integration-tests",
	}
)

func init() {
	if env := os.Getenv("BL_ENV"); env == "dev" {
		defaultRegion = "eu-dub-1"
	}
}

// uniqueName generates a unique name for testing
func uniqueName(prefix string) string {
	bytes := make([]byte, 4)
	rand.Read(bytes)
	return prefix + "-" + hex.EncodeToString(bytes)
}

// newTestClient creates a new Blaxel client for testing
func newTestClient(t *testing.T) blaxel.Client {
	t.Helper()
	client, err := blaxel.NewDefaultClient()
	if err != nil {
		t.Fatalf("failed to create client: %v", err)
		return blaxel.Client{}
	}
	return client
}

// waitForSandboxDeletion waits for a sandbox to be fully deleted
func waitForSandboxDeletion(ctx context.Context, client blaxel.Client, sandboxName string, maxAttempts int) bool {
	for i := 0; i < maxAttempts; i++ {
		select {
		case <-ctx.Done():
			return false
		default:
		}
		_, err := client.Sandboxes.Get(ctx, sandboxName, blaxel.SandboxGetParams{})
		if err != nil {
			// Sandbox no longer exists
			return true
		}
		time.Sleep(1 * time.Second)
	}
	return false
}

// waitForVolumeDeletion waits for a volume to be fully deleted
func waitForVolumeDeletion(ctx context.Context, client blaxel.Client, volumeName string, maxAttempts int) bool {
	for i := 0; i < maxAttempts; i++ {
		select {
		case <-ctx.Done():
			return false
		default:
		}
		_, err := client.Volumes.Get(ctx, volumeName)
		if err != nil {
			// Volume no longer exists
			return true
		}
		time.Sleep(1 * time.Second)
	}
	return false
}

// sleep is a helper to pause execution
func sleep(d time.Duration) {
	time.Sleep(d)
}

// cleanupSandboxes cleans up sandboxes in parallel
func cleanupSandboxes(ctx context.Context, client blaxel.Client, names []string) {
	for _, name := range names {
		go func(n string) {
			client.Sandboxes.Delete(ctx, n)
		}(name)
	}
}

// cleanupVolumes cleans up volumes in parallel
func cleanupVolumes(ctx context.Context, client blaxel.Client, names []string) {
	for _, name := range names {
		go func(n string) {
			client.Volumes.Delete(ctx, n)
		}(name)
	}
}
