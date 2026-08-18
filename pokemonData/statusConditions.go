package pokemondata

import "fmt"

func (p *Pokemon) ApplyStatus(status StatusCondition) bool {
	if p.Status != None {
		return false
	}

	switch status {
	case Burn:
		if len(p.Types) == 2 {
			if p.Types[0] == Fire || p.Types[1] == Fire {
				fmt.Printf("It doesn't affect %v", p.Name)
				return false
			}
		}
		if p.Types[0] == Fire {
			fmt.Printf("It doesn't affect %v", p.Name)
			return false
		}
	case PoisonReg:

	}

	p.Status = status
	return true
}

func ProcessStatus(p *Pokemon) {
	switch p.Status {
	case Burn:
		ApplyBurnDamage(p)
	}
}

func ApplyBurnDamage(p *Pokemon) {
	// max returns the larger of the 2 variables
	dmg := max(p.MaxHP/16, 1)

	p.HP -= dmg
	fmt.Printf("%v was hurt by its burn! %v took %v damage!\n", p.Name, p.Name, dmg)

	if p.HP < 0 || p.HP == 0 {
		p.HP = 0
		fmt.Printf("%v fainted!\n", p.Name)
	}
}

type StatusCondition int

const (
	None StatusCondition = iota
	Burn
	PoisonReg
	Paralysis
	Sleep
	Freeze
	PoisonBad
)
