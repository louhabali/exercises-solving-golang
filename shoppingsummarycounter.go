package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	summary := make(map[string]int)
	word := ""

	for _, char := range str {
		if char != ' ' {
			word += string(char)
		} else if word != "" {
			summary[word]++
			word = ""
		}
	}
	if word != "" {
		summary[word]++
	}

	return summary
}
