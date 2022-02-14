package transport

import (
	"io"
	"net/http"
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
	resp, err := c.client.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
