package cmd_test

import (
	"fmt"
	"testing"

	"github.com/stevengt/mppm/config"
	"github.com/stevengt/mppm/config/configtest"

	"github.com/stevengt/mppm/cmd"
	"github.com/stevengt/mppm/util/utiltest"
)

func TestProjectRestoreCmd(t *testing.T) {

	testCases := []*ProjectRestoreCmdTestCase{

		&ProjectRestoreCmdTestCase{
			description: "Test that all supported compressed file types are restored from uncompressed files.",
			args:        []string{"project", "restore"},
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockFileSystemDelegaterBuilder(
					utiltest.NewMockFileSystemDelegaterBuilder().
						SetMockFileBuilders(
							configtest.ConfigWithValidVersionAndApplicationNameAndApplicationVersion.AsMockFileBuilder().
								SetFilePath(config.MppmConfigFileName),
							utiltest.GetFakeUncompressedAbletonLiveSetFileBuilder(),
							utiltest.GetFakeUncompressedAbletonLiveClipFileBuilder(),
							utiltest.GetPlainTextFileBuilder(),
						),
				),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetMockFileBuilders(
					configtest.ConfigWithValidVersionAndApplicationNameAndApplicationVersion.AsMockFileBuilder().
						SetFilePath(config.MppmConfigFileName).
						SetWasClosed(true),
					utiltest.GetFakeUncompressedAbletonLiveSetFileBuilder().
						SetWasClosed(true),
					utiltest.GetFakeAbletonLiveSetFileBuilder().
						SetWasClosed(true),
					utiltest.GetFakeUncompressedAbletonLiveClipFileBuilder().
						SetWasClosed(true),
					utiltest.GetFakeAbletonLiveClipFileBuilder().
						SetWasClosed(true),
					utiltest.GetPlainTextFileBuilder(),
				),
		},

		&ProjectRestoreCmdTestCase{
			description: "Test that all affected file changes are displayed without actually making the changes.",
			args:        []string{"project", "restore", "--preview"},
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockFileSystemDelegaterBuilder(
					utiltest.NewMockFileSystemDelegaterBuilder().
						SetMockFileBuilders(
							configtest.ConfigWithValidVersionAndApplicationNameAndApplicationVersion.AsMockFileBuilder().
								SetFilePath(config.MppmConfigFileName),
							utiltest.GetFakeUncompressedAbletonLiveSetFileBuilder(),
							utiltest.GetFakeUncompressedAbletonLiveClipFileBuilder(),
							utiltest.GetPlainTextFileBuilder(),
						),
				),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetMockFileBuilders(
					configtest.ConfigWithValidVersionAndApplicationNameAndApplicationVersion.AsMockFileBuilder().
						SetFilePath(config.MppmConfigFileName).
						SetWasClosed(true),
					utiltest.GetFakeUncompressedAbletonLiveSetFileBuilder(),
					utiltest.GetFakeUncompressedAbletonLiveClipFileBuilder(),
					utiltest.GetPlainTextFileBuilder(),
				).
				SetWritePrinterOutputContents(
					[]byte(
						fmt.Sprintf(
							"%s will be restored from %s\n%s will be restored from %s\n",
							utiltest.GetFakeAbletonLiveClipFileBuilder().FilePath,
							utiltest.GetFakeUncompressedAbletonLiveClipFileBuilder().FilePath,
							utiltest.GetFakeAbletonLiveSetFileBuilder().FilePath,
							utiltest.GetFakeUncompressedAbletonLiveSetFileBuilder().FilePath,
						),
					),
				),
		},

		&ProjectRestoreCmdTestCase{
			description: "Test that any error resulting from an invalid config file is properly raised.",
			args:        []string{"project", "restore"},
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockFileSystemDelegaterBuilder(
					utiltest.NewMockFileSystemDelegaterBuilder().
						SetMockFileBuilders(
							configtest.ConfigWithInvalidVersionAndNoApplications.AsMockFileBuilder().
								SetFilePath(config.MppmConfigFileName),
						),
				),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetExiterWasExited(true).
				SetExiterError(configtest.ConfigWithInvalidVersionAndNoApplications.ExpectedError).
				SetMockFileBuilders(
					configtest.ConfigWithInvalidVersionAndNoApplications.AsMockFileBuilder().
						SetFilePath(config.MppmConfigFileName).
						SetWasClosed(true),
				),
		},

		&ProjectRestoreCmdTestCase{
			description: "Test that any error from filepath.Walk() is properly raised.",
			args:        []string{"project", "restore"},
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockFileSystemDelegaterBuilder(
					utiltest.NewMockFileSystemDelegaterBuilder().
						SetMockFileBuilders(
							configtest.ConfigWithValidVersionAndApplicationNameAndApplicationVersion.AsMockFileBuilder().
								SetFilePath(config.MppmConfigFileName),
							utiltest.GetFakeUncompressedAbletonLiveSetFileBuilder(),
							utiltest.GetFakeUncompressedAbletonLiveClipFileBuilder(),
							utiltest.GetPlainTextFileBuilder(),
						).
						SetUseDefaultWalkFilePathError(true),
				),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetExiterWasExited(true).
				SetExiterError(utiltest.DefaultWalkFilePathError).
				SetMockFileBuilders(
					configtest.ConfigWithValidVersionAndApplicationNameAndApplicationVersion.AsMockFileBuilder().
						SetFilePath(config.MppmConfigFileName).
						SetWasClosed(true),
					utiltest.GetFakeUncompressedAbletonLiveSetFileBuilder(),
					utiltest.GetFakeUncompressedAbletonLiveClipFileBuilder(),
					utiltest.GetPlainTextFileBuilder(),
				),
		},

		&ProjectRestoreCmdTestCase{
			description: "Test that any error from os.Create() is properly raised.",
			args:        []string{"project", "restore"},
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockFileSystemDelegaterBuilder(
					utiltest.NewMockFileSystemDelegaterBuilder().
						SetMockFileBuilders(
							configtest.ConfigWithValidVersionAndApplicationNameAndApplicationVersion.AsMockFileBuilder().
								SetFilePath(config.MppmConfigFileName),
							utiltest.GetFakeUncompressedAbletonLiveSetFileBuilder(),
						).
						SetUseDefaultCreateFileError(true),
				),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetExiterWasExited(true).
				SetExiterError(utiltest.DefaultCreateFileError).
				SetMockFileBuilders(
					configtest.ConfigWithValidVersionAndApplicationNameAndApplicationVersion.AsMockFileBuilder().
						SetFilePath(config.MppmConfigFileName).
						SetWasClosed(true),
					utiltest.GetFakeUncompressedAbletonLiveSetFileBuilder().
						SetWasClosed(true),
				),
		},

		&ProjectRestoreCmdTestCase{
			description: "Test that any error from os.Rename() is properly raised.",
			args:        []string{"project", "restore"},
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockFileSystemDelegaterBuilder(
					utiltest.NewMockFileSystemDelegaterBuilder().
						SetMockFileBuilders(
							configtest.ConfigWithValidVersionAndApplicationNameAndApplicationVersion.AsMockFileBuilder().
								SetFilePath(config.MppmConfigFileName),
							utiltest.GetFakeUncompressedAbletonLiveSetFileBuilder(),
						).
						SetUseDefaultRenameFileError(true),
				),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetExiterWasExited(true).
				SetExiterError(utiltest.DefaultRenameFileError).
				SetMockFileBuilders(
					configtest.ConfigWithValidVersionAndApplicationNameAndApplicationVersion.AsMockFileBuilder().
						SetFilePath(config.MppmConfigFileName).
						SetWasClosed(true),
					utiltest.GetFakeUncompressedAbletonLiveSetFileBuilder().
						SetWasClosed(true),
					utiltest.GetFakeAbletonLiveSetFileBuilder().
						SetFilePath(
							utiltest.GetFakeUncompressedAbletonLiveSetFileBuilder().FilePath+".gz",
						).
						SetWasClosed(true),
				),
		},
	}

	for _, testCase := range testCases {
		testCase.Run(t)
	}

}

type ProjectRestoreCmdTestCase struct {
	description                              string
	args                                     []string
	mockExecutionEnvironmentBuilder          *utiltest.MockExecutionEnvironmentBuilder
	expectedExecutionEnvironmentStateBuilder *utiltest.MockExecutionEnvironmentStateBuilder
}

func (testCase *ProjectRestoreCmdTestCase) Run(t *testing.T) {

	mockExecutionEnvironment := testCase.mockExecutionEnvironmentBuilder.BuildAndInit()

	cmd.RootCmd.SetArgs(testCase.args)
	cmd.RootCmd.Execute()

	expectedExecutionEnvironmentState := testCase.expectedExecutionEnvironmentStateBuilder.Build()
	mockExecutionEnvironment.GetCurrentState().AssertEquals(t, expectedExecutionEnvironmentState, testCase.description)

}
