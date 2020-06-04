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

// func TestSaveProjectConfig(t *testing.T) {}

// func TestSaveGlobalConfig(t *testing.T) {}

// func TestSaveDefaultProjectConfig(t *testing.T) {}
