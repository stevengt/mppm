package config

type LibraryConfig struct {
	FilePath              string `json:"location"`
	MostRecentGitCommitId string `json:"most-recent-version"`
	CurrentGitCommitId    string `json:"current-version"`
}
