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
			description:              "Test that the project config info is correctly returned from a valid config file.",
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
			description:   "Test that any error from os.Open() is correctly raised while opening the project config file.",
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
			description:              "Test that the global config info is correctly returned from a valid config file.",
			expectedConfigInfoAsJson: configtest.ConfigWithValidVersionAndApplicationNameAndApplicationVersion.ConfigAsJson,
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockFileSystemDelegaterBuilder(
					utiltest.NewMockFileSystemDelegaterBuilder().
						SetMockFileBuilders(
							configtest.ConfigWithValidVersionAndApplicationNameAndApplicationVersion.AsMockFileBuilder().
								SetFilePath("/home/testuser/.mppm.json"),
						),
				),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetMockFileBuilders(
					configtest.ConfigWithValidVersionAndApplicationNameAndApplicationVersion.AsMockFileBuilder().
						SetFilePath("/home/testuser/.mppm.json").
						SetWasClosed(true),
				),
		},

		&GetGlobalConfigTestCase{
			description:                     "Test that a default global config file is created if it does not already exist.",
			expectedConfigInfoAsJson:        configtest.GetDefaultMppmConfigAsJson(),
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder(),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetMockFileBuilders(
					utiltest.NewMockFileBuilder().
						SetFilePath("/home/testuser/.mppm.json").
						SetContentsFromBytes(configtest.GetDefaultMppmConfigAsJson()).
						SetWasClosed(true),
				),
		},

		&GetGlobalConfigTestCase{
			description:   "Test that an error is correctly raised when the global config file has an invalid mppm version.",
			expectedError: configtest.ConfigWithInvalidVersionAndNoApplications.ExpectedError,
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockFileSystemDelegaterBuilder(
					utiltest.NewMockFileSystemDelegaterBuilder().
						SetMockFileBuilders(
							configtest.ConfigWithInvalidVersionAndNoApplications.AsMockFileBuilder().
								SetFilePath("/home/testuser/.mppm.json"),
						),
				),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetMockFileBuilders(
					configtest.ConfigWithInvalidVersionAndNoApplications.AsMockFileBuilder().
						SetFilePath("/home/testuser/.mppm.json").
						SetWasClosed(true),
				),
		},

		&GetGlobalConfigTestCase{
			description:   "Test that an error is correctly raised when the global config file has an invalid application name.",
			expectedError: configtest.ConfigWithValidVersionAndInvalidApplicationName.ExpectedError,
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockFileSystemDelegaterBuilder(
					utiltest.NewMockFileSystemDelegaterBuilder().
						SetMockFileBuilders(
							configtest.ConfigWithValidVersionAndInvalidApplicationName.AsMockFileBuilder().
								SetFilePath("/home/testuser/.mppm.json"),
						),
				),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetMockFileBuilders(
					configtest.ConfigWithValidVersionAndInvalidApplicationName.AsMockFileBuilder().
						SetFilePath("/home/testuser/.mppm.json").
						SetWasClosed(true),
				),
		},

		&GetGlobalConfigTestCase{
			description:   "Test that an error is correctly raised when the global config file has an invalid application version.",
			expectedError: configtest.ConfigWithValidVersionAndApplicationNameAndInvalidApplicationVersion.ExpectedError,
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockFileSystemDelegaterBuilder(
					utiltest.NewMockFileSystemDelegaterBuilder().
						SetMockFileBuilders(
							configtest.ConfigWithValidVersionAndApplicationNameAndInvalidApplicationVersion.AsMockFileBuilder().
								SetFilePath("/home/testuser/.mppm.json"),
						),
				),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetMockFileBuilders(
					configtest.ConfigWithValidVersionAndApplicationNameAndInvalidApplicationVersion.AsMockFileBuilder().
						SetFilePath("/home/testuser/.mppm.json").
						SetWasClosed(true),
				),
		},

		&GetGlobalConfigTestCase{
			description:   "Test that any error from os.Open() is correctly raised while opening the global config file.",
			expectedError: errors.New("\nThere was a problem while opening the mppm config file.\nIf the file doesn't exist, try running 'mppm project init' first.\nThere was a problem opening the file.\n"),
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockFileSystemDelegaterBuilder(
					utiltest.NewMockFileSystemDelegaterBuilder().
						SetMockFileBuilders(
							configtest.ConfigWithValidVersionAndApplicationNameAndApplicationVersion.AsMockFileBuilder().
								SetFilePath("/home/testuser/.mppm.json"),
						).
						SetUseDefaultOpenFileError(true),
				),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetMockFileBuilders(
					configtest.ConfigWithValidVersionAndApplicationNameAndApplicationVersion.AsMockFileBuilder().
						SetFilePath("/home/testuser/.mppm.json"),
				),
		},

		&GetGlobalConfigTestCase{
			description:   "Test that any error from os.Create() is correctly raised while creating the global config file.",
			expectedError: utiltest.DefaultCreateFileError,
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockFileSystemDelegaterBuilder(
					utiltest.NewMockFileSystemDelegaterBuilder().
						SetUseDefaultCreateFileError(true),
				),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder(),
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
	description                              string
	expectedConfigInfoAsJson                 []byte
	expectedError                            error
	mockExecutionEnvironmentBuilder          *utiltest.MockExecutionEnvironmentBuilder
	expectedExecutionEnvironmentStateBuilder *utiltest.MockExecutionEnvironmentStateBuilder
}

func (testCase *GetGlobalConfigTestCase) Run(t *testing.T) {

	mockExecutionEnvironment := testCase.mockExecutionEnvironmentBuilder.BuildAndInit()

	actualConfigInfo, actualError := config.MppmConfigFileManager.GetGlobalConfig()
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
