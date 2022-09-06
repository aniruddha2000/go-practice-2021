package app

import (
	"net/http"
)

type Client struct {
	C http.Client
}

func NewClient() *Client {
	return &Client{}
}
