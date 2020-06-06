package utiltest

import "fmt"

type MockPrinter struct {
	OutputHistory []string
}

func NewMockPrinter() *MockPrinter {
	return &MockPrinter{
		OutputHistory: make([]string, 0),
	}
}

func (mockPrinter *MockPrinter) Print(v ...interface{}) {
	mockPrinter.OutputHistory = append(
		mockPrinter.OutputHistory,
		fmt.Sprint(v...),
	)
}

func (mockPrinter *MockPrinter) Printf(format string, v ...interface{}) {
	mockPrinter.OutputHistory = append(
		mockPrinter.OutputHistory,
		fmt.Sprintf(format, v...),
	)
}

func (mockPrinter *MockPrinter) Println(v ...interface{}) {
	mockPrinter.OutputHistory = append(
		mockPrinter.OutputHistory,
		fmt.Sprintln(v...),
	)
}
