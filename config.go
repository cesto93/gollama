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

func (c *Gollama) SetSeed(seed int) *Gollama {
	c.SeedOrNegative = seed
	return c
}

func (c *Gollama) SetRandomSeed() *Gollama {
	c.SeedOrNegative = -1
	return c
}

func (c *Gollama) SetTemperature(temperature float64) *Gollama {
	c.TemperatureIfNegativeSeed = temperature
	return c
}

func (c *Gollama) SetContextLength(contextLength int64) *Gollama {
	c.ContextLength = contextLength
	return c
}

func getEnv(key, defaultValue string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return defaultValue
}
