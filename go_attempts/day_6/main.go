package main

import (
	_ "embed"
	"fmt"
	"strings"
)

//go:embed input.txt
var inputString string

func main() {
	patrolGrid := loadFileAsStructuredData(inputString)

	directions := [][2]int{
		{-1, 0}, // North
		{0, 1},  // East
		{1, 0},  // South
		{0, -1}, // West
	}

	possibleGuardOrientations := map[rune]string{
		'<': "W",
		'>': "E",
		'^': "N",
		'v': "S",
	}

	directionIndexes := map[string]int{
		"N": 0,
		"E": 1,
		"S": 2,
		"W": 3,
	}

	// Part 1: Calculate total distinct guard locations
	currentGuardPositionX, currentGuardPositionY, guardOrientation := findGuardInitialPosition(patrolGrid, possibleGuardOrientations)
	currentDirectionIndex := directionIndexes[guardOrientation]
	totalVisitedCells := 0
	visitedCells := make(map[[2]int]bool)

	visitedCells[[2]int{currentGuardPositionX, currentGuardPositionY}] = true
	totalVisitedCells++

	for {
		currentVelocityX := directions[currentDirectionIndex][0]
		currentVelocityY := directions[currentDirectionIndex][1]

		nextRowIndex := currentGuardPositionX + currentVelocityX
		nextColumnIndex := currentGuardPositionY + currentVelocityY

		if nextRowIndex < 0 || nextRowIndex >= len(patrolGrid) || nextColumnIndex < 0 || nextColumnIndex >= len(patrolGrid[0]) {
			break
		}

		if patrolGrid[nextRowIndex][nextColumnIndex] == '#' {
			currentDirectionIndex = (currentDirectionIndex + 1) % 4
			continue
		}

		currentGuardPositionX = nextRowIndex
		currentGuardPositionY = nextColumnIndex

		currentState := [2]int{currentGuardPositionX, currentGuardPositionY}
		if !visitedCells[currentState] {
			totalVisitedCells++
			visitedCells[currentState] = true
		}
	}

	fmt.Printf("Total distinct guard locations: %d\n", totalVisitedCells)

	// Part 2: Detect possible looping positions
	initialGuardPositionX, initialGuardPositionY, initialGuardOrientation := findGuardInitialPosition(patrolGrid, possibleGuardOrientations)
	initialDirectionIndex := directionIndexes[initialGuardOrientation]
	detectLoopingPositions(patrolGrid, initialGuardPositionX, initialGuardPositionY, initialDirectionIndex, directions)
}

func loadFileAsStructuredData(inputFileString string) [][]rune {
	lines := strings.Split(strings.TrimSpace(inputFileString), "\n")
	grid := make([][]rune, len(lines))

	for lineIndex, line := range lines {
		grid[lineIndex] = []rune(line)
	}
	return grid
}

func findGuardInitialPosition(grid [][]rune, orientations map[rune]string) (int, int, string) {
	for rowIndex := range grid {
		for columnIndex := range grid[rowIndex] {
			if value, exists := orientations[grid[rowIndex][columnIndex]]; exists {
				return rowIndex, columnIndex, value
			}
		}
	}
	return -1, -1, "N"
}

func causesLoop(grid [][]rune, startX, startY, startDirection int, directions [][2]int) bool {
	visitedLocations := make(map[[3]int]bool)
	currentX, currentY, currentDirection := startX, startY, startDirection

	for {
		guardState := [3]int{currentX, currentY, currentDirection}
		if visitedLocations[guardState] {
			return true
		}

		visitedLocations[guardState] = true

		nextX := currentX + directions[currentDirection][0]
		nextY := currentY + directions[currentDirection][1]

		if nextX < 0 || nextX >= len(grid) || nextY < 0 || nextY >= len(grid[0]) {
			return false
		}

		if grid[nextX][nextY] == '#' {
			currentDirection = (currentDirection + 1) % 4
		} else {
			currentX = nextX
			currentY = nextY
		}
	}
}

func detectLoopingPositions(patrolGrid [][]rune, startX, startY, startDirection int, directions [][2]int) {
	loopCausingObstacles := make(map[[2]int]bool)
	totalPotentialNewObstacles := 0

	for row := 0; row < len(patrolGrid); row++ {
		for col := 0; col < len(patrolGrid[row]); col++ {
			if patrolGrid[row][col] != '.' {
				continue
			}

			patrolGrid[row][col] = '#'

			if causesLoop(patrolGrid, startX, startY, startDirection, directions) {
				loopCausingObstacles[[2]int{row, col}] = true
				totalPotentialNewObstacles++
			}

			patrolGrid[row][col] = '.'
		}
	}

	fmt.Printf("Potential new obstacle locations: %d\n", totalPotentialNewObstacles)
}
