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
			{Name: "Scratch", Power: 40, Type: pokemondata.Normal, Accuracy: 100, NeverMisses: false, TotalPP: 35, PPLeft: 35, Category: pokemondata.Physical},
			{Name: "Ember", Power: 40, Type: pokemondata.Fire, Accuracy: 100, NeverMisses: false, TotalPP: 25, PPLeft: 25, Category: pokemondata.Special},
			{Name: "Flamethrower", Power: 90, Type: pokemondata.Fire, Accuracy: 100, NeverMisses: false, TotalPP: 15, PPLeft: 15, Category: pokemondata.Special},
		},
		Types: []pokemondata.Type{
			pokemondata.Fire,
		},
	}
	wild.HP = pokemondata.HPForumla(wild)
	wild.MaxHP = wild.HP
	wild.CalculateBattleStats()

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
			{Name: "Stun Spore", Power: 0, Type: pokemondata.Grass, Accuracy: 75, NeverMisses: false, TotalPP: 30, PPLeft: 30, Category: pokemondata.Status},
			{Name: "Venoshock", Power: 65, Type: pokemondata.Poison, Accuracy: 100, NeverMisses: false, TotalPP: 10, PPLeft: 10, Category: pokemondata.Special},
			{Name: "Magical Leaf", Power: 60, Type: pokemondata.Grass, Accuracy: 100, NeverMisses: true, TotalPP: 20, PPLeft: 20, Category: pokemondata.Special},
			{Name: "Razor Leaf", Power: 55, Type: pokemondata.Grass, Accuracy: 95, NeverMisses: false, TotalPP: 25, PPLeft: 25, Category: pokemondata.Physical},
		},
		Types: []pokemondata.Type{
			pokemondata.Grass,
			pokemondata.Poison,
		},
	}
	ally.HP = pokemondata.HPForumla(ally)
	ally.MaxHP = ally.HP
	ally.CalculateBattleStats()

	fmt.Printf("A wild %v appeared!\n", wild.Name)
	fmt.Printf("Go! %v!\n", ally.Name)

	battledata.BattleSequence(&ally, &wild)
}
