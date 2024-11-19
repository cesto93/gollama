# gollama
Easy Ollama package for Golang

### Example use

```go
package main

import (
    "fmt"

    "github.com/jonathanhecl/gollama/v1"
)

func main() {
    model := "llama3.2"

    g := gollama.New(model)
    g.Verbose = true
    if err := g.PullIfMissing(model); err != nil {
        fmt.Println("Error:", err)
        return
    }
    prompt := "what is the capital of Argentina?"
    output, err := g.Chat(gollama.GollamaInput{Prompt: prompt})
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Printf("\n%s\n", output.Content)
}
```
