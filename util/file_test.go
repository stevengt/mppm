package util_test

import (
	"errors"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stevengt/mppm/util"
	"github.com/stevengt/mppm/util/utiltest"
)

func TestCopyFile(t *testing.T) {

	testCases := []*CopyFileTestCase{
		&CopyFileTestCase{
			sourceFileName: "file1.txt",
			targetFileName: "file2.bin",
			mockFileSystemDelegaterBuilder: &utiltest.MockFileSystemDelegaterBuilder{
				FileNamesAndContentsAsBytes: utiltest.GetTestFileNamesAndContents(),
			},
		},
		&CopyFileTestCase{
			sourceFileName: "file1.txt",
			targetFileName: "file2.bin",
			mockFileSystemDelegaterBuilder: &utiltest.MockFileSystemDelegaterBuilder{
				FileNamesAndContentsAsBytes: utiltest.GetTestFileNamesAndContents(),
				UseDefaultOpenFileError:     true,
			},
		},
		&CopyFileTestCase{
			sourceFileName: "file1.txt",
			targetFileName: "file2.bin",
			mockFileSystemDelegaterBuilder: &utiltest.MockFileSystemDelegaterBuilder{
				FileNamesAndContentsAsBytes: utiltest.GetTestFileNamesAndContents(),
				UseDefaultCreateFileError:   true,
			},
		},
		&CopyFileTestCase{
			sourceFileName: "file1.txt",
			targetFileName: "new-file",
			mockFileSystemDelegaterBuilder: &utiltest.MockFileSystemDelegaterBuilder{
				FileNamesAndContentsAsBytes: utiltest.GetTestFileNamesAndContents(),
			},
		},
		&CopyFileTestCase{
			sourceFileName: "file1.txt",
			targetFileName: "new-file",
			mockFileSystemDelegaterBuilder: &utiltest.MockFileSystemDelegaterBuilder{
				FileNamesAndContentsAsBytes: utiltest.GetTestFileNamesAndContents(),
				UseDefaultOpenFileError:     true,
			},
		},
		&CopyFileTestCase{
			sourceFileName: "file1.txt",
			targetFileName: "new-file",
			mockFileSystemDelegaterBuilder: &utiltest.MockFileSystemDelegaterBuilder{
				FileNamesAndContentsAsBytes: utiltest.GetTestFileNamesAndContents(),
				UseDefaultCreateFileError:   true,
			},
		},
		&CopyFileTestCase{
			sourceFileName: "does-not-exist",
			targetFileName: "new-file",
			mockFileSystemDelegaterBuilder: &utiltest.MockFileSystemDelegaterBuilder{
				FileNamesAndContentsAsBytes: utiltest.GetTestFileNamesAndContents(),
			},
		},
		&CopyFileTestCase{
			sourceFileName: "empty-file.bin",
			targetFileName: "new-file",
			mockFileSystemDelegaterBuilder: &utiltest.MockFileSystemDelegaterBuilder{
				FileNamesAndContentsAsBytes: utiltest.GetTestFileNamesAndContents(),
			},
		},
	}

	for _, testCase := range testCases {
		testCase.Run(t)
	}

}

func TestGzipFile(t *testing.T) {

	testCases := []*GzipFileTestCase{
		&GzipFileTestCase{
			fileName:                       "file1.txt",
			expectedCompressedFileContents: []byte{0x4e, 0xb0, 0xa0, 0xe3, 0xf, 0x0, 0x0, 0x0},
			mockFileSystemDelegaterBuilder: &utiltest.MockFileSystemDelegaterBuilder{
				FileNamesAndContentsAsBytes: utiltest.GetTestFileNamesAndContents(),
			},
		},
		&GzipFileTestCase{
			fileName:                       "file1.txt",
			expectedCompressedFileContents: nil,
			mockFileSystemDelegaterBuilder: &utiltest.MockFileSystemDelegaterBuilder{
				FileNamesAndContentsAsBytes: utiltest.GetTestFileNamesAndContents(),
				UseDefaultOpenFileError:     true,
			},
		},
		&GzipFileTestCase{
			fileName:                       "file1.txt",
			expectedCompressedFileContents: nil,
			mockFileSystemDelegaterBuilder: &utiltest.MockFileSystemDelegaterBuilder{
				FileNamesAndContentsAsBytes: utiltest.GetTestFileNamesAndContents(),
				UseDefaultCreateFileError:   true,
			},
		},
		&GzipFileTestCase{
			fileName:                       "file2.bin",
			expectedCompressedFileContents: []byte{0x5a, 0xa3, 0x9c, 0x7c, 0x4, 0x0, 0x0, 0x0},
			mockFileSystemDelegaterBuilder: &utiltest.MockFileSystemDelegaterBuilder{
				FileNamesAndContentsAsBytes: utiltest.GetTestFileNamesAndContents(),
			},
		},
		&GzipFileTestCase{
			fileName:                       "file2.bin",
			expectedCompressedFileContents: nil,
			mockFileSystemDelegaterBuilder: &utiltest.MockFileSystemDelegaterBuilder{
				FileNamesAndContentsAsBytes: utiltest.GetTestFileNamesAndContents(),
				UseDefaultOpenFileError:     true,
			},
		},
		&GzipFileTestCase{
			fileName:                       "file2.bin",
			expectedCompressedFileContents: nil,
			mockFileSystemDelegaterBuilder: &utiltest.MockFileSystemDelegaterBuilder{
				FileNamesAndContentsAsBytes: utiltest.GetTestFileNamesAndContents(),
				UseDefaultCreateFileError:   true,
			},
		},
		&GzipFileTestCase{
			fileName:                       "empty-file.bin",
			expectedCompressedFileContents: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
			mockFileSystemDelegaterBuilder: &utiltest.MockFileSystemDelegaterBuilder{
				FileNamesAndContentsAsBytes: utiltest.GetTestFileNamesAndContents(),
			},
		},
		&GzipFileTestCase{
			fileName:                       "does-not-exist",
			expectedCompressedFileContents: nil,
			mockFileSystemDelegaterBuilder: &utiltest.MockFileSystemDelegaterBuilder{
				FileNamesAndContentsAsBytes: utiltest.GetTestFileNamesAndContents(),
			},
		},
	}

	for _, testCase := range testCases {
		testCase.Run(t)
	}

}

func TestGunzipFile(t *testing.T) {

	testCases := []*GunzipFileTestCase{
		&GunzipFileTestCase{
			fileName:                         "file1.txt.gz",
			expectedUncompressedFileContents: append([]byte("file 1 contents"), 0xa),
			mockFileSystemDelegaterBuilder: &utiltest.MockFileSystemDelegaterBuilder{
				FileNamesAndContentsAsBytes: utiltest.GetTestFileNamesAndContents(),
			},
		},
		&GunzipFileTestCase{
			fileName:                         "file1.txt.gz",
			expectedUncompressedFileContents: nil,
			mockFileSystemDelegaterBuilder: &utiltest.MockFileSystemDelegaterBuilder{
				FileNamesAndContentsAsBytes: utiltest.GetTestFileNamesAndContents(),
				UseDefaultOpenFileError:     true,
			},
		},
		&GunzipFileTestCase{
			fileName:                         "file1.txt.gz",
			expectedUncompressedFileContents: nil,
			mockFileSystemDelegaterBuilder: &utiltest.MockFileSystemDelegaterBuilder{
				FileNamesAndContentsAsBytes: utiltest.GetTestFileNamesAndContents(),
				UseDefaultCreateFileError:   true,
			},
		},
		&GunzipFileTestCase{
			fileName:                         "does-not-exist.gz",
			expectedUncompressedFileContents: append([]byte("does-not-exist"), 0xa),
			mockFileSystemDelegaterBuilder: &utiltest.MockFileSystemDelegaterBuilder{
				FileNamesAndContentsAsBytes: utiltest.GetTestFileNamesAndContents(),
			},
		},
		&GunzipFileTestCase{
			fileName:                         "does-not-exist.gz",
			expectedUncompressedFileContents: nil,
			mockFileSystemDelegaterBuilder: &utiltest.MockFileSystemDelegaterBuilder{
				FileNamesAndContentsAsBytes: utiltest.GetTestFileNamesAndContents(),
				UseDefaultOpenFileError:     true,
			},
		},
		&GunzipFileTestCase{
			fileName:                         "does-not-exist.gz",
			expectedUncompressedFileContents: nil,
			mockFileSystemDelegaterBuilder: &utiltest.MockFileSystemDelegaterBuilder{
				FileNamesAndContentsAsBytes: utiltest.GetTestFileNamesAndContents(),
				UseDefaultCreateFileError:   true,
			},
		},
		&GunzipFileTestCase{
			fileName:                         "empty-file.bin",
			expectedUncompressedFileContents: make([]byte, 0),
			mockFileSystemDelegaterBuilder: &utiltest.MockFileSystemDelegaterBuilder{
				FileNamesAndContentsAsBytes: utiltest.GetTestFileNamesAndContents(),
			},
		},
		&GunzipFileTestCase{
			fileName:                         "does-not-exist",
			expectedUncompressedFileContents: nil,
			mockFileSystemDelegaterBuilder: &utiltest.MockFileSystemDelegaterBuilder{
				FileNamesAndContentsAsBytes: utiltest.GetTestFileNamesAndContents(),
			},
		},
	}

	for _, testCase := range testCases {
		testCase.Run(t)
	}

}

func TestGetAllFileNamesWithExtension(t *testing.T) {

	fileExtensionsAndExpectedFileNames := map[string][]string{
		"txt":  []string{"file1.txt"},
		"bin":  []string{"file2.bin", "empty-file.bin"},
		"fake": []string{},
	}

	for fileExtension, expectedFileNames := range fileExtensionsAndExpectedFileNames {

		mockFileSystemDelegater := &utiltest.MockFileSystemDelegater{}
		mockFileSystemDelegater.InitFiles(utiltest.GetTestFileNamesAndContents())
		util.FileSystemProxy = mockFileSystemDelegater

		actualFileNames, err := util.GetAllFileNamesWithExtension(fileExtension)
		sort.Strings(expectedFileNames)
		sort.Strings(actualFileNames)

		assert.Nil(t, err)
		assert.Exactly(t, expectedFileNames, actualFileNames)
	}

	for fileExtension, _ := range fileExtensionsAndExpectedFileNames {

		expectedFileNames := make([]string, 0)
		expectedError := errors.New("There was a problem while walking the filepath.")

		mockFileSystemDelegater := &utiltest.MockFileSystemDelegater{
			WalkFilePathError: expectedError,
		}
		mockFileSystemDelegater.InitFiles(utiltest.GetTestFileNamesAndContents())
		util.FileSystemProxy = mockFileSystemDelegater

		actualFileNames, actualError := util.GetAllFileNamesWithExtension(fileExtension)
		assert.NotNil(t, actualError)
		assert.Exactly(t, expectedFileNames, actualFileNames)
		assert.Exactly(t, expectedError, actualError)

	}

}

// ------------------------------------------------------------------------------

type CopyFileTestCase struct {
	sourceFileName                 string
	targetFileName                 string
	mockFileSystemDelegaterBuilder *utiltest.MockFileSystemDelegaterBuilder
}

func (testCase *CopyFileTestCase) Run(t *testing.T) {

	mockFileSystemDelegater := utiltest.GetMockFileSystemDelegaterFromBuilderOrNil(testCase.mockFileSystemDelegaterBuilder)
	util.FileSystemProxy = mockFileSystemDelegater

	sourceFileBeforeCopy := mockFileSystemDelegater.Files[testCase.sourceFileName]
	targetFileBeforeCopy := mockFileSystemDelegater.Files[testCase.targetFileName]

	var sourceFileContentsBeforeCopy, targetFileContentsBeforeCopy []byte
	if sourceFileBeforeCopy != nil {
		sourceFileContentsBeforeCopy = sourceFileBeforeCopy.Contents
	}
	if targetFileBeforeCopy != nil {
		targetFileContentsBeforeCopy = targetFileBeforeCopy.Contents
	}

	actualError := util.CopyFile(testCase.sourceFileName, testCase.targetFileName)

	if sourceFileBeforeCopy == nil {
		assert.NotNil(t, actualError)
		return
	}

	sourceFileAfterCopy := mockFileSystemDelegater.Files[testCase.sourceFileName]
	targetFileAfterCopy := mockFileSystemDelegater.Files[testCase.targetFileName]

	var sourceFileContentsAfterCopy, targetFileContentsAfterCopy []byte
	if sourceFileAfterCopy != nil {
		sourceFileContentsAfterCopy = sourceFileAfterCopy.Contents
	}
	if targetFileAfterCopy != nil {
		targetFileContentsAfterCopy = targetFileAfterCopy.Contents
	}

	if testCase.mockFileSystemDelegaterBuilder.UseDefaultOpenFileError {

		expectedError := utiltest.DefaultOpenFileError
		assert.Exactly(t, expectedError, actualError)

		if sourceFileBeforeCopy != nil {
			assert.Same(t, sourceFileBeforeCopy, sourceFileAfterCopy)
			assert.Equal(t, sourceFileContentsBeforeCopy, sourceFileContentsAfterCopy)
			assert.False(t, sourceFileAfterCopy.WasClosed)
		}

		if targetFileBeforeCopy != nil {
			assert.Same(t, targetFileBeforeCopy, targetFileAfterCopy)
			assert.Equal(t, targetFileContentsBeforeCopy, targetFileContentsAfterCopy)
			assert.False(t, targetFileAfterCopy.WasClosed)
		} else {
			assert.Nil(t, targetFileAfterCopy)
		}

	} else if testCase.mockFileSystemDelegaterBuilder.UseDefaultCreateFileError {

		expectedError := utiltest.DefaultCreateFileError
		assert.Exactly(t, expectedError, actualError)

		if sourceFileBeforeCopy != nil {
			assert.Same(t, sourceFileBeforeCopy, sourceFileAfterCopy)
			assert.Equal(t, sourceFileContentsBeforeCopy, sourceFileContentsAfterCopy)
			assert.True(t, sourceFileAfterCopy.WasClosed)
		}

		if targetFileBeforeCopy != nil {
			assert.Same(t, targetFileBeforeCopy, targetFileAfterCopy)
			assert.Equal(t, targetFileContentsBeforeCopy, targetFileContentsAfterCopy)
			assert.False(t, targetFileAfterCopy.WasClosed)
		} else {
			assert.Nil(t, targetFileAfterCopy)
		}

	} else {

		assert.Nil(t, actualError)

		if sourceFileBeforeCopy != nil {
			assert.Same(t, sourceFileBeforeCopy, sourceFileAfterCopy)
			assert.Equal(t, sourceFileContentsBeforeCopy, sourceFileContentsAfterCopy)
			assert.True(t, sourceFileAfterCopy.WasClosed)
			assert.NotNil(t, targetFileAfterCopy)
			assert.True(t, targetFileAfterCopy.WasClosed)
			assert.Exactly(t, sourceFileContentsBeforeCopy, targetFileContentsAfterCopy)
		}

	}

}

// ------------------------------------------------------------------------------

type GzipFileTestCase struct {
	fileName                       string
	expectedCompressedFileContents []byte
	mockFileSystemDelegaterBuilder *utiltest.MockFileSystemDelegaterBuilder
}

func (testCase *GzipFileTestCase) Run(t *testing.T) {

	mockFileSystemDelegater := utiltest.GetMockFileSystemDelegaterFromBuilderOrNil(testCase.mockFileSystemDelegaterBuilder)
	util.FileSystemProxy = mockFileSystemDelegater

	uncompressedFileName := testCase.fileName
	compressedFileName := uncompressedFileName + ".gz"

	uncompressedFileBeforeGzip := mockFileSystemDelegater.Files[uncompressedFileName]
	compressedFileBeforeGzip := mockFileSystemDelegater.Files[compressedFileName]

	var uncompressedFileContentsBeforeGzip, compressedFileContentsBeforeGzip []byte
	if uncompressedFileBeforeGzip != nil {
		uncompressedFileContentsBeforeGzip = uncompressedFileBeforeGzip.Contents
	}
	if compressedFileBeforeGzip != nil {
		compressedFileContentsBeforeGzip = compressedFileBeforeGzip.Contents
	}

	actualError := util.GzipFile(testCase.fileName)

	uncompressedFileAfterGzip := mockFileSystemDelegater.Files[uncompressedFileName]
	compressedFileAfterGzip := mockFileSystemDelegater.Files[compressedFileName]

	var uncompressedFileContentsAfterGzip, compressedFileContentsAfterGzip []byte
	if uncompressedFileAfterGzip != nil {
		uncompressedFileContentsAfterGzip = uncompressedFileAfterGzip.Contents
	}
	if compressedFileAfterGzip != nil {
		compressedFileContentsAfterGzip = compressedFileAfterGzip.Contents
	}

	if uncompressedFileBeforeGzip == nil {
		assert.NotNil(t, actualError)
		if compressedFileBeforeGzip != nil {
			assert.Same(t, compressedFileBeforeGzip, compressedFileAfterGzip)
			assert.Equal(t, compressedFileContentsBeforeGzip, compressedFileContentsAfterGzip)
			assert.False(t, compressedFileBeforeGzip.WasClosed)
		}
		return
	}

	if testCase.mockFileSystemDelegaterBuilder.UseDefaultOpenFileError {

		expectedError := utiltest.DefaultOpenFileError
		assert.Exactly(t, expectedError, actualError)

		if uncompressedFileBeforeGzip != nil {
			assert.Same(t, uncompressedFileBeforeGzip, uncompressedFileAfterGzip)
			assert.Equal(t, uncompressedFileContentsBeforeGzip, uncompressedFileContentsAfterGzip)
			assert.False(t, uncompressedFileBeforeGzip.WasClosed)
		}

		if compressedFileBeforeGzip != nil {
			assert.Same(t, compressedFileBeforeGzip, compressedFileAfterGzip)
			assert.Equal(t, compressedFileContentsBeforeGzip, compressedFileContentsAfterGzip)
			assert.False(t, compressedFileBeforeGzip.WasClosed)
		} else {
			assert.Nil(t, compressedFileAfterGzip)
		}

	} else if testCase.mockFileSystemDelegaterBuilder.UseDefaultCreateFileError {

		expectedError := utiltest.DefaultCreateFileError
		assert.Exactly(t, expectedError, actualError)

		if uncompressedFileBeforeGzip != nil {
			assert.Same(t, uncompressedFileBeforeGzip, uncompressedFileAfterGzip)
			assert.Equal(t, uncompressedFileContentsBeforeGzip, uncompressedFileContentsAfterGzip)
			assert.True(t, uncompressedFileAfterGzip.WasClosed)
		}

		if compressedFileBeforeGzip != nil {
			assert.Same(t, compressedFileBeforeGzip, compressedFileAfterGzip)
			assert.Equal(t, compressedFileContentsBeforeGzip, compressedFileContentsAfterGzip)
			assert.False(t, compressedFileAfterGzip.WasClosed)
		} else {
			assert.Nil(t, compressedFileAfterGzip)
		}

	} else {

		assert.Nil(t, actualError)

		if uncompressedFileBeforeGzip != nil {
			assert.Same(t, uncompressedFileBeforeGzip, uncompressedFileAfterGzip)
			assert.Equal(t, uncompressedFileContentsBeforeGzip, uncompressedFileContentsAfterGzip)
			assert.True(t, uncompressedFileAfterGzip.WasClosed)
			assert.NotNil(t, compressedFileAfterGzip)
			assert.True(t, compressedFileAfterGzip.WasClosed)
			assert.Exactly(t, testCase.expectedCompressedFileContents, compressedFileContentsAfterGzip)
		}

	}

}

// ------------------------------------------------------------------------------

type GunzipFileTestCase struct {
	fileName                         string
	expectedUncompressedFileContents []byte
	mockFileSystemDelegaterBuilder   *utiltest.MockFileSystemDelegaterBuilder
}

func (testCase *GunzipFileTestCase) Run(t *testing.T) {

	mockFileSystemDelegater := utiltest.GetMockFileSystemDelegaterFromBuilderOrNil(testCase.mockFileSystemDelegaterBuilder)
	util.FileSystemProxy = mockFileSystemDelegater

	compressedFileName := testCase.fileName
	uncompressedFileName := strings.TrimSuffix(compressedFileName, ".gz")

	uncompressedFileBeforeGunzip := mockFileSystemDelegater.Files[uncompressedFileName]
	compressedFileBeforeGunzip := mockFileSystemDelegater.Files[compressedFileName]

	var uncompressedFileContentsBeforeGunzip, compressedFileContentsBeforeGunzip []byte
	if uncompressedFileBeforeGunzip != nil {
		uncompressedFileContentsBeforeGunzip = uncompressedFileBeforeGunzip.Contents
	}
	if compressedFileBeforeGunzip != nil {
		compressedFileContentsBeforeGunzip = compressedFileBeforeGunzip.Contents
	}

	actualError := util.GunzipFile(testCase.fileName)

	uncompressedFileAfterGunzip := mockFileSystemDelegater.Files[uncompressedFileName]
	compressedFileAfterGunzip := mockFileSystemDelegater.Files[compressedFileName]

	var uncompressedFileContentsAfterGunzip, compressedFileContentsAfterGunzip []byte
	if uncompressedFileAfterGunzip != nil {
		uncompressedFileContentsAfterGunzip = uncompressedFileAfterGunzip.Contents
	}
	if compressedFileAfterGunzip != nil {
		compressedFileContentsAfterGunzip = compressedFileAfterGunzip.Contents
	}

	if compressedFileBeforeGunzip == nil {
		assert.NotNil(t, actualError)
		if uncompressedFileBeforeGunzip != nil {
			assert.Same(t, uncompressedFileBeforeGunzip, uncompressedFileAfterGunzip)
			assert.Equal(t, uncompressedFileContentsBeforeGunzip, uncompressedFileContentsAfterGunzip)
			assert.False(t, uncompressedFileBeforeGunzip.WasClosed)
		}
		return
	} else {
		if len(compressedFileContentsBeforeGunzip) == 0 {
			assert.NotNil(t, actualError)
			assert.Same(t, compressedFileBeforeGunzip, compressedFileAfterGunzip)
			assert.Equal(t, compressedFileContentsBeforeGunzip, compressedFileContentsAfterGunzip)
			assert.True(t, compressedFileAfterGunzip.WasClosed)
			assert.True(t, uncompressedFileAfterGunzip.WasClosed)
			return
		}
	}

	if testCase.mockFileSystemDelegaterBuilder.UseDefaultOpenFileError {

		expectedError := utiltest.DefaultOpenFileError
		assert.Exactly(t, expectedError, actualError)

		if compressedFileBeforeGunzip != nil {
			assert.Same(t, compressedFileBeforeGunzip, compressedFileAfterGunzip)
			assert.Equal(t, compressedFileContentsBeforeGunzip, compressedFileContentsAfterGunzip)
			assert.False(t, compressedFileBeforeGunzip.WasClosed)
		}

		if uncompressedFileBeforeGunzip != nil {
			assert.Same(t, uncompressedFileBeforeGunzip, uncompressedFileAfterGunzip)
			assert.Equal(t, uncompressedFileContentsBeforeGunzip, uncompressedFileContentsAfterGunzip)
			assert.False(t, uncompressedFileBeforeGunzip.WasClosed)
		} else {
			assert.Nil(t, uncompressedFileAfterGunzip)
		}

	} else if testCase.mockFileSystemDelegaterBuilder.UseDefaultCreateFileError {

		expectedError := utiltest.DefaultCreateFileError
		assert.Exactly(t, expectedError, actualError)

		if compressedFileBeforeGunzip != nil {
			assert.Same(t, compressedFileBeforeGunzip, compressedFileAfterGunzip)
			assert.Equal(t, compressedFileContentsBeforeGunzip, compressedFileContentsAfterGunzip)
			assert.True(t, compressedFileAfterGunzip.WasClosed)
		}

		if uncompressedFileBeforeGunzip != nil {
			assert.Same(t, uncompressedFileBeforeGunzip, uncompressedFileAfterGunzip)
			assert.Equal(t, uncompressedFileContentsBeforeGunzip, uncompressedFileContentsAfterGunzip)
			assert.False(t, uncompressedFileAfterGunzip.WasClosed)
		} else {
			assert.Nil(t, uncompressedFileAfterGunzip)
		}

	} else {

		assert.Nil(t, actualError)

		if compressedFileBeforeGunzip != nil {
			assert.Nil(t, compressedFileAfterGunzip)
			assert.NotNil(t, uncompressedFileAfterGunzip)
			assert.True(t, uncompressedFileAfterGunzip.WasClosed)
			assert.Exactly(t, testCase.expectedUncompressedFileContents, uncompressedFileContentsAfterGunzip)
		}

	}

}
