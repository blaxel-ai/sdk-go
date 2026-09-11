package integration_tests

import (
	"context"
	"os"
	"testing"
	"time"

	blaxel "github.com/blaxel-ai/sdk-go"
)

// Snapshots are workspace-level objects: workspace routes address them by id,
// sandbox-nested routes by name (or id) within the source sandbox, and they
// outlive the sandbox they were captured from.
func TestWorkspaceSnapshots(t *testing.T) {
	ctx := context.Background()
	client := newTestClient(t)

	sandboxName := uniqueName("snap-src")
	snapshotName := uniqueName("snap")
	forkName := uniqueName("snap-fork")
	var snapshotID string

	t.Cleanup(func() {
		if snapshotID != "" {
			_ = client.Snapshots.Delete(ctx, snapshotID)
		}
		_, _ = client.Sandboxes.Delete(ctx, sandboxName)
		_, _ = client.Sandboxes.Delete(ctx, forkName)
	})

	sandbox, err := client.Sandboxes.NewInstance(ctx, blaxel.SandboxNewParams{
		Sandbox: blaxel.SandboxParam{
			Metadata: blaxel.MetadataParam{
				Name:   sandboxName,
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
	if sandbox.Metadata.Name != sandboxName {
		t.Fatalf("expected sandbox %s, got %s", sandboxName, sandbox.Metadata.Name)
	}

	t.Run("creates a named snapshot of a sandbox", func(t *testing.T) {
		snapshot, err := client.Sandboxes.Snapshots.New(ctx, sandboxName, blaxel.SandboxSnapshotNewParams{
			SandboxSnapshotRequest: blaxel.SandboxSnapshotRequestParam{
				Name: blaxel.String(snapshotName),
			},
		})
		if err != nil {
			t.Fatalf("failed to create snapshot: %v", err)
		}
		if snapshot.ID == "" {
			t.Fatal("expected snapshot id to be set")
		}
		snapshotID = snapshot.ID
		if snapshot.Name != snapshotName {
			t.Errorf("expected snapshot name %s, got %s", snapshotName, snapshot.Name)
		}
		if snapshot.Source.Kind != blaxel.SandboxSnapshotSourceKindSandbox {
			t.Errorf("expected source kind sandbox, got %s", snapshot.Source.Kind)
		}
		if snapshot.Source.Name != sandboxName {
			t.Errorf("expected source %s, got %s", sandboxName, snapshot.Source.Name)
		}
	})
	if snapshotID == "" {
		t.FailNow()
	}

	t.Run("lists the snapshot under its sandbox and in the workspace", func(t *testing.T) {
		nested, err := client.Sandboxes.Snapshots.List(ctx, sandboxName)
		if err != nil {
			t.Fatalf("failed to list sandbox snapshots: %v", err)
		}
		found := false
		for _, s := range *nested {
			if s.Name == snapshotName {
				found = true
			}
		}
		if !found {
			t.Errorf("snapshot %s not listed on sandbox %s", snapshotName, sandboxName)
		}

		found = false
		pager := client.Snapshots.ListAutoPaging(ctx, blaxel.SnapshotListParams{Limit: blaxel.Int(200)})
		for pager.Next() {
			if pager.Current().ID == snapshotID {
				found = true
				break
			}
		}
		if err := pager.Err(); err != nil {
			t.Fatalf("failed to list workspace snapshots: %v", err)
		}
		if !found {
			t.Errorf("snapshot %s not listed in the workspace", snapshotID)
		}
	})

	t.Run("becomes ready", func(t *testing.T) {
		deadline := time.Now().Add(90 * time.Second)
		for {
			snapshot, err := client.Snapshots.Get(ctx, snapshotID)
			if err != nil {
				t.Fatalf("failed to get snapshot %s: %v", snapshotID, err)
			}
			if snapshot.Status == "ready" {
				return
			}
			if snapshot.Status == "failed" {
				t.Fatalf("snapshot %s failed", snapshotID)
			}
			if time.Now().After(deadline) {
				t.Fatalf("snapshot %s still %s after 90s", snapshotID, snapshot.Status)
			}
			sleep(500 * time.Millisecond)
		}
	})

	t.Run("outlives the sandbox it was captured from", func(t *testing.T) {
		if _, err := client.Sandboxes.Delete(ctx, sandboxName); err != nil {
			t.Fatalf("failed to delete sandbox: %v", err)
		}
		if !waitForSandboxDeletion(ctx, client, sandboxName, 60) {
			t.Fatalf("sandbox %s was not deleted", sandboxName)
		}

		orphan, err := client.Snapshots.Get(ctx, snapshotID)
		if err != nil {
			t.Fatalf("snapshot %s vanished with its sandbox: %v", snapshotID, err)
		}
		if orphan.Name != snapshotName {
			t.Errorf("expected snapshot name %s, got %s", snapshotName, orphan.Name)
		}
		if orphan.Source.Name != sandboxName {
			t.Errorf("expected source %s, got %s", sandboxName, orphan.Source.Name)
		}
		// What a fork needs to run is on the snapshot itself, not on the source.
		if orphan.Spec.Image == "" {
			t.Error("expected snapshot spec to carry the image")
		}
	})

	// A fork is a full sandbox start, past the one-minute budget of the default run.
	t.Run("forks a sandbox from a snapshot whose source is gone", func(t *testing.T) {
		if os.Getenv("RUN_SLOW_SNAPSHOT_FORK") == "" {
			t.Skip("slow test; set RUN_SLOW_SNAPSHOT_FORK=1 to enable")
		}
		fork, err := client.Snapshots.Fork(ctx, snapshotID, blaxel.SnapshotForkParams{
			SandboxForkRequest: blaxel.SandboxForkRequestParam{
				TargetName: forkName,
				TargetType: "sandbox",
			},
		})
		if err != nil {
			t.Fatalf("failed to fork snapshot: %v", err)
		}
		if fork.Name != forkName {
			t.Errorf("expected fork %s, got %s", forkName, fork.Name)
		}
		if fork.Type != blaxel.SandboxForkResponseTypeSandbox {
			t.Errorf("expected fork type sandbox, got %s", fork.Type)
		}
		forked, err := client.Sandboxes.Get(ctx, forkName, blaxel.SandboxGetParams{})
		if err != nil {
			t.Fatalf("forked sandbox %s not found: %v", forkName, err)
		}
		if forked.Metadata.Name != forkName {
			t.Errorf("expected sandbox %s, got %s", forkName, forked.Metadata.Name)
		}
	})

	t.Run("deletes the snapshot by id", func(t *testing.T) {
		if err := client.Snapshots.Delete(ctx, snapshotID); err != nil {
			t.Fatalf("failed to delete snapshot: %v", err)
		}
		if _, err := client.Snapshots.Get(ctx, snapshotID); err == nil {
			t.Errorf("snapshot %s still exists after delete", snapshotID)
		}
		snapshotID = ""
	})
}
