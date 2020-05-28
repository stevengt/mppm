package config

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"os"
	"strings"

	"github.com/stevengt/mppm/util"
)

var MppmProjectConfig *MppmConfigInfo

var MppmConfigFileName = ".mppm.json"

type MppmConfigInfo struct {
	Version      string               `json:"version"`
	Applications []*ApplicationConfig `json:"applications"`
}

func (config *MppmConfigInfo) Save() (err error) {

	configAsJson, err := json.Marshal(config)
	if err != nil {
		return
	}

	filePermissionsCode := os.FileMode(0644)
	err = ioutil.WriteFile(MppmConfigFileName, configAsJson, filePermissionsCode)
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

func LoadMppmProjectConfig() {

	configFile, err := os.Open(MppmConfigFileName)
	if err != nil {
		errorMessage := getOpeningMppmProjectConfigFileErrorMessage(err)
		util.ExitWithErrorMessage(errorMessage)
	}
	defer configFile.Close()

	MppmProjectConfig = &MppmConfigInfo{}

	jsonDecoder := json.NewDecoder(configFile)
	jsonDecoder.DisallowUnknownFields()

	err = jsonDecoder.Decode(MppmProjectConfig)
	if err != nil {
		errorMessage := getInvalidMppmProjectConfigFileErrorMessage(err)
		util.ExitWithErrorMessage(errorMessage)
	}

	err = MppmProjectConfig.CheckIfCompatibleWithInstalledMppmVersion()
	if err != nil {
		util.ExitWithError(err)
	}

	err = MppmProjectConfig.CheckIfCompatibleWithSupportedApplications()
	if err != nil {
		util.ExitWithError(err)
	}

}

func GetDefaultMppmProjectConfig() (mppmProjectConfig *MppmConfigInfo) {

	applicationConfigList := make([]*ApplicationConfig, 0)

	for _, supportedApplication := range SupportedApplications {
		applicationConfig := &ApplicationConfig{
			Name:    supportedApplication.Name,
			Version: supportedApplication.DefaultVersion,
		}
		applicationConfigList = append(applicationConfigList, applicationConfig)
	}

	mppmProjectConfig = &MppmConfigInfo{
		Version:      Version,
		Applications: applicationConfigList,
	}

	return

}
