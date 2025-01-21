package gollama

import (
	"os"
	"time"
)

type Gollama struct {
	ServerAddr                string
	ModelName                 string
	SeedOrNegative            int
	TemperatureIfNegativeSeed float64
	TopK                      int
	TopP                      float64
	PullTimeout               time.Duration
	HTTPTimeout               time.Duration
	TrimSpace                 bool
	Verbose                   bool
	ContextLength             int64
	SystemPrompt              string
}

const (
	defaultModel       = "llama3.2"
	defaultHTTPTimeout = 10 * time.Minute // per HTTP request to Ollama
	defaultFixedSeed   = 256              // for when generated output should not be random, but have temperature 0 and a specific seed
	defaultPullTimeout = 48 * time.Hour   // pretty generous, in case someone has a poor connection
	mimeJSON           = "application/json"
)

func (c *Gollama) SetHTTPTimeout(timeout time.Duration) *Gollama {
	c.HTTPTimeout = timeout
	return c
}

// SetSeed Sets the random number seed to use for generation. Setting this to a specific number will make the model generate the same text for the same prompt. (Default: 0)
func (c *Gollama) SetSeed(seed int) *Gollama {
	c.SeedOrNegative = seed
	return c
}

func (c *Gollama) SetRandomSeed() *Gollama {
	c.SeedOrNegative = -1
	return c
}

// SetTemperature The temperature of the model. Increasing the temperature will make the model answer more creatively. (Default: 0.8)
func (c *Gollama) SetTemperature(temperature float64) *Gollama {
	c.TemperatureIfNegativeSeed = temperature
	return c
}

// SetTopK Reduces the probability of generating nonsense. A higher value (e.g. 100) will give more diverse answers, while a lower value (e.g. 10) will be more conservative. (Default: 40)
func (c *Gollama) SetTopK(topK int) *Gollama {
	c.TopK = topK
	return c
}

// SetTopP Works together with top-k. A higher value (e.g., 0.95) will lead to more diverse text, while a lower value (e.g., 0.5) will generate more focused and conservative text. (Default: 0.9)
func (c *Gollama) SetTopP(topP float64) *Gollama {
	c.TopP = topP
	return c
}

// SetContextLength Sets the size of the context window used to generate the next token. (Default: 2048)
func (c *Gollama) SetContextLength(contextLength int64) *Gollama {
	c.ContextLength = contextLength
	return c
}

func (c *Gollama) SetSystemPrompt(prompt string) *Gollama {
	c.SystemPrompt = prompt
	return c
}

func getEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
