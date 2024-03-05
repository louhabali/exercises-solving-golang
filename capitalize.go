package piscine

func Capitalize(s string) string {
	words := []rune(s)
	capitalize := true

	for i := 0; i < len(words); i++ {
		if IsLetter(words[i]) {
			if capitalize && isLower(words[i]) {
				words[i] -= 32
			} else if !capitalize && isUpper(words[i]) {
				words[i] += 32
			}

			capitalize = false

		} else {
			capitalize = true
		}
	}

	return string(words)
}

func IsLetter(r rune) bool {
	return (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9')
}

func isLower(r rune) bool {
	return (r >= 'a' && r <= 'z')
}

func isUpper(r rune) bool {
	return (r >= 'A' && r <= 'Z')
}
