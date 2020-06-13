package config_test

import (
	"errors"
	"testing"

	"github.com/stevengt/mppm/util/utiltest"

	"github.com/stevengt/mppm/config/applications"

	"github.com/stevengt/mppm/config"
	"github.com/stevengt/mppm/config/configtest"
	"github.com/stretchr/testify/assert"
)

func TestGetProjectConfig(t *testing.T) {

	testCases := []*GetProjectConfigTestCase{

		&GetProjectConfigTestCase{
			description:              "Test that config info is correctly returned from valid project config file.",
			expectedConfigInfoAsJson: configtest.ConfigWithValidVersionAndApplicationNameAndApplicationVersion.ConfigAsJson,
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockFileSystemDelegaterBuilder(
					utiltest.NewMockFileSystemDelegaterBuilder().
						SetMockFileBuilders(
							configtest.ConfigWithValidVersionAndApplicationNameAndApplicationVersion.AsMockFileBuilder().
								SetFilePath(config.MppmConfigFileName),
						),
				),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetMockFileBuilders(
					configtest.ConfigWithValidVersionAndApplicationNameAndApplicationVersion.AsMockFileBuilder().
						SetFilePath(config.MppmConfigFileName).
						SetWasClosed(true),
				),
		},

		&GetProjectConfigTestCase{
			description:   "Test that an error is correctly raised when the project config file has an invalid mppm version.",
			expectedError: configtest.ConfigWithInvalidVersionAndNoApplications.ExpectedError,
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockFileSystemDelegaterBuilder(
					utiltest.NewMockFileSystemDelegaterBuilder().
						SetMockFileBuilders(
							configtest.ConfigWithInvalidVersionAndNoApplications.AsMockFileBuilder().
								SetFilePath(config.MppmConfigFileName),
						),
				),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetMockFileBuilders(
					configtest.ConfigWithInvalidVersionAndNoApplications.AsMockFileBuilder().
						SetFilePath(config.MppmConfigFileName).
						SetWasClosed(true),
				),
		},

		&GetProjectConfigTestCase{
			description:   "Test that an error is correctly raised when the project config file has an invalid application name.",
			expectedError: configtest.ConfigWithValidVersionAndInvalidApplicationName.ExpectedError,
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockFileSystemDelegaterBuilder(
					utiltest.NewMockFileSystemDelegaterBuilder().
						SetMockFileBuilders(
							configtest.ConfigWithValidVersionAndInvalidApplicationName.AsMockFileBuilder().
								SetFilePath(config.MppmConfigFileName),
						),
				),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetMockFileBuilders(
					configtest.ConfigWithValidVersionAndInvalidApplicationName.AsMockFileBuilder().
						SetFilePath(config.MppmConfigFileName).
						SetWasClosed(true),
				),
		},

		&GetProjectConfigTestCase{
			description:   "Test that an error is correctly raised when the project config file has an invalid application version.",
			expectedError: configtest.ConfigWithValidVersionAndApplicationNameAndInvalidApplicationVersion.ExpectedError,
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockFileSystemDelegaterBuilder(
					utiltest.NewMockFileSystemDelegaterBuilder().
						SetMockFileBuilders(
							configtest.ConfigWithValidVersionAndApplicationNameAndInvalidApplicationVersion.AsMockFileBuilder().
								SetFilePath(config.MppmConfigFileName),
						),
				),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetMockFileBuilders(
					configtest.ConfigWithValidVersionAndApplicationNameAndInvalidApplicationVersion.AsMockFileBuilder().
						SetFilePath(config.MppmConfigFileName).
						SetWasClosed(true),
				),
		},

		&GetProjectConfigTestCase{
			description:                              "Test that an error is correctly raised when the project config file does not exist.",
			expectedError:                            errors.New("\nThere was a problem while opening the mppm config file.\nIf the file doesn't exist, try running 'mppm project init' first.\nUnable to open file .mppm.json\n"),
			mockExecutionEnvironmentBuilder:          utiltest.NewMockExecutionEnvironmentBuilder(),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder(),
		},

		&GetProjectConfigTestCase{
			description:   "Test that any error from os.Open() is correctly raised.",
			expectedError: errors.New("\nThere was a problem while opening the mppm config file.\nIf the file doesn't exist, try running 'mppm project init' first.\nThere was a problem opening the file.\n"),
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockFileSystemDelegaterBuilder(
					utiltest.NewMockFileSystemDelegaterBuilder().
						SetUseDefaultOpenFileError(true),
				),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder(),
		},
	}

	for _, testCase := range testCases {
		testCase.Run(t)
	}

}

func TestGetGlobalConfig(t *testing.T) {

	testCases := []*GetGlobalConfigTestCase{

		&GetGlobalConfigTestCase{
			mockFileSystemDelegaterBuilder: nil,
			expectedErrorIfNotConfigError:  nil,
			mppmConfigInfoAndExpectedError: configtest.ConfigWithValidVersionAndApplicationNameAndApplicationVersion,
		},

		&GetGlobalConfigTestCase{
			mockFileSystemDelegaterBuilder: nil,
			expectedErrorIfNotConfigError:  nil,
			mppmConfigInfoAndExpectedError: configtest.ConfigWithValidVersionAndNoApplications,
		},

		&GetGlobalConfigTestCase{
			mockFileSystemDelegaterBuilder: nil,
			expectedErrorIfNotConfigError:  nil,
			mppmConfigInfoAndExpectedError: configtest.ConfigWithInvalidVersionAndNoApplications,
		},

		&GetGlobalConfigTestCase{
			mockFileSystemDelegaterBuilder: nil,
			expectedErrorIfNotConfigError:  nil,
			mppmConfigInfoAndExpectedError: configtest.ConfigWithValidVersionAndInvalidApplicationName,
		},

		&GetGlobalConfigTestCase{
			mockFileSystemDelegaterBuilder: nil,
			expectedErrorIfNotConfigError:  nil,
			mppmConfigInfoAndExpectedError: configtest.ConfigWithValidVersionAndApplicationNameAndInvalidApplicationVersion,
		},

		&GetGlobalConfigTestCase{
			mockFileSystemDelegaterBuilder: nil,
			expectedErrorIfNotConfigError:  nil,
			mppmConfigInfoAndExpectedError: nil,
		},

		&GetGlobalConfigTestCase{
			mockFileSystemDelegaterBuilder: &utiltest.MockFileSystemDelegaterBuilder{
				UseDefaultOpenFileError: true,
			},
			expectedErrorIfNotConfigError:  errors.New("\nThere was a problem while opening the mppm config file.\nIf the file doesn't exist, try running 'mppm project init' first.\nThere was a problem opening the file.\n"),
			mppmConfigInfoAndExpectedError: configtest.ConfigWithValidVersionAndApplicationNameAndApplicationVersion,
		},

		&GetGlobalConfigTestCase{
			mockFileSystemDelegaterBuilder: &utiltest.MockFileSystemDelegaterBuilder{
				UseDefaultCreateFileError: true,
			},
			expectedErrorIfNotConfigError:  errors.New("There was a problem creating the file."),
			mppmConfigInfoAndExpectedError: nil,
		},
	}

	for _, testCase := range testCases {
		testCase.Run(t)
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

	mockFileSystemDelegater := utiltest.NewMockFileSystemDelegater()
	configtest.InitMockFileSystemDelegaterWithConfigFiles(mockFileSystemDelegater, nil, nil)
	configManager := config.MppmConfigFileManager
	expectedFilePath := "/home/testuser/.mppm.json"
	expectedError = nil
	actualFilePath, actualError := configManager.GetMppmGlobalConfigFilePath()
	assert.Exactly(t, expectedFilePath, actualFilePath)
	assert.Exactly(t, expectedError, actualError)

	mockFileSystemDelegater = (&utiltest.MockFileSystemDelegaterBuilder{
		UseDefaultUserHomeDirError: true,
	}).Build()
	configtest.InitMockFileSystemDelegaterWithConfigFiles(mockFileSystemDelegater, nil, nil)
	configManager = config.MppmConfigFileManager
	expectedFilePath = ""
	expectedError = utiltest.DefaultUserHomeDirError
	actualFilePath, actualError = configManager.GetMppmGlobalConfigFilePath()
	assert.Exactly(t, expectedFilePath, actualFilePath)
	assert.Exactly(t, expectedError, actualError)

}

func TestSaveProjectConfig(t *testing.T) {

	mockFileSystemDelegater := configtest.InitAndReturnMockFileSystemDelegaterWithDefaultConfigFiles()
	configManager := config.MppmConfigFileManager
	expectedError := errors.New("Unable to save uninitialized project config.")
	actualError := configManager.SaveProjectConfig()
	assert.Exactly(t, expectedError, actualError)

	mockFileSystemDelegater = configtest.InitAndReturnMockFileSystemDelegaterWithDefaultConfigFiles()
	configManager = config.MppmConfigFileManager
	projectConfig, actualError := configManager.GetProjectConfig()
	assert.Nil(t, actualError)
	assert.NotNil(t, projectConfig)
	assert.True(t, mockFileSystemDelegater.Files[".mppm.json"].WasClosed)
	projectConfig.Version = "1.9999.9999"
	expectedError = nil
	actualError = configManager.SaveProjectConfig()
	assert.Exactly(t, expectedError, actualError)
	expectedConfigInfo := projectConfig
	actualConfigInfo, actualError := config.NewMppmConfigInfoFromJsonReader(mockFileSystemDelegater.Files[".mppm.json"])
	assert.Nil(t, actualError)
	assert.NotNil(t, actualConfigInfo)
	assert.Equal(t, "1.9999.9999", actualConfigInfo.Version)
	assert.Exactly(t, expectedConfigInfo, actualConfigInfo)

}

func TestSaveGlobalConfig(t *testing.T) {

	mockFileSystemDelegater := configtest.InitAndReturnMockFileSystemDelegaterWithDefaultConfigFiles()
	configManager := config.MppmConfigFileManager
	expectedError := errors.New("Unable to save uninitialized global config.")
	actualError := configManager.SaveGlobalConfig()
	assert.Exactly(t, expectedError, actualError)

	mockFileSystemDelegater = configtest.InitAndReturnMockFileSystemDelegaterWithDefaultConfigFiles()
	configManager = config.MppmConfigFileManager
	actualConfigInfo, actualError := configManager.GetGlobalConfig()
	assert.Nil(t, actualError)
	assert.NotNil(t, actualConfigInfo)
	assert.True(t, mockFileSystemDelegater.Files["/home/testuser/.mppm.json"].WasClosed)
	actualConfigInfo.Version = "1.9999.9999"
	expectedError = nil
	actualError = configManager.SaveGlobalConfig()
	assert.Exactly(t, expectedError, actualError)
	expectedConfigInfo := actualConfigInfo
	expectedError = nil
	actualConfigInfo, actualError = config.NewMppmConfigInfoFromJsonReader(mockFileSystemDelegater.Files["/home/testuser/.mppm.json"])
	assert.Nil(t, actualError)
	assert.NotNil(t, actualConfigInfo)
	assert.Equal(t, "1.9999.9999", actualConfigInfo.Version)
	assert.Exactly(t, expectedConfigInfo, actualConfigInfo)

}

func TestSaveDefaultProjectConfig(t *testing.T) {

	mockFileSystemDelegater := configtest.InitAndReturnMockFileSystemDelegaterWithNoConfigFiles()
	configManager := config.MppmConfigFileManager

	_, actualError := configManager.GetProjectConfig()
	assert.NotNil(t, actualError)
	actualError = configManager.SaveDefaultProjectConfig()
	assert.Nil(t, actualError)

	config.MppmConfigFileManager = config.NewMppmConfigFileManager()
	configManager = config.MppmConfigFileManager

	expectedProjectConfig := configManager.GetDefaultMppmConfig()
	actualProjectConfig, actualError := configManager.GetProjectConfig()
	assert.Exactly(t, expectedProjectConfig, actualProjectConfig)
	assert.Nil(t, actualError)

	mockFileSystemDelegater = (&utiltest.MockFileSystemDelegaterBuilder{
		UseDefaultCreateFileError: true,
	}).Build()
	configtest.InitMockFileSystemDelegaterWithConfigFiles(mockFileSystemDelegater, nil, nil)
	configManager = config.MppmConfigFileManager
	expectedError := utiltest.DefaultCreateFileError
	actualError = configManager.SaveDefaultProjectConfig()
	assert.Exactly(t, expectedError, actualError)

}

// ------------------------------------------------------------------------------

type GetProjectConfigTestCase struct {
	description                              string
	expectedConfigInfoAsJson                 []byte
	expectedError                            error
	mockExecutionEnvironmentBuilder          *utiltest.MockExecutionEnvironmentBuilder
	expectedExecutionEnvironmentStateBuilder *utiltest.MockExecutionEnvironmentStateBuilder
}

func (testCase *GetProjectConfigTestCase) Run(t *testing.T) {

	mockExecutionEnvironment := testCase.mockExecutionEnvironmentBuilder.BuildAndInit()

	actualConfigInfo, actualError := config.MppmConfigFileManager.GetProjectConfig()
	assert.Exactlyf(t, testCase.expectedError, actualError, testCase.description)

	var actualConfigInfoAsJson []byte
	if actualConfigInfo != nil {
		actualConfigInfoAsJson, actualError = actualConfigInfo.AsJson()
		assert.Nilf(t, actualError, testCase.description)
	}
	assert.Exactlyf(t, testCase.expectedConfigInfoAsJson, actualConfigInfoAsJson, testCase.description)

	expectedExecutionEnvironmentState := testCase.expectedExecutionEnvironmentStateBuilder.Build()
	mockExecutionEnvironment.GetCurrentState().AssertEquals(t, expectedExecutionEnvironmentState, testCase.description)

}

// ------------------------------------------------------------------------------

type GetGlobalConfigTestCase struct {
	mockFileSystemDelegaterBuilder *utiltest.MockFileSystemDelegaterBuilder
	expectedErrorIfNotConfigError  error // The expected error, if not mppmConfigInfoAndExpectedError.ExpectedError
	mppmConfigInfoAndExpectedError *configtest.MppmConfigInfoAndExpectedError
}

func (testCase *GetGlobalConfigTestCase) Run(t *testing.T) {

	mockFileSystemDelegater := utiltest.GetMockFileSystemDelegaterFromBuilderOrNil(testCase.mockFileSystemDelegaterBuilder)

	globalConfigFile := configtest.ReturnMppmConfigInfoAsMockFileIfNotNilElseReturnNil(testCase.mppmConfigInfoAndExpectedError)
	configtest.InitMockFileSystemDelegaterWithConfigFiles(mockFileSystemDelegater, nil, globalConfigFile)

	var actualError error
	expectedError := configtest.GetExpectedError(
		testCase.expectedErrorIfNotConfigError,
		testCase.mppmConfigInfoAndExpectedError,
	)

	configManager := config.MppmConfigFileManager

	expectedConfigInfoAsJson := configtest.ReturnMppmConfigInfoAsJsonIfNotNilAndErrorIsNilElseReturnNil(
		testCase.mppmConfigInfoAndExpectedError,
		expectedError,
	)
	if expectedConfigInfoAsJson == nil {
		expectedConfigInfoAsJson, actualError = configManager.GetDefaultMppmConfig().AsJson()
		assert.Nil(t, actualError)
	}

	var actualConfigInfoAsJson []byte
	actualConfigInfo, actualError := configManager.GetGlobalConfig()
	assert.Exactly(t, expectedError, actualError)

	if actualConfigInfo != nil {
		actualConfigInfoAsJson, actualError = actualConfigInfo.AsJson()
		assert.Nil(t, actualError)
	}

	if expectedError == nil {
		assert.Exactly(t, expectedConfigInfoAsJson, actualConfigInfoAsJson)
		assert.True(t, mockFileSystemDelegater.DoesFileExist("/home/testuser/.mppm.json"))
		assert.True(t, mockFileSystemDelegater.Files["/home/testuser/.mppm.json"].WasClosed)
	}

}
