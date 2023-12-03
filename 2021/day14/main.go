package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	fmt.Printf("Part 1: %d\n", polymerInsertion("input.txt", 10))
	fmt.Printf("Part 2: %d\n", polymerInsertion("input.txt", 40))
}

type Rules map[string]string
type Pairs map[string]int

// Given input, performs "steps" insertion steps on the parsed polymer
// using the parsed rules, then returns the score of the resulting polymer.
func polymerInsertion(path string, steps int) int {
	polymer, rules := parsePolymer(path)

	pairs := map[string]int{}
	for i := 0; i < len(polymer)-1; i++ {
		pairs[polymer[i:i+2]]++
	}

	for i := 0; i < steps; i++ {
		pairs = polymerize(pairs, rules)
	}

	return score(pairs, polymer[len(polymer)-1:])
}

// polymerize evaluates pairs on provided rules, splitting them when there
// is a match.
// Given a pair "NN" and a rule "NN -> B", the new map of pairs will contains
// "NB" and "BN" with the same value.
func polymerize(pairs Pairs, rules Rules) Pairs {
	newPairs := map[string]int{}
	for k := range pairs {
		if _, ok := rules[k]; ok {
			// There's an applicable rule, so split this pair
			newPairOne := string(k[0]) + rules[k]
			newPairTwo := rules[k] + string(k[1])
			newPairs[newPairOne] += pairs[k]
			newPairs[newPairTwo] += pairs[k]
		} else {
			// No rules match this pair
			newPairs[k] += pairs[k]
		}
	}

	return newPairs
}

// score computes the difference between the most and least common letter
// in the polymer
func score(pairs Pairs, last string) int {
	// Since there are overlapping letters between pairs,
	// only count the first letter of each pair and manually add
	// the last letter.
	letters := map[string]int{last: 1}
	for i := range pairs {
		letters[string(i[0])] += pairs[i]
	}

	max := 0
	min := math.MaxInt64
	for i := range letters {
		if letters[i] < min {
			min = letters[i]
		}
		if letters[i] > max {
			max = letters[i]
		}
	}

	return max - min
}

func parsePolymer(path string) (string, Rules) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	// First line is the polymer
	scanner.Scan()
	polymer := scanner.Text()

	// The rest of the lines are rules
	rules := map[string]string{}
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		var pair, insert string
		fmt.Sscanf(scanner.Text(), "%s -> %s", &pair, &insert)
		rules[pair] = insert
	}

	return polymer, rules
}
