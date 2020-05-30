package util

import (
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

var FileSystemProxy FileSystemDelegater = &fileSystemProxy{}

func CreateFile(fileName string) (file *os.File, err error) {
	return FileSystemProxy.CreateFile(fileName)
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
	CreateFile(fileName string) (file *os.File, err error)
	CopyFile(sourceFileName string, targetFileName string) (err error)
	GzipFile(fileName string) (err error)
	GunzipFile(compressedFileName string) (err error)
	GetAllFileNamesWithExtension(extension string) (fileNames []string, err error)
}

type fileSystemProxy struct{}

func (proxy *fileSystemProxy) CreateFile(fileName string) (file *os.File, err error) {

	err = os.RemoveAll(fileName)
	if err != nil {
		return
	}

	file, err = os.Create(fileName)
	if err != nil {
		return
	}

	return
}

func (proxy *fileSystemProxy) CopyFile(sourceFileName string, targetFileName string) (err error) {

	source, err := os.Open(sourceFileName)
	if err != nil {
		return
	}
	defer source.Close()

	destination, err := os.Create(targetFileName)
	if err != nil {
		return
	}
	defer destination.Close()

	_, err = io.Copy(destination, source)
	return
}

func (proxy *fileSystemProxy) GzipFile(fileName string) (err error) {

	uncompressedFile, err := os.Open(fileName)
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

	compressedFile, err := os.Open(compressedFileName)
	if err != nil {
		return
	}
	defer func() {
		compressedFile.Close()
		err = os.RemoveAll(compressedFileName)
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
