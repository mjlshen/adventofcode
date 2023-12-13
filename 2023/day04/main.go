package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", winningNumbers("input.txt"))
	fmt.Printf("Part 2: %d\n", totalScratchcards("input.txt"))
}

func winningNumbers(path string) int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	ans := 0

	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		ans += int(math.Pow(2, float64(matches(strings.Split(scanner.Text(), ":")[1])-1)))
	}

	return ans
}

func matches(line string) int {
	winningString := strings.Split(line, "|")[0]
	numbersString := strings.Split(line, "|")[1]
	winningNumbers := map[int]struct{}{}
	numbers := map[int]struct{}{}

	for _, s := range strings.Fields(winningString) {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}

		winningNumbers[num] = struct{}{}
	}

	for _, s := range strings.Fields(numbersString) {
		num, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}

		numbers[num] = struct{}{}
	}

	matches := 0
	for num := range winningNumbers {
		if _, ok := numbers[num]; ok {
			matches++
		}
	}

	return matches
}

func totalScratchcards(path string) int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scratchcards := map[int]int{}

	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		var cardNum int

		if _, err := fmt.Sscanf(strings.Split(line, ":")[0], "Card %d", &cardNum); err != nil {
			panic(err)
		}
		scratchcards[cardNum]++

		for i := 0; i < matches(strings.Split(line, ":")[1]); i++ {
			scratchcards[cardNum+i+1] += scratchcards[cardNum]
		}
	}

	ans := 0
	for _, v := range scratchcards {
		ans += v
	}

	return ans
}
