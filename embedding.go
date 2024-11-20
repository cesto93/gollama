package gollama

// Embedding generates a vector embedding for a given string of text using the
// currently set model. The model must support the "embeddings" capability.
//
// The function will return an error if the model does not support the
// "embeddings" capability. The function will also return an error if the
// request fails.
//
// The function returns a slice of floats, representing the vector
// embedding of the input text.
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
