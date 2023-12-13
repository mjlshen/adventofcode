package main

import "testing"

func TestWinningNumbers(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected int
	}{
		{
			name:     "Part 1",
			path:     "test.txt",
			expected: 13,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := winningNumbers(test.path)
			if test.expected != actual {
				t.Errorf("expected: %d, got: %d", test.expected, actual)
			}
		})
	}
}

func TestTotalScratchcards(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected int
	}{
		{
			name:     "Part 2",
			path:     "test.txt",
			expected: 30,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := totalScratchcards(test.path)
			if test.expected != actual {
				t.Errorf("expected: %d, got: %d", test.expected, actual)
			}
		})
	}
}
