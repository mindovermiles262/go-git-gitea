package main

import (
	"context"
	"fmt"
	gt "go-git-gitea/pkg/gitea"
	gh "go-git-gitea/pkg/github"
	gl "go-git-gitea/pkg/gitlab"

	"github.com/google/go-github/github"
)

func main() {
	ctx := context.Background()
	fmt.Println("Hello World")

	gt.Mirror()

	gitlabRepo := gl.GitlabRepository{
		RepoUrl:  "https://gitlab.com/my/repository.git",
		RepoName: "Repository",
	}

	gl.ListRepos(gitlabRepo)

	ghClient := gh.NewClient()
	repos, resp, err := ghClient.Repositories.List(ctx, "andyduss", &github.RepositoryListOptions{})
	fmt.Println(repos)
	fmt.Println(resp)
	fmt.Println(err)
}
