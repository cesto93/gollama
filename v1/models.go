package gollama

type ModelInfo struct {
	Name string
	Size int
}

func (c *Gollama) ListModels() ([]ModelInfo, error) {
	type modelStr struct {
		Model string `json:"model"`
		Size  int    `json:"size"`
	}

	type responseStr struct {
		Models []modelStr `json:"models"`
	}

	var r responseStr
	c.apiGet("/api/tags", &r)

	var resp []ModelInfo
	for _, m := range r.Models {
		resp = append(resp, ModelInfo{
			Name: m.Model,
			Size: m.Size,
		})
	}

	return resp, nil
}

func (c *Gollama) HasModel(model string) (bool, error) {
	models, err := c.ListModels()
	if err != nil {
		return false, err
	}

	for _, m := range models {
		if m.Name == model || m.Name == model+":latest" {
			return true, nil
		}
	}

	return false, nil
}
