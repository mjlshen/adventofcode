package main

import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", alignCrabs("input.txt", false))
	fmt.Printf("Part 2: %d\n", alignCrabs("input.txt", true))
}

// alignCrabs computes the amount of fuel needed for all crabs to
// converge to a location requiring the minimunm amount of fuel.
// With exp: false, crabs consume 1 fuel per step.
// With exp: true, crabs consume 1 fuel for the first step, 2 for the second,
// 3 for the third, and so on, following n(n+1)/2 for n steps.
func alignCrabs(path string, exp bool) int {
	var (
		crabs []int
		min   int
	)

	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	cstrings := strings.Split(string(file), ",")
	for _, crab := range cstrings {
		c, err := strconv.Atoi(crab)
		if err != nil {
			panic(err)
		}
		crabs = append(crabs, c)
	}
	sort.Ints(crabs)

	for i := crabs[0]; i < len(crabs); i++ {
		fuel := 0
		if !exp {
			for _, crab := range crabs {
				fuel += Abs(crab - i)
			}
		} else {
			for _, crab := range crabs {
				fuel += (Abs(crab-i)*Abs(crab-i) + Abs(crab-i)) / 2
			}
		}
		if fuel < min || min == 0 {
			min = fuel
		}
	}

	return min
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
