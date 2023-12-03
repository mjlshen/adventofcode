package main

import (
	"fmt"
	"testing"
)

func TestBeaconExclusion(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		row      int
		expected int
	}{
		{
			name:     "Part 1",
			path:     "test.txt",
			row:      10,
			expected: 26,
		},
		{
			name:     "Part 2",
			path:     "test.txt",
			row:      20,
			expected: 56000011,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := beaconExclusion(test.path, test.row)
			if test.expected != actual {
				t.Errorf("expected %d, got %d", test.expected, actual)
			}
		})
	}
}

func TestGenerateOOR(t *testing.T) {
	sensors := map[coord]int{
		coord{0, 0}: 2,
	}
	actual := generateOOR(sensors, 20)
	fmt.Println(actual, len(actual))
}
