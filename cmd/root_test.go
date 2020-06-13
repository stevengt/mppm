package cmd_test

import (
	"testing"

	"github.com/stevengt/mppm/cmd"
	"github.com/stevengt/mppm/util/utiltest"
)

var rootCmdHelpMessage string = "Short for 'Music Production Project Manager', mppm provides utilities for managing music production projects, such as:\n\n\t- Simplified version control using 'git' and 'git-lfs'.\n\t- Extraction of 'Ableton Live Set' files to/from raw XML files.\n\nUsage:\n  mppm [flags]\n  mppm [command]\n\nAvailable Commands:\n  help        Help about any command\n  library     Provides utilities for globally managing multiple libraries (folders).\n  project     Provides utilities for managing a specific project.\n\nFlags:\n  -h, --help             help for mppm\n  -s, --show-supported   Shows what file types are supported by mppm.\n  -v, --version          version for mppm\n\nUse \"mppm [command] --help\" for more information about a command.\n"

func TestRootCmd(t *testing.T) {

	testCases := []*RootCmdTestCase{

		&RootCmdTestCase{
			description: "Test that the root help message is displayed if no args are given.",
			args:        nil,
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetWritePrinterOutputContents(
					[]byte(rootCmdHelpMessage),
				),
		},

		&RootCmdTestCase{
			description: "Test that the help message is displayed if invalid args are given.",
			args:        []string{"invalid", "args"},
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetWritePrinterOutputContents(
					[]byte(rootCmdHelpMessage),
				),
		},

		&RootCmdTestCase{
			description: "Test that the --show-supported flag correctly prints all supported file types.",
			args:        []string{"--show-supported"},
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetWritePrinterOutputContents(
					[]byte("Audio\n\n\tGit Ignore Patterns\n\t\t\n\tGit LFS Track Patterns\n\t\t*.3gp\n\t\t*.aa\n\t\t*.aac\n\t\t*.aax\n\t\t*.act\n\t\t*.aiff\n\t\t*.alac\n\t\t*.amr\n\t\t*.ape\n\t\t*.au\n\t\t*.awb\n\t\t*.dct\n\t\t*.dss\n\t\t*.dvf\n\t\t*.flac\n\t\t*.gsm\n\t\t*.iklax\n\t\t*.ivs\n\t\t*.m4a\n\t\t*.m4b\n\t\t*.m4p\n\t\t*.mmf\n\t\t*.mp3\n\t\t*.mpc\n\t\t*.msv\n\t\t*.nmf\n\t\t*.nsf\n\t\t*.ogg\n\t\t*.oga\n\t\t*.mogg\n\t\t*.opus\n\t\t*.ra\n\t\t*.rm\n\t\t*.raw\n\t\t*.rf64\n\t\t*.sln\n\t\t*.tta\n\t\t*.voc\n\t\t*.vox\n\t\t*.wav\n\t\t*.wma\n\t\t*.wv\n\t\t*.webm\n\t\t*.8svx\n\t\t*.cda\n\tGzipped XML File Types\n\t\t\nAbleton 10\n\n\tGit Ignore Patterns\n\t\tBackup/\n\t\t*.als\n\t\t*.alc\n\t\t*.adv\n\t\t*.adg\n\tGit LFS Track Patterns\n\t\t*.alp\n\t\t*.asd\n\t\t*.agr\n\t\t*.ams\n\t\t*.amxd\n\tGzipped XML File Types\n\t\tals\n\t\talc\n\t\tadv\n\t\tadg\n"),
				),
		},
	}

	for _, testCase := range testCases {
		testCase.Run(t)
	}

}

type RootCmdTestCase struct {
	description                              string
	args                                     []string
	expectedExecutionEnvironmentStateBuilder *utiltest.MockExecutionEnvironmentStateBuilder
}

func (testCase *RootCmdTestCase) Run(t *testing.T) {

	mockExecutionEnvironment := utiltest.NewMockExecutionEnvironmentBuilder().BuildAndInit()

	cmd.RootCmd.SetArgs(testCase.args)
	cmd.RootCmd.Execute()

	expectedExecutionEnvironmentState := testCase.expectedExecutionEnvironmentStateBuilder.Build()
	mockExecutionEnvironment.GetCurrentState().AssertEquals(t, expectedExecutionEnvironmentState, testCase.description)
}
