package piscine

func Join(strs []string, sep string) string {
	newarr := ""
	for i, ch := range strs {
		if i < len(strs)-1 {
			newarr += ch + sep
		}
	}
	return newarr
}
