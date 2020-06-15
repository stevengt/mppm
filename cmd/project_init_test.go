package cmd_test

import (
	"strings"
	"testing"

	"github.com/stevengt/mppm/config/applications"

	"github.com/stevengt/mppm/config/configtest"

	"github.com/stevengt/mppm/cmd"
	"github.com/stevengt/mppm/config"
	"github.com/stevengt/mppm/util/utiltest"
)

func TestProjectInitCmd(t *testing.T) {

	testCases := []*ProjectInitCmdTestCase{

		&ProjectInitCmdTestCase{
			description:                     "Test that the the project git and git-lfs settings are initialized, and a config file is created.",
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder(),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetMockFileBuilders(
					utiltest.NewMockFileBuilder().
						SetFilePath(config.MppmConfigFileName).
						SetContentsFromBytes(
							configtest.GetDefaultMppmConfigAsJson(),
						).
						SetWasClosed(true),
					utiltest.NewMockFileBuilder().
						SetFilePath(".gitignore").
						SetContentsFromString(
							strings.Join(applications.GetAllFilePatternsConfig().GitIgnorePatterns, "\n"),
						).
						SetWasClosed(true),
				).
				SetGitManagerInputHistoriesIndexedByRepoPath(
					map[string][][]string{
						".": [][]string{
							[]string{"init"},
							[]string{"lfs", "install"},
							append([]string{"lfs", "track"}, applications.GetAllFilePatternsConfig().GitLfsTrackPatterns...),
							[]string{"add", ".gitignore", ".gitattributes", config.MppmConfigFileName},
							[]string{"commit", "-m", "Initial commit."},
						},
					},
				),
		},
	}

	for _, testCase := range testCases {
		testCase.Run(t)
	}

}

type ProjectInitCmdTestCase struct {
	description                              string
	mockExecutionEnvironmentBuilder          *utiltest.MockExecutionEnvironmentBuilder
	expectedExecutionEnvironmentStateBuilder *utiltest.MockExecutionEnvironmentStateBuilder
}

func (testCase *ProjectInitCmdTestCase) Run(t *testing.T) {

	mockExecutionEnvironment := testCase.mockExecutionEnvironmentBuilder.BuildAndInit()

	cmd.RootCmd.SetArgs([]string{"project", "init"})
	cmd.RootCmd.Execute()

	expectedExecutionEnvironmentState := testCase.expectedExecutionEnvironmentStateBuilder.Build()
	mockExecutionEnvironment.GetCurrentState().AssertEquals(t, expectedExecutionEnvironmentState, testCase.description)

}
