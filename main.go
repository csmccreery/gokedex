package main

import (
	"github.com/csmccreery/gokedex/internal/pokecommands"
	"github.com/csmccreery/gokedex/internal/repl"
)

func main() {
	pokeRepl := repl.Repl{
		Description: "A CLI pokedex!",
		StopMessage: "Closing the Pokedex... Goodbye!",
		HelpMessage: "Welcome to the Pokedex!",
		History:     map[string]string{},
		Commands:    pokecommands.GetPokeCommands(),
	}

	pokeRepl.Start("Pokedex > ")
}
