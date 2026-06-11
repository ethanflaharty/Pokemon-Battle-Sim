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
}

func calculateMove(attacker, defender pokemondata.Pokemon, move pokemondata.Move) MoveResult {
	result := MoveResult{}
	// Check for a hit
	result.Hit = rand.IntN(100) < move.Accuracy
	if !result.Hit {
		return result
	}
	// Check for a Crit
	if rand.Float64() < 0.0625 {
		result.Crit = true
	}
	// Damage Calculation
	result.Damage = ((((((2 * attacker.Level) / 5) + 2) * move.Power * attacker.Attack / defender.Defense) / 50) + 2)

	multiplier := float64(rand.IntN(16)+85) / 100.0

	if result.Crit {
		result.Damage = int(float64(result.Damage) * multiplier * 1.5)
	} else {
		result.Damage = int(float64(result.Damage) * multiplier)
	}
	return result
}

func applyDamage(attacker, defender *pokemondata.Pokemon, move pokemondata.Move) {
	fmt.Printf("%v used %v!\n", attacker.Name, move.Name)

	result := calculateMove(*attacker, *defender, move)

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
