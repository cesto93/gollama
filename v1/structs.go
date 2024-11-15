package gollama

type message struct {
	Role    string   `json:"role"`
	Content string   `json:"content"`
	Images  []string `json:"images"`
}

type requestOptions struct {
	Seed          int     `json:"seed"`
	Temperature   float64 `json:"temperature"`
	ContextLength int64   `json:"context_length"`
}

type requestChat struct {
	Model    string         `json:"model"`
	Messages []message      `json:"messages"`
	Options  requestOptions `json:"options"`
}

type responseMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
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
