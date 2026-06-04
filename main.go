package main

import (
	"fmt"
	pokemondata "pokemonBattleSim/pokemonData"
)

func main() {
	wild := pokemondata.Pokemon{
		Name: "Charmander",
		BaseStats: pokemondata.PokemonBaseStats{
			HP:      39,
			Attack:  52,
			Defense: 43,
		},
		Level: 5,
	}
	wild.HP = pokemondata.HPForumla(wild)
	trainerPokemon := pokemondata.Pokemon{
		Name: "Bulbasaur",
		BaseStats: pokemondata.PokemonBaseStats{
			HP:      45,
			Attack:  49,
			Defense: 49,
		},
		Level: 5,
	}
	trainerPokemon.HP = pokemondata.HPForumla(trainerPokemon)
	fmt.Printf("A wild %v appeared!\n", wild.Name)
	fmt.Printf("Go! %v!\n", trainerPokemon.Name)

}
