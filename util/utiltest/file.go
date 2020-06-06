package utiltest

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"path/filepath"
	"strings"
)

// ------------------------------------------------------------------------------

var DefaultOpenFileError error = errors.New("There was a problem opening the file.")

var DefaultCreateFileError error = errors.New("There was a problem creating the file.")

var DefaultRemoveFileError error = errors.New("There was a problem removing the file.")

var DefaultWalkFilePathError error = errors.New("There was a problem walking the file path.")

var DefaultUserHomeDirError error = errors.New("There was a problem getting the user's home directory.")

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

func GetMockFileSystemDelegaterFromBuilderOrNil(mockFileSystemDelegaterBuilder *MockFileSystemDelegaterBuilder) *MockFileSystemDelegater {
	if mockFileSystemDelegaterBuilder != nil {
		return mockFileSystemDelegaterBuilder.Build()
	} else {
		return &MockFileSystemDelegater{}
	}
}

// ------------------------------------------------------------------------------

type MockFileSystemDelegaterBuilder struct {
	Files                       map[string]*MockFile
	FileNamesAndContentsAsBytes map[string][]byte // Use this if you want the builder to create MockFile instances for you.
	UseDefaultOpenFileError     bool
	UseDefaultCreateFileError   bool
	UseDefaultRemoveFileError   bool
	UseDefaultWalkFilePathError bool
	UseDefaultUserHomeDirError  bool
}

func (builder *MockFileSystemDelegaterBuilder) Build() *MockFileSystemDelegater {

	mockFileSystemDelegater := NewDefaultMockFileSystemDelegater()

	if builder.Files != nil {
		mockFileSystemDelegater.Files = builder.Files
	}

	if builder.FileNamesAndContentsAsBytes != nil {
		for fileName, fileContents := range builder.FileNamesAndContentsAsBytes {
			mockFileSystemDelegater.Files[fileName] = NewMockFile(fileContents)
		}
	}

	if builder.UseDefaultOpenFileError {
		mockFileSystemDelegater.OpenFileError = DefaultOpenFileError
	}

	if builder.UseDefaultCreateFileError {
		mockFileSystemDelegater.CreateFileError = DefaultCreateFileError
	}

	if builder.UseDefaultRemoveFileError {
		mockFileSystemDelegater.RemoveFileError = DefaultRemoveFileError
	}

	if builder.UseDefaultWalkFilePathError {
		mockFileSystemDelegater.WalkFilePathError = DefaultWalkFilePathError
	}

	if builder.UseDefaultUserHomeDirError {
		mockFileSystemDelegater.UserHomeDirError = DefaultUserHomeDirError
	}

	return mockFileSystemDelegater

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

func NewDefaultMockFileSystemDelegater() *MockFileSystemDelegater {
	return &MockFileSystemDelegater{
		Files: make(map[string]*MockFile),
	}
}

func (mockFileSystemDelegater *MockFileSystemDelegater) InitFiles(fileNamesAndContents map[string][]byte) {
	files := make(map[string]*MockFile)
	for fileName, fileContents := range fileNamesAndContents {
		files[fileName] = NewMockFile(fileContents)
	}
	mockFileSystemDelegater.Files = files
}

func (mockFileSystemDelegater *MockFileSystemDelegater) GetMockFileAndContentsIfFileExistsElseReturnNil(fileName string) (file *MockFile, contents []byte) {
	if mockFileSystemDelegater.DoesFileExist(fileName) {
		file = mockFileSystemDelegater.Files[fileName]
		contents = file.Contents
		return
	}
	return nil, nil
}

func (mockFileSystemDelegater *MockFileSystemDelegater) OpenFile(fileName string) (file io.ReadWriteCloser, err error) {
	err = mockFileSystemDelegater.OpenFileError
	if err == nil {
		var doesFileExist bool
		file, doesFileExist = mockFileSystemDelegater.Files[fileName]
		if !doesFileExist {
			err = errors.New("Unable to open file " + fileName)
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
	mockFile := &MockFile{
		Contents:  contents,
		WasClosed: false,
	}
	mockFile.resetBuffer()
	return mockFile
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
	mockFile.resetBuffer()
	return nil
}

func (mockFile *MockFile) resetBuffer() {
	buffer := bytes.NewBuffer(mockFile.Contents)
	bufferReader := bufio.NewReader(buffer)
	bufferWriter := bufio.NewWriter(buffer)
	mockFile.bufferReadWriter = bufio.NewReadWriter(bufferReader, bufferWriter)
}
