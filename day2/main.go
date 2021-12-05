package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Part 1: %d\n", dive("input.txt", false))
	fmt.Printf("Part 2: %d\n", dive("input.txt", true))
}

// dive returns the product of the horizontal and vertical distance traveled
func dive(path string, aim bool) int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	a, h, d := 0, 0, 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var (
			cmd string
			v   int
		)
		fmt.Sscanf(scanner.Text(), "%s %d", &cmd, &v)

		if cmd == "forward" {
			h += v
			if aim {
				d += a * v
			}
		} else if cmd == "up" {
			if aim {
				a -= v
			} else {
				d -= v
			}
		} else if cmd == "down" {
			if aim {
				a += v
			} else {
				d += v
			}
		}
	}

	return h * d
}
