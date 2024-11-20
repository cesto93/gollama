package gollama

func (c *Gollama) Embedding(in GollamaInput) ([]float64, error) {
	type embeddingsRequest struct {
		Model  string `json:"model"`
		Prompt string `json:"prompt"`
	}

	req := embeddingsRequest{
		Model:  c.ModelName,
		Prompt: in.Prompt,
	}

	var resp responseEmbedding
	err := c.apiPost("/api/embeddings", &resp, req)
	if err != nil {
		return nil, err
	}

	return resp.Embedding, nil
}
