package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", numDots("input.txt", true))
	fmt.Printf("Part 2: %d\n", numDots("input.txt", false))
}

type Coord struct {
	x, y int
}

type Dots map[Coord]bool

type Fold struct {
	axis  string
	value int
}

func numDots(path string, oneFold bool) int {
	dots, folds := parseDotsAndFolds(path)
	if oneFold {
		dots = fold(dots, folds[0])
	} else {
		for _, f := range folds {
			dots = fold(dots, f)
		}
	}

	if !oneFold {
		fmt.Printf("%s\n", dots)
	}
	return len(dots)
}

func fold(d Dots, f *Fold) Dots {
	newDots := map[Coord]bool{}

	for d := range d {
		newX, newY := d.x, d.y
		if f.axis == "x" {
			if newX >= f.value {
				newX -= 2 * (newX - f.value)
			}
		} else {
			if newY >= f.value {
				newY -= 2 * (newY - f.value)
			}
		}

		newDots[Coord{x: newX, y: newY}] = true
	}

	return newDots
}

// Prints Dots in a grid with █ where there are dots and an empty space otherwise
func (d Dots) String() string {
	maxX, maxY := 0, 0
	for k := range d {
		if k.x > maxX {
			maxX = k.x
		}
		if k.y > maxY {
			maxY = k.y
		}
	}
	output := make([][]string, maxY+1)
	for i := range output {
		output[i] = make([]string, maxX+1)
		for j := range output[i] {
			output[i][j] = " "
		}
	}

	for k := range d {
		output[k.y][k.x] = "█"
	}

	os := ""
	for i := range output {
		os += fmt.Sprintf("%s\n", strings.Join(output[i], ""))
	}
	return os
}

func parseDotsAndFolds(path string) (Dots, []*Fold) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	dots := map[Coord]bool{}
	folds := []*Fold{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}

		var x, y int
		fmt.Sscanf(scanner.Text(), "%d,%d", &x, &y)
		dots[Coord{x: x, y: y}] = true
	}

	for scanner.Scan() {
		var (
			axis  string
			value int
		)

		fmt.Sscanf(scanner.Text(), "fold along %1s=%d:", &axis, &value)
		folds = append(folds, &Fold{axis: axis, value: value})
	}

	return dots, folds
}
