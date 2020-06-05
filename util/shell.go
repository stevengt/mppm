package util

import (
	"os/exec"
)

var ShellProxy ShellCommandDelegater = &shellProxy{}

func ExecuteShellCommand(commandName string, args ...string) (err error) {
	return ShellProxy.ExecuteShellCommand(commandName, args...)
}

func ExecuteShellCommandAndReturnOutput(commandName string, args ...string) (stdout string, err error) {
	return ShellProxy.ExecuteShellCommandAndReturnOutput(commandName, args...)
}

// ------------------------------------------------------------------------------

type ShellCommandDelegater interface {
	ExecuteShellCommand(commandName string, args ...string) (err error)
	ExecuteShellCommandAndReturnOutput(commandName string, args ...string) (stdout string, err error)
}

type shellProxy struct{}

func (proxy *shellProxy) ExecuteShellCommand(commandName string, args ...string) (err error) {
	stdout, err := proxy.ExecuteShellCommandAndReturnOutput(commandName, args...)
	if len(stdout) > 0 {
		Println(stdout)
	}
	return
}

func (proxy *shellProxy) ExecuteShellCommandAndReturnOutput(commandName string, args ...string) (stdout string, err error) {
	stdoutAsByteArray, err := exec.Command(commandName, args...).CombinedOutput()
	if stdoutAsByteArray != nil && len(stdoutAsByteArray) > 0 {
		stdout = string(stdoutAsByteArray)
	}
	return
}
