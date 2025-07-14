package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Mossblac/pokedexcli/assist"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		_, exists := assist.CommandMap[input]
		if exists {

		}

	}

}
