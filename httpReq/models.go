package httpReq

import (
	"net/http"
	"time"
)

type PostRequestConfig struct {
	Url            string
	Payload        interface{}
	ExpectedStatus int
	ResponseType   interface{}
	Headers        map[string]string
}

type GetRequestConfig struct {
	Url            string
	ExpectedStatus int
	ResponseType   interface{}
	Headers        map[string]string
	QueryParams    map[string]string
}

type Validate interface {
	ValidateSelf() error
}

type Client struct {
	*http.Client
}

func NewClient(timeOut time.Duration) *Client {
	return &Client{
		Client: &http.Client{
			Timeout: timeOut,
		},
	}
}
