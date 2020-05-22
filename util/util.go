package util

import (
	"fmt"
	"io"
	"os"
	"os/exec"
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
	cmdArgsArray := append([]string{commandName}, args...)
	stdoutAsByteArray, err := exec.Command("/usr/bin/env", cmdArgsArray...).CombinedOutput()
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

	compressedFileName := fileName + ".gz"
	err = os.RemoveAll(compressedFileName)
	if err != nil {
		return
	}

	err = ExecuteShellCommand("gzip", fileName)
	if err != nil {
		return
	}

	return
}

func GunzipFile(fileName string) (err error) {

	uncompressedFileName := strings.TrimSuffix(fileName, ".gz")
	err = os.RemoveAll(uncompressedFileName)
	if err != nil {
		return
	}

	err = ExecuteShellCommand("gunzip", fileName)
	if err != nil {
		return
	}

	return
}

func GetAllFileNamesWithExtension(extension string) (fileNames []string, err error) {
	fileNames = make([]string, 0)
	stdout, err := ExecuteShellCommandAndReturnOutput("find", ".", "-name", "*."+extension)
	if err == nil {
		stdoutLines := strings.Split(stdout, "\n")
		for i := 0; i < len(stdoutLines); i++ {
			line := stdoutLines[i]
			line = strings.Trim(line, " \n")
			line = strings.TrimPrefix(line, "./")
			if len(line) > 0 {
				fileNames = append(fileNames, line)
			}
		}
	}
	return
}
