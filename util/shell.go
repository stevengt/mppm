package util

import (
	"fmt"
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
	stdoutAsByteArray, err := exec.Command(commandName, args...).CombinedOutput()
	if stdoutAsByteArray != nil && len(stdoutAsByteArray) > 0 {
		stdout = string(stdoutAsByteArray)
	}
	return
}
