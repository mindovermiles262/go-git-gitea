package ggGitea

import (
	"fmt"

	"github.com/spf13/cobra"
)

var mirrorSource string

var mirrorCmd = &cobra.Command{
	Use:     "mirror <GIT URL> <GITEA PATH> [flags]",
	Short:   "Set up a Mirror",
	Long:    "Use this command to set up a new mirror between a source git repository and your Gitea instance",
	Args:    cobra.MinimumNArgs(2),
	Example: `gggitea mirror https://server.com/my/repo.git /gitea/path/repository`,
	Run: func(cmd *cobra.Command, args []string) {
		mirror()
	},
}

func mirror() error {
	env, err := getGiteaEnvironment()
	if err != nil {
		return err
	}

	fmt.Println("URL: " + env.baseUrl)
	fmt.Println("API Token: " + env.apiToken)
	fmt.Println("Mirror Source: " + env.mirrorSource)
	return nil
}

func init() {
	mirrorCmd.Flags().StringVarP(&mirrorSource, "mirror-source", "s", "", "The site that is hosting the repository you wish to mirror. Gitlab, Github, etc.")
	rootCmd.AddCommand(mirrorCmd)
}
