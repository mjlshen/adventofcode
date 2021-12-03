package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", dive("input.txt", false))
	fmt.Printf("Part 2: %d\n", dive("input.txt", true))
}

// dive returns the product of the horizontal and vertical distance traveled
func dive(path string, aim bool) int {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	a, h, d := 0, 0, 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		cmd := strings.Fields(line)
		v, err := strconv.Atoi(cmd[1])
		if err != nil {
			log.Fatal(err)
		}

		if cmd[0] == "forward" {
			h += v
			if aim {
				d += a * v
			}
		} else if cmd[0] == "up" {
			if aim {
				a -= v
			} else {
				d -= v
			}
		} else if cmd[0] == "down" {
			if aim {
				a += v
			} else {
				d += v
			}
		}
	}

	return h * d
}
