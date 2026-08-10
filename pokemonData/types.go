package pokemondata

import (
	"log"
	"slices"
)

func EffectivenessCalc(defTypes []Type, atkType Type) float64 {
	if len(defTypes) == 0 || len(defTypes) > 2 {
		log.Fatalln("defending pokemon needs one or two types")
		return 0
	}

	effectiveness := 1.0

	if len(defTypes) == 2 {
		effectiveness = singleTypeEffectiveness(defTypes[0], atkType) * singleTypeEffectiveness(defTypes[1], atkType)
	} else {
		effectiveness = singleTypeEffectiveness(defTypes[0], atkType)
	}

	return effectiveness
}

func singleTypeEffectiveness(defType, atkType Type) float64 {
	defenderEffectiveness := typeChart[defType]

	if slices.Contains(defenderEffectiveness[Weaknesses], atkType) {
		return 2
	}

	if slices.Contains(defenderEffectiveness[Resistances], atkType) {
		return 0.5
	}

	if slices.Contains(defenderEffectiveness[Immunities], atkType) {
		return 0
	}

	return 1
}

type Effectiveness int

const (
	Weaknesses Effectiveness = iota
	Resistances
	Immunities
)

type Type int

const (
	Normal Type = iota
	Fire
	Water
	Electric
	Grass
	Ice
	Fighting
	Poison
	Ground
	Flying
	Psychic
	Bug
	Rock
	Ghost
	Dragon
	Dark
	Steel
	Fairy
)

var typeChart = map[Type]map[Effectiveness][]Type{
	Normal: {
		Weaknesses: {
			Fighting,
		},
		Resistances: {},
		Immunities: {
			Ghost,
		},
	},
	Fire: {
		Weaknesses: {
			Water,
			Ground,
			Rock,
		},
		Resistances: {
			Fire,
			Grass,
			Ice,
			Bug,
			Steel,
			Fairy,
		},
		Immunities: {},
	},
	Water: {
		Weaknesses: {
			Electric,
			Grass,
		},
		Resistances: {
			Fire,
			Water,
			Ice,
			Steel,
		},
		Immunities: {},
	},
	Electric: {
		Weaknesses: {
			Ground,
		},
		Resistances: {
			Electric,
			Flying,
			Steel,
		},
		Immunities: {},
	},
	Grass: {
		Weaknesses: {
			Fire,
			Ice,
			Poison,
			Flying,
			Bug,
		},
		Resistances: {
			Water,
			Electric,
			Grass,
			Ground,
		},
		Immunities: {},
	},
	Ice: {
		Weaknesses: {
			Fire,
			Fighting,
			Rock,
			Steel,
		},
		Resistances: {
			Ice,
		},
		Immunities: {},
	},
	Fighting: {
		Weaknesses: {
			Flying,
			Psychic,
			Fairy,
		},
		Resistances: {
			Bug,
			Rock,
			Dark,
		},
		Immunities: {},
	},
	Poison: {
		Weaknesses: {
			Ground,
			Psychic,
		},
		Resistances: {
			Grass,
			Fighting,
			Poison,
			Bug,
			Fairy,
		},
		Immunities: {},
	},
	Ground: {
		Weaknesses: {
			Water,
			Grass,
			Ice,
		},
		Resistances: {
			Poison,
			Rock,
		},
		Immunities: {
			Electric,
		},
	},
	Flying: {
		Weaknesses: {
			Electric,
			Ice,
			Rock,
		},
		Resistances: {
			Grass,
			Fighting,
			Bug,
		},
		Immunities: {
			Ground,
		},
	},
	Psychic: {
		Weaknesses: {
			Bug,
			Ghost,
			Dark,
		},
		Resistances: {
			Fighting,
			Psychic,
		},
		Immunities: {},
	},
	Bug: {
		Weaknesses: {
			Fire,
			Flying,
			Rock,
		},
		Resistances: {
			Grass,
			Fighting,
			Ground,
		},
		Immunities: {},
	},
	Rock: {
		Weaknesses: {
			Water,
			Grass,
			Fighting,
			Ground,
			Steel,
		},
		Resistances: {
			Normal,
			Fire,
			Poison,
			Flying,
		},
		Immunities: {},
	},
	Ghost: {
		Weaknesses: {
			Ghost,
			Dark,
		},
		Resistances: {
			Poison,
			Bug,
		},
		Immunities: {
			Normal,
			Fighting,
		},
	},
	Dragon: {
		Weaknesses: {
			Ice,
			Dragon,
			Fairy,
		},
		Resistances: {
			Fire,
			Water,
			Electric,
			Grass,
		},
		Immunities: {},
	},
	Dark: {
		Weaknesses: {
			Fighting,
			Bug,
			Fairy,
		},
		Resistances: {
			Ghost,
			Dark,
		},
		Immunities: {
			Psychic,
		},
	},
	Steel: {
		Weaknesses: {
			Fire,
			Fighting,
			Ground,
		},
		Resistances: {
			Normal,
			Grass,
			Ice,
			Flying,
			Psychic,
			Bug,
			Rock,
			Dragon,
			Steel,
			Fairy,
		},
		Immunities: {
			Poison,
		},
	},
	Fairy: {
		Weaknesses: {
			Poison,
			Steel,
		},
		Resistances: {
			Fighting,
			Bug,
			Dark,
		},
		Immunities: {
			Dragon,
		},
	},
}
