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

	prompt := "what is the capital of Argentina? and how many people live there?"

	type Capital struct {
		Capital    string `json:"capital" required:"true" description:"the capital of a country"`
		Population string `json:"population" required:"true" description:"how many people live in a city"`
	}

	option := gollama.StructToStructuredFormat(Capital{})

	fmt.Printf("Option: %+v\n", option)

	output, err := g.Chat(ctx, prompt, option)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("\n%s\n", output.Content)
}
