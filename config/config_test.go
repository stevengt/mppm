package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stevengt/mppm/config"

	"github.com/stevengt/mppm/config/configtest"
)

func TestGetFilePatternsConfigListFromProjectConfig(t *testing.T) {

	for _, testCase := range configtest.TestMppmConfigInfoAndExpectedConfigFunctionResponses {

		configAsJson := testCase.ConfigAsJson
		configtest.InitMockFileSystemDelegaterWithConfigFiles(configAsJson, configAsJson)

		expectedError := testCase.ExpectedError
		expectedFilePatternsConfigList := testCase.ExpectedFilePatternsConfigList
		if expectedFilePatternsConfigList != nil {
			for _, filePatternConfig := range expectedFilePatternsConfigList {
				filePatternConfig.SortAllLists()
			}
		}

		actualFilePatternsConfigList, actualError := config.GetFilePatternsConfigListFromProjectConfig()
		if actualFilePatternsConfigList != nil {
			for _, filePatternConfig := range actualFilePatternsConfigList {
				filePatternConfig.SortAllLists()
			}
		}

		assert.Exactly(t, expectedFilePatternsConfigList, actualFilePatternsConfigList)
		assert.Exactly(t, expectedError, actualError)

	}

}

func TestGetAllFilePatternsConfigFromProjectConfig(t *testing.T) {

	for _, testCase := range configtest.TestMppmConfigInfoAndExpectedConfigFunctionResponses {

		configAsJson := testCase.ConfigAsJson
		configtest.InitMockFileSystemDelegaterWithConfigFiles(configAsJson, configAsJson)

		expectedError := testCase.ExpectedError
		expectedFilePatternsConfig := testCase.ExpectedFilePatternsConfig
		if expectedFilePatternsConfig != nil {
			expectedFilePatternsConfig.SortAllLists()
		}

		actualFilePatternsConfig, actualError := config.GetAllFilePatternsConfigFromProjectConfig()
		if actualFilePatternsConfig != nil {
			actualFilePatternsConfig.SortAllLists()
		}

		assert.Exactly(t, expectedFilePatternsConfig, actualFilePatternsConfig)
		assert.Exactly(t, expectedError, actualError)

	}

}
