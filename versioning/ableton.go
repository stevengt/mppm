package versioning

import (
	"io/ioutil"
	"os"
	"strings"

	"github.com/stevengt/mppm/util"
	"github.com/stevengt/mppm/versioning/project"
)

type AbletonVersioner struct{}

func (v *AbletonVersioner) Init() (err error) {

	config := &project.MppmConfig{
		ProjectType:     project.Ableton,
		GitAnnexFolders: []string{"Samples"},
	}

	err = config.Save()
	if err != nil {
		return
	}

	err = createGitIgnore()
	if err != nil {
		return
	}

	err = util.CreateFoldersIfNotExists(config.GitAnnexFolders...)
	if err != nil {
		return
	}

	err = util.ExecuteShellCommand("git", "init")
	if err != nil {
		return
	}

	err = util.ExecuteShellCommand("git", "annex", "init")
	if err != nil {
		return
	}

	return
}

func (v *AbletonVersioner) Git(args ...string) (err error) {

	err = runPreGitHook()
	if err != nil {
		return
	}

	err = util.ExecuteShellCommand("git", args...)
	if err != nil {
		return
	}

	err = runPostGitHook()
	if err != nil {
		return
	}

	return
}

// Code to execute before invoking any "git" command.
func runPreGitHook() (err error) {

	err = copyAllAlsFilesToUncompressedXmlFiles()
	if err != nil {
		return
	}

	err = runGitAnnexAdd()
	if err != nil {
		return
	}

	return
}

// Code to execute after invoking any "git" command.
func runPostGitHook() (err error) {

	// err = dropUnusedGitAnnexFiles()
	// if err != nil {
	// 	return
	// }

	return
}

func dropUnusedGitAnnexFiles() (err error) {
	err = util.ExecuteShellCommand("git", "annex", "dropunused", "1-1000")
	return
}

func copyAllAlsFilesToUncompressedXmlFiles() (err error) {
	fileNames, err := getAllAlsFileNamesInProject()
	if err != nil {
		return
	}

	for i := 0; i < len(fileNames); i++ {
		originalFileName := fileNames[i]
		newFileName := fileNames[i] + ".xml.gz"

		err = util.CopyFile(originalFileName, newFileName)
		if err != nil {
			return
		}

		err = util.GunzipFile(newFileName)
		if err != nil {
			return
		}
	}

	return
}

func getAllAlsFileNamesInProject() (fileNames []string, err error) {
	fileNames = make([]string, 0)
	stdout, err := util.ExecuteShellCommandAndReturnOutput("find", ".", "-name", "*.als")
	if err == nil {
		stdoutLines := strings.Split(stdout, "\n")
		for i := 0; i < len(stdoutLines); i++ {
			line := stdoutLines[i]
			line = strings.Trim(line, " \n")
			line = strings.TrimPrefix(line, "./")
			if len(line) > 0 {
				fileNames = append(fileNames, line)
			}
		}
	}
	return
}

func runGitAnnexAdd() (err error) {
	config, err := project.LoadMppmConfig()
	if err != nil {
		return
	}

	for i := 0; i < len(config.GitAnnexFolders); i++ {
		gitAnnexFolder := config.GitAnnexFolders[i]
		err = util.ExecuteShellCommand("git", "annex", "add", gitAnnexFolder)
		if err != nil {
			return
		}
	}

	return
}

func createGitIgnore() (err error) {
	fileName := ".gitignore"
	filePermissionsCode := os.FileMode(0644)
	fileContents := "Backup/\n"
	fileContents += "*.als\n"
	err = ioutil.WriteFile(fileName, []byte(fileContents), filePermissionsCode)
	return
}
