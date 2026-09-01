package battle

import (
	"fmt"
	"math/rand/v2"
	battledata "pokemonBattleSim/battleSystem/battleData"
	pokemondata "pokemonBattleSim/pokemonData"
)

func determineSpeedOrder(ally, foe *pokemondata.Pokemon, bs battledata.BattleState) (*pokemondata.Pokemon, *pokemondata.Pokemon) {
	// BattleState will be used later from what I think will make sense for speed control but this may change
	if ally.BattleStats.Speed > foe.BattleStats.Speed {
		return ally, foe
	}

	if foe.BattleStats.Speed > ally.BattleStats.Speed {
		return foe, ally
	}

	if rand.IntN(2) == 0 {
		return ally, foe
	}

	return foe, ally
}

func selectPlayerMove(pokemon *pokemondata.Pokemon) (*pokemondata.Move, int) {
	for {
		fmt.Println("Choose a Move:")

		for i, move := range pokemon.Moves {
			fmt.Printf("%v. %v  %v/%v\n", i+1, move.Name, move.PPLeft, move.TotalPP)
		}

		var choice int
		fmt.Scanln(&choice)

		move := pokemon.Moves[choice-1]

		if move.PPLeft == 0 {
			fmt.Printf("There's no PP left for this move!\n")
			continue
		}
		// TODO: Force Struggle if there is no PP remaining at all
		pokemon.Moves[choice-1].PPLeft--

		return &pokemon.Moves[choice-1], choice - 1
	}
}

func selectAIMove(pokemon *pokemondata.Pokemon) *pokemondata.Move {
	return &pokemon.Moves[rand.IntN(len(pokemon.Moves))]
}

type TurnAction struct {
	User      *pokemondata.Pokemon
	Move      *pokemondata.Move
	MoveIndex int
}

func CheckPriority(first, second *pokemondata.Pokemon, firstMove, secondMove pokemondata.Move) (*pokemondata.Pokemon, *pokemondata.Pokemon) {
	if firstMove.Priority > 5 {
		firstMove.Priority = 5
	} else if firstMove.Priority < -7 {
		firstMove.Priority = -7
	}

	if secondMove.Priority > 5 {
		secondMove.Priority = 5
	} else if secondMove.Priority < -7 {
		secondMove.Priority = -7
	}

	if firstMove.Priority < secondMove.Priority {
		return second, first
	}

	return first, second
}

func BattleSequence(ally, foe *pokemondata.Pokemon) {
	ally.StatStages = pokemondata.StatStages{}
	foe.StatStages = pokemondata.StatStages{}
	bs := battledata.BattleState{}
	turn := 0

	for ally.HP > 0 && foe.HP > 0 {
		turn++
		fmt.Println()
		fmt.Printf("Turn %v\n", turn)
		fmt.Println()

		ally.UpdateBattleStats(bs)
		foe.UpdateBattleStats(bs)

		playerMove, playerMoveIndex := selectPlayerMove(ally)

		playerAction := TurnAction{
			User:      ally,
			Move:      playerMove,
			MoveIndex: playerMoveIndex,
		}
		enemyAction := TurnAction{
			User: foe,
			Move: selectAIMove(foe),
		}

		first, second := determineSpeedOrder(ally, foe, bs)

		var firstAction TurnAction
		var secondAction TurnAction

		if first == ally {
			firstAction = playerAction
			secondAction = enemyAction
		} else {
			firstAction = enemyAction
			secondAction = playerAction
		}

		first, second = CheckPriority(first, second, *firstAction.Move, *secondAction.Move)
		if first == ally {
			firstAction = playerAction
			secondAction = enemyAction
		} else {
			firstAction = enemyAction
			secondAction = playerAction
		}

		if pokemondata.CanMove(firstAction.User) {
			applyDamage(firstAction.User, secondAction.User, *firstAction.Move, &bs)
			ally.UpdateBattleStats(bs)
			foe.UpdateBattleStats(bs)
		}

		if second.HP > 0 {
			if pokemondata.CanMove(secondAction.User) {
				applyDamage(secondAction.User, firstAction.User, *secondAction.Move, &bs)
				ally.UpdateBattleStats(bs)
				foe.UpdateBattleStats(bs)
			}
		}

		if ally.HP > 0 && foe.HP > 0 {
			if ally.Status != pokemondata.None {
				pokemondata.ProcessStatus(ally)
			}
			if foe.Status != pokemondata.None {
				pokemondata.ProcessStatus(foe)
			}
		}

		if ally.HP > 0 && foe.HP > 0 && bs.Conditions.Weather != battledata.Clear {
			ProcessWeather(ally, foe, &bs)
		}

		if ally.HP > 0 && foe.HP > 0 && bs.Conditions.Terrain != battledata.Regular {
			ProcessTerrain(ally, foe, &bs)
		}
	}
}
