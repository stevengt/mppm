package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
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

func (config *MppmProjectConfigInfo) IsCompatibleWithInstalledMppmVersion() (isCompatible bool, installedVersion string, configVersion string) {

	installedVersion = Version
	configVersion = config.Version
	installedMajorVersion := strings.Split(installedVersion, ".")[0]
	configMajorVersion := strings.Split(configVersion, ".")[0]

	isCompatible = installedMajorVersion == configMajorVersion

	return
}

func init() {

	configAsJson, err := ioutil.ReadFile(MppmProjectConfigFileName)
	if err != nil {
		return
	}

	MppmProjectConfig = &MppmProjectConfigInfo{}

	err = json.Unmarshal(configAsJson, MppmProjectConfig)
	if err != nil {
		MppmProjectConfig = nil
		fmt.Println("WARN: Invalid config file detected in " + MppmProjectConfigFileName)
	}

}
