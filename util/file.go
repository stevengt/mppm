package util

import (
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var FileSystemProxy FileSystemDelegater = &fileSystemProxy{}

func OpenFile(fileName string) (file io.ReadWriteCloser, err error) {
	return FileSystemProxy.OpenFile(fileName)
}

func CreateFile(fileName string) (file io.ReadWriteCloser, err error) {
	return FileSystemProxy.CreateFile(fileName)
}

func RemoveFile(fileName string) (err error) {
	return FileSystemProxy.RemoveFile(fileName)
}

func CopyFile(sourceFileName string, targetFileName string) (err error) {

	source, err := FileSystemProxy.OpenFile(sourceFileName)
	if err != nil {
		return
	}
	defer source.Close()

	destination, err := FileSystemProxy.CreateFile(targetFileName)
	if err != nil {
		return
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return
}

func GzipFile(fileName string) (err error) {

	uncompressedFile, err := FileSystemProxy.OpenFile(fileName)
	if err != nil {
		return
	}
	defer uncompressedFile.Close()

	compressedFileName := fileName + ".gz"
	compressedFile, err := FileSystemProxy.CreateFile(compressedFileName)
	if err != nil {
		return
	}
	defer compressedFile.Close()

	gzipWriter := gzip.NewWriter(compressedFile)
	defer gzipWriter.Close()

	_, err = io.Copy(gzipWriter, uncompressedFile)
	if err != nil {
		return
	}

	return
}

func GunzipFile(compressedFileName string) (err error) {

	compressedFile, err := FileSystemProxy.OpenFile(compressedFileName)
	if err != nil {
		return
	}
	defer func() {
		compressedFile.Close()
		err = FileSystemProxy.RemoveFile(compressedFileName)
	}()

	gzipReader, err := gzip.NewReader(compressedFile)
	if err != nil {
		return
	}
	defer gzipReader.Close()

	uncompressedFileName := strings.TrimSuffix(compressedFileName, ".gz")
	uncompressedFile, err := FileSystemProxy.CreateFile(uncompressedFileName)
	if err != nil {
		return
	}
	defer uncompressedFile.Close()

	_, err = io.Copy(uncompressedFile, gzipReader)
	if err != nil {
		return
	}

	return
}

func GetAllFileNamesWithExtension(extension string) (fileNames []string, err error) {

	fileNames = make([]string, 0)

	err = FileSystemProxy.WalkFilePath(".", func(fileName string, fileInfo os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if strings.HasSuffix(fileName, "."+extension) {
			fileNames = append(fileNames, fileName)
		}
		return nil
	})

	if err != nil {
		return
	}

	return

}

type FileSystemDelegater interface {
	OpenFile(fileName string) (file io.ReadWriteCloser, err error)
	CreateFile(fileName string) (file io.ReadWriteCloser, err error)
	RemoveFile(fileName string) (err error)
	WalkFilePath(root string, walkFn filepath.WalkFunc) (err error)
}

type fileSystemProxy struct{}

func (proxy *fileSystemProxy) OpenFile(fileName string) (file io.ReadWriteCloser, err error) {
	file, err = os.Open(fileName)
	return
}

func (proxy *fileSystemProxy) CreateFile(fileName string) (file io.ReadWriteCloser, err error) {

	err = proxy.RemoveFile(fileName)
	if err != nil {
		return
	}

	file, err = os.Create(fileName)
	if err != nil {
		return
	}

	return
}

func (proxy *fileSystemProxy) RemoveFile(fileName string) (err error) {
	err = os.RemoveAll(fileName)
	return
}

func (proxy *fileSystemProxy) WalkFilePath(root string, walkFn filepath.WalkFunc) (err error) {
	err = filepath.Walk(root, walkFn)
	return
}
