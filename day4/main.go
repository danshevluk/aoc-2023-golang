package main

import (
	"aoc_2023/common"
	"flag"
	"fmt"
	"math"
	"slices"
	"strconv"
	"strings"
)

func main() {
	flag.Parse()
	inputContent := common.ReadInputFile()

	fmt.Println("Part 1 solution: ", solvePart1(inputContent))
}

func solvePart1(input string) int {
	lines := strings.Split(input, "\n")

	gamesListScore := 0
	for _, line := range lines {
		lineParts := strings.Split(line, ":")
		if len(lineParts) != 2 {
			fmt.Println("Invalid line:", line)
			continue
		}

		numberSets := strings.Split(lineParts[1], "|")
		if len(numberSets) != 2 {
			fmt.Println("Invalid number sets:", numberSets)
			continue
		}
		numberSetToArray := func(numberSet string) []int {
			numberSet = strings.TrimSpace(numberSet)
			if len(numberSet) == 0 {
				return make([]int, 0)
			}

			numbersArray := make([]int, 0)
			for _, number := range strings.Split(numberSet, " ") {
				number = strings.TrimSpace(number)
				if len(number) == 0 {
					continue
				}

				numberInt, err := strconv.Atoi(number)
				if err != nil {
					fmt.Println("Invalid number:", number)
					continue
				}

				numbersArray = append(numbersArray, numberInt)
			}

			return numbersArray
		}

		winningNumbers := numberSetToArray(numberSets[0])
		cardNumbers := numberSetToArray(numberSets[1])

		foundWinningNumbersCount := 0
		for _, cardNumber := range cardNumbers {
			if slices.Contains(winningNumbers, cardNumber) {
				foundWinningNumbersCount += 1
			}
		}
		score := 0
		if foundWinningNumbersCount == 1 {
			score = 1
		} else if foundWinningNumbersCount > 1 {
			score += int(math.Pow(2, float64(foundWinningNumbersCount-1)))
		}
		gamesListScore += score
	}

	return gamesListScore
}
