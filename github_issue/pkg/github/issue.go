package github

import (
	"encoding/json"
	"fmt"
)

func NewIssue() *Issue {
	return &Issue{}
}

func (i *Issue) Search(items []string) error {
	return nil
}

func (i *Issue) Create(title, body string, label []string) ([]byte, error) {
	i.Title = title
	i.Body = body
	i.Lables = label
	ctn, err := json.Marshal(i)
	if err != nil {
		return nil, fmt.Errorf("json marshal err: %v", err)
	}

	fmt.Println(i)
	return ctn, nil
}
