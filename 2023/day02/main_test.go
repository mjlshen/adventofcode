package main

import "testing"

func TestPossibleGames(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected int
	}{
		{
			name:     "Part 1",
			path:     "test.txt",
			expected: 8,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := possibleGames(test.path)
			if test.expected != actual {
				t.Errorf("expected: %d, got: %d", test.expected, actual)
			}
		})
	}
}

func TestMinimumBag(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected int
	}{
		{
			name:     "Part 2",
			path:     "test.txt",
			expected: 2286,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := minimumBag(test.path)
			if test.expected != actual {
				t.Errorf("expected: %d, got: %d", test.expected, actual)
			}
		})
	}
}
