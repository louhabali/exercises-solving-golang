package piscine

func LoafOfBread(str string) string {
	if len(str) < 5 {
		return "Invalid Output\n"
	}
	var result string
	count := 0
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			result += string(str[i])
			count++
			if count == 5 {
				i++
				count = 0
				if i < len(str)-1 {
					result += " "
				}

			}
		}
	}
	return result + "\n"
}
