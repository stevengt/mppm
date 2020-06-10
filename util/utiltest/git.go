package utiltest

import (
	"errors"

	"github.com/stevengt/mppm/util"
)

// ------------------------------------------------------------------------------

var DefaultInitError error = errors.New("There was a problem initializing the git repository.")

var DefaultAddError error = errors.New("There was a problem adding items to the git repository.")

var DefaultCommitError error = errors.New("There was a problem committing items to the git repository.")

var DefaultCheckoutError error = errors.New("There was a problem checking out the git repository.")

var DefaultRevParseError error = errors.New("Not a git repository.")

var DefaultLfsInstallError error = errors.New("There was a problem while setting up git lfs.")

var DefaultLfsTrackError error = errors.New("There was a problem trying to track files with git lfs.")

// ------------------------------------------------------------------------------

func GetMockGitManagerCreatorFromBuilderOrNil(mockGitManagerCreatorBuilder *MockGitManagerCreatorBuilder) *MockGitManagerCreator {
	if mockGitManagerCreatorBuilder != nil {
		return mockGitManagerCreatorBuilder.Build()
	} else {
		return NewDefaultMockGitManagerCreatorBuilder().Build()
	}
}

// ------------------------------------------------------------------------------

type MockGitManagerCreatorBuilder struct {
	RevParseStdout            string
	UseDefaultInitError       bool
	UseDefaultAddError        bool
	UseDefaultCommitError     bool
	UseDefaultCheckoutError   bool
	UseDefaultRevParseError   bool
	UseDefaultLfsInstallError bool
	UseDefaultLfsTrackError   bool
}

func NewDefaultMockGitManagerCreatorBuilder() *MockGitManagerCreatorBuilder {
	return &MockGitManagerCreatorBuilder{}
}

func (builder *MockGitManagerCreatorBuilder) Build() *MockGitManagerCreator {

	mockGitManager := &MockGitManager{
		RevParseStdout: builder.RevParseStdout,
	}

	if builder.UseDefaultInitError {
		mockGitManager.InitError = DefaultInitError
	}

	if builder.UseDefaultAddError {
		mockGitManager.AddError = DefaultAddError
	}

	if builder.UseDefaultCommitError {
		mockGitManager.CommitError = DefaultCommitError
	}

	if builder.UseDefaultCheckoutError {
		mockGitManager.CheckoutError = DefaultCheckoutError
	}

	if builder.UseDefaultRevParseError {
		mockGitManager.RevParseError = DefaultRevParseError
	}

	if builder.UseDefaultLfsInstallError {
		mockGitManager.LfsInstallError = DefaultLfsInstallError
	}

	if builder.UseDefaultLfsTrackError {
		mockGitManager.LfsTrackError = DefaultLfsTrackError
	}

	return &MockGitManagerCreator{
		MockGitManager: mockGitManager,
	}

}

// ------------------------------------------------------------------------------

type MockGitManagerCreator struct {
	MockGitManager                   *MockGitManager
	MockGitManagersIndexedByRepoPath map[string]*MockGitManager
}

func (mockGitManagerCreator *MockGitManagerCreator) NewGitManager(repoFilePath string) util.GitManager {
	mockGitManager := mockGitManagerCreator.MockGitManager
	mockGitManagerCreator.MockGitManagersIndexedByRepoPath[repoFilePath] = mockGitManager
	return mockGitManager
}

// ------------------------------------------------------------------------------

type MockGitManager struct {
	InputHistory    [][]string
	InitError       error
	AddError        error
	CommitError     error
	CheckoutError   error
	RevParseStdout  string
	RevParseError   error
	LfsInstallError error
	LfsTrackError   error
}

func (mockGitManager *MockGitManager) Init() (err error) {
	mockGitManager.appendToInputHistory("init")
	return mockGitManager.InitError
}

func (mockGitManager *MockGitManager) Add(args ...string) (err error) {
	mockGitManager.appendToInputHistory("add", args...)
	return mockGitManager.AddError
}

func (mockGitManager *MockGitManager) Commit(args ...string) (err error) {
	mockGitManager.appendToInputHistory("commit", args...)
	return mockGitManager.CommitError
}

func (mockGitManager *MockGitManager) Checkout(args ...string) (err error) {
	mockGitManager.appendToInputHistory("checkout", args...)
	return mockGitManager.CheckoutError
}

func (mockGitManager *MockGitManager) RevParse(args ...string) (stdout string, err error) {
	mockGitManager.appendToInputHistory("rev-parse", args...)
	return mockGitManager.RevParseStdout, mockGitManager.RevParseError
}

func (mockGitManager *MockGitManager) LfsInstall() (err error) {
	mockGitManager.appendToInputHistory("lfs", "install")
	return mockGitManager.LfsInstallError
}

func (mockGitManager *MockGitManager) LfsTrack(args ...string) (err error) {
	gitLfsCommandArgs := append(
		[]string{"track"},
		args...,
	)
	mockGitManager.appendToInputHistory("lfs", gitLfsCommandArgs...)
	return mockGitManager.LfsTrackError
}

func (mockGitManager *MockGitManager) AddAllAndCommit(commitMessage string) (err error) {
	err = mockGitManager.Add("-A", ".")
	if err != nil {
		return
	}
	err = mockGitManager.Commit("-m", commitMessage)
	if err != nil {
		return
	}
	return
}

func (mockGitManager *MockGitManager) appendToInputHistory(commandName string, args ...string) {
	input := append(
		[]string{commandName},
		args...,
	)
	mockGitManager.InputHistory = append(mockGitManager.InputHistory, input)
}
