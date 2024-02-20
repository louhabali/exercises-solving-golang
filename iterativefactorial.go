package piscine

func IterativeFactorial(nb int) int {
		c := 1
	for i := 2; i <= nb; i++ {
		c *= i
	}
	return c
}
