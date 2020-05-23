package config

var AbletonFilePatternsConfig *FilePatternsConfig = &FilePatternsConfig{

	GitIgnorePatterns: []string{
		"Backup/",
		"*.als",
		"*.alc",
		"*.adv",
		"*.adg",
	},

	GitLfsTrackPatterns: []string{
		"*.alp",
		"*.asd",
		"*.agr",
		"*.ams",
		"*.amxd",
	},

	GzippedXmlFileExtensions: []string{
		"als",
		"alc",
		"adv",
		"adg",
	},
}
