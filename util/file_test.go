package util_test

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"path/filepath"
	"sort"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/stevengt/mppm/util"
)

func TestCopyFile(t *testing.T) {

	testCases := []*CopyFileTestCase{
		&CopyFileTestCase{
			sourceFileName:       "file1.txt",
			targetFileName:       "file2.bin",
			fileNamesAndContents: GetTestFileNamesAndContents(),
			openFileError:        nil,
			createFileError:      nil,
		},
		&CopyFileTestCase{
			sourceFileName:       "file1.txt",
			targetFileName:       "file2.bin",
			fileNamesAndContents: GetTestFileNamesAndContents(),
			openFileError:        errors.New("There was a problem opening the file."),
			createFileError:      nil,
		},
		&CopyFileTestCase{
			sourceFileName:       "file1.txt",
			targetFileName:       "file2.bin",
			fileNamesAndContents: GetTestFileNamesAndContents(),
			openFileError:        nil,
			createFileError:      errors.New("There was a problem creating the file."),
		},
		&CopyFileTestCase{
			sourceFileName:       "file1.txt",
			targetFileName:       "new-file",
			fileNamesAndContents: GetTestFileNamesAndContents(),
			openFileError:        nil,
			createFileError:      nil,
		},
		&CopyFileTestCase{
			sourceFileName:       "file1.txt",
			targetFileName:       "new-file",
			fileNamesAndContents: GetTestFileNamesAndContents(),
			openFileError:        errors.New("There was a problem opening the file."),
			createFileError:      nil,
		},
		&CopyFileTestCase{
			sourceFileName:       "file1.txt",
			targetFileName:       "new-file",
			fileNamesAndContents: GetTestFileNamesAndContents(),
			openFileError:        nil,
			createFileError:      errors.New("There was a problem creating the file."),
		},
		&CopyFileTestCase{
			sourceFileName:       "does-not-exist",
			targetFileName:       "new-file",
			fileNamesAndContents: GetTestFileNamesAndContents(),
			openFileError:        nil,
			createFileError:      nil,
		},
		&CopyFileTestCase{
			sourceFileName:       "empty-file.bin",
			targetFileName:       "new-file",
			fileNamesAndContents: GetTestFileNamesAndContents(),
			openFileError:        nil,
			createFileError:      nil,
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
			fileNamesAndContents:           GetTestFileNamesAndContents(),
			openFileError:                  nil,
			createFileError:                nil,
		},
		&GzipFileTestCase{
			fileName:                       "file1.txt",
			expectedCompressedFileContents: nil,
			fileNamesAndContents:           GetTestFileNamesAndContents(),
			openFileError:                  errors.New("There was a problem opening the file."),
			createFileError:                nil,
		},
		&GzipFileTestCase{
			fileName:                       "file1.txt",
			expectedCompressedFileContents: nil,
			fileNamesAndContents:           GetTestFileNamesAndContents(),
			openFileError:                  nil,
			createFileError:                errors.New("There was a problem creating the file."),
		},
		&GzipFileTestCase{
			fileName:                       "file2.bin",
			expectedCompressedFileContents: []byte{0x5a, 0xa3, 0x9c, 0x7c, 0x4, 0x0, 0x0, 0x0},
			fileNamesAndContents:           GetTestFileNamesAndContents(),
			openFileError:                  nil,
			createFileError:                nil,
		},
		&GzipFileTestCase{
			fileName:                       "file2.bin",
			expectedCompressedFileContents: nil,
			fileNamesAndContents:           GetTestFileNamesAndContents(),
			openFileError:                  errors.New("There was a problem opening the file."),
			createFileError:                nil,
		},
		&GzipFileTestCase{
			fileName:                       "file2.bin",
			expectedCompressedFileContents: nil,
			fileNamesAndContents:           GetTestFileNamesAndContents(),
			openFileError:                  nil,
			createFileError:                errors.New("There was a problem creating the file."),
		},
		&GzipFileTestCase{
			fileName:                       "empty-file.bin",
			expectedCompressedFileContents: []byte{0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0, 0x0},
			fileNamesAndContents:           GetTestFileNamesAndContents(),
			openFileError:                  nil,
			createFileError:                nil,
		},
		&GzipFileTestCase{
			fileName:                       "does-not-exist",
			expectedCompressedFileContents: nil,
			fileNamesAndContents:           GetTestFileNamesAndContents(),
			openFileError:                  nil,
			createFileError:                nil,
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
			fileNamesAndContents:             GetTestFileNamesAndContents(),
			openFileError:                    nil,
			createFileError:                  nil,
		},
		&GunzipFileTestCase{
			fileName:                         "file1.txt.gz",
			expectedUncompressedFileContents: nil,
			fileNamesAndContents:             GetTestFileNamesAndContents(),
			openFileError:                    errors.New("There was a problem opening the file."),
			createFileError:                  nil,
		},
		&GunzipFileTestCase{
			fileName:                         "file1.txt.gz",
			expectedUncompressedFileContents: nil,
			fileNamesAndContents:             GetTestFileNamesAndContents(),
			openFileError:                    nil,
			createFileError:                  errors.New("There was a problem creating the file."),
		},
		&GunzipFileTestCase{
			fileName:                         "does-not-exist.gz",
			expectedUncompressedFileContents: append([]byte("does-not-exist"), 0xa),
			fileNamesAndContents:             GetTestFileNamesAndContents(),
			openFileError:                    nil,
			createFileError:                  nil,
		},
		&GunzipFileTestCase{
			fileName:                         "does-not-exist.gz",
			expectedUncompressedFileContents: nil,
			fileNamesAndContents:             GetTestFileNamesAndContents(),
			openFileError:                    errors.New("There was a problem opening the file."),
			createFileError:                  nil,
		},
		&GunzipFileTestCase{
			fileName:                         "does-not-exist.gz",
			expectedUncompressedFileContents: nil,
			fileNamesAndContents:             GetTestFileNamesAndContents(),
			openFileError:                    nil,
			createFileError:                  errors.New("There was a problem creating the file."),
		},
		&GunzipFileTestCase{
			fileName:                         "empty-file.bin",
			expectedUncompressedFileContents: make([]byte, 0),
			fileNamesAndContents:             GetTestFileNamesAndContents(),
			openFileError:                    nil,
			createFileError:                  nil,
		},
		&GunzipFileTestCase{
			fileName:                         "does-not-exist",
			expectedUncompressedFileContents: nil,
			fileNamesAndContents:             GetTestFileNamesAndContents(),
			openFileError:                    nil,
			createFileError:                  nil,
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

		mockFileSystemDelegater := &MockFileSystemDelegater{}
		mockFileSystemDelegater.InitFiles(GetTestFileNamesAndContents())
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

		mockFileSystemDelegater := &MockFileSystemDelegater{
			WalkFilePathError: expectedError,
		}
		mockFileSystemDelegater.InitFiles(GetTestFileNamesAndContents())
		util.FileSystemProxy = mockFileSystemDelegater

		actualFileNames, actualError := util.GetAllFileNamesWithExtension(fileExtension)
		assert.NotNil(t, actualError)
		assert.Exactly(t, expectedFileNames, actualFileNames)
		assert.Exactly(t, expectedError, actualError)

	}

}

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
}

func (mockFileSystemDelegater *MockFileSystemDelegater) InitFiles(fileNamesAndContents map[string][]byte) {
	files := make(map[string]*MockFile)
	for fileName, fileContents := range fileNamesAndContents {
		files[fileName] = newMockFile(fileContents)
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
		mockFileSystemDelegater.Files[fileName] = newMockFile(fileContents)
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

// ------------------------------------------------------------------------------

type MockFile struct {
	contents         []byte
	bufferReadWriter *bufio.ReadWriter
	WasClosed        bool
}

func newMockFile(contents []byte) *MockFile {
	buffer := bytes.NewBuffer(contents)
	bufferReader := bufio.NewReader(buffer)
	bufferWriter := bufio.NewWriter(buffer)
	return &MockFile{
		contents:         contents,
		bufferReadWriter: bufio.NewReadWriter(bufferReader, bufferWriter),
		WasClosed:        false,
	}
}

func (mockFile *MockFile) Read(p []byte) (n int, err error) {
	return mockFile.bufferReadWriter.Read(p)
}

func (mockFile *MockFile) Write(p []byte) (n int, err error) {
	mockFile.contents = p
	return mockFile.bufferReadWriter.Write(p)
}

func (mockFile *MockFile) Close() error {
	mockFile.WasClosed = true
	return nil
}

// ------------------------------------------------------------------------------

type CopyFileTestCase struct {
	sourceFileName       string
	targetFileName       string
	fileNamesAndContents map[string][]byte
	openFileError        error
	createFileError      error
}

func (testCase *CopyFileTestCase) Run(t *testing.T) {

	mockFileSystemDelegater := &MockFileSystemDelegater{
		OpenFileError:   testCase.openFileError,
		CreateFileError: testCase.createFileError,
	}
	mockFileSystemDelegater.InitFiles(testCase.fileNamesAndContents)
	util.FileSystemProxy = mockFileSystemDelegater

	sourceFileBeforeCopy := mockFileSystemDelegater.Files[testCase.sourceFileName]
	targetFileBeforeCopy := mockFileSystemDelegater.Files[testCase.targetFileName]

	var sourceFileContentsBeforeCopy, targetFileContentsBeforeCopy []byte
	if sourceFileBeforeCopy != nil {
		sourceFileContentsBeforeCopy = sourceFileBeforeCopy.contents
	}
	if targetFileBeforeCopy != nil {
		targetFileContentsBeforeCopy = targetFileBeforeCopy.contents
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
		sourceFileContentsAfterCopy = sourceFileAfterCopy.contents
	}
	if targetFileAfterCopy != nil {
		targetFileContentsAfterCopy = targetFileAfterCopy.contents
	}

	if testCase.openFileError != nil {

		expectedError := testCase.openFileError
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

	} else if testCase.createFileError != nil {

		expectedError := testCase.createFileError
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
	fileNamesAndContents           map[string][]byte
	openFileError                  error
	createFileError                error
}

func (testCase *GzipFileTestCase) Run(t *testing.T) {

	mockFileSystemDelegater := &MockFileSystemDelegater{
		OpenFileError:   testCase.openFileError,
		CreateFileError: testCase.createFileError,
	}
	mockFileSystemDelegater.InitFiles(testCase.fileNamesAndContents)
	util.FileSystemProxy = mockFileSystemDelegater

	uncompressedFileName := testCase.fileName
	compressedFileName := uncompressedFileName + ".gz"

	uncompressedFileBeforeGzip := mockFileSystemDelegater.Files[uncompressedFileName]
	compressedFileBeforeGzip := mockFileSystemDelegater.Files[compressedFileName]

	var uncompressedFileContentsBeforeGzip, compressedFileContentsBeforeGzip []byte
	if uncompressedFileBeforeGzip != nil {
		uncompressedFileContentsBeforeGzip = uncompressedFileBeforeGzip.contents
	}
	if compressedFileBeforeGzip != nil {
		compressedFileContentsBeforeGzip = compressedFileBeforeGzip.contents
	}

	actualError := util.GzipFile(testCase.fileName)

	uncompressedFileAfterGzip := mockFileSystemDelegater.Files[uncompressedFileName]
	compressedFileAfterGzip := mockFileSystemDelegater.Files[compressedFileName]

	var uncompressedFileContentsAfterGzip, compressedFileContentsAfterGzip []byte
	if uncompressedFileAfterGzip != nil {
		uncompressedFileContentsAfterGzip = uncompressedFileAfterGzip.contents
	}
	if compressedFileAfterGzip != nil {
		compressedFileContentsAfterGzip = compressedFileAfterGzip.contents
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

	if testCase.openFileError != nil {

		expectedError := testCase.openFileError
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

	} else if testCase.createFileError != nil {

		expectedError := testCase.createFileError
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
	fileNamesAndContents             map[string][]byte
	openFileError                    error
	createFileError                  error
}

func (testCase *GunzipFileTestCase) Run(t *testing.T) {

	mockFileSystemDelegater := &MockFileSystemDelegater{
		OpenFileError:   testCase.openFileError,
		CreateFileError: testCase.createFileError,
	}
	mockFileSystemDelegater.InitFiles(testCase.fileNamesAndContents)
	util.FileSystemProxy = mockFileSystemDelegater

	compressedFileName := testCase.fileName
	uncompressedFileName := strings.TrimSuffix(compressedFileName, ".gz")

	uncompressedFileBeforeGunzip := mockFileSystemDelegater.Files[uncompressedFileName]
	compressedFileBeforeGunzip := mockFileSystemDelegater.Files[compressedFileName]

	var uncompressedFileContentsBeforeGunzip, compressedFileContentsBeforeGunzip []byte
	if uncompressedFileBeforeGunzip != nil {
		uncompressedFileContentsBeforeGunzip = uncompressedFileBeforeGunzip.contents
	}
	if compressedFileBeforeGunzip != nil {
		compressedFileContentsBeforeGunzip = compressedFileBeforeGunzip.contents
	}

	actualError := util.GunzipFile(testCase.fileName)

	uncompressedFileAfterGunzip := mockFileSystemDelegater.Files[uncompressedFileName]
	compressedFileAfterGunzip := mockFileSystemDelegater.Files[compressedFileName]

	var uncompressedFileContentsAfterGunzip, compressedFileContentsAfterGunzip []byte
	if uncompressedFileAfterGunzip != nil {
		uncompressedFileContentsAfterGunzip = uncompressedFileAfterGunzip.contents
	}
	if compressedFileAfterGunzip != nil {
		compressedFileContentsAfterGunzip = compressedFileAfterGunzip.contents
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

	if testCase.openFileError != nil {

		expectedError := testCase.openFileError
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

	} else if testCase.createFileError != nil {

		expectedError := testCase.createFileError
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
