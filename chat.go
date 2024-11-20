package gollama

import (
	"fmt"
	"strings"
)

// Chat starts a conversation with the LLaMA model.
//
// The function will append the SystemPrompt to the conversation if it is set.
//
// The function will set the temperature to TemperatureIfNegativeSeed if the seed is negative.
//
// The SeedOrNegative, TemperatureIfNegativeSeed, ContextLength and TrimSpace fields of the Gollama object will be used
// to set the seed, temperature, context length and whether to trim the output respectively.
//
// The function will also pass through the Tools field of the input to the API.
//
// The function will return the content of the message from the model, the content of the message will be trimmed if
// TrimSpace is true.
//
// The function will return an error if the model is not found.
func (c *Gollama) Chat(in ChatInput) (*ChatResponse, error) {
	var (
		temperature   float64
		seed          = c.SeedOrNegative
		contextLength = c.ContextLength
	)

	if seed < 0 {
		temperature = c.TemperatureIfNegativeSeed
	}

	messages := []chatMessage{}
	if c.SystemPrompt != "" {
		messages = append(messages, chatMessage{
			Role:    "system",
			Content: c.SystemPrompt,
		})
	}

	userMessage := chatMessage{
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

	req := chatRequest{
		Stream:   false,
		Model:    c.ModelName,
		Messages: messages,
		Options: chatOptionsRequest{
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

	var resp chatResponse
	err := c.apiPost("/api/chat", &resp, req)
	if err != nil {
		return nil, err
	}

	if resp.Model != c.ModelName {
		return nil, fmt.Errorf("model don't found")
	}

	out := &ChatResponse{
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
