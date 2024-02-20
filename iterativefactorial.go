package piscine

func IterativeFactorial(nb int) int {
	c := int(1)
	for i := int(2); i <= nb; i++ {
		c = c * i
	}
	return c
}
