package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// https://pokeapi.co/api/v2/location-area/- area url example
// skarmory, torkoal
//https://pokeapi.co/api/v2/pokemon/skarmory - pokemon info example url

/* you can make your own struct that only has the fields you need. as long as
the field matches the structure of the json it will unmarshal into that field,
but anything not in your struct is ignored - see Explore struct for example of how
to cut it down.- each key on left, needs a 'json:"identifier"' to match
*/

type AreaStruct struct {
	Count    int     `json:"count"`
	Next     string  `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func CreateGoStruct(url string) (AreaStruct, error) {
	var areas AreaStruct
	res, err := http.Get(url)
	if err != nil {
		return AreaStruct{}, fmt.Errorf("error obtaining response from http")
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&areas); err != nil {
		return AreaStruct{}, fmt.Errorf("error decoding resposne")
	}
	return areas, nil

}

type Explore struct {
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

func CreateExploreStruct(url string) (Explore, error) {
	var ex Explore
	res, err := http.Get(url)
	if err != nil {
		return Explore{}, fmt.Errorf("error obtaining response from http")
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&ex); err != nil {
		return Explore{}, fmt.Errorf("error decoding resposne")
	}
	return ex, nil

}

type Poke struct {
	BaseExperience int `json:"base_experience"`
}

func CreatePokeStruct(url string) (Poke, error) {
	var poke Poke
	res, err := http.Get(url)
	if err != nil {
		return Poke{}, fmt.Errorf("error obtaining response from http")
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&poke); err != nil {
		return Poke{}, fmt.Errorf("error decoding resposne")
	}
	return poke, nil
}
