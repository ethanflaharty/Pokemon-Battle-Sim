package pokemondata

import (
	"encoding/json"
	"fmt"
	"os"
)

func GetPokemon(name string) (Pokemon, error) {
	info, exists := PokemonData[name]
	if !exists {
		return Pokemon{}, fmt.Errorf("pokemon does not exist")
	}

	pokemon := Pokemon{}

	type1, err := typeConversion(info.Type1)
	if err != nil {
		return Pokemon{}, err
	}

	if info.Type2 != "" {
		type2, err := typeConversion(info.Type2)
		if err != nil {
			return Pokemon{}, err
		}

		pokemon = Pokemon{
			Name:      name,
			BaseStats: info.BaseStats,
			Types:     []Type{type1, type2},
		}
	} else {
		pokemon = Pokemon{
			Name:      name,
			BaseStats: info.BaseStats,
			Types:     []Type{type1},
		}
	}

	return pokemon, nil
}

func LoadPokemon() error {
	data, err := os.ReadFile("persistentData/pokemon.json")
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &PokemonData)
	if err != nil {
		return err
	}

	return nil
}

var PokemonData map[string]pokemonJSON

type pokemonJSON struct {
	Type1     string   `json:"type1"`
	Type2     string   `json:"type2"`
	BaseStats Stats    `json:"baseStats"`
	Abilities []string `json:"abilities"`
	Learnset  []string `json:"learnset"`
}

type Pokemon struct {
	Name      string
	BaseStats Stats
	Level     int
	Moves     []Move

	Types       []Type
	IVs         Stats
	EVs         Stats
	Nature      Nature
	StatStages  StatStages
	BattleStats BattleStats

	Status      StatusCondition
	StatusTurns int

	MaxHP     int
	HP        int
	Attack    int
	Defense   int
	SpAttack  int
	SpDefense int
	Speed     int
}

type Nature int

const (
	Lonely Nature = iota
	Adamant
	Naughty
	Brave
	Bold
	Impish
	Lax
	Relaxed
	Modest
	Mild
	Rash
	Quiet
	Calm
	Gentle
	Careful
	Sassy
	Timid
	Hasty
	Jolly
	Naive
	Serious
)
