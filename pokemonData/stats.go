package pokemondata

func HPForumla(p Pokemon) int {
	dividend := (2 * p.BaseStats.HP) * p.Level
	division := dividend / 100
	final := division + p.Level + 10
	return final
}

func calculateStat(baseStat, level int) int {
	return ((2 * baseStat * level) / 100) + 5
}

func (p *Pokemon) CalculateStats() {
	p.Attack = calculateStat(p.BaseStats.Attack, p.Level)
	p.Defense = calculateStat(p.BaseStats.Defense, p.Level)
	p.SpAttack = calculateStat(p.BaseStats.SpAttack, p.Level)
	p.SpDefense = calculateStat(p.BaseStats.SpDefense, p.Level)
	p.Speed = calculateStat(p.BaseStats.Speed, p.Level)
}
