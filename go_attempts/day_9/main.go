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
	inputString = strings.TrimSpace(inputString)

	blockArrayPart1 := createBlockArrayFromInput(inputString)
	compactBlocks(blockArrayPart1)
	checksumPart1 := calculateChecksum(blockArrayPart1)
	fmt.Println("Part 1 Checksum:", checksumPart1)

	blockArrayPart2 := createBlockArrayFromInput(inputString)
	compactFiles(blockArrayPart2)
	checksumPart2 := calculateChecksum(blockArrayPart2)
	fmt.Println("Part 2 Checksum:", checksumPart2)
}

func createBlockArrayFromInput(input string) []rune {
	var blockArray []rune
	isAddingBlocks := true
	currentFileID := 0

	for _, inputRune := range input {
		blockCount, _ := strconv.Atoi(string(inputRune))

		if isAddingBlocks {
			blockArray = appendRepeatedRune(blockArray, rune('0'+currentFileID), blockCount)
			currentFileID++
		} else {
			blockArray = appendRepeatedRune(blockArray, '.', blockCount)
		}

		isAddingBlocks = !isAddingBlocks
	}

	return blockArray
}

func appendRepeatedRune(blockArray []rune, char rune, count int) []rune {
	for i := 0; i < count; i++ {
		blockArray = append(blockArray, char)
	}
	return blockArray
}

func compactBlocks(blockArray []rune) {
	leftPointer := 0
	rightPointer := len(blockArray) - 1

	for blockArray[leftPointer] != '.' {
		leftPointer++
	}

	for blockArray[rightPointer] == '.' {
		rightPointer--
	}

	for leftPointer <= rightPointer {
		blockArray[leftPointer], blockArray[rightPointer] = blockArray[rightPointer], '.'

		for blockArray[leftPointer] != '.' && leftPointer <= rightPointer {
			leftPointer++
		}

		for blockArray[rightPointer] == '.' && rightPointer >= leftPointer {
			rightPointer--
		}
	}
}

func compactFiles(blockArray []rune) {
	fileIDPositions := getFileIDPositions(blockArray)

	for currentFileID := len(fileIDPositions) - 1; currentFileID >= 0; currentFileID-- {
		fileBlocks := fileIDPositions[currentFileID]

		if len(fileBlocks) == 0 {
			continue
		}

		freeStart, freeLength := findLeftmostContiguousFreeSpace(blockArray, len(fileBlocks), fileBlocks[0])

		if freeLength == len(fileBlocks) {
			for i := 0; i < freeLength; i++ {
				blockArray[fileBlocks[i]] = '.'
				blockArray[freeStart+i] = rune('0' + currentFileID)
			}
		}
	}
}

func getFileIDPositions(blockArray []rune) map[int][]int {
	fileIDPositions := make(map[int][]int)
	for index, block := range blockArray {
		if block != '.' {
			fileID := int(block - '0')
			fileIDPositions[fileID] = append(fileIDPositions[fileID], index)
		}
	}
	return fileIDPositions
}

func findLeftmostContiguousFreeSpace(blockArray []rune, requiredSpace int, limitIndex int) (int, int) {
	freeStart := -1
	freeLength := 0

	for i := 0; i < limitIndex; i++ {
		if blockArray[i] != '.' {
			freeStart = -1
			freeLength = 0
			continue
		}

		if freeStart == -1 {
			freeStart = i
		}

		freeLength++

		if freeLength == requiredSpace {
			return freeStart, freeLength
		}
	}

	return -1, 0
}

func calculateChecksum(blockArray []rune) int {
	totalChecksum := 0
	for index, block := range blockArray {
		if block != '.' {
			fileID := int(block - '0')
			totalChecksum += index * fileID
		}
	}
	return totalChecksum
}
