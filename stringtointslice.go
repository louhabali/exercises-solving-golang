package piscine

func StringToIntSlice(str string) []int {
	intarr := make([]int, 0)
	for _, ch := range str {
		intarr = append(intarr, int(ch))
	}
	if len(intarr) == 0 {
		return nil
	}
	return intarr
}
