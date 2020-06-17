package cmd_test

import (
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
