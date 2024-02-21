package piscine

func IterativePower(nb int, power int) int {
	if power < 0 {
		return 0
	}
	itp := 1
	for i := 1; i <= power; i++ {
		itp *= nb
	}
	return itp
}
