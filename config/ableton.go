package config

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
