package server

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	gh "github.com/aniruddha2000/github_issue/pkg/github"
	"github.com/aniruddha2000/github_issue/pkg/utils"
)

func (c *Client) SearchIssue(name string) {
	url := gh.RepoIssueURL + name
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		log.Fatal(err)
	}
	res, err := c.C.Do(req)
	if err != nil {
		log.Fatal(err)
	}

	c.GH = gh.NewIssue()

	resBody, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	if err := c.GH.Search(resBody); err != nil {
		log.Fatal(err)
	}
}

func (c *Client) CreateIssue(repoName, title, body, labels string) {
	c.GH = gh.NewIssue()

	splitLabels := strings.Split(labels, ",")

	ctn, err := c.GH.Create(title, body, splitLabels)
	if err != nil {
		log.Fatal(err)
	}

	url := gh.RepoIssueURL + repoName
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
