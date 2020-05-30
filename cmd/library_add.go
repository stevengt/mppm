package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stevengt/mppm/config"
	"github.com/stevengt/mppm/util"
)

func init() {

	libraryCmd.AddCommand(libraryAddCmd)

}

var libraryAddCmd = &cobra.Command{

	Use: "add",

	Short: "Adds a library (folder) to track globally on your system.",

	Long: "Adds a library (folder) to track globally on your system.",

	Args: cobra.MinimumNArgs(1),

	Run: func(cmd *cobra.Command, args []string) {
		err := addLibrary(args[0])
		if err != nil {
			util.ExitWithError(err)
		}
	},
}

func addLibrary(libraryFilePath string) (err error) {

	if isLibraryPreviouslyAdded(libraryFilePath) {
		return
	}

	if !isGitRepository(libraryFilePath) {

		gitManager := util.NewGitManager(libraryFilePath)

		err = gitManager.Init()
		if err != nil {
			return
		}

		err = gitManager.LfsInstall()
		if err != nil {
			return
		}

		err = gitManager.LfsTrack("*")
		if err != nil {
			return
		}

		err = gitManager.AddAllAndCommit("Initial commit.")
		if err != nil {
			return
		}

	}

	libraryConfig := &config.LibraryConfig{
		FilePath: libraryFilePath,
	}

	err = libraryConfig.UpdateCurrentGitCommitId()
	if err != nil {
		return
	}
	libraryConfig.MostRecentGitCommitId = libraryConfig.CurrentGitCommitId

	config.MppmGlobalConfig.Libraries = append(config.MppmGlobalConfig.Libraries, libraryConfig)
	err = config.MppmGlobalConfig.SaveAsGlobalConfig()
	if err != nil {
		return
	}

	return

}

func isGitRepository(libraryFilePath string) bool {
	gitManager := util.NewGitManager(libraryFilePath)
	_, err := gitManager.RevParse()
	return err == nil
}

func isLibraryPreviouslyAdded(libraryFilePath string) bool {
	for _, libraryConfig := range config.MppmGlobalConfig.Libraries {
		if libraryConfig.FilePath == libraryFilePath {
			return true
		}
	}
	return false
}
