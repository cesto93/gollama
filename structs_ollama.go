package gollama

// Version

type versionResponse struct {
	Version string `json:"version"`
}

// Tags

type tagsResponse struct {
	Models []ModelInfo `json:"models"`
}

// Pull

type pullRequest struct {
	Model  string `json:"model"`
	Stream bool   `json:"stream"`
}

type pullResponse struct {
	Status string `json:"status"`
}

// Show

type showRequest struct {
	Model string `json:"model"`
}

// Embeddings

type embeddingsRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
}
type embeddingsResponse struct {
	Embedding []float64 `json:"embedding"`
}

// Chat

type chatMessage struct {
	Role    string   `json:"role"`
	Content string   `json:"content"`
	Images  []string `json:"images,omitempty"`
}

type chatOptionsRequest struct {
	Seed          int     `json:"seed"`
	Temperature   float64 `json:"temperature"`
	TopK          int     `json:"top_k"`
	TopP          float64 `json:"top_p"`
	ContextLength int64   `json:"num_ctx,omitempty"`
}

type chatRequest struct {
	Model    string             `json:"model"`
	Stream   bool               `json:"stream"`
	Messages []chatMessage      `json:"messages"`
	Tools    *[]Tool            `json:"tools,omitempty"`
	Format   *StructuredFormat  `json:"format,omitempty"`
	Options  chatOptionsRequest `json:"options"`
}

// ResponseChat is the response from the Ollama API

type messageResponse struct {
	Role      string     `json:"role"`
	Content   string     `json:"content"`
	ToolCalls []ToolCall `json:"tool_calls"`
}

type chatResponse struct {
	Model              string          `json:"model"`
	CreatedAt          string          `json:"created_at"`
	Message            messageResponse `json:"message"`
	DoneReason         string          `json:"done_reason"`
	Done               bool            `json:"done"`
	TotalDuration      int64           `json:"total_duration,omitempty"`
	LoadDuration       int64           `json:"load_duration,omitempty"`
	PromptEvalCount    int             `json:"prompt_eval_count,omitempty"`
	PromptEvalDuration int64           `json:"prompt_eval_duration,omitempty"`
	EvalCount          int             `json:"eval_count,omitempty"`
	EvalDuration       int64           `json:"eval_duration,omitempty"`
}
