package main

import "testing"

func TestPolymerInsertion(t *testing.T) {
	tests := []struct {
		input    string
		steps    int
		expected int
	}{
		{
			input:    "test.txt",
			steps:    10,
			expected: 1588,
		},
		{
			input:    "test.txt",
			steps:    40,
			expected: 2188189693529,
		},
	}

	for _, test := range tests {
		actual := polymerInsertion(test.input, test.steps)
		if actual != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, actual)
		}
	}
}
