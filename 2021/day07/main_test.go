package main

import "testing"

func TestAlignCrabs(t *testing.T) {
	tests := []struct {
		input    string
		exp      bool
		expected int
	}{
		{
			input:    "test.txt",
			exp:      false,
			expected: 37,
		},
		{
			input:    "test.txt",
			exp:      true,
			expected: 168,
		},
	}

	for _, test := range tests {
		actual := alignCrabs(test.input, test.exp)
		if actual != test.expected {
			t.Errorf("actual %d, expected %d", actual, test.expected)
		}
	}
}
