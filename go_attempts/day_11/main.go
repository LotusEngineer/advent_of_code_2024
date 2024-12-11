package main

import (
	_ "embed"
	"fmt"
	"math"
	"strconv"
	"strings"
	"time"
)

//go:embed input.txt
var inputString string

func main() {
	start := time.Now()
	stonesInput := strings.Fields(strings.TrimSpace(inputString))

	stoneValues := make(map[int]int)
	for _, s := range stonesInput {
		stoneValue, _ := strconv.Atoi(s)
		stoneValues[stoneValue]++
	}

	transformationCache := make(map[int][]int)
	transformationCache[0] = []int{1}

	for i := 0; i < 75; i++ {
		nextStoneValues := make(map[int]int)

		for stoneValue, count := range stoneValues {
			if _, found := transformationCache[stoneValue]; !found {
				transformationCache[stoneValue] = transformStone(stoneValue)
			}

			transformedStones := transformationCache[stoneValue]
			for _, newStoneValue := range transformedStones {
				nextStoneValues[newStoneValue] += count
			}
		}

		stoneValues = nextStoneValues
	}

	totalStonesCount := 0
	for _, count := range stoneValues {
		totalStonesCount += count
	}

	elapsed := time.Since(start)

	fmt.Println("Total final stones (75 iterations):", totalStonesCount)
	fmt.Println("Time taken:", elapsed)
}

func transformStone(stoneValue int) []int {
	if stoneValue == 0 {
		return []int{1}
	}

	valueLength := digitLength(stoneValue)
	if valueLength%2 == 0 {
		halfLength := valueLength / 2
		tenPower := int(math.Pow10(halfLength))

		leftStoneValue := stoneValue / tenPower
		rightStoneValue := stoneValue % tenPower
		return []int{leftStoneValue, rightStoneValue}
	} else {
		return []int{stoneValue * 2024}
	}
}

func digitLength(n int) int {
	if n == 0 {
		return 1
	}
	length := 0
	for n > 0 {
		n /= 10
		length++
	}
	return length
}

func Part1() {
	stones := strings.Fields(strings.TrimSpace(inputString))

	for i := 0; i < 25; i++ {
		newStones := make([]string, 0, len(stones)*2)

		for _, stone := range stones {
			if stone == "0" {
				newStones = append(newStones, "1")
			} else if len(stone)%2 == 0 {
				midpoint := len(stone) / 2
				firstHalf := strings.TrimLeft(stone[:midpoint], "0")
				secondHalf := strings.TrimLeft(stone[midpoint:], "0")

				if firstHalf == "" {
					firstHalf = "0"
				}
				if secondHalf == "" {
					secondHalf = "0"
				}

				newStones = append(newStones, firstHalf, secondHalf)
			} else {
				stoneValue, _ := strconv.Atoi(stone)
				finalNewValue := strconv.Itoa(stoneValue * 2024)
				newStones = append(newStones, finalNewValue)
			}
		}

		stones = newStones
	}

	fmt.Println("Total final stones (25 iterations):", len(stones))
}
