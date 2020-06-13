package util_test

import (
	"errors"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stevengt/mppm/util"
	"github.com/stevengt/mppm/util/utiltest"
)

func TestCopyFile(t *testing.T) {

	testCases := []*CopyFileTestCase{

		&CopyFileTestCase{
			description:    "Test if contents of one file are correctly overwritten with the contents of another file.",
			sourceFileName: utiltest.GetPlainTextFileBuilder().FilePath,
			targetFileName: utiltest.GetBinaryContaining0xdeadbeefFileBuilder().FilePath,
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockFileSystemDelegaterBuilder(
					utiltest.NewMockFileSystemDelegaterBuilder().
						SetMockFileBuilders(
							utiltest.GetPlainTextFileBuilder(),
							utiltest.GetBinaryContaining0xdeadbeefFileBuilder(),
						),
				),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetMockFileBuilders(
					utiltest.GetPlainTextFileBuilder().
						SetWasClosed(true),
					utiltest.GetBinaryContaining0xdeadbeefFileBuilder().
						SetContentsFromBytes(utiltest.GetPlainTextFileBuilder().Contents).
						SetWasClosed(true),
				),
		},

		&CopyFileTestCase{
			description:    "Test if error is correctly raised when unable to open a file.",
			sourceFileName: utiltest.GetPlainTextFileBuilder().FilePath,
			targetFileName: utiltest.GetBinaryContaining0xdeadbeefFileBuilder().FilePath,
			expectedError:  utiltest.DefaultOpenFileError,
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockFileSystemDelegaterBuilder(
					utiltest.NewMockFileSystemDelegaterBuilder().
						SetMockFileBuilders(
							utiltest.GetPlainTextFileBuilder(),
							utiltest.GetBinaryContaining0xdeadbeefFileBuilder(),
						).
						SetUseDefaultOpenFileError(true),
				),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetMockFileBuilders(
					utiltest.GetPlainTextFileBuilder().
						SetWasClosed(false),
					utiltest.GetBinaryContaining0xdeadbeefFileBuilder().
						SetWasClosed(false),
				),
		},

		&CopyFileTestCase{
			description:    "Test if error is correctly raised when unable to recreate an existing file.",
			sourceFileName: utiltest.GetPlainTextFileBuilder().FilePath,
			targetFileName: utiltest.GetBinaryContaining0xdeadbeefFileBuilder().FilePath,
			expectedError:  utiltest.DefaultCreateFileError,
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockFileSystemDelegaterBuilder(
					utiltest.NewMockFileSystemDelegaterBuilder().
						SetMockFileBuilders(
							utiltest.GetPlainTextFileBuilder(),
							utiltest.GetBinaryContaining0xdeadbeefFileBuilder(),
						).
						SetUseDefaultCreateFileError(true),
				),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetMockFileBuilders(
					utiltest.GetPlainTextFileBuilder().
						SetWasClosed(true),
					utiltest.GetBinaryContaining0xdeadbeefFileBuilder().
						SetWasClosed(false),
				),
		},

		&CopyFileTestCase{
			description:    "Test if new file is created with the correct contents.",
			sourceFileName: utiltest.GetPlainTextFileBuilder().FilePath,
			targetFileName: "new-file",
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockFileSystemDelegaterBuilder(
					utiltest.NewMockFileSystemDelegaterBuilder().
						SetMockFileBuilders(
							utiltest.GetPlainTextFileBuilder(),
						),
				),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetMockFileBuilders(
					utiltest.GetPlainTextFileBuilder().
						SetWasClosed(true),
					utiltest.NewMockFileBuilder().
						SetFilePath("new-file").
						SetContentsFromBytes(utiltest.GetPlainTextFileBuilder().Contents).
						SetWasClosed(true),
				),
		},

		&CopyFileTestCase{
			description:    "Test if error is correctly raised when unable to create a new file.",
			sourceFileName: utiltest.GetPlainTextFileBuilder().FilePath,
			targetFileName: "new-file",
			expectedError:  utiltest.DefaultCreateFileError,
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockFileSystemDelegaterBuilder(
					utiltest.NewMockFileSystemDelegaterBuilder().
						SetMockFileBuilders(
							utiltest.GetPlainTextFileBuilder(),
						).
						SetUseDefaultCreateFileError(true),
				),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetMockFileBuilders(
					utiltest.GetPlainTextFileBuilder().
						SetWasClosed(true),
				),
		},

		&CopyFileTestCase{
			description:                              "Test if error is correctly raised when attempting to copy a file that does not exist.",
			sourceFileName:                           "does-not-exist",
			targetFileName:                           "new-file",
			expectedError:                            errors.New("Unable to open file does-not-exist"),
			mockExecutionEnvironmentBuilder:          utiltest.NewMockExecutionEnvironmentBuilder(),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder(),
		},

		&CopyFileTestCase{
			description:    "Test if empty file is correctly copied.",
			sourceFileName: utiltest.GetEmptyFileBuilder().FilePath,
			targetFileName: "new-file",
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockFileSystemDelegaterBuilder(
					utiltest.NewMockFileSystemDelegaterBuilder().
						SetMockFileBuilders(
							utiltest.GetEmptyFileBuilder(),
						),
				),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetMockFileBuilders(
					utiltest.GetEmptyFileBuilder().
						SetWasClosed(true),
					utiltest.GetEmptyFileBuilder().
						SetFilePath("new-file").
						SetWasClosed(true),
				),
		},
	}

	for _, testCase := range testCases {
		testCase.Run(t)
	}

}

func TestGzipFile(t *testing.T) {

	testCases := []*GzipFileTestCase{

		&GzipFileTestCase{
			description: "Test if a new file is created with correctly compressed contents.",
			fileName:    utiltest.GetPlainTextFileBuilder().FilePath,
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockFileSystemDelegaterBuilder(
					utiltest.NewMockFileSystemDelegaterBuilder().
						SetMockFileBuilders(
							utiltest.GetPlainTextFileBuilder(),
						),
				),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetMockFileBuilders(
					utiltest.GetPlainTextFileBuilder().
						SetWasClosed(true),
					utiltest.GetGzippedPlainTextFileBuilder().
						SetWasClosed(true),
				),
		},

		&GzipFileTestCase{
			description: "Test if a file is overwritten with correctly compressed contents.",
			fileName:    utiltest.GetPlainTextFileBuilder().FilePath,
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockFileSystemDelegaterBuilder(
					utiltest.NewMockFileSystemDelegaterBuilder().
						SetMockFileBuilders(
							utiltest.GetPlainTextFileBuilder(),
							utiltest.GetGzippedPlainTextFileBuilder().
								SetContentsFromString("fake orgininal gzipped contents."),
						),
				),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetMockFileBuilders(
					utiltest.GetPlainTextFileBuilder().
						SetWasClosed(true),
					utiltest.GetGzippedPlainTextFileBuilder().
						SetWasClosed(true),
				),
		},

		&GzipFileTestCase{
			description:   "Test that no files are changed if there is an os.Open error.",
			fileName:      utiltest.GetPlainTextFileBuilder().FilePath,
			expectedError: utiltest.DefaultOpenFileError,
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockFileSystemDelegaterBuilder(
					utiltest.NewMockFileSystemDelegaterBuilder().
						SetMockFileBuilders(
							utiltest.GetPlainTextFileBuilder(),
						).
						SetUseDefaultOpenFileError(true),
				),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetMockFileBuilders(
					utiltest.GetPlainTextFileBuilder(),
				),
		},

		&GzipFileTestCase{
			description:   "Test that no files are changed if there is an os.Create error, and the compresed file does not already exist.",
			fileName:      utiltest.GetPlainTextFileBuilder().FilePath,
			expectedError: utiltest.DefaultCreateFileError,
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockFileSystemDelegaterBuilder(
					utiltest.NewMockFileSystemDelegaterBuilder().
						SetMockFileBuilders(
							utiltest.GetPlainTextFileBuilder(),
						).
						SetUseDefaultCreateFileError(true),
				),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetMockFileBuilders(
					utiltest.GetPlainTextFileBuilder().
						SetWasClosed(true),
				),
		},

		&GzipFileTestCase{
			description:   "Test that no files are changed if there is an os.Create error, and the compresed file already exists.",
			fileName:      utiltest.GetPlainTextFileBuilder().FilePath,
			expectedError: utiltest.DefaultCreateFileError,
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockFileSystemDelegaterBuilder(
					utiltest.NewMockFileSystemDelegaterBuilder().
						SetMockFileBuilders(
							utiltest.GetPlainTextFileBuilder(),
							utiltest.GetGzippedPlainTextFileBuilder(),
						).
						SetUseDefaultCreateFileError(true),
				),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetMockFileBuilders(
					utiltest.GetPlainTextFileBuilder().
						SetWasClosed(true),
					utiltest.GetGzippedPlainTextFileBuilder(),
				),
		},

		&GzipFileTestCase{
			description:                              "Test if error is properly raised if uncompressed file does not exist.",
			fileName:                                 "does-not-exist",
			expectedError:                            errors.New("Unable to open file does-not-exist"),
			mockExecutionEnvironmentBuilder:          utiltest.NewMockExecutionEnvironmentBuilder(),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder(),
		},
	}

	for _, testCase := range testCases {
		testCase.Run(t)
	}

}

func TestGunzipFile(t *testing.T) {

	testCases := []*GunzipFileTestCase{

		&GunzipFileTestCase{
			description: "Test if a new file is created with correctly uncompressed contents, and the compressed file is removed.",
			fileName:    utiltest.GetGzippedPlainTextFileBuilder().FilePath,
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockFileSystemDelegaterBuilder(
					utiltest.NewMockFileSystemDelegaterBuilder().
						SetMockFileBuilders(
							utiltest.GetGzippedPlainTextFileBuilder(),
						),
				),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetMockFileBuilders(
					utiltest.GetPlainTextFileBuilder().
						SetWasClosed(true),
				),
		},

		&GunzipFileTestCase{
			description: "Test if a file is overwritten with correctly uncompressed contents, and the compressed file is removed.",
			fileName:    utiltest.GetGzippedPlainTextFileBuilder().FilePath,
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockFileSystemDelegaterBuilder(
					utiltest.NewMockFileSystemDelegaterBuilder().
						SetMockFileBuilders(
							utiltest.GetPlainTextFileBuilder().
								SetContentsFromString("fake orgininal gzipped contents."),
							utiltest.GetGzippedPlainTextFileBuilder(),
						),
				),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetMockFileBuilders(
					utiltest.GetPlainTextFileBuilder().
						SetWasClosed(true),
				),
		},

		&GunzipFileTestCase{
			description:   "Test that no files are changed if there is an os.Open error.",
			fileName:      utiltest.GetGzippedPlainTextFileBuilder().FilePath,
			expectedError: utiltest.DefaultOpenFileError,
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockFileSystemDelegaterBuilder(
					utiltest.NewMockFileSystemDelegaterBuilder().
						SetMockFileBuilders(
							utiltest.GetGzippedPlainTextFileBuilder(),
						).
						SetUseDefaultOpenFileError(true),
				),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetMockFileBuilders(
					utiltest.GetGzippedPlainTextFileBuilder(),
				),
		},

		&GunzipFileTestCase{
			description:   "Test that no files are changed if there is an os.Create error, and the uncompresed file does not already exist.",
			fileName:      utiltest.GetGzippedPlainTextFileBuilder().FilePath,
			expectedError: utiltest.DefaultCreateFileError,
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockFileSystemDelegaterBuilder(
					utiltest.NewMockFileSystemDelegaterBuilder().
						SetMockFileBuilders(
							utiltest.GetGzippedPlainTextFileBuilder(),
						).
						SetUseDefaultCreateFileError(true),
				),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetMockFileBuilders(
					utiltest.GetGzippedPlainTextFileBuilder().
						SetWasClosed(true),
				),
		},

		&GunzipFileTestCase{
			description:   "Test that no files are changed if there is an os.Create error, and the uncompresed file already exists.",
			fileName:      utiltest.GetGzippedPlainTextFileBuilder().FilePath,
			expectedError: utiltest.DefaultCreateFileError,
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockFileSystemDelegaterBuilder(
					utiltest.NewMockFileSystemDelegaterBuilder().
						SetMockFileBuilders(
							utiltest.GetPlainTextFileBuilder(),
							utiltest.GetGzippedPlainTextFileBuilder(),
						).
						SetUseDefaultCreateFileError(true),
				),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder().
				SetMockFileBuilders(
					utiltest.GetGzippedPlainTextFileBuilder().
						SetWasClosed(true),
					utiltest.GetPlainTextFileBuilder(),
				),
		},

		&GunzipFileTestCase{
			description:                              "Test if error is properly raised if compressed file does not exist.",
			fileName:                                 "does-not-exist",
			expectedError:                            errors.New("Unable to open file does-not-exist"),
			mockExecutionEnvironmentBuilder:          utiltest.NewMockExecutionEnvironmentBuilder(),
			expectedExecutionEnvironmentStateBuilder: utiltest.NewMockExecutionEnvironmentStateBuilder(),
		},
	}

	for _, testCase := range testCases {
		testCase.Run(t)
	}

}

func TestGetAllFileNamesWithExtension(t *testing.T) {

	testCases := []*GetAllFileNamesWithExtensionTestCase{

		&GetAllFileNamesWithExtensionTestCase{
			description:       "Test that only file-names with the correct file extension are returned.",
			fileExtension:     "txt",
			expectedFileNames: []string{"file1.txt", "/path/to/file2.txt", "file3.txt"},
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockFileSystemDelegaterBuilder(
					utiltest.NewMockFileSystemDelegaterBuilder().
						SetMockFileBuilders(
							utiltest.GetPlainTextFileBuilder().
								SetFilePath("file1.txt"),
							utiltest.GetPlainTextFileBuilder().
								SetFilePath("/path/to/file2.txt"),
							utiltest.GetPlainTextFileBuilder().
								SetFilePath("file3.txt"),
							utiltest.GetGzippedPlainTextFileBuilder(),
							utiltest.GetBinaryContaining0xdeadbeefFileBuilder(),
						),
				),
		},

		&GetAllFileNamesWithExtensionTestCase{
			description:       "Test that errors from filepath.Walk() are properly raised.",
			fileExtension:     "txt",
			expectedFileNames: nil,
			expectedError:     utiltest.DefaultWalkFilePathError,
			mockExecutionEnvironmentBuilder: utiltest.NewMockExecutionEnvironmentBuilder().
				SetMockFileSystemDelegaterBuilder(
					utiltest.NewMockFileSystemDelegaterBuilder().
						SetMockFileBuilders(
							utiltest.GetPlainTextFileBuilder().
								SetFilePath("file1.txt"),
							utiltest.GetPlainTextFileBuilder().
								SetFilePath("/path/to/file2.txt"),
							utiltest.GetPlainTextFileBuilder().
								SetFilePath("file3.txt"),
							utiltest.GetGzippedPlainTextFileBuilder(),
							utiltest.GetBinaryContaining0xdeadbeefFileBuilder(),
						).
						SetUseDefaultWalkFilePathError(true),
				),
		},
	}

	for _, testCase := range testCases {
		testCase.Run(t)
	}

}

// ------------------------------------------------------------------------------

type CopyFileTestCase struct {
	description                              string
	sourceFileName                           string
	targetFileName                           string
	expectedError                            error
	mockExecutionEnvironmentBuilder          *utiltest.MockExecutionEnvironmentBuilder
	expectedExecutionEnvironmentStateBuilder *utiltest.MockExecutionEnvironmentStateBuilder
}

func (testCase *CopyFileTestCase) Run(t *testing.T) {

	mockExecutionEnvironment := testCase.mockExecutionEnvironmentBuilder.BuildAndInit()

	actualError := util.CopyFile(testCase.sourceFileName, testCase.targetFileName)
	assert.Exactly(t, testCase.expectedError, actualError)

	expectedExecutionEnvironmentState := testCase.expectedExecutionEnvironmentStateBuilder.Build()
	mockExecutionEnvironment.GetCurrentState().AssertEquals(t, expectedExecutionEnvironmentState, testCase.description)

}

// ------------------------------------------------------------------------------

type GzipFileTestCase struct {
	description                              string
	fileName                                 string
	expectedError                            error
	mockExecutionEnvironmentBuilder          *utiltest.MockExecutionEnvironmentBuilder
	expectedExecutionEnvironmentStateBuilder *utiltest.MockExecutionEnvironmentStateBuilder
}

func (testCase *GzipFileTestCase) Run(t *testing.T) {

	mockExecutionEnvironment := testCase.mockExecutionEnvironmentBuilder.BuildAndInit()

	actualError := util.GzipFile(testCase.fileName)
	assert.Exactly(t, testCase.expectedError, actualError)

	expectedExecutionEnvironmentState := testCase.expectedExecutionEnvironmentStateBuilder.Build()
	mockExecutionEnvironment.GetCurrentState().AssertEquals(t, expectedExecutionEnvironmentState, testCase.description)

}

// ------------------------------------------------------------------------------

type GunzipFileTestCase struct {
	description                              string
	fileName                                 string
	expectedError                            error
	mockExecutionEnvironmentBuilder          *utiltest.MockExecutionEnvironmentBuilder
	expectedExecutionEnvironmentStateBuilder *utiltest.MockExecutionEnvironmentStateBuilder
}

func (testCase *GunzipFileTestCase) Run(t *testing.T) {

	mockExecutionEnvironment := testCase.mockExecutionEnvironmentBuilder.BuildAndInit()

	actualError := util.GunzipFile(testCase.fileName)
	assert.Exactly(t, testCase.expectedError, actualError)

	expectedExecutionEnvironmentState := testCase.expectedExecutionEnvironmentStateBuilder.Build()
	mockExecutionEnvironment.GetCurrentState().AssertEquals(t, expectedExecutionEnvironmentState, testCase.description)

}

// ------------------------------------------------------------------------------

type GetAllFileNamesWithExtensionTestCase struct {
	description                     string
	fileExtension                   string
	expectedFileNames               []string
	expectedError                   error
	mockExecutionEnvironmentBuilder *utiltest.MockExecutionEnvironmentBuilder
}

func (testCase *GetAllFileNamesWithExtensionTestCase) Run(t *testing.T) {

	_ = testCase.mockExecutionEnvironmentBuilder.BuildAndInit()

	actualFileNames, actualError := util.GetAllFileNamesWithExtension(testCase.fileExtension)
	assert.Exactly(t, testCase.expectedError, actualError)

	if testCase.expectedFileNames != nil {
		sort.Strings(testCase.expectedFileNames)
	}
	if actualFileNames != nil {
		sort.Strings(actualFileNames)
	}
	assert.Exactly(t, testCase.expectedFileNames, actualFileNames)

}
