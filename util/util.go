package util

import (
	"fmt"
	"io"
	"os"
	"os/exec"
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
