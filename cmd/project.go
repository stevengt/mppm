package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stevengt/mppm/config"
	"github.com/stevengt/mppm/util"
)

func init() {

	cobra.OnInitialize(
		func() {
			isPreviewCommand, _ = projectCmd.PersistentFlags().GetBool("preview")
			isCommitAllCommand, _ = projectCmd.Flags().GetBool("commit-all")
		},
	)

	projectCmd.PersistentFlags().BoolVarP(
		&isPreviewCommand,
		"preview",
		"p",
		false,
		"Shows what files will be affected without actually making changes.",
	)

	projectCmd.Flags().BoolVarP(
		&isCommitAllCommand,
		"commit-all",
		"c",
		false,
		"Equivalent to running 'mppm project extract; git add . -A; git commit -m '<commit message>'.",
	)

	rootCmd.AddCommand(projectCmd)

}

var isPreviewCommand bool
var isCommitAllCommand bool

var projectCmd = &cobra.Command{

	Use: "project",

	Short: "Provides utilities for managing a specific project.",

	Long: "Provides utilities for managing a specific project.",

	Args: cobra.MinimumNArgs(1),

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if cmd.Use != "init" {
			config.LoadMppmProjectConfig()
		}
	},

	Run: func(cmd *cobra.Command, args []string) {
		if isCommitAllCommand {
			commitAll(args[0])
		} else {
			cmd.Help()
		}
	},
}

func commitAll(commitMessage string) {

	err := extractAllCompressedFiles()
	if err != nil {
		util.ExitWithError(err)
	}

	err = util.ExecuteShellCommand("git", "add", ".", "-A")
	if err != nil {
		util.ExitWithError(err)
	}

	err = util.ExecuteShellCommand("git", "commit", "-m", commitMessage)
	if err != nil {
		util.ExitWithError(err)
	}

}
