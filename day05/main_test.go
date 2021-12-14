package main

import (
	"testing"
)

func TestHydroVents(t *testing.T) {
	tests := []struct {
		input    string
		diag     bool
		expected int
	}{
		{
			input:    "test.txt",
			diag:     false,
			expected: 5,
		},
		{
			input:    "test.txt",
			diag:     true,
			expected: 12,
		},
	}

	for _, test := range tests {
		actual := hydroVents(test.input, test.diag)
		if actual != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, actual)
		}
	}
}

func TestGenerateLineCoords(t *testing.T) {
	tests := []struct {
		start             *Coord
		end               *Coord
		diag              bool
		expectedNumCoords int
	}{
		{
			start:             &Coord{x: 0, y: 0},
			end:               &Coord{x: 3, y: 3},
			diag:              false,
			expectedNumCoords: 0,
		},
		{
			start:             &Coord{x: 0, y: 0},
			end:               &Coord{x: 3, y: 3},
			diag:              true,
			expectedNumCoords: 4,
		},
		{
			start:             &Coord{x: 0, y: 0},
			end:               &Coord{x: 3, y: 0},
			diag:              false,
			expectedNumCoords: 4,
		},
	}

	for _, test := range tests {
		actual := generateLineCoords(test.start, test.end, test.diag)
		if len(actual) != test.expectedNumCoords {
			t.Errorf("Expected %d, got %d\n%v\n", test.expectedNumCoords, len(actual), actual)
		}
	}
}
