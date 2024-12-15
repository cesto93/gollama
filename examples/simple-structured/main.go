package main

import (
	"fmt"

	"github.com/jonathanhecl/gollama"
)

func main() {
	g := gollama.New("llama3.2")
	g.Verbose = true
	if err := g.PullIfMissing(); err != nil {
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

	output, err := g.Chat(prompt, option)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("\n%s\n", output.Content)
}
