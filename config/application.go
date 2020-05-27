package config

type ApplicationInfo struct {
	Name               ApplicationName
	SupportedVersions  []ApplicationVersion
	DefaultVersion     ApplicationVersion
	FilePatternConfigs map[ApplicationVersion]*FilePatternsConfig
}

type ApplicationConfig struct {
	Name    ApplicationName    `json:"name"`
	Version ApplicationVersion `json:"version"`
}

type ApplicationName string
type ApplicationVersion string

var SupportedApplications = map[string]*ApplicationInfo{
	"Ableton": AbletonInfo,
}
