package util_test

type MockShellCommandDelegater struct {
	Stdout string
	Err    error
}

func (mockShellCommandDelegater *MockShellCommandDelegater) ExecuteShellCommand(commandName string, args ...string) {
	return
}

func (mockShellCommandDelegater *MockShellCommandDelegater) ExecuteShellCommandAndReturnOutput(commandName string, args ...string) (stdout string, err error) {
	stdout = mockShellCommandDelegater.Stdout
	err = mockShellCommandDelegater.Err
	return
}
