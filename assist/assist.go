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
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

type CliCommand struct {
	Name        string
	Description string
	Callback    func() error
}

var CommandMap = map[string]CliCommand{
	"exit": {
		Name:        "exit",
		Description: "EXit the Pokedex",
		Callback:    CommandExit,
	},
}
