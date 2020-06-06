package utiltest

import (
	"fmt"
)

type MockWritePrinter struct {
	OutputContents []byte
}

func NewMockWritePrinter() *MockWritePrinter {
	return &MockWritePrinter{
		OutputContents: make([]byte, 0),
	}
}

func (mockWritePrinter *MockWritePrinter) GetOutputContentsAsString() string {
	return string(mockWritePrinter.OutputContents)
}

func (mockWritePrinter *MockWritePrinter) Print(v ...interface{}) {
	outputAsBytes := []byte(fmt.Sprint(v...))
	mockWritePrinter.OutputContents = append(
		mockWritePrinter.OutputContents,
		outputAsBytes...,
	)
}

func (mockWritePrinter *MockWritePrinter) Printf(format string, v ...interface{}) {
	outputAsBytes := []byte(fmt.Sprintf(format, v...))
	mockWritePrinter.OutputContents = append(
		mockWritePrinter.OutputContents,
		outputAsBytes...,
	)
}

func (mockWritePrinter *MockWritePrinter) Println(v ...interface{}) {
	outputAsBytes := []byte(fmt.Sprintln(v...))
	mockWritePrinter.OutputContents = append(
		mockWritePrinter.OutputContents,
		outputAsBytes...,
	)
}

func (mockWritePrinter *MockWritePrinter) Write(p []byte) (n int, err error) {
	mockWritePrinter.OutputContents = append(
		mockWritePrinter.OutputContents,
		p...,
	)
	return len(p), nil
}
