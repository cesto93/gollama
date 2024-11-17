package gollama

import (
	"errors"
	"strings"
)

func (c *Gollama) Vision(in GollamaInput) (*GollamaResponse, error) {
	if len(in.VisionImages) == 0 {
		return nil, errors.New("no images provided")
	}

	base64images := make([]string, len(in.VisionImages))
	for i, image := range in.VisionImages {
		base64image, err := base64EncodeFile(image)
		if err != nil {
			return nil, err
		}
		base64images[i] = base64image
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

	messages = append(messages, message{
		Role:    "user",
		Content: in.Prompt,
		Images:  base64images,
	})

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
		Response:       resp.Message.Content,
		ToolCalls:      resp.Message.ToolCalls,
		PromptTokens:   resp.PromptEvalCount,
		ResponseTokens: resp.EvalCount,
	}

	if c.TrimSpace {
		res.Response = strings.TrimSpace(res.Response)
	}

	return res, nil
}
