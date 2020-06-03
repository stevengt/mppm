package utiltest

import "github.com/stevengt/mppm/util"

type MockExiter struct {
	WasExited    bool
	ErrorMessage string
}

func InitializeAndReturnNewMockExiter() *MockExiter {
	mockExiter := NewMockExiter()
	util.CurrentProcessExiter = mockExiter
	return mockExiter
}

func NewMockExiter() *MockExiter {
	return &MockExiter{
		WasExited: false,
	}
}

func (mockExiter *MockExiter) ExitWithError(err error) {
	mockExiter.WasExited = true
	mockExiter.ErrorMessage = err.Error()
	return
}

func (mockExiter *MockExiter) ExitWithErrorMessage(errorMessage string) {
	mockExiter.WasExited = true
	mockExiter.ErrorMessage = errorMessage
	return
}
