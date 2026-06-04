package pokemondata

func Damage(attacker, defender Pokemon) int {
	return (((((2 * attacker.Level) / 5) + 2) * 50 * attacker.Attack / defender.Defense) / 50) + 2
}
