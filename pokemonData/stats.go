package pokemondata

import (
	"log"
)

func HPForumla(p Pokemon) int {
	dividend := (2 * p.BaseStats.HP) * p.Level
	division := dividend / 100
	final := division + p.Level + 10
	return final
}

func calculateStat(baseStat, level int) int {
	return ((2 * baseStat * level) / 100) + 5
}

func (p *Pokemon) CalculateBattleStats() {
	p.Attack = calculateStat(p.BaseStats.Attack, p.Level)
	p.Defense = calculateStat(p.BaseStats.Defense, p.Level)
	p.SpAttack = calculateStat(p.BaseStats.SpAttack, p.Level)
	p.SpDefense = calculateStat(p.BaseStats.SpDefense, p.Level)
	p.Speed = calculateStat(p.BaseStats.Speed, p.Level)
}

func (p *Pokemon) UpdateBattleStats() {
	p.BattleStats.Attack = p.CalculateStatChange(
		p.Attack,
		p.StatStages.Attack,
	)

	p.BattleStats.Defense = p.CalculateStatChange(
		p.Defense,
		p.StatStages.Defense,
	)

	p.BattleStats.SpAttack = p.CalculateStatChange(
		p.SpAttack,
		p.StatStages.SpAttack,
	)

	p.BattleStats.SpDefense = p.CalculateStatChange(
		p.SpDefense,
		p.StatStages.SpDefense,
	)

	p.BattleStats.Speed = p.CalculateStatChange(
		p.Speed,
		p.StatStages.Speed,
	)
}

func (p *Pokemon) ChangeStatStage(stat Stat, change int) {
	switch stat {
	case Attack:
		p.StatStages.Attack += change
	case Defense:
		p.StatStages.Defense += change
	case SpAttack:
		p.StatStages.SpAttack += change
	case SpDefense:
		p.StatStages.SpDefense += change
	case Speed:
		p.StatStages.Speed += change
	case Accuracy:
		p.StatStages.Accuracy += change
	case Evasion:
		p.StatStages.Evasion += change
	}
}

func (p *Pokemon) CalculateStatChange(stat, stage int) int {
	switch stage {
	case -6:
		return stat * 2 / 8
	case -5:
		return stat * 2 / 7
	case -4:
		return stat * 2 / 6
	case -3:
		return stat * 2 / 5
	case -2:
		return stat * 2 / 4
	case -1:
		return stat * 2 / 3
	case 0:
		return stat
	case 1:
		return stat * 3 / 2
	case 2:
		return stat * 4 / 2
	case 3:
		return stat * 5 / 2
	case 4:
		return stat * 6 / 2
	case 5:
		return stat * 7 / 2
	case 6:
		return stat * 8 / 2
	default:
		log.Fatalln("stat changes only go from -6 to 6")
		return stat
	}
}

type PokemonBaseStats struct {
	HP        int
	Attack    int
	Defense   int
	SpAttack  int
	SpDefense int
	Speed     int
}

type BattleStats struct {
	Attack    int
	Defense   int
	SpAttack  int
	SpDefense int
	Speed     int
}

type StatStages struct {
	Attack    int
	Defense   int
	SpAttack  int
	SpDefense int
	Speed     int
	// Implement Calculation For Accuracy and Evasion later on
	Accuracy int
	Evasion  int
}

type Stat int

const (
	Attack Stat = iota
	Defense
	SpAttack
	SpDefense
	Speed
	Accuracy
	Evasion
)
