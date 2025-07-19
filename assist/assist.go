package assist

import (
	"encoding/json"
	"fmt"
	"math/rand"
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

var PokeCatalogue map[string]internal.Poke

var CommandInfo map[string]CliCommand

var cache *pokecache.Cache

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
	if selection == "" {
		fmt.Printf("invalid entry: input 'explore' and name of area to explore\n")
	} else {
		fmt.Printf("Exploring %s...\n", selection)
		fmt.Printf("Found Pokemon: \n")

	}
	var Ex internal.Explore
	baseUrl := "https://pokeapi.co/api/v2/location-area/"
	fullUrl := baseUrl + selection
	Data, ok := cache.Get(fullUrl)
	if ok {
		err := json.Unmarshal(Data, &Ex)
		if err != nil {
			return err
		}

	} else {

		var err error
		Ex, err = internal.CreateExploreStruct(fullUrl)
		if err != nil {
			return err
		}

		Data, err := json.Marshal(Ex)
		if err != nil {
			return err
		}

		cache.Add(fullUrl, Data)
	}

	pokeEncounters := Ex.PokemonEncounters

	for _, pokemon := range pokeEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}
	return nil
}

func CommandCatch(cfg *Config, selection string) error {
	if selection == "" {
		fmt.Printf("invalid entry: input 'catch' and name of pokemon\n")
	} else {
		fmt.Printf("Throwing a Pokeball at %s...\n", selection)
	}
	var poke internal.Poke
	pUrl := "https://pokeapi.co/api/v2/pokemon/"
	fullUrl := pUrl + selection
	Data, ok := cache.Get(fullUrl)
	if ok {
		err := json.Unmarshal(Data, &poke)
		if err != nil {
			return err
		}

	} else {

		var err error
		poke, err = internal.CreatePokeStruct(fullUrl)
		if err != nil {
			return err
		}

		Data, err := json.Marshal(poke)
		if err != nil {
			return err
		}

		cache.Add(fullUrl, Data)
	}
	var PokeCatalogue map[string]internal.Poke

	exp := poke.BaseExperience
	source := rand.NewSource(time.Now().UnixNano())
	//creates seed for random from time
	r := rand.New(source)
	// creates random number from seed
	rand := r.Intn(exp)
	// sets range of random number and outputs to rand

	if rand > exp/2 {
		fmt.Printf("%s was caught\n", selection)
		PokeCatalogue = map[string]internal.Poke{
			selection: {
				BaseExperience: exp,
			},
		}
		fmt.Printf("%v added to pokedex!\n", PokeCatalogue)
		// printing because you are not returning or using PokeCatalogue yet
	} else {
		fmt.Printf("%s escaped\n", selection)
	}

	return nil
}

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
	}
}

func init() {
	cache = pokecache.NewCache(5 * time.Minute)
}
