package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stevengt/mppm/config"
	"github.com/stevengt/mppm/util"
)

func init() {

	cobra.OnInitialize(
		func() {
			isListAllLibrariesCommand, _ = LibraryCmd.Flags().GetBool("list")
			isCommitAllLibrariesCommand, _ = LibraryCmd.Flags().GetBool("commit-all")
		},
	)

	LibraryCmd.Flags().BoolVarP(
		&isListAllLibrariesCommand,
		"list",
		"l",
		false,
		"Lists all libraries (folders) currently tracked globally on your system.",
	)

	LibraryCmd.Flags().BoolVarP(
		&isCommitAllLibrariesCommand,
		"commit-all",
		"c",
		false,
		"Commits (snapshots) all changes made to all libraries (folders) currently tracked globally on your system.",
	)

	RootCmd.AddCommand(LibraryCmd)

}

var LibraryCmd = &cobra.Command{

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

		var err error

		if isListAllLibrariesCommand {
			err = listAllLibraries()
		} else if isCommitAllLibrariesCommand {
			err = commitAllLibraries()
		} else {
			util.Println(cmd.UsageString())
		}

		if err != nil {
			util.ExitWithError(err)
		}

	},
}

var isListAllLibrariesCommand bool
var isCommitAllLibrariesCommand bool

func listAllLibraries() (err error) {

	globalConfig, err := configManager.GetGlobalConfig()
	if err != nil {
		return
	}

	for _, libraryConfig := range globalConfig.Libraries {
		libraryConfig.Print()
	}

	return

}

func commitAllLibraries() (err error) {

	globalConfig, err := configManager.GetGlobalConfig()
	if err != nil {
		return
	}

	for _, libraryConfig := range globalConfig.Libraries {
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
