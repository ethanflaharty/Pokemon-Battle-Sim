package pokemondata

import (
	"fmt"
	"math/rand/v2"
)

type DamageResult struct {
	Damage int
	Crit   bool
}

func calculateDamage(attacker, defender Pokemon) DamageResult {
	damage := ((((((2 * attacker.Level) / 5) + 2) * 50 * attacker.Attack / defender.Defense) / 50) + 2)

	multiplier := float64(rand.IntN(16)+85) / 100.0

	result := DamageResult{}

	if rand.Float64() < 0.0625 {
		result.Crit = true
		result.Damage = int(float64(damage) * multiplier * 1.5)
	} else {
		result.Damage = int(float64(damage) * multiplier)
	}
	return result
}

func applyDamage(attacker, defender *Pokemon) {
	damageTaken := calculateDamage(*attacker, *defender)
	defender.HP -= damageTaken.Damage
	if damageTaken.Crit {
		fmt.Print("A critical hit! ")
	}
	if defender.HP <= 0 {
		defender.HP = 0
		fmt.Printf("%v took %v damage! ", defender.Name, damageTaken.Damage)
		fmt.Printf("%v has fainted!\n", defender.Name)

	} else {
		fmt.Printf("%v took %v damage! ", defender.Name, damageTaken.Damage)
		fmt.Printf("%v now has %v HP!\n", defender.Name, defender.HP)
	}
}

func determineSpeedOrder(ally, foe *Pokemon) (*Pokemon, *Pokemon) {
	if ally.Speed > foe.Speed {
		return ally, foe
	}

	if foe.Speed > ally.Speed {
		return foe, ally
	}

	if rand.IntN(2) == 0 {
		return ally, foe
	}

	return foe, ally
}

func BattleSequence(ally, foe *Pokemon) {
	turn := 0
	for ally.HP > 0 && foe.HP > 0 {
		turn++
		fmt.Println()
		fmt.Printf("Turn %v\n", turn)
		fmt.Println()

		first, second := determineSpeedOrder(ally, foe)

		applyDamage(first, second)

		if second.HP > 0 {
			applyDamage(second, first)
		}
	}
}
