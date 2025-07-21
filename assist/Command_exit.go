package assist

import (
	"fmt"
	"os"
)

func CommandExit(*Config, string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}
