package util_test

import (
	"testing"

	"github.com/stevengt/mppm/util"
	"github.com/stevengt/mppm/util/utiltest"
	"github.com/stretchr/testify/assert"
)

type gitManagerMethodType int

const (
	gitInit gitManagerMethodType = iota
	add
	commit
	checkout
	revParse
	lfsInstall
	lfsTrack
	addAllAndCommit
)

func TestInit(t *testing.T) {

	testCases := []*GitManagerTestCase{

		&GitManagerTestCase{
			description:            "Test that the correct 'git init' shell command is invoked.",
			gitManagerRepoFilePath: ".",
			methodType:             gitInit,
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockShellCommandDelegaterBuilder(
					utiltest.NewMockShellCommandDelegaterBuilder().
						SetOutputSequence(
							&utiltest.MockShellCommandOutput{
								Stdout: "Initialized git repository.",
							},
						),
				),

			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetShellCommandDelegaterInputHistory(
					"git -C . init",
				).
				SetShellCommandDelegaterOutputHistory(
					&utiltest.MockShellCommandOutput{
						Stdout: "Initialized git repository.",
					},
				).
				SetWritePrinterOutputContents(
					[]byte("Initialized git repository.\n"),
				),
		},

		&GitManagerTestCase{
			description:            "Test that any error from running 'git init' is correctly raised.",
			gitManagerRepoFilePath: ".",
			methodType:             gitInit,
			expectedError:          utiltest.DefaultInitError,
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockShellCommandDelegaterBuilder(
					utiltest.NewMockShellCommandDelegaterBuilder().
						SetOutputSequence(
							&utiltest.MockShellCommandOutput{
								Err: utiltest.DefaultInitError,
							},
						),
				),

			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetShellCommandDelegaterInputHistory(
					"git -C . init",
				).
				SetShellCommandDelegaterOutputHistory(
					&utiltest.MockShellCommandOutput{
						Err: utiltest.DefaultInitError,
					},
				),
		},
	}

	for _, testCase := range testCases {
		testCase.Run(t)
	}

}

func TestAdd(t *testing.T) {

	testCases := []*GitManagerTestCase{

		&GitManagerTestCase{
			description:            "Test that the correct 'git add' shell command is invoked.",
			gitManagerRepoFilePath: ".",
			methodType:             add,
			gitManagerMethodArgs:   []string{".", "-A"},
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockShellCommandDelegaterBuilder(
					utiltest.NewMockShellCommandDelegaterBuilder().
						SetOutputSequence(
							&utiltest.MockShellCommandOutput{
								Stdout: "Added items to git repository.",
							},
						),
				),

			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetShellCommandDelegaterInputHistory(
					"git -C . add . -A",
				).
				SetShellCommandDelegaterOutputHistory(
					&utiltest.MockShellCommandOutput{
						Stdout: "Added items to git repository.",
					},
				).
				SetWritePrinterOutputContents(
					[]byte("Added items to git repository.\n"),
				),
		},

		&GitManagerTestCase{
			description:            "Test that any error from running 'git add' is correctly raised.",
			gitManagerRepoFilePath: ".",
			methodType:             add,
			gitManagerMethodArgs:   []string{".", "-A"},
			expectedError:          utiltest.DefaultAddError,
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockShellCommandDelegaterBuilder(
					utiltest.NewMockShellCommandDelegaterBuilder().
						SetOutputSequence(
							&utiltest.MockShellCommandOutput{
								Err: utiltest.DefaultAddError,
							},
						),
				),

			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetShellCommandDelegaterInputHistory(
					"git -C . add . -A",
				).
				SetShellCommandDelegaterOutputHistory(
					&utiltest.MockShellCommandOutput{
						Err: utiltest.DefaultAddError,
					},
				),
		},
	}

	for _, testCase := range testCases {
		testCase.Run(t)
	}

}

func TestCommit(t *testing.T) {

	testCases := []*GitManagerTestCase{

		&GitManagerTestCase{
			description:            "Test that the correct 'git commit' shell command is invoked.",
			gitManagerRepoFilePath: ".",
			methodType:             commit,
			gitManagerMethodArgs:   []string{"-m", "fake commit message"},
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockShellCommandDelegaterBuilder(
					utiltest.NewMockShellCommandDelegaterBuilder().
						SetOutputSequence(
							&utiltest.MockShellCommandOutput{
								Stdout: "Committed items to git repository.",
							},
						),
				),

			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetShellCommandDelegaterInputHistory(
					"git -C . commit -m fake commit message",
				).
				SetShellCommandDelegaterOutputHistory(
					&utiltest.MockShellCommandOutput{
						Stdout: "Committed items to git repository.",
					},
				).
				SetWritePrinterOutputContents(
					[]byte("Committed items to git repository.\n"),
				),
		},

		&GitManagerTestCase{
			description:            "Test that any error from running 'git commit' is correctly raised.",
			gitManagerRepoFilePath: ".",
			methodType:             commit,
			gitManagerMethodArgs:   []string{"-m", "fake commit message"},
			expectedError:          utiltest.DefaultCommitError,
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockShellCommandDelegaterBuilder(
					utiltest.NewMockShellCommandDelegaterBuilder().
						SetOutputSequence(
							&utiltest.MockShellCommandOutput{
								Err: utiltest.DefaultCommitError,
							},
						),
				),

			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetShellCommandDelegaterInputHistory(
					"git -C . commit -m fake commit message",
				).
				SetShellCommandDelegaterOutputHistory(
					&utiltest.MockShellCommandOutput{
						Err: utiltest.DefaultCommitError,
					},
				),
		},
	}

	for _, testCase := range testCases {
		testCase.Run(t)
	}

}

func TestCheckout(t *testing.T) {

	testCases := []*GitManagerTestCase{

		&GitManagerTestCase{
			description:            "Test that the correct 'git checkout' shell command is invoked.",
			gitManagerRepoFilePath: ".",
			methodType:             checkout,
			gitManagerMethodArgs:   []string{"master"},
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockShellCommandDelegaterBuilder(
					utiltest.NewMockShellCommandDelegaterBuilder().
						SetOutputSequence(
							&utiltest.MockShellCommandOutput{
								Stdout: "Checked out master branch.",
							},
						),
				),

			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetShellCommandDelegaterInputHistory(
					"git -C . checkout master",
				).
				SetShellCommandDelegaterOutputHistory(
					&utiltest.MockShellCommandOutput{
						Stdout: "Checked out master branch.",
					},
				).
				SetWritePrinterOutputContents(
					[]byte("Checked out master branch.\n"),
				),
		},

		&GitManagerTestCase{
			description:            "Test that any error from running 'git checkout' is correctly raised.",
			gitManagerRepoFilePath: ".",
			methodType:             checkout,
			gitManagerMethodArgs:   []string{"master"},
			expectedError:          utiltest.DefaultCheckoutError,
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockShellCommandDelegaterBuilder(
					utiltest.NewMockShellCommandDelegaterBuilder().
						SetOutputSequence(
							&utiltest.MockShellCommandOutput{
								Err: utiltest.DefaultCheckoutError,
							},
						),
				),

			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetShellCommandDelegaterInputHistory(
					"git -C . checkout master",
				).
				SetShellCommandDelegaterOutputHistory(
					&utiltest.MockShellCommandOutput{
						Err: utiltest.DefaultCheckoutError,
					},
				),
		},
	}

	for _, testCase := range testCases {
		testCase.Run(t)
	}

}

func TestRevParse(t *testing.T) {

	testCases := []*GitManagerTestCase{

		&GitManagerTestCase{
			description:            "Test that the correct 'git rev-parse' shell command is invoked.",
			gitManagerRepoFilePath: ".",
			methodType:             revParse,
			gitManagerMethodArgs:   []string{"HEAD"},
			expectedStdout:         "git commit id = 012345",
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockShellCommandDelegaterBuilder(
					utiltest.NewMockShellCommandDelegaterBuilder().
						SetOutputSequence(
							&utiltest.MockShellCommandOutput{
								Stdout: "git commit id = 012345",
							},
						),
				),

			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetShellCommandDelegaterInputHistory(
					"git -C . rev-parse HEAD",
				).
				SetShellCommandDelegaterOutputHistory(
					&utiltest.MockShellCommandOutput{
						Stdout: "git commit id = 012345",
					},
				),
		},

		&GitManagerTestCase{
			description:            "Test that any error from running 'git rev-parse' is correctly raised.",
			gitManagerRepoFilePath: ".",
			methodType:             revParse,
			gitManagerMethodArgs:   []string{"HEAD"},
			expectedError:          utiltest.DefaultRevParseError,
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockShellCommandDelegaterBuilder(
					utiltest.NewMockShellCommandDelegaterBuilder().
						SetOutputSequence(
							&utiltest.MockShellCommandOutput{
								Err: utiltest.DefaultRevParseError,
							},
						),
				),

			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetShellCommandDelegaterInputHistory(
					"git -C . rev-parse HEAD",
				).
				SetShellCommandDelegaterOutputHistory(
					&utiltest.MockShellCommandOutput{
						Err: utiltest.DefaultRevParseError,
					},
				),
		},
	}

	for _, testCase := range testCases {
		testCase.Run(t)
	}

}

func TestLfsInstall(t *testing.T) {

	testCases := []*GitManagerTestCase{

		&GitManagerTestCase{
			description:            "Test that the correct 'git lfs install' shell command is invoked.",
			gitManagerRepoFilePath: ".",
			methodType:             lfsInstall,
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockShellCommandDelegaterBuilder(
					utiltest.NewMockShellCommandDelegaterBuilder().
						SetOutputSequence(
							&utiltest.MockShellCommandOutput{
								Stdout: "git lfs is now set up.",
							},
						),
				),

			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetShellCommandDelegaterInputHistory(
					"git -C . lfs install",
				).
				SetShellCommandDelegaterOutputHistory(
					&utiltest.MockShellCommandOutput{
						Stdout: "git lfs is now set up.",
					},
				).
				SetWritePrinterOutputContents(
					[]byte("git lfs is now set up.\n"),
				),
		},

		&GitManagerTestCase{
			description:            "Test that any error from running 'git lfs install' is correctly raised.",
			gitManagerRepoFilePath: ".",
			methodType:             lfsInstall,
			expectedError:          utiltest.DefaultLfsInstallError,
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockShellCommandDelegaterBuilder(
					utiltest.NewMockShellCommandDelegaterBuilder().
						SetOutputSequence(
							&utiltest.MockShellCommandOutput{
								Err: utiltest.DefaultLfsInstallError,
							},
						),
				),

			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetShellCommandDelegaterInputHistory(
					"git -C . lfs install",
				).
				SetShellCommandDelegaterOutputHistory(
					&utiltest.MockShellCommandOutput{
						Err: utiltest.DefaultLfsInstallError,
					},
				),
		},
	}

	for _, testCase := range testCases {
		testCase.Run(t)
	}

}

func TestLfsTrack(t *testing.T) {

	testCases := []*GitManagerTestCase{

		&GitManagerTestCase{
			description:            "Test that the correct 'git lfs track' shell command is invoked.",
			gitManagerRepoFilePath: ".",
			methodType:             lfsTrack,
			gitManagerMethodArgs:   []string{"*.txt", "*.bin"},
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockShellCommandDelegaterBuilder(
					utiltest.NewMockShellCommandDelegaterBuilder().
						SetOutputSequence(
							&utiltest.MockShellCommandOutput{
								Stdout: "Added items to track with git lfs.",
							},
						),
				),

			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetShellCommandDelegaterInputHistory(
					"git -C . lfs track *.txt *.bin",
				).
				SetShellCommandDelegaterOutputHistory(
					&utiltest.MockShellCommandOutput{
						Stdout: "Added items to track with git lfs.",
					},
				).
				SetWritePrinterOutputContents(
					[]byte("Added items to track with git lfs.\n"),
				),
		},

		&GitManagerTestCase{
			description:            "Test that any error from running 'git lfs track' is correctly raised.",
			gitManagerRepoFilePath: ".",
			methodType:             lfsTrack,
			gitManagerMethodArgs:   []string{"*.txt", "*.bin"},
			expectedError:          utiltest.DefaultLfsTrackError,
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockShellCommandDelegaterBuilder(
					utiltest.NewMockShellCommandDelegaterBuilder().
						SetOutputSequence(
							&utiltest.MockShellCommandOutput{
								Err: utiltest.DefaultLfsTrackError,
							},
						),
				),

			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetShellCommandDelegaterInputHistory(
					"git -C . lfs track *.txt *.bin",
				).
				SetShellCommandDelegaterOutputHistory(
					&utiltest.MockShellCommandOutput{
						Err: utiltest.DefaultLfsTrackError,
					},
				),
		},
	}

	for _, testCase := range testCases {
		testCase.Run(t)
	}

}

type GitManagerTestCase struct {
	description                              string
	gitManagerRepoFilePath                   string
	methodType                               gitManagerMethodType
	gitManagerMethodArgs                     []string
	expectedStdout                           string
	expectedError                            error
	mockExecutionEnvironmentBuilder          *utiltest.MockExecutionEnvironmentBuilder
	expectedExecutionEnvironmentStateBuilder *utiltest.MockExecutionEnvironmentStateBuilder
}

func (testCase *GitManagerTestCase) Run(t *testing.T) {

	// Set up an execution environment without a MockGitManagerCreator. Instead, use the
	// default GitManagerCreator that delegates shell commands to a MockShellCommandDelegater.
	mockExecutionEnvironment := testCase.mockExecutionEnvironmentBuilder.Build()
	mockExecutionEnvironment.MockGitManagerCreator = nil
	mockExecutionEnvironment.Init()

	gitManager := util.NewGitManager(testCase.gitManagerRepoFilePath)

	var actualStdout string
	var actualError error

	switch testCase.methodType {
	case gitInit:
		actualError = gitManager.Init()
	case add:
		actualError = gitManager.Add(testCase.gitManagerMethodArgs...)
	case commit:
		actualError = gitManager.Commit(testCase.gitManagerMethodArgs...)
	case checkout:
		actualError = gitManager.Checkout(testCase.gitManagerMethodArgs...)
	case revParse:
		actualStdout, actualError = gitManager.RevParse(testCase.gitManagerMethodArgs...)
	case lfsInstall:
		actualError = gitManager.LfsInstall()
	case lfsTrack:
		actualError = gitManager.LfsTrack(testCase.gitManagerMethodArgs...)
	case addAllAndCommit:
		actualError = gitManager.AddAllAndCommit(testCase.gitManagerMethodArgs[0])
	}

	assert.Exactly(t, testCase.expectedStdout, actualStdout)
	assert.Exactly(t, testCase.expectedError, actualError)

	expectedExecutionEnvironmentState := testCase.expectedExecutionEnvironmentStateBuilder.Build()
	mockExecutionEnvironment.GetCurrentState().AssertEquals(t, expectedExecutionEnvironmentState, testCase.description)

}
