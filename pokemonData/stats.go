package pokemondata

import (
	"log"
	battledata "pokemonBattleSim/battleSystem/battleData"
)

func (p Pokemon) TotalEVs() int {
	return p.EVs.HP +
		p.EVs.Attack +
		p.EVs.Defense +
		p.EVs.SpAttack +
		p.EVs.SpDefense +
		p.EVs.Speed
}

func (p Pokemon) ValidateIVs() bool {
	if p.IVs.HP > 31 ||
		p.IVs.Attack > 31 ||
		p.IVs.Defense > 31 ||
		p.IVs.SpAttack > 31 ||
		p.IVs.SpDefense > 31 ||
		p.IVs.Speed > 31 {
		return false
	}

	if p.IVs.HP < 0 ||
		p.IVs.Attack < 0 ||
		p.IVs.Defense < 0 ||
		p.IVs.SpAttack < 0 ||
		p.IVs.SpDefense < 0 ||
		p.IVs.Speed < 0 {
		return false
	}

	return true
}

func (p Pokemon) ValidateEVs() bool {
	if p.EVs.HP > 252 ||
		p.EVs.Attack > 252 ||
		p.EVs.Defense > 252 ||
		p.EVs.SpAttack > 252 ||
		p.EVs.SpDefense > 252 ||
		p.EVs.Speed > 252 {
		return false
	}

	if p.EVs.HP < 0 ||
		p.EVs.Attack < 0 ||
		p.EVs.Defense < 0 ||
		p.EVs.SpAttack < 0 ||
		p.EVs.SpDefense < 0 ||
		p.EVs.Speed < 0 {
		return false
	}

	// Used 508 instead of 510 since the missing 2 can
	// never add a stat point to a Pokemon
	return p.TotalEVs() <= 508
}

func HPForumla(p Pokemon) int {
	dividend := (2*p.BaseStats.HP + p.IVs.HP + p.EVs.HP/4) * p.Level
	division := dividend / 100
	return division + p.Level + 10
}

func calculateStat(baseStat, level, iv, ev int, natureMult float64) int {
	stat := (((2*baseStat + iv + ev/4) * level) / 100) + 5

	return int(float64(stat) * natureMult)
}

func (p Pokemon) NatureMultiplier(stat Stat) float64 {
	switch p.Nature {
	case Lonely:
		switch stat {
		case Attack:
			return 1.1
		case Defense:
			return 0.9
		}
	case Adamant:
		switch stat {
		case Attack:
			return 1.1
		case SpAttack:
			return 0.9
		}
	case Naughty:
		switch stat {
		case Attack:
			return 1.1
		case SpDefense:
			return 0.9
		}
	case Brave:
		switch stat {
		case Attack:
			return 1.1
		case Speed:
			return 0.9
		}
	case Bold:
		switch stat {
		case Defense:
			return 1.1
		case Attack:
			return 0.9
		}
	case Impish:
		switch stat {
		case Defense:
			return 1.1
		case SpAttack:
			return 0.9
		}
	case Lax:
		switch stat {
		case Defense:
			return 1.1
		case SpDefense:
			return 0.9
		}
	case Relaxed:
		switch stat {
		case Defense:
			return 1.1
		case Speed:
			return 0.9
		}
	case Modest:
		switch stat {
		case SpAttack:
			return 1.1
		case Attack:
			return 0.9
		}
	case Mild:
		switch stat {
		case SpAttack:
			return 1.1
		case Defense:
			return 0.9
		}
	case Rash:
		switch stat {
		case SpAttack:
			return 1.1
		case SpDefense:
			return 0.9
		}
	case Quiet:
		switch stat {
		case SpAttack:
			return 1.1
		case Speed:
			return 0.9
		}
	case Calm:
		switch stat {
		case SpDefense:
			return 1.1
		case Attack:
			return 0.9
		}
	case Gentle:
		switch stat {
		case SpDefense:
			return 1.1
		case Defense:
			return 0.9
		}
	case Careful:
		switch stat {
		case SpDefense:
			return 1.1
		case SpAttack:
			return 0.9
		}
	case Sassy:
		switch stat {
		case SpDefense:
			return 1.1
		case Speed:
			return 0.9
		}
	case Timid:
		switch stat {
		case Speed:
			return 1.1
		case Attack:
			return 0.9
		}
	case Hasty:
		switch stat {
		case Speed:
			return 1.1
		case Defense:
			return 0.9
		}
	case Jolly:
		switch stat {
		case Speed:
			return 1.1
		case SpAttack:
			return 0.9
		}
	case Naive:
		switch stat {
		case Speed:
			return 1.1
		case SpDefense:
			return 0.9
		}
	}

	return 1
}

func (p *Pokemon) CalculateBattleStats() {
	p.Attack = calculateStat(p.BaseStats.Attack, p.Level, p.IVs.Attack, p.EVs.Attack, p.NatureMultiplier(Attack))
	p.Defense = calculateStat(p.BaseStats.Defense, p.Level, p.IVs.Defense, p.EVs.Defense, p.NatureMultiplier(Defense))
	p.SpAttack = calculateStat(p.BaseStats.SpAttack, p.Level, p.IVs.SpAttack, p.EVs.SpAttack, p.NatureMultiplier(SpAttack))
	p.SpDefense = calculateStat(p.BaseStats.SpDefense, p.Level, p.IVs.SpDefense, p.EVs.SpDefense, p.NatureMultiplier(SpDefense))
	p.Speed = calculateStat(p.BaseStats.Speed, p.Level, p.IVs.Speed, p.EVs.Speed, p.NatureMultiplier(Speed))
}

func (p *Pokemon) UpdateBattleStats(bs battledata.BattleState) {
	p.BattleStats.Attack = p.CalculateStatChange(
		p.Attack,
		p.StatStages.Attack,
	)

	p.BattleStats.Defense = p.CalculateStatChange(
		p.Defense,
		p.StatStages.Defense,
	)

	p.BattleStats.SpAttack = p.CalculateStatChange(
		p.SpAttack,
		p.StatStages.SpAttack,
	)

	p.BattleStats.SpDefense = p.CalculateStatChange(
		p.SpDefense,
		p.StatStages.SpDefense,
	)

	p.BattleStats.Speed = p.CalculateStatChange(
		p.Speed,
		p.StatStages.Speed,
	)
	if p.Status == Paralysis {
		p.BattleStats.Speed /= 2
	}

	switch bs.Conditions.Weather {
	case battledata.Sandstorm:
		if len(p.Types) == 2 {
			if p.Types[0] == Rock || p.Types[1] == Rock {
				p.BattleStats.SpDefense = int(float64(p.BattleStats.SpDefense) * 1.5)
			}
		} else {
			if p.Types[0] == Rock {
				p.BattleStats.SpDefense = int(float64(p.BattleStats.SpDefense) * 1.5)
			}
		}
	case battledata.Snow:
		if len(p.Types) == 2 {
			if p.Types[0] == Ice || p.Types[1] == Ice {
				p.BattleStats.Defense = int(float64(p.BattleStats.SpDefense) * 1.5)
			}
		} else {
			if p.Types[0] == Ice {
				p.BattleStats.Defense = int(float64(p.BattleStats.SpDefense) * 1.5)
			}
		}
	}
}

func (p *Pokemon) ChangeStatStage(stat Stat, change int) {
	switch stat {
	case Attack:
		p.StatStages.Attack += change
		if p.StatStages.Attack > 6 {
			p.StatStages.Attack = 6
		} else if p.StatStages.Attack < -6 {
			p.StatStages.Attack = -6
		}
	case Defense:
		p.StatStages.Defense += change
		if p.StatStages.Defense > 6 {
			p.StatStages.Defense = 6
		} else if p.StatStages.Defense < -6 {
			p.StatStages.Defense = -6
		}
	case SpAttack:
		p.StatStages.SpAttack += change
		if p.StatStages.SpAttack > 6 {
			p.StatStages.SpAttack = 6
		} else if p.StatStages.SpAttack < -6 {
			p.StatStages.SpAttack = -6
		}
	case SpDefense:
		p.StatStages.SpDefense += change
		if p.StatStages.SpDefense > 6 {
			p.StatStages.SpDefense = 6
		} else if p.StatStages.SpDefense < -6 {
			p.StatStages.SpDefense = -6
		}
	case Speed:
		p.StatStages.Speed += change
		if p.StatStages.Speed > 6 {
			p.StatStages.Speed = 6
		} else if p.StatStages.Speed < -6 {
			p.StatStages.Speed = -6
		}
	case Accuracy:
		p.StatStages.Accuracy += change
		if p.StatStages.Accuracy > 6 {
			p.StatStages.Accuracy = 6
		} else if p.StatStages.Accuracy < -6 {
			p.StatStages.Accuracy = -6
		}
	case Evasion:
		p.StatStages.Evasion += change
		if p.StatStages.Evasion > 6 {
			p.StatStages.Evasion = 6
		} else if p.StatStages.Evasion < -6 {
			p.StatStages.Evasion = -6
		}
	}
}

func (p *Pokemon) CalculateStatChange(stat, stage int) int {
	switch stage {
	case -6:
		return stat * 2 / 8
	case -5:
		return stat * 2 / 7
	case -4:
		return stat * 2 / 6
	case -3:
		return stat * 2 / 5
	case -2:
		return stat * 2 / 4
	case -1:
		return stat * 2 / 3
	case 0:
		return stat
	case 1:
		return stat * 3 / 2
	case 2:
		return stat * 4 / 2
	case 3:
		return stat * 5 / 2
	case 4:
		return stat * 6 / 2
	case 5:
		return stat * 7 / 2
	case 6:
		return stat * 8 / 2
	default:
		log.Fatalln("stat changes only go from -6 to 6")
		return stat
	}
}

type Stats struct {
	HP        int
	Attack    int
	Defense   int
	SpAttack  int
	SpDefense int
	Speed     int
}

type BattleStats struct {
	Attack    int
	Defense   int
	SpAttack  int
	SpDefense int
	Speed     int
}

type StatStages struct {
	Attack    int
	Defense   int
	SpAttack  int
	SpDefense int
	Speed     int
	Accuracy  int
	Evasion   int
}

type Stat int

const (
	Attack Stat = iota
	Defense
	SpAttack
	SpDefense
	Speed
	Accuracy
	Evasion
)
