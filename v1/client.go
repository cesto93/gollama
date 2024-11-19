package gollama

func New(model string) *Gollama {
	if len(model) == 0 {
		model = defaultModel
	}

	oc := Gollama{
		ServerAddr:                getEnv("OLLAMA_HOST", "http://localhost:11434"),
		ModelName:                 getEnv("OLLAMA_MODEL", model),
		SeedOrNegative:            defaultFixedSeed,
		TemperatureIfNegativeSeed: 0.8,
		PullTimeout:               defaultPullTimeout,
		HTTPTimeout:               defaultHTTPTimeout,
		TrimSpace:                 true,
		Verbose:                   getEnv("OLLAMA_VERBOSE", "false") == "true",
	}

	return &oc
}
