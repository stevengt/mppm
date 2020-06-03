package config_test

import (
	"testing"

	"github.com/stevengt/mppm/config"
	"github.com/stevengt/mppm/config/configtest"
	"github.com/stretchr/testify/assert"
)

func TestGetProjectConfig(t *testing.T) {

	expectedConfigInfo, configAsJson := configtest.GetTestMppmConfigInfo()
	configtest.InitMockFileSystemDelegaterWithConfigFiles(configAsJson, configAsJson)
	configManager := config.MppmConfigFileManager
	actualConfigInfo, actualError := configManager.GetProjectConfig()
	assert.Nil(t, actualError)
	assert.NotNil(t, actualConfigInfo)
	assert.Exactly(t, expectedConfigInfo, actualConfigInfo)

	for _, testCase := range configtest.TestMppmConfigInfoAsJsonAndExpectedConfigFunctionResponses {

		configAsJson = testCase.ConfigAsJson
		configtest.InitMockFileSystemDelegaterWithConfigFiles(configAsJson, configAsJson)
		configManager = config.MppmConfigFileManager

		expectedError := testCase.ExpectedError

		_, actualError = configManager.GetProjectConfig()

		assert.Exactly(t, expectedError, actualError)

	}

}

func TestGetGlobalConfig(t *testing.T) {

	expectedConfigInfo, configAsJson := configtest.GetTestMppmConfigInfo()
	configtest.InitMockFileSystemDelegaterWithConfigFiles(configAsJson, configAsJson)
	configManager := config.MppmConfigFileManager
	actualConfigInfo, actualError := configManager.GetGlobalConfig()
	assert.Nil(t, actualError)
	assert.NotNil(t, actualConfigInfo)
	assert.Exactly(t, expectedConfigInfo, actualConfigInfo)

	for _, testCase := range configtest.TestMppmConfigInfoAsJsonAndExpectedConfigFunctionResponses {

		configAsJson = testCase.ConfigAsJson
		configtest.InitMockFileSystemDelegaterWithConfigFiles(configAsJson, configAsJson)
		configManager = config.MppmConfigFileManager

		expectedError := testCase.ExpectedError
		_, actualError = configManager.GetGlobalConfig()

		assert.Exactly(t, expectedError, actualError)

	}

}

// func TestGetDefaultMppmConfig(t *testing.T) {}

// func TestGetMppmGlobalConfigFilePath(t *testing.T) {}

// func TestSaveProjectConfig(t *testing.T) {}

// func TestSaveGlobalConfig(t *testing.T) {}

// func TestSaveDefaultProjectConfig(t *testing.T) {}
