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
		mockExecutionEnvironment = NewMockExecutionEnvironmentBuilder().Build()
	}
	mockExecutionEnvironment.Init()
	return mockExecutionEnvironment
}

// ------------------------------------------------------------------------------

type MockExecutionEnvironmentBuilder struct {
	MockShellCommandDelegaterBuilder *MockShellCommandDelegaterBuilder
	MockFileSystemDelegaterBuilder   *MockFileSystemDelegaterBuilder
	MockGitManagerCreatorBuilder     *MockGitManagerCreatorBuilder
}

func NewMockExecutionEnvironmentBuilder() *MockExecutionEnvironmentBuilder {
	return &MockExecutionEnvironmentBuilder{}
}

func (builder *MockExecutionEnvironmentBuilder) SetMockShellCommandDelegaterBuilder(mockShellCommandDelegaterBuilder *MockShellCommandDelegaterBuilder) *MockExecutionEnvironmentBuilder {
	builder.MockShellCommandDelegaterBuilder = mockShellCommandDelegaterBuilder
	return builder
}

func (builder *MockExecutionEnvironmentBuilder) SetMockFileSystemDelegaterBuilder(mockFileSystemDelegaterBuilder *MockFileSystemDelegaterBuilder) *MockExecutionEnvironmentBuilder {
	builder.MockFileSystemDelegaterBuilder = mockFileSystemDelegaterBuilder
	return builder
}

func (builder *MockExecutionEnvironmentBuilder) SetMockGitManagerCreatorBuilder(mockGitManagerCreatorBuilder *MockGitManagerCreatorBuilder) *MockExecutionEnvironmentBuilder {
	builder.MockGitManagerCreatorBuilder = mockGitManagerCreatorBuilder
	return builder
}

func (builder *MockExecutionEnvironmentBuilder) Build() *MockExecutionEnvironment {

	return &MockExecutionEnvironment{
		MockExiter:                NewMockExiter(),
		MockWritePrinter:          NewMockWritePrinter(),
		MockShellCommandDelegater: GetMockShellCommandDelegaterFromBuilderOrNil(builder.MockShellCommandDelegaterBuilder),
		MockFileSystemDelegater:   GetMockFileSystemDelegaterFromBuilderOrNil(builder.MockFileSystemDelegaterBuilder),
		MockGitManagerCreator:     GetMockGitManagerCreatorFromBuilderOrNil(builder.MockGitManagerCreatorBuilder),
	}

}

// Builds the MockExecutionEnvironment, initializes it with
// the MockExecutionEnvironment.Init() method, and returns the new MockExecutionEnvironment.
func (builder *MockExecutionEnvironmentBuilder) BuildAndInit() *MockExecutionEnvironment {
	mockExecutionEnvironment := builder.Build()
	mockExecutionEnvironment.Init()
	return mockExecutionEnvironment
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

	environment.MockExiter.Init()
	environment.MockWritePrinter.Init()
	environment.MockShellCommandDelegater.Init()
	environment.MockFileSystemDelegater.Init()

	if environment.MockGitManagerCreator != nil {
		environment.MockGitManagerCreator.Init()
	} else {
		util.GitManagerFactory = util.NewGitShellCommandProxyCreator()
	}

}

func (environment *MockExecutionEnvironment) GetCurrentState() *MockExecutionEnvironmentState {

	mockFileBuilders := make([]*MockFileBuilder, 0)
	for mockFilePath, mockFile := range environment.MockFileSystemDelegater.Files {
		mockFileBuilder := NewMockFileBuilder().
			SetFilePath(mockFilePath).
			SetContentsFromBytes(mockFile.Contents).
			SetWasClosed(mockFile.WasClosed)
		mockFileBuilders = append(mockFileBuilders, mockFileBuilder)
	}

	gitManagerInputHistoriesIndexedByRepoPath := make(map[string][][]string)
	if environment.MockGitManagerCreator != nil {
		for repoPath, gitManager := range environment.MockGitManagerCreator.MockGitManagersIndexedByRepoPath {
			gitManagerInputHistoriesIndexedByRepoPath[repoPath] = gitManager.InputHistory
		}
	}

	return NewMockExecutionEnvironmentStateBuilder().
		SetExiterWasExited(environment.MockExiter.WasExited).
		SetExiterError(environment.MockExiter.Error).
		SetMockFileBuilders(mockFileBuilders...).
		SetGitManagerInputHistoriesIndexedByRepoPath(gitManagerInputHistoriesIndexedByRepoPath).
		SetWritePrinterOutputContents(environment.MockWritePrinter.OutputContents).
		SetShellCommandDelegaterInputHistory(environment.MockShellCommandDelegater.InputHistory...).
		SetShellCommandDelegaterOutputHistory(environment.MockShellCommandDelegater.OutputHistory...).
		Build()

}

// ------------------------------------------------------------------------------

type MockExecutionEnvironmentStateBuilder struct {
	ExiterWasExited                           bool
	ExiterError                               error
	MockFileBuilders                          []*MockFileBuilder
	GitManagerInputHistoriesIndexedByRepoPath map[string][][]string
	WritePrinterOutputContents                []byte
	ShellCommandDelegaterInputHistory         []string
	ShellCommandDelegaterOutputHistory        []*MockShellCommandOutput
}

func NewMockExecutionEnvironmentStateBuilder() *MockExecutionEnvironmentStateBuilder {
	return &MockExecutionEnvironmentStateBuilder{
		ExiterWasExited:  false,
		ExiterError:      nil,
		MockFileBuilders: make([]*MockFileBuilder, 0),
		GitManagerInputHistoriesIndexedByRepoPath: make(map[string][][]string),
		WritePrinterOutputContents:                make([]byte, 0),
		ShellCommandDelegaterInputHistory:         make([]string, 0),
		ShellCommandDelegaterOutputHistory:        make([]*MockShellCommandOutput, 0),
	}
}

func (builder *MockExecutionEnvironmentStateBuilder) SetExiterWasExited(exiterWasExited bool) *MockExecutionEnvironmentStateBuilder {
	builder.ExiterWasExited = exiterWasExited
	return builder
}

func (builder *MockExecutionEnvironmentStateBuilder) SetExiterError(exiterError error) *MockExecutionEnvironmentStateBuilder {
	builder.ExiterError = exiterError
	return builder
}

func (builder *MockExecutionEnvironmentStateBuilder) SetMockFileBuilders(mockFileBuilders ...*MockFileBuilder) *MockExecutionEnvironmentStateBuilder {
	builder.MockFileBuilders = mockFileBuilders
	return builder
}

func (builder *MockExecutionEnvironmentStateBuilder) SetGitManagerInputHistoriesIndexedByRepoPath(gitManagerInputHistoriesIndexedByRepoPath map[string][][]string) *MockExecutionEnvironmentStateBuilder {
	builder.GitManagerInputHistoriesIndexedByRepoPath = gitManagerInputHistoriesIndexedByRepoPath
	return builder
}

func (builder *MockExecutionEnvironmentStateBuilder) SetWritePrinterOutputContents(writePrinterOutputContents []byte) *MockExecutionEnvironmentStateBuilder {
	builder.WritePrinterOutputContents = writePrinterOutputContents
	return builder
}

func (builder *MockExecutionEnvironmentStateBuilder) SetShellCommandDelegaterInputHistory(shellCommandDelegaterInputHistory ...string) *MockExecutionEnvironmentStateBuilder {
	builder.ShellCommandDelegaterInputHistory = shellCommandDelegaterInputHistory
	return builder
}

func (builder *MockExecutionEnvironmentStateBuilder) SetShellCommandDelegaterOutputHistory(shellCommandDelegaterOutputHistory ...*MockShellCommandOutput) *MockExecutionEnvironmentStateBuilder {
	builder.ShellCommandDelegaterOutputHistory = shellCommandDelegaterOutputHistory
	return builder
}

func (builder *MockExecutionEnvironmentStateBuilder) Build() *MockExecutionEnvironmentState {

	mockExecutionEnvironmentState := &MockExecutionEnvironmentState{
		ExiterWasExited:          builder.ExiterWasExited,
		ExiterError:              builder.ExiterError,
		FileSystemDelegaterFiles: make(map[string]*MockFile),
		GitManagerInputHistoriesIndexedByRepoPath: builder.GitManagerInputHistoriesIndexedByRepoPath,
		WritePrinterOutputContents:                builder.WritePrinterOutputContents,
		ShellCommandDelegaterInputHistory:         builder.ShellCommandDelegaterInputHistory,
		ShellCommandDelegaterOutputHistory:        builder.ShellCommandDelegaterOutputHistory,
	}

	for _, mockFileBuilder := range builder.MockFileBuilders {
		mockFile := mockFileBuilder.Build()
		mockExecutionEnvironmentState.FileSystemDelegaterFiles[mockFile.FilePath] = mockFile
	}

	return mockExecutionEnvironmentState

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

func (environmentState *MockExecutionEnvironmentState) AssertEquals(t *testing.T, environmentStateToCompare *MockExecutionEnvironmentState, testCaseDescription string) {

	assert.Equalf(
		t,
		environmentState.ExiterWasExited,
		environmentStateToCompare.ExiterWasExited,
		testCaseDescription,
	)

	assert.Exactlyf(
		t,
		environmentState.ExiterError,
		environmentStateToCompare.ExiterError,
		testCaseDescription,
	)

	assert.Exactlyf(
		t,
		environmentState.FileSystemDelegaterFiles,
		environmentStateToCompare.FileSystemDelegaterFiles,
		testCaseDescription,
	)

	assert.Exactlyf(
		t,
		environmentState.GitManagerInputHistoriesIndexedByRepoPath,
		environmentStateToCompare.GitManagerInputHistoriesIndexedByRepoPath,
		testCaseDescription,
	)

	assert.Exactlyf(
		t,
		environmentState.WritePrinterOutputContents,
		environmentStateToCompare.WritePrinterOutputContents,
		testCaseDescription,
	)

	assert.Exactlyf(
		t,
		environmentState.ShellCommandDelegaterInputHistory,
		environmentStateToCompare.ShellCommandDelegaterInputHistory,
		testCaseDescription,
	)

	assert.Exactlyf(
		t,
		environmentState.ShellCommandDelegaterOutputHistory,
		environmentStateToCompare.ShellCommandDelegaterOutputHistory,
		testCaseDescription,
	)

}
