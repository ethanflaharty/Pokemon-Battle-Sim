package pokemondata

import (
	"fmt"
)

func CalculateStatusMove(user, target *Pokemon, move Move) {
	switch move.Name {
	case "Growth":
		Growth(user)
	case "Growl":
		Growl(target)
	case "Iron Defense":
		IronDefense(user)
	case "Tail Whip":
		TailWhip(target)
	case "Amnesia":
		Amnesia(user)
	case "Agility":
		Agility(user)
	case "Scary Face":
		ScaryFace(target)
	case "Hone Claws":
		HoneClaws(user)
	case "Sand Attack":
		SandAttack(target)
	case "Double Team":
		DoubleTeam(user)
	case "Will-O-Wisp":
		WillOWisp(target)
	case "Poison Powder":
		PoisonPowder(target)
	case "Toxic":
		Toxic(target)
	case "Stun Spore":
		StunSpore(target)
	default:
		fmt.Println("specified status move not implemented yet")
	}
}

func StunSpore(p *Pokemon) {
	if len(p.Types) == 2 {
		if p.Types[0] == Grass || p.Types[1] == Grass {
			fmt.Printf("It doesn't affect %v\n", p.Name)
		}
	}
	if p.Types[0] == Grass {
		fmt.Printf("It doesn't affect %v\n", p.Name)
	}

	result := p.ApplyStatus(Paralysis)
	if result == true {
		fmt.Printf("%v is paralyzed, so it may be unable to move!\n", p.Name)
	}
}

func Toxic(p *Pokemon) {
	result := p.ApplyStatus(PoisonBad)
	if result == true {
		fmt.Printf("%v was badly poisoned!\n", p.Name)
	}
}

func PoisonPowder(p *Pokemon) {
	result := p.ApplyStatus(PoisonReg)
	if result == true {
		fmt.Printf("%v was poisoned!\n", p.Name)
	}
}

func WillOWisp(p *Pokemon) {
	result := p.ApplyStatus(Burn)
	if result == true {
		fmt.Printf("%v was burned!\n", p.Name)
	}
}

func DoubleTeam(p *Pokemon) {
	if p.StatStages.Evasion == 6 {
		fmt.Printf("%v Evasion won't go any higher!\n", p.Name)
		return
	} else {
		p.ChangeStatStage(Evasion, 1)
		fmt.Printf("%v Evasion rose!\n", p.Name)
		return
	}
}

func SandAttack(p *Pokemon) {
	if p.StatStages.Accuracy == -6 {
		fmt.Printf("%v Accuracy won't go any lower!\n", p.Name)
		return
	} else {
		p.ChangeStatStage(Accuracy, -1)
		fmt.Printf("%v Accuracy fell!\n", p.Name)
		return
	}
}

func HoneClaws(p *Pokemon) {
	if p.StatStages.Attack == 6 && p.StatStages.Accuracy == 6 {
		fmt.Printf("%v Attack won't go any higher!\n", p.Name)
		fmt.Printf("%v Accuracy won't go any higher!\n", p.Name)
		return
	} else if p.StatStages.Attack == 6 && p.StatStages.Accuracy != 6 {
		fmt.Printf("%v Attack won't go any higher!\n", p.Name)
		p.ChangeStatStage(Accuracy, 1)
		fmt.Printf("%v Accuracy rose!\n", p.Name)
		return
	} else if p.StatStages.Attack != 6 && p.StatStages.Accuracy == 6 {
		p.ChangeStatStage(Attack, 1)
		fmt.Printf("%v Attack rose!\n", p.Name)
		fmt.Printf("%v Accuracy won't go any higher!\n", p.Name)
		return
	} else {
		p.ChangeStatStage(Attack, 1)
		p.ChangeStatStage(Accuracy, 1)
		fmt.Printf("%v Attack rose!\n", p.Name)
		fmt.Printf("%v Accuracy rose!\n", p.Name)
		return
	}
}

func ScaryFace(p *Pokemon) {
	switch p.StatStages.Speed {
	case -6:
		fmt.Printf("%v Speed won't go any lower!\n", p.Name)
		return
	case -5:
		p.ChangeStatStage(Speed, -1)
		fmt.Printf("%v Speed harshly fell!\n", p.Name)
		return
	default:
		p.ChangeStatStage(Speed, -2)
		fmt.Printf("%v Speed harshly fell!\n", p.Name)
		return
	}
}

func Agility(p *Pokemon) {
	switch p.StatStages.Speed {
	case 6:
		fmt.Printf("%v Speed won't go any higher!\n", p.Name)
		return
	case 5:
		p.ChangeStatStage(Speed, 1)
		fmt.Printf("%v Speed rose sharply!\n", p.Name)
		return
	default:
		p.ChangeStatStage(Speed, 2)
		fmt.Printf("%v Speed rose sharply!\n", p.Name)
		return
	}
}

func Amnesia(p *Pokemon) {
	switch p.StatStages.SpDefense {
	case 6:
		fmt.Printf("%v Special Defense won't go any higher!\n", p.Name)
		return
	case 5:
		p.ChangeStatStage(SpDefense, 1)
		fmt.Printf("%v Special Defense rose sharply!\n", p.Name)
		return
	default:
		p.ChangeStatStage(SpDefense, 2)
		fmt.Printf("%v Special Defense rose sharply!\n", p.Name)
		return
	}
}

func TailWhip(p *Pokemon) {
	if p.StatStages.Defense == -6 {
		fmt.Printf("%v Defense won't go any lower!\n", p.Name)
		return
	} else {
		p.ChangeStatStage(Defense, -1)
		fmt.Printf("%v Defense fell!\n", p.Name)
		return
	}
}

func IronDefense(p *Pokemon) {
	switch p.StatStages.Defense {
	case 6:
		fmt.Printf("%v Defense won't go any higher!\n", p.Name)
		return
	case 5:
		p.ChangeStatStage(Defense, 1)
		fmt.Printf("%v Defense rose sharply!\n", p.Name)
		return
	default:
		p.ChangeStatStage(Defense, 2)
		fmt.Printf("%v Defense rose sharply!\n", p.Name)
		return
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

func Growl(p *Pokemon) {
	if p.StatStages.Attack == -6 {
		fmt.Printf("%v Attack won't go any lower!\n", p.Name)
		return
	} else {
		p.ChangeStatStage(Attack, -1)
		fmt.Printf("%v Attack fell!\n", p.Name)
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
