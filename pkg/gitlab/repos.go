package goGitGiteaGitlab

import "fmt"

type GitlabRepository struct {
	RepoUrl  string
	RepoName string
}

func ListRepos(glr GitlabRepository) {
	fmt.Printf("URL %s\n", glr.RepoUrl)
	fmt.Printf("Name %s\n", glr.RepoName)
}
