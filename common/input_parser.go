package common

import (
	"flag"
	"fmt"
	"os"
)

func ReadInputFile() string {
	args := flag.Args()
	if len(args) < 1 {
		fmt.Println("Error: Please provide the input file path as a command-line argument.")
		os.Exit(1)
	}

	inputFilePath := args[0]

	contents, err := os.ReadFile(inputFilePath)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	return string(contents)
}
