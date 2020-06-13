package cmd_test

import (
	"errors"
	"testing"

	"github.com/stevengt/mppm/config"
	"github.com/stevengt/mppm/config/configtest"

	"github.com/stevengt/mppm/cmd"

	"github.com/stevengt/mppm/util/utiltest"
)

var projectCmdHelpMessage string = "Provides utilities for managing a specific project.\n\nUsage:\n  mppm project [flags]\n  mppm project [command]\n\nAvailable Commands:\n  extract     Extracts all binary files of supported types into plain-text files, such as XML.\n  init        Initializes version control settings for a project using git and git-lfs.\n  restore     Restores all plain-text files of supported types to their original binary files.\n\nFlags:\n  -c, --commit-all         Equivalent to running 'mppm project extract; git add . -A; git commit -m '<commit message>'.\n  -h, --help               help for project\n  -p, --preview            Shows what files will be affected without actually making changes.\n  -u, --update-libraries   Updates the library versions in the project config file to match the\n                           current versions in the global config file.\n                           To see the global current versions, run 'mppm library --list'.\n\nUse \"mppm project [command] --help\" for more information about a command.\n"

func TestProjectCmd(t *testing.T) {

	testCases := []*ProjectCmdTestCase{

		&ProjectCmdTestCase{
			description:                     "Test that the project help message is displayed if no args are given.",
			args:                            []string{"project"},
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder(),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetWritePrinterOutputContents(
					[]byte(projectCmdHelpMessage),
				),
		},

		&ProjectCmdTestCase{
			description:                     "Test that the project help message is displayed if invalid args are given.",
			args:                            []string{"project", "invalid", "args"},
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder(),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetWritePrinterOutputContents(
					[]byte(projectCmdHelpMessage),
				),
		},

		&ProjectCmdTestCase{
			description: "Test that the library versions in the project config are updated to match those in the global config.",
			args:        []string{"project", "--update-libraries"},
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockFileSystemDelegaterBuilder(
					utiltest.NewMockFileSystemDelegaterBuilder().
						SetMockFileBuilders(
							configtest.ConfigWithAllValidInfoAndPreviousLibraryVersion.AsMockFileBuilder().
								SetFilePath(config.MppmConfigFileName),
							configtest.ConfigWithAllValidInfoAndMostRecentLibraryVersion.AsMockFileBuilder().
								SetFilePath("/home/testuser/.mppm.json"),
						),
				),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetMockFileBuilders(
					configtest.ConfigWithAllValidInfoAndMostRecentLibraryVersion.AsMockFileBuilder().
						SetFilePath(config.MppmConfigFileName).
						SetWasClosed(true),
					configtest.ConfigWithAllValidInfoAndMostRecentLibraryVersion.AsMockFileBuilder().
						SetFilePath("/home/testuser/.mppm.json").
						SetWasClosed(true),
				),
		},

		&ProjectCmdTestCase{
			description: "Test that any errors from os.Open() are correctly raised and the process is exited.",
			args:        []string{"project", "--update-libraries"},
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockFileSystemDelegaterBuilder(
					utiltest.NewMockFileSystemDelegaterBuilder().
						SetUseDefaultOpenFileError(true),
				),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetExiterWasExited(true).
				SetExiterError(errors.New("\nThere was a problem while opening the mppm config file.\nIf the file doesn't exist, try running 'mppm project init' first.\nThere was a problem opening the file.\n")),
		},

		&ProjectCmdTestCase{
			description: "Test that the process extracts compressed files and invokes 'git add -A . && git commit -m <commit-message>'.",
			args:        []string{"project", "--commit-all", "Made changes"},
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockFileSystemDelegaterBuilder(
					utiltest.NewMockFileSystemDelegaterBuilder().
						SetMockFileBuilders(
							configtest.ConfigWithValidVersionAndApplicationNameAndApplicationVersion.AsMockFileBuilder().
								SetFilePath(config.MppmConfigFileName),
							utiltest.GetFakeAbletonLiveSetFileBuilder(),
							utiltest.GetFakeAbletonLiveClipFileBuilder(),
						),
				),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetMockFileBuilders(
					configtest.ConfigWithValidVersionAndApplicationNameAndApplicationVersion.AsMockFileBuilder().
						SetFilePath(config.MppmConfigFileName).
						SetWasClosed(true),
					utiltest.GetFakeAbletonLiveSetFileBuilder().
						SetWasClosed(true),
					utiltest.GetFakeUncompressedAbletonLiveSetFileBuilder().
						SetWasClosed(true),
					utiltest.GetFakeAbletonLiveClipFileBuilder().
						SetWasClosed(true),
					utiltest.GetFakeUncompressedAbletonLiveClipFileBuilder().
						SetWasClosed(true),
				).
				SetGitManagerInputHistoriesIndexedByRepoPath(
					map[string][][]string{
						".": [][]string{
							[]string{"add", "-A", "."},
							[]string{"commit", "-m", "Made changes"},
						},
					},
				),
		},

		&ProjectCmdTestCase{
			description: "Test that any error from 'git add' is properly raised.",
			args:        []string{"project", "--commit-all", "Made changes"},
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockFileSystemDelegaterBuilder(
					utiltest.NewMockFileSystemDelegaterBuilder().
						SetMockFileBuilders(
							configtest.ConfigWithValidVersionAndApplicationNameAndApplicationVersion.AsMockFileBuilder().
								SetFilePath(config.MppmConfigFileName),
						),
				).
				SetMockGitManagerCreatorBuilder(
					utiltest.NewMockGitManagerCreatorBuilder().
						SetUseDefaultAddError(true),
				),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetExiterError(utiltest.DefaultAddError).
				SetExiterWasExited(true).
				SetMockFileBuilders(
					configtest.ConfigWithValidVersionAndApplicationNameAndApplicationVersion.AsMockFileBuilder().
						SetFilePath(config.MppmConfigFileName).
						SetWasClosed(true),
				).
				SetGitManagerInputHistoriesIndexedByRepoPath(
					map[string][][]string{
						".": [][]string{
							[]string{"add", "-A", "."},
						},
					},
				),
		},

		&ProjectCmdTestCase{
			description: "Test that any error from 'git commit' is properly raised.",
			args:        []string{"project", "--commit-all", "Made changes"},
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockFileSystemDelegaterBuilder(
					utiltest.NewMockFileSystemDelegaterBuilder().
						SetMockFileBuilders(
							configtest.ConfigWithValidVersionAndApplicationNameAndApplicationVersion.AsMockFileBuilder().
								SetFilePath(config.MppmConfigFileName),
						),
				).
				SetMockGitManagerCreatorBuilder(
					utiltest.NewMockGitManagerCreatorBuilder().
						SetUseDefaultCommitError(true),
				),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetExiterError(utiltest.DefaultCommitError).
				SetExiterWasExited(true).
				SetMockFileBuilders(
					configtest.ConfigWithValidVersionAndApplicationNameAndApplicationVersion.AsMockFileBuilder().
						SetFilePath(config.MppmConfigFileName).
						SetWasClosed(true),
				).
				SetGitManagerInputHistoriesIndexedByRepoPath(
					map[string][][]string{
						".": [][]string{
							[]string{"add", "-A", "."},
							[]string{"commit", "-m", "Made changes"},
						},
					},
				),
		},
	}

	for _, testCase := range testCases {
		testCase.Run(t)
	}

}

type ProjectCmdTestCase struct {
	description                              string
	args                                     []string
	mockExecutionEnvironmentBuilder          *utiltest.MockExecutionEnvironmentBuilder
	expectedExecutionEnvironmentStateBuilder *utiltest.MockExecutionEnvironmentStateBuilder
}

func (testCase *ProjectCmdTestCase) Run(t *testing.T) {

	mockExecutionEnvironment := testCase.mockExecutionEnvironmentBuilder.BuildAndInit()

	cmd.RootCmd.SetArgs(testCase.args)
	cmd.RootCmd.Execute()

	expectedExecutionEnvironmentState := testCase.expectedExecutionEnvironmentStateBuilder.Build()
	mockExecutionEnvironment.GetCurrentState().AssertEquals(t, expectedExecutionEnvironmentState, testCase.description)

}
