package cmd_test

import (
	"fmt"
	"testing"

	"github.com/stevengt/mppm/config"
	"github.com/stevengt/mppm/config/configtest"

	"github.com/stevengt/mppm/cmd"
	"github.com/stevengt/mppm/util/utiltest"
)

func TestProjectExtractCmd(t *testing.T) {

	testCases := []*ProjectExtractCmdTestCase{

		&ProjectExtractCmdTestCase{
			description: "Test that all supported compressed file types are extracted to uncompressed files.",
			args:        []string{"project", "extract"},
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockFileSystemDelegaterBuilder(
					utiltest.NewMockFileSystemDelegaterBuilder().
						SetMockFileBuilders(
							configtest.ConfigWithValidVersionAndApplicationNameAndApplicationVersion.AsMockFileBuilder().
								SetFilePath(config.MppmConfigFileName),
							utiltest.GetFakeAbletonLiveSetFileBuilder(),
							utiltest.GetFakeAbletonLiveClipFileBuilder(),
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

		&ProjectExtractCmdTestCase{
			description: "Test that all affected file changes are displayed without actually making the changes.",
			args:        []string{"project", "extract", "--preview"},
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockFileSystemDelegaterBuilder(
					utiltest.NewMockFileSystemDelegaterBuilder().
						SetMockFileBuilders(
							configtest.ConfigWithValidVersionAndApplicationNameAndApplicationVersion.AsMockFileBuilder().
								SetFilePath(config.MppmConfigFileName),
							utiltest.GetFakeAbletonLiveSetFileBuilder(),
							utiltest.GetFakeAbletonLiveClipFileBuilder(),
							utiltest.GetPlainTextFileBuilder(),
						),
				),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetMockFileBuilders(
					configtest.ConfigWithValidVersionAndApplicationNameAndApplicationVersion.AsMockFileBuilder().
						SetFilePath(config.MppmConfigFileName).
						SetWasClosed(true),
					utiltest.GetFakeAbletonLiveSetFileBuilder(),
					utiltest.GetFakeAbletonLiveClipFileBuilder(),
					utiltest.GetPlainTextFileBuilder(),
				).
				SetWritePrinterOutputContents(
					[]byte(
						fmt.Sprintf(
							"%s will be extracted to %s\n%s will be extracted to %s\n",
							utiltest.GetFakeAbletonLiveClipFileBuilder().FilePath,
							utiltest.GetFakeUncompressedAbletonLiveClipFileBuilder().FilePath,
							utiltest.GetFakeAbletonLiveSetFileBuilder().FilePath,
							utiltest.GetFakeUncompressedAbletonLiveSetFileBuilder().FilePath,
						),
					),
				),
		},

		&ProjectExtractCmdTestCase{
			description: "Test that any error resulting from an invalid config file is properly raised.",
			args:        []string{"project", "extract"},
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

		&ProjectExtractCmdTestCase{
			description: "Test that any error from filepath.Walk() is properly raised.",
			args:        []string{"project", "extract"},
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockFileSystemDelegaterBuilder(
					utiltest.NewMockFileSystemDelegaterBuilder().
						SetMockFileBuilders(
							configtest.ConfigWithValidVersionAndApplicationNameAndApplicationVersion.AsMockFileBuilder().
								SetFilePath(config.MppmConfigFileName),
							utiltest.GetFakeAbletonLiveSetFileBuilder(),
							utiltest.GetFakeAbletonLiveClipFileBuilder(),
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
					utiltest.GetFakeAbletonLiveSetFileBuilder(),
					utiltest.GetFakeAbletonLiveClipFileBuilder(),
					utiltest.GetPlainTextFileBuilder(),
				),
		},

		&ProjectExtractCmdTestCase{
			description: "Test that any error from os.Create() is properly raised.",
			args:        []string{"project", "extract"},
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockFileSystemDelegaterBuilder(
					utiltest.NewMockFileSystemDelegaterBuilder().
						SetMockFileBuilders(
							configtest.ConfigWithValidVersionAndApplicationNameAndApplicationVersion.AsMockFileBuilder().
								SetFilePath(config.MppmConfigFileName),
							utiltest.GetFakeAbletonLiveSetFileBuilder(),
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
					utiltest.GetFakeAbletonLiveSetFileBuilder().
						SetWasClosed(true),
				),
		},
	}

	for _, testCase := range testCases {
		testCase.Run(t)
	}

}

type ProjectExtractCmdTestCase struct {
	description                              string
	args                                     []string
	mockExecutionEnvironmentBuilder          *utiltest.MockExecutionEnvironmentBuilder
	expectedExecutionEnvironmentStateBuilder *utiltest.MockExecutionEnvironmentStateBuilder
}

func (testCase *ProjectExtractCmdTestCase) Run(t *testing.T) {

	mockExecutionEnvironment := testCase.mockExecutionEnvironmentBuilder.BuildAndInit()

	cmd.RootCmd.SetArgs(testCase.args)
	cmd.RootCmd.Execute()

	expectedExecutionEnvironmentState := testCase.expectedExecutionEnvironmentStateBuilder.Build()
	mockExecutionEnvironment.GetCurrentState().AssertEquals(t, expectedExecutionEnvironmentState, testCase.description)

}
