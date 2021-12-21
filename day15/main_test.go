package main

import (
	"os"
	"testing"
)

func Test_CaveDjikstra(t *testing.T) {
	tests := []struct {
		input    string
		fullCave bool
		expected int
	}{
		{
			input:    "test.txt",
			fullCave: false,
			expected: 40,
		},
		{
			input:    "test.txt",
			fullCave: true,
			expected: 315,
		},
	}

	for _, test := range tests {
		actual := parseCave(test.input, test.fullCave).djikstra()
		if actual != test.expected {
			t.Errorf("actual = %d, expected %d", actual, test.expected)
		}
	}
}

func Test_CaveString(t *testing.T) {
	tests := []struct {
		input string
	}{
		{
			input: "test.txt",
		},
	}

	for _, test := range tests {
		expected, _ := os.ReadFile(test.input)
		actual := parseCave(test.input, false)
		if string(expected) != actual.String() {
			t.Errorf("actual = %s, expected %s", string(expected), actual.String())
		}
	}
}
