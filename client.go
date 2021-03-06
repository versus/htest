package htest

import (
	"github.com/stretchr/testify/assert"
	"io"
	"net/http"
	"testing"
)

type (
	Client struct {
		handler http.Handler
		*testing.T
	}
)

func NewClient(t *testing.T) *Client {
	return &Client{T: t}
}

func (c Client) To(handler http.Handler) *Client {
	c.handler = handler
	return &c
}

func (c Client) ToFunc(handlerFunc http.HandlerFunc) *Client {
	c.handler = handlerFunc
	return &c
}

func (c Client) NewRequest(req *http.Request) *Request {
	return &Request{
		Request: req,
		Handler: c.handler,
		T:       c.T,
	}
}

func (c Client) request(method, path string, body io.Reader) *Request {
	req, err := http.NewRequest(method, path, body)
	assert.Nil(c.T, err)
	return c.NewRequest(req)
}

func (c Client) Get(path string) *Request {
	return c.request(GET, path, nil)
}

func (c Client) Head(path string) *Request {
	return c.request(HEAD, path, nil)
}

func (c Client) Trace(path string) *Request {
	return c.request(TRACE, path, nil)
}

func (c Client) Options(path string) *Request {
	return c.request(OPTIONS, path, nil)
}

func (c Client) Connect(path string) *Request {
	return c.request(CONNECT, path, nil)
}

func (c Client) Delete(path string) *Request {
	return c.request(DELETE, path, nil)
}

func (c Client) Post(path string, body io.Reader) *Request {
	return c.request(POST, path, body)
}

func (c Client) Put(path string, body io.Reader) *Request {
	return c.request(PUT, path, body)
}

func (c Client) Patch(path string, body io.Reader) *Request {
	return c.request(PATCH, path, body)
}
