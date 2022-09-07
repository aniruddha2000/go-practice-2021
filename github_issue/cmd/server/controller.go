package server

import (
	"bytes"
	"fmt"
	"log"
	"net/http"
	"os"

	gh "github.com/aniruddha2000/github_issue/pkg/github"
	"github.com/aniruddha2000/github_issue/pkg/utils"
)

func (c *Client) CreateIssue() {
	c.GH = gh.NewIssue()

	ctn, err := c.GH.Create("hello", "world", []string{"go", "pher"})
	if err != nil {
		log.Fatal(err)
	}

	url := gh.RepoIssueURL + os.Args[1]
	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(ctn))
	if err != nil {
		log.Fatal(err)
	}
	req.Header = http.Header{
		"Accept":        {"application/vnd.github+json"},
		"Authorization": {fmt.Sprintf("Bearer %s", utils.GetToken())},
	}

	res, err := c.C.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	if res.StatusCode == http.StatusCreated {
		log.Println("Content created")
	}
}
