package util_test

import (
	"bufio"
	"bytes"
	"io"
)

type MockFileSystemDelegater struct {
	Err       error
	FileNames []string
}

func (mockFileSystemDelegater *MockFileSystemDelegater) OpenFile(fileName string) (file io.ReadWriteCloser, err error) {
	file = newMockFile()
	err = mockFileSystemDelegater.Err
	return
}

func (mockFileSystemDelegater *MockFileSystemDelegater) CreateFile(fileName string) (file io.ReadWriteCloser, err error) {
	return mockFileSystemDelegater.OpenFile(fileName)
}

func (mockFileSystemDelegater *MockFileSystemDelegater) RemoveFile(fileName string) (err error) {
	err = mockFileSystemDelegater.Err
	return
}

func (mockFileSystemDelegater *MockFileSystemDelegater) CopyFile(sourceFileName string, targetFileName string) (err error) {
	err = mockFileSystemDelegater.Err
	return
}

func (mockFileSystemDelegater *MockFileSystemDelegater) GzipFile(fileName string) (err error) {
	err = mockFileSystemDelegater.Err
	return
}

func (mockFileSystemDelegater *MockFileSystemDelegater) GunzipFile(compressedFileName string) (err error) {
	err = mockFileSystemDelegater.Err
	return
}

func (mockFileSystemDelegater *MockFileSystemDelegater) GetAllFileNamesWithExtension(extension string) (fileNames []string, err error) {
	fileNames = mockFileSystemDelegater.FileNames
	err = mockFileSystemDelegater.Err
	return
}

type MockFile struct {
	bufferReadWriter *bufio.ReadWriter
}

func newMockFile() *MockFile {
	buffer := new(bytes.Buffer)
	bufferReader := bufio.NewReader(buffer)
	bufferWriter := bufio.NewWriter(buffer)
	return &MockFile{
		bufferReadWriter: bufio.NewReadWriter(bufferReader, bufferWriter),
	}
}

func (mockFile *MockFile) Read(p []byte) (n int, err error) {
	return mockFile.bufferReadWriter.Read(p)
}

func (mockFile *MockFile) Write(p []byte) (n int, err error) {
	return mockFile.bufferReadWriter.Write(p)
}

func (mockFile *MockFile) Close() error {
	return nil
}
