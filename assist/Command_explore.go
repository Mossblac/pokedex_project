package assist

import (
	"encoding/json"
	"fmt"

	"github.com/Mossblac/pokedexcli/internal"
)

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
