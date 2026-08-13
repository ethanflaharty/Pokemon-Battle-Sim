package battledata

import (
	"fmt"
	"math/rand/v2"
	pokemondata "pokemonBattleSim/pokemonData"
)

type MoveResult struct {
	Damage int
	Crit   bool
	Hit    bool
	STAB   bool
}

func calculateMove(attacker, defender *pokemondata.Pokemon, move pokemondata.Move) MoveResult {
	result := MoveResult{}

	// Check for a hit
	if move.NeverMisses {
		result.Hit = true
	} else {
		result.Hit = rand.IntN(100) < move.Accuracy
		if !result.Hit {
			return result
		}
	}

	// Check for a Crit
	if rand.Float64() < 0.0625 {
		result.Crit = true
	}

	// Check for STAB
	if len(attacker.Types) == 2 {
		if attacker.Types[0] == move.Type || attacker.Types[1] == move.Type {
			result.STAB = true
		}
	} else {
		if attacker.Types[0] == move.Type {
			result.STAB = true
		}
	}

	// Calculate Type Effectiveness
	effectiveness := pokemondata.EffectivenessCalc(defender.Types, move.Type)

	// Calculate RNG Multiplier
	multiplier := float64(rand.IntN(16)+85) / 100.0

	// Damage Calculation
	switch move.Category {
	case pokemondata.Physical:
		result.Damage = int(float64(((((((2 * attacker.Level) / 5) + 2) * move.Power * attacker.BattleStats.Attack / defender.BattleStats.Defense) / 50) + 2)) * multiplier * effectiveness)
	case pokemondata.Special:
		result.Damage = int(float64(((((((2 * attacker.Level) / 5) + 2) * move.Power * attacker.BattleStats.SpAttack / defender.BattleStats.SpDefense) / 50) + 2)) * multiplier * effectiveness)
	case pokemondata.Status:
		pokemondata.CalculateStatusMove(attacker, defender, move)
	default:
		panic("unknown move category")
	}

	// Apply STAB
	if result.STAB {
		result.Damage = int(float64(result.Damage) * 1.5)
	}

	// Apply Crit
	if result.Crit {
		result.Damage = int(float64(result.Damage) * 1.5)
	}

	return result
}

func applyDamage(attacker, defender *pokemondata.Pokemon, move pokemondata.Move) {
	fmt.Printf("%v used %v!\n", attacker.Name, move.Name)

	result := calculateMove(attacker, defender, move)

	if !result.Hit {
		fmt.Printf("%v's attack missed!\n", attacker.Name)
	}

	if result.Crit {
		fmt.Println("A critical hit!")
	}

	defender.HP -= result.Damage

	if defender.HP <= 0 {
		defender.HP = 0
		fmt.Printf("%v took %v damage! ", defender.Name, result.Damage)
		fmt.Printf("%v has fainted!\n", defender.Name)

	} else {
		fmt.Printf("%v took %v damage! ", defender.Name, result.Damage)
		fmt.Printf("%v has %v HP left!\n", defender.Name, defender.HP)
	}
}
