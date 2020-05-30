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

var MppmConfigFileName = ".mppm.json"

var mppmConfigManagerFactory MppmConfigManagerCreator = &MppmConfigFileManagerCreator{}
var mppmConfigManager MppmConfigManager

func GetMppmConfigManager() MppmConfigManager {
	if mppmConfigManager == nil {
		mppmConfigManager = mppmConfigManagerFactory.NewMppmConfigManager()
	}
	return mppmConfigManager
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
	GetDefaultMppmConfig() (mppmConfig *MppmConfigInfo)
	GetMppmGlobalConfigFilePath() (filePath string)
	SaveProjectConfig() (err error)
	SaveGlobalConfig() (err error)
	SaveDefaultProjectConfig() (err error)
}

type MppmConfigFileManager struct {
	projectConfig *MppmConfigInfo
	globalConfig  *MppmConfigInfo
}

func (configFileManager *MppmConfigFileManager) GetProjectConfig() *MppmConfigInfo {
	if configFileManager.projectConfig == nil {
		configFileManager.projectConfig = &MppmConfigInfo{}
		configFileManager.loadMppmConfig(
			configFileManager.projectConfig,
			MppmConfigFileName,
		)
	}
	return configFileManager.projectConfig
}

func (configFileManager *MppmConfigFileManager) GetGlobalConfig() *MppmConfigInfo {
	if configFileManager.globalConfig == nil {
		configFileManager.globalConfig = &MppmConfigInfo{}
		configFileManager.createMppmGlobalConfigFileIfNotExists()
		configFileManager.loadMppmConfig(
			configFileManager.globalConfig,
			configFileManager.GetMppmGlobalConfigFilePath(),
		)
	}
	return configFileManager.globalConfig
}

func (configFileManager *MppmConfigFileManager) GetDefaultMppmConfig() (mppmConfig *MppmConfigInfo) {

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

func (configFileManager *MppmConfigFileManager) GetMppmGlobalConfigFilePath() (filePath string) {
	homeDirectoryPath, _ := os.UserHomeDir()
	filePath = filepath.Join(homeDirectoryPath, MppmConfigFileName)
	return
}

func (configFileManager *MppmConfigFileManager) SaveProjectConfig() (err error) {
	err = configFileManager.GetProjectConfig().save(MppmConfigFileName)
	return
}

func (configFileManager *MppmConfigFileManager) SaveGlobalConfig() (err error) {

	configFilePath := configFileManager.GetMppmGlobalConfigFilePath()
	err = configFileManager.GetGlobalConfig().save(configFilePath)
	if err != nil {
		return
	}

	return

}

func (configFileManager *MppmConfigFileManager) SaveDefaultProjectConfig() (err error) {
	err = configFileManager.GetDefaultMppmConfig().save(MppmConfigFileName)
	return
}

func (configFileManager *MppmConfigFileManager) loadMppmConfig(config *MppmConfigInfo, configFilePath string) {

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

func (configFileManager *MppmConfigFileManager) createMppmGlobalConfigFileIfNotExists() {
	mppmGlobalConfigFilePath := configFileManager.GetMppmGlobalConfigFilePath()
	if _, err := os.Stat(mppmGlobalConfigFilePath); os.IsNotExist(err) {
		configFileManager.globalConfig = configFileManager.GetDefaultMppmConfig()
		err = configFileManager.SaveGlobalConfig()
		if err != nil {
			util.ExitWithError(err)
		}
	}
}

type MppmConfigInfo struct {
	Version      string               `json:"version"`
	Applications []*ApplicationConfig `json:"applications"`
	Libraries    []*LibraryConfig     `json:"libraries"`
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
