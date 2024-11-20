package gollama

// ModelDetails

type getDetails struct {
	Model string `json:"model"`
}

type ModelDetails struct {
	License    string `json:"license"`
	Modelfile  string `json:"modelfile"`
	Parameters string `json:"parameters"`
	Template   string `json:"template"`
	Details    struct {
		ParentModel       string   `json:"parent_model"`
		Format            string   `json:"format"`
		Family            string   `json:"family"`
		Families          []string `json:"families"`
		ParameterSize     string   `json:"parameter_size"`
		QuantizationLevel string   `json:"quantization_level"`
	} `json:"details"`
	ModelInfo struct {
		GeneralArchitecture               string   `json:"general.architecture"`
		GeneralBasename                   string   `json:"general.basename"`
		GeneralFileType                   int      `json:"general.file_type"`
		GeneralFinetune                   string   `json:"general.finetune"`
		GeneralLanguages                  []string `json:"general.languages"`
		GeneralLicense                    string   `json:"general.license"`
		GeneralParameterCount             int64    `json:"general.parameter_count"`
		GeneralQuantizationVersion        int      `json:"general.quantization_version"`
		GeneralSizeLabel                  string   `json:"general.size_label"`
		GeneralTags                       []string `json:"general.tags"`
		GeneralType                       string   `json:"general.type"`
		LlamaAttentionHeadCount           int      `json:"llama.attention.head_count"`
		LlamaAttentionHeadCountKv         int      `json:"llama.attention.head_count_kv"`
		LlamaAttentionLayerNormRmsEpsilon float64  `json:"llama.attention.layer_norm_rms_epsilon"`
		LlamaBlockCount                   int      `json:"llama.block_count"`
		LlamaContextLength                int      `json:"llama.context_length"`
		LlamaEmbeddingLength              int      `json:"llama.embedding_length"`
		LlamaFeedForwardLength            int      `json:"llama.feed_forward_length"`
		LlamaRopeDimensionCount           int      `json:"llama.rope.dimension_count"`
		LlamaRopeFreqBase                 int      `json:"llama.rope.freq_base"`
		LlamaVocabSize                    int      `json:"llama.vocab_size"`
		TokenizerGgmlBosTokenID           int      `json:"tokenizer.ggml.bos_token_id"`
		TokenizerGgmlEosTokenID           int      `json:"tokenizer.ggml.eos_token_id"`
		TokenizerGgmlMerges               any      `json:"tokenizer.ggml.merges"`
		TokenizerGgmlModel                string   `json:"tokenizer.ggml.model"`
		TokenizerGgmlPre                  string   `json:"tokenizer.ggml.pre"`
		TokenizerGgmlTokenType            any      `json:"tokenizer.ggml.token_type"`
		TokenizerGgmlTokens               any      `json:"tokenizer.ggml.tokens"`
	} `json:"model_info"`
	ModifiedAt string `json:"modified_at"`
}

// Embedding

type embedding struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}

type responseEmbedding struct {
	Embedding []float64 `json:"embedding"`
}

// Chat

type message struct {
	Role    string   `json:"role"`
	Content string   `json:"content"`
	Images  []string `json:"images,omitempty"`
}

type requestOptions struct {
	Seed          int     `json:"seed"`
	Temperature   float64 `json:"temperature"`
	ContextLength int64   `json:"context_length,omitempty"`
}

type requestChat struct {
	Model    string         `json:"model"`
	Stream   bool           `json:"stream"`
	Messages []message      `json:"messages"`
	Tools    []GollamaTool  `json:"tools,omitempty"`
	Options  requestOptions `json:"options"`
}

// Tool structs

type GollamaToolProperty struct {
	Type        string   `json:"type"`
	Description string   `json:"description"`
	Enum        []string `json:"enum"`
}

type GollamaToolParameters struct {
	Type       string                         `json:"type"`
	Properties map[string]GollamaToolProperty `json:"properties"`
	Required   []string                       `json:"required"`
}

type GollamaToolFunction struct {
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Parameters  GollamaToolParameters `json:"parameters"`
}

type GollamaTool struct {
	Type     string              `json:"type"`
	Function GollamaToolFunction `json:"function"`
}

// ResponseChat is the response from the Ollama API

type responseMessage struct {
	Role      string            `json:"role"`
	Content   string            `json:"content"`
	ToolCalls []GollamaToolCall `json:"tool_calls"`
}

type responseChat struct {
	Model              string          `json:"model"`
	CreatedAt          string          `json:"created_at"`
	Message            responseMessage `json:"message"`
	DoneReason         string          `json:"done_reason"`
	Done               bool            `json:"done"`
	TotalDuration      int64           `json:"total_duration,omitempty"`
	LoadDuration       int64           `json:"load_duration,omitempty"`
	PromptEvalCount    int             `json:"prompt_eval_count,omitempty"`
	PromptEvalDuration int64           `json:"prompt_eval_duration,omitempty"`
	EvalCount          int             `json:"eval_count,omitempty"`
	EvalDuration       int64           `json:"eval_duration,omitempty"`
}

// Input structs

type GollamaInput struct {
	Prompt       string        `json:"prompt"`
	VisionImages []string      `json:"vision_images,omitempty"`
	Tools        []GollamaTool `json:"tools,omitempty"`
}

// Output structs

type GollamaToolCallFunction struct {
	Name      string         `json:"name"`
	Arguments map[string]any `json:"arguments"`
}

type GollamaToolCall struct {
	Function GollamaToolCallFunction `json:"function"`
}

type GollamaResponse struct {
	Role           string            `json:"role"`
	Content        string            `json:"content"`
	ToolCalls      []GollamaToolCall `json:"tool_calls"`
	PromptTokens   int               `json:"prompt_tokens"`
	ResponseTokens int               `json:"response_tokens"`
}
