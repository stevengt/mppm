package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stevengt/mppm/config"
	"github.com/stevengt/mppm/util"
)

func init() {

	cobra.OnInitialize(
		func() {
			isListAllLibrariesCommand, _ = libraryCmd.Flags().GetBool("list")
			isCommitAllLibrariesCommand, _ = libraryCmd.Flags().GetBool("commit-all")
		},
	)

	libraryCmd.Flags().BoolVarP(
		&isListAllLibrariesCommand,
		"list",
		"l",
		false,
		"Lists all libraries (folders) currently tracked globally on your system.",
	)

	libraryCmd.Flags().BoolVarP(
		&isCommitAllLibrariesCommand,
		"commit-all",
		"c",
		false,
		"Commits (snapshots) all changes made to all libraries (folders) currently tracked globally on your system.",
	)

	rootCmd.AddCommand(libraryCmd)

}

var libraryCmd = &cobra.Command{

	Use: "library",

	Short: "Provides utilities for globally managing multiple libraries (folders).",

	Long: `Provides utilities for globally managing multiple libraries (folders).

Specifically, you can specify a list of folders and periodically take "snapshots"
of their contents. mppm can then keep track of which "versions" of the folders
each of your projects depend on, at any given time.

This is useful in the case that you inadvertently change/delete a file that your project
depends on. While working on your project, you can simply revert to a previous snapshot
of your libraries to get the original file back. When you finish working on your project,
you can then restore your libraries to their most recent snapshot.

Libraries can be any collection of audio samples, plugins, presets, etc. that:
	- Your projects might "depend on".
	- You expect, in general, to update less frequently compared to projects.
`,

	Args: cobra.OnlyValidArgs,

	Run: func(cmd *cobra.Command, args []string) {
		if isListAllLibrariesCommand {
			listAllLibraries()
		} else if isCommitAllLibrariesCommand {
			err := commitAllLibraries()
			if err != nil {
				util.ExitWithError(err)
			}
		} else {
			cmd.Help()
		}
	},
}

var isListAllLibrariesCommand bool
var isCommitAllLibrariesCommand bool

func listAllLibraries() {
	for _, libraryConfig := range configManager.GetGlobalConfig().Libraries {
		libraryConfig.Print()
	}
}

func commitAllLibraries() (err error) {
	for _, libraryConfig := range configManager.GetGlobalConfig().Libraries {
		err = commitLibrary(libraryConfig)
		if err != nil {
			return
		}
	}
	return
}

func commitLibrary(libraryConfig *config.LibraryConfig) (err error) {

	gitManager := util.NewGitManager(libraryConfig.FilePath)

	err = gitManager.AddAllAndCommit("Committed all changes.")
	if err != nil {
		return
	}

	err = libraryConfig.UpdateCurrentGitCommitId()
	if err != nil {
		return
	}
	libraryConfig.MostRecentGitCommitId = libraryConfig.CurrentGitCommitId

	err = configManager.SaveGlobalConfig()
	if err != nil {
		return
	}

	return

}
