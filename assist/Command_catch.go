package assist

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	"github.com/Mossblac/pokedexcli/internal"
)

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

	exp := poke.BaseExperience
	source := rand.NewSource(time.Now().UnixNano())
	//creates seed for random from time
	r := rand.New(source)
	// creates random number from seed
	rand := r.Intn(exp)
	// sets range of random number and outputs to rand

	if rand > exp/2 {
		fmt.Printf("%s was caught\n", selection)
		fmt.Printf("learn more with inspect command\n")
		PokeCatalogue[selection] = poke
	} else {
		fmt.Printf("%s escaped\n", selection)
	}

	return nil
}
