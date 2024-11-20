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
