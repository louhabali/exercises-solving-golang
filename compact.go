package piscine

func Compact(ptr *[]string) int {
	p := *ptr
	newarr := make([]string, 0)
	for i := 0; i < len(p); i++ {
		if string(p[i]) != "" {
			newarr = append(newarr, p[i])
		}
		*ptr = newarr
	}
	return len(newarr)
}
