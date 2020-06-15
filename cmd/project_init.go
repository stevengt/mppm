package cmd

import (
	"bytes"
	"io"
	"strings"

	"github.com/spf13/cobra"
	"github.com/stevengt/mppm/config"
	"github.com/stevengt/mppm/config/applications"
	"github.com/stevengt/mppm/util"
)

func init() {
	ProjectCmd.AddCommand(InitCmd)
}

var InitCmd = &cobra.Command{

	Use: "init",

	Short: "Initializes version control settings for a project using git and git-lfs.",

	Long: "Initializes version control settings for a project using git and git-lfs.",

	Args: cobra.NoArgs,

	Run: func(cmd *cobra.Command, args []string) {
		if err := initProject(); err != nil {
			util.ExitWithError(err)
		}
	},
}

func initProject() (err error) {

	gitRepoFilePath := "."
	gitManager := util.NewGitManager(gitRepoFilePath)

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

	filePatternsConfig := applications.GetAllFilePatternsConfig()

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
	fileContents := strings.Join(filePatterns, "\n")
	fileContentsAsReader := bytes.NewReader([]byte(fileContents))

	file, err := util.CreateFile(fileName)
	if err != nil {
		return
	}
	defer file.Close()

	io.Copy(file, fileContentsAsReader)

	return
}

func createMppmProjectConfigFile() (err error) {
	err = configManager.SaveDefaultProjectConfig()
	return
}
