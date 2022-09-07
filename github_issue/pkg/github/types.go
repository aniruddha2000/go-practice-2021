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
	Number       int      `json:"number,omitempty"`
	Title        string   `json:"title"`
	Body         string   `json:"body"`
	Labels       []Labels `json:"omitempty"`
	User         *User    `json:"user,omitempty"`
	State        state    `json:"state,omitempty"`
}

type Labels struct {
	Name string `json:"name"`
}

type Github interface {
	Search([]byte) error
	Create(string, string, []string) ([]byte, error)
}
