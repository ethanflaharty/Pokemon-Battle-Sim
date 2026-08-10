package pokemondata

type Pokemon struct {
	Name      string
	BaseStats PokemonBaseStats
	Level     int
	Moves     []Move

	Types []Type

	HP        int
	Attack    int
	Defense   int
	SpAttack  int
	SpDefense int
	Speed     int
}

type PokemonBaseStats struct {
	HP        int
	Attack    int
	Defense   int
	SpAttack  int
	SpDefense int
	Speed     int
}
