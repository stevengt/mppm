package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stevengt/mppm/util"
)

func init() {

	cobra.OnInitialize(
		func() {
			isPreviewCommand, _ = ProjectCmd.PersistentFlags().GetBool("preview")
			isCommitAllCommand, _ = ProjectCmd.Flags().GetBool("commit-all")
			shouldUpdateLibraries, _ = ProjectCmd.Flags().GetBool("update-libraries")
		},
	)

	ProjectCmd.PersistentFlags().BoolVarP(
		&isPreviewCommand,
		"preview",
		"p",
		false,
		"Shows what files will be affected without actually making changes.",
	)

	ProjectCmd.Flags().BoolVarP(
		&isCommitAllCommand,
		"commit-all",
		"c",
		false,
		"Equivalent to running 'mppm project extract; git add . -A; git commit -m '<commit message>'.",
	)

	ProjectCmd.Flags().BoolVarP(
		&shouldUpdateLibraries,
		"update-libraries",
		"u",
		false,
		`Updates the library versions in the project config file to match the
current versions in the global config file.
To see the global current versions, run 'mppm library --list'.`,
	)

	RootCmd.AddCommand(ProjectCmd)

}

var isPreviewCommand bool
var isCommitAllCommand bool
var shouldUpdateLibraries bool

var ProjectCmd = &cobra.Command{

	Use: "project",

	Short: "Provides utilities for managing a specific project.",

	Long: "Provides utilities for managing a specific project.",

	Args: cobra.OnlyValidArgs,

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		RootCmd.PersistentPreRun(cmd, args)
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

	// Clear any session variables between unit tests.
	PersistentPostRun: func(cmd *cobra.Command, args []string) {
		RootCmd.PersistentPostRun(cmd, args)
		isPreviewCommand = false
		isCommitAllCommand = false
		shouldUpdateLibraries = false
	},
}

func commitAll(commitMessage string) (err error) {

	gitRepoFilePath := "."
	gitManager := util.NewGitManager(gitRepoFilePath)

	err = extractAllCompressedFiles()
	if err != nil {
		return
	}

	err = gitManager.AddAllAndCommit(commitMessage)
	if err != nil {
		return
	}

	return

}

func updateProjectLibraryGitCommitIds() (err error) {

	projectConfig, globalConfig, err := configManager.GetProjectAndGlobalConfigs()
	if err != nil {
		return
	}

	projectConfig.Libraries = globalConfig.Libraries
	err = configManager.SaveProjectConfig()
	return
}
