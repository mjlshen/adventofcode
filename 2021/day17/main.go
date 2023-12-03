package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Part 1: %d\n", maxYHeight("input.txt"))
	fmt.Printf("Part 2: %d\n", countVelocities("input.txt"))
}

type Coord struct {
	x, y int
}

func parseBounds(path string) (Coord, Coord) {
	line, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	var minBound, maxBound Coord
	fmt.Sscanf(
		string(line), "target area: x=%d..%d, y=%d..%d",
		&minBound.x, &maxBound.x,
		&minBound.y, &maxBound.y,
	)

	return minBound, maxBound
}

// maxYHeight returns the max Y coordinate reached by a certain shot
func maxYHeight(path string) int {
	minBound, _ := parseBounds(path)

	// Velocity in the y direction after n steps
	// vy = vy0 - n
	// Position in the y direction after n steps
	// y = vy0 * n - n(n-1)/2
	// When n = 2vy0 + 1, y == 0, with vy = -vy0 - 1
	// So we just need to check -vy0 - 1 >= yMin
	vy0 := -minBound.y - 1

	// Then, given vy0, the max height is vy0 * (vy0 + 1) / 2
	// Since it's the sum of n from 0 ... vy0
	return vy0 * (vy0 + 1) / 2
}

// countVelocities returns the number of possible shots that will hit the target
func countVelocities(path string) int {
	minBound, maxBound := parseBounds(path)

	count := 0
	vxMax := maxBound.x + 1
	vyMax := -minBound.y

	for vx := 1; vx < vxMax; vx++ {
		for vy := -vyMax; vy < vyMax; vy++ {
			if hit(0, 0, vx, vy, minBound, maxBound) {
				count++
			}
		}
	}
	return count
}

// hit returns true if a shot with velocity (vx, vy) will hit the target area
// defined by minBound and maxBound, starting from (x, y)
func hit(x, y, vx, vy int, minBound, maxBound Coord) bool {
	if x > maxBound.x || y < minBound.y {
		// If past the maximum bounds
		return false
	} else if x >= minBound.x && y <= maxBound.y {
		// Otherwise, if we're inside the minimum bounds
		return true
	} else {
		newVx := vx
		newVy := vy - 1
		if vx > 0 {
			newVx--
		}
		return hit(x+vx, y+vy, newVx, newVy, minBound, maxBound)
	}
}
