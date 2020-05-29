package cmd

import (
	"strings"

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

	err = util.ExecuteShellCommand("git", "-C", libraryFilePath, "init")
	if err != nil {
		return
	}

	err = util.ExecuteShellCommand("git", "-C", libraryFilePath, "lfs", "install")
	if err != nil {
		return
	}

	err = util.ExecuteShellCommand("git", "-C", libraryFilePath, "lfs", "track", "*")
	if err != nil {
		return
	}

	err = util.ExecuteShellCommand("git", "-C", libraryFilePath, "add", "-A", ".")
	if err != nil {
		return
	}

	err = util.ExecuteShellCommand("git", "-C", libraryFilePath, "commit", "-m", "Initial commit.")
	if err != nil {
		return
	}

	libraryGitCommitId, err := util.ExecuteShellCommandAndReturnOutput("git", "-C", libraryFilePath, "rev-parse", "HEAD")
	if err != nil {
		return
	}
	libraryGitCommitId = strings.Trim(libraryGitCommitId, " \n")

	libraryConfig := &config.LibraryConfig{
		FilePath:              libraryFilePath,
		MostRecentGitCommitId: libraryGitCommitId,
		CurrentGitCommitId:    libraryGitCommitId,
	}

	config.MppmGlobalConfig.Libraries = append(config.MppmGlobalConfig.Libraries, libraryConfig)
	err = config.MppmGlobalConfig.SaveAsGlobalConfig()
	if err != nil {
		return
	}

	return

}
