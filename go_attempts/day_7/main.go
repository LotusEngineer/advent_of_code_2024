package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

// Wrote this in the early hours of monday morning as I had a busy weekend, as such lazier attempt and didn't refactor in a way that returns both results for part 1 & 2. Comment line 38 for day 1 result.

//go:embed input.txt
var inputString string

func main() {
	lines := strings.Split(strings.TrimSpace(inputString), "\n")
	total := 0

	for _, line := range lines {
		parts := strings.Split(line, ":")
		target, _ := strconv.Atoi(strings.TrimSpace(parts[0]))
		numberStrings := strings.Fields(strings.TrimSpace(parts[1]))
		numbers := make([]int, len(numberStrings))
		for i, numStr := range numberStrings {
			numbers[i], _ = strconv.Atoi(numStr)
		}

		var generateOperators func(n int) [][]string
		generateOperators = func(n int) [][]string {
			if n == 0 {
				return [][]string{{}}
			}
			partial := generateOperators(n - 1)
			var result [][]string
			for _, p := range partial {
				result = append(result, append(append([]string{}, p...), "+"))
				result = append(result, append(append([]string{}, p...), "*"))
				result = append(result, append(append([]string{}, p...), "||"))
			}
			return result
		}

		matches := false
		operatorCombos := generateOperators(len(numbers) - 1)
		for _, operators := range operatorCombos {
			if evaluate(numbers, operators) == target {
				matches = true
				break
			}
		}

		if matches {
			total += target
		}
	}

	fmt.Println("Total Calibration Result:", total)
}

func evaluate(numbers []int, operators []string) int {
	result := numbers[0]
	for i, op := range operators {
		if op == "+" {
			result += numbers[i+1]
		} else if op == "*" {
			result *= numbers[i+1]
		} else if op == "||" {
			left := strconv.Itoa(result)
			right := strconv.Itoa(numbers[i+1])
			concatenated, _ := strconv.Atoi(left + right)
			result = concatenated
		}
	}
	return result
}
