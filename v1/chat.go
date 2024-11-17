package gollama

import (
	"strings"
)

func (c *Gollama) Chat(in GollamaInput) (*GollamaResponse, error) {
	base64VisionImages := make([]string, 0)
	for _, image := range in.VisionImages {
		base64image, err := base64EncodeFile(image)
		if err != nil {
			return nil, err
		}
		base64VisionImages = append(base64VisionImages, base64image)
	}

	var (
		temperature float64
		seed        = c.SeedOrNegative
	)

	if seed < 0 {
		temperature = c.TemperatureIfNegativeSeed
	}

	messages := []message{}
	if c.SystemPrompt != "" {
		messages = append(messages, message{
			Role:    "system",
			Content: c.SystemPrompt,
		})
	}

	userMessage := message{
		Role:    "user",
		Content: in.Prompt,
	}

	if len(base64VisionImages) > 0 {
		userMessage.Images = base64VisionImages
	}

	messages = append(messages, userMessage)

	req := requestChat{
		Stream:   false,
		Model:    c.ModelName,
		Messages: messages,
		Options: requestOptions{
			Seed:        seed,        // set to -1 to make it random
			Temperature: temperature, // set to 0 together with a specific seed to make output reproducible
		},
	}

	if c.ContextLength != 0 {
		req.Options.ContextLength = c.ContextLength
	}

	var resp responseChat
	err := c.apiPost("/api/chat", &resp, req)
	if err != nil {
		return nil, err
	}

	res := &GollamaResponse{
		Role:           resp.Message.Role,
		Content:        resp.Message.Content,
		ToolCalls:      resp.Message.ToolCalls,
		PromptTokens:   resp.PromptEvalCount,
		ResponseTokens: resp.EvalCount,
	}

	if c.TrimSpace {
		res.Content = strings.TrimSpace(res.Content)
	}

	return res, nil
}
