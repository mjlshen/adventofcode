package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

func main() {
	fmt.Printf("Part 1: %d\n", sumMinPath("input.txt", false))
	fmt.Printf("Part 2: %d\n", sumMinPath("input.txt", true))
}

type Coord struct {
	x, y int
}

type Node struct {
	c              Coord
	prev           *Node
	score          int
	heuristicScore int
}

type PriorityQueue struct {
	chitons  map[Coord]int
	queue    []*Node
	expanded map[Node]bool
	end      Coord
}

func sumMinPath(path string, fullCave bool) int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	chitons := map[Coord]int{}
	x, y := 0, 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		y = 0
		for _, i := range scanner.Text() {
			chitons[Coord{x, y}] = int(i - '0')
			if fullCave {
				for j := 1; j < 5; j++ {
					val := int(i-'0') + j
					if val > 9 {
						val -= 9
					}
					chitons[Coord{x, y + len(scanner.Text())*j}] += val
				}
			}
			y++
		}
		x++
	}

	if fullCave {
		for transpositions := 1; transpositions < 5; transpositions++ {
			for i := 0; i < x; i++ {
				for j := 0; j < 5*y; j++ {
					val := chitons[Coord{i, j}] + transpositions
					if val > 9 {
						val -= 9
					}
					chitons[Coord{i + x*transpositions, j}] = val
				}
			}
		}
	}

	maxX, maxY := x, y
	if fullCave {
		maxX = x * 5
		maxY = y * 5
	}

	priorityQueue := &PriorityQueue{
		chitons:  chitons,
		queue:    []*Node{},
		expanded: map[Node]bool{},
		end:      Coord{maxX - 1, maxY - 1},
	}

	priorityQueue.insert(
		Node{
			c:     Coord{0, 0},
			prev:  nil,
			score: 0,
		},
	)

	ans := priorityQueue.aStar()

	return ans
}

func (p *PriorityQueue) String() string {
	var output string
	for _, node := range p.queue {
		output += fmt.Sprintf("Coord: {%d,%d}, Prev: %v, Score: %d\n", node.c.x, node.c.y, *node.prev, node.score)
	}

	return output
}

func (p *PriorityQueue) insert(n Node) {
	for _, node := range p.queue {
		if node.c == n.c {
			if n.score < node.score {
				node.score = n.score
				node.prev = n.prev
			}
			return
		}
	}

	p.queue = append(p.queue, &n)
	p.sort()
}

func (p *PriorityQueue) pop() *Node {
	n := p.queue[0]
	if n.prev != nil {
		n.score = n.prev.score + p.chitons[n.c]
	} else {
		n.score = 0
	}

	p.queue = p.queue[1:]
	return n
}

func (p *PriorityQueue) sort() {
	sort.SliceStable(p.queue, func(i, j int) bool {
		return p.queue[i].heuristicScore < p.queue[j].heuristicScore
	})
}

func manhattanDistance(start Coord, end Coord) int {
	return abs(start.x-end.x) + abs(start.y-end.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (p *PriorityQueue) expand(n *Node) {
	sweep := []Coord{
		{n.c.x - 1, n.c.y},
		{n.c.x + 1, n.c.y},
		{n.c.x, n.c.y - 1},
		{n.c.x, n.c.y + 1},
	}

	for _, neighbor := range sweep {
		if _, ok := p.chitons[neighbor]; ok {
			found := false
			for k := range p.expanded {
				if k.c == neighbor {
					found = true
					break
				}
			}
			if !found {
				p.insert(Node{
					c:              neighbor,
					prev:           n,
					score:          n.score + p.chitons[neighbor],
					heuristicScore: n.score + manhattanDistance(p.end, neighbor) + p.chitons[neighbor],
				})
			}
		}
	}
}

func (p *PriorityQueue) score(end Coord) int {
	score := 0
	current := end
	for {
		if current.x == 0 && current.y == 0 {
			return score
		}

		found := false
		for k := range p.expanded {
			if k.c == current {
				score += p.chitons[k.c]
				current = k.prev.c
				found = true
				break
			}
		}

		if !found {
			panic("Not found")
		}
	}
}

func (p *PriorityQueue) aStar() int {
	numpop := 0
	for {
		n := p.pop()
		numpop++
		fmt.Println(numpop)
		p.expanded[*n] = true
		if n.c == p.end {
			return p.score(p.end)
		}
		p.expand(n)
	}
}
