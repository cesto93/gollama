package gollama

// Embedding gets the embeddings of a given input prompt
//
// The function takes a ChatInput object as input and returns a slice of float64
// representing the embeddings of the input, or an error if something went wrong.
//
// The function will return an error if the model is not found.
func (c *Gollama) Embedding(prompt string) ([]float64, error) {
	req := embeddingsRequest{
		Model:  c.ModelName,
		Prompt: prompt,
	}

	var resp embeddingsResponse
	err := c.apiPost("/api/embeddings", &resp, req)
	if err != nil {
		return nil, err
	}

	return resp.Embedding, nil
}
