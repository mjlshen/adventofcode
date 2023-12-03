package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Printf("Part 1: %d\n", numFlashes("input.txt", 100))
	fmt.Printf("Part 2: %d\n", allFlash("input.txt"))
}

type Coord struct {
	i, j int
}

type Octopi struct {
	grid    [][]int
	flashes int
}

// allFlash returns the step at which all octopi flash
func allFlash(path string) int {
	octopi := parseOctopi(path)
	prev := octopi.flashes

	for i := 0; ; i++ {
		octopi.step()
		if octopi.flashes-prev == 100 {
			return i + 1
		} else {
			prev = octopi.flashes
		}
	}
}

// numFlashes returns the number of flashes for a given octopi grid after some steps
func numFlashes(path string, steps int) int {
	octopi := parseOctopi(path)

	for i := 0; i < steps; i++ {
		octopi.step()
	}

	return octopi.flashes
}

func (o *Octopi) step() {
	// Increment each octopus in the grid by 1
	for i := range o.grid {
		for j := range o.grid[i] {
			o.grid[i][j]++
		}
	}

	// For each octopus, if it's level is > 9, increment its neighbors
	// up/down/diagonally by 1, and continue until the grid is stable
	prev := 0
	flashed := map[Coord]bool{}
	for {
		for i := range o.grid {
			for j := range o.grid[i] {
				sweep := []Coord{
					{i - 1, j - 1},
					{i - 1, j},
					{i - 1, j + 1},
					{i, j - 1},
					{i, j + 1},
					{i + 1, j - 1},
					{i + 1, j},
					{i + 1, j + 1},
				}
				if o.grid[i][j] > 9 {
					if _, ok := flashed[Coord{i, j}]; !ok {
						for _, coord := range sweep {
							if coord.i >= 0 && coord.i < len(o.grid) && coord.j >= 0 && coord.j < len(o.grid[i]) {
								o.grid[coord.i][coord.j]++
							}
						}
						flashed[Coord{i, j}] = true
					}
				}
			}
		}
		if prev == len(flashed) {
			break
		} else {
			prev = len(flashed)
		}
	}

	// For each octopus, if its level is > 9, it flashes and its level becomes 0
	for i := range o.grid {
		for j := range o.grid[i] {
			if o.grid[i][j] > 9 {
				o.grid[i][j] = 0
				o.flashes++
			}
		}
	}
}

func parseOctopi(path string) *Octopi {
	var octopi [][]int

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line []int
		for _, c := range scanner.Text() {
			level, err := strconv.Atoi(string(c))
			if err != nil {
				panic(err)
			}
			line = append(line, level)
		}
		octopi = append(octopi, line)
	}

	return &Octopi{grid: octopi, flashes: 0}
}
