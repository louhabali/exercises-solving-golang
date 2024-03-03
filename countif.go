package piscine

func CountIf(f func(string) bool, tab []string) int {
	countif := 0
	for _, char := range tab {
		if f(char) == true {
			countif++
		}
	}
	return countif
}
