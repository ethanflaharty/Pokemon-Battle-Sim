package pokemondata

import (
	"fmt"
)

type Move struct {
	Name     string
	Power    int
	Accuracy int
}

func (p *Pokemon) AddMove(move Move) error {
	if len(p.Moves) >= 4 {
		return fmt.Errorf("%v already has 4 moves!", p.Name)
	}

	p.Moves = append(p.Moves, move)
	return nil
}
