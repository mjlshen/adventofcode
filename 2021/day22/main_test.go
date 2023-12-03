package main

import "testing"

func TestCountCubes(t *testing.T) {
	tests := []struct {
		input    string
		trim     bool
		expected int
	}{
		{
			input:    "test.txt",
			trim:     true,
			expected: 39,
		},
		{
			input:    "test2.txt",
			trim:     true,
			expected: 590784,
		},
		{
			input:    "test3.txt",
			trim:     true,
			expected: 474140,
		},
		{
			input:    "test3.txt",
			trim:     false,
			expected: 2758514936282235,
		},
	}

	for _, test := range tests {
		actual := countCubes(test.input, test.trim)
		if actual != test.expected {
			t.Errorf("actual = %d, expected %d", actual, test.expected)
		}
	}
}
