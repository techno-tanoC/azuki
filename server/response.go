package main

import (
	"io"
	"net/http"
)

type ResponseLike interface {
	Body() io.ReadCloser
	ContentLength() int64
}

type Response struct {
	response *http.Response
}

func (r *Response) Body() io.ReadCloser {
	return r.response.Body
}

func (r *Response) ContentLength() int64 {
	return r.response.ContentLength
}
