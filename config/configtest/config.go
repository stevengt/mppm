package configtest

import (
	"errors"
	"fmt"

	"github.com/stevengt/mppm/config"
	"github.com/stevengt/mppm/config/applications"
	"github.com/stevengt/mppm/util"
	"github.com/stevengt/mppm/util/utiltest"
)

// ------------------------------------------------------------------------------

var ConfigWithValidVersionAndApplicationNameAndApplicationVersion *MppmConfigInfoAndExpectedError = &MppmConfigInfoAndExpectedError{
	ConfigAsJson: []byte(
		fmt.Sprintf(
			`{"version":"%s.0.0","applications":[{"name":"Ableton","version":"10"}],"libraries":null}`,
			config.GetCurrentlyInstalledMajorVersion(),
		),
	),
	ExpectedError: nil,
}

var ConfigWithValidVersionAndNoApplications *MppmConfigInfoAndExpectedError = &MppmConfigInfoAndExpectedError{
	ConfigAsJson: []byte(
		fmt.Sprintf(
			`{"version":"%s.0.0","applications":[],"libraries":null}`,
			config.GetCurrentlyInstalledMajorVersion(),
		),
	),
	ExpectedError: nil,
}

var ConfigWithInvalidVersionAndNoApplications *MppmConfigInfoAndExpectedError = &MppmConfigInfoAndExpectedError{
	ConfigAsJson: []byte(`{"version":"0.0.0","applications":[],"libraries":null}`),
	ExpectedError: errors.New(
		fmt.Sprintf(
			"Installed mppm version %s is not compatible with this project's configured version 0.0.0",
			config.Version,
		),
	),
}

var ConfigWithValidVersionAndInvalidApplicationName *MppmConfigInfoAndExpectedError = &MppmConfigInfoAndExpectedError{
	ConfigAsJson: []byte(
		fmt.Sprintf(
			`{"version":"%s.0.0","applications":[{"name":"Fake Application","version":"1"}],"libraries":null}`,
			config.GetCurrentlyInstalledMajorVersion(),
		),
	),
	ExpectedError: errors.New("\nFound unsupported application Fake Application 1 in config file .mppm.json\nTo see what applications are supported, please run 'mppm --show-supported'.\n"),
}

var ConfigWithValidVersionAndApplicationNameAndInvalidApplicationVersion *MppmConfigInfoAndExpectedError = &MppmConfigInfoAndExpectedError{
	ConfigAsJson: []byte(
		fmt.Sprintf(
			`{"version":"%s.0.0","applications":[{"name":"Ableton","version":"1"}],"libraries":null}`,
			config.GetCurrentlyInstalledMajorVersion(),
		),
	),
	ExpectedError: errors.New("\nFound unsupported application Ableton 1 in config file .mppm.json\nTo see what applications are supported, please run 'mppm --show-supported'.\n"),
}

// ------------------------------------------------------------------------------

func GetDefaultTestMppmConfigInfo() (testMppmConfigInfo *config.MppmConfigInfo, configAsJson []byte) {

	testMppmConfigInfo = &config.MppmConfigInfo{
		Version: fmt.Sprintf("%s.0.0", config.GetCurrentlyInstalledMajorVersion()),
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

	configAsJson = []byte(
		fmt.Sprintf(
			`{"version":"%s.0.0","applications":[{"name":"Ableton","version":"10"}],"libraries":[{"location":"/home/testuser/library","most-recent-version":"56789","current-version":"01234"}]}`,
			config.GetCurrentlyInstalledMajorVersion(),
		),
	)

	return

}

// ------------------------------------------------------------------------------

// Updates util.FileSystemProxy and config.MppmConfigFileManager so that any
// previously loaded config files are discarded, and no new config files are available.
func InitAndReturnMockFileSystemDelegaterWithNoConfigFiles() *utiltest.MockFileSystemDelegater {
	mockFileSystemDelegater := &utiltest.MockFileSystemDelegater{}
	InitMockFileSystemDelegaterWithConfigFiles(mockFileSystemDelegater, nil, nil)
	return mockFileSystemDelegater
}

// Updates util.FileSystemProxy and config.MppmConfigFileManager so that any
// previously loaded config files are discarded. Instead, new default config files
// are available with appropriate file names.
//
// The config file contents are the same as the results from calling configtest.GetDefaultTestMppmConfigInfo().
func InitAndReturnMockFileSystemDelegaterWithDefaultConfigFiles() *utiltest.MockFileSystemDelegater {

	mockFileSystemDelegater := &utiltest.MockFileSystemDelegater{}

	_, projectConfigAsJson := GetDefaultTestMppmConfigInfo()
	_, globalConfigAsJson := GetDefaultTestMppmConfigInfo()

	InitMockFileSystemDelegaterWithConfigFiles(
		mockFileSystemDelegater,
		utiltest.NewMockFile(projectConfigAsJson),
		utiltest.NewMockFile(globalConfigAsJson),
	)

	return mockFileSystemDelegater

}

// Adds project and global config files to the mock file system with appropriate file names,
// and then updates util.FileSystemProxy and config.MppmConfigFileManager so that any
// previously loaded config files are discarded, and the new ones are used instead.
//
// If mockFileSystemDelegater.Files is nil, a new empty set of files will be created
// before adding the config files.
func InitMockFileSystemDelegaterWithConfigFiles(mockFileSystemDelegater *utiltest.MockFileSystemDelegater, projectConfigFile *utiltest.MockFile, globalConfigFile *utiltest.MockFile) {

	if mockFileSystemDelegater.Files == nil {
		mockFileSystemDelegater.Files = make(map[string]*utiltest.MockFile)
	}

	if projectConfigFile != nil {
		projectConfigFilePath := config.MppmConfigFileName
		mockFileSystemDelegater.Files[projectConfigFilePath] = projectConfigFile
	}

	if globalConfigFile != nil {
		userHomeDirectoryFilePath, _ := mockFileSystemDelegater.UserHomeDir()
		globalConfigFilePath := mockFileSystemDelegater.JoinFilePath(userHomeDirectoryFilePath, config.MppmConfigFileName)
		mockFileSystemDelegater.Files[globalConfigFilePath] = globalConfigFile
	}

	util.FileSystemProxy = mockFileSystemDelegater
	config.MppmConfigFileManager = config.NewMppmConfigFileManager()

}

// ------------------------------------------------------------------------------

// Returns expectedError if it is not nil.
// Otherwise, if mppmConfigInfoAndExpectedError is not nil, returns mppmConfigInfoAndExpectedError.ExpectedError.
// Otherwise, returns nil.
func GetExpectedError(expectedError error, mppmConfigInfoAndExpectedError *MppmConfigInfoAndExpectedError) error {
	if expectedError != nil {
		return expectedError
	} else if mppmConfigInfoAndExpectedError != nil {
		return mppmConfigInfoAndExpectedError.ExpectedError
	} else {
		return nil
	}
}

func ReturnMppmConfigInfoAsMockFileIfNotNilElseReturnNil(mppmConfigInfoAndExpectedError *MppmConfigInfoAndExpectedError) *utiltest.MockFile {
	if mppmConfigInfoAndExpectedError != nil {
		return mppmConfigInfoAndExpectedError.AsMockFile()
	} else {
		return nil
	}
}

func ReturnMppmConfigInfoAsJsonIfNotNilAndErrorIsNilElseReturnNil(mppmConfigInfoAndExpectedError *MppmConfigInfoAndExpectedError, err error) []byte {
	if mppmConfigInfoAndExpectedError != nil && err == nil {
		return mppmConfigInfoAndExpectedError.ConfigAsJson
	} else {
		return nil
	}
}

// ------------------------------------------------------------------------------

// Contains the contents of a mppm config file as a JSON byte array,
// and the expected error to receive when attempting to load the config file.
type MppmConfigInfoAndExpectedError struct {
	ConfigAsJson  []byte
	ExpectedError error
}

func (mppmConfigInfoAndExpectedError *MppmConfigInfoAndExpectedError) AsMockFile() *utiltest.MockFile {
	return utiltest.NewMockFile(mppmConfigInfoAndExpectedError.ConfigAsJson)
}
