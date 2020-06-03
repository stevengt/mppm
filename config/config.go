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

type MppmConfigInfo struct {
	Version      string                            `json:"version"`
	Applications []*applications.ApplicationConfig `json:"applications"`
	Libraries    []*LibraryConfig                  `json:"libraries"`
}

func (config *MppmConfigInfo) save(filePath string) (err error) {

	configAsJson, err := json.Marshal(config)
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

	installedVersion := Version
	configVersion := config.Version

	installedMajorVersion := strings.Split(installedVersion, ".")[0]
	configMajorVersion := strings.Split(configVersion, ".")[0]

	isCompatible := installedMajorVersion == configMajorVersion

	if !isCompatible {
		errorMessage := getIncompatibleMppmVersionErrorMessage(installedVersion, configVersion)
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

// Returns a list of *applications.FilePatternsConfig, including all non-application-specific configs
// and any supported application-specific configs specified in the project config file.
func GetFilePatternsConfigListFromProjectConfig() (filePatternsConfigList []*applications.FilePatternsConfig) {

	configManager := MppmConfigFileManager

	filePatternsConfigList = applications.GetNonApplicationSpecificFilePatternsConfigList()
	projectApplicationConfigs := configManager.GetProjectConfig().Applications

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
func GetAllFilePatternsConfigFromProjectConfig() (allFilePatternsConfig *applications.FilePatternsConfig) {
	allFilePatternsConfig = applications.NewFilePatternsConfig()
	for _, filePatternsConfig := range GetFilePatternsConfigListFromProjectConfig() {
		allFilePatternsConfig = allFilePatternsConfig.AppendAll(filePatternsConfig)
	}
	return
}
