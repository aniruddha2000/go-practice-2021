package github

const (
	RepoIssueURL = "https://api.github.com/repos/"
)

type state string

const (
	OPEN   state = "open"
	CLOSED state = "closed"
)

type User struct {
	Login   string
	HTMLURL string `json:"html_url"`
}

type Issue struct {
	Title  string   `json:"title"`
	Body   string   `json:"body"`
	Lables []string `json:"lables"`
	User   *User    `json:"user,omitempty"`
	State  state    `json:"state,omitempty"`
}

type Github interface {
	Search([]string) error
	Create(string, string, []string) ([]byte, error)
}
