package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Printf("Part 1: %d\n", firstBingo("input.txt"))
	fmt.Printf("Part 2: %d\n", lastBingo("input.txt"))
}

// firstBingo returns the score of the first winning board
func firstBingo(path string) int {
	moves, boards := parse(path)
	for _, move := range moves {
		for _, board := range boards {
			board.play(move)
			if board.won() {
				return board.score()
			}
		}
	}
	// Should never get here
	return 0
}

// lastBingo returns the score of the last winning board
func lastBingo(path string) int {
	moves, boards := parse(path)
	winners := make([]bool, len(boards))
	for _, move := range moves {
		for i, board := range boards {
			board.play(move)
			if board.won() {
				winners[i] = true
			}
			if allTrue(winners) {
				return board.score()
			}
		}
	}
	// Should never get here
	return 0
}

// Coord stores the row/col coordinate of a number on a BingoBoard
type Coord struct {
	row, col int
}

// BingoBoard tracks the last move and the remaining numbers
type BingoBoard struct {
	// numbers is a map with the key being the number and the value being the
	// row/col coordinate of the number on the board
	numbers map[int]Coord
	// Played is a 2d array of booleans indicating whether a space has been played
	played [5][5]bool
	// lastmove stores the last move played
	lastmove int
}

func (b *BingoBoard) play(number int) {
	if k, ok := b.numbers[number]; ok {
		b.played[k.row][k.col] = true
		b.lastmove = number
		delete(b.numbers, number)
	}
}

func (b *BingoBoard) won() bool {
	// Check rows
	for i := range b.played {
		if b.played[i][0] && b.played[i][1] && b.played[i][2] && b.played[i][3] && b.played[i][4] {
			return true
		}
	}
	// Check columns
	for j := range b.played {
		if b.played[0][j] && b.played[1][j] && b.played[2][j] && b.played[3][j] && b.played[4][j] {
			return true
		}
	}
	return false
}

// The score for a winning BingoBoard is the sum of all unmarked numbers on the board
// times the last called number.
func (b *BingoBoard) score() int {
	score := 0
	for k := range b.numbers {
		score += k
	}
	return score * b.lastmove
}

func allTrue(a []bool) bool {
	if len(a) == 0 {
		return false
	}

	for v := range a {
		if !a[v] {
			return false
		}
	}
	return true
}

func parse(path string) ([]int, []*BingoBoard) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if scanner.Text() != "" {
			lines = append(lines, scanner.Text())
		}
	}

	// Parse the moves
	s := strings.Split(lines[0], ",")
	moves := make([]int, len(s))
	for i := range s {
		moves[i], _ = strconv.Atoi(s[i])
	}

	// Initialize n bingo boards
	boards := make([]*BingoBoard, (len(lines)-1)/5)
	for i := range boards {
		boards[i] = &BingoBoard{
			numbers: make(map[int]Coord),
		}
	}

	// Parse the boards
	for i := 1; i < len(lines); i++ {
		s := strings.Fields(lines[i])
		for j := range s {
			n, _ := strconv.Atoi(s[j])
			boards[(i-1)/5].numbers[n] = Coord{row: (i - 1) % 5, col: j}
		}
	}

	return moves, boards
}
