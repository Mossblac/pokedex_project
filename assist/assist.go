package assist

import (
	"fmt"
	"os"
	"strings"
)

func CleanInput(text string) []string {
	output := []string{}
	lowered := strings.ToLower(text)
	words := strings.Fields(lowered)
	output = append(output, words...)

	return output

}

var CommandInfo map[string]CliCommand

func CommandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func CommandHelp() error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n")
	for _, cmd := range CommandInfo {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}
	return nil
}

func CommandMap() error {
	// each call prints 20 locations at a time
	return nil
}

type CliCommand struct {
	Name        string
	Description string
	Callback    func() error
	Page        *Config
}

type Config struct {
	Next     string
	Previous string
}

func init() {
	CommandInfo = map[string]CliCommand{
		"exit": {
			Name:        "exit",
			Description: "EXit the Pokedex",
			Callback:    CommandExit,
		},
		"help": {
			Name:        "help",
			Description: "Displays a help message",
			Callback:    CommandHelp,
		},
		"map": {
			Name:        "map",
			Description: "Displays 20 locations per call",
			Callback:    CommandMap,
		},
	}
}
