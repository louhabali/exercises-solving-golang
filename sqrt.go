package piscine

func Sqrt(nb int) int {
	if nb < 0 {
		return 0
	}
	nombre := 0
	for i := 1; i <= nb/2; i++ {
		if i*i == nb {
			nombre = i
		}
	}
	return nombre
}
