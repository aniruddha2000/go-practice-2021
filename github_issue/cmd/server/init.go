package server

import (
	"flag"
	"log"
)

func Run() {
	var (
		operationTrype string
		resourceType   string
		repoName       string
		IssueTitle     string
		IssueBody      string
		IssueLabels    string
	)

	flag.StringVar(&resourceType, "resource", "issue",
		"Define the resource type which is going to fetch")
	flag.StringVar(&operationTrype, "operation", "search",
		"Define the operation type. e.g. - search, create, update, delete")
	flag.StringVar(&repoName, "repo", "",
		"Define the repo on which the operation is going to take place")
	flag.StringVar(&IssueTitle, "title", "",
		"Define the Issue title")
	flag.StringVar(&IssueBody, "body", "",
		"Define the issue body")
	flag.StringVar(&IssueLabels, "labels", "",
		"Define the issue labels")

	flag.Parse()

	c := NewClient()

	if repoName == "" {
		log.Fatal("Pease specify the repo name.")
	}
	if resourceType == "issue" {
		switch operationTrype {
		case "create":
			if IssueTitle != "" {
				c.CreateIssue(repoName, IssueTitle, IssueBody, IssueLabels)
			}
		case "search":
			c.SearchIssue(repoName)
		}
	}
}
