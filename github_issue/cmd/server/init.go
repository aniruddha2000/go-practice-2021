package server

import (
	"github.com/aniruddha2000/github_issue/pkg/app"
	gh "github.com/aniruddha2000/github_issue/pkg/github"
)

func Run() {
	c := app.NewClient()

	gh.Create(c)
}
