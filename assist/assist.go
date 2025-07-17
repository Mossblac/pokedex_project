package assist

import (
	"fmt"
	"os"
	"strings"

	"github.com/Mossblac/pokedexcli/internal"
)

func CleanInput(text string) []string {
	output := []string{}
	lowered := strings.ToLower(text)
	words := strings.Fields(lowered)
	output = append(output, words...)

	return output

}

var CommandInfo map[string]CliCommand

func CommandExit(*Config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func CommandHelp(*Config) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n")
	for _, cmd := range CommandInfo {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}
	return nil
}

func CommandMap(cfg *Config) error {
	AreaInfo, err := internal.CreateGoStruct(cfg.Next)
	if err != nil {
		return err
	}
	for _, name := range AreaInfo.Results {
		fmt.Printf("%s\n", name.Name)
	}
	cfg.Next = AreaInfo.Next
	cfg.Previous = AreaInfo.Previous

	return nil
}

func CommandMapb(cfg *Config) error {
	if cfg.Previous == nil {
		fmt.Print("your're on the first page\n")
	} else {
		AreaInfo, err := internal.CreateGoStruct(*cfg.Previous)
		if err != nil {
			return err
		}
		for _, name := range AreaInfo.Results {
			fmt.Printf("%s\n", name.Name)
		}
		cfg.Next = AreaInfo.Next
		cfg.Previous = AreaInfo.Previous
	}
	return nil
}

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*Config) error
}

type Config struct {
	Next     string
	Previous *string
}

func init() {
	CommandInfo = map[string]CliCommand{
		"exit": {
			Name:        "exit",
			Description: "Exit the Pokedex",
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
		"mapb": {
			Name:        "mapb",
			Description: "Displays previous 20 entries",
			Callback:    CommandMapb,
		},
	}

}
