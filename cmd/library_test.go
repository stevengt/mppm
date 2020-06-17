package cmd_test

import (
	"testing"

	"github.com/stevengt/mppm/config/configtest"

	"github.com/stevengt/mppm/cmd"
	"github.com/stevengt/mppm/util/utiltest"
)

func TestLibraryCmd(t *testing.T) {

	testCases := []*LibraryCmdTestCase{

		&LibraryCmdTestCase{
			description: "Test that all libraries in the global config file are displayed.",
			args:        []string{"library", "--list"},
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockFileSystemDelegaterBuilder(
					utiltest.NewMockFileSystemDelegaterBuilder().
						SetMockFileBuilders(
							configtest.ConfigWithAllValidInfoAndMostRecentLibraryVersion.AsMockFileBuilder().
								SetFilePath("/home/testuser/.mppm.json"),
						),
				),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetMockFileBuilders(
					configtest.ConfigWithAllValidInfoAndMostRecentLibraryVersion.AsMockFileBuilder().
						SetFilePath("/home/testuser/.mppm.json").
						SetWasClosed(true),
				).
				SetWritePrinterOutputContents(
					[]byte("\n/home/testuser/library\n\tmost-recent-version=\"56789\"\n\tcurrent-version=\"56789\"\n\n"),
				),
		},

		&LibraryCmdTestCase{
			description: "Test that changes to all libraries are committed with git, and that the global config file is updated.",
			args:        []string{"library", "--commit-all"},
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockFileSystemDelegaterBuilder(
					utiltest.NewMockFileSystemDelegaterBuilder().
						SetMockFileBuilders(
							configtest.ConfigWithAllValidInfoAndPreviousLibraryVersion.AsMockFileBuilder().
								SetFilePath("/home/testuser/.mppm.json"),
						),
				).
				SetMockGitManagerCreatorBuilder(
					utiltest.NewMockGitManagerCreatorBuilder().
						SetRevParseStdout("56789"),
				),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetMockFileBuilders(
					configtest.ConfigWithAllValidInfoAndMostRecentLibraryVersion.AsMockFileBuilder().
						SetFilePath("/home/testuser/.mppm.json").
						SetWasClosed(true),
				).
				SetGitManagerInputHistoriesIndexedByRepoPath(
					map[string][][]string{
						"/home/testuser/library": [][]string{
							[]string{"add", "-A", "."},
							[]string{"commit", "-m", "Committed all changes."},
							[]string{"rev-parse", "HEAD"},
						},
					},
				),
		},
	}

	for _, testCase := range testCases {
		testCase.Run(t)
	}

}

type LibraryCmdTestCase struct {
	description                              string
	args                                     []string
	mockExecutionEnvironmentBuilder          *utiltest.MockExecutionEnvironmentBuilder
	expectedExecutionEnvironmentStateBuilder *utiltest.MockExecutionEnvironmentStateBuilder
}

func (testCase *LibraryCmdTestCase) Run(t *testing.T) {

	mockExecutionEnvironment := testCase.mockExecutionEnvironmentBuilder.BuildAndInit()

	cmd.RootCmd.SetArgs(testCase.args)
	cmd.RootCmd.Execute()

	expectedExecutionEnvironmentState := testCase.expectedExecutionEnvironmentStateBuilder.Build()
	mockExecutionEnvironment.GetCurrentState().AssertEquals(t, expectedExecutionEnvironmentState, testCase.description)

}
