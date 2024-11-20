package gollama

func (c *Gollama) Version() (string, error) {
	var resp versionResponse
	c.apiGet("/api/version", &resp)

	return resp.Version, nil
}
