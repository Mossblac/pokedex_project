package assist

import (
	"encoding/json"
	"fmt"

	"github.com/Mossblac/pokedexcli/internal"
)

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
