package ggGitea

import (
	"errors"

	"github.com/spf13/viper"
)

type environment struct {
	baseUrl      string // The Base URL to the destination Gitea Server
	apiToken     string // The Gitea API Token
	mirrorSource string // The source site. Gitlab, Github, etc.
}

func getGiteaEnvironment() (*environment, error) {
	e := environment{}
	// Base URL. Default comes from Cobra CMD
	viper.BindPFlag("gt_base_url", rootCmd.Flags().Lookup("url"))
	url := viper.GetString("gt_base_url")
	logger.Infof("Gitea URL is %s", url)

	// API Key
	viper.BindPFlag("gt_api_token", rootCmd.Flags().Lookup("token"))
	apiToken := viper.GetString("gt_api_token")
	if apiToken == "" {
		logger.Fatal("API Token not set. Exiting.")
		return nil, errors.New("gitea api token is required")
	}

	// Mirror Source.
	var ms string
	ms = mirrorSource
	if ms == "" {
		ms = viper.GetString("mirror_source")
	}

	// Default
	if ms == "" {
		ms = "github"
	}
	logger.Infof("Mirror Source is %s", ms)

	e.baseUrl = url
	e.apiToken = apiToken
	e.mirrorSource = ms
	return &e, nil
}
