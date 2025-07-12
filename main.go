package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello World!")
}

func cleanInput(text string) []string {
	output := []string{}
	lowered := strings.ToLower(text)
	words := strings.Fields(lowered)
	output = append(output, words...)

	return output

}
