package util_test

import (
	"bufio"
	"bytes"
	"errors"
	"io"
	"path/filepath"
)

type MockFileSystemDelegater struct {
	Err   error                // Error to return for a specific instance of a mocked method.
	Files map[string]*MockFile // Map of file names to mocked file instances.
}

func (mockFileSystemDelegater *MockFileSystemDelegater) OpenFile(fileName string) (file io.ReadWriteCloser, err error) {
	err = mockFileSystemDelegater.Err
	if err != nil {
		file = mockFileSystemDelegater.Files[fileName]
		if file == nil {
			err = errors.New("Unable to open file" + fileName)
		}
	}
	return
}

func (mockFileSystemDelegater *MockFileSystemDelegater) CreateFile(fileName string) (file io.ReadWriteCloser, err error) {
	err = mockFileSystemDelegater.Err
	if err != nil {
		if mockFileSystemDelegater.Files[fileName] != nil {
			err = mockFileSystemDelegater.RemoveFile(fileName)
			if err != nil {
				return
			}
		}
		mockFileSystemDelegater.Files[fileName] = newMockFile()
		file = mockFileSystemDelegater.Files[fileName]
	}
	return
}

func (mockFileSystemDelegater *MockFileSystemDelegater) RemoveFile(fileName string) (err error) {
	err = mockFileSystemDelegater.Err
	if err != nil {
		delete(mockFileSystemDelegater.Files, fileName)
	}
	return
}

func (mockFileSystemDelegater *MockFileSystemDelegater) WalkFilePath(root string, walkFn filepath.WalkFunc) (err error) {
	err = mockFileSystemDelegater.Err
	if err != nil {
		for fileName, _ := range mockFileSystemDelegater.Files {
			err = walkFn(fileName, nil, nil)
			if err != nil {
				return
			}
		}
	}
	return
}

type MockFile struct {
	bufferReadWriter *bufio.ReadWriter
	IsClosed         bool
}

func newMockFile() *MockFile {
	buffer := new(bytes.Buffer)
	bufferReader := bufio.NewReader(buffer)
	bufferWriter := bufio.NewWriter(buffer)
	return &MockFile{
		bufferReadWriter: bufio.NewReadWriter(bufferReader, bufferWriter),
		IsClosed:         false,
	}
}

func (mockFile *MockFile) Read(p []byte) (n int, err error) {
	return mockFile.bufferReadWriter.Read(p)
}

func (mockFile *MockFile) Write(p []byte) (n int, err error) {
	return mockFile.bufferReadWriter.Write(p)
}

func (mockFile *MockFile) Close() error {
	mockFile.IsClosed = true
	return nil
}
