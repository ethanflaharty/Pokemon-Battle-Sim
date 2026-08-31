package battle

import (
	"fmt"
	"math/rand/v2"
	battledata "pokemonBattleSim/battleSystem/battleData"
	pokemondata "pokemonBattleSim/pokemonData"
)

type MoveResult struct {
	Damage        int
	Effectiveness float64
	Crit          bool
	Hit           bool
	STAB          bool
}

func calculateMove(attacker, defender *pokemondata.Pokemon, move pokemondata.Move, bs *battledata.BattleState) MoveResult {
	result := MoveResult{}

	// Check for a hit
	result.Hit = CheckAccuracy(attacker, defender, move)
	if !result.Hit {
		return result
	}

	if result.Hit {
		if defender.Status == pokemondata.Freeze {
			if move.Type == pokemondata.Fire {
				fmt.Printf("%v thawed out!\n", defender.Name)
				defender.ApplyStatus(pokemondata.None)
			}
		}
	}

	// Check Specific Circumstances
	CheckSpecificMoveCases(&move, *bs)

	// Check for a Crit
	if rand.Float64() < 0.0625 {
		result.Crit = true
	}

	// Check for STAB
	if len(attacker.Types) == 2 {
		if attacker.Types[0] == move.Type || attacker.Types[1] == move.Type {
			result.STAB = true
		}
	} else {
		if attacker.Types[0] == move.Type {
			result.STAB = true
		}
	}

	// Check for Weather Mult
	weatherMult := CheckWeatherMult(move, *bs)

	// Calculate Type Effectiveness
	result.Effectiveness = pokemondata.EffectivenessCalc(defender.Types, move.Type)

	// Calculate RNG Multiplier
	multiplier := float64(rand.IntN(16)+85) / 100.0

	// Damage Calculation
	switch move.Category {
	case pokemondata.Physical:
		if attacker.Status == pokemondata.Burn {
			result.Damage = int(float64(((((((2 * attacker.Level) / 5) + 2) * move.Power * attacker.BattleStats.Attack / defender.BattleStats.Defense) / 50) + 2)) * weatherMult * multiplier * result.Effectiveness * 0.5)
		}
		result.Damage = int(float64(((((((2 * attacker.Level) / 5) + 2) * move.Power * attacker.BattleStats.Attack / defender.BattleStats.Defense) / 50) + 2)) * weatherMult * multiplier * result.Effectiveness)
		CheckSecondaryEffect(attacker, defender, move)
	case pokemondata.Special:
		result.Damage = int(float64(((((((2 * attacker.Level) / 5) + 2) * move.Power * attacker.BattleStats.SpAttack / defender.BattleStats.SpDefense) / 50) + 2)) * weatherMult * multiplier * result.Effectiveness)
		CheckSecondaryEffect(attacker, defender, move)
	case pokemondata.Status:
		pokemondata.CalculateStatusMove(attacker, defender, move, bs)
		if result.Crit == true {
			result.Crit = false
		}
	default:
		panic("unknown move category")
	}

	// Apply STAB
	if result.STAB {
		result.Damage = int(float64(result.Damage) * 1.5)
	}

	// Apply Crit
	if result.Crit {
		result.Damage = int(float64(result.Damage) * 1.5)
	}

	return result
}
func CheckSpecificMoveCases(move *pokemondata.Move, bs battledata.BattleState) {
	switch move.Name {
	case "Weather Ball":
		switch bs.Conditions.Weather {
		case battledata.HarshSunlight:
			move.Type = pokemondata.Fire
			move.Power = move.Power * 2
		case battledata.ExtremeHarshSunlight:
			move.Type = pokemondata.Fire
			move.Power = move.Power * 2
		case battledata.Rain:
			move.Type = pokemondata.Water
			move.Power = move.Power * 2
		case battledata.HeavyRain:
			move.Type = pokemondata.Water
			move.Power = move.Power * 2
		case battledata.Sandstorm:
			move.Type = pokemondata.Rock
			move.Power = move.Power * 2
		case battledata.Snow:
			move.Type = pokemondata.Ice
			move.Power = move.Power * 2
		default:
			return
		}
	}
}

func CheckWeatherMult(move pokemondata.Move, bs battledata.BattleState) float64 {
	switch bs.Conditions.Weather {
	case battledata.HarshSunlight:
		switch move.Type {
		case pokemondata.Fire:
			return 1.5
		case pokemondata.Water:
			return 0.5
		}
	case battledata.Rain:
		switch move.Type {
		case pokemondata.Fire:
			return 0.5
		case pokemondata.Water:
			return 1.5
		}
	}
	return 1
}

func CheckSecondaryEffect(attacker, defender *pokemondata.Pokemon, move pokemondata.Move) {
	switch move.Name {
	case "Malignant Chain":
		if len(defender.Types) == 2 {
			if defender.Types[0] == pokemondata.Poison || defender.Types[1] == pokemondata.Poison || defender.Types[0] == pokemondata.Steel || defender.Types[1] == pokemondata.Steel {
				return
			}
		} else {
			if defender.Types[0] == pokemondata.Poison || defender.Types[0] == pokemondata.Steel {
				return
			}
		}

		if rand.IntN(100) < 50 {
			result := defender.ApplyStatus(pokemondata.PoisonBad)
			if result {
				fmt.Printf("%v was badly poisoned!\n", defender.Name)
			}
		}
	case "Sludge Bomb":
		if len(defender.Types) == 2 {
			if defender.Types[0] == pokemondata.Poison || defender.Types[1] == pokemondata.Poison || defender.Types[0] == pokemondata.Steel || defender.Types[1] == pokemondata.Steel {
				return
			}
		} else {
			if defender.Types[0] == pokemondata.Poison || defender.Types[0] == pokemondata.Steel {
				return
			}
		}

		if rand.IntN(100) < 30 {
			result := defender.ApplyStatus(pokemondata.PoisonReg)
			if result {
				fmt.Printf("%v was poisoned!\n", defender.Name)
			}
		}
	case "Discharge":
		if len(defender.Types) == 2 {
			if defender.Types[0] == pokemondata.Electric || defender.Types[1] == pokemondata.Electric || defender.Types[0] == pokemondata.Ground || defender.Types[1] == pokemondata.Ground {
				return
			}
		} else {
			if defender.Types[0] == pokemondata.Electric || defender.Types[0] == pokemondata.Ground {
				return
			}
		}

	case "Ice Beam":
		if len(defender.Types) == 2 {
			if defender.Types[0] == pokemondata.Ice || defender.Types[1] == pokemondata.Ice {
				return
			}
		} else {
			if defender.Types[0] == pokemondata.Ice {
				return
			}
		}

		if rand.IntN(100) < 10 {
			result := defender.ApplyStatus(pokemondata.Freeze)
			if result {
				fmt.Printf("%v was frozen solid!\n", defender.Name)
			}
		}
	case "Ember":
		if len(defender.Types) == 2 {
			if defender.Types[0] == pokemondata.Fire || defender.Types[1] == pokemondata.Fire {
				return
			}
		} else {
			if defender.Types[0] == pokemondata.Fire {
				return
			}
		}

		if rand.IntN(100) < 10 {
			result := defender.ApplyStatus(pokemondata.Burn)
			if result {
				fmt.Printf("%v was burned!\n", defender.Name)
			}
		}
	}
}

func CheckAccuracy(attacker, defender *pokemondata.Pokemon, move pokemondata.Move) bool {
	if move.NeverMisses {
		return true
	}

	if move.Name == "Toxic" && attacker.HasType(pokemondata.Poison) {
		return true
	}

	accStage := attacker.StatStages.Accuracy - defender.StatStages.Evasion
	if accStage < -6 {
		accStage = -6
	} else if accStage > 6 {
		accStage = 6
	}

	var effectiveAcc int
	var result bool
	switch accStage {
	case -6:
		effectiveAcc = move.Accuracy * 3 / 9
		result = rand.IntN(100) < effectiveAcc
		return result
	case -5:
		effectiveAcc = move.Accuracy * 3 / 8
		result = rand.IntN(100) < effectiveAcc
		return result
	case -4:
		effectiveAcc = move.Accuracy * 3 / 7
		result = rand.IntN(100) < effectiveAcc
		return result
	case -3:
		effectiveAcc = move.Accuracy * 3 / 6
		result = rand.IntN(100) < effectiveAcc
		return result
	case -2:
		effectiveAcc = move.Accuracy * 3 / 5
		result = rand.IntN(100) < effectiveAcc
		return result
	case -1:
		effectiveAcc = move.Accuracy * 3 / 4
		result = rand.IntN(100) < effectiveAcc
		return result
	case 0:
		effectiveAcc = move.Accuracy
		result = rand.IntN(100) < effectiveAcc
		return result
	case 1:
		effectiveAcc = move.Accuracy * 4 / 3
		result = rand.IntN(100) < effectiveAcc
		return result
	case 2:
		effectiveAcc = move.Accuracy * 5 / 3
		result = rand.IntN(100) < effectiveAcc
		return result
	case 3:
		effectiveAcc = move.Accuracy * 6 / 3
		result = rand.IntN(100) < effectiveAcc
		return result
	case 4:
		effectiveAcc = move.Accuracy * 7 / 3
		result = rand.IntN(100) < effectiveAcc
		return result
	case 5:
		effectiveAcc = move.Accuracy * 8 / 3
		result = rand.IntN(100) < effectiveAcc
		return result
	case 6:
		effectiveAcc = move.Accuracy * 9 / 3
		result = rand.IntN(100) < effectiveAcc
		return result
	}
	return result
}

func applyDamage(attacker, defender *pokemondata.Pokemon, move pokemondata.Move, bs *battledata.BattleState) {
	fmt.Printf("%v used %v!\n", attacker.Name, move.Name)

	result := calculateMove(attacker, defender, move, bs)

	if !result.Hit {
		fmt.Printf("%v's attack missed!\n", attacker.Name)
		return
	}

	if move.Category == pokemondata.Special || move.Category == pokemondata.Physical {
		if result.Crit {
			fmt.Println("A critical hit!")
		}
		if result.Effectiveness > 1 {
			fmt.Println("It's super effective!")
		} else if result.Effectiveness < 1 {
			fmt.Println("It's not very effective!")
		} else if result.Effectiveness == 0 {
			fmt.Printf("It doesn't effect %v!\n", defender.Name)
			return
		}

		defender.HP -= result.Damage

		if defender.HP <= 0 {
			defender.HP = 0
			fmt.Printf("%v took %v damage! ", defender.Name, result.Damage)
			fmt.Printf("%v has fainted!\n", defender.Name)
		} else {
			if result.Damage != 0 {
				fmt.Printf("%v took %v damage! ", defender.Name, result.Damage)
			}
			fmt.Printf("%v has %v HP left!\n", defender.Name, defender.HP)
		}
	}
}
