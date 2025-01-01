package gollama

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
)

// apiGet sends a GET request to the specified path on the Ollama server,
// and unmarshals the response into the given interface.
//
// The URL is built by joining the server address with the path.
//
// The Ollama server must respond with a JSON object that can be
// unmarshaled into the given interface.
//
// The Verbose flag is respected, and the URL is printed if it is set.
//
// The HTTPTimeout is used as the timeout for the HTTP request.
//
// If the request fails, or the response cannot be unmarshaled, an error
// is returned.
func (c *Gollama) apiGet(path string, v interface{}) error {
	url, _ := url.JoinPath(c.ServerAddr, path)
	if c.Verbose {
		fmt.Printf("Sending a request to GET %s\n", url)
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

// apiPost sends a POST request to the specified path on the Ollama server,
// and unmarshals the response into the given interface.
//
// The URL is built by joining the server address with the path.
//
// The Ollama server must respond with a JSON object that can be
// unmarshaled into the given interface.
//
// The Verbose flag is respected, and the URL is printed if it is set.
//
// If the request fails, or the response cannot be unmarshaled, an error
// is returned.
//
// The HTTPTimeout is used as the timeout for the HTTP request, except for
// requests to the /api/pull endpoint, which is given the PullTimeout.
func (c *Gollama) apiPost(path string, v interface{}, data interface{}) error {
	url, _ := url.JoinPath(c.ServerAddr, path)
	if c.Verbose {
		fmt.Printf("Sending a request to POST %s\n", url)
	}

	reqBytes, err := json.Marshal(data)
	if err != nil {
		if c.Verbose {
			fmt.Printf("Failed to marshal request data: %s\n", err)
		}
		return err
	}

	HTTPClient := &http.Client{
		Timeout: c.HTTPTimeout,
	}

	if path == "/api/pull" {
		HTTPClient.Timeout = c.PullTimeout
	}

	resp, err := HTTPClient.Post(url, "application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		if c.Verbose {
			fmt.Printf("Failed to send request: %s\n", err)
		}
		return err
	}
	defer resp.Body.Close()

	if c.Verbose {
		fmt.Printf("Response status code: %d\n", resp.StatusCode)
	}

	return json.NewDecoder(resp.Body).Decode(v)
}
