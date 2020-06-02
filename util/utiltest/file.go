package utiltest

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stevengt/mppm/util"
	"github.com/stretchr/testify/assert"
)

// ------------------------------------------------------------------------------

func GetTestFileNamesAndContents() map[string][]byte {
	return map[string][]byte{
		"file1.txt": []byte("file 1 contents"),
		"file1.txt.gz": []byte{
			0x1f, 0x8b, 0x08, 0x08, 0x8a, 0xab, 0xd4, 0x5e, 0x00, 0x03,
			0x74, 0x65, 0x73, 0x74, 0x00, 0x4b, 0xcb, 0xcc, 0x49, 0x55,
			0x30, 0x54, 0x48, 0xce, 0xcf, 0x2b, 0x49, 0xcd, 0x2b, 0x29,
			0xe6, 0x02, 0x00, 0xb4, 0xca, 0x50, 0xa3, 0x10, 0x00, 0x00, 0x00,
		},
		"file2.bin":      []byte{0xDE, 0xAD, 0xBE, 0xEF},
		"empty-file.bin": make([]byte, 0),
		"does-not-exist.gz": []byte{
			0x1f, 0x8b, 0x08, 0x08, 0x96, 0xa8, 0xd4, 0x5e, 0x00, 0x03,
			0x74, 0x65, 0x73, 0x74, 0x00, 0x4b, 0xc9, 0x4f, 0x2d, 0xd6,
			0xcd, 0xcb, 0x2f, 0xd1, 0x4d, 0xad, 0xc8, 0x2c, 0x2e, 0xe1,
			0x02, 0x00, 0x2a, 0x53, 0xd8, 0x28, 0x0f, 0x00, 0x00, 0x00,
		},
	}
}

// ------------------------------------------------------------------------------

type MockFileSystemDelegater struct {
	Files             map[string]*MockFile // Map of file names to mocked file instances.
	OpenFileError     error
	CreateFileError   error
	RemoveFileError   error
	WalkFilePathError error
	UserHomeDirError  error
}

func (mockFileSystemDelegater *MockFileSystemDelegater) InitFiles(fileNamesAndContents map[string][]byte) {
	files := make(map[string]*MockFile)
	for fileName, fileContents := range fileNamesAndContents {
		files[fileName] = NewMockFile(fileContents)
	}
	mockFileSystemDelegater.Files = files
}

func (mockFileSystemDelegater *MockFileSystemDelegater) OpenFile(fileName string) (file io.ReadWriteCloser, err error) {
	err = mockFileSystemDelegater.OpenFileError
	if err == nil {
		var doesFileExist bool
		file, doesFileExist = mockFileSystemDelegater.Files[fileName]
		if !doesFileExist {
			err = errors.New("Unable to open file" + fileName)
		}
	}
	return
}

func (mockFileSystemDelegater *MockFileSystemDelegater) CreateFile(fileName string) (file io.ReadWriteCloser, err error) {
	err = mockFileSystemDelegater.CreateFileError
	if err == nil {
		if mockFileSystemDelegater.Files[fileName] != nil {
			err = mockFileSystemDelegater.RemoveFile(fileName)
			if err != nil {
				return
			}
		}
		fileContents := make([]byte, 0)
		mockFileSystemDelegater.Files[fileName] = NewMockFile(fileContents)
		file = mockFileSystemDelegater.Files[fileName]
	}
	return
}

func (mockFileSystemDelegater *MockFileSystemDelegater) RemoveFile(fileName string) (err error) {
	err = mockFileSystemDelegater.RemoveFileError
	if err == nil {
		delete(mockFileSystemDelegater.Files, fileName)
	}
	return
}

func (mockFileSystemDelegater *MockFileSystemDelegater) WalkFilePath(root string, walkFn filepath.WalkFunc) (err error) {
	err = mockFileSystemDelegater.WalkFilePathError
	if err == nil {
		for fileName, _ := range mockFileSystemDelegater.Files {
			err = walkFn(fileName, nil, nil)
			if err != nil {
				return
			}
		}
	}
	return
}

func (mockFileSystemDelegater *MockFileSystemDelegater) UserHomeDir() (string, error) {
	err := mockFileSystemDelegater.UserHomeDirError
	if err == nil {
		return "/home/testuser", err
	}
	return "", err
}

func (mockFileSystemDelegater *MockFileSystemDelegater) JoinFilePath(elem ...string) string {
	return strings.Join(elem, "/")
}

func (mockFileSystemDelegater *MockFileSystemDelegater) DoesFileExist(filePath string) bool {
	var doesFileExist bool
	_, doesFileExist = mockFileSystemDelegater.Files[filePath]
	return doesFileExist
}

// ------------------------------------------------------------------------------

type MockFile struct {
	Contents         []byte
	bufferReadWriter *bufio.ReadWriter
	WasClosed        bool
}

func NewMockFile(contents []byte) *MockFile {
	buffer := bytes.NewBuffer(contents)
	bufferReader := bufio.NewReader(buffer)
	bufferWriter := bufio.NewWriter(buffer)
	return &MockFile{
		Contents:         contents,
		bufferReadWriter: bufio.NewReadWriter(bufferReader, bufferWriter),
		WasClosed:        false,
	}
}

func (mockFile *MockFile) Read(p []byte) (n int, err error) {
	return mockFile.bufferReadWriter.Read(p)
}

func (mockFile *MockFile) Write(p []byte) (n int, err error) {
	mockFile.Contents = p
	return mockFile.bufferReadWriter.Write(p)
}

func (mockFile *MockFile) Close() error {
	mockFile.WasClosed = true
	return nil
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

	mockFileSystemDelegater := &MockFileSystemDelegater{
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

	mockFileSystemDelegater := &MockFileSystemDelegater{
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

	mockFileSystemDelegater := &MockFileSystemDelegater{
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
