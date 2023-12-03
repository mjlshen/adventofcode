package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	fmt.Printf("Part 1: %d\n", maxModelNumber(generateRuleTable(parseRules("input.txt"))))
	fmt.Printf("Part 2: %d\n", minModelNumber(generateRuleTable(parseRules("input.txt"))))
}

type Rules [14][3]int
type RuleTable [7][4]int

func maxModelNumber(rt RuleTable) int {
	var digits [14]int

	for _, rule := range rt {
		// Maximize the digit that appears first
		for i := 9; i > 0; i-- {
			lessSignificantDigit := i + rule[1] + rule[2]
			if lessSignificantDigit > 0 && lessSignificantDigit < 10 {
				digits[rule[0]] = i
				digits[rule[3]] = lessSignificantDigit
				break
			}
		}
	}

	ans := 0
	for i := range digits {
		ans *= 10
		ans += digits[i]
	}
	return ans
}

func minModelNumber(rt RuleTable) int {
	var digits [14]int

	for _, rule := range rt {
		// Minimize the digit that appears first
		for i := 1; i < 10; i++ {
			lessSignificantDigit := i + rule[1] + rule[2]
			if lessSignificantDigit > 0 && lessSignificantDigit < 10 {
				digits[rule[0]] = i
				digits[rule[3]] = lessSignificantDigit
				break
			}
		}
	}

	ans := 0
	for i := range digits {
		ans *= 10
		ans += digits[i]
	}
	return ans
}

func generateRuleTable(rs Rules) RuleTable {
	var push []int
	var pushpop [][2]int
	for i, rule := range rs {
		if rule[0] == 1 {
			push = append(push, i)
		} else {
			pushpop = append(pushpop, [2]int{push[len(push)-1], i})
			push = push[:len(push)-1]
		}
	}

	var rt [7][4]int
	for i := range pushpop {
		rt[i][0] = pushpop[i][0]
		rt[i][1] = rs[pushpop[i][0]][2]
		rt[i][2] = rs[pushpop[i][1]][1]
		rt[i][3] = pushpop[i][1]
	}

	return rt
}

// The instructions follow an 18-line pattern, with only three lines varying
func repeatedInstructions() []string {
	return []string{
		"inp w",
		"mul x 0",
		"add x z",
		"mod x 26",
		// "div z 1/26",
		"",
		// "add x %d",
		"",
		"eql x w",
		"eql x 0",
		"mul y 0",
		"add y 25",
		"mul y x",
		"add y 1",
		"mul z y",
		"mul y 0",
		"add y w",
		// "add y %d",
		"",
		"mul y x",
		"add z y",
	}
}

func parseRules(path string) Rules {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var rules [14][3]int
	scanner := bufio.NewScanner(file)
	for rule := 0; rule < 14; rule++ {
		j := 0
		for i := 0; i < 18; i++ {
			scanner.Scan()
			if repeatedInstructions()[i] == "" {
				var a, b string
				// Don't care about a, b, but forced to use something as a placeholder
				fmt.Sscanf(scanner.Text(), "%s %s %d", &a, &b, &rules[rule][j])
				j++
			}
		}
	}

	return rules
}

type Instruction struct {
	action string
	first  string
	second string
}

type ALU struct {
	vars map[string]int
}

func feedInstructions(instructions []Instruction, input []int) bool {
	for _, i := range input {
		if i == 0 {
			return false
		}
	}

	alu := ALU{
		vars: map[string]int{},
	}
	j := 0
	for _, i := range instructions {
		if i.action == "inp" {
			if _, ok := alu.vars["z"]; ok {
				fmt.Println(alu.vars["z"])
			}
			i.second = strconv.Itoa(input[j])
			j++
		}
		alu.execute(i)
	}

	return alu.isValid()
}

func (a *ALU) execute(i Instruction) {
	switch i.action {
	case "inp":
		val, err := strconv.Atoi(i.second)
		if err != nil {
			panic(err)
		}
		a.vars[i.first] = val
	case "add":
		val, err := strconv.Atoi(i.second)
		if err == nil {
			a.vars[i.first] += val
		} else {
			a.vars[i.first] += a.vars[i.second]
		}
	case "mul":
		val, err := strconv.Atoi(i.second)
		if err == nil {
			a.vars[i.first] *= val
		} else {
			a.vars[i.first] *= a.vars[i.second]
		}
	case "div":
		val, err := strconv.Atoi(i.second)
		if err != nil {
			val = a.vars[i.second]
		}

		if val == 0 {
			panic("divide by zero")
		}

		ans := a.vars[i.first] / val
		if ans < 0 {
			ans++
		}
		a.vars[i.first] = ans
	case "mod":
		val, err := strconv.Atoi(i.second)
		if err == nil {
			a.vars[i.first] %= val
		} else {
			a.vars[i.first] %= a.vars[i.second]
		}
	case "eql":
		val, err := strconv.Atoi(i.second)
		if err != nil {
			val = a.vars[i.second]
		}

		if a.vars[i.first] == val {
			a.vars[i.first] = 1
		} else {
			a.vars[i.first] = 0
		}
	}
}

func (a ALU) isValid() bool {
	if v, ok := a.vars["z"]; ok {
		if v == 0 {
			return true
		}
	}
	return false
}

func parseInstructions(path string) []Instruction {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	is := []Instruction{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var (
			action, first, second string
		)

		_, err := fmt.Sscanf(scanner.Text(), "%s %s %s", &action, &first, &second)
		if err != nil {
			_, err = fmt.Sscanf(scanner.Text(), "%s %s", &action, &first)
			if err != nil {
				panic(err)
			}
		}
		is = append(is, Instruction{action, first, second})
	}

	return is
}
