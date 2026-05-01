package main

import (
	process "ascii-art/ascii"
	"fmt"
	"os"
	"strings"
)

func main() {
	var words []string
	if len(os.Args) < 2 {
		fmt.Println("Argument can't be less then 2. Usage go run . <string>")
		return
	}
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments, wrap your input in quotes")
		return
	}

	inputfile := os.Args[1]
	if inputfile == "" {
		return
	}
	// for _, val := range inputfile {
	// 	_, err := process.IsValidInput(inputfile)
	// 	if err != nil {
	// 		fmt.Println("Invalid Input", val)
	// 	}
	// }

	// To check for space in the input
	new_inputfile := strings.Split(inputfile, " ")
	isspace := true
	for _, val := range new_inputfile {
		if val != "" {
			isspace = false
		}
	}
	if isspace {
		return
	}

	inputfile = strings.TrimSpace(inputfile)
	inputfile = strings.ReplaceAll(inputfile, "\\n", "\n")

	new_text := process.AsciiArt("ascii/standard.txt")

	splitinputfile := strings.Split(inputfile, "\n")
	var start int
	for _, word := range splitinputfile {
		if word == "" {
			fmt.Println()
			return
		}
		for i := range 8 {
			for _, ch := range word {
				_, err := process.IsValidInput(string(ch))
				if err != nil {
					fmt.Println("Invalid Input", string(ch))
					return
				}
				start = (int(ch)-32)*9 + 1 // the starting point of you the character.
				words = new_text[start : start+8]
				fmt.Print(words[i])
			}
			fmt.Println()
		}
	}
}
