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

		&ProjectInitCmdTestCase{
			description: "Test that any error from os.Create() is properly raised.",
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockFileSystemDelegaterBuilder(
					utiltest.NewMockFileSystemDelegaterBuilder().
						SetUseDefaultCreateFileError(true),
				),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetExiterWasExited(true).
				SetExiterError(utiltest.DefaultCreateFileError).
				SetGitManagerInputHistoriesIndexedByRepoPath(
					map[string][][]string{
						".": make([][]string, 0),
					},
				),
		},

		&ProjectInitCmdTestCase{
			description: "Test that any error from running 'git init' is properly raised.",
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockGitManagerCreatorBuilder(
					utiltest.NewMockGitManagerCreatorBuilder().
						SetUseDefaultInitError(true),
				),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetExiterWasExited(true).
				SetExiterError(utiltest.DefaultInitError).
				SetMockFileBuilders(
					utiltest.NewMockFileBuilder().
						SetFilePath(config.MppmConfigFileName).
						SetContentsFromBytes(
							configtest.GetDefaultMppmConfigAsJson(),
						).
						SetWasClosed(true),
				).
				SetGitManagerInputHistoriesIndexedByRepoPath(
					map[string][][]string{
						".": [][]string{
							[]string{"init"},
						},
					},
				),
		},

		&ProjectInitCmdTestCase{
			description: "Test that any error from running 'git lfs install' is properly raised.",
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockGitManagerCreatorBuilder(
					utiltest.NewMockGitManagerCreatorBuilder().
						SetUseDefaultLfsInstallError(true),
				),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetExiterWasExited(true).
				SetExiterError(utiltest.DefaultLfsInstallError).
				SetMockFileBuilders(
					utiltest.NewMockFileBuilder().
						SetFilePath(config.MppmConfigFileName).
						SetContentsFromBytes(
							configtest.GetDefaultMppmConfigAsJson(),
						).
						SetWasClosed(true),
				).
				SetGitManagerInputHistoriesIndexedByRepoPath(
					map[string][][]string{
						".": [][]string{
							[]string{"init"},
							[]string{"lfs", "install"},
						},
					},
				),
		},

		&ProjectInitCmdTestCase{
			description: "Test that any error from running 'git lfs track' is properly raised.",
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockGitManagerCreatorBuilder(
					utiltest.NewMockGitManagerCreatorBuilder().
						SetUseDefaultLfsTrackError(true),
				),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetExiterWasExited(true).
				SetExiterError(utiltest.DefaultLfsTrackError).
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
						},
					},
				),
		},

		&ProjectInitCmdTestCase{
			description: "Test that any error from running 'git add' is properly raised.",
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockGitManagerCreatorBuilder(
					utiltest.NewMockGitManagerCreatorBuilder().
						SetUseDefaultAddError(true),
				),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetExiterWasExited(true).
				SetExiterError(utiltest.DefaultAddError).
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
						},
					},
				),
		},

		&ProjectInitCmdTestCase{
			description: "Test that any error from running 'git commit' is properly raised.",
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockGitManagerCreatorBuilder(
					utiltest.NewMockGitManagerCreatorBuilder().
						SetUseDefaultCommitError(true),
				),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetExiterWasExited(true).
				SetExiterError(utiltest.DefaultCommitError).
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
