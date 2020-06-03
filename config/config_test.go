package config_test

import (
	"testing"

	"github.com/stevengt/mppm/util/utiltest"

	"github.com/stretchr/testify/assert"

	"github.com/stevengt/mppm/config"
	"github.com/stevengt/mppm/config/applications"

	"github.com/stevengt/mppm/config/configtest"
)

func TestGetFilePatternsConfigListFromProjectConfig(t *testing.T) {

	configAsJson := configtest.TestMppmConfigInfosAsJson["valid version, valid application name, valid application version"]
	mockCurrentProcessExiter := utiltest.InitializeAndReturnNewMockExiter()
	configtest.InitMockFileSystemDelegaterWithConfigFiles(configAsJson, configAsJson)
	expectedFilePatternsConfigList := []*applications.FilePatternsConfig{
		applications.AudioFilePatternsConfig,
		applications.Ableton10FilePatternsConfig,
	}
	actualFilePatternsConfigList := config.GetFilePatternsConfigListFromProjectConfig()
	assert.NotNil(t, actualFilePatternsConfigList)
	assert.Exactly(t, expectedFilePatternsConfigList, actualFilePatternsConfigList)
	assert.False(t, mockCurrentProcessExiter.WasExited)

	configAsJson = configtest.TestMppmConfigInfosAsJson["valid version, no applications"]
	mockCurrentProcessExiter = utiltest.InitializeAndReturnNewMockExiter()
	configtest.InitMockFileSystemDelegaterWithConfigFiles(configAsJson, configAsJson)
	expectedFilePatternsConfigList = []*applications.FilePatternsConfig{
		applications.AudioFilePatternsConfig,
	}
	actualFilePatternsConfigList = config.GetFilePatternsConfigListFromProjectConfig()
	assert.NotNil(t, actualFilePatternsConfigList)
	assert.Exactly(t, expectedFilePatternsConfigList, actualFilePatternsConfigList)
	assert.False(t, mockCurrentProcessExiter.WasExited)

	configAsJson = configtest.TestMppmConfigInfosAsJson["invalid version, no applications"]
	mockCurrentProcessExiter = utiltest.InitializeAndReturnNewMockExiter()
	configtest.InitMockFileSystemDelegaterWithConfigFiles(configAsJson, configAsJson)
	_ = config.GetFilePatternsConfigListFromProjectConfig()
	assert.True(t, mockCurrentProcessExiter.WasExited)

	configAsJson = configtest.TestMppmConfigInfosAsJson["valid version, invalid application name"]
	mockCurrentProcessExiter = utiltest.InitializeAndReturnNewMockExiter()
	configtest.InitMockFileSystemDelegaterWithConfigFiles(configAsJson, configAsJson)
	_ = config.GetFilePatternsConfigListFromProjectConfig()
	assert.True(t, mockCurrentProcessExiter.WasExited)

	configAsJson = configtest.TestMppmConfigInfosAsJson["valid version, valid application name, invalid application version"]
	mockCurrentProcessExiter = utiltest.InitializeAndReturnNewMockExiter()
	configtest.InitMockFileSystemDelegaterWithConfigFiles(configAsJson, configAsJson)
	_ = config.GetFilePatternsConfigListFromProjectConfig()
	assert.True(t, mockCurrentProcessExiter.WasExited)

}

func TestGetAllFilePatternsConfigFromProjectConfig(t *testing.T) {

	configAsJson := configtest.TestMppmConfigInfosAsJson["valid version, valid application name, valid application version"]
	mockCurrentProcessExiter := utiltest.InitializeAndReturnNewMockExiter()
	configtest.InitMockFileSystemDelegaterWithConfigFiles(configAsJson, configAsJson)
	expectedFilePatternsConfig := &applications.FilePatternsConfig{
		Name:              "",
		GitIgnorePatterns: []string{"Backup/", "*.als", "*.alc", "*.adv", "*.adg"},
		GitLfsTrackPatterns: []string{"*.flac", "*.iklax", "*.m4a", "*.alac", "*.au",
			"*.mpc", "*.ogg", "*.mogg", "*.tta", "*.wma", "*.aax", "*.act", "*.ivs",
			"*.aa", "*.dvf", "*.m4b", "*.nsf", "*.raw", "*.webm", "*.cda", "*.dct",
			"*.gsm", "*.dss", "*.msv", "*.nmf", "*.sln", "*.3gp", "*.aac", "*.voc",
			"*.wv", "*.m4p", "*.rm", "*.ape", "*.awb", "*.mmf", "*.oga", "*.opus",
			"*.rf64", "*.aiff", "*.amr", "*.vox", "*.wav", "*.8svx", "*.mp3", "*.ra",
			"*.ams", "*.amxd", "*.alp", "*.asd", "*.agr",
		},
		GzippedXmlFileExtensions: []string{"adv", "adg", "als", "alc"},
	}
	expectedFilePatternsConfig.SortAllLists()
	actualFilePatternsConfig := config.GetAllFilePatternsConfigFromProjectConfig()
	assert.NotNil(t, actualFilePatternsConfig)
	actualFilePatternsConfig.SortAllLists()
	assert.Exactly(t, expectedFilePatternsConfig, actualFilePatternsConfig)
	assert.False(t, mockCurrentProcessExiter.WasExited)

	configAsJson = configtest.TestMppmConfigInfosAsJson["valid version, no applications"]
	mockCurrentProcessExiter = utiltest.InitializeAndReturnNewMockExiter()
	configtest.InitMockFileSystemDelegaterWithConfigFiles(configAsJson, configAsJson)
	expectedFilePatternsConfig = &applications.FilePatternsConfig{
		Name:              "",
		GitIgnorePatterns: []string{},
		GitLfsTrackPatterns: []string{"*.flac", "*.iklax", "*.m4a", "*.alac", "*.au",
			"*.mpc", "*.ogg", "*.mogg", "*.tta", "*.wma", "*.aax", "*.act", "*.ivs",
			"*.aa", "*.dvf", "*.m4b", "*.nsf", "*.raw", "*.webm", "*.cda", "*.dct",
			"*.gsm", "*.dss", "*.msv", "*.nmf", "*.sln", "*.3gp", "*.aac", "*.voc",
			"*.wv", "*.m4p", "*.rm", "*.ape", "*.awb", "*.mmf", "*.oga", "*.opus",
			"*.rf64", "*.aiff", "*.amr", "*.vox", "*.wav", "*.8svx", "*.mp3", "*.ra",
		},
		GzippedXmlFileExtensions: []string{},
	}
	expectedFilePatternsConfig.SortAllLists()
	actualFilePatternsConfig = config.GetAllFilePatternsConfigFromProjectConfig()
	assert.NotNil(t, actualFilePatternsConfig)
	actualFilePatternsConfig.SortAllLists()
	assert.Exactly(t, expectedFilePatternsConfig, actualFilePatternsConfig)
	assert.False(t, mockCurrentProcessExiter.WasExited)

	configAsJson = configtest.TestMppmConfigInfosAsJson["invalid version, no applications"]
	mockCurrentProcessExiter = utiltest.InitializeAndReturnNewMockExiter()
	configtest.InitMockFileSystemDelegaterWithConfigFiles(configAsJson, configAsJson)
	_ = config.GetAllFilePatternsConfigFromProjectConfig()
	assert.True(t, mockCurrentProcessExiter.WasExited)

	configAsJson = configtest.TestMppmConfigInfosAsJson["valid version, invalid application name"]
	mockCurrentProcessExiter = utiltest.InitializeAndReturnNewMockExiter()
	configtest.InitMockFileSystemDelegaterWithConfigFiles(configAsJson, configAsJson)
	_ = config.GetAllFilePatternsConfigFromProjectConfig()
	assert.True(t, mockCurrentProcessExiter.WasExited)

	configAsJson = configtest.TestMppmConfigInfosAsJson["valid version, valid application name, invalid application version"]
	mockCurrentProcessExiter = utiltest.InitializeAndReturnNewMockExiter()
	configtest.InitMockFileSystemDelegaterWithConfigFiles(configAsJson, configAsJson)
	_ = config.GetAllFilePatternsConfigFromProjectConfig()
	assert.True(t, mockCurrentProcessExiter.WasExited)

}
