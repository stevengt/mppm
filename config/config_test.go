package config_test

import (
	"errors"
	"testing"

	"github.com/stevengt/mppm/config/applications"
	"github.com/stevengt/mppm/util/utiltest"

	"github.com/stretchr/testify/assert"

	"github.com/stevengt/mppm/config"

	"github.com/stevengt/mppm/config/configtest"
)

func TestGetFilePatternsConfigListFromProjectConfig(t *testing.T) {

	testCases := []*GetFilePatternsConfigListFromProjectConfigTestCase{

		&GetFilePatternsConfigListFromProjectConfigTestCase{
			mockFileSystemDelegaterBuilder: nil,
			expectedErrorIfNotConfigError:  nil,
			mppmConfigInfoAndExpectedError: configtest.ConfigWithValidVersionAndApplicationNameAndApplicationVersion,
			expectedFilePatternsConfigList: []*applications.FilePatternsConfig{
				applications.AudioFilePatternsConfig,
				applications.Ableton10FilePatternsConfig,
			},
		},

		&GetFilePatternsConfigListFromProjectConfigTestCase{
			mockFileSystemDelegaterBuilder: nil,
			expectedErrorIfNotConfigError:  nil,
			mppmConfigInfoAndExpectedError: configtest.ConfigWithValidVersionAndNoApplications,
			expectedFilePatternsConfigList: []*applications.FilePatternsConfig{
				applications.AudioFilePatternsConfig,
			},
		},

		&GetFilePatternsConfigListFromProjectConfigTestCase{
			mockFileSystemDelegaterBuilder: nil,
			expectedErrorIfNotConfigError:  nil,
			mppmConfigInfoAndExpectedError: configtest.ConfigWithInvalidVersionAndNoApplications,
			expectedFilePatternsConfigList: nil,
		},

		&GetFilePatternsConfigListFromProjectConfigTestCase{
			mockFileSystemDelegaterBuilder: nil,
			expectedErrorIfNotConfigError:  nil,
			mppmConfigInfoAndExpectedError: configtest.ConfigWithValidVersionAndInvalidApplicationName,
			expectedFilePatternsConfigList: nil,
		},

		&GetFilePatternsConfigListFromProjectConfigTestCase{
			mockFileSystemDelegaterBuilder: nil,
			expectedErrorIfNotConfigError:  nil,
			mppmConfigInfoAndExpectedError: configtest.ConfigWithValidVersionAndApplicationNameAndInvalidApplicationVersion,
			expectedFilePatternsConfigList: nil,
		},

		&GetFilePatternsConfigListFromProjectConfigTestCase{
			mockFileSystemDelegaterBuilder: nil,
			expectedErrorIfNotConfigError: errors.New(`
There was a problem while opening the mppm config file.
If the file doesn't exist, try running 'mppm project init' first.
Unable to open file .mppm.json
`),
			mppmConfigInfoAndExpectedError: nil,
			expectedFilePatternsConfigList: nil,
		},

		&GetFilePatternsConfigListFromProjectConfigTestCase{
			mockFileSystemDelegaterBuilder: &utiltest.MockFileSystemDelegaterBuilder{
				UseDefaultOpenFileError: true,
			},
			expectedErrorIfNotConfigError: errors.New(`
There was a problem while opening the mppm config file.
If the file doesn't exist, try running 'mppm project init' first.
There was a problem opening the file.
`),
			mppmConfigInfoAndExpectedError: configtest.ConfigWithValidVersionAndApplicationNameAndApplicationVersion,
			expectedFilePatternsConfigList: nil,
		},
	}

	for _, testCase := range testCases {
		testCase.Run(t)
	}

}

func TestGetAllFilePatternsConfigFromProjectConfig(t *testing.T) {

	testCases := []*GetAllFilePatternsConfigFromProjectConfigTestCase{

		&GetAllFilePatternsConfigFromProjectConfigTestCase{
			mockFileSystemDelegaterBuilder: nil,
			expectedErrorIfNotConfigError:  nil,
			mppmConfigInfoAndExpectedError: configtest.ConfigWithValidVersionAndApplicationNameAndApplicationVersion,
			expectedFilePatternsConfig: &applications.FilePatternsConfig{
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
			},
		},

		&GetAllFilePatternsConfigFromProjectConfigTestCase{
			mockFileSystemDelegaterBuilder: nil,
			expectedErrorIfNotConfigError:  nil,
			mppmConfigInfoAndExpectedError: configtest.ConfigWithValidVersionAndNoApplications,
			expectedFilePatternsConfig: &applications.FilePatternsConfig{
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
			},
		},

		&GetAllFilePatternsConfigFromProjectConfigTestCase{
			mockFileSystemDelegaterBuilder: nil,
			expectedErrorIfNotConfigError:  nil,
			mppmConfigInfoAndExpectedError: configtest.ConfigWithInvalidVersionAndNoApplications,
			expectedFilePatternsConfig:     nil,
		},

		&GetAllFilePatternsConfigFromProjectConfigTestCase{
			mockFileSystemDelegaterBuilder: nil,
			expectedErrorIfNotConfigError:  nil,
			mppmConfigInfoAndExpectedError: configtest.ConfigWithValidVersionAndInvalidApplicationName,
			expectedFilePatternsConfig:     nil,
		},

		&GetAllFilePatternsConfigFromProjectConfigTestCase{
			mockFileSystemDelegaterBuilder: nil,
			expectedErrorIfNotConfigError:  nil,
			mppmConfigInfoAndExpectedError: configtest.ConfigWithValidVersionAndApplicationNameAndInvalidApplicationVersion,
			expectedFilePatternsConfig:     nil,
		},

		&GetAllFilePatternsConfigFromProjectConfigTestCase{
			mockFileSystemDelegaterBuilder: nil,
			expectedErrorIfNotConfigError: errors.New(`
There was a problem while opening the mppm config file.
If the file doesn't exist, try running 'mppm project init' first.
Unable to open file .mppm.json
`),
			mppmConfigInfoAndExpectedError: nil,
			expectedFilePatternsConfig:     nil,
		},

		&GetAllFilePatternsConfigFromProjectConfigTestCase{
			mockFileSystemDelegaterBuilder: &utiltest.MockFileSystemDelegaterBuilder{
				UseDefaultOpenFileError: true,
			},
			expectedErrorIfNotConfigError: errors.New(`
There was a problem while opening the mppm config file.
If the file doesn't exist, try running 'mppm project init' first.
There was a problem opening the file.
`),
			mppmConfigInfoAndExpectedError: configtest.ConfigWithValidVersionAndApplicationNameAndApplicationVersion,
			expectedFilePatternsConfig:     nil,
		},
	}

	for _, testCase := range testCases {
		testCase.Run(t)
	}

}

// ------------------------------------------------------------------------------

type GetFilePatternsConfigListFromProjectConfigTestCase struct {
	mockFileSystemDelegaterBuilder *utiltest.MockFileSystemDelegaterBuilder
	expectedErrorIfNotConfigError  error // The expected error, if not mppmConfigInfoAndExpectedError.ExpectedError
	mppmConfigInfoAndExpectedError *configtest.MppmConfigInfoAndExpectedError
	expectedFilePatternsConfigList []*applications.FilePatternsConfig
}

func (testCase *GetFilePatternsConfigListFromProjectConfigTestCase) Run(t *testing.T) {

	mockFileSystemDelegater := utiltest.GetMockFileSystemDelegaterFromBuilderOrNil(testCase.mockFileSystemDelegaterBuilder)

	projectConfigFile := configtest.ReturnMppmConfigInfoAsMockFileIfNotNilElseReturnNil(testCase.mppmConfigInfoAndExpectedError)
	configtest.InitMockFileSystemDelegaterWithConfigFiles(mockFileSystemDelegater, projectConfigFile, nil)

	expectedError := configtest.GetExpectedError(
		testCase.expectedErrorIfNotConfigError,
		testCase.mppmConfigInfoAndExpectedError,
	)

	expectedFilePatternsConfigList := testCase.expectedFilePatternsConfigList

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
	if expectedError == nil {
		assert.True(t, mockFileSystemDelegater.Files[".mppm.json"].WasClosed)
	}

}

// ------------------------------------------------------------------------------

type GetAllFilePatternsConfigFromProjectConfigTestCase struct {
	mockFileSystemDelegaterBuilder *utiltest.MockFileSystemDelegaterBuilder
	expectedErrorIfNotConfigError  error // The expected error, if not mppmConfigInfoAndExpectedError.ExpectedError
	mppmConfigInfoAndExpectedError *configtest.MppmConfigInfoAndExpectedError
	expectedFilePatternsConfig     *applications.FilePatternsConfig
}

func (testCase *GetAllFilePatternsConfigFromProjectConfigTestCase) Run(t *testing.T) {

	mockFileSystemDelegater := utiltest.GetMockFileSystemDelegaterFromBuilderOrNil(testCase.mockFileSystemDelegaterBuilder)

	projectConfigFile := configtest.ReturnMppmConfigInfoAsMockFileIfNotNilElseReturnNil(testCase.mppmConfigInfoAndExpectedError)
	configtest.InitMockFileSystemDelegaterWithConfigFiles(mockFileSystemDelegater, projectConfigFile, nil)

	expectedError := configtest.GetExpectedError(
		testCase.expectedErrorIfNotConfigError,
		testCase.mppmConfigInfoAndExpectedError,
	)

	expectedFilePatternsConfig := testCase.expectedFilePatternsConfig
	if expectedFilePatternsConfig != nil {
		expectedFilePatternsConfig.SortAllLists()
	}

	actualFilePatternsConfig, actualError := config.GetAllFilePatternsConfigFromProjectConfig()
	if actualFilePatternsConfig != nil {
		actualFilePatternsConfig.SortAllLists()
	}

	assert.Exactly(t, expectedFilePatternsConfig, actualFilePatternsConfig)
	assert.Exactly(t, expectedError, actualError)
	if expectedError == nil {
		assert.True(t, mockFileSystemDelegater.Files[".mppm.json"].WasClosed)
	}

}
