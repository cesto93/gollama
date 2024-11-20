package gollama

// New is a constructor for Gollama that takes a model name as an argument. Any
// unset fields will be populated with default values.
//
// This is useful for cases where you want to create a Gollama with a specific
// model name, but don't want to specify every single field.
//
// If the provided model name is empty, the default model name will be used.
//
// The following environment variables can be set to override the default values
// of the fields:
//
// OLLAMA_HOST: the address of the Ollama server
// OLLAMA_MODEL: the model name
// OLLAMA_VERBOSE: whether to print debug information
//
// The returned Gollama can be used to call other methods on the Ollama API.
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
		SystemPrompt:              "",
	}

	return &oc
}

// NewWithConfig is a constructor for Gollama that takes a pre-populated Gollama
// as an argument. Any unset fields will be populated with default values.
//
// This is useful for cases where you want to create a Gollama with some custom
// configuration, but don't want to specify every single field.
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
		SystemPrompt:              "",
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

	if oc.SystemPrompt != config.SystemPrompt {
		oc.SystemPrompt = config.SystemPrompt
	}

	return &oc
}
