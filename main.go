package main

import (
	"fmt"
	battledata "pokemonBattleSim/battleData"
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
		Moves: []pokemondata.Move{
			{Name: "Scratch", Power: 40},
			{Name: "Ember", Power: 40},
			{Name: "Flamethrower", Power: 90},
		},
	}
	wild.HP = pokemondata.HPForumla(wild)
	wild.CalculateStats()
	ally := pokemondata.Pokemon{
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
		Moves: []pokemondata.Move{
			{Name: "Tackle", Power: 40},
			{Name: "Vine Whip", Power: 45},
			{Name: "Seed Bomb", Power: 80},
		},
	}
	ally.HP = pokemondata.HPForumla(ally)
	ally.CalculateStats()
	fmt.Printf("A wild %v appeared!\n", wild.Name)
	fmt.Printf("Go! %v!\n", ally.Name)
	battledata.BattleSequence(&ally, &wild)
}
