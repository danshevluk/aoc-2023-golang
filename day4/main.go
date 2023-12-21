package main

import (
	"aoc_2023/common"
	"errors"
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
	fmt.Println("Part 2 solution: ", solvePart2(inputContent))
}

func numberSetsFromString(numbersString string) ([]int, []int, error) {
	numberSets := strings.Split(numbersString, "|")
	if len(numberSets) != 2 {
		return nil, nil, errors.New("invalid number sets")
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
	return winningNumbers, cardNumbers, nil
}

// Part 1

func solvePart1(input string) int {
	lines := strings.Split(input, "\n")

	gamesListScore := 0
	for _, line := range lines {
		parts := strings.Split(line, ":")
		if len(parts) != 2 {
			fmt.Println("Invalid line:", line)
			continue
		}

		winningNumbers, cardNumbers, err := numberSetsFromString(parts[1])
		if err != nil {
			fmt.Println("Failed to extract numbers from string:", parts[1])
			continue
		}
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

// Part 2

func solvePart2(input string) int {
	lines := strings.Split(input, "\n")

	numberOfTicketsByType := make([]int, len(lines))
	for i := range numberOfTicketsByType {
		numberOfTicketsByType[i] = 1
	}

	for lineIndex, line := range lines {
		lineParts := strings.Split(line, ":")
		if len(lineParts) != 2 {
			fmt.Println("Invalid line:", line)
			continue
		}

		winningNumbers, cardNumbers, err := numberSetsFromString(lineParts[1])
		if err != nil {
			fmt.Println("Failed to extract numbers from string:", lineParts[1])
			continue
		}

		foundWinningNumbersCount := 0
		for _, cardNumber := range cardNumbers {
			if slices.Contains(winningNumbers, cardNumber) {
				foundWinningNumbersCount += 1
			}
		}

		if foundWinningNumbersCount == 0 {
			continue
		}

		// Save new ticket counts that we just won
		nextTicketIndex := lineIndex + 1
		if nextTicketIndex >= len(lines) {
			continue
		}

		// Each of tickets wins plus one for each next ticket
		numberOfTicketsOfThisType := numberOfTicketsByType[lineIndex]
		for index := range numberOfTicketsByType[nextTicketIndex:min(nextTicketIndex+foundWinningNumbersCount, len(lines))] {
			ticketTypeIndex := nextTicketIndex + index
			numberOfTicketsByType[ticketTypeIndex] += numberOfTicketsOfThisType
		}
	}

	totalNumberOfTickets := 0
	for _, numberOfTickets := range numberOfTicketsByType {
		totalNumberOfTickets += numberOfTickets

	}

	return totalNumberOfTickets
}
