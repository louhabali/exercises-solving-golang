package piscine

func SplitWhiteSpaces(s string) []string {
	var t []string
	sarr := []rune(s)
	st := ""
	for i := 0; i < len(sarr); i++ {
		if sarr[i] != ' ' {
			st += string(sarr[i])
		} else {
			if st != "" && st != " " && i < len(sarr) && sarr[i] == ' ' {
				t = append(t, st)
			}
			st = ""
		}
	}
	if st == "" {
		return nil
	}
	return t
}
