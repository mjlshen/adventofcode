package main

import (
	"fmt"
	"testing"
)

func TestPart1(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "test.txt",
			expected: 4140,
		},
	}

	for _, test := range tests {
		actual := part1(test.input)
		if actual != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, actual)
		}
	}
}

func TestPart2(t *testing.T) {
	tests := []struct {
		input    string
		expected int
	}{
		{
			input:    "test.txt",
			expected: 3993,
		},
	}

	for _, test := range tests {
		actual := part2(test.input)
		if actual != test.expected {
			t.Errorf("Expected %d, got %d", test.expected, actual)
		}
	}
}

func Test_PairSum(t *testing.T) {
	tests := []struct {
		a        string
		b        string
		expected string
	}{
		{
			a:        "[1,2]",
			b:        "[[3,4],5]",
			expected: "[[1,2],[[3,4],5]]",
		},
		{
			a:        "[[[[4,3],4],4],[7,[[8,4],9]]]",
			b:        "[1,1]",
			expected: "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
		},
	}

	for _, test := range tests {
		pair := newPair(test.a)
		fmt.Println(test.b)
		actual := pair.sum(newPair(test.b))
		if actual.String() != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, pair.String())
		}
	}
}

func Test_PairReduce(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "[[[[[9,8],1],2],3],4]",
			expected: "[[[[0,9],2],3],4]",
		},
		{
			input:    "[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]",
			expected: "[[3,[2,[8,0]]],[9,[5,[7,0]]]]",
		},
		{
			input:    "[[[[[4,3],4],4],[7,[[8,4],9]]],[1,1]]",
			expected: "[[[[0,7],4],[[7,8],[6,0]]],[8,1]]",
		},
	}

	for _, test := range tests {
		pair := newPair(test.input)
		pair.reduce()
		if pair.String() != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, pair.String())
		}
	}
}

func Test_PairExplode(t *testing.T) {
	tests := []struct {
		input    string
		expected string
		left     int
		right    int
	}{
		{
			input:    "[[[[[9,8],1],2],3],4]",
			expected: "[[[[0,9],2],3],4]",
			left:     4,
			right:    0,
		},
		{
			input:    "[[6,[5,[4,[3,2]]]],1]",
			expected: "[[6,[5,[7,0]]],3]",
			left:     1,
			right:    3,
		},
		{
			input:    "[[3,[2,[1,[7,3]]]],[6,[5,[4,[3,2]]]]]",
			expected: "[[3,[2,[8,0]]],[9,[5,[4,[3,2]]]]]",
			left:     1,
			right:    3,
		},
	}

	for _, test := range tests {
		pair := newPair(test.input)
		temp := pair
		for i := 0; i < test.left; i++ {
			temp = temp.left
			fmt.Println(temp)
		}
		for i := 0; i < test.right; i++ {
			temp = temp.right
			fmt.Println(temp)
		}
		temp.explode()
		fmt.Println("hi?")
		if pair.String() != test.expected {
			t.Errorf("Expected %s, got %s", test.expected, pair.String())
		}
	}
}

func Test_PairString(t *testing.T) {
	tests := []struct {
		input string
	}{
		{
			input: "[[[0,[5,8]],[[1,7],[9,6]]],[[4,[1,2]],[[1,4],2]]]",
		},
		{
			input: "[6,[[[6,2],[5,6]],[[7,6],[4,7]]]]",
		},
	}

	for _, test := range tests {
		pair := newPair(test.input)
		if pair.String() != test.input {
			t.Errorf("Expected %s, got %s", test.input, pair.String())
		}
	}
}
