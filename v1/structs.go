package gollama

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
	Tools    []requestTool  `json:"tools,omitempty"`
	Options  requestOptions `json:"options"`
}

// Tool structs

type requestToolProperty struct {
	Type        string   `json:"type"`
	Description string   `json:"description"`
	Enum        []string `json:"enum"`
}

type requestToolParameters struct {
	Type       string                         `json:"type"`
	Properties map[string]requestToolProperty `json:"properties"`
	Required   []string                       `json:"required"`
}

type requestToolFunction struct {
	Name        string                `json:"name"`
	Description string                `json:"description"`
	Parameters  requestToolParameters `json:"parameters"`
}

type requestTool struct {
	Type     string              `json:"type"`
	Function requestToolFunction `json:"function"`
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
	Prompt       string   `json:"prompt"`
	VisionImages []string `json:"vision_images,omitempty"`
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
	Response       string            `json:"response"`
	ToolCalls      []GollamaToolCall `json:"tool_calls"`
	PromptTokens   int               `json:"prompt_tokens"`
	ResponseTokens int               `json:"response_tokens"`
}
