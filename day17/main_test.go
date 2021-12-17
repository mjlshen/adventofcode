package main

import "testing"

func TestMaxYHeight(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "test.txt",
			expected: 45,
		},
	}

	for _, test := range tests {
		actual := maxYHeight(test.input)
		if actual != test.expected {
			t.Errorf("actual == %v, expected %v", actual, test.expected)
		}
	}
}

func TestCountVelocities(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "test.txt",
			expected: 112,
		},
	}

	for _, test := range tests {
		actual := countVelocities(test.input)
		if actual != test.expected {
			t.Errorf("actual == %v, expected %v", actual, test.expected)
		}
	}
}
