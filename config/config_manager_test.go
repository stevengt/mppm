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
	actualConfigInfo, err := configManager.GetProjectConfig()
	assert.Nil(t, err)
	assert.NotNil(t, actualConfigInfo)
	assert.Exactly(t, expectedConfigInfo, actualConfigInfo)

	configAsJson = configtest.TestMppmConfigInfosAsJson["invalid version, no applications"]
	configtest.InitMockFileSystemDelegaterWithConfigFiles(configAsJson, configAsJson)
	configManager = config.MppmConfigFileManager
	actualConfigInfo, err = configManager.GetProjectConfig()
	assert.NotNil(t, err)

	configAsJson = configtest.TestMppmConfigInfosAsJson["valid version, invalid application name"]
	configtest.InitMockFileSystemDelegaterWithConfigFiles(configAsJson, configAsJson)
	configManager = config.MppmConfigFileManager
	actualConfigInfo, err = configManager.GetProjectConfig()
	assert.NotNil(t, err)

	configAsJson = configtest.TestMppmConfigInfosAsJson["valid version, valid application name, invalid application version"]
	configtest.InitMockFileSystemDelegaterWithConfigFiles(configAsJson, configAsJson)
	configManager = config.MppmConfigFileManager
	actualConfigInfo, err = configManager.GetProjectConfig()
	assert.NotNil(t, err)

}

func TestGetGlobalConfig(t *testing.T) {

	expectedConfigInfo, configAsJson := configtest.GetTestMppmConfigInfo()
	configtest.InitMockFileSystemDelegaterWithConfigFiles(configAsJson, configAsJson)
	configManager := config.MppmConfigFileManager
	actualConfigInfo, err := configManager.GetGlobalConfig()
	assert.Nil(t, err)
	assert.NotNil(t, actualConfigInfo)
	assert.Exactly(t, expectedConfigInfo, actualConfigInfo)

	configAsJson = configtest.TestMppmConfigInfosAsJson["invalid version, no applications"]
	configtest.InitMockFileSystemDelegaterWithConfigFiles(configAsJson, configAsJson)
	configManager = config.MppmConfigFileManager
	actualConfigInfo, err = configManager.GetGlobalConfig()
	assert.NotNil(t, err)

	configAsJson = configtest.TestMppmConfigInfosAsJson["valid version, invalid application name"]
	configtest.InitMockFileSystemDelegaterWithConfigFiles(configAsJson, configAsJson)
	configManager = config.MppmConfigFileManager
	actualConfigInfo, err = configManager.GetGlobalConfig()
	assert.NotNil(t, err)

	configAsJson = configtest.TestMppmConfigInfosAsJson["valid version, valid application name, invalid application version"]
	configtest.InitMockFileSystemDelegaterWithConfigFiles(configAsJson, configAsJson)
	configManager = config.MppmConfigFileManager
	actualConfigInfo, err = configManager.GetGlobalConfig()
	assert.NotNil(t, err)

}

// func TestGetDefaultMppmConfig(t *testing.T) {}

// func TestGetMppmGlobalConfigFilePath(t *testing.T) {}

// func TestSaveProjectConfig(t *testing.T) {}

// func TestSaveGlobalConfig(t *testing.T) {}

// func TestSaveDefaultProjectConfig(t *testing.T) {}
