package piscine

func ReverseMenuIndex(menu []string) []string {
	newarr := make([]string, 0)
	for i := len(menu) - 1; i >= 0; i-- {
		newarr = append(newarr, menu[i])
	}
	return newarr
}
