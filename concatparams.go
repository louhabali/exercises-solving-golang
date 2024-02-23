package piscine

func ConcatParams(args []string) string {
	strarray := make([]byte, 0)
	for _, arg := range args {
		strarray = append(strarray, arg...)
		strarray = append(strarray, '\n')
	}
	return string(strarray)
}
