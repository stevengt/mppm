package config

type FilePatternsConfig struct {
	GitIgnorePatterns        []string
	GitLfsTrackPatterns      []string
	GzippedXmlFileExtensions []string // List of file extensions that represent Gzipped XML files.
}

func newFilePatternsConfig() (filePatternsConfig *FilePatternsConfig) {
	return &FilePatternsConfig{
		GitIgnorePatterns:        make([]string, 0),
		GitLfsTrackPatterns:      make([]string, 0),
		GzippedXmlFileExtensions: make([]string, 0),
	}
}

func (config1 *FilePatternsConfig) appendAll(config2 *FilePatternsConfig) (filePatternsConfig *FilePatternsConfig) {
	config1.GitIgnorePatterns = append(config1.GitIgnorePatterns, config2.GitIgnorePatterns...)
	config1.GitLfsTrackPatterns = append(config1.GitLfsTrackPatterns, config2.GitLfsTrackPatterns...)
	config1.GzippedXmlFileExtensions = append(config1.GzippedXmlFileExtensions, config2.GzippedXmlFileExtensions...)
	filePatternsConfig = config1
	return
}

func GetAllFilePatternsConfig() (allFilePatternsConfig *FilePatternsConfig) {

	allFilePatternsConfig = newFilePatternsConfig()

	filePatternsConfigList := []*FilePatternsConfig{
		AudioFilePatternsConfig,
		AbletonFilePatternsConfig,
	}

	for i := 0; i < len(filePatternsConfigList); i++ {
		filePatternsConfig := filePatternsConfigList[i]
		allFilePatternsConfig = allFilePatternsConfig.appendAll(filePatternsConfig)
	}

	return

}

var AudioFilePatternsConfig *FilePatternsConfig = &FilePatternsConfig{

	GitIgnorePatterns: []string{},

	GitLfsTrackPatterns: []string{
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
	},

	GzippedXmlFileExtensions: []string{},
}

var AbletonFilePatternsConfig *FilePatternsConfig = &FilePatternsConfig{

	GitIgnorePatterns: []string{
		"Backup/",
		"*.als",
		"*.alc",
	},

	GitLfsTrackPatterns: []string{
		"*.alp",
		"*.asd",
		"*.ask",
		"*.adg",
		"*.adv",
		"*.agr",
		"*.ams",
		"*.amxd",
	},

	GzippedXmlFileExtensions: []string{
		"als",
		"alc",
	},
}
