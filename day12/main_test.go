package main

import "testing"

func TestNumPaths(t *testing.T) {
	tests := []struct {
		input      string
		smallTwice bool
		expected   int
	}{
		{
			input:      "test.txt",
			smallTwice: false,
			expected:   10,
		},
		{
			input:      "test2.txt",
			smallTwice: false,
			expected:   19,
		},
		{
			input:      "test3.txt",
			smallTwice: false,
			expected:   226,
		},
		{
			input:      "test.txt",
			smallTwice: true,
			expected:   36,
		},
		{
			input:      "test2.txt",
			smallTwice: true,
			expected:   103,
		},
		{
			input:      "test3.txt",
			smallTwice: true,
			expected:   3509,
		},
	}

	for _, test := range tests {
		actual := numPaths(test.input, test.smallTwice)
		if actual != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, actual)
		}
	}
}
