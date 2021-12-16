package main

import "testing"

func TestSumMinPath(t *testing.T) {
	tests := []struct {
		input    string
		fullCave bool
		expected int
	}{
		// {
		// 	input:    "test.txt",
		// 	fullCave: false,
		// 	expected: 40,
		// },
		{
			input:    "test.txt",
			fullCave: true,
			expected: 315,
		},
	}

	for _, test := range tests {
		actual := sumMinPath(test.input, test.fullCave)
		if actual != test.expected {
			t.Errorf("actual = %d, expected %d", actual, test.expected)
		}
	}
}

func TestManhattanDistance(t *testing.T) {
	tests := []struct {
		start    Coord
		end      Coord
		expected int
	}{
		{
			start:    Coord{0, 1},
			end:      Coord{1, 0},
			expected: 2,
		},
		{
			start:    Coord{1, 1},
			end:      Coord{0, 0},
			expected: 2,
		},
	}

	for _, test := range tests {
		actual := manhattanDistance(test.start, test.end)
		if actual != test.expected {
			t.Errorf("actual = %d, expected %d", actual, test.expected)
		}
	}
}
