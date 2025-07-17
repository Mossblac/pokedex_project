package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/Mossblac/pokedexcli/assist"
)

func main() {
	PageInfo := &assist.Config{
		Next:     "https://pokeapi.co/api/v2/location-area/",
		Previous: nil,
	}
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		input := scanner.Text()
		_, exists := assist.CommandInfo[input]
		if exists {
			c := assist.CommandInfo[input]
			err := c.Callback(PageInfo)
			if err != nil {
				fmt.Println("Error:", err)
			}

		} else {
			fmt.Println("Unknown command")
		}
	}

}
