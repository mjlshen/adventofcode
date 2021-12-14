package main

import "testing"

func TestCount1478(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "test.txt",
			expected: 26,
		},
	}

	for _, test := range tests {
		actual := count1478(test.input)
		if actual != test.expected {
			t.Errorf("Count1478(%s) = %d, expected %d", test.input, actual, test.expected)
		}
	}
}

func TestSumOutput(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "test.txt",
			expected: 61229,
		},
	}

	for _, test := range tests {
		actual := sumOutput(test.input)
		if actual != test.expected {
			t.Errorf("SumOutput(%s) = %d, expected %d", test.input, actual, test.expected)
		}
	}
}
