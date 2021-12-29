package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	fmt.Printf("Part 1: %d\n", numIncreases("input.txt"))
	fmt.Printf("Part 2: %d\n", numIncreasesWindow("input.txt", 3))
}

// numIncreases returns the number of times a measurement increases
func numIncreases(path string) int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	prev, ans := math.MinInt, -1

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var m int
		if _, err := fmt.Sscanf(scanner.Text(), "%d", &m); err != nil {
			panic(err)
		}
		if m > prev {
			ans++
		}
		prev = m
	}
	return ans
}

// numIncreasesWindow returns the number of times a sliding-window of size w
// increases, where a sliding window of three are three consecutive measurements
func numIncreasesWindow(path string, w int) int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	prev, ans := math.MinInt, -1
	window := make([]int, w)

	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		if i < w-1 {
			if _, err := fmt.Sscanf(scanner.Text(), "%d", &window[i]); err != nil {
				panic(err)
			}
			continue
		}

		if _, err := fmt.Sscanf(scanner.Text(), "%d", &window[len(window)-1]); err != nil {
			panic(err)
		}

		m := sum(window)
		if m > prev {
			ans++
		}

		prev = m
		window = shift(window)
	}

	return ans
}

func sum(a []int) int {
	s := 0
	for _, v := range a {
		s += v
	}
	return s
}

// Shift the values to the left, don't care about the last value
func shift(a []int) []int {
	for i := 1; i < len(a); i++ {
		a[i-1] = a[i]
	}
	return a
}
