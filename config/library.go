package config

import "fmt"

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

	fmt.Println(libraryConfigAsString)
}
