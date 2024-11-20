# gollama
Easy Ollama package for Golang

### Example use

> go get -u github.com/jonathanhecl/gollama

```go
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
    output, err := g.Chat(prompt)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    fmt.Printf("\n%s\n", output.Content)
}
```

### Features

- Supports Vision models
- Supports Tools models
- Downloads model if missing
- Chat with model
- Generates embeddings with model
- Get model details

### Functions

- `New(model string) *Gollama`
- `NewWithConfig(config Gollama) *Gollama`
- `Chat(prompt string, ...ChatOption) (*gollama.ChatOutput, error)`
- `Embedding(prompt string) ([]float64, error)`
- `ListModels() ([]ModelInfo, error)`
- `HasModel(model string) (bool, error)`
- `ModelSize(model string) (int, error)`
- `PullModel(model string) error`
- `PullIfMissing(model ...string) error`
- `GetDetails(model ...string) ([]ModelDetails, error)`
- `Version() (string, error)`
