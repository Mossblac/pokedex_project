package assist

import (
	"fmt"
)

func CommandInspect(cfg *Config, selection string) error {
	key, ok := PokeCatalogue[selection]
	if !ok {
		fmt.Printf("%v has not been captured\n", selection)
	}
	if ok {
		fmt.Printf("Height: %v\nWeight: %v\n", key.Height, key.Weight)
		fmt.Printf("Stats:\n")
		for i := range key.Stats {
			val := key.Stats[i].BaseStat
			name := key.Stats[i].Stat.Name
			fmt.Printf("  -%v: %v\n", name, val)
		}
		fmt.Printf("Types:\n")
		for i := range key.Types {
			fmt.Printf("  -%v\n", key.Types[i].Type.Name)
		}
	}
	return nil
}
