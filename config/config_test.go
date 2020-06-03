package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stevengt/mppm/config"
	"github.com/stevengt/mppm/config/applications"

	"github.com/stevengt/mppm/config/configtest"
)

func TestGetFilePatternsConfigListFromProjectConfig(t *testing.T) {

	_, configAsJson := configtest.GetTestMppmConfigInfo()
	configtest.InitMockFileSystemDelegaterWithConfigFiles(configAsJson, configAsJson)

	expectedFilePatternsConfigList := []*applications.FilePatternsConfig{
		applications.Ableton10FilePatternsConfig,
	}

	actualFilePatternsConfigList := config.GetFilePatternsConfigListFromProjectConfig()
	assert.NotNil(t, actualFilePatternsConfigList)
	assert.Exactly(t, expectedFilePatternsConfigList, actualFilePatternsConfigList)

	configAsJson = []byte(`{"version":"1.0.0","applications":[]}`)
	configtest.InitMockFileSystemDelegaterWithConfigFiles(configAsJson, configAsJson)

	expectedFilePatternsConfigList = make([]*applications.FilePatternsConfig, 0)
	actualFilePatternsConfigList = config.GetFilePatternsConfigListFromProjectConfig()

	assert.NotNil(t, actualFilePatternsConfigList)
	assert.Exactly(t, expectedFilePatternsConfigList, actualFilePatternsConfigList)
	assert.Equal(t, 0, len(actualFilePatternsConfigList))

}

func TestGetAllFilePatternsConfigFromProjectConfig(t *testing.T) {

	_, configAsJson := configtest.GetTestMppmConfigInfo()
	configtest.InitMockFileSystemDelegaterWithConfigFiles(configAsJson, configAsJson)

	expectedFilePatternsConfig := &applications.FilePatternsConfig{
		Name:                     "",
		GitIgnorePatterns:        []string{"Backup/", "*.als", "*.alc", "*.adv", "*.adg"},
		GitLfsTrackPatterns:      []string{"*.ams", "*.amxd", "*.alp", "*.asd", "*.agr"},
		GzippedXmlFileExtensions: []string{"adv", "adg", "als", "alc"},
	}
	expectedFilePatternsConfig.SortAllLists()

	actualFilePatternsConfig := config.GetAllFilePatternsConfigFromProjectConfig()
	assert.NotNil(t, actualFilePatternsConfig)
	actualFilePatternsConfig.SortAllLists()
	assert.Exactly(t, expectedFilePatternsConfig, actualFilePatternsConfig)

	configAsJson = []byte(`{"version":"1.0.0","applications":[]}`)
	configtest.InitMockFileSystemDelegaterWithConfigFiles(configAsJson, configAsJson)

	expectedFilePatternsConfig = applications.NewFilePatternsConfig()
	actualFilePatternsConfig = config.GetAllFilePatternsConfigFromProjectConfig()
	assert.NotNil(t, actualFilePatternsConfig)
	assert.Exactly(t, expectedFilePatternsConfig, actualFilePatternsConfig)
	assert.Equal(t, 0, len(actualFilePatternsConfig.GitIgnorePatterns))
	assert.Equal(t, 0, len(actualFilePatternsConfig.GitLfsTrackPatterns))
	assert.Equal(t, 0, len(actualFilePatternsConfig.GzippedXmlFileExtensions))

}
