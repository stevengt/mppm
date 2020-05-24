package config

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/stevengt/mppm/util"
)

var MppmProjectConfig *MppmProjectConfigInfo

var MppmProjectConfigFileName = ".mppm.json"

type MppmProjectConfigInfo struct {
	Version      string               `json:"version"`
	Applications []*ApplicationConfig `json:"applications"`
}

func (config *MppmProjectConfigInfo) Save() (err error) {

	configAsJson, err := json.Marshal(config)
	if err != nil {
		return
	}

	filePermissionsCode := os.FileMode(0644)
	err = ioutil.WriteFile(MppmProjectConfigFileName, configAsJson, filePermissionsCode)
	if err != nil {
		return
	}

	return

}

func (config *MppmProjectConfigInfo) CheckIfCompatibleWithInstalledMppmVersion() (err error) {

	installedVersion := Version
	configVersion := config.Version

	installedMajorVersion := strings.Split(installedVersion, ".")[0]
	configMajorVersion := strings.Split(configVersion, ".")[0]

	isCompatible := installedMajorVersion == configMajorVersion

	if !isCompatible {
		err = errors.New("Installed mppm version " + installedVersion +
			" is not compatible with this project's configured version " + configVersion)
	}

	return
}

func (config *MppmProjectConfigInfo) CheckIfCompatibleWithSupportedApplications() (err error) {

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
			errorMessage := fmt.Sprintf("Found unsupported application %s %s in %s", application.Name, application.Version, MppmProjectConfigFileName)
			err = errors.New(errorMessage)
			return
		}

	}

	return

}

func LoadMppmProjectConfig() {

	configFile, err := os.Open(MppmProjectConfigFileName)
	if err != nil {
		errorMessage := openingMppmProjectConfigFileErrorMessage + err.Error()
		util.ExitWithErrorMessage(errorMessage)
	}
	defer configFile.Close()

	MppmProjectConfig = &MppmProjectConfigInfo{}

	jsonDecoder := json.NewDecoder(configFile)
	jsonDecoder.DisallowUnknownFields()

	err = jsonDecoder.Decode(MppmProjectConfig)
	if err != nil {
		errorMessage := invalidMppmProjectConfigFileErrorMessage + err.Error()
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
