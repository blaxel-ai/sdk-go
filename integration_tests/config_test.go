package integration_tests

import (
	"os"
	"path/filepath"
	"testing"

	blaxel "github.com/blaxel-ai/sdk-go"
)

func TestSaveCredentialsWorkspaceID(t *testing.T) {
	workspace := os.Getenv("BL_WORKSPACE")
	if workspace == "" {
		workspace = blaxel.GetDefaultWorkspace()
	}
	if workspace == "" {
		t.Skip("No workspace configured (set BL_WORKSPACE or login via bl login)")
	}

	// Resolve real credentials from env or config
	creds := blaxel.Credentials{}
	if apiKey := os.Getenv("BL_API_KEY"); apiKey != "" {
		creds.APIKey = apiKey
	} else if clientCreds := os.Getenv("BL_CLIENT_CREDENTIALS"); clientCreds != "" {
		creds.ClientCredentials = clientCreds
	} else {
		loaded, err := blaxel.LoadCredentials(workspace)
		if err != nil || (loaded.APIKey == "" && loaded.AccessToken == "" && loaded.ClientCredentials == "") {
			t.Skip("No credentials available (need BL_API_KEY, BL_CLIENT_CREDENTIALS, or bl login)")
		}
		creds = loaded
	}

	// Use a temp HOME so we don't clobber the real config
	tmpDir := t.TempDir()
	configDir := filepath.Join(tmpDir, ".blaxel")
	os.MkdirAll(configDir, 0o700)

	// Seed a config with just the workspace name — no ID
	seed := "context:\n  workspace: " + workspace + "\nworkspaces:\n- name: " + workspace + "\n  credentials: {}\ntracking: false\n"
	os.WriteFile(filepath.Join(configDir, "config.yaml"), []byte(seed), 0o600)

	origHome := os.Getenv("HOME")
	os.Setenv("HOME", tmpDir)
	t.Cleanup(func() { os.Setenv("HOME", origHome) })

	// Verify ID is not set before
	cfgBefore, _ := blaxel.LoadConfig()
	for _, ws := range cfgBefore.Workspaces {
		if ws.Name == workspace && ws.ID != "" {
			t.Fatal("workspace ID should not be set before SaveCredentials")
		}
	}

	// Call SaveCredentials — this should fetch and persist the workspace ID
	err := blaxel.SaveCredentials(workspace, creds)
	if err != nil {
		t.Fatalf("SaveCredentials failed: %v", err)
	}

	// Reload config and verify the ID was saved
	cfgAfter, err := blaxel.LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig failed after SaveCredentials: %v", err)
	}

	var found bool
	for _, ws := range cfgAfter.Workspaces {
		if ws.Name == workspace {
			found = true
			if ws.ID == "" {
				t.Error("workspace ID was not saved by SaveCredentials")
			} else {
				t.Logf("workspace ID correctly saved: %s", ws.ID)
			}
			break
		}
	}
	if !found {
		t.Errorf("workspace %q not found in config after SaveCredentials", workspace)
	}
}
