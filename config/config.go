package config

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"
	"strings"

	"github.com/stevengt/mppm/util"
)

var Version = "1.2.1"

var MppmConfigFileName = ".mppm.json"

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
