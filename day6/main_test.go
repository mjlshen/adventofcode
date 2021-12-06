package main

import "testing"

func TestModelLanternfish(t *testing.T) {
	tests := []struct {
		input    string
		days     int
		expected int
	}{
		{
			input:    "test.txt",
			days:     18,
			expected: 26,
		},
		{
			input:    "test.txt",
			days:     80,
			expected: 5934,
		},
	}

	for _, test := range tests {
		actual := modelLanternfish(test.input, test.days)
		if actual != test.expected {
			t.Errorf("ModelLanternfish(%q, %d) = %d, expected %d", test.input, test.days, actual, test.expected)
		}
	}
}
