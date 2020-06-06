package utiltest

import (
	"github.com/stevengt/mppm/util"
)

func GetAndInitMockExecutionEnvironmentFromBuilderOrNil(mockExecutionEnvironmentBuilder *MockExecutionEnvironmentBuilder) *MockExecutionEnvironment {
	var mockExecutionEnvironment *MockExecutionEnvironment
	if mockExecutionEnvironmentBuilder != nil {
		mockExecutionEnvironment = mockExecutionEnvironmentBuilder.Build()
	} else {
		mockExecutionEnvironment = NewDefaultMockExecutionEnvironmentBuilder().Build()
	}
	mockExecutionEnvironment.Init()
	return mockExecutionEnvironment
}

// ------------------------------------------------------------------------------

type MockExecutionEnvironmentBuilder struct {
	MockShellCommandDelegater      *MockShellCommandDelegater
	MockFileSystemDelegaterBuilder *MockFileSystemDelegaterBuilder
	MockGitManagerCreatorBuilder   *MockGitManagerCreatorBuilder
}

func NewDefaultMockExecutionEnvironmentBuilder() *MockExecutionEnvironmentBuilder {
	return &MockExecutionEnvironmentBuilder{}
}

func (builder *MockExecutionEnvironmentBuilder) Build() *MockExecutionEnvironment {

	return &MockExecutionEnvironment{
		MockExiter:                NewMockExiter(),
		MockWritePrinter:          NewMockWritePrinter(),
		MockShellCommandDelegater: builder.MockShellCommandDelegater,
		MockFileSystemDelegater:   GetMockFileSystemDelegaterFromBuilderOrNil(builder.MockFileSystemDelegaterBuilder),
		MockGitManagerCreator:     GetMockGitManagerCreatorFromBuilderOrNil(builder.MockGitManagerCreatorBuilder),
	}

}

// ------------------------------------------------------------------------------

type MockExecutionEnvironment struct {
	MockExiter                *MockExiter
	MockWritePrinter          *MockWritePrinter
	MockShellCommandDelegater *MockShellCommandDelegater
	MockFileSystemDelegater   *MockFileSystemDelegater
	MockGitManagerCreator     *MockGitManagerCreator
}

func (environment *MockExecutionEnvironment) Init() {

	if environment.MockExiter != nil {
		util.CurrentProcessExiter = environment.MockExiter
	}

	if environment.MockWritePrinter != nil {
		util.Logger = environment.MockWritePrinter
	}

	if environment.MockShellCommandDelegater != nil {
		util.ShellProxy = environment.MockShellCommandDelegater
	}

	if environment.MockFileSystemDelegater != nil {
		util.FileSystemProxy = environment.MockFileSystemDelegater
	}

	if environment.MockGitManagerCreator != nil {
		util.GitManagerFactory = environment.MockGitManagerCreator
	}

}
