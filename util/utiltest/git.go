package utiltest

import "github.com/stevengt/mppm/util"

// ------------------------------------------------------------------------------

type MockGitManagerCreator struct {
	MockGitManager *MockGitManager
}

func (mockGitManagerCreator *MockGitManagerCreator) NewGitManager(repoFilePath string) util.GitManager {
	return mockGitManagerCreator.MockGitManager
}

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
