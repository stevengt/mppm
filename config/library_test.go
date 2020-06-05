package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stevengt/mppm/config"
	"github.com/stevengt/mppm/util"
	"github.com/stevengt/mppm/util/utiltest"
)

func TestUpdateCurrentGitCommitId(t *testing.T) {

	libraryConfig := &config.LibraryConfig{
		FilePath:              ".",
		MostRecentGitCommitId: "Git commit = 789",
		CurrentGitCommitId:    "Git commit = 123",
	}

	mockGitManagerCreatorBuilder := &utiltest.MockGitManagerCreatorBuilder{
		RevParseStdout: "Git commit = 456",
	}
	util.GitManagerFactory = mockGitManagerCreatorBuilder.Build()

	err := libraryConfig.UpdateCurrentGitCommitId()
	assert.Nil(t, err)
	assert.Equal(t, "Git commit = 456", libraryConfig.CurrentGitCommitId)

	mockGitManagerCreatorBuilder = &utiltest.MockGitManagerCreatorBuilder{
		RevParseStdout:          "",
		UseDefaultRevParseError: true,
	}
	util.GitManagerFactory = mockGitManagerCreatorBuilder.Build()

	err = libraryConfig.UpdateCurrentGitCommitId()
	assert.Exactly(t, utiltest.DefaultRevParseError, err)
	assert.Equal(t, "Git commit = 456", libraryConfig.CurrentGitCommitId)

}
