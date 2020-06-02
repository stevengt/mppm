package utiltest

import "strings"

type MockShellCommandDelegater struct {
	OutputSequence         []*MockShellCommandOutput
	InputHistory           []string
	OutputHistory          []*MockShellCommandOutput
	curOutputSequenceIndex int
}

type MockShellCommandOutput struct {
	Stdout string
	Err    error
}

func NewMockShellCommandDelegater(outputSequence []*MockShellCommandOutput) *MockShellCommandDelegater {
	return &MockShellCommandDelegater{
		OutputSequence:         outputSequence,
		InputHistory:           make([]string, 0),
		OutputHistory:          make([]*MockShellCommandOutput, 0),
		curOutputSequenceIndex: 0,
	}
}

func (mockShellCommandDelegater *MockShellCommandDelegater) ExecuteShellCommand(commandName string, args ...string) (err error) {
	_, err = mockShellCommandDelegater.ExecuteShellCommandAndReturnOutput(commandName, args...)
	return
}

func (mockShellCommandDelegater *MockShellCommandDelegater) ExecuteShellCommandAndReturnOutput(commandName string, args ...string) (stdout string, err error) {

	inputArgs := append([]string{commandName}, args...)
	input := strings.Join(inputArgs, " ")
	mockShellCommandDelegater.InputHistory = append(mockShellCommandDelegater.InputHistory, input)

	outputSequenceIndex := mockShellCommandDelegater.curOutputSequenceIndex
	output := mockShellCommandDelegater.OutputSequence[outputSequenceIndex]

	stdout = output.Stdout
	err = output.Err

	mockShellCommandDelegater.OutputHistory = append(mockShellCommandDelegater.OutputHistory, output)
	mockShellCommandDelegater.curOutputSequenceIndex++

	return
}
