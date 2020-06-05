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

var DefaultAddAllAndCommitError error = errors.New("There was a problem trying to add all and commit.")

// ------------------------------------------------------------------------------

type MockGitManagerCreatorBuilder struct {
	RevParseStdout                 string
	UseDefaultInitError            bool
	UseDefaultAddError             bool
	UseDefaultCommitError          bool
	UseDefaultCheckoutError        bool
	UseDefaultRevParseError        bool
	UseDefaultLfsInstallError      bool
	UseDefaultLfsTrackError        bool
	UseDefaultAddAllAndCommitError bool
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

	if builder.UseDefaultAddAllAndCommitError {
		mockGitManager.AddAllAndCommitError = DefaultAddAllAndCommitError
	}

	return &MockGitManagerCreator{
		MockGitManager: mockGitManager,
	}

}

// ------------------------------------------------------------------------------

type MockGitManagerCreator struct {
	MockGitManager *MockGitManager
}

func (mockGitManagerCreator *MockGitManagerCreator) NewGitManager(repoFilePath string) util.GitManager {
	return mockGitManagerCreator.MockGitManager
}

// ------------------------------------------------------------------------------

type MockGitManager struct {
	InitError            error
	AddError             error
	CommitError          error
	CheckoutError        error
	RevParseStdout       string
	RevParseError        error
	LfsInstallError      error
	LfsTrackError        error
	AddAllAndCommitError error
}

func (mockGitManager *MockGitManager) Init() (err error) {
	return mockGitManager.InitError
}

func (mockGitManager *MockGitManager) Add(args ...string) (err error) {
	return mockGitManager.AddError
}

func (mockGitManager *MockGitManager) Commit(args ...string) (err error) {
	return mockGitManager.CommitError
}

func (mockGitManager *MockGitManager) Checkout(args ...string) (err error) {
	return mockGitManager.CheckoutError
}

func (mockGitManager *MockGitManager) RevParse(args ...string) (stdout string, err error) {
	return mockGitManager.RevParseStdout, mockGitManager.RevParseError
}

func (mockGitManager *MockGitManager) LfsInstall() (err error) {
	return mockGitManager.LfsInstallError
}

func (mockGitManager *MockGitManager) LfsTrack(args ...string) (err error) {
	return mockGitManager.LfsTrackError
}

func (mockGitManager *MockGitManager) AddAllAndCommit(commitMessage string) (err error) {
	return mockGitManager.AddAllAndCommitError
}
