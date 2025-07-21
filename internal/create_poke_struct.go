package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
