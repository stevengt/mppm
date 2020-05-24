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

	err = createMppmProjectConfigFile()
	if err != nil {
		return
	}

	err = util.ExecuteShellCommand("git", "init")
	if err != nil {
		return
	}

	err = util.ExecuteShellCommand("git", "lfs", "install")
	if err != nil {
		return
	}

	filePatternsConfig := config.GetAllFilePatternsConfig()

	err = createGitIgnoreFile(filePatternsConfig.GitIgnorePatterns...)
	if err != nil {
		return
	}

	err = runGitLfsTrack(filePatternsConfig.GitLfsTrackPatterns...)
	if err != nil {
		return
	}

	err = util.ExecuteShellCommand("git", "add", ".gitignore", ".gitattributes", config.MppmProjectConfigFileName)
	if err != nil {
		return
	}

	err = util.ExecuteShellCommand("git", "commit", "-m", "Initial commit.")
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

func runGitLfsTrack(filePatterns ...string) (err error) {
	commandName := "git"
	commandArgs := append([]string{"lfs", "track"}, filePatterns...)
	err = util.ExecuteShellCommand(commandName, commandArgs...)
	if err != nil {
		return
	}
	return
}

func createMppmProjectConfigFile() (err error) {
	mppmProjectConfig := config.GetDefaultMppmProjectConfig()
	err = mppmProjectConfig.Save()
	return
}
