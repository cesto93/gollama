package main

import (
	"context"
	"fmt"

	"github.com/jonathanhecl/gollama"
)

func main() {
	ctx := context.Background()
	g := gollama.New("llama3.2")
	g.Verbose = true
	if err := g.PullIfMissing(ctx); err != nil {
		fmt.Println("Error:", err)
		return
	}

	prompt := "what is the capital of Argentina?"
	option := gollama.StructuredFormat{
		Type: "object",
		Properties: map[string]gollama.FormatProperty{
			"capital": {
				Type: "string",
			}},
		Required: []string{"capital"},
	}

	output, err := g.Chat(ctx, prompt, option)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("\n%s\n", output.Content)
}
