package config_test

import (
	"errors"
	"testing"

	"github.com/stevengt/mppm/util"
	"github.com/stevengt/mppm/util/utiltest"

	"github.com/stevengt/mppm/config/applications"

	"github.com/stevengt/mppm/config"
	"github.com/stevengt/mppm/config/configtest"
	"github.com/stretchr/testify/assert"
)

func TestGetProjectConfig(t *testing.T) {

	for _, testCase := range configtest.TestMppmConfigInfoAndExpectedConfigFunctionResponses {

		configAsJson := testCase.ConfigAsJson
		configtest.InitMockFileSystemDelegaterWithConfigFiles(configAsJson, configAsJson)
		configManager := config.MppmConfigFileManager

		expectedConfigInfo := testCase.ConfigInfo
		expectedError := testCase.ExpectedError

		actualConfigInfo, actualError := configManager.GetProjectConfig()

		assert.Exactly(t, expectedConfigInfo, actualConfigInfo)
		assert.Exactly(t, expectedError, actualError)

	}

	configtest.InitMockFileSystemDelegaterWithNoConfigFiles()
	configManager := config.MppmConfigFileManager
	expectedError := errors.New(`
There was a problem while opening the mppm config file.
If the file doesn't exist, try running 'mppm project init' first.
Unable to open file.mppm.json
`)
	actualConfigInfo, actualError := configManager.GetProjectConfig()
	assert.Nil(t, actualConfigInfo)
	assert.Exactly(t, expectedError, actualError)

}

func TestGetGlobalConfig(t *testing.T) {

	for _, testCase := range configtest.TestMppmConfigInfoAndExpectedConfigFunctionResponses {

		configAsJson := testCase.ConfigAsJson
		configtest.InitMockFileSystemDelegaterWithConfigFiles(configAsJson, configAsJson)
		configManager := config.MppmConfigFileManager

		expectedConfigInfo := testCase.ConfigInfo
		expectedError := testCase.ExpectedError

		actualConfigInfo, actualError := configManager.GetGlobalConfig()

		assert.Exactly(t, expectedConfigInfo, actualConfigInfo)
		assert.Exactly(t, expectedError, actualError)

	}

	configtest.InitMockFileSystemDelegaterWithNoConfigFiles()
	configManager := config.MppmConfigFileManager
	expectedConfigInfo := configManager.GetDefaultMppmConfig()
	actualConfigInfo, actualError := configManager.GetGlobalConfig()
	assert.Exactly(t, expectedConfigInfo, actualConfigInfo)
	assert.Nil(t, actualError)
	config.MppmConfigFileManager = config.NewMppmConfigFileManager()
	configManager = config.MppmConfigFileManager
	actualConfigInfo, actualError = configManager.GetGlobalConfig()
	assert.Exactly(t, expectedConfigInfo, actualConfigInfo)
	assert.Nil(t, actualError)

	mockFileSystemDelegater := &utiltest.MockFileSystemDelegater{
		Files:         make(map[string]*utiltest.MockFile),
		OpenFileError: errors.New("There was a problem opening the file."),
	}
	util.FileSystemProxy = mockFileSystemDelegater
	config.MppmConfigFileManager = config.NewMppmConfigFileManager()
	configManager = config.MppmConfigFileManager
	expectedError := errors.New(`
There was a problem while opening the mppm config file.
If the file doesn't exist, try running 'mppm project init' first.
There was a problem opening the file.
`)
	actualConfigInfo, actualError = configManager.GetGlobalConfig()
	assert.Nil(t, actualConfigInfo)
	assert.Exactly(t, expectedError, actualError)

}

func TestGetDefaultMppmConfig(t *testing.T) {

	expectedConfigInfo := &config.MppmConfigInfo{
		Version: config.Version,
		Applications: []*applications.ApplicationConfig{
			&applications.ApplicationConfig{
				Name:    "Ableton",
				Version: "10",
			},
		},
		Libraries: make([]*config.LibraryConfig, 0),
	}

	configManager := config.MppmConfigFileManager
	actualConfigInfo := configManager.GetDefaultMppmConfig()

	assert.Exactly(t, expectedConfigInfo, actualConfigInfo)

}

func TestGetMppmGlobalConfigFilePath(t *testing.T) {

	var expectedError error

	mockFileSystemDelegater := &utiltest.MockFileSystemDelegater{}
	util.FileSystemProxy = mockFileSystemDelegater
	configManager := config.MppmConfigFileManager
	expectedFilePath := "/home/testuser/.mppm.json"
	expectedError = nil
	actualFilePath, actualError := configManager.GetMppmGlobalConfigFilePath()
	assert.Exactly(t, expectedFilePath, actualFilePath)
	assert.Exactly(t, expectedError, actualError)

	mockFileSystemDelegater = &utiltest.MockFileSystemDelegater{
		UserHomeDirError: errors.New("There was a problem getting the user's home directory."),
	}
	util.FileSystemProxy = mockFileSystemDelegater
	configManager = config.MppmConfigFileManager
	expectedFilePath = ""
	expectedError = mockFileSystemDelegater.UserHomeDirError
	actualFilePath, actualError = configManager.GetMppmGlobalConfigFilePath()
	assert.Exactly(t, expectedFilePath, actualFilePath)
	assert.Exactly(t, expectedError, actualError)

}

func TestSaveProjectConfig(t *testing.T) {

	configtest.InitMockFileSystemDelegaterWithDefaultConfigFiles()
	configManager := config.MppmConfigFileManager
	expectedError := errors.New("Unable to save uninitialized project config.")
	actualError := configManager.SaveProjectConfig()
	assert.Exactly(t, expectedError, actualError)

	configtest.InitMockFileSystemDelegaterWithDefaultConfigFiles()
	configManager = config.MppmConfigFileManager
	projectConfig, actualError := configManager.GetProjectConfig()
	assert.Nil(t, actualError)
	assert.NotNil(t, projectConfig)
	projectConfig.Version = "1.9999.9999"
	expectedError = nil
	actualError = configManager.SaveProjectConfig()
	assert.Exactly(t, expectedError, actualError)
	config.MppmConfigFileManager = config.NewMppmConfigFileManager()
	configManager = config.MppmConfigFileManager
	projectConfig, actualError = configManager.GetProjectConfig()
	assert.Nil(t, actualError)
	assert.NotNil(t, projectConfig)
	assert.Equal(t, "1.9999.9999", projectConfig.Version)

}

func TestSaveGlobalConfig(t *testing.T) {

	configtest.InitMockFileSystemDelegaterWithDefaultConfigFiles()
	configManager := config.MppmConfigFileManager
	expectedError := errors.New("Unable to save uninitialized global config.")
	actualError := configManager.SaveGlobalConfig()
	assert.Exactly(t, expectedError, actualError)

	configtest.InitMockFileSystemDelegaterWithDefaultConfigFiles()
	configManager = config.MppmConfigFileManager
	globalConfig, actualError := configManager.GetGlobalConfig()
	assert.Nil(t, actualError)
	assert.NotNil(t, globalConfig)
	globalConfig.Version = "1.9999.9999"
	expectedError = nil
	actualError = configManager.SaveGlobalConfig()
	assert.Exactly(t, expectedError, actualError)
	config.MppmConfigFileManager = config.NewMppmConfigFileManager()
	configManager = config.MppmConfigFileManager
	globalConfig, actualError = configManager.GetGlobalConfig()
	assert.Nil(t, actualError)
	assert.NotNil(t, globalConfig)
	assert.Equal(t, "1.9999.9999", globalConfig.Version)

}

func TestSaveDefaultProjectConfig(t *testing.T) {

	mockFileSystemDelegater := &utiltest.MockFileSystemDelegater{
		Files: make(map[string]*utiltest.MockFile),
	}
	util.FileSystemProxy = mockFileSystemDelegater
	configManager := config.MppmConfigFileManager

	_, actualError := configManager.GetProjectConfig()
	assert.NotNil(t, actualError)
	actualError = configManager.SaveDefaultProjectConfig()
	assert.Nil(t, actualError)

	config.MppmConfigFileManager = config.NewMppmConfigFileManager()
	configManager = config.MppmConfigFileManager

	expectedProjectConfig := configManager.GetDefaultMppmConfig()
	actualProjectConfig, actualError := configManager.GetProjectConfig()
	assert.Exactly(t, expectedProjectConfig, actualProjectConfig)
	assert.Nil(t, actualError)

}
