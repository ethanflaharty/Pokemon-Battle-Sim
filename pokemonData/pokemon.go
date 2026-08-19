package pokemondata

type Pokemon struct {
	Name      string
	BaseStats PokemonBaseStats
	Level     int
	Moves     []Move

	Types       []Type
	StatStages  StatStages
	BattleStats BattleStats
	Status      StatusCondition
	StatusTurns int

	MaxHP     int
	HP        int
	Attack    int
	Defense   int
	SpAttack  int
	SpDefense int
	Speed     int
}
