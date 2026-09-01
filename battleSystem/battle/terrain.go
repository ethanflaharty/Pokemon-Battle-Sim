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
	case battledata.Grassy:
		first, _ := determineSpeedOrder(ally, foe, *bs)
		if first == ally {
			if ally.HP < ally.MaxHP {
				ally.HP += ally.MaxHP / 16
				if ally.HP > ally.MaxHP {
					ally.HP = ally.MaxHP
				}
				fmt.Printf("%v restored its HP using the Grassy Terrain! %v now has %v HP left!\n", ally.Name, ally.Name, ally.HP)
			}

			if foe.HP < foe.MaxHP {
				foe.HP += foe.MaxHP / 16
				if foe.HP > foe.MaxHP {
					foe.HP = foe.MaxHP
				}
				fmt.Printf("%v restored its HP using the Grassy Terrain! %v now has %v HP left!\n", foe.Name, foe.Name, foe.HP)
			}
		} else {
			if foe.HP < foe.MaxHP {
				foe.HP += foe.MaxHP / 16
				if foe.HP > foe.MaxHP {
					foe.HP = foe.MaxHP
				}
				fmt.Printf("%v restored its HP using the Grassy Terrain! %v now has %v HP left!\n", foe.Name, foe.Name, foe.HP)
			}

			if ally.HP < ally.MaxHP {
				ally.HP += ally.MaxHP / 16
				if ally.HP > ally.MaxHP {
					ally.HP = ally.MaxHP
				}
				fmt.Printf("%v restored its HP using the Grassy Terrain! %v now has %v HP left!\n", ally.Name, ally.Name, ally.HP)
			}
		}

		bs.Conditions.TerrainTurns--
		if bs.Conditions.TerrainTurns == 0 {
			bs.Conditions.Terrain = battledata.Regular
			fmt.Println("The battlefield returned to normal.")
		}
	case battledata.Misty:
		bs.Conditions.TerrainTurns--
		if bs.Conditions.TerrainTurns == 0 {
			bs.Conditions.Terrain = battledata.Regular
			fmt.Println("The mist disappeared from the battlefield.")
		}
	case battledata.Psych:
		bs.Conditions.TerrainTurns--
		if bs.Conditions.TerrainTurns == 0 {
			bs.Conditions.Terrain = battledata.Regular
			fmt.Println("The Psychic Terrain disappeared!")
		}
	}
}
