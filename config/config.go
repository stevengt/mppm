package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/stevengt/mppm/util"
)

var Version = "1.2.1"

var MppmProjectConfig, MppmGlobalConfig *MppmConfigInfo

var MppmConfigFileName = ".mppm.json"

var mppmConfigManagerFactory MppmConfigManagerCreator = &MppmConfigFileManagerCreator{}

func NewMppmConfigManager() MppmConfigManager {
	return mppmConfigManagerFactory.NewMppmConfigManager()
}

type MppmConfigManagerCreator interface {
	NewMppmConfigManager() MppmConfigManager
}

type MppmConfigFileManagerCreator struct{}

func (mppmConfigFileManagerCreator *MppmConfigFileManagerCreator) NewMppmConfigManager() MppmConfigManager {
	return &MppmConfigFileManager{}
}

type MppmConfigManager interface {
	GetProjectConfig() *MppmConfigInfo
	GetGlobalConfig() *MppmConfigInfo
}

type MppmConfigFileManager struct {
	projectConfig *MppmConfigInfo
	globalConfig  *MppmConfigInfo
}

func (configFileManager *MppmConfigFileManager) GetProjectConfig() *MppmConfigInfo {
	if configFileManager.projectConfig == nil {
		configFileManager.projectConfig = &MppmConfigInfo{}
		loadMppmConfig(configFileManager.projectConfig, MppmConfigFileName)
	}
	return configFileManager.projectConfig
}

func (configFileManager *MppmConfigFileManager) GetGlobalConfig() *MppmConfigInfo {
	if configFileManager.globalConfig == nil {
		configFileManager.globalConfig = &MppmConfigInfo{}
		createMppmGlobalConfigFileIfNotExists()
		loadMppmConfig(configFileManager.globalConfig, GetMppmGlobalConfigFilePath())
	}
	return configFileManager.globalConfig
}

type MppmConfigInfo struct {
	Version      string               `json:"version"`
	Applications []*ApplicationConfig `json:"applications"`
	Libraries    []*LibraryConfig     `json:"libraries"`
}

func (config *MppmConfigInfo) SaveAsProjectConfig() (err error) {
	err = config.save(MppmConfigFileName)
	return
}

func (config *MppmConfigInfo) SaveAsGlobalConfig() (err error) {

	configFilePath := GetMppmGlobalConfigFilePath()
	err = config.save(configFilePath)
	if err != nil {
		return
	}

	return

}

func (config *MppmConfigInfo) save(filePath string) (err error) {

	configAsJson, err := json.Marshal(config)
	if err != nil {
		return
	}

	filePermissionsCode := os.FileMode(0644)
	err = ioutil.WriteFile(filePath, configAsJson, filePermissionsCode)
	if err != nil {
		return
	}

	return

}

func (config *MppmConfigInfo) CheckIfCompatibleWithInstalledMppmVersion() (err error) {

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

func (config *MppmConfigInfo) CheckIfCompatibleWithSupportedApplications() (err error) {

	for _, application := range config.Applications {

		isApplicationSupported := false
		for _, supportedApplication := range SupportedApplications {
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

// TODO Override global config settings with project config settings.
func LoadMppmProjectConfig() {
	MppmProjectConfig = &MppmConfigInfo{}
	loadMppmConfig(MppmProjectConfig, MppmConfigFileName)
}

func LoadMppmGlobalConfig() {
	MppmGlobalConfig = &MppmConfigInfo{}
	createMppmGlobalConfigFileIfNotExists()
	loadMppmConfig(MppmGlobalConfig, GetMppmGlobalConfigFilePath())
}

func loadMppmConfig(config *MppmConfigInfo, configFilePath string) {

	configFile, err := os.Open(configFilePath)
	if err != nil {
		errorMessage := getOpeningMppmProjectConfigFileErrorMessage(err)
		util.ExitWithErrorMessage(errorMessage)
	}
	defer configFile.Close()

	jsonDecoder := json.NewDecoder(configFile)
	jsonDecoder.DisallowUnknownFields()

	err = jsonDecoder.Decode(config)
	if err != nil {
		errorMessage := getInvalidMppmProjectConfigFileErrorMessage(err)
		util.ExitWithErrorMessage(errorMessage)
	}

	err = config.CheckIfCompatibleWithInstalledMppmVersion()
	if err != nil {
		util.ExitWithError(err)
	}

	err = config.CheckIfCompatibleWithSupportedApplications()
	if err != nil {
		util.ExitWithError(err)
	}

}

func GetDefaultMppmConfig() (mppmConfig *MppmConfigInfo) {

	applicationConfigList := make([]*ApplicationConfig, 0)
	libraryConfigList := make([]*LibraryConfig, 0)

	for _, supportedApplication := range SupportedApplications {
		applicationConfig := &ApplicationConfig{
			Name:    supportedApplication.Name,
			Version: supportedApplication.DefaultVersion,
		}
		applicationConfigList = append(applicationConfigList, applicationConfig)
	}

	mppmConfig = &MppmConfigInfo{
		Version:      Version,
		Applications: applicationConfigList,
		Libraries:    libraryConfigList,
	}

	return

}

func GetMppmGlobalConfigFilePath() (filePath string) {
	homeDirectoryPath, _ := os.UserHomeDir()
	filePath = filepath.Join(homeDirectoryPath, MppmConfigFileName)
	return
}

func createMppmGlobalConfigFileIfNotExists() {
	mppmGlobalConfigFilePath := GetMppmGlobalConfigFilePath()
	defaultMppmGlobalConfig := GetDefaultMppmConfig()

	if _, err := os.Stat(mppmGlobalConfigFilePath); os.IsNotExist(err) {
		err = defaultMppmGlobalConfig.SaveAsGlobalConfig()
		if err != nil {
			util.ExitWithError(err)
		}
	}
}
