package utiltest

import (
	"github.com/stevengt/mppm/util"
)

type MockExecutionEnvironmentBuilder struct {
	MockShellCommandDelegater      *MockShellCommandDelegater
	MockFileSystemDelegaterBuilder *MockFileSystemDelegaterBuilder
	MockGitManagerCreatorBuilder   *MockGitManagerCreatorBuilder
}

func (builder *MockExecutionEnvironmentBuilder) Build() *MockExecutionEnvironment {

	return &MockExecutionEnvironment{
		MockExiter:                NewMockExiter(),
		MockPrinter:               NewMockPrinter(),
		MockShellCommandDelegater: builder.MockShellCommandDelegater,
		MockFileSystemDelegater:   GetMockFileSystemDelegaterFromBuilderOrNil(builder.MockFileSystemDelegaterBuilder),
		MockGitManagerCreator:     GetMockGitManagerCreatorFromBuilderOrNil(builder.MockGitManagerCreatorBuilder),
	}

}

// ------------------------------------------------------------------------------

type MockExecutionEnvironment struct {
	MockExiter                *MockExiter
	MockPrinter               *MockPrinter
	MockShellCommandDelegater *MockShellCommandDelegater
	MockFileSystemDelegater   *MockFileSystemDelegater
	MockGitManagerCreator     *MockGitManagerCreator
}

func (environment *MockExecutionEnvironment) Init() {

	if environment.MockExiter != nil {
		util.CurrentProcessExiter = environment.MockExiter
	}

	if environment.MockPrinter != nil {
		util.Logger = environment.MockPrinter
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
