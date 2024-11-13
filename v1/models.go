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

func (c *Gollama) HasModel(model string) (bool, error) {
	models, err := c.ListModels()
	if err != nil {
		return false, err
	}

	for _, m := range models {
		if m == model || m == model+":latest" {
			return true, nil
		}
	}

	return false, nil
}
