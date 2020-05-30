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
			shouldUpdateLibraries, _ = projectCmd.Flags().GetBool("update-libraries")
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

	projectCmd.Flags().BoolVarP(
		&shouldUpdateLibraries,
		"update-libraries",
		"u",
		false,
		`Updates the library versions in the project config file to match the
current versions in the global config file.
To see the global current versions, run 'mppm library --list'.`,
	)

	rootCmd.AddCommand(projectCmd)

}

var isPreviewCommand bool
var isCommitAllCommand bool
var shouldUpdateLibraries bool

var projectCmd = &cobra.Command{

	Use: "project",

	Short: "Provides utilities for managing a specific project.",

	Long: "Provides utilities for managing a specific project.",

	Args: cobra.OnlyValidArgs,

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		shouldLoadMppmProjectConfig := cmd.Use != "init" && (cmd.Use != "project" || isCommitAllCommand || shouldUpdateLibraries)
		if shouldLoadMppmProjectConfig {
			config.LoadMppmProjectConfig()
		}
		if shouldUpdateLibraries {
			err := updateProjectLibraryGitCommitIds()
			if err != nil {
				util.ExitWithError(err)
			}
		}
	},

	Run: func(cmd *cobra.Command, args []string) {
		isCommandValid := isCommitAllCommand || shouldUpdateLibraries
		if !isCommandValid {
			cmd.Help()
		} else if isCommitAllCommand {
			if len(args) == 0 {
				util.ExitWithErrorMessage("Please provide a commit message.")
			}
			err := commitAll(args[0])
			if err != nil {
				util.ExitWithError(err)
			}
		}
	},
}

func commitAll(commitMessage string) (err error) {

	gitRepoFilePath := "."
	gitManager := util.NewGitManager(gitRepoFilePath)

	err = extractAllCompressedFiles()
	if err != nil {
		return
	}

	err = gitManager.Add(".", "-A")
	if err != nil {
		return
	}

	err = gitManager.Commit("-m", commitMessage)
	if err != nil {
		return
	}

	return

}

func updateProjectLibraryGitCommitIds() (err error) {
	config.LoadMppmGlobalConfig()
	config.MppmProjectConfig.Libraries = config.MppmGlobalConfig.Libraries
	err = config.MppmProjectConfig.SaveAsProjectConfig()
	return
}
