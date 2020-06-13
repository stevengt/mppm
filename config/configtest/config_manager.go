package configtest

import (
	"errors"

	"github.com/stevengt/mppm/config"
)

var DefaultGetProjectConfigError error = errors.New("There was a problem getting the project config.")

var DefaultGetGlobalConfigError error = errors.New("There was a problem getting the global config.")

var DefaultGetProjectAndGlobalConfigsError error = errors.New("There was a problem getting the project and global configs.")

var DefaultGetMppmGlobalConfigFilePathError error = errors.New("There was a problem getting the global config's file path.")

var DefaultSaveProjectConfigError error = errors.New("There was a problem saving the project config.")

var DefaultSaveGlobalConfigError error = errors.New("There was a problem saving the global config.")

var DefaultSaveDefaultProjectConfigError error = errors.New("There was a problem saving the default project config.")

// ------------------------------------------------------------------------------

type MockMppmConfigManagerBuilder struct {
	ProjectConfig                              *config.MppmConfigInfo
	GlobalConfig                               *config.MppmConfigInfo
	UseDefaultGetProjectConfigError            bool
	UseDefaultGetGlobalConfigError             bool
	UseDefaultGetProjectAndGlobalConfigsError  bool
	UseDefaultGetMppmGlobalConfigFilePathError bool
	UseDefaultSaveProjectConfigError           bool
	UseDefaultSaveGlobalConfigError            bool
	UseDefaultSaveDefaultProjectConfigError    bool
}

func NewMockMppmConfigManagerBuilder() *MockMppmConfigManagerBuilder {
	return &MockMppmConfigManagerBuilder{}
}

func (builder *MockMppmConfigManagerBuilder) SetProjectConfig(projectConfig *config.MppmConfigInfo) *MockMppmConfigManagerBuilder {
	builder.ProjectConfig = projectConfig
	return builder
}

func (builder *MockMppmConfigManagerBuilder) SetProjectConfigFromJson(projectConfigAsJson []byte) *MockMppmConfigManagerBuilder {
	builder.ProjectConfig, _ = config.NewMppmConfigInfoFromJson(projectConfigAsJson)
	return builder
}

func (builder *MockMppmConfigManagerBuilder) SetGlobalConfig(globalConfig *config.MppmConfigInfo) *MockMppmConfigManagerBuilder {
	builder.GlobalConfig = globalConfig
	return builder
}

func (builder *MockMppmConfigManagerBuilder) SetGlobalConfigFromJson(globalConfigAsJson []byte) *MockMppmConfigManagerBuilder {
	builder.GlobalConfig, _ = config.NewMppmConfigInfoFromJson(globalConfigAsJson)
	return builder
}

func (builder *MockMppmConfigManagerBuilder) SetUseDefaultGetProjectConfigError(useDefaultGetProjectConfigError bool) *MockMppmConfigManagerBuilder {
	builder.UseDefaultGetProjectConfigError = useDefaultGetProjectConfigError
	return builder
}

func (builder *MockMppmConfigManagerBuilder) SetUseDefaultGetGlobalConfigError(useDefaultGetGlobalConfigError bool) *MockMppmConfigManagerBuilder {
	builder.UseDefaultGetGlobalConfigError = useDefaultGetGlobalConfigError
	return builder
}

func (builder *MockMppmConfigManagerBuilder) SetUseDefaultGetProjectAndGlobalConfigsError(useDefaultGetProjectAndGlobalConfigsError bool) *MockMppmConfigManagerBuilder {
	builder.UseDefaultGetProjectAndGlobalConfigsError = useDefaultGetProjectAndGlobalConfigsError
	return builder
}

func (builder *MockMppmConfigManagerBuilder) SetUseDefaultGetMppmGlobalConfigFilePathError(useDefaultGetMppmGlobalConfigFilePathError bool) *MockMppmConfigManagerBuilder {
	builder.UseDefaultGetMppmGlobalConfigFilePathError = useDefaultGetMppmGlobalConfigFilePathError
	return builder
}

func (builder *MockMppmConfigManagerBuilder) SetUseDefaultSaveProjectConfigError(useDefaultSaveProjectConfigError bool) *MockMppmConfigManagerBuilder {
	builder.UseDefaultSaveProjectConfigError = useDefaultSaveProjectConfigError
	return builder
}

func (builder *MockMppmConfigManagerBuilder) SetUseDefaultSaveGlobalConfigError(useDefaultSaveGlobalConfigError bool) *MockMppmConfigManagerBuilder {
	builder.UseDefaultSaveGlobalConfigError = useDefaultSaveGlobalConfigError
	return builder
}

func (builder *MockMppmConfigManagerBuilder) SetUseDefaultSaveDefaultProjectConfigError(useDefaultSaveDefaultProjectConfigError bool) *MockMppmConfigManagerBuilder {
	builder.UseDefaultSaveDefaultProjectConfigError = useDefaultSaveDefaultProjectConfigError
	return builder
}

func (builder *MockMppmConfigManagerBuilder) Build() *MockMppmConfigManager {

	mockMppmConfigManager := &MockMppmConfigManager{
		ProjectConfig: builder.ProjectConfig,
		GlobalConfig:  builder.GlobalConfig,
	}

	if builder.UseDefaultGetProjectConfigError {
		mockMppmConfigManager.GetProjectConfigError = DefaultGetProjectConfigError
	}

	if builder.UseDefaultGetGlobalConfigError {
		mockMppmConfigManager.GetGlobalConfigError = DefaultGetGlobalConfigError
	}

	if builder.UseDefaultGetProjectAndGlobalConfigsError {
		mockMppmConfigManager.GetProjectAndGlobalConfigsError = DefaultGetProjectAndGlobalConfigsError
	}

	if builder.UseDefaultGetMppmGlobalConfigFilePathError {
		mockMppmConfigManager.GetMppmGlobalConfigFilePathError = DefaultGetMppmGlobalConfigFilePathError
	}

	if builder.UseDefaultSaveProjectConfigError {
		mockMppmConfigManager.SaveProjectConfigError = DefaultSaveProjectConfigError
	}

	if builder.UseDefaultSaveGlobalConfigError {
		mockMppmConfigManager.SaveGlobalConfigError = DefaultSaveGlobalConfigError
	}

	if builder.UseDefaultSaveDefaultProjectConfigError {
		mockMppmConfigManager.SaveDefaultProjectConfigError = DefaultSaveDefaultProjectConfigError
	}

	return mockMppmConfigManager

}

// Builds the MockMppmConfigManager, initializes it with
// the MockMppmConfigManager.Init() method, and returns the new MockMppmConfigManager.
func (builder *MockMppmConfigManagerBuilder) BuildAndInit() *MockMppmConfigManager {
	mockMppmConfigManager := builder.Build()
	mockMppmConfigManager.Init()
	return mockMppmConfigManager
}

// ------------------------------------------------------------------------------

type MockMppmConfigManager struct {
	ProjectConfig                    *config.MppmConfigInfo
	GlobalConfig                     *config.MppmConfigInfo
	GetProjectConfigError            error
	GetGlobalConfigError             error
	GetProjectAndGlobalConfigsError  error
	GetMppmGlobalConfigFilePathError error
	SaveProjectConfigError           error
	SaveGlobalConfigError            error
	SaveDefaultProjectConfigError    error
}

func (mockMppmConfigManager *MockMppmConfigManager) Init() {
	config.MppmConfigFileManager = mockMppmConfigManager
}

func (mockMppmConfigManager *MockMppmConfigManager) GetProjectConfig() (projectConfig *config.MppmConfigInfo, err error) {
	err = mockMppmConfigManager.GetProjectConfigError
	if err == nil {
		projectConfig = mockMppmConfigManager.ProjectConfig
	}
	return
}

func (mockMppmConfigManager *MockMppmConfigManager) GetGlobalConfig() (globalConfig *config.MppmConfigInfo, err error) {
	err = mockMppmConfigManager.GetGlobalConfigError
	if err == nil {
		globalConfig = mockMppmConfigManager.GlobalConfig
	}
	return
}

func (mockMppmConfigManager *MockMppmConfigManager) GetProjectAndGlobalConfigs() (projectConfig *config.MppmConfigInfo, globalConfig *config.MppmConfigInfo, err error) {
	err = mockMppmConfigManager.GetProjectAndGlobalConfigsError
	if err == nil {
		projectConfig = mockMppmConfigManager.ProjectConfig
		globalConfig = mockMppmConfigManager.GlobalConfig
	}
	return
}

func (mockMppmConfigManager *MockMppmConfigManager) GetDefaultMppmConfig() (mppmConfig *config.MppmConfigInfo) {
	return GetDefaultTestMppmConfigInfo()
}

func (mockMppmConfigManager *MockMppmConfigManager) GetMppmGlobalConfigFilePath() (filePath string, err error) {
	err = mockMppmConfigManager.GetMppmGlobalConfigFilePathError
	if err == nil {
		filePath = "/home/testuser/.mppm.json"
	}
	return
}

func (mockMppmConfigManager *MockMppmConfigManager) SaveProjectConfig() (err error) {
	return mockMppmConfigManager.SaveProjectConfigError
}

func (mockMppmConfigManager *MockMppmConfigManager) SaveGlobalConfig() (err error) {
	return mockMppmConfigManager.SaveGlobalConfigError
}

func (mockMppmConfigManager *MockMppmConfigManager) SaveDefaultProjectConfig() (err error) {
	return mockMppmConfigManager.SaveDefaultProjectConfigError
}
