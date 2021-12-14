package main

import "testing"

func TestDive(t *testing.T) {
	tests := []struct {
		input    string
		aim      bool
		expected int
	}{
		{
			input:    "test.txt",
			aim:      false,
			expected: 150,
		},
		{
			input:    "test.txt",
			aim:      true,
			expected: 900,
		},
	}

	for _, test := range tests {
		actual := dive(test.input, test.aim)
		if actual != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, actual)
		}
	}
}
