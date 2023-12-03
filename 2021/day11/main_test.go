package main

import "testing"

func TestNumFlashes(t *testing.T) {
	tests := []struct {
		input    string
		steps    int
		expected int
	}{
		{
			input:    "test.txt",
			steps:    10,
			expected: 204,
		},
		{
			input:    "test.txt",
			steps:    100,
			expected: 1656,
		},
	}

	for _, test := range tests {
		actual := numFlashes(test.input, test.steps)
		if actual != test.expected {
			t.Errorf("Expected %d, but got %d", test.expected, actual)
		}
	}
}

func TestAllFlash(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "test.txt",
			expected: 195,
		},
	}

	for _, test := range tests {
		actual := allFlash(test.input)
		if actual != test.expected {
			t.Errorf("Expected %d, but got %d", test.expected, actual)
		}
	}
}
