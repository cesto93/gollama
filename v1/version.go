package gollama

func (c *Gollama) Version() (string, error) {
	type responseStr struct {
		Version string `json:"version"`
	}

	var resp responseStr
	c.apiGet("/api/version", &resp)

	return resp.Version, nil
}
