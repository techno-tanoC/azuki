package main

import (
	"io"
	"net"
	"net/http"
	"time"
)

type ResponseLike interface {
	io.ReadCloser
	ContentLength() int64
}

type Response struct {
	response *http.Response
}

func (r *Response) Read(p []byte) (int, error) {
	return r.response.Body.Read(p)
}

func (r *Response) Close() error {
	return r.response.Body.Close()
}

func (r *Response) ContentLength() int64 {
	return r.response.ContentLength
}

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
