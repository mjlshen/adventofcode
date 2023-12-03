package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	fmt.Printf("Part 1: %d\n", (gammaEpsilon("input.txt")))
	fmt.Printf("Part 2: %d\n", (o2CO2("input.txt", 12)))
}

// gammaEpsilon returns the product of gamma and epsilon
// based on an input of binary numbers of equal length.
func gammaEpsilon(path string) int {
	var (
		numOnes []int
		gamma   int
		epsilon int
	)

	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	numLines := 0
	for scanner.Scan() {
		for i, c := range scanner.Text() {
			if len(numOnes) <= i {
				numOnes = append(numOnes, 0)
			}
			if c == rune('1') {
				numOnes[i] += 1
			}
		}
		numLines++
	}

	// gamma in binary has digits represented by the most common digit from the input
	for i, c := range numOnes {
		if c > numLines/2 {
			gamma += int(math.Pow(2, float64(len(numOnes)-1-i)))
		}
	}

	// epsilon in binary has digits represented by the least common digit from the input
	// which makes it the one's complement of gamma
	epsilon = (int(math.Pow(2, float64(len(numOnes)))) - 1) - gamma

	return gamma * epsilon
}

// Readings is a struct that holds the binary readings, sorted based on a sort depth
// For example, 0010 with a sort depth of 0 would be sorted into Zeros, but with a
// sort depth of 2, would be sorted into Ones.
type Readings struct {
	// Zeros holds all of the binary readings that have a 0 at the sort depth
	Zeros []string
	// Ones holds all of the binary readings that have a 1 at the sort depth
	Ones []string
	// SortDepth is the digit that is used to differentiate the readings
	SortDepth int
	// Length is the length of the binary strings, they must all have the same length
	Length int
}

// Insert places a reading into the correct slice based on the sort depth
func (rs *Readings) insert(reading string) {
	if []rune(reading)[rs.SortDepth] == '1' {
		rs.Ones = append(rs.Ones, reading)
	} else {
		rs.Zeros = append(rs.Zeros, reading)
	}
}

// O2CO2 returns the product of the O2 and CO2 ratings
func o2CO2(path string, length int) int {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	readings := &Readings{
		SortDepth: 0,
		Length:    length,
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		readings.insert(scanner.Text())
	}

	return binaryToInt(readings.O2()) * binaryToInt(readings.CO2())
}

func binaryToInt(binary string) int {
	num, err := strconv.ParseInt(binary, 2, 64)
	if err != nil {
		panic(err)
	}

	return int(num)
}

// O2 iterates through, picking the Zeros/Ones slice with the most elements
// with ties going to the Ones slice, until there is one reading left.
func (rs *Readings) O2() string {
	for {
		if rs.SortDepth >= rs.Length {
			break
		}
		readings := &Readings{
			SortDepth: rs.SortDepth + 1,
			Length:    rs.Length,
		}
		if len(rs.Ones) >= len(rs.Zeros) {
			if len(rs.Ones) == 1 {
				return rs.Ones[0]
			}
			for _, r := range rs.Ones {
				readings.insert(r)
			}
		} else {
			if len(rs.Zeros) == 1 {
				return rs.Zeros[0]
			}
			for _, r := range rs.Zeros {
				readings.insert(r)
			}
		}

		rs = readings
	}

	return ""
}

// CO2 iterates through, picking the Zeros/Ones slice with the fewest elements
// with ties going to the Zeros slice, until there is one reading left.
func (rs *Readings) CO2() string {
	for {
		if rs.SortDepth >= rs.Length {
			break
		}
		readings := &Readings{
			SortDepth: rs.SortDepth + 1,
			Length:    rs.Length,
		}
		if len(rs.Zeros) <= len(rs.Ones) {
			if len(rs.Zeros) == 1 {
				return rs.Zeros[0]
			}
			for _, r := range rs.Zeros {
				readings.insert(r)
			}
		} else {
			if len(rs.Ones) == 1 {
				return rs.Ones[0]
			}
			for _, r := range rs.Ones {
				readings.insert(r)
			}
		}

		rs = readings
	}

	return ""
}
