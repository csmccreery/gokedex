package main

import (
	"github.com/csmccreery/gokedex/internal/repl"
	"github.com/csmccreery/gokedex/internal/pokecommands"
)

func main() {
    pokeRepl := repl.Repl{}
    pokeRepl.SetCommands(pokecommands.GetPokeCommands())
    pokeRepl.Start("Pokedex > ")
}
