package pokemondata

import (
	"fmt"
	"math/rand/v2"
)

func (p *Pokemon) ApplyStatus(status StatusCondition) bool {
	if p.Status != None {
		fmt.Printf("It doesn't affect %v\n", p.Name)
		return false
	}

	switch status {
	case Burn:
		if len(p.Types) == 2 {
			if p.Types[0] == Fire || p.Types[1] == Fire {
				fmt.Printf("It doesn't affect %v\n", p.Name)
				return false
			}
		}
		if p.Types[0] == Fire {
			fmt.Printf("It doesn't affect %v\n", p.Name)
			return false
		}
	case PoisonReg:
		if len(p.Types) == 2 {
			if p.Types[0] == Poison || p.Types[1] == Poison || p.Types[0] == Steel || p.Types[1] == Steel {
				fmt.Printf("It doesn't affect %v\n", p.Name)
				return false
			}
		}
		if p.Types[0] == Poison || p.Types[0] == Steel {
			fmt.Printf("It doesn't affect %v\n", p.Name)
			return false
		}
	case Paralysis:
		if len(p.Types) == 2 {
			if p.Types[0] == Electric || p.Types[1] == Electric {
				fmt.Printf("It doesn't affect %v\n", p.Name)
				return false
			}
		}
		if p.Types[0] == Electric {
			fmt.Printf("It doesn't affect %v\n", p.Name)
			return false
		}
	case Sleep:

	case Freeze:

	case PoisonBad:
		if len(p.Types) == 2 {
			if p.Types[0] == Poison || p.Types[1] == Poison || p.Types[0] == Steel || p.Types[1] == Steel {
				fmt.Printf("It doesn't affect %v\n", p.Name)
				return false
			}
		}
		if p.Types[0] == Poison || p.Types[0] == Steel {
			fmt.Printf("It doesn't affect %v\n", p.Name)
			return false
		}
	}

	p.Status = status
	p.StatusTurns = 0
	return true
}

func ProcessStatus(p *Pokemon) {
	switch p.Status {
	case Burn:
		ApplyBurnDamage(p)
	case PoisonReg:
		ApplyRegPoisonDamage(p)
	case Sleep:

	case Freeze:

	case PoisonBad:
		p.StatusTurns++
		ApplyBadPoisonDamage(p)
	}
}

func CanMove(p *Pokemon) bool {
	switch p.Status {
	case Paralysis:
		if rand.IntN(1000) < 125 {
			fmt.Printf("%v couldn't move because it's paralyzed!\n", p.Name)
			return false
		}
	case Freeze:

	}

	return true
}

func ApplyBadPoisonDamage(p *Pokemon) {
	var dmg int
	switch p.StatusTurns {
	case 1:
		dmg = max(p.MaxHP/16, 1)
	case 2:
		dmg = max(p.MaxHP*2/16, 1)
	case 3:
		dmg = max(p.MaxHP*3/16, 1)
	case 4:
		dmg = max(p.MaxHP*4/16, 1)
	case 5:
		dmg = max(p.MaxHP*5/16, 1)
	case 6:
		dmg = max(p.MaxHP*6/16, 1)
	case 7:
		dmg = max(p.MaxHP*7/16, 1)
	case 8:
		dmg = max(p.MaxHP*8/16, 1)
	case 9:
		dmg = max(p.MaxHP*9/16, 1)
	case 10:
		dmg = max(p.MaxHP*10/16, 1)
	case 11:
		dmg = max(p.MaxHP*11/16, 1)
	case 12:
		dmg = max(p.MaxHP*12/16, 1)
	case 13:
		dmg = max(p.MaxHP*13/16, 1)
	case 14:
		dmg = max(p.MaxHP*14/16, 1)
	case 15:
		dmg = max(p.MaxHP*15/16, 1)
	default:
		dmg = max(p.MaxHP, 1)
	}

	p.HP -= dmg
	fmt.Printf("%v was hurt by its poison! %v took %v damage!\n", p.Name, p.Name, dmg)

	if p.HP <= 0 {
		p.HP = 0
		fmt.Printf("%v fainted!\n", p.Name)
	}
}

func ApplyRegPoisonDamage(p *Pokemon) {
	dmg := max(p.MaxHP/8, 1)

	p.HP -= dmg
	fmt.Printf("%v was hurt by its poison! %v took %v damage!\n", p.Name, p.Name, dmg)

	if p.HP <= 0 {
		p.HP = 0
		fmt.Printf("%v fainted!\n", p.Name)
	}
}

func ApplyBurnDamage(p *Pokemon) {
	// max returns the larger of the 2 variables
	dmg := max(p.MaxHP/16, 1)

	p.HP -= dmg
	fmt.Printf("%v was hurt by its burn! %v took %v damage!\n", p.Name, p.Name, dmg)

	if p.HP <= 0 {
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
