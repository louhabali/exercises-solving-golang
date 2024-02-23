package piscine

func MakeRange(min, max int) []int {
	if min >= max {
		return nil
	}
	sizeoffarr := max - min
	arr := make([]int, sizeoffarr)
	for i := 0; i < sizeoffarr; i++ {
		arr[i] = min + i
	}
	return arr
}
