package main

import (
	"testing"
)

func TestBlizzardBasin(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected int
	}{
		{
			name:     "Part 1",
			path:     "test.txt",
			expected: 110,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := blizzardBasin(test.path)
			if test.expected != actual {
				t.Errorf("expected %d, got %d", test.expected, actual)
			}
		})
	}
}
