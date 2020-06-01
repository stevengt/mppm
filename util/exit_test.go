package util_test

type MockExiter struct {
	WasExited bool
}

func (mockExiter *MockExiter) ExitWithError(err error) {
	mockExiter.WasExited = true
	return
}

func (mockExiter *MockExiter) ExitWithErrorMessage(errorMessage string) {
	mockExiter.WasExited = true
	return
}
