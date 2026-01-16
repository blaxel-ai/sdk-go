package integration_tests

import (
	"context"
	"strings"
	"sync"
	"testing"
	"time"

	blaxel "github.com/blaxel-ai/sdk-go"
)

func TestSandboxSessions(t *testing.T) {
	ctx := context.Background()
	client := newTestClient(t)

	sandboxName := uniqueName("session-test")
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

	t.Run("Create", func(t *testing.T) {
		t.Run("creates a session with expiration", func(t *testing.T) {
			expiresAt := time.Now().Add(24 * time.Hour)

			session, err := sandbox.Sessions.Create(ctx, &blaxel.SessionCreateOptions{
				ExpiresAt: &expiresAt,
			})
			if err != nil {
				t.Fatalf("failed to create session: %v", err)
			}

			if session.Name == "" {
				t.Error("expected session name to be set")
			}
			if session.Token == "" {
				t.Error("expected session token to be set")
			}
			if session.URL == "" {
				t.Error("expected session URL to be set")
			}

			_ = sandbox.Sessions.Delete(ctx, session.Name)
		})

		t.Run("session has valid URL and token", func(t *testing.T) {
			expiresAt := time.Now().Add(1 * time.Hour)

			session, err := sandbox.Sessions.Create(ctx, &blaxel.SessionCreateOptions{
				ExpiresAt: &expiresAt,
			})
			if err != nil {
				t.Fatalf("failed to create session: %v", err)
			}

			if !strings.HasPrefix(session.URL, "http") {
				t.Errorf("expected URL to start with 'http', got %s", session.URL)
			}
			if len(session.Token) == 0 {
				t.Error("expected token to have length > 0")
			}

			_ = sandbox.Sessions.Delete(ctx, session.Name)
		})
	})

	t.Run("List", func(t *testing.T) {
		t.Run("lists all sessions", func(t *testing.T) {
			expiresAt := time.Now().Add(1 * time.Hour)

			session1, err := sandbox.Sessions.Create(ctx, &blaxel.SessionCreateOptions{
				ExpiresAt: &expiresAt,
			})
			if err != nil {
				t.Fatalf("failed to create session: %v", err)
			}

			session2, err := sandbox.Sessions.Create(ctx, &blaxel.SessionCreateOptions{
				ExpiresAt: &expiresAt,
			})
			if err != nil {
				t.Fatalf("failed to create session: %v", err)
			}

			sessions, err := sandbox.Sessions.List(ctx)
			if err != nil {
				t.Fatalf("failed to list sessions: %v", err)
			}

			if len(sessions) < 2 {
				t.Errorf("expected at least 2 sessions, got %d", len(sessions))
			}

			found1, found2 := false, false
			for _, s := range sessions {
				if s.Name == session1.Name {
					found1 = true
				}
				if s.Name == session2.Name {
					found2 = true
				}
			}

			if !found1 || !found2 {
				t.Error("expected to find both sessions in list")
			}

			_ = sandbox.Sessions.Delete(ctx, session1.Name)
			_ = sandbox.Sessions.Delete(ctx, session2.Name)
		})
	})

	t.Run("Delete", func(t *testing.T) {
		t.Run("deletes a session", func(t *testing.T) {
			expiresAt := time.Now().Add(1 * time.Hour)
			session, err := sandbox.Sessions.Create(ctx, &blaxel.SessionCreateOptions{
				ExpiresAt: &expiresAt,
			})
			if err != nil {
				t.Fatalf("failed to create session: %v", err)
			}

			err = sandbox.Sessions.Delete(ctx, session.Name)
			if err != nil {
				t.Fatalf("failed to delete session: %v", err)
			}

			sessions, err := sandbox.Sessions.List(ctx)
			if err != nil {
				t.Fatalf("failed to list sessions: %v", err)
			}

			for _, s := range sessions {
				if s.Name == session.Name {
					t.Error("expected session to be deleted")
				}
			}
		})
	})

	t.Run("FromSession", func(t *testing.T) {
		t.Run("creates sandbox instance from session", func(t *testing.T) {
			expiresAt := time.Now().Add(1 * time.Hour)
			session, err := sandbox.Sessions.Create(ctx, &blaxel.SessionCreateOptions{
				ExpiresAt: &expiresAt,
			})
			if err != nil {
				t.Fatalf("failed to create session: %v", err)
			}

			sandboxFromSession := client.Sandboxes.FromSession(*session)

			// Should be able to perform operations
			listing, err := sandboxFromSession.FS.LS(ctx, "/")
			if err != nil {
				t.Fatalf("failed to list: %v", err)
			}
			if len(listing.Subdirectories) == 0 {
				t.Error("expected to find subdirectories")
			}

			_ = sandbox.Sessions.Delete(ctx, session.Name)
		})

		t.Run("session sandbox can execute processes", func(t *testing.T) {
			expiresAt := time.Now().Add(1 * time.Hour)
			session, err := sandbox.Sessions.Create(ctx, &blaxel.SessionCreateOptions{
				ExpiresAt: &expiresAt,
			})
			if err != nil {
				t.Fatalf("failed to create session: %v", err)
			}

			sandboxFromSession := client.Sandboxes.FromSession(*session)

			result, err := sandboxFromSession.Process.New(ctx, blaxel.ProcessRequestParam{
				Command:           "echo 'from session'",
				WaitForCompletion: blaxel.Bool(true),
			})
			if err != nil {
				t.Fatalf("failed to exec: %v", err)
			}

			if !strings.Contains(result.Stdout, "from session") {
				t.Errorf("expected stdout to contain 'from session', got %s", result.Stdout)
			}

			_ = sandbox.Sessions.Delete(ctx, session.Name)
		})

		t.Run("session sandbox can stream logs", func(t *testing.T) {
			expiresAt := time.Now().Add(1 * time.Hour)
			session, err := sandbox.Sessions.Create(ctx, &blaxel.SessionCreateOptions{
				ExpiresAt: &expiresAt,
			})
			if err != nil {
				t.Fatalf("failed to create session: %v", err)
			}

			sandboxFromSession := client.Sandboxes.FromSession(*session)

			_, err = sandboxFromSession.Process.New(ctx, blaxel.ProcessRequestParam{
				Name:              blaxel.String("stream-session"),
				Command:           "for i in 1 2 3; do echo $i; sleep 1; done",
				WaitForCompletion: blaxel.Bool(false),
			})
			if err != nil {
				t.Fatalf("failed to exec: %v", err)
			}

			var logs []string
			var mu sync.Mutex

			stream := sandboxFromSession.Process.StreamLogs(ctx, "stream-session", blaxel.ProcessStreamOptions{
				OnLog: func(log string) {
					mu.Lock()
					logs = append(logs, log)
					mu.Unlock()
				},
			})

			_, err = sandboxFromSession.Process.Wait(ctx, "stream-session", 30*time.Second, time.Second)
			if err != nil {
				t.Fatalf("failed to wait: %v", err)
			}
			time.Sleep(100 * time.Millisecond)
			stream.Close()

			mu.Lock()
			logCount := len(logs)
			mu.Unlock()

			if logCount == 0 {
				t.Error("expected to receive logs")
			}

			_ = sandbox.Sessions.Delete(ctx, session.Name)
		})

		t.Run("session sandbox can watch files", func(t *testing.T) {
			expiresAt := time.Now().Add(1 * time.Hour)
			session, err := sandbox.Sessions.Create(ctx, &blaxel.SessionCreateOptions{
				ExpiresAt: &expiresAt,
			})
			if err != nil {
				t.Fatalf("failed to create session: %v", err)
			}

			sandboxFromSession := client.Sandboxes.FromSession(*session)

			var changeDetected bool
			var mu sync.Mutex

			handle := sandboxFromSession.FS.Watch(ctx, "/", func(event blaxel.WatchEvent) {
				mu.Lock()
				changeDetected = true
				mu.Unlock()
			}, nil)

			time.Sleep(100 * time.Millisecond)
			_, err = sandboxFromSession.FS.Write(ctx, "/session-test.txt", "content")
			if err != nil {
				t.Fatalf("failed to write: %v", err)
			}

			time.Sleep(200 * time.Millisecond)
			handle.Close()

			mu.Lock()
			detected := changeDetected
			mu.Unlock()

			if !detected {
				t.Error("expected change to be detected")
			}

			_ = sandbox.Sessions.Delete(ctx, session.Name)
		})
	})
}
