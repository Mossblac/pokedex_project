package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		if input != "" {
			input = strings.ToLower(input)
			fields := strings.Fields(input)
			firstWord := fields[0]
			fmt.Printf("Your command was: %v\n", firstWord)
		}

	}
}

func cleanInput(text string) []string {
	output := []string{}
	lowered := strings.ToLower(text)
	words := strings.Fields(lowered)
	output = append(output, words...)

	return output

}
