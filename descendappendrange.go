package piscine

func DescendAppendRange(max, min int) []int {
	if min > max {
		return nil
	}
	arr := make([]int, 0, max-min)
	for i := max; i > min; i-- {
		arr = append(arr, i)
	}
	return arr
}
