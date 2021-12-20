package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	numBeacons, maxDistance := numBeacons("input.txt")
	fmt.Printf("Part 1: %d\nPart 2: %d\n", numBeacons, maxDistance)
}

type Coord3D struct {
	x, y, z int
}

type Mat3D [3][3]int
type Scanner []Coord3D
type Scanners []Scanner

func numBeacons(path string) (int, int) {
	scanners := parseScanners(path)
	assembled := scanners[0]
	scrambled := scanners[1:]
	shifts := []Coord3D{{0, 0, 0}}

	for len(scrambled) > 0 {
		for i, scanner := range scrambled {
			overlap, rot, shift := numOverlap(assembled, scanner)
			if overlap >= 12 {
				shifts = append(shifts, shift)
				new := make(Scanner, len(scanner))
				for j, b := range scanner {
					new[j] = rot.mul(b).add(shift)
				}
				assembled = append(assembled, new...)
				scrambled = append(scrambled[:i], scrambled[i+1:]...)
				break
			}
		}
	}

	uniqueBeacons := map[Coord3D]bool{}
	for _, beacon := range assembled {
		uniqueBeacons[beacon] = true
	}

	maxDistance := 0
	for i := 0; i < len(shifts)-1; i++ {
		for j := i + 1; j < len(shifts); j++ {
			distance := abs(shifts[i].x-shifts[j].x) + abs(shifts[i].y-shifts[j].y) + abs(shifts[i].z-shifts[j].z)
			if distance > maxDistance {
				maxDistance = distance
			}
		}
	}

	return len(uniqueBeacons), maxDistance
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func numOverlap(sc1, sc2 Scanner) (int, Mat3D, Coord3D) {
	var (
		maxOverlap int
		bestRot    Mat3D
		bestShift  Coord3D
	)

	for _, mat := range getRotations() {
		shifts := map[Coord3D]int{}
		for _, b1 := range sc1 {
			for _, b2 := range sc2 {
				shifts[b1.subtract(mat.mul(b2))]++
			}
		}

		for shift, count := range shifts {
			if count >= 12 {
				return count, mat, shift
			}
		}
	}

	// less than 12 overlapping beacons
	return maxOverlap, bestRot, bestShift
}

func (m Mat3D) mul(c Coord3D) Coord3D {
	return Coord3D{
		x: m[0][0]*c.x + m[0][1]*c.y + m[0][2]*c.z,
		y: m[1][0]*c.x + m[1][1]*c.y + m[1][2]*c.z,
		z: m[2][0]*c.x + m[2][1]*c.y + m[2][2]*c.z,
	}
}

func (m Mat3D) determinant() int {
	return (m[0][0] * (m[1][1]*m[2][2] - m[1][2]*m[2][1])) -
		(m[1][0] * (m[0][1]*m[2][2] - m[0][2]*m[2][1])) +
		(m[2][0] * (m[0][1]*m[1][2] - m[0][2]*m[1][1]))
}

func (c Coord3D) subtract(d Coord3D) Coord3D {
	return Coord3D{
		x: c.x - d.x,
		y: c.y - d.y,
		z: c.z - d.z,
	}
}

func (c Coord3D) add(d Coord3D) Coord3D {
	return Coord3D{
		x: c.x + d.x,
		y: c.y + d.y,
		z: c.z + d.z,
	}
}

// getRotations goes through all permutations, to find the 24 rotation
// matrices using the fact that rotation matrices have a determinant of 1.
func getRotations() []Mat3D {
	x := [2][3]int{
		{1, 0, 0},
		{-1, 0, 0},
	}
	y := [2][3]int{
		{0, 1, 0},
		{0, -1, 0},
	}
	z := [2][3]int{
		{0, 0, 1},
		{0, 0, -1},
	}

	potentialRotations := []Mat3D{}
	for i := range x {
		for j := range y {
			for k := range z {
				m := []Mat3D{
					{x[i], y[j], z[k]},
					{x[i], z[k], y[j]},
					{y[j], x[i], z[k]},
					{y[j], z[k], x[i]},
					{z[k], x[i], y[j]},
					{z[k], y[j], x[i]},
				}
				for _, rot := range m {
					if rot.determinant() == 1 {
						potentialRotations = append(potentialRotations, rot)
					}
				}
			}
		}
	}

	return potentialRotations
}

func parseScanners(path string) Scanners {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var (
		scanners Scanners
		scanner  Scanner
	)

	s := bufio.NewScanner(file)
	for s.Scan() {
		if s.Text() == "" {
			scanners = append(scanners, scanner)
			scanner = []Coord3D{}
		} else if strings.Split(s.Text(), " ")[0] == "---" {
			continue
		} else {
			var x, y, z int
			fmt.Sscanf(s.Text(), "%d,%d,%d", &x, &y, &z)
			scanner = append(scanner, Coord3D{x, y, z})
		}
	}
	scanners = append(scanners, scanner)

	return scanners
}

func (ss Scanners) String() string {
	var output string
	for i, scanner := range ss {
		output += fmt.Sprintf("--- scanner %d ---\n", i)
		for _, beacon := range scanner {
			output += fmt.Sprintf("%d,%d,%d\n", beacon.x, beacon.y, beacon.z)
		}
		output += "\n"
	}

	return output[:len(output)-1]
}
