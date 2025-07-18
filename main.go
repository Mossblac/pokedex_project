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
		//for words in input, input[0] is
		//c.input use that to check if command exists
		//anything after add to string called "selection"
		_, exists := assist.CommandInfo[input] //c.input
		if exists {
			//"selection" is
			//the string passed to callback
			c := assist.CommandInfo[input] //c.input
			err := c.Callback(PageInfo)    // pass in "selection" to match new signatures
			if err != nil {
				fmt.Println("Error:", err)
			}

		} else {
			fmt.Println("Unknown command")
		}
	}

}
