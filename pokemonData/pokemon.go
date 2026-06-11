package pokemondata

type Pokemon struct {
	Name      string
	BaseStats PokemonBaseStats
	Level     int
	Moves     []Move

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

type Type struct {
}
