package blaxel_test

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/blaxel-ai/sdk-go"
)

func TestSandboxRuntimeParamStorageMBJSON(t *testing.T) {
	data, err := json.Marshal(blaxel.SandboxRuntimeParam{
		StorageMB: blaxel.Int(102400),
	})
	if err != nil {
		t.Fatalf("json.Marshal error = %v", err)
	}
	if got := string(data); !strings.Contains(got, `"storageMb":102400`) {
		t.Fatalf("expected storageMb in marshaled runtime, got %s", got)
	}
}

func TestSandboxRuntimeStorageMBJSON(t *testing.T) {
	var runtime blaxel.SandboxRuntime
	if err := json.Unmarshal([]byte(`{"storageMb":102400}`), &runtime); err != nil {
		t.Fatalf("json.Unmarshal error = %v", err)
	}
	if runtime.StorageMB != 102400 {
		t.Fatalf("StorageMB = %d, want 102400", runtime.StorageMB)
	}
	if !runtime.JSON.StorageMB.Valid() {
		t.Fatal("expected StorageMB JSON metadata to be valid")
	}
}
