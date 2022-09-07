package github

import (
	"encoding/json"
	"fmt"
)

func NewIssue() *Issue {
	return &Issue{}
}

func (i *Issue) Search(res []byte) error {
	var issues []Issue
	err := json.Unmarshal(res, &issues)
	if err != nil {
		return fmt.Errorf("json unmarshal: %v", err)
	}

	for _, issue := range issues {
		fmt.Printf("%d - %s - %s\n", issue.Number, issue.Title, issue.State)
	}

	return nil
}

func (i *Issue) Create(title, body string, label []string) ([]byte, error) {
	i.Title = title
	i.Body = body
	ctn, err := json.Marshal(i)
	if err != nil {
		return nil, fmt.Errorf("json marshal err: %v", err)
	}
	return ctn, nil
}
