package main

import "testing"

func TestFeedInstructions(t *testing.T) {
	tests := []struct {
		path     string
		input    []int
		expected bool
	}{
		{
			path:     "test.txt",
			input:    []int{1, 1},
			expected: true,
		},
		{
			path:     "test.txt",
			input:    []int{1, 3},
			expected: false,
		},
	}

	for _, test := range tests {
		is := parseInstructions(test.path)
		actual := feedInstructions(is, test.input)
		if actual != test.expected {
			t.Errorf("feedInstructions(%v, %v) = %v, expected %v", test.path, test.input, actual, test.expected)
		}
	}
}
