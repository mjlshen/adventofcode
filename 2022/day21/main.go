package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", monkeyMath("input.txt"))
	//fmt.Printf("Part 2: %d\n", monkeyMath("input.txt", 2))
}

type monkey struct {
	first, second string
	op            string
	val           int
	yelled        bool
}

func monkeyMath(path string) int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}

	monkeys := map[string]*monkey{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var name string
		var m *monkey

		if _, err := fmt.Sscanf(scanner.Text(), "%s %s %s %s", &name, &m.first, &m.op, &m.second); err != nil {
			if _, nerr := fmt.Sscanf(scanner.Text(), "%s %d", &name, &m.val); nerr != nil {
				panic(nerr)
			}
		}

		monkeys[strings.TrimSuffix(name, ":")] = m
	}

	return 0
}

func solve(ans string, monkeys map[string]*monkey) int {
	stack := []string{ans}
	visited := map[string]struct{}{}

	return monkeys[ans].val
}
