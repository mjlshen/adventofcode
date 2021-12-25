package main

import (
	"os"
	"testing"
)

func Test_SeafloorStable(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "test.txt",
			expected: 58,
		},
	}

	for _, test := range tests {
		actual := parseSeafloor(test.input).stable()
		if actual != test.expected {
			t.Errorf("expected %d, got %d", test.expected, actual)
		}
	}
}

func Test_SeafloorMove(t *testing.T) {
	tests := []struct {
		moves    int
		expected string
	}{
		{
			moves: 1,
			expected: `....>.>v.>
v.v>.>v.v.
>v>>..>v..
>>v>v>.>.v
.>v.v...v.
v>>.>vvv..
..v...>>..
vv...>>vv.
>.v.v..v.v`,
		},
		{
			moves: 58,
			expected: `..>>v>vv..
..v.>>vv..
..>>v>>vv.
..>>>>>vv.
v......>vv
v>v....>>v
vvv.....>>
>vv......>
.>v.vv.v..`,
		},
	}

	for _, test := range tests {
		sf := parseSeafloor("test.txt")
		for i := 0; i < test.moves; i++ {
			sf, _ = sf.move()
		}
		actual := sf.String()
		if actual != test.expected {
			t.Errorf("moves: %d\nexpected \n%s, got \n%s", test.moves, test.expected, actual)
		}
	}
}
func Test_SeafloorString(t *testing.T) {
	tests := []struct {
		path string
	}{
		{
			path: "test.txt",
		},
	}

	for _, test := range tests {
		expected, err := os.ReadFile(test.path)
		if err != nil {
			t.Error(err)
		}
		actual := parseSeafloor(test.path).String()
		if actual != string(expected) {
			t.Errorf("expected \n%s, got \n%s", expected, actual)
		}
	}
}
