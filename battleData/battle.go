package battledata

import (
	"fmt"
	"math/rand/v2"
	pokemondata "pokemonBattleSim/pokemonData"
)

func determineSpeedOrder(ally, foe *pokemondata.Pokemon) (*pokemondata.Pokemon, *pokemondata.Pokemon) {
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

func selectPlayerMove(pokemon *pokemondata.Pokemon) pokemondata.Move {
	fmt.Println("Choose a Move:")

	for i, move := range pokemon.Moves {
		fmt.Printf("%v. %v\n", i+1, move.Name)
	}

	var choice int
	fmt.Scanln(&choice)

	return pokemon.Moves[choice-1]
}

func selectAIMove(pokemon *pokemondata.Pokemon) pokemondata.Move {
	return pokemon.Moves[rand.IntN(len(pokemon.Moves))]
}

type TurnAction struct {
	User *pokemondata.Pokemon
	Move pokemondata.Move
}

func BattleSequence(ally, foe *pokemondata.Pokemon) {
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
