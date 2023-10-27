package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	//"text/tabwriter"

	"pokeDSA/pokemonapi"
)

func parseCMD(str string) {
	str = strings.Trim(str, "\n")
	args := strings.Split(str, " ")
	cmd := args[0]

	if cmd == "exit" {
		fmt.Println("Gotta catch'em all! Goodbye Trainer...")
		os.Exit(0)
	} else {
		if cmd == "check" {
			var pokemon pokemonapi.Pokemon
			arg := args[1]
			pokemon.GetPokemonData(arg)
			pokemon.DisplayPokemonData()
		} else if cmd == "location" { 
			var pokemons pokemonapi.Pokemons
			location := args[1]
			pokemons.GetPokemonsData(location)
			pokemons.DisplayPokemonsNameByLocation()
		} else if cmd == "moves" {
			var pokemon pokemonapi.Pokemon
			arg := args[1]
			pokemon.GetPokemonData(arg)
			pokemon.DisplayPokemonMoves()
		} else {
			fmt.Println("Invalid input")
		}
	}
}
	
func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("err:", err)
		}
		parseCMD(input)
	}
}


