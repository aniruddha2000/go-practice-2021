package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/aniruddha2000/github_issue/pkg/app"
)

func searchIssues(items []string) (*IssueSearchResult, error) {
	q := url.QueryEscape(strings.Join(items, " "))
	resp, err := http.Get(IssueURL + "?q=" + q)
	log.Println(IssueURL + "?q=" + q)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	ctn, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("ioutil read: %v", err)
	}

	var result IssueSearchResult
	if err = json.Unmarshal(ctn, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

func createIssue(c *app.Client, issue Issue, path string) error {
	ctn, err := json.Marshal(issue)
	if err != nil {
		return fmt.Errorf("json marshal err: %v", err)
	}

	url := CreateIssueURL + path

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(ctn))
	if err != nil {
		return fmt.Errorf("error making request: %v", err)
	}
	req.Header = http.Header{
		"Accept":        {"application/vnd.github+json"},
		"Authorization": {fmt.Sprintf("Bearer %s", getToken())},
	}

	res, err := c.C.Do(req)
	if err != nil {
		return nil
	}

	if res.StatusCode == http.StatusCreated {
		return nil
	}

	return fmt.Errorf("error creating issue: status code: %d", res.StatusCode)
}

func getToken() string {
	return os.Getenv("GITHUB_TOKEN")
}
