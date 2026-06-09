package pokemondata

import "fmt"

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

func (p *Pokemon) AddMove(move Move) error {
	if len(p.Moves) >= 4 {
		return fmt.Errorf("%v already has 4 moves!", p.Name)
	}

	p.Moves = append(p.Moves, move)
	return nil
}

type PokemonBaseStats struct {
	HP        int
	Attack    int
	Defense   int
	SpAttack  int
	SpDefense int
	Speed     int
}

type Move struct {
	Name  string
	Power int
}

type Type struct {
}
