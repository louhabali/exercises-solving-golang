package piscine

func DescendAppendRange(max, min int) []int {
	arr := []int{}
	if min >= max {
		return arr
	}
	for i := max; i > min; i-- {
		arr = append(arr, i)
	}
	return arr
}
