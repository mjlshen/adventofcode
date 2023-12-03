package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Part 1: %d\n", pyroclasticFlow("input.txt"))
	//fmt.Printf("Part 2: %d\n", pyroclasticFlow("input.txt", true))
}

type coord struct {
	x, y int
}

const chamberWidth = 7

func pyroclasticFlow(path string) int {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return 0
}

func generateRocks() [][][]bool {
	return [][][]bool{
		{
			{true, true, true, true},
		},
		{
			{false, true, false, false},
			{true, true, true, false},
			{false, true, false, false},
		},
		{
			{false, false, true, false},
			{false, false, true, false},
			{true, true, true, false},
		},
		{
			{true, false, false, false},
			{true, false, false, false},
			{true, false, false, false},
			{true, false, false, false},
		},
		{
			{true, true, false, false},
			{true, true, false, false},
		},
	}
}
