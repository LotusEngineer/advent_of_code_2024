package main

import (
	_ "embed"
	"fmt"
	"strings"
)

type Antenna struct {
	X         int
	Y         int
	Frequency rune
}

//go:embed input.txt
var inputString string

func main() {
	areaMap := loadFileAsStructuredData(inputString)

	antennas := make(map[rune][][2]int)
	for x := 0; x < len(areaMap); x++ {
		for y := 0; y < len(areaMap[x]); y++ {
			cell := areaMap[x][y]
			if cell != '.' {
				antennas[cell] = append(antennas[cell], [2]int{x, y})
			}
		}
	}

	antiNodesPart1 := make(map[[2]int]bool)

	for _, positions := range antennas {
		for i := 0; i < len(positions); i++ {
			for j := i + 1; j < len(positions); j++ {
				pos1 := positions[i]
				pos2 := positions[j]

				deltaX := pos2[0] - pos1[0]
				deltaY := pos2[1] - pos1[1]

				antiNode1 := [2]int{pos1[0] - deltaX, pos1[1] - deltaY}
				antiNode2 := [2]int{pos2[0] + deltaX, pos2[1] + deltaY}

				if antiNode1[0] >= 0 && antiNode1[0] < len(areaMap) && antiNode1[1] >= 0 && antiNode1[1] < len(areaMap[0]) {
					antiNodesPart1[antiNode1] = true
				}
				if antiNode2[0] >= 0 && antiNode2[0] < len(areaMap) && antiNode2[1] >= 0 && antiNode2[1] < len(areaMap[0]) {
					antiNodesPart1[antiNode2] = true
				}
			}
		}
	}

	antiNodesPart2 := make(map[[2]int]bool)

	for _, positions := range antennas {
		for _, pos := range positions {
			antiNodesPart2[pos] = true
		}

		for i := 0; i < len(positions); i++ {
			for j := i + 1; j < len(positions); j++ {
				pos1 := positions[i]
				pos2 := positions[j]

				deltaX := pos2[0] - pos1[0]
				deltaY := pos2[1] - pos1[1]

				antiNodeRow, antiNodeCol := pos1[0], pos1[1]
				for antiNodeRow >= 0 && antiNodeRow < len(areaMap) && antiNodeCol >= 0 && antiNodeCol < len(areaMap[0]) {
					antiNodesPart2[[2]int{antiNodeRow, antiNodeCol}] = true
					antiNodeRow -= deltaX
					antiNodeCol -= deltaY
				}

				antiNodeRow, antiNodeCol = pos2[0], pos2[1]
				for antiNodeRow >= 0 && antiNodeRow < len(areaMap) && antiNodeCol >= 0 && antiNodeCol < len(areaMap[0]) {
					antiNodesPart2[[2]int{antiNodeRow, antiNodeCol}] = true
					antiNodeRow += deltaX
					antiNodeCol += deltaY
				}
			}
		}
	}

	fmt.Println("Part 1:", len(antiNodesPart1))
	fmt.Println("Part 2:", len(antiNodesPart2))
}

func loadFileAsStructuredData(inputFileString string) [][]rune {
	lines := strings.Split(strings.TrimSpace(inputFileString), "\n")
	grid := make([][]rune, len(lines))

	for lineIndex, line := range lines {
		grid[lineIndex] = []rune(line)
	}
	return grid
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
