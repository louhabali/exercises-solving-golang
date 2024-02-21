package piscine

func RecursiveFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	if nb == 0 {
		return 1
	}
	c := nb * RecursiveFactorial(nb-1)
	if c < 0 {
		return 0
	}
	return c
}
