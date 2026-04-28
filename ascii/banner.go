package process

import (
	"fmt"
	"os"
	"strings"
)

func AsciiArt(filename string) {
	if len(os.Args) < 2 {
		fmt.Println("Argument can't be less then 2. Usage go run . <string>")
		return
	}
	if len(os.Args) > 2 {
		fmt.Println("Too many arguments, wrap your input in quotes")
		return
	}
	input := os.Args[1]
	if input == "" { // handle empty input
		return
	}
	// checking if the input is only a space and if there are space before a word.
	spacesplite := strings.Split(input, " ")
	space := true

	for _, x := range spacesplite {
		if x != "" {
			space = false
		}
	}
	if space {
		return
	}
	input = strings.TrimSpace(input) // trim space before you continue with the input.
	input = strings.ReplaceAll(input, "\\n", "\n")
	var text []string
	cont, err := os.ReadFile("ascii/standard.txt")
	if err != nil {
		fmt.Println("Error occured", err)
		return
	}
	splitcont := strings.Split(string(cont), "\n")
	inputtxt := strings.Split(input, "\n")

	var start int
	for _, cha := range inputtxt {
		if cha == "" {
			fmt.Println()
		}
		for i := 0; i < 8; i++ {
			for _, words := range cha {
				if words < 32 || words > 126 {
					words = ' '
				}
				start = (int(words)-32)*9 + 1 // the beginning of the characters
				text = splitcont[start : start+8]
				fmt.Print(text[i])
			}

			fmt.Println()
		}

	}
}
