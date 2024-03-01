package piscine

func Map(f func(int) bool, a []int) []bool {
	result := make([]bool, len(a))
	for i, nbr := range a {
		result[i] = f(nbr)
	}
	return result
}
