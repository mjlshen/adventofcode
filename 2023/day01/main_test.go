package main

import "testing"

func TestMaxCalories(t *testing.T) {
	tests := []struct {
		name     string
		path     string
		words    bool
		expected int
	}{
		{
			name:     "Part 1",
			path:     "test.txt",
			words:    false,
			expected: 142,
		},
		{
			name:     "Part 2",
			path:     "test2.txt",
			words:    true,
			expected: 281,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := calibrate(test.path, test.words)
			if test.expected != actual {
				t.Errorf("expected: %d, got: %d", test.expected, actual)
			}
		})
	}

}
