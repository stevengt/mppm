package utiltest

import (
	"testing"

	"github.com/stevengt/mppm/util"
	"github.com/stretchr/testify/assert"
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

func (environment *MockExecutionEnvironment) GetCurrentState() *MockExecutionEnvironmentState {

	gitManagerInputHistoriesIndexedByRepoPath := make(map[string][][]string)
	for repoPath, gitManager := range environment.MockGitManagerCreator.MockGitManagersIndexedByRepoPath {
		gitManagerInputHistoriesIndexedByRepoPath[repoPath] = gitManager.InputHistory
	}

	return &MockExecutionEnvironmentState{
		ExiterWasExited:          environment.MockExiter.WasExited,
		ExiterError:              environment.MockExiter.Error,
		FileSystemDelegaterFiles: environment.MockFileSystemDelegater.Files,
		GitManagerInputHistoriesIndexedByRepoPath: gitManagerInputHistoriesIndexedByRepoPath,
		WritePrinterOutputContents:                environment.MockWritePrinter.OutputContents,
		ShellCommandDelegaterInputHistory:         environment.MockShellCommandDelegater.InputHistory,
		ShellCommandDelegaterOutputHistory:        environment.MockShellCommandDelegater.OutputHistory,
	}

}

// ------------------------------------------------------------------------------

// Describes the current or expected state of a MockExecutionEnvironment.
type MockExecutionEnvironmentState struct {
	ExiterWasExited                           bool
	ExiterError                               error
	FileSystemDelegaterFiles                  map[string]*MockFile // Map of file names to mocked file instances.
	GitManagerInputHistoriesIndexedByRepoPath map[string][][]string
	WritePrinterOutputContents                []byte
	ShellCommandDelegaterInputHistory         []string
	ShellCommandDelegaterOutputHistory        []*MockShellCommandOutput
}

func (environmentState *MockExecutionEnvironmentState) AssertEquals(t *testing.T, environmentStateToCompare *MockExecutionEnvironmentState) {

	assert.Equal(
		t,
		environmentState.ExiterWasExited,
		environmentStateToCompare.ExiterWasExited,
	)

	assert.Exactly(
		t,
		environmentState.ExiterError,
		environmentStateToCompare.ExiterError,
	)

	assert.Exactly(
		t,
		environmentState.FileSystemDelegaterFiles,
		environmentStateToCompare.FileSystemDelegaterFiles,
	)

	assert.Exactly(
		t,
		environmentState.GitManagerInputHistoriesIndexedByRepoPath,
		environmentStateToCompare.GitManagerInputHistoriesIndexedByRepoPath,
	)

	assert.Exactly(
		t,
		environmentState.WritePrinterOutputContents,
		environmentStateToCompare.WritePrinterOutputContents,
	)

	assert.Exactly(
		t,
		environmentState.ShellCommandDelegaterInputHistory,
		environmentStateToCompare.ShellCommandDelegaterInputHistory,
	)

	assert.Exactly(
		t,
		environmentState.ShellCommandDelegaterOutputHistory,
		environmentStateToCompare.ShellCommandDelegaterOutputHistory,
	)

}
