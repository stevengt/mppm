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

func RenameFile(fileName string, newFileName string) (err error) {
	return FileSystemProxy.RenameFile(fileName, newFileName)
}

func RemoveFile(fileName string) (err error) {
	return FileSystemProxy.RemoveFile(fileName)
}

func UserHomeDir() (string, error) {
	return FileSystemProxy.UserHomeDir()
}

func JoinFilePath(elem ...string) string {
	return FileSystemProxy.JoinFilePath(elem...)
}

func DoesFileExist(filePath string) bool {
	return FileSystemProxy.DoesFileExist(filePath)
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

	// Copy the uncompressedFile contents to the gzipWriter, then immediately Close() the gzipWriter.
	// This flushes all compressed contents and the GZIP footer to the compressedFile without closing
	// the compressedFile. If gzipWriter.Close() is deferred, the GZIP footer might not be written.
	_, err = io.Copy(gzipWriter, uncompressedFile)
	gzipWriter.Close()
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
	defer compressedFile.Close()

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

	err = FileSystemProxy.RemoveFile(compressedFileName)
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
		fileNames = nil
		return
	}

	return

}

// ------------------------------------------------------------------------------

type FileSystemDelegater interface {
	OpenFile(fileName string) (file io.ReadWriteCloser, err error)
	CreateFile(fileName string) (file io.ReadWriteCloser, err error)
	RenameFile(fileName string, newFileName string) (err error)
	RemoveFile(fileName string) (err error)
	WalkFilePath(root string, walkFn filepath.WalkFunc) (err error)
	UserHomeDir() (string, error)
	JoinFilePath(elem ...string) string
	DoesFileExist(filePath string) bool
}

type fileSystemProxy struct{}

func (proxy *fileSystemProxy) OpenFile(fileName string) (file io.ReadWriteCloser, err error) {
	file, err = os.Open(fileName)
	return
}

func (proxy *fileSystemProxy) CreateFile(fileName string) (file io.ReadWriteCloser, err error) {

	file, err = os.Create(fileName)
	if err != nil {
		return
	}

	return
}

func (proxy *fileSystemProxy) RenameFile(fileName string, newFileName string) (err error) {

	err = os.Rename(fileName, newFileName)
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

func (proxy *fileSystemProxy) UserHomeDir() (string, error) {
	return os.UserHomeDir()
}

func (proxy *fileSystemProxy) JoinFilePath(elem ...string) string {
	return filepath.Join(elem...)
}

func (proxy *fileSystemProxy) DoesFileExist(filePath string) bool {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return false
	}
	return true
}
