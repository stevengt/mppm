package util_test

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stevengt/mppm/util"
	"github.com/stevengt/mppm/util/utiltest"
)

func TestInit(t *testing.T) {

	mockShellCommandDelegater := utiltest.NewMockShellCommandDelegater(
		[]*utiltest.MockShellCommandOutput{
			&utiltest.MockShellCommandOutput{
				Stdout: "Initialized git repository.",
				Err:    nil,
			},
			&utiltest.MockShellCommandOutput{
				Stdout: "Something went wrong.",
				Err:    errors.New("There was a problem initializing the git repository."),
			},
		},
	)
	util.ShellProxy = mockShellCommandDelegater

	gitManager := util.NewGitManager(".")

	err := gitManager.Init()
	assert.Equal(t, mockShellCommandDelegater.InputHistory[0], "git -C . init")
	assert.Equal(t, mockShellCommandDelegater.OutputHistory[0].Stdout, "Initialized git repository.")
	assert.Nil(t, err)

	err = gitManager.Init()
	assert.Equal(t, mockShellCommandDelegater.InputHistory[1], "git -C . init")
	assert.Equal(t, mockShellCommandDelegater.OutputHistory[1].Stdout, "Something went wrong.")
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "There was a problem initializing the git repository.")

}

func TestAdd(t *testing.T) {

	mockShellCommandDelegater := utiltest.NewMockShellCommandDelegater(
		[]*utiltest.MockShellCommandOutput{
			&utiltest.MockShellCommandOutput{
				Stdout: "Added items to git repository.",
				Err:    nil,
			},
			&utiltest.MockShellCommandOutput{
				Stdout: "Added items to git repository.",
				Err:    nil,
			},
			&utiltest.MockShellCommandOutput{
				Stdout: "Something went wrong.",
				Err:    errors.New("There was a problem adding items to the git repository."),
			},
		},
	)
	util.ShellProxy = mockShellCommandDelegater

	gitManager := util.NewGitManager(".")

	err := gitManager.Add("-A", ".")
	assert.Equal(t, mockShellCommandDelegater.InputHistory[0], "git -C . add -A .")
	assert.Equal(t, mockShellCommandDelegater.OutputHistory[0].Stdout, "Added items to git repository.")
	assert.Nil(t, err)

	err = gitManager.Add("item1", "item2")
	assert.Equal(t, mockShellCommandDelegater.InputHistory[1], "git -C . add item1 item2")
	assert.Equal(t, mockShellCommandDelegater.OutputHistory[1].Stdout, "Added items to git repository.")
	assert.Nil(t, err)

	err = gitManager.Add("-A", ".")
	assert.Equal(t, mockShellCommandDelegater.InputHistory[2], "git -C . add -A .")
	assert.Equal(t, mockShellCommandDelegater.OutputHistory[2].Stdout, "Something went wrong.")
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "There was a problem adding items to the git repository.")

}

func TestCommit(t *testing.T) {

	mockShellCommandDelegater := utiltest.NewMockShellCommandDelegater(
		[]*utiltest.MockShellCommandOutput{
			&utiltest.MockShellCommandOutput{
				Stdout: "Committed items to git repository.",
				Err:    nil,
			},
			&utiltest.MockShellCommandOutput{
				Stdout: "Something went wrong.",
				Err:    errors.New("There was a problem committing items to the git repository."),
			},
		},
	)
	util.ShellProxy = mockShellCommandDelegater

	gitManager := util.NewGitManager(".")

	err := gitManager.Commit("-m", "commit message")
	assert.Equal(t, mockShellCommandDelegater.InputHistory[0], "git -C . commit -m commit message")
	assert.Equal(t, mockShellCommandDelegater.OutputHistory[0].Stdout, "Committed items to git repository.")
	assert.Nil(t, err)

	err = gitManager.Commit("-m", "commit message")
	assert.Equal(t, mockShellCommandDelegater.InputHistory[1], "git -C . commit -m commit message")
	assert.Equal(t, mockShellCommandDelegater.OutputHistory[1].Stdout, "Something went wrong.")
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "There was a problem committing items to the git repository.")

}

func TestCheckout(t *testing.T) {

	mockShellCommandDelegater := utiltest.NewMockShellCommandDelegater(
		[]*utiltest.MockShellCommandOutput{
			&utiltest.MockShellCommandOutput{
				Stdout: "Checked out git repository.",
				Err:    nil,
			},
			&utiltest.MockShellCommandOutput{
				Stdout: "Something went wrong.",
				Err:    errors.New("There was a problem checking out the git repository."),
			},
		},
	)
	util.ShellProxy = mockShellCommandDelegater

	gitManager := util.NewGitManager(".")

	err := gitManager.Checkout("master")
	assert.Equal(t, mockShellCommandDelegater.InputHistory[0], "git -C . checkout master")
	assert.Equal(t, mockShellCommandDelegater.OutputHistory[0].Stdout, "Checked out git repository.")
	assert.Nil(t, err)

	err = gitManager.Checkout("master")
	assert.Equal(t, mockShellCommandDelegater.InputHistory[1], "git -C . checkout master")
	assert.Equal(t, mockShellCommandDelegater.OutputHistory[1].Stdout, "Something went wrong.")
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "There was a problem checking out the git repository.")

}

func TestRevParse(t *testing.T) {

	mockShellCommandDelegater := utiltest.NewMockShellCommandDelegater(
		[]*utiltest.MockShellCommandOutput{
			&utiltest.MockShellCommandOutput{
				Stdout: "",
				Err:    nil,
			},
			&utiltest.MockShellCommandOutput{
				Stdout: "git commit id = 012345",
				Err:    nil,
			},
			&utiltest.MockShellCommandOutput{
				Stdout: "",
				Err:    errors.New("Not a git repository."),
			},
		},
	)
	util.ShellProxy = mockShellCommandDelegater

	gitManager := util.NewGitManager(".")

	stdout, err := gitManager.RevParse()
	assert.Equal(t, mockShellCommandDelegater.InputHistory[0], "git -C . rev-parse")
	assert.Equal(t, stdout, "")
	assert.Nil(t, err)

	stdout, err = gitManager.RevParse("HEAD")
	assert.Equal(t, mockShellCommandDelegater.InputHistory[1], "git -C . rev-parse HEAD")
	assert.Equal(t, stdout, "git commit id = 012345")
	assert.Nil(t, err)

	stdout, err = gitManager.RevParse()
	assert.Equal(t, mockShellCommandDelegater.InputHistory[2], "git -C . rev-parse")
	assert.Equal(t, stdout, "")
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "Not a git repository.")

}

func TestLfsInstall(t *testing.T) {

	mockShellCommandDelegater := utiltest.NewMockShellCommandDelegater(
		[]*utiltest.MockShellCommandOutput{
			&utiltest.MockShellCommandOutput{
				Stdout: "git lfs is now set up.",
				Err:    nil,
			},
			&utiltest.MockShellCommandOutput{
				Stdout: "",
				Err:    errors.New("There was a problem while setting up git lfs."),
			},
		},
	)
	util.ShellProxy = mockShellCommandDelegater

	gitManager := util.NewGitManager(".")

	err := gitManager.LfsInstall()
	assert.Equal(t, mockShellCommandDelegater.InputHistory[0], "git -C . lfs install")
	assert.Equal(t, mockShellCommandDelegater.OutputHistory[0].Stdout, "git lfs is now set up.")
	assert.Nil(t, err)

	err = gitManager.LfsInstall()
	assert.Equal(t, mockShellCommandDelegater.InputHistory[1], "git -C . lfs install")
	assert.Equal(t, mockShellCommandDelegater.OutputHistory[1].Stdout, "")
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "There was a problem while setting up git lfs.")

}

func TestLfsTrack(t *testing.T) {

	mockShellCommandDelegater := utiltest.NewMockShellCommandDelegater(
		[]*utiltest.MockShellCommandOutput{
			&utiltest.MockShellCommandOutput{
				Stdout: "Tracking files with git lfs.",
				Err:    nil,
			},
			&utiltest.MockShellCommandOutput{
				Stdout: "",
				Err:    errors.New("There was a problem trying to track files with git lfs."),
			},
		},
	)
	util.ShellProxy = mockShellCommandDelegater

	gitManager := util.NewGitManager(".")

	err := gitManager.LfsTrack("*.txt")
	assert.Equal(t, mockShellCommandDelegater.InputHistory[0], "git -C . lfs track *.txt")
	assert.Equal(t, mockShellCommandDelegater.OutputHistory[0].Stdout, "Tracking files with git lfs.")
	assert.Nil(t, err)

	err = gitManager.LfsTrack("*.txt", "big-file.bin")
	assert.Equal(t, mockShellCommandDelegater.InputHistory[1], "git -C . lfs track *.txt big-file.bin")
	assert.Equal(t, mockShellCommandDelegater.OutputHistory[1].Stdout, "")
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "There was a problem trying to track files with git lfs.")

}

func TestAddAllAndCommit(t *testing.T) {

	mockShellCommandDelegater := utiltest.NewMockShellCommandDelegater(
		[]*utiltest.MockShellCommandOutput{
			&utiltest.MockShellCommandOutput{
				Stdout: "Added items to git repository.",
				Err:    nil,
			},
			&utiltest.MockShellCommandOutput{
				Stdout: "Committed items to git repository.",
				Err:    nil,
			},

			&utiltest.MockShellCommandOutput{
				Stdout: "",
				Err:    errors.New("There was a problem adding items to the git repository."),
			},

			&utiltest.MockShellCommandOutput{
				Stdout: "Added items to git repository.",
				Err:    nil,
			},
			&utiltest.MockShellCommandOutput{
				Stdout: "",
				Err:    errors.New("There was a problem committing items to the git repository."),
			},
		},
	)
	util.ShellProxy = mockShellCommandDelegater

	gitManager := util.NewGitManager(".")

	err := gitManager.AddAllAndCommit("commit message")
	assert.Equal(t, mockShellCommandDelegater.InputHistory[0], "git -C . add -A .")
	assert.Equal(t, mockShellCommandDelegater.OutputHistory[0].Stdout, "Added items to git repository.")
	assert.Equal(t, mockShellCommandDelegater.InputHistory[1], "git -C . commit -m commit message")
	assert.Equal(t, mockShellCommandDelegater.OutputHistory[1].Stdout, "Committed items to git repository.")
	assert.Nil(t, err)

	err = gitManager.AddAllAndCommit("commit message")
	assert.Equal(t, mockShellCommandDelegater.InputHistory[2], "git -C . add -A .")
	assert.Equal(t, mockShellCommandDelegater.OutputHistory[2].Stdout, "")
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "There was a problem adding items to the git repository.")

	err = gitManager.AddAllAndCommit("commit message")
	assert.Equal(t, mockShellCommandDelegater.InputHistory[3], "git -C . add -A .")
	assert.Equal(t, mockShellCommandDelegater.OutputHistory[3].Stdout, "Added items to git repository.")
	assert.Equal(t, mockShellCommandDelegater.InputHistory[4], "git -C . commit -m commit message")
	assert.Equal(t, mockShellCommandDelegater.OutputHistory[4].Stdout, "")
	assert.NotNil(t, err)
	assert.Equal(t, err.Error(), "There was a problem committing items to the git repository.")

}
