package piscine

func Unmatch(a []int) int {
	justone := 0
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a); j++ {
			if a[j] == a[i] {
				justone++
			}
		}
		if justone%2 != 0 {
			return a[i]
		}
	}

	return -1
}
