package gollama

import (
	"fmt"
	"strings"
)

func (c *Gollama) Chat(in GollamaInput) (*GollamaResponse, error) {
	var (
		temperature   float64
		seed          = c.SeedOrNegative
		contextLength = c.ContextLength
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

	base64VisionImages := make([]string, 0)
	for _, image := range in.VisionImages {
		base64image, err := base64EncodeFile(image)
		if err != nil {
			return nil, err
		}
		base64VisionImages = append(base64VisionImages, base64image)
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
			Seed:          seed,
			Temperature:   temperature,
			ContextLength: contextLength,
		},
	}

	if len(in.Tools) > 0 {
		req.Tools = in.Tools
	}

	if c.ContextLength != 0 {
		req.Options.ContextLength = c.ContextLength
	}

	var resp responseChat
	err := c.apiPost("/api/chat", &resp, req)
	if err != nil {
		return nil, err
	}

	if resp.Model != c.ModelName {
		return nil, fmt.Errorf("model don't found")
	}

	out := &GollamaResponse{
		Role:           resp.Message.Role,
		Content:        resp.Message.Content,
		ToolCalls:      resp.Message.ToolCalls,
		PromptTokens:   resp.PromptEvalCount,
		ResponseTokens: resp.EvalCount,
	}

	if c.TrimSpace {
		out.Content = strings.TrimSpace(out.Content)
	}

	return out, nil
}
