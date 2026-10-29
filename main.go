package main

import (
	"fmt"
	"log"
	battle "pokemonBattleSim/battleSystem/battle"
	pokemondata "pokemonBattleSim/pokemonData"
)

func main() {
	err := pokemondata.LoadMoves()
	if err != nil {
		log.Fatal(err)
	}

	err = pokemondata.LoadPokemon()
	if err != nil {
		log.Fatal(err)
	}

	wild, err := pokemondata.GetPokemon("Charmander")
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	ally, err := pokemondata.GetPokemon("Venusaur")
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	scratch, err := pokemondata.GetMove("Scratch")
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	ember, err := pokemondata.GetMove("Ember")
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	pound, err := pokemondata.GetMove("Pound")
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	wild.Level = 50
	wild.Moves = []pokemondata.Move{
		scratch,
		ember,
	}

	ally.Level = 50
	ally.Moves = []pokemondata.Move{
		scratch,
		ember,
		pound,
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
