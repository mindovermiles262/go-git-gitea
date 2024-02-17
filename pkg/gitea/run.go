package ggGitea

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var GtUrl string
var GtApiKey string

var logger = logrus.Logger{
	Out: os.Stdout,
	Formatter: &logrus.TextFormatter{
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05.000",
	},
	Level: logrus.InfoLevel,
}

var rootCmd = &cobra.Command{
	Use:     "gggitea",
	Short:   "Go Git Gitea is the Swiss Army Knife for Gitea",
	Long:    "Git Gitea is the Swiss Army Knife for Gitea",
	Version: "0.1.0",
	Args:    cobra.MinimumNArgs(1),
}

func Run() {
	// Configure Viper to read from .env file
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("fatal error config file: %w", err))
	}

	logger.Info("Env .env Loaded")

	if err := rootCmd.Execute(); err != nil {
		logger.Fatal("error executing command")
		logger.Info(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&GtUrl, "url", "u", "https://gitea.com", "The URL to your Gitea instance")
	rootCmd.PersistentFlags().StringVarP(&GtApiKey, "token", "t", "", "Your Gitea API Key")
}
