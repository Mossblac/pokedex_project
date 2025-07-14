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

func CommandExit() error {
	fmt.Printf("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

type CliCommand struct {
	name        string
	description string
	callback    func() error
}

var CommandMap = map[string]CliCommand{
	"exit": {
		name:        "exit",
		description: "EXit the Pokedex",
		callback:    CommandExit,
	},
}
