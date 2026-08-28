package battledata

import "fmt"

func ProcessWeather(bs *BattleState) {
	switch bs.Conditions.Weather {
	case HarshSunlight:
		bs.Conditions.WeatherTurns--
		if bs.Conditions.WeatherTurns == 0 {
			fmt.Println("The sunlight faded.")
		}
		fmt.Println("The sunlight is strong.")
	}
}
