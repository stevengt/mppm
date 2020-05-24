package config

var AbletonInfo = &ApplicationInfo{

	Name: "Ableton",

	SupportedVersions: []ApplicationVersion{
		"10",
	},

	DefaultVersion: "10",

	FilePatternConfigs: map[ApplicationVersion]*FilePatternsConfig{
		"10": Ableton10FilePatternsConfig,
	},
}

var Ableton10FilePatternsConfig *FilePatternsConfig = &FilePatternsConfig{

	Name: "Ableton 10",

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
