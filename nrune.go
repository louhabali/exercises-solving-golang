package piscine

func NRune(s string, n int) rune {
	arr := []rune(s)
	if n < 1 || n > len(arr) {
		return 0
	} else {
		return arr[n-1]
	}
}
