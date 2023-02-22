package util

import (
	"errors"
	"fmt"
	"testing"
)

func TestCheckPanic(t *testing.T) {
	// Define test cases
	testCases := []struct {
		err      error
		expected string
	}{
		{nil, ""}, // no error
		{errors.New("test error"), "test error"},
	}

	// Run tests
	for _, tc := range testCases {
		// Capture output to check for panic message
		var output string
		func() {
			defer func() {
				if r := recover(); r != nil {
					output = fmt.Sprintf("%v", r)
				}
			}()
			CheckPanic(tc.err)
		}()

		// Check that the panic message (if any) matches the expected value
		if output != tc.expected {
			t.Errorf("Expected panic message %q, but got %q", tc.expected, output)
		}
	}
}
