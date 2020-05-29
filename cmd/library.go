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

	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		config.LoadMppmGlobalConfig()
	},

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
	for _, libraryConfig := range config.MppmGlobalConfig.Libraries {
		libraryConfig.Print()
	}
}

func commitAllLibraries() (err error) {
	for _, libraryConfig := range config.MppmGlobalConfig.Libraries {
		err = commitLibrary(libraryConfig)
		if err != nil {
			return
		}
	}
	return
}

func commitLibrary(libraryConfig *config.LibraryConfig) (err error) {

	err = addAllAndCommit(libraryConfig.FilePath)
	if err != nil {
		return
	}

	err = libraryConfig.UpdateCurrentGitCommitId()
	if err != nil {
		return
	}

	err = config.MppmGlobalConfig.SaveAsGlobalConfig()
	if err != nil {
		return
	}

	return

}

func addAllAndCommit(gitRepoFilePath string) (err error) {

	err = util.ExecuteGitCommandInDirectory(gitRepoFilePath, "add", "-A", ".")
	if err != nil {
		return
	}

	err = util.ExecuteGitCommandInDirectory(gitRepoFilePath, "commit", "-m", "Committed all changes.")
	if err != nil {
		return
	}

	return

}
