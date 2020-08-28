package cmd

import (
	"github.com/spf13/cobra"
	"github.com/stevengt/mppm/config"
	"github.com/stevengt/mppm/util"
)

func init() {

	LibraryCmd.AddCommand(LibraryAddCmd)

}

var LibraryAddCmd = &cobra.Command{

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

	isLibraryPreviouslyAdded, err := isLibraryPreviouslyAdded(libraryFilePath)
	if isLibraryPreviouslyAdded || err != nil {
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

		err = gitManager.LfsTrack("[A-Za-z0-9]*") // Do not track hidden files with git-lfs.
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

	globalConfig, err := configManager.GetGlobalConfig()
	if err != nil {
		return
	}
	globalConfig.Libraries = append(globalConfig.Libraries, libraryConfig)

	err = configManager.SaveGlobalConfig()
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

func isLibraryPreviouslyAdded(libraryFilePath string) (bool, error) {

	globalConfig, err := configManager.GetGlobalConfig()
	if err != nil {
		return false, err
	}

	for _, libraryConfig := range globalConfig.Libraries {
		if libraryConfig.FilePath == libraryFilePath {
			return true, nil
		}
	}
	return false, nil

}
