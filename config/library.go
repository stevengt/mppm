package config

import (
	"fmt"
	"strings"

	"github.com/stevengt/mppm/util"
)

type LibraryConfig struct {
	FilePath              string `json:"location"`
	MostRecentGitCommitId string `json:"most-recent-version"`
	CurrentGitCommitId    string `json:"current-version"`
}

func (libraryConfig *LibraryConfig) Print() {
	libraryConfigAsStringTemplate := `
%s
	most-recent-version="%s"
	current-version="%s"
`
	libraryConfigAsString := fmt.Sprintf(
		libraryConfigAsStringTemplate,
		libraryConfig.FilePath,
		libraryConfig.MostRecentGitCommitId,
		libraryConfig.CurrentGitCommitId,
	)

	util.Println(libraryConfigAsString)
}

func (libraryConfig *LibraryConfig) UpdateCurrentGitCommitId() (err error) {
	gitManager := util.NewGitManager(libraryConfig.FilePath)
	libraryGitCommitId, err := gitManager.RevParse("HEAD")
	if err != nil {
		return
	}
	libraryGitCommitId = strings.Trim(libraryGitCommitId, " \n")
	libraryConfig.CurrentGitCommitId = libraryGitCommitId
	return
}
