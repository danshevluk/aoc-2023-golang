package main

import (
	"aoc_2023/common"
	"errors"
	"flag"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"
)

func main() {
	isVerbose := flag.Bool("v", false, "Verbose")
	flag.Parse()

	inputContent := common.ReadInputFile()

	calibrationValuesTotal := 0
	for _, line := range strings.Split(inputContent, "\n") {
		if *isVerbose {
			fmt.Println("Line:", line)
		}
		firstDigit, err := getFristDigit(line, true)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		secondDigit, err := getFristDigit(line, false)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		calibrationValue := firstDigit*10 + secondDigit
		calibrationValuesTotal += calibrationValue
	}

	fmt.Println("Calibration values total:", calibrationValuesTotal)
}

func getFristDigit(line string, forward bool) (int, error) {
	var digit *int
	forEachRune(line, forward, func(charRune rune) error {
		if unicode.IsDigit(charRune) {
			convertedRune := int(charRune - '0')
			digit = &convertedRune
		}
		if digit != nil {
			return errors.New("Loop finished")
		}

		return nil
	})

	if digit == nil {
		return 0, errors.New("No digit found")
	}

	return *digit, nil
}

func forEachRune(str string, forward bool, block func(r rune) error) {
	if forward {
		for _, charRune := range str {
			err := block(charRune)
			if err != nil {
				break
			}
		}
	} else {
		lastIndex := len(str) - 1

		for i := lastIndex; i >= 0; {
			charRune, size := utf8.DecodeLastRuneInString(str[:i+1])
			if charRune != utf8.RuneError || size > 0 {
				err := block(charRune)
				if err != nil {
					break
				}
			} else {
				fmt.Printf("Invalid index %d\n", i)
				break
			}

			i -= size
		}
	}
}
