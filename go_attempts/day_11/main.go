package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputString string

func main() {
	stones := strings.Fields(strings.TrimSpace(inputString))

	for i := 0; i < 75; i++ {
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

	fmt.Println("Total final stones:", len(stones))
}
