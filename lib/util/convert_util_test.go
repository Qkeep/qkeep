package util

import "testing"

func TestStringToInterface(t *testing.T) {
	// Define test cases
	testCases := []struct {
		input    string
		expected interface{}
	}{
		{"hello", "hello"},
		{"123", "123"},
		{"", ""},
	}

	// Run tests
	for _, tc := range testCases {
		actual := StringToInterface(tc.input)
		if actual != tc.expected {
			t.Errorf("Expected %v, but got %v for input %s", tc.expected, actual, tc.input)
		}
	}
}

func TestInterfaceToString(t *testing.T) {
	// Define test cases
	testCases := []struct {
		input    interface{}
		expected string
	}{
		{"hello", "hello"},
		{"123", "123"},
		{"", ""},
	}

	// Run tests
	for _, tc := range testCases {
		actual := InterfaceToString(tc.input)
		if actual != tc.expected {
			t.Errorf("Expected %s, but got %s for input %v", tc.expected, actual, tc.input)
		}
	}
}

func TestSnakeCaseToUpperCamelCase(t *testing.T) {

	testCases := &[]struct {
		input  string
		output string
	}{
		{"hello_world", "HelloWorld"},
		{"hello_world_", "HelloWorld"},
		{"_hello_world", "HelloWorld"},
		{"hello_world_123", "HelloWorld123"},
	}

	for _, tc := range *testCases {
		actual := SnakeCaseToUpperCamelCase(tc.input)
		if actual != tc.output {
			t.Errorf("Expected %s, but got %s for input %s", tc.output, actual, tc.input)
		}
	}

}

func TestSplitBy(t *testing.T) {
	testCases := &[]struct {
		input  string
		by     string
		output []string
	}{
		{"hello_world", "_", []string{"hello", "world"}},
		{"hello_world_", "_", []string{"hello", "world", ""}},
		{"_hello_world", "_", []string{"", "hello", "world"}},
		{"hello_world_123", "_", []string{"hello", "world", "123"}},
	}

	for _, tc := range *testCases {
		actual := SplitBy(tc.input, tc.by)
		if len(actual) != len(tc.output) {
			t.Errorf("Expected %v, but got %v for input %s", tc.output, actual, tc.input)
		}
		for i, v := range actual {
			if v != tc.output[i] {
				t.Errorf("Expected %v, but got %v for input %s", tc.output, actual, tc.input)
			}
		}
	}
}
