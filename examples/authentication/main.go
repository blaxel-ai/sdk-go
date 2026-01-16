package main

import (
	"context"
	"fmt"
	"os"

	blaxel "github.com/blaxel-ai/sdk-go"
	"github.com/blaxel-ai/sdk-go/option"
)

func main() {
	fmt.Println("Blaxel Go SDK Authentication Examples")
	fmt.Println("=====================================")

	// Example 1: Using environment variables (BL_API_KEY)
	fmt.Println("Example 1: Authentication with BL_API_KEY environment variable")
	if apiKey := os.Getenv("BL_API_KEY"); apiKey != "" {
		client := blaxel.NewClient(option.WithAPIKey(apiKey))
		testClient(&client, "BL_API_KEY")
	} else {
		fmt.Println("  BL_API_KEY not set, skipping...")
	}

	// Example 2: Using BL_CLIENT_CREDENTIALS environment variable
	fmt.Println("Example 2: Authentication with BL_CLIENT_CREDENTIALS environment variable")
	if clientCreds := os.Getenv("BL_CLIENT_CREDENTIALS"); clientCreds != "" {
		client := blaxel.NewClient(option.WithClientCredentials(clientCreds))
		testClient(&client, "BL_CLIENT_CREDENTIALS")
	} else {
		fmt.Println("  BL_CLIENT_CREDENTIALS not set, skipping...")
	}

	// Example 3: Using config file
	fmt.Println("Example 3: Authentication from ~/.blaxel/config.yaml")
	config, err := blaxel.LoadConfig()
	if err != nil {
		fmt.Printf("  Error loading config: %v\n\n", err)
	} else if len(config.Workspaces) > 0 {
		workspace := config.Context.Workspace
		if workspace == "" && len(config.Workspaces) > 0 {
			workspace = config.Workspaces[0].Name
		}

		client, err := blaxel.NewClientFromConfig(workspace)
		if err != nil {
			fmt.Printf("  Error creating client from config: %v\n\n", err)
		} else {
			testClient(client, fmt.Sprintf("config file (workspace: %s)", workspace))
		}
	} else {
		fmt.Println("  No workspaces configured in ~/.blaxel/config.yaml, skipping...")
	}

	// Example 4: Using NewClientFromEnv (automatic detection)
	fmt.Println("Example 4: Automatic authentication with NewClientFromEnv")
	client, err := blaxel.NewDefaultClient()
	if err != nil {
		fmt.Printf("  Error creating client from environment: %v\n\n", err)
	} else {
		testClient(&client, "NewClientFromEnv (auto-detection)")
	}

	// Example 5: Custom workspace header
	fmt.Println("Example 5: Using custom workspace header")
	if workspace := os.Getenv("BL_WORKSPACE"); workspace != "" {
		client := blaxel.NewClient(
			option.WithAPIKey(os.Getenv("BL_API_KEY")),
			option.WithWorkspace(workspace),
		)
		fmt.Printf("  Created client with workspace header: %s\n\n", workspace)
		testClient(&client, fmt.Sprintf("custom workspace: %s", workspace))
	} else {
		fmt.Println("  BL_WORKSPACE not set, skipping...")
	}
}

func testClient(client *blaxel.Client, method string) {
	ctx := context.Background()

	fmt.Printf("  Testing with %s...\n", method)

	// Try to list agents to verify authentication works
	// Note: This is just a connectivity test, actual API calls depend on permissions
	_, err := client.Agents.List(ctx)
	if err != nil {
		fmt.Printf("  ❌ Authentication failed: %v\n\n", err)
	} else {
		fmt.Printf("  ✅ Authentication successful!\n\n")
	}
}
