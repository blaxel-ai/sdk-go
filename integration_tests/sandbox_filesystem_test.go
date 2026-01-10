package integration_tests

import (
	"context"
	"strings"
	"sync"
	"testing"
	"time"

	blaxel "github.com/stainless-sdks/blaxel-go"
)

func TestSandboxFilesystem(t *testing.T) {
	ctx := context.Background()
	client := newTestClient(t)

	sandboxName := uniqueName("fs-test")
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

	t.Run("WriteAndRead", func(t *testing.T) {
		t.Run("writes and reads a text file", func(t *testing.T) {
			content := "Hello, World!"
			path := "/tmp/test-write.txt"

			_, err := sandbox.FS.Write(ctx, path, content)
			if err != nil {
				t.Fatalf("failed to write: %v", err)
			}

			result, err := sandbox.FS.Read(ctx, path)
			if err != nil {
				t.Fatalf("failed to read: %v", err)
			}
			if result != content {
				t.Errorf("expected content %s, got %s", content, result)
			}
		})

		t.Run("writes and reads unicode content", func(t *testing.T) {
			content := "Hello 世界 🌍 émojis"
			path := "/tmp/test-unicode.txt"

			_, err := sandbox.FS.Write(ctx, path, content)
			if err != nil {
				t.Fatalf("failed to write: %v", err)
			}

			result, err := sandbox.FS.Read(ctx, path)
			if err != nil {
				t.Fatalf("failed to read: %v", err)
			}
			if result != content {
				t.Errorf("expected content %s, got %s", content, result)
			}
		})

		t.Run("writes and reads multiline content", func(t *testing.T) {
			content := "Line 1\nLine 2\nLine 3"
			path := "/tmp/test-multiline.txt"

			_, err := sandbox.FS.Write(ctx, path, content)
			if err != nil {
				t.Fatalf("failed to write: %v", err)
			}

			result, err := sandbox.FS.Read(ctx, path)
			if err != nil {
				t.Fatalf("failed to read: %v", err)
			}
			if result != content {
				t.Errorf("expected content %s, got %s", content, result)
			}
		})

		t.Run("overwrites existing file", func(t *testing.T) {
			path := "/tmp/test-overwrite.txt"

			_, err := sandbox.FS.Write(ctx, path, "original")
			if err != nil {
				t.Fatalf("failed to write: %v", err)
			}

			_, err = sandbox.FS.Write(ctx, path, "updated")
			if err != nil {
				t.Fatalf("failed to write: %v", err)
			}

			result, err := sandbox.FS.Read(ctx, path)
			if err != nil {
				t.Fatalf("failed to read: %v", err)
			}
			if result != "updated" {
				t.Errorf("expected content 'updated', got %s", result)
			}
		})
	})

	t.Run("WriteBinaryAndReadBinary", func(t *testing.T) {
		t.Run("readBinary works on text files too", func(t *testing.T) {
			path := "/tmp/test-binary-text.txt"
			_, err := sandbox.FS.Write(ctx, path, "text content")
			if err != nil {
				t.Fatalf("failed to write: %v", err)
			}

			blob, err := sandbox.FS.ReadBinary(ctx, path)
			if err != nil {
				t.Fatalf("failed to read binary: %v", err)
			}

			if string(blob) != "text content" {
				t.Errorf("expected 'text content', got %s", string(blob))
			}
		})
	})

	t.Run("Ls", func(t *testing.T) {
		t.Run("lists files in a directory", func(t *testing.T) {
			// Create some test files
			_, err := sandbox.FS.Write(ctx, "/tmp/ls-test/file1.txt", "content1")
			if err != nil {
				t.Fatalf("failed to write: %v", err)
			}
			_, err = sandbox.FS.Write(ctx, "/tmp/ls-test/file2.txt", "content2")
			if err != nil {
				t.Fatalf("failed to write: %v", err)
			}

			listing, err := sandbox.FS.LS(ctx, "/tmp/ls-test")
			if err != nil {
				t.Fatalf("failed to list: %v", err)
			}

			if len(listing.Files) < 2 {
				t.Errorf("expected at least 2 files, got %d", len(listing.Files))
			}

			names := make(map[string]bool)
			for _, f := range listing.Files {
				names[f.Name] = true
			}
			if !names["file1.txt"] || !names["file2.txt"] {
				t.Error("expected to find file1.txt and file2.txt")
			}
		})

		t.Run("lists subdirectories", func(t *testing.T) {
			_, err := sandbox.FS.Mkdir(ctx, "/tmp/ls-subdir-test/subdir1", "755")
			if err != nil {
				t.Fatalf("failed to mkdir: %v", err)
			}
			_, err = sandbox.FS.Mkdir(ctx, "/tmp/ls-subdir-test/subdir2", "755")
			if err != nil {
				t.Fatalf("failed to mkdir: %v", err)
			}

			listing, err := sandbox.FS.LS(ctx, "/tmp/ls-subdir-test")
			if err != nil {
				t.Fatalf("failed to list: %v", err)
			}

			names := make(map[string]bool)
			for _, d := range listing.Subdirectories {
				names[d.Name] = true
			}
			if !names["subdir1"] || !names["subdir2"] {
				t.Error("expected to find subdir1 and subdir2")
			}
		})

		t.Run("returns file metadata", func(t *testing.T) {
			_, err := sandbox.FS.Write(ctx, "/tmp/meta-test.txt", "some content")
			if err != nil {
				t.Fatalf("failed to write: %v", err)
			}

			listing, err := sandbox.FS.LS(ctx, "/tmp")
			if err != nil {
				t.Fatalf("failed to list: %v", err)
			}

			var found *blaxel.FilesystemRead
			for i, f := range listing.Files {
				if f.Name == "meta-test.txt" {
					found = &listing.Files[i]
					break
				}
			}
			if found == nil {
				t.Error("expected to find meta-test.txt")
			}
			if found.Path != "/tmp/meta-test.txt" {
				t.Errorf("expected path '/tmp/meta-test.txt', got %s", found.Path)
			}
		})
	})

	t.Run("Mkdir", func(t *testing.T) {
		t.Run("creates a directory", func(t *testing.T) {
			path := "/tmp/new-dir-" + uniqueName("")
			_, err := sandbox.FS.Mkdir(ctx, path, "755")
			if err != nil {
				t.Fatalf("failed to mkdir: %v", err)
			}

			_, err = sandbox.FS.LS(ctx, path)
			if err != nil {
				t.Fatalf("expected directory to exist: %v", err)
			}
		})

		t.Run("creates nested directories", func(t *testing.T) {
			path := "/tmp/nested-" + uniqueName("") + "/level1/level2"
			_, err := sandbox.FS.Mkdir(ctx, path, "755")
			if err != nil {
				t.Fatalf("failed to mkdir: %v", err)
			}

			_, err = sandbox.FS.LS(ctx, path)
			if err != nil {
				t.Fatalf("expected directory to exist: %v", err)
			}
		})
	})

	t.Run("Cp", func(t *testing.T) {
		t.Run("copies a file", func(t *testing.T) {
			src := "/tmp/cp-src.txt"
			dst := "/tmp/cp-dst.txt"

			_, err := sandbox.FS.Write(ctx, src, "copy me")
			if err != nil {
				t.Fatalf("failed to write: %v", err)
			}

			err = sandbox.FS.CP(ctx, src, dst, sandbox.Process)
			if err != nil {
				t.Fatalf("failed to copy: %v", err)
			}

			content, err := sandbox.FS.Read(ctx, dst)
			if err != nil {
				t.Fatalf("failed to read: %v", err)
			}
			if content != "copy me" {
				t.Errorf("expected content 'copy me', got %s", content)
			}
		})

		t.Run("copies a directory", func(t *testing.T) {
			srcDir := "/tmp/cp-dir-src"
			dstDir := "/tmp/cp-dir-dst"

			_, err := sandbox.FS.Write(ctx, srcDir+"/file.txt", "content")
			if err != nil {
				t.Fatalf("failed to write: %v", err)
			}

			err = sandbox.FS.CP(ctx, srcDir, dstDir, sandbox.Process)
			if err != nil {
				t.Fatalf("failed to copy: %v", err)
			}

			content, err := sandbox.FS.Read(ctx, dstDir+"/file.txt")
			if err != nil {
				t.Fatalf("failed to read: %v", err)
			}
			if content != "content" {
				t.Errorf("expected content 'content', got %s", content)
			}
		})
	})

	t.Run("Rm", func(t *testing.T) {
		t.Run("removes a file", func(t *testing.T) {
			path := "/tmp/rm-file.txt"
			_, err := sandbox.FS.Write(ctx, path, "delete me")
			if err != nil {
				t.Fatalf("failed to write: %v", err)
			}

			_, err = sandbox.FS.RM(ctx, path, false)
			if err != nil {
				t.Fatalf("failed to remove: %v", err)
			}

			// File should no longer exist
			_, err = sandbox.FS.Read(ctx, path)
			if err == nil {
				t.Error("expected error reading deleted file")
			}
		})

		t.Run("removes a directory recursively", func(t *testing.T) {
			dir := "/tmp/rm-dir"
			_, err := sandbox.FS.Write(ctx, dir+"/file.txt", "content")
			if err != nil {
				t.Fatalf("failed to write: %v", err)
			}
			_, err = sandbox.FS.Mkdir(ctx, dir+"/subdir", "755")
			if err != nil {
				t.Fatalf("failed to mkdir: %v", err)
			}

			_, err = sandbox.FS.RM(ctx, dir, true)
			if err != nil {
				t.Fatalf("failed to remove: %v", err)
			}

			_, err = sandbox.FS.LS(ctx, dir)
			if err == nil {
				t.Error("expected error listing deleted directory")
			}
		})

		t.Run("fails to remove non-empty directory without recursive flag", func(t *testing.T) {
			dir := "/tmp/rm-nonempty"
			_, err := sandbox.FS.Write(ctx, dir+"/file.txt", "content")
			if err != nil {
				t.Fatalf("failed to write: %v", err)
			}

			_, err = sandbox.FS.RM(ctx, dir, false)
			if err == nil {
				t.Error("expected error removing non-empty directory without recursive flag")
			}

			// Cleanup
			_, _ = sandbox.FS.RM(ctx, dir, true)
		})
	})

	t.Run("Watch", func(t *testing.T) {
		t.Run("watches for file changes", func(t *testing.T) {
			dir := "/tmp/watch-test-" + uniqueName("")
			_, err := sandbox.FS.Mkdir(ctx, dir, "755")
			if err != nil {
				t.Fatalf("failed to mkdir: %v", err)
			}

			var changeDetected bool
			var mu sync.Mutex

			handle := sandbox.FS.Watch(ctx, dir, func(event blaxel.WatchEvent) {
				if event.Name == "watched-file.txt" {
					mu.Lock()
					changeDetected = true
					mu.Unlock()
				}
			}, nil)

			time.Sleep(200 * time.Millisecond)
			// Trigger a file change
			_, err = sandbox.FS.Write(ctx, dir+"/watched-file.txt", "new content")
			if err != nil {
				t.Fatalf("failed to write: %v", err)
			}

			// Wait for callback
			time.Sleep(100 * time.Millisecond)
			handle.Close()

			mu.Lock()
			detected := changeDetected
			mu.Unlock()

			if !detected {
				t.Error("expected change to be detected")
			}
		})

		t.Run("watches with content option", func(t *testing.T) {
			dir := "/tmp/watch-content-" + uniqueName("")
			_, err := sandbox.FS.Mkdir(ctx, dir, "755")
			if err != nil {
				t.Fatalf("failed to mkdir: %v", err)
			}

			var receivedContent string
			var mu sync.Mutex

			handle := sandbox.FS.Watch(ctx, dir, func(event blaxel.WatchEvent) {
				if event.Content != "" {
					mu.Lock()
					receivedContent = event.Content
					mu.Unlock()
				}
			}, &blaxel.WatchOptions{WithContent: true})

			time.Sleep(1 * time.Second)
			_, err = sandbox.FS.Write(ctx, dir+"/content-file.txt", "the content")
			if err != nil {
				t.Fatalf("failed to write: %v", err)
			}

			time.Sleep(1 * time.Second)
			handle.Close()

			mu.Lock()
			content := receivedContent
			mu.Unlock()

			if content != "the content" {
				t.Errorf("expected content 'the content', got %s", content)
			}
		})
	})

	t.Run("MultipartUpload", func(t *testing.T) {
		t.Run("uploads small file via regular upload", func(t *testing.T) {
			content := strings.Repeat("Hello, world! ", 1000) // ~14KB
			path := "/tmp/small-upload.txt"

			_, err := sandbox.FS.Write(ctx, path, content)
			if err != nil {
				t.Fatalf("failed to write: %v", err)
			}

			result, err := sandbox.FS.Read(ctx, path)
			if err != nil {
				t.Fatalf("failed to read: %v", err)
			}
			if result != content {
				t.Errorf("content mismatch")
			}
		})

		t.Run("uploads large text file via multipart", func(t *testing.T) {
			content := strings.Repeat("Large file content line. ", 50000) // ~1.2MB
			path := "/tmp/large-upload.txt"

			_, err := sandbox.FS.Write(ctx, path, content)
			if err != nil {
				t.Fatalf("failed to write: %v", err)
			}

			result, err := sandbox.FS.Read(ctx, path)
			if err != nil {
				t.Fatalf("failed to read: %v", err)
			}
			if len(result) != len(content) {
				t.Errorf("expected length %d, got %d", len(content), len(result))
			}
			if result != content {
				t.Error("content mismatch")
			}
		})

		t.Run("uploads large binary file via multipart", func(t *testing.T) {
			size := 2 * 1024 * 1024 // 2MB
			binaryContent := make([]byte, size)
			// Fill with pattern for verification
			for i := 0; i < size; i++ {
				binaryContent[i] = byte(i % 256)
			}
			path := "/tmp/large-binary-upload.bin"

			_, err := sandbox.FS.WriteBinary(ctx, path, binaryContent, "644")
			if err != nil {
				t.Fatalf("failed to write binary: %v", err)
			}

			result, err := sandbox.FS.ReadBinary(ctx, path)
			if err != nil {
				t.Fatalf("failed to read binary: %v", err)
			}

			if len(result) != size {
				t.Errorf("expected length %d, got %d", size, len(result))
			}
			for i := 0; i < size; i++ {
				if result[i] != byte(i%256) {
					t.Errorf("content mismatch at index %d", i)
					break
				}
			}
		})

		t.Run("uploads very large file with multiple parts", func(t *testing.T) {
			content := strings.Repeat("X", 6*1024*1024) // 6MB
			path := "/tmp/very-large-upload.txt"

			_, err := sandbox.FS.Write(ctx, path, content)
			if err != nil {
				t.Fatalf("failed to write: %v", err)
			}

			result, err := sandbox.FS.Read(ctx, path)
			if err != nil {
				t.Fatalf("failed to read: %v", err)
			}
			if len(result) != len(content) {
				t.Errorf("expected length %d, got %d", len(content), len(result))
			}
		})
	})

	t.Run("ParallelOperations", func(t *testing.T) {
		t.Run("handles 100 parallel file reads", func(t *testing.T) {
			// Create a test file
			content := strings.Repeat("A", 200*1024) // 200KB
			path := "/tmp/parallel-read.txt"
			_, err := sandbox.FS.Write(ctx, path, content)
			if err != nil {
				t.Fatalf("failed to write: %v", err)
			}

			// Make 100 parallel read calls
			var wg sync.WaitGroup
			results := make(chan int, 100)
			errors := make(chan error, 100)

			for i := 0; i < 100; i++ {
				wg.Add(1)
				go func() {
					defer wg.Done()
					fileContent, err := sandbox.FS.Read(ctx, path)
					if err != nil {
						errors <- err
						return
					}
					results <- len(fileContent)
				}()
			}

			wg.Wait()
			close(results)
			close(errors)

			// Check for errors
			for err := range errors {
				t.Errorf("read error: %v", err)
			}

			// All reads should return the same size
			for size := range results {
				if size != len(content) {
					t.Errorf("expected size %d, got %d", len(content), size)
				}
			}
		})
	})
}
