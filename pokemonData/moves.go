package pokemondata

import (
	"fmt"
)

func CalculateStatusMove(user, target *Pokemon, move Move) {
	switch move.Name {
	case "Growth":
		Growth(user)
	default:
		fmt.Println("specified status move not implemented yet")
	}
}

func Growth(p *Pokemon) {
	if p.StatStages.Attack == 6 && p.StatStages.SpAttack == 6 {
		fmt.Printf("%v Attack won't go any higher!\n", p.Name)
		fmt.Printf("%v Special Attack won't go any higher!\n", p.Name)
		return
	} else if p.StatStages.Attack == 6 && p.StatStages.SpAttack != 6 {
		fmt.Printf("%v Attack won't go any higher!\n", p.Name)
		p.ChangeStatStage(SpAttack, 1)
		fmt.Printf("%v Special Attack rose!\n", p.Name)
		return
	} else if p.StatStages.Attack != 6 && p.StatStages.SpAttack == 6 {
		p.ChangeStatStage(Attack, 1)
		fmt.Printf("%v Attack rose!\n", p.Name)
		fmt.Printf("%v Special Attack won't go any higher!\n", p.Name)
		return
	} else {
		p.ChangeStatStage(Attack, 1)
		p.ChangeStatStage(SpAttack, 1)
		fmt.Printf("%v Attack rose!\n", p.Name)
		fmt.Printf("%v Special Attack rose!\n", p.Name)
		return
	}
}

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
