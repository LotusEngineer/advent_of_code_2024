package main

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputString string

type Rule struct {
	Before int
	After  int
}

func main() {
	fmt.Println("Loading file as structured data...")
	rules, updates := loadFileAsStructuredData(inputString)

	totalValidMiddlePageValue := 0
	totalFixedMiddlePageValue := 0

	// fireDebugInfo(rules, updates)

	fmt.Println("Creating page dependencies map...")
	pageDependencies := make(map[int][]int)
	for _, rule := range rules {
		pageDependencies[rule.Before] = append(pageDependencies[rule.Before], rule.After)
	}

	fmt.Println("Processing updates...")
	for _, update := range updates {
		fmt.Printf("Processing update: %v\n", update)
		if validateUpdate(update, pageDependencies) {
			fmt.Println("Update is valid.")
			updateMiddleValue := update[len(update)/2]
			fmt.Printf("Middle value of valid update: %d\n", updateMiddleValue)
			totalValidMiddlePageValue += updateMiddleValue
		} else {
			fmt.Println("Update is invalid, attempting to reorder...")
			reorderedUpdate := reorderUpdate(update, pageDependencies)
			fmt.Printf("Reordered update: %v\n", reorderedUpdate)
			if validateUpdate(reorderedUpdate, pageDependencies) {
				fmt.Println("Reordered update is valid.")
				reorderedUpdateMiddleValue := reorderedUpdate[len(reorderedUpdate)/2]
				fmt.Printf("Middle value of reordered update: %d\n", reorderedUpdateMiddleValue)
				totalFixedMiddlePageValue += reorderedUpdateMiddleValue
			} else {
				fmt.Println("Reordered update is still invalid. Detailed validation log:")
				logValidationFailure(reorderedUpdate, pageDependencies)
			}
		}
	}

	fmt.Printf("Valid middle page value total: %d\nCorrected middle page value total: %d\n", totalValidMiddlePageValue, totalFixedMiddlePageValue)
}

func reorderUpdate(update []int, pageDependencies map[int][]int) []int {
	fmt.Printf("Reordering update: %v\n", update)

	remaining := make(map[int]bool)
	dependencies := make(map[int][]int)

	for _, num := range update {
		remaining[num] = true
		for _, dep := range pageDependencies[num] {
			dependencies[dep] = append(dependencies[dep], num)
		}
	}

	ordered := []int{}

	for len(remaining) > 0 {
		for num := range remaining {
			if len(dependencies[num]) == 0 {
				ordered = append(ordered, num)
				delete(remaining, num)

				for key, val := range dependencies {
					newList := []int{}
					for _, n := range val {
						if n != num {
							newList = append(newList, n)
						}
					}
					dependencies[key] = newList
				}
			}
		}
	}

	fmt.Printf("Sorted update: %v\n", ordered)
	return ordered
}

func sort(graph map[int][]int) []int {
	fmt.Println("Calculating dependencies...")
	dependencies := make(map[int]int)
	for _, neighbors := range graph {
		for _, neighbor := range neighbors {
			dependencies[neighbor]++
		}
	}

	fmt.Println("Initializing queue with nodes that have no dependencies...")
	queue := []int{}
	for node := range graph {
		if dependencies[node] == 0 {
			queue = append(queue, node)
		}
	}

	var sortedOrder []int
	visited := make(map[int]bool)
	for len(queue) > 0 {
		current := queue[0]
		queue = queue[1:]
		sortedOrder = append(sortedOrder, current)
		visited[current] = true
		fmt.Printf("Current node: %d, queue: %v\n", current, queue)

		for _, neighbor := range graph[current] {
			dependencies[neighbor]--
			if dependencies[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	fmt.Printf("Sorted order: %v\n", sortedOrder)
	return sortedOrder
}

func validateUpdate(update []int, pageDependencies map[int][]int) bool {
	fmt.Printf("Validating update: %v\n", update)
	position := make(map[int]int)
	for i, page := range update {
		position[page] = i
	}

	for before, afterList := range pageDependencies {
		for _, after := range afterList {
			if posBefore, ok := position[before]; ok {
				if posAfter, ok := position[after]; ok {
					if posBefore > posAfter {
						fmt.Printf("Invalid update: page %d comes before page %d\n", before, after)
						return false
					}
				}
			}
		}
	}

	fmt.Println("Update is valid.")
	return true
}

func logValidationFailure(update []int, pageDependencies map[int][]int) {
	position := make(map[int]int)
	for i, page := range update {
		position[page] = i
	}

	for before, afterList := range pageDependencies {
		for _, after := range afterList {
			if posBefore, ok := position[before]; ok {
				if posAfter, ok := position[after]; ok {
					if posBefore > posAfter {
						fmt.Printf("Validation failed: page %d comes before page %d in update %v\n", before, after, update)
					}
				}
			}
		}
	}
}

func loadFileAsStructuredData(inputFileString string) ([]Rule, [][]int) {
	fmt.Println("Loading rules and updates from input file...")
	parts := strings.Split(inputFileString, "\n\n")

	var rules []Rule
	ruleLines := strings.Split(parts[0], "\n")

	for _, line := range ruleLines {
		fmt.Printf("Parsing rule line: %s\n", line)
		ruleParts := strings.Split(line, "|")
		before, _ := strconv.Atoi(ruleParts[0])
		after, _ := strconv.Atoi(ruleParts[1])

		rules = append(rules, Rule{
			Before: before,
			After:  after,
		})
	}

	updateLines := strings.Split(parts[1], "\n")

	var updates [][]int

	for _, line := range updateLines {
		fmt.Printf("Parsing update line: %s\n", line)
		updateValueSlice := strings.Split(line, ",")
		var convertedUpdateValueSlice []int
		for _, updateItem := range updateValueSlice {
			convertedUpdateItem, _ := strconv.Atoi(updateItem)
			convertedUpdateValueSlice = append(convertedUpdateValueSlice, convertedUpdateItem)
		}
		updates = append(updates, convertedUpdateValueSlice)
	}

	fmt.Printf("Loaded %d rules and %d updates\n", len(rules), len(updates))
	return rules, updates
}

func fireDebugInfo(rules []Rule, updates [][]int) {
	fmt.Println("Rules:")
	for i, rule := range rules {
		fmt.Printf("  Rule %d: Before = %d, After = %d\n", i+1, rule.Before, rule.After)
	}

	fmt.Println("\nUpdates:")
	for i, update := range updates {
		fmt.Printf("  Update %d: %v\n", i+1, update)
	}
}
