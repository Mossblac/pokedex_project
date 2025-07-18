package assist

import (
	"encoding/json"
	"fmt"
	"os"
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

var CommandInfo map[string]CliCommand

var cache *pokecache.Cache

//each command also accepts a string, even though "explore"
//will be the only Command that uses it

func CommandExit(*Config, string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func CommandHelp(*Config, string) error {
	fmt.Printf("Welcome to the Pokedex!\nUsage:\n")
	for _, cmd := range CommandInfo {
		fmt.Printf("%s: %s\n", cmd.Name, cmd.Description)
	}
	return nil
}

func CommandMap(cfg *Config, s string) error {
	var AreaInfo internal.AreaStruct
	Data, ok := cache.Get(cfg.Next)
	if ok {
		//fmt.Println("Cache Hit - using cached data")
		err := json.Unmarshal(Data, &AreaInfo)
		if err != nil {
			return err
		}
	} else {
		//fmt.Println("Cache MISS - making network request")
		var err error
		AreaInfo, err = internal.CreateGoStruct(cfg.Next)
		if err != nil {
			return err
		}

		Data, err := json.Marshal(AreaInfo)
		if err != nil {
			return err
		}

		cache.Add(cfg.Next, Data)
	}
	for _, name := range AreaInfo.Results {
		fmt.Printf("%s\n", name.Name)
	}

	cfg.Next = AreaInfo.Next
	cfg.Previous = AreaInfo.Previous

	return nil
}

func CommandMapb(cfg *Config, s string) error {
	if cfg.Previous == nil {
		fmt.Print("your're on the first page\n")
	} else {
		var AreaInfo internal.AreaStruct
		Data, ok := cache.Get(*cfg.Previous)
		if ok {
			//fmt.Println("Cache Hit - using cached data")
			err := json.Unmarshal(Data, &AreaInfo)
			if err != nil {
				return err
			}
		} else {
			//fmt.Println("Cache MISS - making network request")
			var err error
			AreaInfo, err = internal.CreateGoStruct(*cfg.Previous)
			if err != nil {
				return err
			}

			Data, err := json.Marshal(AreaInfo)
			if err != nil {
				return err
			}

			cache.Add(*cfg.Previous, Data)
		}
		for _, name := range AreaInfo.Results {
			fmt.Printf("%s\n", name.Name)
		}
		cfg.Next = AreaInfo.Next
		cfg.Previous = AreaInfo.Previous
	}
	return nil
}

func CommandEx(cfg *Config, selection string) error {
	fmt.Printf("explore function called with input: %v\n", selection)
	return nil
}

/*"explore" triggers the callback
using the cache- makes a call to the same
url with name or id (selection string) added to end.
convert response to struct to obtain info to print*/

type CliCommand struct {
	Name        string
	Description string
	Callback    func(*Config, string) error //change this to accept a string also
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
	}

}

func init() {
	cache = pokecache.NewCache(5 * time.Minute)
}
