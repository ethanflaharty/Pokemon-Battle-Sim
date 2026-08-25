package pokemondata

type Pokemon struct {
	Name      string
	BaseStats Stats
	Level     int
	Moves     []Move

	Types       []Type
	StatStages  StatStages
	BattleStats BattleStats
	IVs         Stats
	EVs         Stats
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
