package config

import (
	"encoding/json"

	"github.com/stevengt/mppm/config/applications"
	"github.com/stevengt/mppm/util"
)

var MppmConfigFileManager MppmConfigManager = &mppmConfigFileManager{}

type MppmConfigManager interface {
	GetProjectConfig() *MppmConfigInfo
	GetGlobalConfig() *MppmConfigInfo
	GetDefaultMppmConfig() (mppmConfig *MppmConfigInfo)
	GetMppmGlobalConfigFilePath() (filePath string)
	SaveProjectConfig() (err error)
	SaveGlobalConfig() (err error)
	SaveDefaultProjectConfig() (err error)
}

type mppmConfigFileManager struct {
	projectConfig *MppmConfigInfo
	globalConfig  *MppmConfigInfo
}

func (configFileManager *mppmConfigFileManager) GetProjectConfig() *MppmConfigInfo {
	if configFileManager.projectConfig == nil {
		configFileManager.projectConfig = &MppmConfigInfo{}
		configFileManager.loadMppmConfig(
			configFileManager.projectConfig,
			MppmConfigFileName,
		)
	}
	return configFileManager.projectConfig
}

func (configFileManager *mppmConfigFileManager) GetGlobalConfig() *MppmConfigInfo {
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

func (configFileManager *mppmConfigFileManager) GetDefaultMppmConfig() (mppmConfig *MppmConfigInfo) {

	applicationConfigList := make([]*applications.ApplicationConfig, 0)
	libraryConfigList := make([]*LibraryConfig, 0)

	for _, supportedApplication := range applications.SupportedApplications {
		applicationConfig := &applications.ApplicationConfig{
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

func (configFileManager *mppmConfigFileManager) GetMppmGlobalConfigFilePath() (filePath string) {
	homeDirectoryPath, _ := util.UserHomeDir()
	filePath = util.JoinFilePath(homeDirectoryPath, MppmConfigFileName)
	return
}

func (configFileManager *mppmConfigFileManager) SaveProjectConfig() (err error) {
	err = configFileManager.GetProjectConfig().save(MppmConfigFileName)
	return
}

func (configFileManager *mppmConfigFileManager) SaveGlobalConfig() (err error) {

	configFilePath := configFileManager.GetMppmGlobalConfigFilePath()
	err = configFileManager.GetGlobalConfig().save(configFilePath)
	if err != nil {
		return
	}

	return

}

func (configFileManager *mppmConfigFileManager) SaveDefaultProjectConfig() (err error) {
	err = configFileManager.GetDefaultMppmConfig().save(MppmConfigFileName)
	return
}

func (configFileManager *mppmConfigFileManager) loadMppmConfig(config *MppmConfigInfo, configFilePath string) {

	configFile, err := util.OpenFile(configFilePath)
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

	err = config.checkIfCompatibleWithInstalledMppmVersion()
	if err != nil {
		util.ExitWithError(err)
	}

	err = config.checkIfCompatibleWithSupportedApplications()
	if err != nil {
		util.ExitWithError(err)
	}

}

func (configFileManager *mppmConfigFileManager) createMppmGlobalConfigFileIfNotExists() {
	mppmGlobalConfigFilePath := configFileManager.GetMppmGlobalConfigFilePath()
	if !util.DoesFileExist(mppmGlobalConfigFilePath) {
		configFileManager.globalConfig = configFileManager.GetDefaultMppmConfig()
		err := configFileManager.SaveGlobalConfig()
		if err != nil {
			util.ExitWithError(err)
		}
	}
}
