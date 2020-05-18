package versioning

import (
	"errors"
	"io/ioutil"
	"os"

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

func (v *AbletonVersioner) Git(args []string) (err error) {
	err = errors.New("Not Implemented.")
	return
}

func createGitIgnore() (err error) {
	fileName := ".gitignore"
	filePermissionsCode := os.FileMode(0644)
	fileContents := "Backup/\n"
	err = ioutil.WriteFile(fileName, []byte(fileContents), filePermissionsCode)
	return
}
