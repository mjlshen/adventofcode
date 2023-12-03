package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Part 1: %d\n", parseSeafloor("input.txt").stable())
}

type Coord struct {
	x, y int
}

type Seafloor struct {
	east  map[Coord]bool
	south map[Coord]bool
	xMax  int
	yMax  int
}

func (sf Seafloor) stable() int {
	moved := true
	moves := 0
	for ; moved; moves++ {
		sf, moved = sf.move()
	}

	return moves
}

func (sf Seafloor) move() (Seafloor, bool) {
	moved := false
	nsf := Seafloor{
		east:  map[Coord]bool{},
		south: map[Coord]bool{},
		xMax:  sf.xMax,
		yMax:  sf.yMax,
	}

	for k := range sf.east {
		newCoord := Coord{k.x, k.y + 1}
		if k.y+1 >= sf.yMax {
			newCoord = Coord{k.x, 0}
		}
		_, ok := sf.south[newCoord]
		_, eok := sf.east[newCoord]
		if !ok && !eok {
			nsf.east[newCoord] = true
			moved = true
		} else {
			nsf.east[k] = true
		}
	}

	for k := range sf.south {
		newCoord := Coord{k.x + 1, k.y}
		if k.x+1 >= sf.xMax {
			newCoord = Coord{0, k.y}
		}

		_, ok := sf.south[newCoord]
		_, eok := nsf.east[newCoord]
		if !ok && !eok {
			nsf.south[newCoord] = true
			moved = true
		} else {
			nsf.south[k] = true
		}
	}

	return nsf, moved
}

func parseSeafloor(path string) Seafloor {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sf := &Seafloor{
		east:  make(map[Coord]bool),
		south: make(map[Coord]bool),
	}

	scanner := bufio.NewScanner(file)
	for i := 0; scanner.Scan(); i++ {
		for j, cucumber := range scanner.Text() {
			if cucumber == 'v' {
				sf.south[Coord{i, j}] = true
			} else if cucumber == '>' {
				sf.east[Coord{i, j}] = true
			}
			if j > sf.yMax {
				sf.yMax = j
			}
		}
		if i > sf.xMax {
			sf.xMax = i
		}
	}

	sf.xMax++
	sf.yMax++

	return *sf
}

func (sf Seafloor) String() string {
	var output string
	for i := 0; i < sf.xMax; i++ {
		for j := 0; j < sf.yMax; j++ {
			if _, ok := sf.east[Coord{i, j}]; ok {
				output += ">"
			} else if _, ok := sf.south[Coord{i, j}]; ok {
				output += "v"
			} else {
				output += "."
			}
		}
		output += "\n"
	}

	return output[:len(output)-1]
}
