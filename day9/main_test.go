package main

import "testing"

func TestScoreHeightmap(t *testing.T) {
	tests := []struct {
		input    string
		basin    bool
		expected int
	}{
		{
			input:    "test.txt",
			basin:    false,
			expected: 15,
		},
		{
			input:    "test.txt",
			basin:    true,
			expected: 1134,
		},
	}

	for _, test := range tests {
		actual := scoreHeightmap(test.input, test.basin)
		if actual != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, actual)
		}
	}
}
