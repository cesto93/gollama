package gollama

import "github.com/xyproto/env"

func New(model string) *Config {
	if len(model) == 0 {
		model = defaultModel
	}

	oc := Config{
		ServerAddr:                env.Str("OLLAMA_HOST", "http://localhost:11434"),
		ModelName:                 env.Str("OLLAMA_MODEL", model),
		SeedOrNegative:            defaultFixedSeed,
		TemperatureIfNegativeSeed: 0.8,
		PullTimeout:               defaultPullTimeout,
		HTTPTimeout:               defaultHTTPTimeout,
		TrimSpace:                 true,
		Verbose:                   env.Bool("OLLAMA_VERBOSE"),
	}

	return &oc
}
