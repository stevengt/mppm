package util

import (
	"fmt"
	"os/exec"
)

func ExecuteShellCommand(commandName string, args ...string) (err error) {
	return getShellCommandManager().ExecuteShellCommand(commandName, args...)
}

func ExecuteShellCommandAndReturnOutput(commandName string, args ...string) (stdout string, err error) {
	return getShellCommandManager().ExecuteShellCommandAndReturnOutput(commandName, args...)
}

var shellCommandManagerFactory ShellCommandManagerCreator = &LocalShellCommandManagerCreator{}
var shellCommandManager ShellCommandManager

func getShellCommandManager() ShellCommandManager {
	if shellCommandManager == nil {
		shellCommandManager = shellCommandManagerFactory.NewShellCommandManager()
	}
	return shellCommandManager
}

type ShellCommandManagerCreator interface {
	NewShellCommandManager() ShellCommandManager
}

type LocalShellCommandManagerCreator struct{}

func (localShellCommandManagerCreator *LocalShellCommandManagerCreator) NewShellCommandManager() ShellCommandManager {
	return &LocalShellCommandManager{}
}

type ShellCommandManager interface {
	ExecuteShellCommand(commandName string, args ...string) (err error)
	ExecuteShellCommandAndReturnOutput(commandName string, args ...string) (stdout string, err error)
}

type LocalShellCommandManager struct{}

func (localShellCommandManager *LocalShellCommandManager) ExecuteShellCommand(commandName string, args ...string) (err error) {
	stdout, err := localShellCommandManager.ExecuteShellCommandAndReturnOutput(commandName, args...)
	if len(stdout) > 0 {
		fmt.Println(stdout)
	}
	return
}

func (localShellCommandManager *LocalShellCommandManager) ExecuteShellCommandAndReturnOutput(commandName string, args ...string) (stdout string, err error) {
	stdoutAsByteArray, err := exec.Command(commandName, args...).CombinedOutput()
	if stdoutAsByteArray != nil && len(stdoutAsByteArray) > 0 {
		stdout = string(stdoutAsByteArray)
	}
	return
}
