package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func main() {
	fmt.Printf("Part 1: %d\n", enhanceImg("input.txt", 2))
	fmt.Printf("Part 2: %d\n", enhanceImg("input.txt", 50))
}

type Coord struct {
	x, y int
}

type TrenchImage struct {
	img                map[Coord]bool
	minX, maxX         int
	minY, maxY         int
	minBound, maxBound Coord
}

func enhanceImg(path string, times int) int {
	alg, img := parse(path)
	for i := 0; i < times; i++ {
		img = img.enhance(alg)
	}
	return img.score()
}

func (t TrenchImage) enhance(alg []bool) TrenchImage {
	img := make(map[Coord]bool)
	minX, maxX, minY, maxY := t.minX, t.maxX, t.minY, t.maxY

	for x := t.minBound.x; x <= t.maxBound.x; x++ {
		for y := t.minBound.y; y <= t.maxBound.y; y++ {
			conversion := ""
			for i := x - 1; i <= x+1; i++ {
				for j := y - 1; j <= y+1; j++ {
					if val, ok := t.img[Coord{i, j}]; ok && val {
						conversion += "1"
					} else {
						conversion += "0"
					}
				}
			}
			newPixel := alg[binaryToDecimal(conversion)]
			if newPixel {
				if x >= t.minX-1 && x <= t.maxX+1 && y >= t.minY-1 && y <= t.maxY+1 {
					if x < minX {
						minX = x
					}
					if x > maxX {
						maxX = x
					}
					if y < minY {
						minY = y
					}
					if y > maxY {
						maxY = y
					}
				}
			}
			img[Coord{x, y}] = newPixel
		}
	}

	return TrenchImage{
		img:      img,
		minX:     minX,
		maxX:     maxX,
		minY:     minY,
		maxY:     maxY,
		minBound: t.minBound,
		maxBound: t.maxBound,
	}
}

func parse(path string) ([]bool, TrenchImage) {
	file, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	pixelMap := map[Coord]bool{}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	imgEnhanceAlg := make([]bool, len(scanner.Text()))
	for i, c := range scanner.Text() {
		imgEnhanceAlg[i] = (c == '#')
	}

	for x := 0; scanner.Scan(); x++ {
		if scanner.Text() == "" {
			x--
			continue
		} else {
			for y, c := range scanner.Text() {
				if c == '#' {
					pixelMap[Coord{x, y}] = true
				}
			}
		}
	}

	minX, minY := math.MaxInt, math.MaxInt
	maxX, maxY := math.MinInt, math.MinInt
	for k, v := range pixelMap {
		if v {
			if k.x < minX {
				minX = k.x
			}
			if k.x > maxX {
				maxX = k.x
			}
			if k.y < minY {
				minY = k.y
			}
			if k.y > maxY {
				maxY = k.y
			}
		}
	}

	return imgEnhanceAlg, TrenchImage{
		img:      pixelMap,
		minX:     minX,
		maxX:     maxX,
		minY:     minY,
		maxY:     maxY,
		minBound: Coord{minX - 100, minY - 100},
		maxBound: Coord{maxX + 100, maxY + 100},
	}
}

func (t TrenchImage) score() int {
	var score int
	for k, v := range t.img {
		if k.x >= t.minX && k.x <= t.maxX && k.y >= t.minY && k.y <= t.maxY {
			if v {
				score++
			}
		}
	}

	return score
}

func (t TrenchImage) String() string {
	var s string
	for x := t.minX; x <= t.maxX; x++ {
		for y := t.minY; y <= t.maxY; y++ {
			if v, ok := t.img[Coord{x, y}]; ok && v {
				s += "#"
			} else {
				s += "."
			}
		}
		s += "\n"
	}
	return s
}

func binaryToDecimal(bin string) int {
	var result int
	for _, b := range bin {
		switch b {
		case '0':
			result = result*2 + 0
		case '1':
			result = result*2 + 1
		}
	}
	return result
}
