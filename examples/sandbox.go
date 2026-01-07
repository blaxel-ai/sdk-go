package main

import (
	"context"
	"fmt"
	"time"

	"github.com/stainless-sdks/blaxel-go"
)

func main() {
	client, err := blaxel.NewDefaultClient()
	if err != nil {
		fmt.Printf("Could not connect to blaxel")
		return
	}
	ctx := context.Background()
	sbx, err := client.Sandboxes.NewInstance(ctx, blaxel.SandboxNewParams{
		Sandbox: blaxel.SandboxParam{
			Metadata: blaxel.MetadataParam{
				Name: "my-sandbox",
			},
			Spec: blaxel.SandboxSpecParam{
				Runtime: blaxel.SandboxRuntimeParam{
					Image:  blaxel.Opt("blaxel/nextjs"),
					Memory: blaxel.Opt[int64](4096),
					Ports: []blaxel.PortParam{
						{
							Target: blaxel.Opt[int64](3000),
						},
					},
				},
			},
		},
	})
	if err != nil {
		fmt.Printf("Could not create sandbox: %v", err)
		return
	}
	fmt.Printf("Sandbox created: %s\n", sbx.Metadata.URL)
	preview, err := sbx.Previews.New(ctx, blaxel.SandboxPreviewNewParams{
		Preview: blaxel.PreviewParam{
			Metadata: blaxel.PreviewMetadataParam{
				Name: "my-sandbox-preview",
			},
			Spec: blaxel.PreviewSpecParam{
				Port:   blaxel.Opt[int64](3000),
				Public: blaxel.Opt(false),
			},
		},
	})
	if err != nil {
		fmt.Printf("Could not create preview: %v", err)
		return
	}

	token, err := preview.Tokens.New(ctx, blaxel.PreviewTokenParam{
		Metadata: blaxel.PreviewTokenMetadataParam{
			Name: "my-sandbox-preview-token",
		},
		Spec: blaxel.PreviewTokenSpecParam{
			ExpiresAt: blaxel.Opt(time.Now().Add(1 * time.Hour).Format(time.RFC3339)),
		},
	})
	if err != nil {
		fmt.Printf("Could not create token: %v", err)
		return
	}
	fmt.Printf("Preview created: %s?bl_preview_token=%s\n", preview.Spec.URL, token.Spec.Token)
}
