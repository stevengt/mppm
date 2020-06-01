package util_test

type MockExiter struct {
	WasExited    bool
	ErrorMessage string
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
