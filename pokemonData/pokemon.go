package pokemondata

type Pokemon struct {
	Name      string
	BaseStats PokemonBaseStats
	Level     int
	HP        int
	Attack    int
	Defense   int
}

type PokemonBaseStats struct {
	HP      int
	Attack  int
	Defense int
}

type Type struct {
}
