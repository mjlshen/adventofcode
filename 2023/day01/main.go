package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", calibrate("input.txt", false))
	fmt.Printf("Part 2: %d\n", calibrate("input.txt", true))
}

// calibrate returns the calibration value
func calibrate(path string, words bool) int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	calibrationValue := 0
	validSubstrings := map[string]int{"0": 0, "1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9}
	if words {
		validSubstrings["zero"] = 0
		validSubstrings["one"] = 1
		validSubstrings["two"] = 2
		validSubstrings["three"] = 3
		validSubstrings["four"] = 4
		validSubstrings["five"] = 5
		validSubstrings["six"] = 6
		validSubstrings["seven"] = 7
		validSubstrings["eight"] = 8
		validSubstrings["nine"] = 9
	}

	// Each elf's items are separated by an empty line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		value := 0
		value += 10 * firstOccurence(line, validSubstrings)
		value += firstOccurence(reverse(line), reverseKeys(validSubstrings))

		calibrationValue += value
	}

	return calibrationValue
}

func reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func reverseKeys(vals map[string]int) map[string]int {
	reversed := make(map[string]int, len(vals))
	for k, v := range vals {
		reversed[reverse(k)] = v
	}

	return reversed
}

func firstOccurence(s string, substrings map[string]int) int {
	min, minVal := len(s), 0

	for substring, val := range substrings {
		if i := strings.Index(s, substring); i < min && i != -1 {
			min = i
			minVal = val
		}
	}

	return minVal
}
