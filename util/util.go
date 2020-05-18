package util

import (
	"fmt"
	"os/exec"
)

func ExecuteShellCommand(commandName string, args ...string) (err error) {
	cmd := exec.Command(commandName, args...)
	stdout, err := cmd.Output()

	if stdout != nil {
		fmt.Println(string(stdout))
	}
	if err != nil {
		fmt.Println(err.Error())
	}

	return
}