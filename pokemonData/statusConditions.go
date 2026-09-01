package pokemondata

import (
	"fmt"
	"math/rand/v2"
	battledata "pokemonBattleSim/battleSystem/battleData"
)

func (p *Pokemon) ApplyStatus(status StatusCondition, bs battledata.BattleState) bool {
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
		if bs.Conditions.Terrain == battledata.Elec {
			return false
		}
	case Freeze:
		if len(p.Types) == 2 {
			if p.Types[0] == Ice || p.Types[1] == Ice {
				fmt.Printf("It doesn't affect %v\n", p.Name)
				return false
			}
		}
		if p.Types[0] == Ice {
			fmt.Printf("It doesn't affect %v\n", p.Name)
			return false
		}
		if bs.Conditions.Weather == battledata.HarshSunlight || bs.Conditions.Weather == battledata.ExtremeHarshSunlight {
			return false
		}
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

	p.StatusTurns = 0
	p.Status = status

	return true
}

func ProcessStatus(p *Pokemon) {
	switch p.Status {
	case Burn:
		ApplyBurnDamage(p)
	case PoisonReg:
		ApplyRegPoisonDamage(p)
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
	case Sleep:
		p.StatusTurns++
		switch p.StatusTurns {
		case 1:
			fmt.Printf("%v is fast asleep!\n", p.Name)
			return false
		case 2:
			if rand.IntN(100) < 33 {
				p.Status = None
				fmt.Printf("%v woke up!\n", p.Name)
				return true
			} else {
				fmt.Printf("%v is fast asleep!\n", p.Name)
				return false
			}
		case 3:
			p.Status = None
			fmt.Printf("%v woke up!\n", p.Name)
			return true
		}
	case Freeze:
		p.StatusTurns++
		switch p.StatusTurns {
		case 1:
			if rand.IntN(100) < 25 {
				fmt.Printf("%v thawed out!\n", p.Name)
				return true
			} else {
				fmt.Printf("%v is frozen solid!\n", p.Name)
				return false
			}
		case 2:
			if rand.IntN(100) < 25 {
				fmt.Printf("%v thawed out!\n", p.Name)
				return true
			} else {
				fmt.Printf("%v is frozen solid!\n", p.Name)
				return false
			}
		case 3:
			fmt.Printf("%v thawed out!\n", p.Name)
			return true
		}
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
	Confusion
	Flinch
	Love
	Drowsy
)
