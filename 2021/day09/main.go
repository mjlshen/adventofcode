package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	fmt.Printf("Part 1: %d\n", scoreHeightmap("input.txt", false))
	fmt.Printf("Part 2: %d\n", scoreHeightmap("input.txt", true))
}

type Coord struct {
	i, j int
}

// If basin is false, find all low points and return their sum
// If basin is true, map out all the basins and return the product of the three largest
func scoreHeightmap(path string, basin bool) int {
	var (
		risklevel int
		basins    []Coord
	)

	heightmap := parseHeightmap(path)

	for i := 0; i < len(heightmap); i++ {
		for j := 0; j < len(heightmap[i]); j++ {
			if isBasin(heightmap, Coord{i, j}) {
				risklevel += heightmap[i][j] + 1
				basins = append(basins, Coord{i, j})
			}
		}
	}

	if basin {
		var sizes []int

		for _, basin := range basins {
			basinMap := searchBasin(heightmap, map[Coord]bool{basin: true})
			sizes = append(sizes, len(basinMap))
		}

		sort.Ints(sizes)
		return sizes[len(sizes)-1] * sizes[len(sizes)-2] * sizes[len(sizes)-3]
	} else {
		return risklevel
	}
}

// For each Coord in basin, search adjacent horizontal/vertical neighbors
// Any of them that are not 9 are added to the basin
func searchBasin(heightmap [][]int, basin map[Coord]bool) map[Coord]bool {
	addedPoints := false

	for b := range basin {
		sweep := []Coord{
			{b.i - 1, b.j},
			{b.i + 1, b.j},
			{b.i, b.j - 1},
			{b.i, b.j + 1},
		}

		// Ensure c is lower than horizontal/vertical neighbors
		for _, adj := range sweep {
			if adj.i >= 0 && adj.i < len(heightmap) && adj.j >= 0 && adj.j < len(heightmap[adj.i]) {
				if _, ok := basin[Coord{adj.i, adj.j}]; !ok && heightmap[adj.i][adj.j] != 9 {
					addedPoints = true
					basin[Coord{adj.i, adj.j}] = true
				}
			}
		}
	}

	if addedPoints {
		return searchBasin(heightmap, basin)
	} else {
		// The entire basin is mapped out
		return basin
	}
}

func isBasin(heightmap [][]int, c Coord) bool {
	sweep := []Coord{
		{c.i - 1, c.j},
		{c.i + 1, c.j},
		{c.i, c.j - 1},
		{c.i, c.j + 1},
	}

	// Ensure c is lower than horizontal/vertical neighbors
	for _, adj := range sweep {
		if adj.i >= 0 && adj.i < len(heightmap) && adj.j >= 0 && adj.j < len(heightmap[adj.i]) {
			if heightmap[c.i][c.j] > heightmap[adj.i][adj.j] {
				return false
			}
		}
	}

	return true
}

func parseHeightmap(path string) [][]int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var heightmap [][]int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var line []int
		for _, r := range scanner.Text() {
			val, err := strconv.Atoi(string(r))
			if err != nil {
				panic(err)
			}
			line = append(line, val)
		}
		heightmap = append(heightmap, line)
	}

	return heightmap
}
