package assist

import (
	"fmt"
)

func CommandHelp(*Config, string) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n")
	for _, cmd := range CommandInfo {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}
	return nil
}
