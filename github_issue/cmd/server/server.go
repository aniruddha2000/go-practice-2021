package server

import (
	"net/http"

	"github.com/aniruddha2000/github_issue/pkg/github"
)

type Client struct {
	C  http.Client
	GH github.Github
}

func NewClient() *Client {
	return &Client{}
}
