package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Printf("Part 1: %d\n", part1("input.txt"))
	fmt.Printf("Part 2: %d\n", part2("input.txt"))
}

type Pair struct {
	value  int
	parent *Pair
	left   *Pair
	right  *Pair
}

// Returns the magnitude of the sum of all pairs in order
func part1(path string) int {
	pairs := parsePairs(path)
	ans := pairs[0]
	for i := 0; i < len(pairs)-1; i++ {
		ans = ans.sum(pairs[i+1])
	}

	return ans.magnitude()
}

// Returns the maximum magnitude from the sum of any two pairs
func part2(path string) int {
	pairs := parsePairs(path)

	max := 0
	for i := 0; i < len(pairs); i++ {
		for j := 0; j < len(pairs); j++ {
			ans := pairs[i].copy()
			if j != i {
				ans = ans.sum(pairs[j].copy())
				if magnitude := ans.magnitude(); magnitude > max {
					max = magnitude
				}
			}
		}
	}

	return max
}

func (p *Pair) magnitude() int {
	magnitude := 0
	if p.left == nil && p.right == nil {
		return p.value
	}

	if p.left != nil {
		magnitude += 3 * p.left.magnitude()
	}

	if p.right != nil {
		magnitude += 2 * p.right.magnitude()
	}

	return magnitude
}

func (p *Pair) sum(x *Pair) *Pair {
	sum := &Pair{
		left:  p,
		right: x,
	}
	sum.left.parent = sum
	sum.right.parent = sum
	sum.reduce()

	return sum
}

func (p *Pair) reduce() *Pair {
	for {
		if exploded := p.reduceExplode(0); exploded {
			continue
		} else if split := p.reduceSplit(); split {
			continue
		} else {
			break
		}
	}

	return p
}

func (p *Pair) reduceExplode(depth int) bool {
	if depth > 4 {
		p.parent.explode()
		return true
	}

	if p.left != nil {
		if exploded := p.left.reduceExplode(depth + 1); exploded {
			return true
		}
	}

	if p.right != nil {
		return p.right.reduceExplode(depth + 1)
	}

	return false
}

func (p *Pair) explode() {
	if ln := p.leftNeighbor(); ln != nil {
		ln.value += p.left.value
	}
	if rn := p.rightNeighbor(); rn != nil {
		rn.value += p.right.value
	}
	p.left.parent = nil
	p.left = nil
	p.right.parent = nil
	p.right = nil
	p.value = 0
}

func (p *Pair) reduceSplit() bool {
	if p.left == nil && p.right == nil {
		if p.value > 9 {
			p.left = &Pair{value: p.value / 2, parent: p}
			p.right = &Pair{value: (p.value + 1) / 2, parent: p}
			p.value = 0
			return true
		}
	}

	if p.left != nil {
		if split := p.left.reduceSplit(); split {
			return true
		}
	}

	if p.right != nil {
		return p.right.reduceSplit()
	}

	return false
}

func (p *Pair) leftNeighbor() *Pair {
	prev := p
	for pair := p.parent; pair != nil; {
		if pair.left == prev {
			// The left member is exploding
			if pair.parent == nil {
				// Nothing to the left, no more outer pairs
				return nil
			}
			prev = pair
			pair = pair.parent
		} else if pair.right == prev {
			// The right member is exploding
			prev = pair
			pair = pair.left
		} else if pair.left == nil && pair.right == nil {
			// Found a value
			return pair
		} else {
			pair = pair.right
		}
	}

	return nil
}

func (p *Pair) rightNeighbor() *Pair {
	prev := p
	for pair := p.parent; pair != nil; {
		if pair.right == prev {
			// The right member is exploding
			if pair.parent == nil {
				// Nothing to the right, no more outer pairs
				return nil
			}
			prev = pair
			pair = pair.parent
		} else if pair.left == prev {
			// The left member is exploding
			prev = pair
			pair = pair.right
		} else if pair.left == nil && pair.right == nil {
			// Found a value
			return pair
		} else {
			pair = pair.left
		}
	}

	return nil
}

func (p *Pair) copy() *Pair {
	copy := &Pair{
		value: p.value,
	}
	if p.left != nil {
		copy.left = p.left.copy()
		copy.left.parent = copy
	}
	if p.right != nil {
		copy.right = p.right.copy()
		copy.right.parent = copy
	}

	return copy
}

func (p *Pair) String() string {
	if p.left != nil || p.right != nil {
		return fmt.Sprintf("[%s,%s]", p.left.String(), p.right.String())
	} else {
		return strconv.Itoa(p.value)
	}
}

func newPair(input string) *Pair {
	pair := &Pair{}
	for _, c := range input {
		if c == '[' {
			pair.left = &Pair{parent: pair}
			pair.right = &Pair{parent: pair}
			pair = pair.left
		} else if c == ']' {
			pair = pair.parent
		} else if c == ',' {
			pair = pair.parent.right
		} else {
			pair.value = int(c - '0')
		}
	}
	return pair
}

func parsePairs(path string) []*Pair {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	pairs := make([]*Pair, 0)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		pair := newPair(scanner.Text())
		pairs = append(pairs, pair)
	}

	return pairs
}
