package integration_tests

import (
	"context"
	"strings"
	"sync"
	"testing"

	blaxel "github.com/stainless-sdks/blaxel-go"
)

func TestVolumes(t *testing.T) {
	ctx := context.Background()
	client := newTestClient(t)

	var createdSandboxes []string
	var createdVolumes []string
	var mu sync.Mutex

	trackSandbox := func(name string) {
		mu.Lock()
		createdSandboxes = append(createdSandboxes, name)
		mu.Unlock()
	}

	trackVolume := func(name string) {
		mu.Lock()
		createdVolumes = append(createdVolumes, name)
		mu.Unlock()
	}

	t.Cleanup(func() {
		// Clean up sandboxes first
		for _, name := range createdSandboxes {
			_, _ = client.Sandboxes.Delete(ctx, name)
			waitForSandboxDeletion(ctx, client, name, 15)
		}
		// Then volumes
		for _, name := range createdVolumes {
			_, _ = client.Volumes.Delete(ctx, name)
		}
	})

	t.Run("VolumeInstanceCRUD", func(t *testing.T) {
		t.Run("creates a volume", func(t *testing.T) {
			name := uniqueName("volume")
			volume, err := client.Volumes.NewInstance(ctx, blaxel.VolumeCreateConfiguration{
				Name:   name,
				Labels: defaultLabels,
				Size:   1024, // 1GB
				Region: defaultRegion,
			})
			if err != nil {
				t.Fatalf("failed to create volume: %v", err)
			}
			trackVolume(name)

			if volume.Name() != name {
				t.Errorf("expected name %s, got %s", name, volume.Name())
			}
		})

		t.Run("creates a volume with display name", func(t *testing.T) {
			name := uniqueName("volume-display")
			volume, err := client.Volumes.NewInstance(ctx, blaxel.VolumeCreateConfiguration{
				Name:        name,
				DisplayName: "My Test Volume",
				Labels:      defaultLabels,
				Size:        1024,
				Region:      defaultRegion,
			})
			if err != nil {
				t.Fatalf("failed to create volume: %v", err)
			}
			trackVolume(name)

			if volume.DisplayName() != "My Test Volume" {
				t.Errorf("expected displayName 'My Test Volume', got %s", volume.DisplayName())
			}
		})

		t.Run("gets a volume", func(t *testing.T) {
			name := uniqueName("volume-get")
			_, err := client.Volumes.NewInstance(ctx, blaxel.VolumeCreateConfiguration{
				Name:   name,
				Labels: defaultLabels,
				Size:   1024,
				Region: defaultRegion,
			})
			if err != nil {
				t.Fatalf("failed to create volume: %v", err)
			}
			trackVolume(name)

			volume, err := client.Volumes.GetInstance(ctx, name)
			if err != nil {
				t.Fatalf("failed to get volume: %v", err)
			}
			if volume.Name() != name {
				t.Errorf("expected name %s, got %s", name, volume.Name())
			}
		})

		t.Run("lists volumes", func(t *testing.T) {
			name := uniqueName("volume-list")
			_, err := client.Volumes.NewInstance(ctx, blaxel.VolumeCreateConfiguration{
				Name:   name,
				Labels: defaultLabels,
				Size:   1024,
				Region: defaultRegion,
			})
			if err != nil {
				t.Fatalf("failed to create volume: %v", err)
			}
			trackVolume(name)

			volumes, err := client.Volumes.ListInstances(ctx)
			if err != nil {
				t.Fatalf("failed to list volumes: %v", err)
			}

			found := false
			for _, v := range volumes {
				if v.Name() == name {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("expected to find volume %s in list", name)
			}
		})

		t.Run("deletes a volume", func(t *testing.T) {
			name := uniqueName("volume-delete")
			volume, err := client.Volumes.NewInstance(ctx, blaxel.VolumeCreateConfiguration{
				Name:   name,
				Labels: defaultLabels,
				Size:   1024,
				Region: defaultRegion,
			})
			if err != nil {
				t.Fatalf("failed to create volume: %v", err)
			}

			err = volume.Delete(ctx)
			if err != nil {
				t.Fatalf("failed to delete volume: %v", err)
			}
			waitForVolumeDeletion(ctx, client, name, 30)

			// Volume should no longer exist
			_, err = client.Volumes.GetInstance(ctx, name)
			if err == nil {
				t.Error("expected error getting deleted volume")
			}
		})
	})

	t.Run("MountingVolumesToSandboxes", func(t *testing.T) {
		t.Run("mounts a volume to a sandbox", func(t *testing.T) {
			volumeName := uniqueName("mount-vol")
			sandboxName := uniqueName("mount-sandbox")

			_, err := client.Volumes.NewInstance(ctx, blaxel.VolumeCreateConfiguration{
				Name:   volumeName,
				Labels: defaultLabels,
				Size:   1024,
				Region: defaultRegion,
			})
			if err != nil {
				t.Fatalf("failed to create volume: %v", err)
			}
			trackVolume(volumeName)

			sandbox, err := client.Sandboxes.NewInstance(ctx, blaxel.SandboxNewParams{
				Sandbox: blaxel.SandboxParam{
					Metadata: blaxel.MetadataParam{
						Name:   sandboxName,
						Labels: defaultLabels,
					},
					Spec: blaxel.SandboxSpecParam{
						Region: blaxel.String(defaultRegion),
						Runtime: blaxel.SandboxRuntimeParam{
							Image:  blaxel.String(defaultImage),
							Memory: blaxel.Int(2048),
						},
						Volumes: []blaxel.VolumeAttachmentParam{
							{
								Name:      blaxel.String(volumeName),
								MountPath: blaxel.String("/data"),
								ReadOnly:  blaxel.Bool(false),
							},
						},
					},
				},
			})
			if err != nil {
				t.Fatalf("failed to create sandbox: %v", err)
			}
			trackSandbox(sandboxName)

			// Verify mount by writing a file
			_, err = sandbox.Process.New(ctx, blaxel.ProcessRequestParam{
				Command:           "echo 'mounted' > /data/test.txt",
				WaitForCompletion: blaxel.Bool(true),
			})
			if err != nil {
				t.Fatalf("failed to write to volume: %v", err)
			}

			result, err := sandbox.Process.New(ctx, blaxel.ProcessRequestParam{
				Command:           "cat /data/test.txt",
				WaitForCompletion: blaxel.Bool(true),
			})
			if err != nil {
				t.Fatalf("failed to read from volume: %v", err)
			}

			if !strings.Contains(result.Stdout, "mounted") {
				t.Errorf("expected stdout to contain 'mounted', got %s", result.Stdout)
			}
		})
	})

	t.Run("VolumePersistence", func(t *testing.T) {
		t.Run("data persists across sandbox recreations", func(t *testing.T) {
			volumeName := uniqueName("persist-vol")
			fileContent := "persistent data " + uniqueName("")

			_, err := client.Volumes.NewInstance(ctx, blaxel.VolumeCreateConfiguration{
				Name:   volumeName,
				Labels: defaultLabels,
				Size:   1024,
				Region: defaultRegion,
			})
			if err != nil {
				t.Fatalf("failed to create volume: %v", err)
			}
			trackVolume(volumeName)

			// First sandbox - write data
			sandbox1Name := uniqueName("persist-1")
			sandbox1, err := client.Sandboxes.NewInstance(ctx, blaxel.SandboxNewParams{
				Sandbox: blaxel.SandboxParam{
					Metadata: blaxel.MetadataParam{
						Name:   sandbox1Name,
						Labels: defaultLabels,
					},
					Spec: blaxel.SandboxSpecParam{
						Region: blaxel.String(defaultRegion),
						Runtime: blaxel.SandboxRuntimeParam{
							Image: blaxel.String(defaultImage),
						},
						Volumes: []blaxel.VolumeAttachmentParam{
							{
								Name:      blaxel.String(volumeName),
								MountPath: blaxel.String("/persistent"),
								ReadOnly:  blaxel.Bool(false),
							},
						},
					},
				},
			})
			if err != nil {
				t.Fatalf("failed to create sandbox1: %v", err)
			}

			_, err = sandbox1.Process.New(ctx, blaxel.ProcessRequestParam{
				Command:           "echo '" + fileContent + "' > /persistent/data.txt",
				WaitForCompletion: blaxel.Bool(true),
			})
			if err != nil {
				t.Fatalf("failed to write data: %v", err)
			}

			// Delete first sandbox and wait for full deletion
			_, _ = client.Sandboxes.Delete(ctx, sandbox1Name)
			waitForSandboxDeletion(ctx, client, sandbox1Name, 15)

			// Second sandbox - read data
			sandbox2Name := uniqueName("persist-2")
			sandbox2, err := client.Sandboxes.NewInstance(ctx, blaxel.SandboxNewParams{
				Sandbox: blaxel.SandboxParam{
					Metadata: blaxel.MetadataParam{
						Name:   sandbox2Name,
						Labels: defaultLabels,
					},
					Spec: blaxel.SandboxSpecParam{
						Region: blaxel.String(defaultRegion),
						Runtime: blaxel.SandboxRuntimeParam{
							Image: blaxel.String(defaultImage),
						},
						Volumes: []blaxel.VolumeAttachmentParam{
							{
								Name:      blaxel.String(volumeName),
								MountPath: blaxel.String("/data"),
								ReadOnly:  blaxel.Bool(false),
							},
						},
					},
				},
			})
			if err != nil {
				t.Fatalf("failed to create sandbox2: %v", err)
			}
			trackSandbox(sandbox2Name)

			result, err := sandbox2.Process.New(ctx, blaxel.ProcessRequestParam{
				Command:           "cat /data/data.txt",
				WaitForCompletion: blaxel.Bool(true),
			})
			if err != nil {
				t.Fatalf("failed to read data: %v", err)
			}

			if !strings.Contains(strings.TrimSpace(result.Logs), fileContent) {
				t.Errorf("expected logs to contain '%s', got %s", fileContent, result.Logs)
			}
		})
	})
}
