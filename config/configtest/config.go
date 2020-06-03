package configtest

import (
	"github.com/stevengt/mppm/config"
	"github.com/stevengt/mppm/config/applications"
	"github.com/stevengt/mppm/util"
	"github.com/stevengt/mppm/util/utiltest"
)

var TestMppmConfigInfosAsJson map[string][]byte = map[string][]byte{
	"valid version, valid application name, valid application version":   []byte(`{"version":"1.0.0","applications":[{"name":"Ableton","version":"10"}]}`),
	"valid version, no applications":                                     []byte(`{"version":"1.0.0","applications":[]}`),
	"invalid version, no applications":                                   []byte(`{"version":"0.0.0","applications":[]}`),
	"valid version, invalid application name":                            []byte(`{"version":"1.0.0","applications":[{"name":"Fake Application","version":"1"}]}`),
	"valid version, valid application name, invalid application version": []byte(`{"version":"1.0.0","applications":[{"name":"Ableton","version":"-1"}]}`),
}

func InitMockFileSystemDelegaterWithDefaultConfigFiles() {

	_, projectConfigAsJson := GetTestMppmConfigInfo()
	_, globalConfigAsJson := GetTestMppmConfigInfo()
	InitMockFileSystemDelegaterWithConfigFiles(projectConfigAsJson, globalConfigAsJson)

}

func InitMockFileSystemDelegaterWithConfigFiles(projectConfigAsJson []byte, globalConfigAsJson []byte) {

	mockFileSystemDelegater := &utiltest.MockFileSystemDelegater{}

	projectConfigFilePath := config.MppmConfigFileName
	userHomeDirectoryFilePath, _ := mockFileSystemDelegater.UserHomeDir()
	globalConfigFilePath := mockFileSystemDelegater.JoinFilePath(userHomeDirectoryFilePath, config.MppmConfigFileName)

	projectConfigFile := utiltest.NewMockFile(projectConfigAsJson)
	globalConfigFile := utiltest.NewMockFile(globalConfigAsJson)

	mockFileSystemDelegater.Files = map[string]*utiltest.MockFile{
		projectConfigFilePath: projectConfigFile,
		globalConfigFilePath:  globalConfigFile,
	}

	util.FileSystemProxy = mockFileSystemDelegater
	config.MppmConfigFileManager = config.NewMppmConfigFileManager()

}

func GetTestMppmConfigInfo() (testMppmConfigInfo *config.MppmConfigInfo, configAsJson []byte) {

	testMppmConfigInfo = &config.MppmConfigInfo{
		Version: "1.0.0",
		Applications: []*applications.ApplicationConfig{
			&applications.ApplicationConfig{
				Name:    "Ableton",
				Version: "10",
			},
		},
		Libraries: []*config.LibraryConfig{
			&config.LibraryConfig{
				FilePath:              "/home/testuser/library",
				MostRecentGitCommitId: "56789",
				CurrentGitCommitId:    "01234",
			},
		},
	}

	configAsJson = []byte(`{"version":"1.0.0","applications":[{"name":"Ableton","version":"10"}],"libraries":[{"location":"/home/testuser/library","most-recent-version":"56789","current-version":"01234"}]}`)

	return

}
