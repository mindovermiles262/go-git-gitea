package main

import (
	"fmt"
	"go-git-gitea/pkg/gitea"
	"go-git-gitea/pkg/gitlab"
)

func main() {
	fmt.Println("Hello World")
	gitea.Mirror()
	gitlabRepo := gitlab.GitlabRepository{
		RepoUrl:  "https://gitlab.com/my/repository.git",
		RepoName: "Repository",
	}
	// gitlabRepo.RepoUrl = "URL"
	gitlab.ListRepos(gitlabRepo)
}
