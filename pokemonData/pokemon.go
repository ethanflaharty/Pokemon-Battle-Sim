package pokemondata

type Pokemon struct {
	Name      string
	BaseStats PokemonBaseStats
	Level     int
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
