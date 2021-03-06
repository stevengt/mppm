package applications

func GetNonApplicationSpecificFilePatternsConfigList() (filePatternsConfigList []*FilePatternsConfig) {
	return []*FilePatternsConfig{
		AudioFilePatternsConfig,
	}
}

var AudioFilePatternsConfig *FilePatternsConfig = &FilePatternsConfig{

	Name: "Audio",

	GitIgnorePatterns: []string{},

	GitLfsTrackPatterns: []string{
		"*.3gp",
		"*.aa",
		"*.aac",
		"*.aax",
		"*.act",
		"*.aiff",
		"*.alac",
		"*.amr",
		"*.ape",
		"*.au",
		"*.awb",
		"*.dct",
		"*.dss",
		"*.dvf",
		"*.flac",
		"*.gsm",
		"*.iklax",
		"*.ivs",
		"*.m4a",
		"*.m4b",
		"*.m4p",
		"*.mmf",
		"*.mp3",
		"*.mpc",
		"*.msv",
		"*.nmf",
		"*.nsf",
		"*.ogg",
		"*.oga",
		"*.mogg",
		"*.opus",
		"*.ra",
		"*.rm",
		"*.raw",
		"*.rf64",
		"*.sln",
		"*.tta",
		"*.voc",
		"*.vox",
		"*.wav",
		"*.wma",
		"*.wv",
		"*.webm",
		"*.8svx",
		"*.cda",
	},

	GzippedXmlFileExtensions: []string{},
}
