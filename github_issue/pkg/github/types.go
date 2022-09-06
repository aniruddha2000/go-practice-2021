package github

import "time"

const (
	IssueURL       = "https://api.github.com/search/issues"
	CreateIssueURL = "https://api.github.com/repos/"
)

type IssueSearchResult struct {
	TotalCount int `json:"total_count"`
	Items      []*SearchIssue
}

type SearchIssue struct {
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
