package piscine

func Compact(ptr *[]string) []string {
	p := *ptr
	newarr := make([]string, 0)
	for i := 0; i < len(p); i++ {
		if p[i] != "" {
			newarr = append(newarr, p[i])
		}
		*ptr = newarr
	}
	return newarr
}
