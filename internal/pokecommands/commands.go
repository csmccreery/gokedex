package pokecommands

import (
	"github.com/csmccreery/gokedex/internal/repl"
)

func Map(arg CommandArg) error {
	if arg != nil {
		arg = nil
	}

	// Check to see if the maps are already cached, if so just retrieve them from file

	// If not cached, call GetLocationMaps store them in memory, print the first 20 and cache the rest.
}

func Mapb(arg CommandArg) error {}

func GetPokeCommands() map[string]repl.Command {
	return map[string]repl.Command{}
}
