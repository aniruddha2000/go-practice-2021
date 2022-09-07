package github

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/aniruddha2000/github_issue/pkg/app"
)

func (i *SearchIssueResult) search(items []string) error {
	q := url.QueryEscape(strings.Join(items, " "))
	resp, err := http.Get(SearchIssueURL + "?q=" + q)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("search query failed: %s", resp.Status)
	}

	ctn, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("ioutil read: %v", err)
	}

	if err = json.Unmarshal(ctn, i); err != nil {
		return err
	}

	return nil
}

func (i *Issue) create(c *app.Client, path string) error {
	ctn, err := json.Marshal(i)
	if err != nil {
		return fmt.Errorf("json marshal err: %v", err)
	}

	url := RepoIssueURL + path

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
