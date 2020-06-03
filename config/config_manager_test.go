package config_test

import (
	"testing"

	"github.com/stevengt/mppm/config"
	"github.com/stevengt/mppm/config/configtest"
	"github.com/stevengt/mppm/util/utiltest"
	"github.com/stretchr/testify/assert"
)

func TestGetProjectConfig(t *testing.T) {

	expectedConfigInfo, configAsJson := configtest.GetTestMppmConfigInfo()
	mockCurrentProcessExiter := utiltest.InitializeAndReturnNewMockExiter()
	configtest.InitMockFileSystemDelegaterWithConfigFiles(configAsJson, configAsJson)
	configManager := config.MppmConfigFileManager
	actualConfigInfo := configManager.GetProjectConfig()
	assert.NotNil(t, actualConfigInfo)
	assert.Exactly(t, expectedConfigInfo, actualConfigInfo)
	assert.False(t, mockCurrentProcessExiter.WasExited)

	configAsJson = configtest.TestMppmConfigInfosAsJson["invalid version, no applications"]
	mockCurrentProcessExiter = utiltest.InitializeAndReturnNewMockExiter()
	configtest.InitMockFileSystemDelegaterWithConfigFiles(configAsJson, configAsJson)
	configManager = config.MppmConfigFileManager
	actualConfigInfo = configManager.GetProjectConfig()
	assert.True(t, mockCurrentProcessExiter.WasExited)

	configAsJson = configtest.TestMppmConfigInfosAsJson["valid version, invalid application name"]
	mockCurrentProcessExiter = utiltest.InitializeAndReturnNewMockExiter()
	configtest.InitMockFileSystemDelegaterWithConfigFiles(configAsJson, configAsJson)
	configManager = config.MppmConfigFileManager
	actualConfigInfo = configManager.GetProjectConfig()
	assert.True(t, mockCurrentProcessExiter.WasExited)

	configAsJson = configtest.TestMppmConfigInfosAsJson["valid version, valid application name, invalid application version"]
	mockCurrentProcessExiter = utiltest.InitializeAndReturnNewMockExiter()
	configtest.InitMockFileSystemDelegaterWithConfigFiles(configAsJson, configAsJson)
	configManager = config.MppmConfigFileManager
	actualConfigInfo = configManager.GetProjectConfig()
	assert.True(t, mockCurrentProcessExiter.WasExited)

}

func TestGetGlobalConfig(t *testing.T) {}

func TestGetDefaultMppmConfig(t *testing.T) {}

func TestGetMppmGlobalConfigFilePath(t *testing.T) {}

func TestSaveProjectConfig(t *testing.T) {}

func TestSaveGlobalConfig(t *testing.T) {}

func TestSaveDefaultProjectConfig(t *testing.T) {}
