package pokemondata

import (
	"fmt"
)

const (
	Physical = "physical"
	Special  = "special"
	Status   = "status"
)

type Move struct {
	Name     string
	Power    int
	Type     string
	Category string

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
