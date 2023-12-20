package main

import (
	"aoc_2023/common"
	"flag"
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	flag.Parse()

	inputContent := common.ReadInputFile()

	fmt.Println("Solution for Part 1:", solvePart1(inputContent))
}

func solvePart1(inputContent string) int {
	lines := strings.Split(inputContent, "\n")

	partNumbers := make([]int, 0)
	for lineIndex, line := range lines {
		numberStart := -1

		for charIndex, char := range line {
			if isDigit(char) {
				if numberStart == -1 {
					numberStart = charIndex
				}

				// Continue if we are not at the end of the line
				if charIndex != len(line)-1 {
					continue
				}
			}

			if numberStart != -1 { // Found end of a number or end of a line
				numberEnd := charIndex
				if charIndex == len(line)-1 && isDigit(char) {
					numberEnd = len(line)
				}
				number, err := strconv.Atoi(line[numberStart:numberEnd])
				if err != nil {
					fmt.Println("Error:", err)
					numberStart = -1
					continue
				}
				fmt.Println("Checking number:", number)
				if hasAdjacentSymbols(&lines, lineIndex, numberStart, numberEnd) {
					fmt.Println("Found adjusted symbols for:", number)
					partNumbers = append(partNumbers, number)
				}
				numberStart = -1
			}
		}
	}

	partNumbersSum := 0
	for _, number := range partNumbers {
		partNumbersSum += number
	}
	return partNumbersSum
}

func isSymbol(char rune) bool {
	return char != '.' && !isDigit(char)
}

func isSymbolAtIndex(line string, index int) bool {
	if index < 0 || index >= len(line) {
		return false
	}
	runeAtIndex, _ := utf8.DecodeRuneInString(line[index:])
	fmt.Printf("Checking rune %c at index %d\n", runeAtIndex, index)
	return isSymbol(runeAtIndex)
}

func isDigit(char rune) bool {
	return unicode.IsDigit(char)
}

func hasAdjacentSymbols(schematic *[]string, lineIndex int, numberStart int, numberEnd int) bool {
	// Check around number on the same line
	if isSymbolAtIndex((*schematic)[lineIndex], numberStart-1) || isSymbolAtIndex((*schematic)[lineIndex], numberEnd) {
		return true
	}

	hasSymbolsInLine := func(line string) bool {
		for _, char := range line[max(0, numberStart-1):min(numberEnd+1, len(line))] {
			if isSymbol(char) {
				return true
			}
		}

		return false
	}

	// Check line above
	if lineIndex > 0 {
		upperLine := (*schematic)[lineIndex-1]
		if hasSymbolsInLine(upperLine) {
			return true
		}
	}

	// Check line below
	if lineIndex < len(*schematic)-1 {
		lowerLine := (*schematic)[lineIndex+1]
		if hasSymbolsInLine(lowerLine) {
			return true
		}
	}

	return false
}
