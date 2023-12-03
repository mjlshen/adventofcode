package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", possibleGames("input.txt"))
	fmt.Printf("Part 2: %d\n", minimumBag("input.txt"))
}

type bag struct {
	blue  int
	red   int
	green int
}

func possibleGames(path string) int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	ans := 0
	const (
		maxRed   = 12
		maxGreen = 13
		maxBlue  = 14
	)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Sample line
		// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
		s := strings.Split(scanner.Text(), ":")
		var round int
		if _, err := fmt.Sscanf(s[0], "Game %d", &round); err != nil {
			panic(err)
		}

		// sets = ["3 blue, 4 red", "1 red, 2 green" ...]
		sets := strings.Split(s[1], ";")
		b := newBag(sets)
		if b.red <= maxRed && b.green <= maxGreen && b.blue <= maxBlue {
			ans += round
		}
	}

	return ans
}

func minimumBag(path string) int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	ans := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ":")
		sets := strings.Split(s[1], ";")
		b := newBag(sets)

		ans += b.score()
	}

	return ans
}

func newBag(sets []string) bag {
	b := &bag{}
	for _, set := range sets {
		cubes := strings.Split(set, ",")
		for _, cube := range cubes {
			var (
				number int
				color  string
			)
			if _, err := fmt.Sscanf(cube, "%d %s", &number, &color); err != nil {
				panic(err)
			}

			switch {
			case color == "red":
				b.red = max(b.red, number)
			case color == "green":
				b.green = max(b.green, number)
			case color == "blue":
				b.blue = max(b.blue, number)
			default:
				panic(fmt.Errorf("unexpected color: %s", color))
			}
		}
	}

	return *b
}

func (b bag) score() int {
	return b.red * b.green * b.blue
}
