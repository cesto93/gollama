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

	type Capital struct {
		Capital string `json:"capital",required="true"`
	}

	option, _ := gollama.StructToStructuredFormat(Capital{})

	fmt.Printf("Option: %+v\n", option)

	output, err := g.Chat(prompt, option)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Printf("\n%s\n", output.Content)
}
