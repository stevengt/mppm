package config

type ApplicationInfo struct {
	Name               string
	SupportedVersions  []ApplicationVersion
	DefaultVersion     ApplicationVersion
	FilePatternConfigs map[ApplicationVersion]*FilePatternsConfig
}

type ApplicationVersion string

var SupportedApplications = []*ApplicationInfo{
	AbletonInfo,
}
