package piscine

func IterativeFactorial(nb int) int {
	if nb < 0 {
		return 0
	}
	c := 1
	for i := 1; i <= nb; i++ {
		c = c * i
		if c < 0 {
			return 0
		}
	}
	return c
}

