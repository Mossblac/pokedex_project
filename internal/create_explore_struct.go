package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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
