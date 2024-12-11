package main

import (
	_ "embed"
	"fmt"
	"strings"
)

type CellLocation struct {
	X int
	Y int
}

type Traversal struct {
	Location CellLocation
	Path     []CellLocation
}

//go:embed input.txt
var inputString string

func main() {
	directions := [][2]int{
		{0, 1},  // Right
		{0, -1}, // Left
		{-1, 0}, // Up
		{1, 0},  // Down
	}
	overallMapTrailScore := 0
	overallTrailRating := 0
	trailMap := loadFileAsStructuredData(inputString)
	var trailHeads []CellLocation

	for rowIndex := 0; rowIndex < len(trailMap); rowIndex++ {
		for columnIndex := 0; columnIndex < len(trailMap[0]); columnIndex++ {
			if trailMap[rowIndex][columnIndex] == 0 {
				trailHeads = append(trailHeads, CellLocation{X: rowIndex, Y: columnIndex})
			}
		}
	}

	for _, trailHead := range trailHeads {
		locationsToTraverse := []Traversal{{trailHead, []CellLocation{trailHead}}}

		uniqueTargets := make(map[CellLocation]bool)
		distinctPaths := make(map[string]bool)
		trailHeadScore := 0
		trailRating := 0

		for len(locationsToTraverse) > 0 {
			current := locationsToTraverse[0]
			locationsToTraverse = locationsToTraverse[1:]

			currentLocation := current.Location
			currentPath := current.Path

			if trailMap[currentLocation.X][currentLocation.Y] == 9 {
				if !uniqueTargets[currentLocation] {
					uniqueTargets[currentLocation] = true
					trailHeadScore++
				}

				pathKey := hashPath(currentPath)
				if !distinctPaths[pathKey] {
					distinctPaths[pathKey] = true
					trailRating++
				}
				continue
			}

			for _, direction := range directions {
				potentialNextStep := CellLocation{X: currentLocation.X + direction[0], Y: currentLocation.Y + direction[1]}
				if isValidNextStep(trailMap, currentLocation, potentialNextStep) {
					newPath := append([]CellLocation{}, currentPath...)
					newPath = append(newPath, potentialNextStep)
					locationsToTraverse = append(locationsToTraverse, Traversal{
						Location: potentialNextStep,
						Path:     newPath,
					})
				}
			}
		}

		overallMapTrailScore += trailHeadScore
		overallTrailRating += trailRating
	}

	fmt.Println("Total map traversal score (Part 1):", overallMapTrailScore)
	fmt.Println("Total trail ratings (Part 2):", overallTrailRating)
}

func isValidNextStep(trailMap [][]int, from CellLocation, to CellLocation) bool {
	if to.X < 0 || to.Y < 0 || to.X >= len(trailMap) || to.Y >= len(trailMap[0]) {
		return false
	}

	differenceInHeight := trailMap[to.X][to.Y] - trailMap[from.X][from.Y]
	return differenceInHeight == 1
}

func loadFileAsStructuredData(inputFileString string) [][]int {
	lines := strings.Split(strings.TrimSpace(inputFileString), "\n")
	grid := make([][]int, len(lines))

	for lineIndex, line := range lines {
		grid[lineIndex] = make([]int, len(line))
		for charIndex, char := range line {
			grid[lineIndex][charIndex] = int(char - '0')
		}
	}
	return grid
}

func hashPath(path []CellLocation) string {
	var sb strings.Builder
	for _, loc := range path {
		sb.WriteString(fmt.Sprintf("%d,%d->", loc.X, loc.Y))
	}
	return sb.String()
}
