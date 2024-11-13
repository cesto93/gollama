package gollama

func (c *Gollama) ListModels() ([]string, error) {
	type modelStr struct {
		Model string `json:"model"`
	}

	type responseStr struct {
		Models []modelStr `json:"models"`
	}

	var r responseStr
	c.apiGet("/api/tags", &r)

	var resp []string
	for _, m := range r.Models {
		resp = append(resp, m.Model)
	}

	return resp, nil
}
