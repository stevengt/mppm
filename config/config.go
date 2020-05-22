package config

type GitConfigFilePatterns struct {
	GitIgnorePatterns   []string
	GitLfsTrackPatterns []string
}

func GetAllGitConfigFilePatterns() (allFilePatterns *GitConfigFilePatterns) {

	allFilePatterns = newGitConfigFilePatterns()

	filePatternGetters := []func() *GitConfigFilePatterns{
		getAudioGitConfigFilePatterns,
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

func getAudioGitConfigFilePatterns() (filePatterns *GitConfigFilePatterns) {

	gitIgnorePatterns := []string{}

	gitLfsTrackPatterns := []string{
		".3gp",
		".aa",
		".aac",
		".aax",
		".act",
		".aiff",
		".alac",
		".amr",
		".ape",
		".au",
		".awb",
		".dct",
		".dss",
		".dvf",
		".flac",
		".gsm",
		".iklax",
		".ivs",
		".m4a",
		".m4b",
		".m4p",
		".mmf",
		".mp3",
		".mpc",
		".msv",
		".nmf",
		".nsf",
		".ogg",
		".oga",
		".mogg",
		".opus",
		".ra",
		".rm",
		".raw",
		".rf64",
		".sln",
		".tta",
		".voc",
		".vox",
		".wav",
		".wma",
		".wv",
		".webm",
		".8svx",
		".cda",
	}

	return &GitConfigFilePatterns{
		GitIgnorePatterns:   gitIgnorePatterns,
		GitLfsTrackPatterns: gitLfsTrackPatterns,
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
