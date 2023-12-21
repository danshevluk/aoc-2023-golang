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
	fmt.Println("Solution for Part 2:", solvePart2(inputContent))
}

// Part 1

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

// Part 2

type SchematicIndex struct {
	line  int
	index int
}

func solvePart2(inputContent string) int {
	lines := strings.Split(inputContent, "\n")

	gearRaiosSum := 0
	for _, adjustedNumbers := range buildAdjustedStarsMap(&lines) {
		if len(adjustedNumbers) != 2 {
			continue
		}
		gearRaiosSum += adjustedNumbers[0] * adjustedNumbers[1]
	}

	return gearRaiosSum
}

func buildAdjustedStarsMap(schematic *[]string) map[SchematicIndex][]int {
	numbersAdjustedToStars := make(map[SchematicIndex][]int)
	for lineIndex, line := range *schematic {
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
				for _, starIndex := range indexesOfAdjustedStars(schematic, lineIndex, numberStart, numberEnd) {
					numbersAdjustedToStars[starIndex] = append(numbersAdjustedToStars[starIndex], number)
				}
				numberStart = -1
			}
		}
	}
	return numbersAdjustedToStars
}

func indexesOfAdjustedStars(schematic *[]string, lineIndex int, numberStart int, numberEnd int) []SchematicIndex {
	const starSymbol = '*'

	starsIndexies := make([]SchematicIndex, 0)

	// Check around number on the same line
	currentLine := (*schematic)[lineIndex]
	leftSideIndex := max(numberStart-1, 0)
	if currentLine[leftSideIndex] == starSymbol {
		starsIndexies = append(starsIndexies, SchematicIndex{lineIndex, leftSideIndex})
	}
	rightSideIndex := min(numberEnd, len(currentLine)-1)
	if currentLine[rightSideIndex] == starSymbol {
		starsIndexies = append(starsIndexies, SchematicIndex{lineIndex, rightSideIndex})
	}

	starIndexiesInLine := func(line string, i int) []SchematicIndex {
		indexies := make([]SchematicIndex, 0)
		startOffset := max(0, numberStart-1)
		for index, char := range line[startOffset:min(numberEnd+1, len(line))] {
			if char == starSymbol {
				indexies = append(indexies, SchematicIndex{i, index + startOffset})
			}

		}
		return indexies
	}

	if lineIndex > 0 {
		upperLine := (*schematic)[lineIndex-1]
		starsIndexies = append(starsIndexies, starIndexiesInLine(upperLine, lineIndex-1)...)
	}

	if lineIndex < len(*schematic)-1 {
		lowerLine := (*schematic)[lineIndex+1]
		starsIndexies = append(starsIndexies, starIndexiesInLine(lowerLine, lineIndex+1)...)
	}

	return starsIndexies
}
