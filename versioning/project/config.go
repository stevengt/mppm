package project

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

const ConfigFileName string = ".mppm-config.json"

type MppmConfig struct {
	ProjectType         ProjectType `json:"project-type"`
	GitLfsTrackPatterns []string    `json:"git-lfs-track-patterns"`
}

func LoadMppmConfig() (config *MppmConfig, err error) {
	config = &MppmConfig{}
	configFileContents, err := ioutil.ReadFile(ConfigFileName)
	if err == nil {
		err = json.Unmarshal(configFileContents, config)
	}
	return
}

func (config *MppmConfig) Save() (err error) {
	configAsJson, err := json.Marshal(config)
	if err != nil {
		return
	}
	filePermissionsCode := os.FileMode(0644)
	err = ioutil.WriteFile(ConfigFileName, configAsJson, filePermissionsCode)
	return
}

func AddGitLfsTrackPatternsToMppmConfig(trackPatterns ...string) (err error) {
	config, err := LoadMppmConfig()
	if err != nil {
		return
	}
	config.GitLfsTrackPatterns = append(config.GitLfsTrackPatterns, trackPatterns...)
	err = config.Save()
	return
}
