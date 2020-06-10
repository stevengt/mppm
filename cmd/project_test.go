package cmd_test

import (
	"testing"

	"github.com/stevengt/mppm/cmd"
	"github.com/stevengt/mppm/config/configtest"
	"github.com/stretchr/testify/assert"

	"github.com/stevengt/mppm/util/utiltest"
)

func TestProjectCmd(t *testing.T) {

	testCases := []*ProjectCmdTestCase{

		&ProjectCmdTestCase{
			args:           []string{"project"},
			expectedOutput: "Provides utilities for managing a specific project.\n\nUsage:\n  mppm project [flags]\n  mppm project [command]\n\nAvailable Commands:\n  extract     Extracts all binary files of supported types into plain-text files, such as XML.\n  init        Initializes version control settings for a project using git and git-lfs.\n  restore     Restores all plain-text files of supported types to their original binary files.\n\nFlags:\n  -c, --commit-all         Equivalent to running 'mppm project extract; git add . -A; git commit -m '<commit message>'.\n  -h, --help               help for project\n  -p, --preview            Shows what files will be affected without actually making changes.\n  -u, --update-libraries   Updates the library versions in the project config file to match the\n                           current versions in the global config file.\n                           To see the global current versions, run 'mppm library --list'.\n\nUse \"mppm project [command] --help\" for more information about a command.\n",
		},
	}

	for _, testCase := range testCases {
		testCase.Run(t)
	}

}

type ProjectCmdTestCase struct {
	args                             []string
	executionEnvironmentBuilder      *utiltest.MockExecutionEnvironmentBuilder
	projectConfigFile                *utiltest.MockFile
	globalConfigFile                 *utiltest.MockFile
	expectedOutput                   string
	expectedGitManagerInputHistories map[string][][]string // Map of git repo file path -> git manager input history.
	expectedExitError                error
}

func (testCase *ProjectCmdTestCase) Run(t *testing.T) {

	executionEnvironment := utiltest.GetAndInitMockExecutionEnvironmentFromBuilderOrNil(testCase.executionEnvironmentBuilder)
	configtest.InitMockFileSystemDelegaterWithConfigFiles(
		executionEnvironment.MockFileSystemDelegater,
		testCase.projectConfigFile,
		testCase.globalConfigFile,
	)

	cmd.RootCmd.SetArgs(testCase.args)
	cmd.RootCmd.Execute()

	actualOutput := executionEnvironment.MockWritePrinter.GetOutputContentsAsString()
	assert.Equal(t, testCase.expectedOutput, actualOutput)

	for repoFilePath, expectedGitManagerInputHistory := range testCase.expectedGitManagerInputHistories {
		gitManager := executionEnvironment.MockGitManagerCreator.MockGitManagersIndexedByRepoPath[repoFilePath]
		assert.NotNil(t, gitManager)
		actualGitManagerInputHistory := gitManager.InputHistory
		assert.Exactly(t, expectedGitManagerInputHistory, actualGitManagerInputHistory)
	}

	if testCase.expectedExitError != nil {
		assert.True(t, executionEnvironment.MockExiter.WasExited)
	} else {
		assert.False(t, executionEnvironment.MockExiter.WasExited)
	}
	assert.Exactly(t, testCase.expectedExitError, executionEnvironment.MockExiter.Error)

}
