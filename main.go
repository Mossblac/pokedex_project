package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/Mossblac/pokedexcli/assist"
	"github.com/Mossblac/pokedexcli/internal"
)

func main() {
	AreaInfo, err := internal.CreateGoStruct("https://pokeapi.co/api/v2/location-area/")
	if err != nil {
		log.Fatalf("%v", err)
	}
	PageInfo := &assist.Config{
		Next:     AreaInfo.Next,
		Previous: AreaInfo.Previous,
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
