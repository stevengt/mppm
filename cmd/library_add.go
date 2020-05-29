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

	err = util.ExecuteGitCommandInDirectory(libraryFilePath, "init")
	if err != nil {
		return
	}

	err = util.ExecuteGitCommandInDirectory(libraryFilePath, "lfs", "install")
	if err != nil {
		return
	}

	err = util.ExecuteGitCommandInDirectory(libraryFilePath, "lfs", "track", "*")
	if err != nil {
		return
	}

	err = addAllAndCommit(libraryFilePath)
	if err != nil {
		return
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
