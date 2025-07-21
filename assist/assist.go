package assist

import (
	"strings"
	"time"

	"github.com/Mossblac/pokedexcli/internal"
	"github.com/Mossblac/pokedexcli/internal/pokecache"
)

func CleanInput(text string) []string {
	output := []string{}
	lowered := strings.ToLower(text)
	words := strings.Fields(lowered)
	output = append(output, words...)

	return output

}

var PokeCatalogue map[string]internal.Poke

var CommandInfo map[string]CliCommand

var cache *pokecache.Cache

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*Config, string) error
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
		"explore": {
			Name:        "explore",
			Description: "type 'explore' then area to explore to list pokemon in area",
			Callback:    CommandEx,
		},
		"catch": {
			Name:        "catch",
			Description: "type 'catch' then name of pokemon to throw pokeball",
			Callback:    CommandCatch,
		},
		"inspect": {
			Name:        "inspect",
			Description: "type 'inspect' then name of caught pokemon to get stats",
			Callback:    CommandInspect,
		},
		"pokedex": {
			Name:        "pokedex",
			Description: "view list of captured pokemon",
			Callback:    CommandPokedex,
		},
	}

	cache = pokecache.NewCache(5 * time.Minute)
	PokeCatalogue = make(map[string]internal.Poke)
}
