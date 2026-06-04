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
	for trainerPokemon.HP != 0 && wild.HP != 0 {
		if wild.Speed > trainerPokemon.Speed {
			trainerPokemon.HP = trainerPokemon.HP - pokemondata.Damage(wild, trainerPokemon)
			if trainerPokemon.HP <= 0 {
				trainerPokemon.HP = 0
				fmt.Printf("%v has fainted!\n", trainerPokemon.Name)
			} else if trainerPokemon.HP > 0 {
				fmt.Printf("%v now has %v HP!\n", trainerPokemon.Name, trainerPokemon.HP)
			}
			wild.HP = wild.HP - pokemondata.Damage(trainerPokemon, wild)
			if wild.HP <= 0 {
				wild.HP = 0
				fmt.Printf("%v has fainted!\n", wild.Name)
			} else if wild.HP > 0 {
				fmt.Printf("%v now has %v HP!\n", wild.Name, wild.HP)
			}
		} else {
			wild.HP = wild.HP - pokemondata.Damage(trainerPokemon, wild)
			if wild.HP <= 0 {
				wild.HP = 0
				fmt.Printf("%v has fainted!\n", wild.Name)
			} else if wild.HP > 0 {
				fmt.Printf("%v now has %v HP!\n", wild.Name, wild.HP)
			}
			trainerPokemon.HP = trainerPokemon.HP - pokemondata.Damage(wild, trainerPokemon)
			if trainerPokemon.HP <= 0 {
				trainerPokemon.HP = 0
				fmt.Printf("%v has fainted!\n", trainerPokemon.Name)
			} else if trainerPokemon.HP > 0 {
				fmt.Printf("%v now has %v HP!\n", trainerPokemon.Name, trainerPokemon.HP)
			}
		}
	}
}
