package main

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

//go:embed input.txt
var inputString string

func main() {
	regexPattern := `mul\(\d{1,3},\d{1,3}\)|\bdo\(\)|\bdon't\(\)`

	regex := regexp.MustCompile(regexPattern)

	validCommands := regex.FindAllString(inputString, -1)

	multResults := 0

	enableMulCommand := true

	for _, validCommand := range validCommands {
		if strings.Contains(validCommand, "mul") && enableMulCommand {
			arguments := regexp.MustCompile(`\d+`).FindAllString(validCommand, -1)
			if len(arguments) == 2 {
				x, _ := strconv.Atoi(arguments[0])
				y, _ := strconv.Atoi(arguments[1])
				multResults += x * y
			}
		} else if strings.Contains(validCommand, "don't()") {
			enableMulCommand = false
		} else if strings.Contains(validCommand, "do()") {
			enableMulCommand = true
		}
	}

	fmt.Printf("Corrupted commands result: %d\n", multResults)
}
