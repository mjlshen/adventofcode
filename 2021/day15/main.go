package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
	"os"
)

func main() {
	fmt.Printf("Part 1: %d\n", parseCave("input.txt", false).djikstra())
	fmt.Printf("Part 2: %d\n", parseCave("input.txt", true).djikstra())
}

type Cave struct {
	chitons    [][]*Node
	start, end Coord
	expanded   map[Coord]bool
}

type Coord struct {
	x, y int
}

type Node struct {
	c              Coord
	prev           *Node
	score          int
	heuristicScore int
	index          int
}

type PriorityQueue []*Node

func (c *Cave) djikstra() int {
	pq := PriorityQueue{}
	heap.Init(&pq)
	heap.Push(&pq, &Node{c: c.start, heuristicScore: 0})

	for len(pq) > 0 {
		n := heap.Pop(&pq).(*Node)
		if n.c == c.end {
			return n.heuristicScore
		}
		c.expanded[n.c] = true

		sweep := []Coord{
			{n.c.x - 1, n.c.y},
			{n.c.x + 1, n.c.y},
			{n.c.x, n.c.y - 1},
			{n.c.x, n.c.y + 1},
		}
		for _, neighbor := range sweep {
			if neighbor.x >= 0 && neighbor.x < len(c.chitons[0]) && neighbor.y >= 0 && neighbor.y < len(c.chitons) {
				if _, ok := c.expanded[neighbor]; !ok {
					newScore := n.heuristicScore + c.chitons[neighbor.x][neighbor.y].score
					if newScore < c.chitons[neighbor.x][neighbor.y].heuristicScore {
						c.chitons[neighbor.x][neighbor.y].prev = n
						c.chitons[neighbor.x][neighbor.y].heuristicScore = newScore
						heap.Push(&pq, c.chitons[neighbor.x][neighbor.y])
					}
				}
			}
		}
	}

	return 0
}

func parseCave(path string, fullCave bool) *Cave {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	chitons := [][]*Node{}

	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		chitonsRow := []*Node{}
		for j, num := range scanner.Text() {
			chitonsRow = append(chitonsRow, &Node{
				c:              Coord{x: i, y: j},
				score:          int(num - '0'),
				heuristicScore: math.MaxInt,
			})
		}

		if fullCave {
			// Repeat 5 times in the positive x direction
			originalLen := len(chitonsRow)
			for i := 1; i < 5; i++ {
				for j := 0; j < originalLen; j++ {
					val := chitonsRow[j].score + i
					if val > 9 {
						val -= 9
					}
					chitonsRow = append(chitonsRow, &Node{
						c:              Coord{x: chitonsRow[j].c.x, y: originalLen*i + j},
						score:          val,
						heuristicScore: math.MaxInt,
					})
				}
			}
		}
		chitons = append(chitons, chitonsRow)
	}

	if fullCave {
		// Repeat 5 times in the positive y direction
		originalLen := len(chitons[0])
		originalHeight := len(chitons)
		for i := 1; i < 5; i++ {
			for j := 0; j < originalHeight; j++ {
				chitonsRow := []*Node{}
				for k := 0; k < originalLen; k++ {
					val := chitons[j][k].score + i
					if val > 9 {
						val -= 9
					}
					chitonsRow = append(chitonsRow, &Node{
						c:              Coord{x: j + originalHeight*i, y: k},
						score:          val,
						heuristicScore: math.MaxInt,
					})
				}
				chitons = append(chitons, chitonsRow)
			}
		}
	}

	return &Cave{
		chitons:  chitons,
		start:    Coord{0, 0},
		end:      Coord{x: len(chitons) - 1, y: len(chitons[0]) - 1},
		expanded: map[Coord]bool{},
	}
}

func (c Cave) String() string {
	var s string
	for _, row := range c.chitons {
		for _, chiton := range row {
			s += fmt.Sprintf("%d", chiton.score)
		}
		s += "\n"
	}
	return s[:len(s)-1]
}

// https://pkg.go.dev/container/heap#example-package-PriorityQueue
func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].heuristicScore < pq[j].heuristicScore
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	node := old[len(old)-1]
	old[len(old)-1] = nil // avoid memory leak
	node.index = -1       // for safety
	*pq = old[0 : len(old)-1]
	return node
}

func (pq *PriorityQueue) Push(x interface{}) {
	node := x.(*Node)
	node.index = len(*pq)
	*pq = append(*pq, node)
}
