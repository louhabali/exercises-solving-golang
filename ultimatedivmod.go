package piscine

func UltimateDivMod(a *int, b *int) {
	division := *a / *b
	remainder := *a % *b
	*a = remainder
	*b = division
}
