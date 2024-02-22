package piscine

func Capitalize(s string) string {
	words := []rune(s)

	for i := 0; i < len(words); i++ {
		if i == 0 || (words[i-1] < 'a' || words[i-1] > 'z') && (words[i-1] < 'A' || words[i-1] > 'Z') && (words[i-1] < '0' || words[i-1] > '9') {
			if words[i] >= 'a' && words[i] <= 'z' {
				words[i] -= 32
			}
		}
	}
	return string(words)
}