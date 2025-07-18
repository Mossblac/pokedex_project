package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/Mossblac/pokedexcli/assist"
)

//eterna-city-area

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
		words := strings.Fields(input)
		var cInput string
		var selection string
		if len(words) == 1 {
			cInput = words[0]
			selection = ""
		} else {
			cInput = words[0]
			selection = words[1]
		}
		if len(words) > 2 {
			fmt.Printf("invalid input: single word only, use dashes for areas\n")
		} else {
			_, exists := assist.CommandInfo[cInput]
			if exists {
				c := assist.CommandInfo[cInput]
				err := c.Callback(PageInfo, selection)
				if err != nil {
					fmt.Println("Error:", err)
				}

			} else {
				fmt.Println("Unknown command")
			}
		}
	}
}
