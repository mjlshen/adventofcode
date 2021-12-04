package main

import "testing"

func TestFirstBingo(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "test.txt",
			expected: 4512,
		},
	}

	for _, test := range tests {
		actual := firstBingo(test.input)
		if actual != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, actual)
		}
	}
}

func TestLastBingo(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "test.txt",
			expected: 1924,
		},
	}

	for _, test := range tests {
		actual := lastBingo(test.input)
		if actual != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, actual)
		}
	}
}

func TestAllTrue(t *testing.T) {
	tests := []struct {
		input    []bool
		expected bool
	}{
		{
			input:    []bool{},
			expected: false,
		},
		{
			input:    []bool{true, true, true, true, true},
			expected: true,
		},
	}

	for _, test := range tests {
		actual := allTrue(test.input)
		if actual != test.expected {
			t.Errorf("Expected %t, got %t", test.expected, actual)
		}
	}
}
