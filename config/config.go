package config

type FilePatternsConfig struct {
	GitIgnorePatterns   []string
	GitLfsTrackPatterns []string
}

func GetAllFilePatternsConfig() (allFilePatternsConfig *FilePatternsConfig) {

	allFilePatternsConfig = newFilePatternsConfig()

	filePatternsConfigGetters := []func() *FilePatternsConfig{
		getAudioFilePatternsConfig,
		getAbletonFilePatternsConfig,
	}

	for i := 0; i < len(filePatternsConfigGetters); i++ {
		filePatternsConfigGetter := filePatternsConfigGetters[i]
		filePatternsConfig := filePatternsConfigGetter()
		allFilePatternsConfig.GitIgnorePatterns = append(allFilePatternsConfig.GitIgnorePatterns, filePatternsConfig.GitIgnorePatterns...)
		allFilePatternsConfig.GitLfsTrackPatterns = append(allFilePatternsConfig.GitLfsTrackPatterns, filePatternsConfig.GitLfsTrackPatterns...)
	}

	return

}

func newFilePatternsConfig() (filePatternsConfig *FilePatternsConfig) {
	return &FilePatternsConfig{
		GitIgnorePatterns:   make([]string, 0),
		GitLfsTrackPatterns: make([]string, 0),
	}
}

func getAudioFilePatternsConfig() (filePatternsConfig *FilePatternsConfig) {

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

	return &FilePatternsConfig{
		GitIgnorePatterns:   gitIgnorePatterns,
		GitLfsTrackPatterns: gitLfsTrackPatterns,
	}

}

func getAbletonFilePatternsConfig() (filePatternsConfig *FilePatternsConfig) {

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

	return &FilePatternsConfig{
		GitIgnorePatterns:   gitIgnorePatterns,
		GitLfsTrackPatterns: gitLfsTrackPatterns,
	}

}
