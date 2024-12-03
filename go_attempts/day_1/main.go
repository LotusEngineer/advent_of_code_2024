package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	leftList, rightList, err := readFile("./input.txt")

	totalDistance := 0
	matchesMap := make(map[int]int)

	if err != nil {
		panic(err)
	}

	slices.Sort(leftList)
	slices.Sort(rightList)

	for _, num := range rightList {
		matchesMap[num]++
	}

	similarityScore := 0

	for i := 0; i < len(leftList); i++ {
		currentLeftListItem := leftList[i]
		currentRightListItem := rightList[i]

		currentRowDifference := currentLeftListItem - currentRightListItem

		if currentRowDifference < 0 {
			currentRowDifference = -currentRowDifference
		}

		totalDistance += currentRowDifference

		similarityScore += currentLeftListItem * matchesMap[currentLeftListItem]
	}

	fmt.Printf("Total Distance: %d\nTotal Matches: %d\n", totalDistance, similarityScore)
}

func readFile(filename string) ([]int, []int, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, nil, err
	}

	var leftList, rightList []int

	lines := strings.Split(string(data), "\n")
	for _, line := range lines {
		if strings.TrimSpace(line) == "" {
			continue
		}

		parts := strings.Fields(line)
		if len(parts) != 2 {
			return nil, nil, fmt.Errorf("invalid line format: %s", line)
		}

		left, err1 := strconv.Atoi(parts[0])
		right, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			return nil, nil, fmt.Errorf("invalid numbers on line: %s", line)
		}

		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}

	return leftList, rightList, nil
}
