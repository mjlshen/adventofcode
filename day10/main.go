package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	fmt.Printf("Part 1: %d\n", syntaxError("input.txt", false))
	fmt.Printf("Part 2: %d\n", syntaxError("input.txt", true))
}

func syntaxError(path string, completion bool) int {
	var (
		score  int
		scores []int
	)

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lineScore, leftoverStack := syntaxErrorScore(scanner.Text())
		if !completion {
			score += lineScore
		} else {
			if len(leftoverStack) > 0 {
				scores = append(scores, completionScore(leftoverStack))
			}
		}
	}

	if !completion {
		return score
	} else {
		sort.Ints(scores)
		return scores[len(scores)/2]
	}
}

// Given a string of parentheses, return the score of the string
// if a syntax error is found, that is, an incorrect closing parentheses
// pair. Otherwise, the score is 0 and any leftover characters are returned.
func syntaxErrorScore(input string) (int, []rune) {
	var stack []rune
	pairs := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
		'>': '<',
	}
	scoring := map[rune]int{
		')': 3,
		']': 57,
		'}': 1197,
		'>': 25137,
	}

	for _, paren := range input {
		switch paren {
		case '(', '[', '{', '<':
			stack = append(stack, paren)
		case ')', ']', '}', '>':
			if stack[len(stack)-1] != pairs[paren] {
				return scoring[paren], []rune{}
			} else {
				stack = stack[:len(stack)-1]
			}
		}
	}
	return 0, stack
}

// Given a leftover stack of parentheses, return the score of the string
func completionScore(stack []rune) int {
	var score int

	scoring := map[rune]int{
		'(': 1,
		'[': 2,
		'{': 3,
		'<': 4,
	}

	for i := len(stack) - 1; i >= 0; i-- {
		score *= 5
		switch stack[i] {
		case '(', '[', '{', '<':
			score += scoring[stack[i]]
		}
	}
	return score
}
