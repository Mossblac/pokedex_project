package assist

import (
	"fmt"
)

func CommandPokedex(cfg *Config, selection string) error {
	for key := range PokeCatalogue {
		if len(key) == 0 {
			fmt.Printf("no pokemon captured")
		} else {
			fmt.Printf("  -%v\n", key)
		}
	}
	return nil
}
