package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stevengt/mppm/config"
	"github.com/stevengt/mppm/util/utiltest"
)

func TestUpdateCurrentGitCommitId(t *testing.T) {

	libraryConfig := &config.LibraryConfig{
		FilePath:              ".",
		MostRecentGitCommitId: "Git commit = 789",
		CurrentGitCommitId:    "Git commit = 123",
	}

	// Test that the current library git commit id is correctly updated.
	utiltest.NewMockGitManagerCreatorBuilder().
		SetRevParseStdout("Git commit = 456").
		Build().
		Init()

	err := libraryConfig.UpdateCurrentGitCommitId()
	assert.Nil(t, err)
	assert.Equal(t, "Git commit = 456", libraryConfig.CurrentGitCommitId)

	// Test that any error from 'git rev-parse' is correctly raised.
	utiltest.NewMockGitManagerCreatorBuilder().
		SetUseDefaultRevParseError(true).
		Build().
		Init()

	err = libraryConfig.UpdateCurrentGitCommitId()
	assert.Exactly(t, utiltest.DefaultRevParseError, err)
	assert.Equal(t, "Git commit = 456", libraryConfig.CurrentGitCommitId)

}
