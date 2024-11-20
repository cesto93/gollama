package gollama

import (
	"errors"
	"fmt"
)

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

	return false, errors.New("model not found")
}

func (c *Gollama) ModelSize(model string) (int, error) {
	models, err := c.ListModels()
	if err != nil {
		return 0, err
	}

	for _, m := range models {
		if m.Name == model || m.Name == model+":latest" {
			return m.Size, nil
		}
	}

	return 0, nil
}

func (c *Gollama) PullModel(model string) error {
	type requestStr struct {
		Model  string `json:"model"`
		Stream bool   `json:"stream"`
	}

	type responseStr struct {
		Status string `json:"status"`
	}

	req := requestStr{
		Model:  model,
		Stream: false,
	}

	var resp responseStr

	c.apiPost("/api/pull", &resp, req)

	if resp.Status != "success" {
		return fmt.Errorf("failed to pull model %s", model)
	}
	return nil
}

func (c *Gollama) PullIfMissing(model ...string) error {
	if len(model) == 0 {
		model = []string{c.ModelName}
	}

	for _, m := range model {
		hasModel, err := c.HasModel(m)
		if err != nil {
			return err
		}

		if !hasModel {
			return c.PullModel(m)
		}
	}

	return nil
}

func (c *Gollama) GetDetails(model ...string) ([]ModelDetails, error) {
	if len(model) == 0 {
		model = []string{c.ModelName}
	}

	ret := make([]ModelDetails, 0)

	for _, m := range model {
		req := getDetails{
			Model: m,
		}

		var resp ModelDetails
		err := c.apiPost("/api/show", &resp, req)
		if err != nil {
			return nil, err
		}

		ret = append(ret, resp)
	}

	return ret, nil
}
