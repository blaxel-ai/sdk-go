package integration_tests

import (
	"context"
	"io"
	"testing"

	blaxel "github.com/blaxel-ai/sdk-go"
)

func TestSandboxFetch(t *testing.T) {
	ctx := context.Background()
	client := newTestClient(t)

	sandboxName := uniqueName("fetch-test")
	sandbox, err := client.Sandboxes.NewInstance(ctx, blaxel.SandboxNewParams{
		Sandbox: blaxel.SandboxParam{
			Metadata: blaxel.MetadataParam{
				Name:   sandboxName,
				Labels: defaultLabels,
			},
			Spec: blaxel.SandboxSpecParam{
				Region: blaxel.String(defaultRegion),
				Runtime: blaxel.SandboxRuntimeParam{
					Image:  blaxel.String("blaxel/node:latest"),
					Memory: blaxel.Int(2048),
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

	t.Cleanup(func() {
		_, _ = client.Sandboxes.Delete(ctx, sandboxName)
	})

	nodeServerCommand := `sleep 2 && node -e "
const http = require('http');
const server = http.createServer((req, res) => {
  res.writeHead(200, { 'Content-Type': 'text/plain' });
  res.end('OK');
});
server.listen(3000);
"`

	_, err = sandbox.Process.New(ctx, blaxel.ProcessRequestParam{
		Name:         blaxel.String("http-server"),
		Command:      nodeServerCommand,
		WaitForPorts: []int64{3000},
	})
	if err != nil {
		t.Fatalf("failed to start http server: %v", err)
	}

	t.Run("fetch returns response from port", func(t *testing.T) {
		resp, err := sandbox.Fetch(ctx, 3000, "/")
		if err != nil {
			t.Fatalf("fetch failed: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			t.Errorf("expected status 200, got %d", resp.StatusCode)
		}

		body, err := io.ReadAll(resp.Body)
		if err != nil {
			t.Fatalf("failed to read body: %v", err)
		}
		if string(body) != "OK" {
			t.Errorf("expected body 'OK', got %q", string(body))
		}
	})

	t.Run("fetch with path", func(t *testing.T) {
		resp, err := sandbox.Fetch(ctx, 3000, "/some/path")
		if err != nil {
			t.Fatalf("fetch failed: %v", err)
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			t.Errorf("expected status 200, got %d", resp.StatusCode)
		}
	})
}
