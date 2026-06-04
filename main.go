package main

import (
	"fmt"
	pokemondata "pokemonBattleSim/pokemonData"
)

func main() {
	wild := pokemondata.Pokemon{
		Name: "Charmander",
		BaseStats: pokemondata.PokemonBaseStats{
			HP:        39,
			Attack:    52,
			Defense:   43,
			SpAttack:  60,
			SpDefense: 50,
			Speed:     65,
		},
		Level: 50,
	}
	wild.HP = pokemondata.HPForumla(wild)
	wild.CalculateStats()
	trainerPokemon := pokemondata.Pokemon{
		Name: "Bulbasaur",
		BaseStats: pokemondata.PokemonBaseStats{
			HP:        45,
			Attack:    49,
			Defense:   49,
			SpAttack:  65,
			SpDefense: 65,
			Speed:     45,
		},
		Level: 50,
	}
	trainerPokemon.HP = pokemondata.HPForumla(trainerPokemon)
	trainerPokemon.CalculateStats()
	fmt.Printf("A wild %v appeared!\n", wild.Name)
	fmt.Printf("Go! %v!\n", trainerPokemon.Name)

}
