package piscine

func Abort(a, b, c, d, e int) int {
	arr := []int{a, b, c, d, e}
	for i := 0; i < len(arr); i++ {
		for j := 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				c := arr[i]
				arr[i] = arr[j]
				arr[j] = c
			}
		}
	}
	return arr[len(arr)/2]
}
