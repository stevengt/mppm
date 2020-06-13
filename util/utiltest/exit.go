package utiltest

import (
	"errors"

	"github.com/stevengt/mppm/util"
)

func InitializeAndReturnNewMockExiter() *MockExiter {
	mockExiter := NewMockExiter()
	util.CurrentProcessExiter = mockExiter
	return mockExiter
}

// ------------------------------------------------------------------------------

type MockExiter struct {
	WasExited bool
	Error     error
}

func NewMockExiter() *MockExiter {
	return &MockExiter{
		WasExited: false,
	}
}

func (mockExiter *MockExiter) Init() {
	util.CurrentProcessExiter = mockExiter
}

func (mockExiter *MockExiter) ExitWithError(err error) {
	mockExiter.WasExited = true
	mockExiter.Error = err
	return
}

func (mockExiter *MockExiter) ExitWithErrorMessage(errorMessage string) {
	mockExiter.WasExited = true
	mockExiter.Error = errors.New(errorMessage)
	return
}
