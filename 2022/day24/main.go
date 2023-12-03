//package main
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//)
//
//func main() {
//	fmt.Printf("Part 1: %d\n", blizzardBasin("input.txt"))
//	//fmt.Printf("Part 2: %d\n", blizzardBasin("input.txt", decryptionKey, 10))
//}
//
//type coord struct {
//	x, y int
//}
//
//type expedition struct {
//	position coord
//	time     int
//}
//
//func blizzardBasin(path string) int {
//	file, err := os.Open(path)
//	if err != nil {
//		panic(err)
//	}
//
//	maxX, maxY := 0, 0
//	valley := map[coord]rune{}
//	scanner := bufio.NewScanner(file)
//	for y := 0; scanner.Scan(); y++ {
//		for x, r := range scanner.Text() {
//			valley[coord{x, y}] = r
//			if x > maxX {
//				maxX = x
//			}
//		}
//
//		if y > maxY {
//			maxY = y
//		}
//	}
//
//	start := coord{1, 0}
//	end := coord{maxX - 1, maxY}
//
//	fmt.Println(valley)
//	fmt.Println(start, end)
//	return 0
//}
//
//func bfs(start, end coord) int {
//	movements := map[rune]coord{
//		'#': {0, 0},
//		'^': {0, 1},
//		'>': {1, 0},
//		'v': {0, 1},
//		'<': {-1, 0},
//	}
//
//	queue := []expedition{{position: start, time: 0}}
//	seen := map[expedition]struct{}{
//		expedition{position: start, time: 0}: {},
//	}
//
//	for len(queue) > 0 {
//		curr := queue[0]
//		queue = queue[1:]
//
//		for _, movement := range movements {
//			next := expedition{position: curr.position}
//		}
//	}
//
//	return 0
//}
//
//func (c coord) add(a coord) coord {
//	return coord{c.x + a.x, c.y + a.y}
//}

package main

import (
	"fmt"
	"image"
	"os"
	"strings"
)

type State struct {
	P image.Point
	T int
}

func main() {
	input, _ := os.ReadFile("input.txt")

	vall := map[image.Point]rune{}
	for y, s := range strings.Fields(string(input)) {
		for x, r := range s {
			vall[image.Point{x, y}] = r
		}
	}

	var bliz image.Rectangle
	for p := range vall {
		bliz = bliz.Union(image.Rectangle{p, p.Add(image.Point{1, 1})})
	}
	bliz.Min, bliz.Max = bliz.Min.Add(image.Point{1, 1}), bliz.Max.Sub(image.Point{1, 1})

	bfs := func(start image.Point, end image.Point, time int) int {
		delta := map[rune]image.Point{
			'#': {0, 0}, '^': {0, -1}, '>': {1, 0}, 'v': {0, 1}, '<': {-1, 0},
		}

		queue := []State{{start, time}}
		seen := map[State]struct{}{queue[0]: {}}

		for len(queue) > 0 {
			cur := queue[0]
			queue = queue[1:]

		loop:
			for _, d := range delta {
				next := State{cur.P.Add(d), cur.T + 1}
				if next.P == end {
					return next.T
				}

				if _, ok := seen[next]; ok {
					continue
				}
				if r, ok := vall[next.P]; !ok || r == '#' {
					continue
				}

				if next.P.In(bliz) {
					for r, d := range delta {
						if vall[next.P.Sub(d.Mul(next.T)).Mod(bliz)] == r {
							continue loop
						}
					}
				}

				seen[next] = struct{}{}
				queue = append(queue, next)
			}
		}
		return -1
	}

	start, end := bliz.Min.Sub(image.Point{0, 1}), bliz.Max.Sub(image.Point{1, 0})
	fmt.Println(bfs(start, end, 0))
	fmt.Println(bfs(start, end, bfs(end, start, bfs(start, end, 0))))
}
