package cmd_test

import (
	"errors"
	"testing"

	"github.com/stevengt/mppm/cmd"
	"github.com/stevengt/mppm/config"
	"github.com/stevengt/mppm/config/configtest"
	"github.com/stevengt/mppm/util"
	"github.com/stretchr/testify/assert"

	"github.com/stevengt/mppm/util/utiltest"
)

var projectCmdHelpMessage string = "Provides utilities for managing a specific project.\n\nUsage:\n  mppm project [flags]\n  mppm project [command]\n\nAvailable Commands:\n  extract     Extracts all binary files of supported types into plain-text files, such as XML.\n  init        Initializes version control settings for a project using git and git-lfs.\n  restore     Restores all plain-text files of supported types to their original binary files.\n\nFlags:\n  -c, --commit-all         Equivalent to running 'mppm project extract; git add . -A; git commit -m '<commit message>'.\n  -h, --help               help for project\n  -p, --preview            Shows what files will be affected without actually making changes.\n  -u, --update-libraries   Updates the library versions in the project config file to match the\n                           current versions in the global config file.\n                           To see the global current versions, run 'mppm library --list'.\n\nUse \"mppm project [command] --help\" for more information about a command.\n"

func TestProjectCmd(t *testing.T) {

	testCases := []*ProjectCmdTestCase{

		&ProjectCmdTestCase{
			args:           []string{"project"},
			expectedOutput: projectCmdHelpMessage,
		},

		&ProjectCmdTestCase{
			args:           []string{"project", "invalid", "args"},
			expectedOutput: projectCmdHelpMessage,
		},

		&ProjectCmdTestCase{
			args: []string{"project", "--update-libraries"},
			projectConfigFile: utiltest.NewMockFile(
				configtest.ConfigWithAllValidInfoAndPreviousLibraryVersion.ConfigAsJson,
			),
			globalConfigFile: utiltest.NewMockFile(
				configtest.ConfigWithAllValidInfoAndMostRecentLibraryVersion.ConfigAsJson,
			),
			shouldUpdateLibraries: true,
		},

		&ProjectCmdTestCase{
			args: []string{"project", "--update-libraries"},
			projectConfigFile: utiltest.NewMockFile(
				configtest.ConfigWithValidVersionAndApplicationNameAndApplicationVersion.ConfigAsJson,
			),
			globalConfigFile: utiltest.NewMockFile(
				configtest.ConfigWithAllValidInfoAndMostRecentLibraryVersion.ConfigAsJson,
			),
			shouldUpdateLibraries: true,
		},

		&ProjectCmdTestCase{
			args: []string{"project", "--update-libraries"},
			executionEnvironmentBuilder: &utiltest.MockExecutionEnvironmentBuilder{
				MockFileSystemDelegaterBuilder: &utiltest.MockFileSystemDelegaterBuilder{
					UseDefaultOpenFileError: true,
				},
			},
			projectConfigFile: utiltest.NewMockFile(
				configtest.ConfigWithAllValidInfoAndPreviousLibraryVersion.ConfigAsJson,
			),
			globalConfigFile: utiltest.NewMockFile(
				configtest.ConfigWithAllValidInfoAndMostRecentLibraryVersion.ConfigAsJson,
			),
			shouldUpdateLibraries: true,
			expectedExitError:     errors.New("\nThere was a problem while opening the mppm config file.\nIf the file doesn't exist, try running 'mppm project init' first.\nThere was a problem opening the file.\n"),
		},

		&ProjectCmdTestCase{
			args: []string{"project", "--commit-all", "Made changes"},
			projectConfigFile: utiltest.NewMockFile(
				configtest.ConfigWithValidVersionAndApplicationNameAndApplicationVersion.ConfigAsJson,
			),
			expectedGitManagerInputHistories: map[string][][]string{
				".": [][]string{
					[]string{"add", ".", "-A"},
					[]string{"commit", "-m", "Made changes"},
				},
			},
			isCommitAllCommand: true,
		},

		&ProjectCmdTestCase{
			args: []string{"project", "--commit-all", "Made changes"},
			executionEnvironmentBuilder: &utiltest.MockExecutionEnvironmentBuilder{
				MockGitManagerCreatorBuilder: &utiltest.MockGitManagerCreatorBuilder{
					UseDefaultAddError: true,
				},
			},
			projectConfigFile: utiltest.NewMockFile(
				configtest.ConfigWithValidVersionAndApplicationNameAndApplicationVersion.ConfigAsJson,
			),
			expectedGitManagerInputHistories: map[string][][]string{
				".": [][]string{
					[]string{"add", ".", "-A"},
				},
			},
			isCommitAllCommand: true,
			expectedExitError:  utiltest.DefaultAddError,
		},

		&ProjectCmdTestCase{
			args: []string{"project", "--commit-all", "Made changes"},
			executionEnvironmentBuilder: &utiltest.MockExecutionEnvironmentBuilder{
				MockFileSystemDelegaterBuilder: &utiltest.MockFileSystemDelegaterBuilder{
					FileNamesAndContentsAsBytes: utiltest.GetTestFileNamesAndContents(),
				},
			},
			projectConfigFile: utiltest.NewMockFile(
				configtest.ConfigWithValidVersionAndApplicationNameAndApplicationVersion.ConfigAsJson,
			),
			expectedGitManagerInputHistories: map[string][][]string{
				".": [][]string{
					[]string{"add", ".", "-A"},
					[]string{"commit", "-m", "Made changes"},
				},
			},
			isCommitAllCommand: true,
		},
	}

	for _, testCase := range testCases {
		testCase.Run(t)
	}

}

type ProjectCmdTestCase struct {
	args                             []string
	executionEnvironmentBuilder      *utiltest.MockExecutionEnvironmentBuilder
	projectConfigFile                *utiltest.MockFile
	globalConfigFile                 *utiltest.MockFile
	shouldUpdateLibraries            bool
	isCommitAllCommand               bool
	expectedOutput                   string
	expectedGitManagerInputHistories map[string][][]string // Map of git repo file path -> git manager input history.
	expectedExitError                error
}

func (testCase *ProjectCmdTestCase) Run(t *testing.T) {

	executionEnvironment := utiltest.GetAndInitMockExecutionEnvironmentFromBuilderOrNil(testCase.executionEnvironmentBuilder)
	configtest.InitMockFileSystemDelegaterWithConfigFiles(
		executionEnvironment.MockFileSystemDelegater,
		testCase.projectConfigFile,
		testCase.globalConfigFile,
	)

	cmd.RootCmd.SetArgs(testCase.args)
	cmd.RootCmd.Execute()

	actualOutput := executionEnvironment.MockWritePrinter.GetOutputContentsAsString()
	assert.Equal(t, testCase.expectedOutput, actualOutput)

	for repoFilePath, expectedGitManagerInputHistory := range testCase.expectedGitManagerInputHistories {
		gitManager := executionEnvironment.MockGitManagerCreator.MockGitManagersIndexedByRepoPath[repoFilePath]
		assert.NotNil(t, gitManager)
		actualGitManagerInputHistory := gitManager.InputHistory
		assert.Exactly(t, expectedGitManagerInputHistory, actualGitManagerInputHistory)
	}

	if testCase.shouldUpdateLibraries {

		projectConfigFileOriginalContents := testCase.projectConfigFile.Contents
		projectConfigFileNewContents := executionEnvironment.MockFileSystemDelegater.Files[".mppm.json"].Contents

		globalConfigFileOriginalContents := testCase.globalConfigFile.Contents
		globalConfigFileNewContents := executionEnvironment.MockFileSystemDelegater.Files["/home/testuser/.mppm.json"].Contents

		if testCase.expectedExitError != nil {
			assert.Exactly(t, projectConfigFileOriginalContents, projectConfigFileNewContents)
			assert.Exactly(t, globalConfigFileOriginalContents, globalConfigFileNewContents)
		} else {
			projectConfig, err := config.NewMppmConfigInfoFromJson(projectConfigFileNewContents)
			assert.Nil(t, err)
			globalConfig, err := config.NewMppmConfigInfoFromJson(globalConfigFileNewContents)
			assert.Nil(t, err)
			assert.Exactly(t, globalConfig.Libraries, projectConfig.Libraries)
		}

	}

	if testCase.isCommitAllCommand && testCase.expectedExitError == nil {

		filePatternsConfig, _ := config.GetAllFilePatternsConfigFromProjectConfig()
		gzippedXmlFileExtensions := filePatternsConfig.GzippedXmlFileExtensions

		for _, fileExtension := range gzippedXmlFileExtensions {
			compressedFileNames, _ := util.GetAllFileNamesWithExtension(fileExtension)
			for _, compressedFileName := range compressedFileNames {
				uncompressedFileName := compressedFileName + ".xml"
				assert.True(t, executionEnvironment.MockFileSystemDelegater.DoesFileExist(uncompressedFileName))
			}
		}

	}

	assert.Exactly(t, testCase.expectedExitError, executionEnvironment.MockExiter.Error)
	if testCase.expectedExitError != nil {
		assert.True(t, executionEnvironment.MockExiter.WasExited)
	} else {
		assert.False(t, executionEnvironment.MockExiter.WasExited)
	}

}
