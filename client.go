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

func NewWithConfig(config Gollama) *Gollama {
	oc := Gollama{
		ServerAddr:                getEnv("OLLAMA_HOST", "http://localhost:11434"),
		ModelName:                 getEnv("OLLAMA_MODEL", defaultModel),
		SeedOrNegative:            defaultFixedSeed,
		TemperatureIfNegativeSeed: 0.8,
		PullTimeout:               defaultPullTimeout,
		HTTPTimeout:               defaultHTTPTimeout,
		TrimSpace:                 true,
		Verbose:                   getEnv("OLLAMA_VERBOSE", "false") == "true",
	}

	if oc.ServerAddr != config.ServerAddr {
		oc.ServerAddr = config.ServerAddr
	}

	if oc.ModelName != config.ModelName {
		oc.ModelName = config.ModelName
	}

	if oc.SeedOrNegative != config.SeedOrNegative {
		oc.SeedOrNegative = config.SeedOrNegative
	}

	if oc.TemperatureIfNegativeSeed != config.TemperatureIfNegativeSeed {
		oc.TemperatureIfNegativeSeed = config.TemperatureIfNegativeSeed
	}

	if oc.PullTimeout != config.PullTimeout {
		oc.PullTimeout = config.PullTimeout
	}

	if oc.HTTPTimeout != config.HTTPTimeout {
		oc.HTTPTimeout = config.HTTPTimeout
	}

	if oc.TrimSpace != config.TrimSpace {
		oc.TrimSpace = config.TrimSpace
	}

	if oc.ContextLength != config.ContextLength {
		oc.ContextLength = config.ContextLength
	}

	if oc.Verbose != config.Verbose {
		oc.Verbose = config.Verbose
	}

	return &oc
}
