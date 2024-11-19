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

func NewWithConfig(config *Gollama) *Gollama {
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

	if config.ServerAddr != "" {
		oc.ServerAddr = config.ServerAddr
	}

	if config.ModelName != "" {
		oc.ModelName = config.ModelName
	}

	if config.SeedOrNegative != defaultFixedSeed {
		oc.SeedOrNegative = config.SeedOrNegative
	}

	if config.TemperatureIfNegativeSeed != 0.8 {
		oc.TemperatureIfNegativeSeed = config.TemperatureIfNegativeSeed
	}

	if config.PullTimeout != defaultPullTimeout {
		oc.PullTimeout = config.PullTimeout
	}

	if config.HTTPTimeout != defaultHTTPTimeout {
		oc.HTTPTimeout = config.HTTPTimeout
	}

	if config.TrimSpace != true {
		oc.TrimSpace = config.TrimSpace
	}

	if config.Verbose != false {
		oc.Verbose = config.Verbose
	}

	return &oc
}
