package main

import (
	"net"
	"net/http"
	"time"
)

type ClientLike interface {
	Get(url string) (ResponseLike, error)
}

type Client struct {
	client *http.Client
}

func NewClient() *Client {
	dialer := &net.Dialer{
		Timeout: 3 * time.Second,
	}

	client := &http.Client{
		Transport: &http.Transport{
			Dial:                  dialer.Dial,
			TLSHandshakeTimeout:   3 * time.Second,
			ResponseHeaderTimeout: 3 * time.Second,
		},
	}

	return &Client{client}
}

func (c *Client) Get(url string) (ResponseLike, error) {
	res, err := c.client.Get(url)
	return &Response{response: res}, err
}
