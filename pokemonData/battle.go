package pokemondata

import (
	"fmt"
	"math/rand/v2"
)

type DamageResult struct {
	Damage int
	Crit   bool
}

func calculateDamage(attacker, defender Pokemon, move Move) DamageResult {
	damage := ((((((2 * attacker.Level) / 5) + 2) * move.Power * attacker.Attack / defender.Defense) / 50) + 2)

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

func applyDamage(attacker, defender *Pokemon, move Move) {
	damageTaken := calculateDamage(*attacker, *defender, move)
	defender.HP -= damageTaken.Damage

	fmt.Printf("%v used %v!\n", attacker.Name, move.Name)

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

func selectPlayerMove(pokemon *Pokemon) Move {
	fmt.Println("Choose a Move:")

	for i, move := range pokemon.Moves {
		fmt.Printf("%v. %v\n", i+1, move.Name)
	}

	var choice int
	fmt.Scanln(&choice)

	return pokemon.Moves[choice-1]
}

func selectAIMove(pokemon *Pokemon) Move {
	return pokemon.Moves[rand.IntN(len(pokemon.Moves))]
}

type TurnAction struct {
	User *Pokemon
	Move Move
}

func BattleSequence(ally, foe *Pokemon) {
	turn := 0
	for ally.HP > 0 && foe.HP > 0 {
		turn++
		fmt.Println()
		fmt.Printf("Turn %v\n", turn)
		fmt.Println()

		playerAction := TurnAction{
			User: ally,
			Move: selectPlayerMove(ally),
		}
		enemyAction := TurnAction{
			User: foe,
			Move: selectAIMove(foe),
		}

		first, second := determineSpeedOrder(ally, foe)

		var firstAction TurnAction
		var secondAction TurnAction

		if first == ally {
			firstAction = playerAction
			secondAction = enemyAction
		} else {
			firstAction = enemyAction
			secondAction = playerAction
		}

		applyDamage(firstAction.User, secondAction.User, firstAction.Move)

		if second.HP > 0 {
			applyDamage(secondAction.User, firstAction.User, secondAction.Move)
		}
	}
}
