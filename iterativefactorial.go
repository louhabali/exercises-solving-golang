package piscine

func IterativeFactorial(nb int) int {
	if nb < 1 {
		return 0
	} else {
		c := 1
		for i := 2; i <= nb; i++ {
			c = c * i
		}
		return c
	}
}
