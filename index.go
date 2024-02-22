package piscine

func Index(s string, toFind string) int {
	farr := []rune(toFind)
	sarr := []rune(s)
	for i := 0; i <= len(s)-1; i++ {
		if sarr[i] == farr[0] {
			return i
		}
	}
	return -1
}
hh