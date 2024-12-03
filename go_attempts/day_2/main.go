package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var reports [][]int

	for scanner.Scan() {
		reports = append(reports, parseReport(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	part1Count := 0
	for _, report := range reports {
		if isSafe(report) {
			part1Count++
		}
	}
	fmt.Printf("Total safe reports: %d\n", part1Count)

	part2Count := 0
	for _, report := range reports {
		if isSafeWithDampener(report) {
			part2Count++
		}
	}
	fmt.Printf("Total safe reports w/ dampener: %d\n", part2Count)
}

func isSafe(report []int) bool {
	increasing, decreasing := true, true
	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]
		if math.Abs(float64(diff)) < 1 || math.Abs(float64(diff)) > 3 {
			return false
		}
		if diff > 0 {
			decreasing = false
		}
		if diff < 0 {
			increasing = false
		}
	}
	return increasing || decreasing
}

func isSafeWithDampener(report []int) bool {
	if isSafe(report) {
		return true
	}
	for i := 0; i < len(report); i++ {
		temp := append([]int{}, report[:i]...)
		temp = append(temp, report[i+1:]...)
		if isSafe(temp) {
			return true
		}
	}
	return false
}

func parseReport(line string) []int {
	parts := strings.Fields(line)
	report := make([]int, len(parts))
	for i, part := range parts {
		num, _ := strconv.Atoi(part)
		report[i] = num
	}
	return report
}
