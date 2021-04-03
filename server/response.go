package main

import (
	"io"
	"net/http"
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
