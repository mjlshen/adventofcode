package main

import "testing"

func TestDeterministicDie(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "test.txt",
			expected: 739785,
		},
	}

	for _, test := range tests {
		actual := deterministicDie(test.input)
		if actual != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, actual)
		}
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "test.txt",
			expected: 444356092776315,
		},
	}

	for _, test := range tests {
		actual := part2(test.input)
		if actual != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, actual)
		}
	}
}
