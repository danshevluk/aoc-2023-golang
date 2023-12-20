package main

import (
	"aoc_2023/common"
	"errors"
	"flag"
	"fmt"
	"strconv"
	"strings"
)

type Cube byte

const (
	Red Cube = iota
	Green
	Blue
)

var cubeCounts = map[Cube]int{
	Red:   12,
	Green: 13,
	Blue:  14,
}

func cubeFromString(str string) (Cube, error) {
	switch str {
	case "red":
		return Red, nil
	case "green":
		return Green, nil
	case "blue":
		return Blue, nil
	default:
		return Red, errors.New("invalid cube name")
	}
}

func main() {
	flag.Parse()
	inputContent := common.ReadInputFile()

	possibleGameIDs := make([]string, 0)

Lines:
	for _, line := range strings.Split(inputContent, "\n") {
		lineParts := strings.Split(line, ":")
		if len(lineParts) != 2 {
			fmt.Println("Invalid line:", line)
			continue
		}

		// Get game id
		gameInfoPart := strings.Split(lineParts[0], " ")
		gameID := gameInfoPart[len(gameInfoPart)-1]

		// Get game data
		attempts := strings.Split(lineParts[1], ";")
		for _, cubesData := range attempts {
			cubesDataParts := strings.Split(cubesData, ",")
			for _, cubeInfo := range cubesDataParts {
				infoItems := strings.Split(strings.TrimSpace(cubeInfo), " ")
				if len(infoItems) != 2 {
					fmt.Println("Invalid cube info:", cubeInfo)
					continue Lines
				}

				count, err := strconv.Atoi(infoItems[0])
				if err != nil {
					fmt.Println("Invalid cube count:", infoItems[0])
					continue Lines
				}
				cubeName, err := cubeFromString(infoItems[1])
				if err != nil {
					fmt.Println("Invalid cube name:", infoItems[1])
					continue Lines
				}
				if count > cubeCounts[cubeName] {
					// Impossible game
					continue Lines
				}
			}
		}
		possibleGameIDs = append(possibleGameIDs, gameID)
	}

	// Sum all possible game IDs
	result := 0
	for _, gameID := range possibleGameIDs {
		gameIDInt, err := strconv.Atoi(gameID)
		if err != nil {
			fmt.Println("Invalid game ID:", gameID)
			continue
		}
		result += gameIDInt
	}
	fmt.Println("Possible games IDs sum:", result)
}
