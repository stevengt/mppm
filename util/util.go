package util

import (
	"compress/gzip"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func ExecuteShellCommand(commandName string, args ...string) (err error) {
	stdout, err := ExecuteShellCommandAndReturnOutput(commandName, args...)
	if len(stdout) > 0 {
		fmt.Println(stdout)
	}
	return
}

func ExecuteShellCommandAndReturnOutput(commandName string, args ...string) (stdout string, err error) {
	stdoutAsByteArray, err := exec.Command(commandName, args...).CombinedOutput()
	if stdoutAsByteArray != nil && len(stdoutAsByteArray) > 0 {
		stdout = string(stdoutAsByteArray)
	}
	return
}

func CopyFile(sourceFileName string, targetFileName string) (err error) {

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

func GzipFile(fileName string) (err error) {

	uncompressedFile, err := os.Open(fileName)
	if err != nil {
		return
	}
	defer uncompressedFile.Close()

	compressedFileName := fileName + ".gz"
	compressedFile, err := CreateFile(compressedFileName)
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
	uncompressedFile, err := CreateFile(uncompressedFileName)
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

func CreateFile(fileName string) (file *os.File, err error) {

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

func GetAllFileNamesWithExtension(extension string) (fileNames []string, err error) {

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

func ExitWithError(err error) {
	ExitWithErrorMessage(err.Error())
}

func ExitWithErrorMessage(errorMessage string) {
	fmt.Println("ERROR: " + errorMessage)
	os.Exit(1)
}
