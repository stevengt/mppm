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
	return FileSystemProxy.CopyFile(sourceFileName, targetFileName)
}

func GzipFile(fileName string) (err error) {
	return FileSystemProxy.GzipFile(fileName)
}

func GunzipFile(compressedFileName string) (err error) {
	return FileSystemProxy.GunzipFile(compressedFileName)
}

func GetAllFileNamesWithExtension(extension string) (fileNames []string, err error) {
	return FileSystemProxy.GetAllFileNamesWithExtension(extension)
}

type FileSystemDelegater interface {
	OpenFile(fileName string) (file io.ReadWriteCloser, err error)
	CreateFile(fileName string) (file io.ReadWriteCloser, err error)
	RemoveFile(fileName string) (err error)
	CopyFile(sourceFileName string, targetFileName string) (err error)
	GzipFile(fileName string) (err error)
	GunzipFile(compressedFileName string) (err error)
	GetAllFileNamesWithExtension(extension string) (fileNames []string, err error)
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

func (proxy *fileSystemProxy) CopyFile(sourceFileName string, targetFileName string) (err error) {

	source, err := proxy.OpenFile(sourceFileName)
	if err != nil {
		return
	}
	defer source.Close()

	destination, err := proxy.CreateFile(targetFileName)
	if err != nil {
		return
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return
}

func (proxy *fileSystemProxy) GzipFile(fileName string) (err error) {

	uncompressedFile, err := proxy.OpenFile(fileName)
	if err != nil {
		return
	}
	defer uncompressedFile.Close()

	compressedFileName := fileName + ".gz"
	compressedFile, err := proxy.CreateFile(compressedFileName)
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

func (proxy *fileSystemProxy) GunzipFile(compressedFileName string) (err error) {

	compressedFile, err := proxy.OpenFile(compressedFileName)
	if err != nil {
		return
	}
	defer func() {
		compressedFile.Close()
		err = proxy.RemoveFile(compressedFileName)
	}()

	gzipReader, err := gzip.NewReader(compressedFile)
	if err != nil {
		return
	}
	defer gzipReader.Close()

	uncompressedFileName := strings.TrimSuffix(compressedFileName, ".gz")
	uncompressedFile, err := proxy.CreateFile(uncompressedFileName)
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

func (proxy *fileSystemProxy) GetAllFileNamesWithExtension(extension string) (fileNames []string, err error) {

	fileNames = make([]string, 0)

	err = filepath.Walk(".", func(fileName string, fileInfo os.FileInfo, err error) error {
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
