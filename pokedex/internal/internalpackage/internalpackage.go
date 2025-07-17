package internalpackage

import (
	"encoding/json"
	"fmt"
	"net/http"
)

/* You're meant to call the data from the website
for each call of your CommandMap and CommandMapb
functions. The PokeAPI endpoint for location
areas is designed to provide batches of results,
and it also provides Next and Previous URLs in
its response. Your config struct should store
these URLs to allow you to paginate through
the data by making new HTTP requests to those URLs.

So, instead of downloading all the data
upfront, you'll be making live GET requests
to the PokeAPI to fetch the next or previous
set of 20 location areas. This is how you'll
explore the Pokemon world dynamically! */

/*You can curl the PokeAPI endpoint,
copy the JSON response, paste it into
the "JSON to Go" tool, and it will generate
the corresponding Go struct definitions for you.
It's quite a time-saver!

However, remember our previous discussion
about json.Unmarshal versus json.NewDecoder
and directly reading from the HTTP response body?
While the "JSON to Go" tool helps you define
the structure, you'll still need to ensure
your Go code correctly fetches the data from
the API and then uses that defined struct
to parse the incoming JSON stream. */

// https://pokeapi.co/api/v2/location-area/

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
