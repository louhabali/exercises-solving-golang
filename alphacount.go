package piscine

func AlphaCount(s string) int {
	count := 0
	for x := 0; x < len(s); x++ {
		if (97 <= s[x] && s[x] <= 122) || (65 <= s[x] && s[x] <= 90) {
			count++
		}
	}
	return count
}
