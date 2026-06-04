package pokemondata

func HPForumla(p Pokemon) int {
	dividend := (2 * p.BaseStats.HP) * p.Level
	division := dividend / 100
	final := division + p.Level + 10
	return final
}
