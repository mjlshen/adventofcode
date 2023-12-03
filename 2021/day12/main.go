package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	fmt.Printf("Part 1: %d\n", numPaths("input.txt", false))
	fmt.Printf("Part 2: %d\n", numPaths("input.txt", true))
}

type Cave struct {
	name        string
	small       bool
	connections []string
}

type Caves map[string]*Cave

// numPaths returns the number of unique paths from start to end.
// Uppercase caves can be visited more than once
// Lowercase caves can only be visited once. If smallTwice is true,
// one lowercase caves can be visited twice (not including start/end).
func numPaths(input string, smallTwice bool) int {
	caves := parseCaves(input)
	paths := caves.pathsDFS(caves["start"], []string{}, smallTwice)

	return len(paths)
}

// DFS search from start to end.
func (cs Caves) pathsDFS(c *Cave, path []string, smallTwice bool) [][]string {
	path = append(path, c.name)
	if c.name == "end" {
		return [][]string{path}
	}

	paths := [][]string{}
	for _, conn := range c.connections {
		if !smallTwice {
			if !cs[conn].small || !contains(path, conn) {
				paths = append(paths, cs.pathsDFS(cs[conn], path, smallTwice)...)
			}
		} else {
			if !cs[conn].small || !containsTwice(path, conn) {
				paths = append(paths, cs.pathsDFS(cs[conn], path, smallTwice)...)
			}
		}
	}

	return paths
}

func parseCaves(path string) Caves {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	caves := map[string]*Cave{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.Split(scanner.Text(), "-")
		if _, ok := caves[line[0]]; !ok {
			caves[line[0]] = &Cave{
				name:        line[0],
				small:       unicode.IsLower(rune(line[0][0])),
				connections: []string{},
			}
		}
		// Do not add connections from cave --> start or end --> cave
		if line[1] != "start" && line[0] != "end" {
			caves[line[0]].connections = append(caves[line[0]].connections, line[1])
		}

		if _, ok := caves[line[1]]; !ok {
			caves[line[1]] = &Cave{
				name:        line[1],
				small:       unicode.IsLower(rune(line[1][0])),
				connections: []string{},
			}
		}
		// Do not add connections from cave --> start or end --> cave
		if line[0] != "start" && line[1] != "end" {
			caves[line[1]].connections = append(caves[line[1]].connections, line[0])
		}
	}

	return caves
}

// containsTwice returns false as soon as appending cave to the path causes the path
// to contain the same lowercase cave three or more times or if there are
// at least two different lowercase caves that are present two or more times.
func containsTwice(path []string, cave string) bool {
	smallCaves := map[string]int{}
	path = append(path, cave)

	for i := 0; i < len(path)-1; i++ {
		// Since the caves are always all upper/lowercase, we can just check
		// if the first letter is lowercase.
		if unicode.IsLower(rune(path[i][0])) {
			smallCaves[path[i]]++
		}
	}

	numTwice := 0
	for _, v := range smallCaves {
		if v > 1 {
			numTwice += v - 1
		}
	}

	return numTwice > 1
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}
