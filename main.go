package main

import (
	"fmt"
	"log"
	battle "pokemonBattleSim/battleSystem/battle"
	pokemondata "pokemonBattleSim/pokemonData"
)

func main() {
	wild := pokemondata.Pokemon{
		Name: "Charmander",
		BaseStats: pokemondata.Stats{
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
			{Name: "Ember", Power: 40, Type: pokemondata.Fire, Accuracy: 100, NeverMisses: false, TotalPP: 10, PPLeft: 10, Category: pokemondata.Special},
		},
		Types: []pokemondata.Type{
			pokemondata.Fire,
		},
	}
	if !wild.ValidateEVs() {
		log.Fatalf("%v has invalid EVs\n", wild.Name)
	}
	if !wild.ValidateIVs() {
		log.Fatalf("%v has invalid IVs\n", wild.Name)
	}
	wild.HP = pokemondata.HPForumla(wild)
	wild.MaxHP = wild.HP
	wild.CalculateBattleStats()

	ally := pokemondata.Pokemon{
		Name: "Bulbasaur",
		BaseStats: pokemondata.Stats{
			HP:        45,
			Attack:    49,
			Defense:   49,
			SpAttack:  65,
			SpDefense: 65,
			Speed:     45,
		},
		Level: 50,
		IVs: pokemondata.Stats{
			Attack:   0,
			SpAttack: 31,
			Speed:    31,
		},
		EVs: pokemondata.Stats{
			HP:       4,
			SpAttack: 252,
			Speed:    252,
		},
		Nature: pokemondata.Modest,
		Moves: []pokemondata.Move{
			{Name: "Weather Ball", Power: 50, Type: pokemondata.Normal, Accuracy: 100, NeverMisses: false, TotalPP: 15, PPLeft: 15, Category: pokemondata.Special},
			{Name: "Dragon Claw", Power: 80, Type: pokemondata.Dragon, Accuracy: 100, NeverMisses: false, TotalPP: 10, PPLeft: 10, Category: pokemondata.Physical},
			{Name: "Magical Leaf", Power: 60, Type: pokemondata.Grass, Accuracy: 100, NeverMisses: true, TotalPP: 20, PPLeft: 20, Category: pokemondata.Special},
			{Name: "Misty Terrain", Power: 0, Type: pokemondata.Fairy, Accuracy: 100, NeverMisses: true, TotalPP: 25, PPLeft: 25, Category: pokemondata.Status},
		},
		Types: []pokemondata.Type{
			pokemondata.Grass,
			pokemondata.Poison,
		},
	}
	if !ally.ValidateEVs() {
		log.Fatalf("%v has invalid EVs\n", ally.Name)
	}
	if !ally.ValidateIVs() {
		log.Fatalf("%v has invalid IVs\n", ally.Name)
	}
	ally.HP = pokemondata.HPForumla(ally)
	ally.MaxHP = ally.HP
	ally.CalculateBattleStats()

	fmt.Printf("A wild %v appeared!\n", wild.Name)
	fmt.Printf("Go! %v!\n", ally.Name)

	battle.BattleSequence(&ally, &wild)
}
