package util_test

type MockExiter struct{}

func (mockExiter *MockExiter) ExitWithError(err error) {
	return
}

func (mockExiter *MockExiter) ExitWithErrorMessage(errorMessage string) {
	return
}
