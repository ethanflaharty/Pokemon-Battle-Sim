package pokemondata

type Pokemon struct {
	Name      string
	BaseStats Stats
	Level     int
	Moves     []Move

	Types       []Type
	IVs         Stats
	EVs         Stats
	Nature      Nature
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

type Nature int

const (
	Lonely Nature = iota
	Adamant
	Naughty
	Brave
	Bold
	Impish
	Lax
	Relaxed
	Modest
	Mild
	Rash
	Quiet
	Calm
	Gentle
	Careful
	Sassy
	Timid
	Hasty
	Jolly
	Naive
	Serious
)
