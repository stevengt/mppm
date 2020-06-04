package config

import (
	"encoding/json"
	"errors"

	"github.com/stevengt/mppm/config/applications"
	"github.com/stevengt/mppm/util"
)

var MppmConfigFileManager MppmConfigManager = &mppmConfigFileManager{}

type MppmConfigManager interface {
	GetProjectConfig() (projectConfig *MppmConfigInfo, err error)
	GetGlobalConfig() (globalConfig *MppmConfigInfo, err error)
	GetProjectAndGlobalConfigs() (projectConfig *MppmConfigInfo, globalConfig *MppmConfigInfo, err error)
	GetDefaultMppmConfig() (mppmConfig *MppmConfigInfo)
	GetMppmGlobalConfigFilePath() (filePath string, err error)
	SaveProjectConfig() (err error)
	SaveGlobalConfig() (err error)
	SaveDefaultProjectConfig() (err error)
}

type mppmConfigFileManager struct {
	projectConfig *MppmConfigInfo
	globalConfig  *MppmConfigInfo
}

func NewMppmConfigFileManager() *mppmConfigFileManager {
	return &mppmConfigFileManager{}
}

func (configFileManager *mppmConfigFileManager) GetProjectConfig() (projectConfig *MppmConfigInfo, err error) {
	if configFileManager.projectConfig == nil {
		configFileManager.projectConfig = &MppmConfigInfo{}
		err = configFileManager.loadMppmConfig(configFileManager.projectConfig, MppmConfigFileName)
		if err != nil {
			return
		}
	}
	projectConfig = configFileManager.projectConfig
	return
}

func (configFileManager *mppmConfigFileManager) GetGlobalConfig() (globalConfig *MppmConfigInfo, err error) {
	if configFileManager.globalConfig == nil {
		configFileManager.globalConfig = &MppmConfigInfo{}

		err = configFileManager.createMppmGlobalConfigFileIfNotExists()
		if err != nil {
			return nil, err
		}

		globalConfigFilePath, err := configFileManager.GetMppmGlobalConfigFilePath()
		if err != nil {
			return nil, err
		}

		err = configFileManager.loadMppmConfig(configFileManager.globalConfig, globalConfigFilePath)
		if err != nil {
			return nil, err
		}

	}
	globalConfig = configFileManager.globalConfig
	return
}

func (configFileManager *mppmConfigFileManager) GetProjectAndGlobalConfigs() (projectConfig *MppmConfigInfo, globalConfig *MppmConfigInfo, err error) {

	projectConfig, err = configFileManager.GetProjectConfig()
	if err != nil {
		return
	}

	globalConfig, err = configFileManager.GetGlobalConfig()
	if err != nil {
		return
	}

	return

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

func (configFileManager *mppmConfigFileManager) GetMppmGlobalConfigFilePath() (filePath string, err error) {
	homeDirectoryPath, err := util.UserHomeDir()
	if err != nil {
		return
	}
	filePath = util.JoinFilePath(homeDirectoryPath, MppmConfigFileName)
	return
}

func (configFileManager *mppmConfigFileManager) SaveProjectConfig() (err error) {
	if configFileManager.projectConfig == nil {
		err = errors.New("Unable to save uninitialized project config.")
		return
	}
	err = configFileManager.projectConfig.save(MppmConfigFileName)
	return
}

func (configFileManager *mppmConfigFileManager) SaveGlobalConfig() (err error) {

	configFilePath, err := configFileManager.GetMppmGlobalConfigFilePath()
	if err != nil {
		return
	}

	globalConfig, err := configFileManager.GetGlobalConfig()
	if err != nil {
		return
	}

	err = globalConfig.save(configFilePath)
	if err != nil {
		return
	}

	return

}

func (configFileManager *mppmConfigFileManager) SaveDefaultProjectConfig() (err error) {
	err = configFileManager.GetDefaultMppmConfig().save(MppmConfigFileName)
	return
}

func (configFileManager *mppmConfigFileManager) loadMppmConfig(config *MppmConfigInfo, configFilePath string) (err error) {

	configFile, err := util.OpenFile(configFilePath)
	if err != nil {
		errorMessage := getOpeningMppmProjectConfigFileErrorMessage(err)
		err = errors.New(errorMessage)
		return
	}
	defer configFile.Close()

	jsonDecoder := json.NewDecoder(configFile)
	jsonDecoder.DisallowUnknownFields()

	err = jsonDecoder.Decode(config)
	if err != nil {
		errorMessage := getInvalidMppmProjectConfigFileErrorMessage(err)
		err = errors.New(errorMessage)
		return
	}

	err = config.checkIfCompatibleWithInstalledMppmVersion()
	if err != nil {
		return
	}

	err = config.checkIfCompatibleWithSupportedApplications()
	if err != nil {
		return
	}

	return

}

func (configFileManager *mppmConfigFileManager) createMppmGlobalConfigFileIfNotExists() (err error) {

	mppmGlobalConfigFilePath, err := configFileManager.GetMppmGlobalConfigFilePath()
	if err != nil {
		return
	}

	if !util.DoesFileExist(mppmGlobalConfigFilePath) {
		configFileManager.globalConfig = configFileManager.GetDefaultMppmConfig()
		err = configFileManager.SaveGlobalConfig()
		if err != nil {
			return
		}
	}

	return

}
