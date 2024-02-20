package piscine

func IterativeFactorial(nb int) int {
	c := int(1)
	for i := 2; i <= nb; i++ {
		c = c * int(i)
	}
	return c
}
