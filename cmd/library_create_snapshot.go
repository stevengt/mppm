package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stevengt/mppm/config"
	"github.com/stevengt/mppm/util"
)

func init() {

	libraryCmd.AddCommand(libraryCreateSnapshotCmd)

}

var libraryCreateSnapshotCmd = &cobra.Command{

	Use: "create-snapshot",

	Short: "Creates a snapshot of all libraries (folders) currently tracked globally on your system.",

	Long: "Creates a snapshot of all libraries (folders) currently tracked globally on your system.",

	Args: cobra.NoArgs,

	Run: func(cmd *cobra.Command, args []string) {
		err := createAllLibrarySnapshots()
		if err != nil {
			util.ExitWithError(err)
		}
	},
}

func createAllLibrarySnapshots() (err error) {
	for _, libraryConfig := range config.MppmGlobalConfig.Libraries {
		err = createLibrarySnapshot(libraryConfig)
		if err != nil {
			return
		}
	}
	return
}

func createLibrarySnapshot(libraryConfig *config.LibraryConfig) (err error) {

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
