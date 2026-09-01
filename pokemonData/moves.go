package pokemondata

import (
	"fmt"
	battledata "pokemonBattleSim/battleSystem/battleData"
)

func CalculateStatusMove(user, target *Pokemon, move Move, bs *battledata.BattleState) {
	switch move.Name {
	case "Growth":
		Growth(user, *bs)
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
		WillOWisp(target, *bs)
	case "Poison Powder":
		PoisonPowder(target, *bs)
	case "Toxic":
		Toxic(target, *bs)
	case "Stun Spore":
		StunSpore(target, *bs)
	case "Sleep Powder":
		SleepPowder(target, *bs)
	case "Sunny Day":
		SunnyDay(bs)
	case "Rain Dance":
		RainDance(bs)
	case "Sandstorm":
		Sandstorm(bs)
	case "Snowscape":
		Snowscape(bs)
	default:
		fmt.Println("specified status move not implemented yet")
	}
}

func Snowscape(bs *battledata.BattleState) {
	switch bs.Conditions.Weather {
	case battledata.Snow:
		fmt.Println("But it failed!")
	case battledata.ExtremeHarshSunlight:
		fmt.Println("The Snowscape failed!")
		fmt.Println("The extremely harsh sunlight is so strong, it can't be blown away!")
	case battledata.HeavyRain:
		fmt.Println("The Snowscape failed!")
		fmt.Println("There is no relief from this heavy rain!")
	case battledata.StrongWinds:
		fmt.Println("The Snowscape failed!")
		fmt.Println("A mysterious air current is protecting Flying-type Pokémon!")
	default:
		bs.Conditions.WeatherTurns = 5
		bs.Conditions.Weather = battledata.Snow
		fmt.Println("It started to snow!")
	}
}

func Sandstorm(bs *battledata.BattleState) {
	switch bs.Conditions.Weather {
	case battledata.Sandstorm:
		fmt.Println("But it failed!")
	case battledata.ExtremeHarshSunlight:
		fmt.Println("The Sandstorm failed!")
		fmt.Println("The extremely harsh sunlight is so strong, it can't be blown away!")
	case battledata.HeavyRain:
		fmt.Println("The Sandstorm failed!")
		fmt.Println("There is no relief from this heavy rain!")
	case battledata.StrongWinds:
		fmt.Println("The Sandstrom failed!")
		fmt.Println("A mysterious air current is protecting Flying-type Pokémon!")
	default:
		bs.Conditions.WeatherTurns = 5
		bs.Conditions.Weather = battledata.Sandstorm
		fmt.Println("A sandstorm kicked up!")
	}
}

func RainDance(bs *battledata.BattleState) {
	switch bs.Conditions.Weather {
	case battledata.Rain:
		fmt.Println("But it failed!")
	case battledata.ExtremeHarshSunlight:
		fmt.Println("The Rain Dance failed!")
		fmt.Println("The extremely harsh sunlight is so strong, it can't be blown away!")
	case battledata.HeavyRain:
		fmt.Println("The Rain Dance failed!")
		fmt.Println("There is no relief from this heavy rain!")
	case battledata.StrongWinds:
		fmt.Println("The Rain Dance failed!")
		fmt.Println("A mysterious air current is protecting Flying-type Pokémon!")
	default:
		bs.Conditions.WeatherTurns = 5
		bs.Conditions.Weather = battledata.Rain
		fmt.Println("It started to rain!")
	}
}

func SunnyDay(bs *battledata.BattleState) {
	switch bs.Conditions.Weather {
	case battledata.HarshSunlight:
		fmt.Println("But it failed!")
	case battledata.ExtremeHarshSunlight:
		fmt.Println("The Sunny Day failed!")
		fmt.Println("The extremely harsh sunlight is so strong, it can't be blown away!")
	case battledata.HeavyRain:
		fmt.Println("The Sunny Day failed!")
		fmt.Println("There is no relief from this heavy rain!")
	case battledata.StrongWinds:
		fmt.Println("The Sunny Day failed!")
		fmt.Println("A mysterious air current is protecting Flying-type Pokémon!")
	default:
		bs.Conditions.WeatherTurns = 5
		bs.Conditions.Weather = battledata.HarshSunlight
		fmt.Println("The sunlight turned harsh!")
	}
}

func SleepPowder(p *Pokemon, bs battledata.BattleState) {
	if len(p.Types) == 2 {
		if p.Types[0] == Grass || p.Types[1] == Grass {
			fmt.Printf("It doesn't affect %v\n", p.Name)
			return
		}
	}
	if p.Types[0] == Grass {
		fmt.Printf("It doesn't affect %v\n", p.Name)
		return
	}

	result := p.ApplyStatus(Sleep, bs)
	if result == true {
		fmt.Printf("%v fell asleep!\n", p.Name)
	}
}

func StunSpore(p *Pokemon, bs battledata.BattleState) {
	if len(p.Types) == 2 {
		if p.Types[0] == Grass || p.Types[1] == Grass {
			fmt.Printf("It doesn't affect %v\n", p.Name)
			return
		}
	}
	if p.Types[0] == Grass {
		fmt.Printf("It doesn't affect %v\n", p.Name)
		return
	}

	result := p.ApplyStatus(Paralysis, bs)
	if result == true {
		fmt.Printf("%v is paralyzed, so it may be unable to move!\n", p.Name)
	}
}

func Toxic(p *Pokemon, bs battledata.BattleState) {
	result := p.ApplyStatus(PoisonBad, bs)
	if result == true {
		fmt.Printf("%v was badly poisoned!\n", p.Name)
	}
}

func PoisonPowder(p *Pokemon, bs battledata.BattleState) {
	result := p.ApplyStatus(PoisonReg, bs)
	if result == true {
		fmt.Printf("%v was poisoned!\n", p.Name)
	}
}

func WillOWisp(p *Pokemon, bs battledata.BattleState) {
	result := p.ApplyStatus(Burn, bs)
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
	if p.StatStages.Speed == -6 {
		fmt.Printf("%v Speed won't go any lower!\n", p.Name)
		return
	}
	p.ChangeStatStage(Speed, -2)
	fmt.Printf("%v Speed harshly fell!\n", p.Name)
}

func Agility(p *Pokemon) {
	if p.StatStages.Speed == 6 {
		fmt.Printf("%v Speed won't go any higher!\n", p.Name)
		return
	}
	p.ChangeStatStage(Speed, 2)
	fmt.Printf("%v Speed rose sharply!\n", p.Name)
}

func Amnesia(p *Pokemon) {
	if p.StatStages.SpDefense == 6 {
		fmt.Printf("%v Special Defense won't go any higher!\n", p.Name)
		return
	}
	p.ChangeStatStage(SpDefense, 2)
	fmt.Printf("%v Special Defense rose sharply!\n", p.Name)
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
	if p.StatStages.Defense == 6 {
		fmt.Printf("%v Defense won't go any higher!\n", p.Name)
		return
	}
	p.ChangeStatStage(Defense, 2)
	fmt.Printf("%v Defense rose sharply!\n", p.Name)
}

func Growth(p *Pokemon, bs battledata.BattleState) {
	amount := 1

	switch bs.Conditions.Weather {
	case battledata.HarshSunlight, battledata.ExtremeHarshSunlight:
		amount = 2
	}

	if p.StatStages.Attack == 6 {
		fmt.Printf("%v Attack won't go any higher!\n", p.Name)
	} else {
		p.ChangeStatStage(Attack, amount)
		fmt.Printf("%v Attack rose!\n", p.Name)
	}

	if p.StatStages.SpAttack == 6 {
		fmt.Printf("%v Special Attack won't go any higher!\n", p.Name)
	} else {
		p.ChangeStatStage(SpAttack, amount)
		fmt.Printf("%v Special Attack rose!\n", p.Name)
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
