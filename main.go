package main

import (
	"fmt"
	gt "go-git-gitea/pkg/gitea"
	gl "go-git-gitea/pkg/gitlab"
)

func main() {
	fmt.Println("Hello World")

	gt.Mirror()

	gitlabRepo := gl.GitlabRepository{
		RepoUrl:  "https://gitlab.com/my/repository.git",
		RepoName: "Repository",
	}

	gl.ListRepos(gitlabRepo)

	// ctx := context.Background()
	// ghClient := gh.NewClient()
	// repos, resp, err := ghClient.Repositories.List(ctx, "andyduss", &github.RepositoryListOptions{})
	// fmt.Println(repos)
	// fmt.Println(resp)
	// fmt.Println(err)
}
