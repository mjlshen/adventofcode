package main

import (
	"testing"
)

func TestEnhanceImg(t *testing.T) {
	tests := []struct {
		path     string
		times    int
		expected int
	}{
		{
			path:     "test.txt",
			times:    2,
			expected: 35,
		},
		{
			path:     "test.txt",
			times:    50,
			expected: 3351,
		},
	}

	for _, test := range tests {
		actual := enhanceImg(test.path, test.times)
		if actual != test.expected {
			t.Errorf("expected %d, got %d", test.expected, actual)
		}
	}
}

func Test_TrenchImageString(t *testing.T) {
	tests := []struct {
		path     string
		expected string
	}{
		{
			path: "test.txt",
			expected: `#..#.
#....
##..#
..#..
..###
`,
		},
	}

	for _, test := range tests {
		_, tr := parse(test.path)
		actual := tr.String()
		if actual != test.expected {
			t.Errorf("expected %s, got %s", test.expected, actual)
		}
	}
}
