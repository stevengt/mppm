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
			SourceFileName:       "file1.txt",
			TargetFileName:       "file2.bin",
			FileNamesAndContents: utiltest.GetTestFileNamesAndContents(),
			OpenFileError:        nil,
			CreateFileError:      nil,
		},
		&CopyFileTestCase{
			SourceFileName:       "file1.txt",
			TargetFileName:       "file2.bin",
			FileNamesAndContents: utiltest.GetTestFileNamesAndContents(),
			OpenFileError:        errors.New("There was a problem opening the file."),
			CreateFileError:      nil,
		},
		&CopyFileTestCase{
			SourceFileName:       "file1.txt",
			TargetFileName:       "file2.bin",
			FileNamesAndContents: utiltest.GetTestFileNamesAndContents(),
			OpenFileError:        nil,
			CreateFileError:      errors.New("There was a problem creating the file."),
		},
		&CopyFileTestCase{
			SourceFileName:       "file1.txt",
			TargetFileName:       "new-file",
			FileNamesAndContents: utiltest.GetTestFileNamesAndContents(),
			OpenFileError:        nil,
			CreateFileError:      nil,
		},
		&CopyFileTestCase{
			SourceFileName:       "file1.txt",
			TargetFileName:       "new-file",
			FileNamesAndContents: utiltest.GetTestFileNamesAndContents(),
			OpenFileError:        errors.New("There was a problem opening the file."),
			CreateFileError:      nil,
		},
		&CopyFileTestCase{
			SourceFileName:       "file1.txt",
			TargetFileName:       "new-file",
			FileNamesAndContents: utiltest.GetTestFileNamesAndContents(),
			OpenFileError:        nil,
			CreateFileError:      errors.New("There was a problem creating the file."),
		},
		&CopyFileTestCase{
			SourceFileName:       "does-not-exist",
			TargetFileName:       "new-file",
			FileNamesAndContents: utiltest.GetTestFileNamesAndContents(),
			OpenFileError:        nil,
			CreateFileError:      nil,
		},
		&CopyFileTestCase{
			SourceFileName:       "empty-file.bin",
			TargetFileName:       "new-file",
			FileNamesAndContents: utiltest.GetTestFileNamesAndContents(),
			OpenFileError:        nil,
			CreateFileError:      nil,
		},
	}

	for _, testCase := range testCases {
		testCase.Run(t)
	}

}

func TestGzipFile(t *testing.T) {

	testCases := []*GzipFileTestCase{
		&GzipFileTestCase{
			FileName:                       "file1.txt",
			ExpectedCompressedFileContents: []byte{0x4e, 0xb0, 0xa0, 0xe3, 0xf, 0x0, 0x0, 0x0},
			FileNamesAndContents:           utiltest.GetTestFileNamesAndContents(),
			OpenFileError:                  nil,
			CreateFileError:                nil,
		},
		&GzipFileTestCase{
			FileName:                       "file1.txt",
			ExpectedCompressedFileContents: nil,
			FileNamesAndContents:           utiltest.GetTestFileNamesAndContents(),
			OpenFileError:                  errors.New("There was a problem opening the file."),
			CreateFileError:                nil,
		},
		&GzipFileTestCase{
			FileName:                       "file1.txt",
			ExpectedCompressedFileContents: nil,
			FileNamesAndContents:           utiltest.GetTestFileNamesAndContents(),
			OpenFileError:                  nil,
			CreateFileError:                errors.New("There was a problem creating the file."),
		},
		&GzipFileTestCase{
			FileName:                       "file2.bin",
			ExpectedCompressedFileContents: []byte{0x5a, 0xa3, 0x9c, 0x7c, 0x4, 0x0, 0x0, 0x0},
			FileNamesAndContents:           utiltest.GetTestFileNamesAndContents(),
			OpenFileError:                  nil,
			CreateFileError:                nil,
		},
		&GzipFileTestCase{
			FileName:                       "file2.bin",
			ExpectedCompressedFileContents: nil,
			FileNamesAndContents:           utiltest.GetTestFileNamesAndContents(),
			OpenFileError:                  errors.New("There was a problem opening the file."),
			CreateFileError:                nil,
		},
		&GzipFileTestCase{
			FileName:                       "file2.bin",
			ExpectedCompressedFileContents: nil,
			FileNamesAndContents:           utiltest.GetTestFileNamesAndContents(),
			OpenFileError:                  nil,
			CreateFileError:                errors.New("There was a problem creating the file."),
		},
		&GzipFileTestCase{
			FileName:                       "empty-file.bin",
			ExpectedCompressedFileContents: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
			FileNamesAndContents:           utiltest.GetTestFileNamesAndContents(),
			OpenFileError:                  nil,
			CreateFileError:                nil,
		},
		&GzipFileTestCase{
			FileName:                       "does-not-exist",
			ExpectedCompressedFileContents: nil,
			FileNamesAndContents:           utiltest.GetTestFileNamesAndContents(),
			OpenFileError:                  nil,
			CreateFileError:                nil,
		},
	}

	for _, testCase := range testCases {
		testCase.Run(t)
	}

}

func TestGunzipFile(t *testing.T) {

	testCases := []*GunzipFileTestCase{
		&GunzipFileTestCase{
			FileName:                         "file1.txt.gz",
			ExpectedUncompressedFileContents: append([]byte("file 1 contents"), 0xa),
			FileNamesAndContents:             utiltest.GetTestFileNamesAndContents(),
			OpenFileError:                    nil,
			CreateFileError:                  nil,
		},
		&GunzipFileTestCase{
			FileName:                         "file1.txt.gz",
			ExpectedUncompressedFileContents: nil,
			FileNamesAndContents:             utiltest.GetTestFileNamesAndContents(),
			OpenFileError:                    errors.New("There was a problem opening the file."),
			CreateFileError:                  nil,
		},
		&GunzipFileTestCase{
			FileName:                         "file1.txt.gz",
			ExpectedUncompressedFileContents: nil,
			FileNamesAndContents:             utiltest.GetTestFileNamesAndContents(),
			OpenFileError:                    nil,
			CreateFileError:                  errors.New("There was a problem creating the file."),
		},
		&GunzipFileTestCase{
			FileName:                         "does-not-exist.gz",
			ExpectedUncompressedFileContents: append([]byte("does-not-exist"), 0xa),
			FileNamesAndContents:             utiltest.GetTestFileNamesAndContents(),
			OpenFileError:                    nil,
			CreateFileError:                  nil,
		},
		&GunzipFileTestCase{
			FileName:                         "does-not-exist.gz",
			ExpectedUncompressedFileContents: nil,
			FileNamesAndContents:             utiltest.GetTestFileNamesAndContents(),
			OpenFileError:                    errors.New("There was a problem opening the file."),
			CreateFileError:                  nil,
		},
		&GunzipFileTestCase{
			FileName:                         "does-not-exist.gz",
			ExpectedUncompressedFileContents: nil,
			FileNamesAndContents:             utiltest.GetTestFileNamesAndContents(),
			OpenFileError:                    nil,
			CreateFileError:                  errors.New("There was a problem creating the file."),
		},
		&GunzipFileTestCase{
			FileName:                         "empty-file.bin",
			ExpectedUncompressedFileContents: make([]byte, 0),
			FileNamesAndContents:             utiltest.GetTestFileNamesAndContents(),
			OpenFileError:                    nil,
			CreateFileError:                  nil,
		},
		&GunzipFileTestCase{
			FileName:                         "does-not-exist",
			ExpectedUncompressedFileContents: nil,
			FileNamesAndContents:             utiltest.GetTestFileNamesAndContents(),
			OpenFileError:                    nil,
			CreateFileError:                  nil,
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
	SourceFileName       string
	TargetFileName       string
	FileNamesAndContents map[string][]byte
	OpenFileError        error
	CreateFileError      error
}

func (testCase *CopyFileTestCase) Run(t *testing.T) {

	mockFileSystemDelegater := &utiltest.MockFileSystemDelegater{
		OpenFileError:   testCase.OpenFileError,
		CreateFileError: testCase.CreateFileError,
	}
	mockFileSystemDelegater.InitFiles(testCase.FileNamesAndContents)
	util.FileSystemProxy = mockFileSystemDelegater

	sourceFileBeforeCopy := mockFileSystemDelegater.Files[testCase.SourceFileName]
	targetFileBeforeCopy := mockFileSystemDelegater.Files[testCase.TargetFileName]

	var sourceFileContentsBeforeCopy, targetFileContentsBeforeCopy []byte
	if sourceFileBeforeCopy != nil {
		sourceFileContentsBeforeCopy = sourceFileBeforeCopy.Contents
	}
	if targetFileBeforeCopy != nil {
		targetFileContentsBeforeCopy = targetFileBeforeCopy.Contents
	}

	actualError := util.CopyFile(testCase.SourceFileName, testCase.TargetFileName)

	if sourceFileBeforeCopy == nil {
		assert.NotNil(t, actualError)
		return
	}

	sourceFileAfterCopy := mockFileSystemDelegater.Files[testCase.SourceFileName]
	targetFileAfterCopy := mockFileSystemDelegater.Files[testCase.TargetFileName]

	var sourceFileContentsAfterCopy, targetFileContentsAfterCopy []byte
	if sourceFileAfterCopy != nil {
		sourceFileContentsAfterCopy = sourceFileAfterCopy.Contents
	}
	if targetFileAfterCopy != nil {
		targetFileContentsAfterCopy = targetFileAfterCopy.Contents
	}

	if testCase.OpenFileError != nil {

		expectedError := testCase.OpenFileError
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

	} else if testCase.CreateFileError != nil {

		expectedError := testCase.CreateFileError
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
	FileName                       string
	ExpectedCompressedFileContents []byte
	FileNamesAndContents           map[string][]byte
	OpenFileError                  error
	CreateFileError                error
}

func (testCase *GzipFileTestCase) Run(t *testing.T) {

	mockFileSystemDelegater := &utiltest.MockFileSystemDelegater{
		OpenFileError:   testCase.OpenFileError,
		CreateFileError: testCase.CreateFileError,
	}
	mockFileSystemDelegater.InitFiles(testCase.FileNamesAndContents)
	util.FileSystemProxy = mockFileSystemDelegater

	uncompressedFileName := testCase.FileName
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

	actualError := util.GzipFile(testCase.FileName)

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

	if testCase.OpenFileError != nil {

		expectedError := testCase.OpenFileError
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

	} else if testCase.CreateFileError != nil {

		expectedError := testCase.CreateFileError
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
			assert.Exactly(t, testCase.ExpectedCompressedFileContents, compressedFileContentsAfterGzip)
		}

	}

}

// ------------------------------------------------------------------------------

type GunzipFileTestCase struct {
	FileName                         string
	ExpectedUncompressedFileContents []byte
	FileNamesAndContents             map[string][]byte
	OpenFileError                    error
	CreateFileError                  error
}

func (testCase *GunzipFileTestCase) Run(t *testing.T) {

	mockFileSystemDelegater := &utiltest.MockFileSystemDelegater{
		OpenFileError:   testCase.OpenFileError,
		CreateFileError: testCase.CreateFileError,
	}
	mockFileSystemDelegater.InitFiles(testCase.FileNamesAndContents)
	util.FileSystemProxy = mockFileSystemDelegater

	compressedFileName := testCase.FileName
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

	actualError := util.GunzipFile(testCase.FileName)

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

	if testCase.OpenFileError != nil {

		expectedError := testCase.OpenFileError
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

	} else if testCase.CreateFileError != nil {

		expectedError := testCase.CreateFileError
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
			assert.Exactly(t, testCase.ExpectedUncompressedFileContents, uncompressedFileContentsAfterGunzip)
		}

	}

}
