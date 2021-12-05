package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Part 1: %d\n", hydroVents("input.txt", false))
	fmt.Printf("Part 2: %d\n", hydroVents("input.txt", true))
}

type Coord struct {
	x, y int
}

func hydroVents(path string, diag bool) int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	hydro_vents := map[Coord]int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		start, end := new(Coord), new(Coord)
		fmt.Sscanf(scanner.Text(), "%d,%d -> %d,%d", &start.x, &start.y, &end.x, &end.y)
		lines := generateLineCoords(start, end, diag)
		for _, coord := range lines {
			hydro_vents[*coord]++
		}
	}

	return score(hydro_vents)
}

// generateLineCoordinates returns all coordinates making up a line between two
// given coordinates
func generateLineCoords(start, end *Coord, diag bool) []*Coord {
	var coords []*Coord
	if start.x == end.x {
		// Horizontal lines
		if start.y < end.y {
			for i := start.y; i <= end.y; i++ {
				coords = append(coords, &Coord{x: start.x, y: i})
			}
		} else {
			for i := start.y; i >= end.y; i-- {
				coords = append(coords, &Coord{x: start.x, y: i})
			}
		}
	} else if start.y == end.y {
		// Vertical lines
		if start.x < end.x {
			for i := start.x; i <= end.x; i++ {
				coords = append(coords, &Coord{x: i, y: start.y})
			}
		} else {
			for i := start.x; i >= end.x; i-- {
				coords = append(coords, &Coord{x: i, y: start.y})
			}
		}
	} else if diag {
		// Diagonal lines
		if start.x < end.x && start.y < end.y {
			for i, j := start.x, start.y; i <= end.x && j <= end.y; i, j = i+1, j+1 {
				coords = append(coords, &Coord{x: i, y: j})
			}
		} else if start.x < end.x && start.y > end.y {
			for i, j := start.x, start.y; i <= end.x && j >= end.y; i, j = i+1, j-1 {
				coords = append(coords, &Coord{x: i, y: j})
			}
		} else if start.x > end.x && start.y < end.y {
			for i, j := start.x, start.y; i >= end.x && j <= end.y; i, j = i-1, j+1 {
				coords = append(coords, &Coord{x: i, y: j})
			}
		} else {
			for i, j := start.x, start.y; i >= end.x && j >= end.y; i, j = i-1, j-1 {
				coords = append(coords, &Coord{x: i, y: j})
			}
		}
	}

	return coords
}

// score returns the number of coordinates that are the
// intersection of at least two lines
func score(coords map[Coord]int) int {
	score := 0
	for _, v := range coords {
		if v > 1 {
			score++
		}
	}

	return score
}
