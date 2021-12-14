package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", count1478("input.txt"))
	fmt.Printf("Part 2: %d\n", sumOutput("input.txt"))
}

var (
	// SEGMENTS maps a digit to the segments it contains
	SEGMENTS map[int][]rune = map[int][]rune{
		0: {'a', 'b', 'c', 'e', 'f', 'g'},
		1: {'c', 'f'},
		2: {'a', 'c', 'd', 'e', 'g'},
		3: {'a', 'c', 'd', 'f', 'g'},
		4: {'b', 'c', 'd', 'f'},
		5: {'a', 'b', 'd', 'f', 'g'},
		6: {'a', 'b', 'd', 'e', 'f', 'g'},
		7: {'a', 'c', 'f'},
		8: {'a', 'b', 'c', 'd', 'e', 'f', 'g'},
		9: {'a', 'b', 'c', 'd', 'f', 'g'},
	}
)

// count1478 returns the number of 1, 4, 7, and 8s in the output
func count1478(path string) int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	count := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " | ")
		output := strings.Split(line[1], " ")
		for _, o := range output {
			if len(o) == len(SEGMENTS[1]) ||
				len(o) == len(SEGMENTS[4]) ||
				len(o) == len(SEGMENTS[7]) ||
				len(o) == len(SEGMENTS[8]) {
				count++
			}
		}
	}

	return count
}

// sumOutput translates the output into numbers
// then returns the sum of all of the output numbers
func sumOutput(path string) int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), " | ")
		input := strings.Split(line[0], " ")
		output := strings.Split(line[1], " ")
		sum += translateOutput(deduceSegment(input), output)
	}

	return sum
}

// translateOutput translates an array of strings of lit segments
// into a number by concatenating the digits.
// segments is a translation dictionary of the actual --> faulty segments
func translateOutput(segments map[rune]rune, output []string) int {
	answer := 0
	for _, digit := range output {
		answer *= 10
		answer += translateDigit(segments, digit)
	}
	return answer
}

// translateDigit translates a string of lit segments into a digit
func translateDigit(segments map[rune]rune, digit string) int {
	if len(digit) == len(SEGMENTS[1]) {
		return 1
	} else if len(digit) == len(SEGMENTS[4]) {
		return 4
	} else if len(digit) == len(SEGMENTS[7]) {
		return 7
	} else if len(digit) == len(SEGMENTS[8]) {
		return 8
	}

	// 0, 6, 9
	if len(digit) == len(SEGMENTS[0]) {
		if !arrayContains([]rune(digit), segments['d']) {
			return 0
		} else if !arrayContains([]rune(digit), segments['c']) {
			return 6
		} else {
			return 9
		}
	}

	// 2, 3, 5
	if len(digit) == len(SEGMENTS[2]) {
		if arrayContains([]rune(digit), segments['b']) {
			return 5
		} else if arrayContains([]rune(digit), segments['e']) {
			return 2
		} else {
			return 3
		}
	}

	return 0
}

// deduceSegment returns a map of the actual --> faulty segments
// given an array of input segments that have scrambled mappings
func deduceSegment(inputSegments []string) map[rune]rune {
	segments := map[rune]rune{
		'a': 0,
		'b': 0,
		'c': 0,
		'd': 0,
		'e': 0,
		'f': 0,
		'g': 0,
	}
	// Stores the runes contained after deducing which segment they belong to
	possibleSegment := map[int][]rune{
		0: {},
		1: {},
		2: {},
		3: {},
		4: {},
		5: {},
		6: {},
		7: {},
		8: {},
		9: {},
	}

	// Segments 1, 4, 7, and 8 can be deduced solely by their length
	for _, inputSegment := range inputSegments {
		if len(inputSegment) == len(SEGMENTS[1]) {
			possibleSegment[1] = append(possibleSegment[1], []rune(inputSegment)...)
		} else if len(inputSegment) == len(SEGMENTS[4]) {
			possibleSegment[4] = append(possibleSegment[4], []rune(inputSegment)...)
		} else if len(inputSegment) == len(SEGMENTS[7]) {
			possibleSegment[7] = append(possibleSegment[7], []rune(inputSegment)...)
		} else if len(inputSegment) == len(SEGMENTS[8]) {
			possibleSegment[8] = append(possibleSegment[8], []rune(inputSegment)...)
		}
	}

	// Solve 'a'
	// 7 contains all segments that 1 does, except for 'a'
	// and we have already identified 1 and 7.
	for _, char := range possibleSegment[7] {
		if !arrayContains(possibleSegment[1], char) {
			segments['a'] = char
		}
	}

	// Solve 'c' and 'f'
	// 0, 6, and 9 all use 6 segments, but we can uniquely identify 6
	// 0 contains both segments of 1, but is missing 'd'
	// 9 contains both segments of 1, but is missing 'e'
	// 6 only contains 'f' from 1 and is missing 'c', so we can find 6
	for _, inputSegment := range inputSegments {
		if len(inputSegment) == len(SEGMENTS[6]) {
			// We found 'c' and therefore identify 'f' as well
			if !arrayContains([]rune(inputSegment), possibleSegment[1][0]) {
				possibleSegment[6] = append(possibleSegment[6], []rune(inputSegment)...)
				segments['c'] = possibleSegment[1][0]
				segments['f'] = possibleSegment[1][1]
			} else if !arrayContains([]rune(inputSegment), possibleSegment[1][1]) {
				possibleSegment[6] = append(possibleSegment[6], []rune(inputSegment)...)
				segments['c'] = possibleSegment[1][1]
				segments['f'] = possibleSegment[1][0]
			}
		}
	}

	// 2, 3, and 5 both use 5 segments
	// 2 has 'c', but is missing 'f'
	// 3 contains both 'c' and 'f'
	// 5 has 'f', but is missing 'c'
	for _, inputSegment := range inputSegments {
		if len(inputSegment) == len(SEGMENTS[2]) {
			if arrayContains([]rune(inputSegment), segments['c']) && arrayContains([]rune(inputSegment), segments['f']) {
				// We found 3
				possibleSegment[3] = append(possibleSegment[3], []rune(inputSegment)...)
			} else if !arrayContains([]rune(inputSegment), segments['f']) {
				// We found 2
				possibleSegment[2] = append(possibleSegment[2], []rune(inputSegment)...)
			} else if !arrayContains([]rune(inputSegment), segments['c']) {
				// We found 5
				possibleSegment[5] = append(possibleSegment[5], []rune(inputSegment)...)
			}
		}
	}

	// Solve 'b' and 'e'
	// 2 has 'e', but is missing 'b' and 'f'
	// 5 has 'b', but is missing 'c' and 'e'
	// Since 'c' and 'f' are known already, we can solve for 'b' and 'e'
	for _, r := range possibleSegment[2] {
		if r != segments['c'] {
			if !arrayContains(possibleSegment[5], r) {
				segments['e'] = r
			}
		}
	}
	for _, r := range possibleSegment[5] {
		if r != segments['f'] {
			if !arrayContains(possibleSegment[2], r) {
				segments['b'] = r
			}
		}
	}

	// Now that we know 'c' and 'e', we can differentiate between 0, 6, and 9
	for _, inputSegment := range inputSegments {
		if len(inputSegment) == len(SEGMENTS[0]) {
			if !arrayContains([]rune(inputSegment), segments['e']) {
				possibleSegment[9] = append(possibleSegment[9], []rune(inputSegment)...)
			} else if arrayContains([]rune(inputSegment), segments['c']) {
				possibleSegment[0] = append(possibleSegment[0], []rune(inputSegment)...)
			}
		}
	}

	// Solve 'd'
	// 0 has all segments except for 'd'
	for r := 'a'; r <= 'g'; r++ {
		if !arrayContains(possibleSegment[0], r) {
			segments['d'] = r
			break
		}
	}

	// Solve 'g'
	// We know all of the other segments, so this is just the remaining one
	for r := 'a'; r <= 'g'; r++ {
		if !mapContains(segments, r) {
			segments['g'] = r
		}
	}

	return segments
}

func mapContains(a map[rune]rune, x rune) bool {
	for _, v := range a {
		if v == x {
			return true
		}
	}
	return false
}

func arrayContains(a []rune, x rune) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}
