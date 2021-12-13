package main

import "testing"

func TestNumDots(t *testing.T) {
	tests := []struct {
		input    string
		oneFold  bool
		expected int
	}{
		{
			input:    "test.txt",
			oneFold:  true,
			expected: 17,
		},
		{
			input:    "test.txt",
			oneFold:  false,
			expected: 16,
		},
	}

	for _, test := range tests {
		actual := numDots(test.input, test.oneFold)
		if actual != test.expected {
			t.Errorf("numDots(%s) expected %d, actual %d", test.input, test.expected, actual)
		}
	}
}
