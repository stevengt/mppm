package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strings"

	"github.com/stevengt/mppm/config/applications"
	"github.com/stevengt/mppm/util"
)

var Version = "1.2.1"

var MppmConfigFileName = ".mppm.json"

// ------------------------------------------------------------------------------

// Returns a list of *applications.FilePatternsConfig, including all non-application-specific configs
// and any supported application-specific configs specified in the project config file.
func GetFilePatternsConfigListFromProjectConfig() (filePatternsConfigList []*applications.FilePatternsConfig, err error) {

	configManager := MppmConfigFileManager

	projectConfig, err := configManager.GetProjectConfig()
	if err != nil {
		return
	}
	projectApplicationConfigs := projectConfig.Applications

	filePatternsConfigList = applications.GetNonApplicationSpecificFilePatternsConfigList()

	for _, projectApplicationConfig := range projectApplicationConfigs {
		for _, supportedApplication := range applications.SupportedApplications {
			if supportedApplication.Name == projectApplicationConfig.Name {
				if filePatternsConfig, ok := supportedApplication.FilePatternConfigs[projectApplicationConfig.Version]; ok {
					filePatternsConfigList = append(filePatternsConfigList, filePatternsConfig)
				}
			}
		}
	}

	return

}

// Returns a single *applications.FilePatternsConfig containing the aggregate of all file patterns,
// including all non-application-specific configs and any supported application-specific configs
// specified in the project config file.
func GetAllFilePatternsConfigFromProjectConfig() (allFilePatternsConfig *applications.FilePatternsConfig, err error) {

	filePatternsConfigList, err := GetFilePatternsConfigListFromProjectConfig()
	if err != nil {
		return
	}

	allFilePatternsConfig = applications.NewFilePatternsConfig()

	for _, filePatternsConfig := range filePatternsConfigList {
		allFilePatternsConfig = allFilePatternsConfig.AppendAll(filePatternsConfig)
	}

	return

}

func GetCurrentlyInstalledMajorVersion() string {
	return strings.Split(Version, ".")[0]
}

// ------------------------------------------------------------------------------

type MppmConfigInfo struct {
	Version      string                            `json:"version"`
	Applications []*applications.ApplicationConfig `json:"applications"`
	Libraries    []*LibraryConfig                  `json:"libraries"`
}

func NewMppmConfigInfoFromJson(json []byte) (mppmConfig *MppmConfigInfo, err error) {
	jsonReader := bytes.NewReader(json)
	return NewMppmConfigInfoFromJsonReader(jsonReader)
}

func NewMppmConfigInfoFromJsonReader(jsonReader io.Reader) (mppmConfig *MppmConfigInfo, err error) {

	mppmConfig = &MppmConfigInfo{}

	jsonDecoder := json.NewDecoder(jsonReader)
	jsonDecoder.DisallowUnknownFields()

	err = jsonDecoder.Decode(mppmConfig)
	if err != nil {
		errorMessage := getInvalidMppmProjectConfigFileErrorMessage(err)
		err = errors.New(errorMessage)
		return
	}
	return
}

func (config *MppmConfigInfo) AsJson() (configAsJson []byte, err error) {
	return json.Marshal(config)
}

func (config *MppmConfigInfo) save(filePath string) (err error) {

	configAsJson, err := config.AsJson()
	if err != nil {
		return
	}
	configAsJsonReader := bytes.NewReader(configAsJson)

	file, err := util.CreateFile(filePath)
	if err != nil {
		return
	}
	defer file.Close()

	_, err = io.Copy(file, configAsJsonReader)
	if err != nil {
		return
	}

	return

}

func (config *MppmConfigInfo) checkIfCompatibleWithInstalledMppmVersion() (err error) {
	if GetCurrentlyInstalledMajorVersion() != config.getMajorVersion() {
		errorMessage := getIncompatibleMppmVersionErrorMessage(Version, config.Version)
		err = errors.New(errorMessage)
	}
	return
}

func (config *MppmConfigInfo) checkIfCompatibleWithSupportedApplications() (err error) {

	for _, application := range config.Applications {

		isApplicationSupported := false
		for _, supportedApplication := range applications.SupportedApplications {
			if application.Name == supportedApplication.Name {

				isVersionSupported := false
				for _, supportedVersion := range supportedApplication.SupportedVersions {
					if application.Version == supportedVersion {
						isVersionSupported = true
						break
					}
				}
				if isVersionSupported {
					isApplicationSupported = true
					break
				}
			}

		}
		if !isApplicationSupported {
			errorMessage := getUnsupportedApplicationErrorMessage(application)
			err = errors.New(errorMessage)
			return
		}

	}

	return

}

func (config *MppmConfigInfo) getMajorVersion() string {
	return strings.Split(config.Version, ".")[0]
}
