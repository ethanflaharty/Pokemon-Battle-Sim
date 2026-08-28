package battledata

type BattleState struct {
	Conditions BattleConditions
}

type BattleConditions struct {
	Weather      Weather
	WeatherTurns int
	Terrain      Terrain
	TerrainTurns int
}

type Weather int

const (
	Clear Weather = iota
	HarshSunlight
	Rain
	Sandstorm
	Snow
	ExtremeHarshSunlight
	HeavyRain
	StrongWinds
)

type Terrain int

const (
	Regular Terrain = iota
	Elec
)
