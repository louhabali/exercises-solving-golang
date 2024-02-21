package piscine

func IterativePower(nb int, power int) int {

	itp := 1
	for i := 1; i <= power; i++ {
		itp = nb * nb
	}
	return itp
}
