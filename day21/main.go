package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Printf("Part 1: %d\n", deterministicDie("input.txt"))
	fmt.Printf("Part 2: %d\n", part2("input.txt"))
}

type Player struct {
	space, score int
}

type Multiverse struct {
	// Memoization cache for computed simulations
	cache map[[2]Player][2]int
}

func deterministicDie(path string) int {
	p1, p2 := parse(path)
	die, rolls := 0, 0
	p1Turn := true

	for p1.score < 1000 && p2.score < 1000 {
		moves := 0
		for roll := 0; roll < 3; roll++ {
			rolls++
			die++
			if die > 100 {
				die = 1
			}
			moves += die
		}

		if p1Turn {
			p1 = p1.move(moves)
		} else {
			p2 = p2.move(moves)
		}
		p1Turn = !p1Turn
	}

	if p1.score > p2.score {
		return p2.score * rolls
	}
	return p2.score * rolls
}

func part2(path string) int {
	p1, p2 := parse(path)
	m := &Multiverse{cache: map[[2]Player][2]int{}}
	p1Wins, p2Wins := m.quantumDie(p1, p2)

	if p1Wins > p2Wins {
		return p1Wins
	}
	return p2Wins
}

func (m *Multiverse) quantumDie(p1, p2 Player) (p1Wins int, p2Wins int) {
	if p1.score > 20 {
		return 1, 0
	}

	if p2.score > 20 {
		return 0, 1
	}

	if ans, ok := m.cache[[2]Player{p1, p2}]; ok {
		return ans[0], ans[1]
	}

	for roll, multiverses := range multiverseMap() {
		p1c, p2c := m.quantumDie(p2, p1.move(roll))
		p1Wins += p2c * multiverses
		p2Wins += p1c * multiverses
	}

	m.cache[[2]Player{p1, p2}] = [2]int{p1Wins, p2Wins}
	return p1Wins, p2Wins
}

// multiverseMap returns the number of ways three rolls of a three-sided
// die will return a specific sum.
func multiverseMap() map[int]int {
	return map[int]int{
		3: 1, // (1, 1, 1)
		4: 3, // (1, 1, 2)
		5: 6, // (1, 2, 2), (1, 1, 3)
		6: 7, // (1, 2, 3), (2, 2, 2)
		7: 6, // (1, 3, 3), (2, 2, 3)
		8: 3, // (2, 3, 3)
		9: 1, // (3, 3, 3)
	}
}

func (p Player) move(steps int) Player {
	p.space = (p.space + steps) % 10
	if p.space == 0 {
		p.space = 10
	}
	p.score += p.space
	return p
}

func parse(path string) (Player, Player) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var p1, p2 int
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "Player 1 starting position: %d", &p1)
	scanner.Scan()
	fmt.Sscanf(scanner.Text(), "Player 2 starting position: %d", &p2)

	player1 := Player{space: p1, score: 0}
	player2 := Player{space: p2, score: 0}

	return player1, player2
}
