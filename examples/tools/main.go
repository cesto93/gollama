package main

import (
	"fmt"
	"strconv"

	"github.com/jonathanhecl/gollama"
)

func main() {
	g := gollama.New("llama3.2")
	g.Verbose = true
	if err := g.PullIfMissing(); err != nil {
		fmt.Println("Error:", err)
		return
	}

	prompt := "what is 300 more 738.2?"

	option := sumFunc()

	fmt.Printf("Option: %+v\n", option)

	output, err := g.Chat(prompt, option)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("\n%+v\n", output.ToolCalls)

	if len(output.ToolCalls) > 0 {
		for _, call := range output.ToolCalls {
			fmt.Printf("Using tool: %+v\n", call)
			if call.Function.Name == "func_sum" {
				a := call.Function.Arguments["a"]
				b := call.Function.Arguments["b"]

				afloat, _ := strconv.ParseFloat(fmt.Sprintf("%v", a), 64)
				bfloat, _ := strconv.ParseFloat(fmt.Sprintf("%v", b), 64)

				fmt.Printf("Sum: %f + %f = %f\n", afloat, bfloat, sum(afloat, bfloat))
			}
		}
	}
}

func sumFunc() gollama.Tool {
	return gollama.Tool{
		Type: "function",
		Function: gollama.ToolFunction{
			Name:        "func_sum",
			Description: "Sum two numbers and return the result",
			Parameters: gollama.StructuredFormat{
				Type: "object",
				Properties: map[string]gollama.FormatProperty{
					"a": {Type: "number", Description: "first number"},
					"b": {Type: "number", Description: "second number"},
				},
				Required: []string{"a", "b"},
			},
		},
	}
}

func sum(a, b float64) float64 {
	return a + b
}
