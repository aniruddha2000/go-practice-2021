package github

import (
	"fmt"
	"log"
	"os"

	"github.com/aniruddha2000/github_issue/pkg/app"
)

func SearchIssues(c *app.Client) {
	var issues SearchIssueResult
	err := issues.search(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Issue count: %d\n", issues.TotalCount)
	for _, item := range issues.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
	}
}

func CreateIssue(c *app.Client) {
	var issue = Issue{
		Title:  "Test Issue by go",
		Body:   `Hey it's a test issue by go program on Ubuntu 2204`,
		Lables: []string{"test", "go", "github-rest"},
	}

	if err := issue.create(c, os.Args[1]); err != nil {
		log.Fatal(err)
	}
}
