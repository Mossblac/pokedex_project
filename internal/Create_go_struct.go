package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
