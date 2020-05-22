package config

type GitConfigFilePatterns struct {
	GitIgnorePatterns   []string
	GitLfsTrackPatterns []string
}

func GetAllGitConfigFilePatterns() (allFilePatterns *GitConfigFilePatterns) {

	allFilePatterns = newGitConfigFilePatterns()

	filePatternGetters := []func() *GitConfigFilePatterns{
		getAbletonGitConfigFilePatterns,
	}

	for i := 0; i < len(filePatternGetters); i++ {
		filePatternGetter := filePatternGetters[i]
		filePatterns := filePatternGetter()
		allFilePatterns.GitIgnorePatterns = append(allFilePatterns.GitIgnorePatterns, filePatterns.GitIgnorePatterns...)
		allFilePatterns.GitLfsTrackPatterns = append(allFilePatterns.GitLfsTrackPatterns, filePatterns.GitLfsTrackPatterns...)
	}

	return

}

func newGitConfigFilePatterns() (filePatterns *GitConfigFilePatterns) {
	return &GitConfigFilePatterns{
		GitIgnorePatterns:   make([]string, 0),
		GitLfsTrackPatterns: make([]string, 0),
	}
}

func getAbletonGitConfigFilePatterns() (filePatterns *GitConfigFilePatterns) {
	gitIgnorePatterns := []string{
		"Backup/",
		"*.als",
	}

	gitLfsTrackPatterns := []string{
		"*.alp",
		"*.asd",
		"*.ask",
		"*.adg",
		"*.adv",
		"*.alc",
		"*.agr",
		"*.ams",
		"*.amxd",
	}

	return &GitConfigFilePatterns{
		GitIgnorePatterns:   gitIgnorePatterns,
		GitLfsTrackPatterns: gitLfsTrackPatterns,
	}
}
