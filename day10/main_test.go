package main

import "testing"

func TestSyntaxError(t *testing.T) {
	tests := []struct {
		input      string
		completion bool
		expected   int
	}{
		{
			input:      "test.txt",
			completion: false,
			expected:   26397,
		},
		{
			input:      "test.txt",
			completion: true,
			expected:   288957,
		},
	}

	for _, test := range tests {
		actual := syntaxError(test.input, test.completion)
		if actual != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, actual)
		}
	}
}

func TestCompletionScore(t *testing.T) {
	tests := []struct {
		stack    string
		expected int
	}{
		{
			stack:    "[({([[{{",
			expected: 288957,
		},
		{
			stack:    "((((<{<{{",
			expected: 1480781,
		},
	}

	for _, test := range tests {
		actual := completionScore([]rune(test.stack))
		if actual != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, actual)
		}
	}
}
