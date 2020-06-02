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

	expectedFilePatternsList := []*applications.FilePatternsConfig{
		applications.Ableton10FilePatternsConfig,
	}

	actualFilePatternsList := config.GetFilePatternsConfigListFromProjectConfig()
	assert.NotNil(t, actualFilePatternsList)
	assert.Exactly(t, expectedFilePatternsList, actualFilePatternsList)

	configAsJson = []byte(`{"version":"1.0.0","applications":[]}`)
	configtest.InitMockFileSystemDelegaterWithConfigFiles(configAsJson, configAsJson)

	expectedFilePatternsList = make([]*applications.FilePatternsConfig, 0)
	actualFilePatternsList = config.GetFilePatternsConfigListFromProjectConfig()

	assert.NotNil(t, actualFilePatternsList)
	assert.Exactly(t, expectedFilePatternsList, actualFilePatternsList)
	assert.Equal(t, 0, len(actualFilePatternsList))

}

func TestGetAllFilePatternsConfigFromProjectConfig(t *testing.T) {

}
