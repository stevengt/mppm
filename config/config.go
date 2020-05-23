package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

var Version = "1.1.0"
var MppmProjectConfigFileName = ".mppm.json"
var MppmProjectConfig *MppmProjectConfigInfo

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

type MppmProjectConfigInfo struct {
	Version string `json:"version"`
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

type FilePatternsConfig struct {
	GitIgnorePatterns        []string
	GitLfsTrackPatterns      []string
	GzippedXmlFileExtensions []string // List of file extensions that represent Gzipped XML files.
}

func newFilePatternsConfig() (filePatternsConfig *FilePatternsConfig) {
	return &FilePatternsConfig{
		GitIgnorePatterns:        make([]string, 0),
		GitLfsTrackPatterns:      make([]string, 0),
		GzippedXmlFileExtensions: make([]string, 0),
	}
}

func (config1 *FilePatternsConfig) appendAll(config2 *FilePatternsConfig) (filePatternsConfig *FilePatternsConfig) {
	config1.GitIgnorePatterns = append(config1.GitIgnorePatterns, config2.GitIgnorePatterns...)
	config1.GitLfsTrackPatterns = append(config1.GitLfsTrackPatterns, config2.GitLfsTrackPatterns...)
	config1.GzippedXmlFileExtensions = append(config1.GzippedXmlFileExtensions, config2.GzippedXmlFileExtensions...)
	filePatternsConfig = config1
	return
}

func (config *FilePatternsConfig) Print() {
	fmt.Println("Git Ignore Patterns    ---------------------------------")
	fmt.Println(strings.Join(config.GitIgnorePatterns, "\n"))
	fmt.Println("Git LFS Track Patterns ---------------------------------")
	fmt.Println(strings.Join(config.GitLfsTrackPatterns, "\n"))
	fmt.Println("Gzipped XML File Types ---------------------------------")
	fmt.Println(strings.Join(config.GzippedXmlFileExtensions, "\n"))
}

func GetAllFilePatternsConfig() (allFilePatternsConfig *FilePatternsConfig) {

	allFilePatternsConfig = newFilePatternsConfig()

	filePatternsConfigList := []*FilePatternsConfig{
		AudioFilePatternsConfig,
		AbletonFilePatternsConfig,
	}

	for _, filePatternsConfig := range filePatternsConfigList {
		allFilePatternsConfig = allFilePatternsConfig.appendAll(filePatternsConfig)
	}

	return

}
