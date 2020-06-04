package config_test

import (
	"testing"

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

// func TestGetDefaultMppmConfig(t *testing.T) {}

// func TestGetMppmGlobalConfigFilePath(t *testing.T) {}

// func TestSaveProjectConfig(t *testing.T) {}

// func TestSaveGlobalConfig(t *testing.T) {}

// func TestSaveDefaultProjectConfig(t *testing.T) {}
