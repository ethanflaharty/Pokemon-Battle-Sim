package battle

import (
	"fmt"
	battledata "pokemonBattleSim/battleSystem/battleData"
	pokemondata "pokemonBattleSim/pokemonData"
)

func ProcessWeather(ally, foe *pokemondata.Pokemon, bs *battledata.BattleState) {
	switch bs.Conditions.Weather {
	case battledata.HarshSunlight:
		bs.Conditions.WeatherTurns--
		if bs.Conditions.WeatherTurns == 0 {
			bs.Conditions.Weather = battledata.Clear
			fmt.Println("The sunlight faded.")
		}
	case battledata.Rain:
		bs.Conditions.WeatherTurns--
		if bs.Conditions.WeatherTurns == 0 {
			bs.Conditions.Weather = battledata.Clear
			fmt.Println("The rain stopped.")
		}
	case battledata.Sandstorm:
		bs.Conditions.WeatherTurns--
		if bs.Conditions.WeatherTurns == 0 {
			bs.Conditions.Weather = battledata.Clear
			fmt.Println("The sandstorm subsided.")
			return
		}

		first, _ := determineSpeedOrder(ally, foe)
		if first == ally {
			if len(ally.Types) == 2 {
				if ally.Types[0] == pokemondata.Rock || ally.Types[1] == pokemondata.Rock || ally.Types[0] == pokemondata.Steel || ally.Types[1] == pokemondata.Steel || ally.Types[0] == pokemondata.Ground || ally.Types[1] == pokemondata.Ground {

				} else {
					dmg := max(ally.MaxHP/16, 1)
					ally.HP -= dmg
					fmt.Printf("%v is buffeted by the sandstorm! %v took %v damage!\n", ally.Name, ally.Name, dmg)
				}
			} else {
				if ally.Types[0] == pokemondata.Rock || ally.Types[0] == pokemondata.Steel || ally.Types[0] == pokemondata.Ground {

				} else {
					dmg := max(ally.MaxHP/16, 1)
					ally.HP -= dmg
					fmt.Printf("%v is buffeted by the sandstorm! %v took %v damage!\n", ally.Name, ally.Name, dmg)
				}
			}
			if ally.HP <= 0 {
				fmt.Printf("%v has fainted!\n", ally.Name)
			}

			if ally.HP > 0 {
				if len(foe.Types) == 2 {
					if foe.Types[0] == pokemondata.Rock || foe.Types[1] == pokemondata.Rock || foe.Types[0] == pokemondata.Steel || foe.Types[1] == pokemondata.Steel || foe.Types[0] == pokemondata.Ground || foe.Types[1] == pokemondata.Ground {

					} else {
						dmg := max(foe.MaxHP/16, 1)
						foe.HP -= dmg
						fmt.Printf("The wild %v is buffeted by the sandstorm! %v took %v damage!\n", foe.Name, foe.Name, dmg)
					}
				} else {
					if foe.Types[0] == pokemondata.Rock || foe.Types[0] == pokemondata.Steel || foe.Types[0] == pokemondata.Ground {

					} else {
						dmg := max(foe.MaxHP/16, 1)
						foe.HP -= dmg
						fmt.Printf("The wild %v is buffeted by the sandstorm! %v took %v damage!\n", foe.Name, foe.Name, dmg)
					}
				}

				if foe.HP <= 0 {
					fmt.Printf("The wild %v has fainted!\n", foe.Name)
				}
			}

		} else {
			if len(foe.Types) == 2 {
				if foe.Types[0] == pokemondata.Rock || foe.Types[1] == pokemondata.Rock || foe.Types[0] == pokemondata.Steel || foe.Types[1] == pokemondata.Steel || foe.Types[0] == pokemondata.Ground || foe.Types[1] == pokemondata.Ground {

				} else {
					dmg := max(foe.MaxHP/16, 1)
					foe.HP -= dmg
					fmt.Printf("The wild %v is buffeted by the sandstorm! %v took %v damage!\n", foe.Name, foe.Name, dmg)
				}
			} else {
				if foe.Types[0] == pokemondata.Rock || foe.Types[0] == pokemondata.Steel || foe.Types[0] == pokemondata.Ground {

				} else {
					dmg := max(foe.MaxHP/16, 1)
					foe.HP -= dmg
					fmt.Printf("The wild %v is buffeted by the sandstorm! %v took %v damage!\n", foe.Name, foe.Name, dmg)
				}
			}
			if foe.HP <= 0 {
				fmt.Printf("The wild %v has fainted!\n", foe.Name)
			}

			if foe.HP > 0 {
				if len(ally.Types) == 2 {
					if ally.Types[0] == pokemondata.Rock || ally.Types[1] == pokemondata.Rock || ally.Types[0] == pokemondata.Steel || ally.Types[1] == pokemondata.Steel || ally.Types[0] == pokemondata.Ground || ally.Types[1] == pokemondata.Ground {

					} else {
						dmg := max(ally.MaxHP/16, 1)
						ally.HP -= dmg
						fmt.Printf("%v is buffeted by the sandstorm! %v took %v damage!\n", ally.Name, ally.Name, dmg)
					}
				} else {
					if ally.Types[0] == pokemondata.Rock || ally.Types[0] == pokemondata.Steel || ally.Types[0] == pokemondata.Ground {

					} else {
						dmg := max(ally.MaxHP/16, 1)
						ally.HP -= dmg
						fmt.Printf("%v is buffeted by the sandstorm! %v took %v damage!\n", ally.Name, ally.Name, dmg)
					}
				}

				if ally.HP <= 0 {
					fmt.Printf("%v has fainted!\n", ally.Name)
				}
			}
		}
	case battledata.Snow:
		bs.Conditions.WeatherTurns--
		if bs.Conditions.WeatherTurns == 0 {
			bs.Conditions.Weather = battledata.Clear
			fmt.Println("The snow stopped.")
		}
	}
}
