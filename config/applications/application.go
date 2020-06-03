package applications

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

func GetApplicationSpecificFilePatternsConfigList() (filePatternsConfigList []*FilePatternsConfig) {
	filePatternsConfigList = make([]*FilePatternsConfig, 0)
	for _, supprtedApplication := range SupportedApplications {
		for _, supportedVersionConfig := range supprtedApplication.FilePatternConfigs {
			filePatternsConfigList = append(filePatternsConfigList, supportedVersionConfig)
		}
	}
	return
}
