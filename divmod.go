package piscine

func DivMod(a int, b int, div *int, mod *int) {
	division := a / b
	remainder := a % b
	*mod = remainder
	*div = division
}
