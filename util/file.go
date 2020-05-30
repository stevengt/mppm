package util

import (
	"compress/gzip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func CreateFile(fileName string) (file *os.File, err error) {
	return getFileManager().CreateFile(fileName)
}

func CopyFile(sourceFileName string, targetFileName string) (err error) {
	return getFileManager().CopyFile(sourceFileName, targetFileName)
}

func GzipFile(fileName string) (err error) {
	return getFileManager().GzipFile(fileName)
}

func GunzipFile(compressedFileName string) (err error) {
	return getFileManager().GunzipFile(compressedFileName)
}

func GetAllFileNamesWithExtension(extension string) (fileNames []string, err error) {
	return getFileManager().GetAllFileNamesWithExtension(extension)
}

var fileManagerFactory FileManagerCreator = &LocalFileSystemManagerCreator{}
var fileManager FileManager

func getFileManager() FileManager {
	if fileManager == nil {
		fileManager = fileManagerFactory.NewFileManager()
	}
	return fileManager
}

type FileManagerCreator interface {
	NewFileManager() FileManager
}

type LocalFileSystemManagerCreator struct{}

func (localFileSystemManagerCreator *LocalFileSystemManagerCreator) NewFileManager() FileManager {
	return &LocalFileSystemManager{}
}

type FileManager interface {
	CreateFile(fileName string) (file *os.File, err error)
	CopyFile(sourceFileName string, targetFileName string) (err error)
	GzipFile(fileName string) (err error)
	GunzipFile(compressedFileName string) (err error)
	GetAllFileNamesWithExtension(extension string) (fileNames []string, err error)
}

type LocalFileSystemManager struct{}

func (localFileSystemManager *LocalFileSystemManager) CreateFile(fileName string) (file *os.File, err error) {

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

func (localFileSystemManager *LocalFileSystemManager) CopyFile(sourceFileName string, targetFileName string) (err error) {

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

func (localFileSystemManager *LocalFileSystemManager) GzipFile(fileName string) (err error) {

	uncompressedFile, err := os.Open(fileName)
	if err != nil {
		return
	}
	defer uncompressedFile.Close()

	compressedFileName := fileName + ".gz"
	compressedFile, err := localFileSystemManager.CreateFile(compressedFileName)
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

func (localFileSystemManager *LocalFileSystemManager) GunzipFile(compressedFileName string) (err error) {

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
	uncompressedFile, err := localFileSystemManager.CreateFile(uncompressedFileName)
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

func (localFileSystemManager *LocalFileSystemManager) GetAllFileNamesWithExtension(extension string) (fileNames []string, err error) {

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
