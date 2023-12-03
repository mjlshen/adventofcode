package main

import (
	"testing"
)

func TestPyroclasticFlow(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		expected int
	}{
		{
			name:     "Part 1",
			path:     "test.txt",
			expected: 3068,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := pyroclasticFlow(test.path)
			if test.expected != actual {
				t.Errorf("expected %d, got %d", test.expected, actual)
			}
		})
	}
}
