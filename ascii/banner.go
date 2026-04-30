package process

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func AsciiArt(filename string) []string {
	var text []string

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal()
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		splittext := strings.Split(scanner.Text(), "\n")
		for x := range splittext {
			text = append(text, splittext[x])
		}
	}
	err = scanner.Err()
	if err != nil {
		fmt.Print("Error", err)
	}
	file.Close()
	return text
}
