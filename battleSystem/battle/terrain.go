package battle

import (
	"fmt"
	battledata "pokemonBattleSim/battleSystem/battleData"
	pokemondata "pokemonBattleSim/pokemonData"
)

func ProcessTerrain(ally, foe *pokemondata.Pokemon, bs *battledata.BattleState) {
	switch bs.Conditions.Terrain {
	case battledata.Elec:
		bs.Conditions.TerrainTurns--
		if bs.Conditions.TerrainTurns == 0 {
			bs.Conditions.Terrain = battledata.Regular
			fmt.Println("The electric current disappeared from the battlefield!")
		}
	}
}
