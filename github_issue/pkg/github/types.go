package github

import (
	"time"

	"github.com/aniruddha2000/github_issue/pkg/app"
)

const (
	SearchIssueURL = "https://api.github.com/search/issues"
	RepoIssueURL   = "https://api.github.com/repos/"
)

type SearchIssueResult struct {
	TotalCount int `json:"total_count"`
	Items      []*IssueItems
}

type IssueItems struct {
	Number    int
	HTMLURL   string `json:"html_url"`
	Title     string
	State     string
	User      *User
	CreatedAt time.Time `json:"created_at"`
	Body      string
	Lables    []string
}

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type Issue struct {
	Title  string   `json:"title"`
	Body   string   `json:"body"`
	Lables []string `json:"lables"`
}

type Github interface {
	Search(c *app.Client)
	Create(c *app.Client)
}
