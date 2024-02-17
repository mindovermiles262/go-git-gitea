package goGitGiteaGithub

import (
	"fmt"

	"github.com/google/go-github/github"
)

func NewClient() *github.Client {
	fmt.Println("Github")
	return github.NewClient(nil)
}
