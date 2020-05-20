package cmd

import (
	"errors"
	"fmt"
	"os"

	"github.com/spf13/cobra"

	"github.com/stevengt/mppm/versioning"
	"github.com/stevengt/mppm/versioning/project"
)

func init() {
	rootCmd.AddCommand(gitCmd)
}

var gitCmd = &cobra.Command{
	Use:   "git",
	Short: "Run any 'git' command with automated management of binary files using 'git-lfs'.",
	Long: `Run any 'git' command with automated management of binary files using 'git-lfs'.
			Before the command is executed, '.als' files will be extracted into readable
			'.xml' files.`,
	DisableFlagParsing: true,
	Run: func(cmd *cobra.Command, args []string) {
		runGitCommand(args)
	},
}

func runGitCommand(args []string) {
	config, err := project.LoadMppmConfig()

	if err == nil {
		switch config.ProjectType {
		case project.Ableton:
			versioner := &versioning.AbletonVersioner{}
			err = versioner.Git(args...)
		default:
			err = errors.New("A valid project type was not found in the config file '" + project.ConfigFileName + "'.")
		}
	}

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
