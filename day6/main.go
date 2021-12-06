package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", modelLanternfish("input.txt", 80))
	fmt.Printf("Part 2: %d\n", modelLanternfish("input.txt", 256))
}

func modelLanternfish(path string, days int) int {
	line, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	// Parse input and store in a map with key of lifecycle
	// and value of count of lanternfish at that lifecycle
	lanternfish := make(map[int]int)
	for _, l := range strings.Split(string(line), ",") {
		count, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		lanternfish[count]++
	}

	// At each day, the lifecycle decrements by 1
	// If a lifecycle was at 0, it would decrement to -1 temporarily
	// At this point a new lanternfish is spawned at lifecycle 8.
	// Then the original lanternfish is reset to lifecycle 6.
	for i := 0; i < days; i++ {
		nextgen := make(map[int]int)
		for k, v := range lanternfish {
			nextgen[k-1] = v
		}
		nextgen[8] = nextgen[-1]
		nextgen[6] += nextgen[-1]
		nextgen[-1] = 0
		lanternfish = nextgen
	}

	count := 0
	for _, v := range lanternfish {
		count += v
	}

	return count
}
