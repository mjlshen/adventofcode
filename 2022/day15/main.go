package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	fmt.Printf("Part 1: %d\n", beaconExclusion("test.txt", 2000000))
	//fmt.Printf("Part 2: %d\n", beaconExclusion("input.txt", true))
}

type coord struct {
	x, y int
}

func beaconExclusion(path string, row int) int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	minX, maxX := math.MaxInt, math.MinInt
	grid := make(map[int]struct{})
	sensors := map[coord]int{}
	ignore := make(map[int]struct{})
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		sx, sy, bx, by := 0, 0, 0, 0
		if _, err := fmt.Sscanf(scanner.Text(), "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by); err != nil {
			panic(err)
		}

		if min(sx, bx) < minX {
			minX = min(sx, bx)
		}
		if max(sx, bx) > maxX {
			maxX = max(sx, bx)
		}
		if by == row {
			ignore[bx] = struct{}{}
		}

		sensors[coord{sx, sy}] = abs(sx-bx) + abs(sy-by)
		//for k := range exclusion(coord{sx, sy}, coord{bx, by}, row) {
		//	grid[k] = struct{}{}
		//}
	}

	fmt.Println(len(generateOOR(sensors, 20)))

	//fmt.Println(grid)
	//fmt.Println(ignore)
	return len(grid) - len(ignore)
}

func distress(sensors map[coord]int, bound int) coord {
	
}

func generateOOR(sensors map[coord]int, bound int) map[coord]struct{} {
	ans := map[coord]struct{}{}

	for k, v := range sensors {
		for i := 0; i <= v+1; i++ {
			if k.x+i <= bound {
				if k.y+(v+1-i) <= bound {
					ans[coord{k.x + i, k.y + (v + 1 - i)}] = struct{}{}
				}
				if k.y-(v+1-i) <= bound {
					ans[coord{k.x + i, k.y - (v + 1 - i)}] = struct{}{}
				}
			}

			if k.x-i <= bound {
				if k.y+(v+1-i) <= bound {
					ans[coord{k.x - i, k.y + (v + 1 - i)}] = struct{}{}
				}
				if k.y-(v+1-i) <= bound {
					ans[coord{k.x - i, k.y - (v + 1 - i)}] = struct{}{}
				}
			}
		}
	}

	return ans
}

func exclusion(s, b coord, row int) map[int]struct{} {
	dist := abs(s.x-b.x) + abs(s.y-b.y)
	exclusions := map[int]struct{}{}

	if abs(s.y-row) >= 0 {
		for i := 0; i <= dist-abs(s.y-row); i++ {
			exclusions[s.x+i] = struct{}{}
			exclusions[s.x-i] = struct{}{}
		}
	}

	return exclusions
}

func abs(x int) int {
	if x > 0 {
		return x
	}
	return -x
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
