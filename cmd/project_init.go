package cmd

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/stevengt/mppm/config"
	"github.com/stevengt/mppm/util"
)

func init() {
	projectCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{

	Use: "init",

	Short: "Initializes version control settings for a project using git and git-lfs.",

	Long: "Initializes version control settings for a project using git and git-lfs.",

	Args: cobra.NoArgs,

	Run: func(cmd *cobra.Command, args []string) {
		if err := initProject(); err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

func initProject() (err error) {

	gitManager := util.CurrentDirectoryGitManager

	err = createMppmProjectConfigFile()
	if err != nil {
		return
	}

	err = gitManager.Init()
	if err != nil {
		return
	}

	err = gitManager.LfsInstall()
	if err != nil {
		return
	}

	filePatternsConfig := config.GetAllFilePatternsConfig()

	err = createGitIgnoreFile(filePatternsConfig.GitIgnorePatterns...)
	if err != nil {
		return
	}

	err = gitManager.LfsTrack(filePatternsConfig.GitLfsTrackPatterns...)
	if err != nil {
		return
	}

	err = gitManager.Add(".gitignore", ".gitattributes", config.MppmConfigFileName)
	if err != nil {
		return
	}

	err = gitManager.Commit("-m", "Initial commit.")
	if err != nil {
		return
	}

	return

}

func createGitIgnoreFile(filePatterns ...string) (err error) {
	fileName := ".gitignore"
	filePermissionsCode := os.FileMode(0644)
	fileContents := strings.Join(filePatterns, "\n")
	err = ioutil.WriteFile(fileName, []byte(fileContents), filePermissionsCode)
	return
}

func createMppmProjectConfigFile() (err error) {
	mppmProjectConfig := config.GetDefaultMppmConfig()
	err = mppmProjectConfig.SaveAsProjectConfig()
	return
}
