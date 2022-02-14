package transport

import (
	"io"
	"net/http"
	"strings"
)

type Client struct {
	client http.Client
}

func NewClient() *Client {
	return &Client{
		client: http.Client{},
	}
}

func (c *Client) SendRequest(url string) ([]byte, error) {
	if !strings.HasPrefix(url, "https://") && !strings.HasPrefix(url, "http://") {
		url = "https://" + url
	}
	resp, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
