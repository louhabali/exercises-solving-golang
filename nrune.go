package piscine

func NRune(s string, n int) rune {
	arr := []rune(s)
	if n < 0 || n > len(arr)-1 {
		return 0
	} else {
		return arr[n]
	}
}
