package gollama

import (
	"errors"
	"strings"
)

func (c *Gollama) Vision(prompt string, images []string) (string, error) {
	if len(images) == 0 {
		return "", errors.New("no images provided")
	}

	base64images := make([]string, len(images))
	for i, image := range images {
		base64image, err := base64EncodeFile(image)
		if err != nil {
			return "", err
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
		Content: prompt,
		Images:  images,
	})

	req := requestChat{
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
	err := c.apiPost("/api/vision", &resp, req)
	if err != nil {
		return "", err
	}

	response := resp.Message.Content

	if c.TrimSpace {
		response = strings.TrimSpace(response)
	}

	return response, nil
}
