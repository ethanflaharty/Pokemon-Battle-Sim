package pokemondata

import (
	"fmt"
)

type MoveCategory int

const (
	Physical MoveCategory = iota
	Special
	Status
)

type Move struct {
	Name     string
	Power    int
	Type     Type
	Category MoveCategory

	Accuracy    int
	NeverMisses bool

	TotalPP int
	PPLeft  int
}

func (p *Pokemon) AddMove(move Move) error {
	if len(p.Moves) >= 4 {
		return fmt.Errorf("%v already has 4 moves!", p.Name)
	}

	p.Moves = append(p.Moves, move)
	return nil
}
