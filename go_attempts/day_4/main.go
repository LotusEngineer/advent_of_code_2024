package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var inputString string

func main() {
	word := "XMAS"
	grid := loadFileAsStructuredData(inputString)

	xmasCount := 0

	xShapedMasCount := 0

	directions := [][2]int{
		{0, 1},
		{0, -1},
		{-1, 0},
		{1, 0},
		{-1, 1},
		{-1, -1},
		{1, -1},
		{1, 1},
	}

	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid[row]); column++ {
			for _, direction := range directions {
				rowIndex := direction[0]
				columnIndex := direction[1]
				foundWord := true

				for charIndex := 0; charIndex < len(word); charIndex++ {
					rowOffset := row + (rowIndex * charIndex)
					columnOffset := column + (columnIndex * charIndex)

					if rowOffset < 0 || rowOffset >= len(grid) || columnOffset < 0 || columnOffset >= len(grid[row]) {
						foundWord = false
						break
					}

					if grid[rowOffset][columnOffset] != rune(word[charIndex]) {
						foundWord = false
						break
					}
				}

				if foundWord {
					xmasCount++
				}
			}
		}
	}

	for row := 1; row < len(grid)-1; row++ {
		for col := 1; col < len(grid[row])-1; col++ {
			if isXMasCenter(grid, row, col) {
				xShapedMasCount++
			}
		}
	}

	fmt.Printf("Xmas occurences: %d\n X shaped mas occurences: %d\n", xmasCount, xShapedMasCount)
}

func loadFileAsStructuredData(inputFileString string) [][]rune {
	lines := strings.Split(strings.TrimSpace(inputFileString), "\n")

	grid := make([][]rune, len(lines))

	for lineIndex, line := range lines {
		grid[lineIndex] = []rune(line)
	}
	return grid
}

func isXMasCenter(grid [][]rune, row, col int) bool {
	if grid[row][col] != 'A' {
		return false
	}

	topLeft := [2]int{row - 1, col - 1}
	topRight := [2]int{row - 1, col + 1}
	bottomLeft := [2]int{row + 1, col - 1}
	bottomRight := [2]int{row + 1, col + 1}

	if isWithinBounds(grid, topLeft) && isWithinBounds(grid, topRight) &&
		isWithinBounds(grid, bottomLeft) && isWithinBounds(grid, bottomRight) {

		if (grid[topLeft[0]][topLeft[1]] == 'M' && grid[bottomRight[0]][bottomRight[1]] == 'S') ||
			(grid[topLeft[0]][topLeft[1]] == 'S' && grid[bottomRight[0]][bottomRight[1]] == 'M') {
			if (grid[topRight[0]][topRight[1]] == 'S' && grid[bottomLeft[0]][bottomLeft[1]] == 'M') ||
				(grid[topRight[0]][topRight[1]] == 'M' && grid[bottomLeft[0]][bottomLeft[1]] == 'S') {
				return true
			}
		}
	}

	return false
}

func isWithinBounds(grid [][]rune, position [2]int) bool {
	row, col := position[0], position[1]
	return row >= 0 && row < len(grid) && col >= 0 && col < len(grid[row])
}
