package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/stainless-sdks/blaxel-go"
)

func main() {
	client, err := blaxel.NewDefaultClient()
	if err != nil {
		fmt.Printf("Could not connect to blaxel")
		return
	}
	type Input struct {
		Inputs string `json:"inputs"`
	}
	inputBody, err := json.Marshal(Input{Inputs: "Hello, world!"})
	if err != nil {
		fmt.Printf("Could not marshal message: %v", err)
		return
	}
	ctx := context.Background()
	res, err := client.Run(ctx, blaxel.GetDefaultWorkspace(), "agents", "test-agent-ts", "POST", "", inputBody)
	if err != nil {
		fmt.Printf("Could not run agent: %v", err)
		return
	}
	body, err := io.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Could not read response body: %v", err)
		return
	}
	fmt.Printf("Agent run: %v", string(body))
}
