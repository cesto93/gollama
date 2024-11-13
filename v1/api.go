package gollama

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

func (c *Gollama) apiGet(path string, v interface{}) error {
	url, _ := url.JoinPath(c.ServerAddr, path)
	if c.Verbose {
		fmt.Printf("Sending a request to %s\n", url)
	}
	HTTPClient := &http.Client{
		Timeout: c.HTTPTimeout,
	}
	resp, err := HTTPClient.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return json.NewDecoder(resp.Body).Decode(v)
}
