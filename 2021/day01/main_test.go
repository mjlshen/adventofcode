package main

import "testing"

func TestNumIncreases(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "test.txt",
			expected: 7,
		},
	}

	for _, test := range tests {
		actual := numIncreases(test.input)
		if actual != test.expected {
			t.Errorf("numIncreases(%s) = %d, expected %d", test.input, actual, test.expected)
		}
	}
}

func TestNumIncreasesWindow(t *testing.T) {
	tests := []struct {
		input    string
		window   int
		expected int
	}{
		{
			input:    "test.txt",
			window:   3,
			expected: 5,
		},
	}

	for _, test := range tests {
		actual := numIncreasesWindow(test.input, test.window)
		if actual != test.expected {
			t.Errorf("numIncreasesWindow(%s, %d) = %d, expected %d", test.input, test.window, actual, test.expected)
		}
	}
}
